package items

import (
	"github.com/fvukojevic/bookstore_items-api/clients/elasticsearch"
	"github.com/fvukojevic/bookstore_util-go/utils/errors"
)

const(
	indexItems = "items"
)

func(i *Item) Save() *errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)
	if err != nil {
		return errors.NewInternalServerError("error when trying to save item")
	}

	i.Id = result.Id
	return nil
}
