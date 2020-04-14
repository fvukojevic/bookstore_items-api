package app

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrls()

	src := &http.Server{
		Handler:      router,
		Addr:         "localhost:8080",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
	}

	if err := src.ListenAndServe(); err != nil {
		panic(err)
	}
}
