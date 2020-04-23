package controllers

import (
	"encoding/json"
	"github.com/fvukojevic/bookstore_items-api/domains/items"
	"github.com/fvukojevic/bookstore_items-api/services"
	"github.com/fvukojevic/bookstore_items-api/utils/http_utils"
	"github.com/fvukojevic/bookstore_oauth-go/oauth"
	"github.com/fvukojevic/bookstore_util-go/utils/errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strings"
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
		http_utils.RespondError(w, *err)
		return
	}
	sellerId := oauth.GetCallerId(r)
	if oauth.GetCallerId(r) == 0 {
		respErr := errors.NewUnauthorizeError("unable to retrieve user information from given access_token")
		http_utils.RespondError(w, *respErr)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := errors.NewBadRequestError("invalid item request body")
		http_utils.RespondError(w, *respErr)
		return
	}

	defer r.Body.Close()
	var itemRequest items.Item

	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := errors.NewBadRequestError("invalid item request body")
		http_utils.RespondError(w, *respErr)
		return
	}

	itemRequest.Seller = sellerId

	result, createErr := services.GetItemsService().Create(itemRequest)
	if createErr != nil {
		http_utils.RespondError(w, *createErr)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (controller itemController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	item, err := services.GetItemsService().Get(itemId)

	if err != nil {
		http_utils.RespondError(w, *err)
		return
	}

	http_utils.RespondJson(w, http.StatusOK, item)
}
