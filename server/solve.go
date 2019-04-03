package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"sort"
	"strings"

	"github.com/DavidNix/boggle"
)

var re = regexp.MustCompile(`[^a-zA-Z]+`)

const boardSize = 4

type entryPresenter struct {
	boggle.Entry
}

func (p entryPresenter) PathJSArray() string {
	path := p.Entry.Path
	arr := make([]string, len(path))
	for i := range path {
		arr[i] = path[i].String()
	}
	d, err := json.Marshal(arr)
	if err != nil {
		return err.Error()
	}
	return string(d)
}

func solve(dict boggle.Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if e := r.ParseForm(); e != nil {
			panic(e)
		}
		board := make(boggle.Board, boardSize)
		for i := 0; i < boardSize; i++ {
			text := r.Form.Get(fmt.Sprintf("row%d", i))
			text = re.ReplaceAllString(text, "")
			board[i] = strings.Split(text, "")
		}

		for i := range board {
			if len(board[i]) != boardSize {
				mssg := fmt.Sprintf("Row %d is wrong size, should be %d letters from A-Z. Please try again.", i+1, boardSize)
				http.Error(w, mssg, http.StatusUnprocessableEntity)
				return
			}
		}

		visitor := boggle.NewVisitor(dict)
		board.Traverse(visitor)
		renderResults(w, board, visitor.Found)
	}
}

func renderResults(w http.ResponseWriter, board boggle.Board, entries []boggle.Entry) {
	sort.Slice(entries, func(i, j int) bool {
		return len(entries[i].Word) > len(entries[j].Word)
	})
	ep := make([]entryPresenter, len(entries))
	for i := range entries {
		ep[i] = entryPresenter{entries[i]}
	}
	data := struct {
		Board   boggle.Board
		Entries []entryPresenter
	}{Board: board, Entries: ep}
	if e := resultsTmpl.Execute(w, data); e != nil {
		panic(e)
	}
}

var resultsTmpl = template.Must(template.New("results").Parse(`
<!doctype html>

<html lang="en">
<head>
  <meta charset="utf-8">
  <title>Boggle Solver</title>
  <style>
	body {
		margin: 30px;
	}
	#board {
		display: grid;
		grid-template-columns: 2em 2em 2em 2em;
		font-size: 3em;
		margin-bottom: 15px;
		text-transform: uppercase;
	}
	#results {
		display: grid;
		font-size: 1.5em;
		grid-template-columns: repeat(4, 1fr);
		grid-gap: 15px;
		width: 30%;
		text-transform: lowercase;
	}
	#results a {
		color: blue;
		display: block;
	}
	#results a:visited {
		color: blue;
	}
  </style>
</head>

<body>
	<h1>Boggle Results</h1>
	<div id="board">
	{{range $row, $letters := .Board}}
		{{range $col, $letter := $letters}}
		<span id="{{$row}}-{{$col}}" class="letter">{{$letter}}</span>
		{{end}}
	{{end}}
	</div>
	<p>Click a word to highlight it on the board.</p>
	<div id="results">
	{{range .Entries}}
		<a href="#" data-path='{{ .PathJSArray }}'>{{ .Word }}</a>
	{{end}}
	</div>

	<script src="https://code.jquery.com/jquery-3.3.1.slim.min.js"
	integrity="sha256-3edrmyuQ0w65f8gfBsqowzjJe2iM6n0nKciPUp8y+7E="
	crossorigin="anonymous"></script>

	<script>
	$('#results a').on('click', function() {
		$('#board span').css("color", "black");
		path = $(this).data('path');
		path.forEach(function (coord, i){
			$("#"+coord).css("color", "red");
		});
	});
	</script>

</body>
</html>
`))
