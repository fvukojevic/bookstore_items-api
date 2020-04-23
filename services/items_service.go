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
	if err := item.Save(); err != nil {
		return nil, err
	}

	return &item, nil
}

func (service itemService) Get(id string) (*items.Item, *errors.RestErr) {
	item := items.Item{Id: id}

	if err := item.Get(); err != nil {
		return nil, err
	}

	return &item, nil
}
