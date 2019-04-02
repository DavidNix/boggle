package server

import (
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if e := homeTmpl.Execute(w, nil); e != nil {
		panic(e)
	}
}

var homeTmpl = template.Must(template.New("home").Parse(`
<!doctype html>

<html lang="en">
<head>
  <meta charset="utf-8">
  <title>Boggle Solver</title>
  <style>
	input[type=text] {
		margin: 15px 15px 15px 0px;
		width: 7em;
		font-size: 3em;
		letter-spacing: 1em;
		text-transform: uppercase;
	}
	input[type=submit] {
		font-size: 1.5em;
	}
  </style>
</head>

<body>
	<h1>Boggle Solver</h1>
	<form action="/solve" method="post">
		Enter a 4 by 4 Boggle board. 4 letters on each line. Only letters A-Z.<br>
		<input type="text" name="row0"><br>
		<input type="text" name="row1"><br>
		<input type="text" name="row2"><br>
		<input type="text" name="row3"><br>
		<input type="submit" value="Submit">
	</form> 
</body>
</html>
`))
