package server

import (
	"fmt"
	"github.com/DavidNix/boggle"
	"html/template"
	"net/http"
	"sort"
	"strings"
)

const boardSize = 4

func solve(dict boggle.Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if e := r.ParseForm(); e != nil {
			panic(e)
		}
		board := make(boggle.Board, boardSize)
		for i := 0; i < boardSize; i++ {
			text := r.Form.Get(fmt.Sprintf("row%d", i))
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

func renderResults(w http.ResponseWriter, board boggle.Board, results []boggle.Entry) {
	sort.Slice(results, func(i, j int) bool {
		return len(results[i].Word) > len(results[j].Word)
	})
	data := struct{
		Entries []boggle.Entry
	} {Entries: results}
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
	a {
		margin: 5px;
		display: block;	
	}
  </style>
</head>

<body>
	<h1>Boggle Results</h1>
	<div>Board is Here</div>
	{{ range .Entries }}
		<a href="#">{{ .Word }}</a>
	{{end}}

	<script src="https://code.jquery.com/jquery-3.3.1.slim.min.js"
	integrity="sha256-3edrmyuQ0w65f8gfBsqowzjJe2iM6n0nKciPUp8y+7E="
	crossorigin="anonymous"></script>

</body>
</html>
`))
