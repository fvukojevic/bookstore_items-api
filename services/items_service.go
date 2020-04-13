package services

import (
	"github.com/fvukojevic/bookstore_items-api/domains/items"
	"github.com/fvukojevic/bookstore_util-go/utils/errors"
)

type ItemsServiceInterface interface {
	Create(item items.Item) (*items.Item, *errors.RestErr)
	Get(id string) (*items.Item, *errors.RestErr)
}

type itemService struct {
}

func GetItemsService() ItemsServiceInterface {
	return &itemService{}
}

func (service itemService) Create(item items.Item) (*items.Item, *errors.RestErr) {
	return nil, errors.NewBadRequestError("not implemented")
}

func (service itemService) Get(id string) (*items.Item, *errors.RestErr) {
	return nil, errors.NewBadRequestError("not implemented")
}