package controllers

import (
	"fmt"
	"github.com/fvukojevic/bookstore_items-api/domains/items"
	"github.com/fvukojevic/bookstore_items-api/services"
	"github.com/fvukojevic/bookstore_oauth-go/oauth"
	"net/http"
)

type ItemsControlllerInterface interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type itemController struct {
}

func GetNewItemsController() ItemsControlllerInterface {
	return &itemController{}
}

func (controller itemController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		//TODO: Return error to the caller
		return
	}

	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}

	result, err := services.GetItemsService().Create(item)
	if err != nil {
		//TODO: Return error json to the caller
		return
	}

	fmt.Println(result)
	//TODO: Return created item with http status 201 Created
}

func (controller itemController) Get(w http.ResponseWriter, r *http.Request) {

}
