package server

import (
	"github.com/DavidNix/boggle"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"
)

const timeout = 60 * time.Second

func ListenAndServe(port string) error {
	dict, err := boggle.NewDictionary()
	if err != nil {
		return errors.WithStack(err)
	}
	log.Println("Loaded dictionary")

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      mux(dict),
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}
	return srv.ListenAndServe()
}

func mux(dict boggle.Dictionary) chi.Router {
	r := chi.NewMux()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.StripSlashes)
	r.Use(middleware.DefaultCompress)

	r.Get("/", home)
	r.Post("/solve", solve(dict))

	return r
}
