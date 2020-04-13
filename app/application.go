package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrls()

	src := &http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}

	if err := src.ListenAndServe(); err != nil {
		panic(err)
	}
}
