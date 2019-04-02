# Boggle Solver

A boggle solver written in Go. 


Build the main package in `cmd` to run.

You may be able to see the site at: https://davidnix-boggle-solver.herokuapp.com/ (Heroku free dyno. YMMV)


# Known Limitations and TODOs

- Only English is supported but letter lookup supports non-ASCII characters.
- The server portion was quick and dirty given time limitations. Typically, I include unit tests for http handlers.
- For English, "Q" should expand to "QU".
- The models can accept a board of any size (like 5x5) even though currently it's locked to 4x4 on the web app.
- I am not a front end engineer, so yeah, I'm using JQuery.