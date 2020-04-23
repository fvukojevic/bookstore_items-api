package app

import (
	"github.com/fvukojevic/bookstore_items-api/controllers"
	"net/http"
)

func mapUrls() {
	router.HandleFunc("/items", controllers.GetNewItemsController().Create).Methods(http.MethodPost)
	router.HandleFunc("/items/{id}", controllers.GetNewItemsController().Get).Methods(http.MethodGet)
}
