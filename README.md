# Boggle Solver

A boggle solver written in Go. 

You may be able to see the site at: https://davidnix-boggle-solver.herokuapp.com/ (Heroku free dyno. YMMV)

## To Run Locally

```
$ go build cmd/bogglehttpd/bogglehttpd.go

$ PORT=<port number of your choice> ./bogglehttpd
```


# Known Limitations and TODOs

- Only English is supported but letter lookup supports non-ASCII characters. It would be straightforward to extend the model layer to additional languages.
- The server portion was quick and dirty given time limitations. Typically, I include unit tests for http handlers.
- The models can accept a board of any size (like 5x5) even though currently it's locked to 4x4 on the web app.
- I am not a front end engineer.