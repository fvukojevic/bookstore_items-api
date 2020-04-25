package items

import (
	"encoding/json"
	"fmt"
	"github.com/fvukojevic/bookstore_items-api/clients/elasticsearch"
	"github.com/fvukojevic/bookstore_items-api/domains/queries"
	"github.com/fvukojevic/bookstore_util-go/utils/errors"
	"strings"
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

func(i *Item) Get() *errors.RestErr {
	itemId := i.Id
	result, err := elasticsearch.Client.Get(indexItems, i.Id)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			return errors.NewNotFoundError(fmt.Sprintf("item with id %s not found", i.Id))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get item by id %s", i.Id))
	}

	bytes, _ := result.Source.MarshalJSON()
	if err := json.Unmarshal(bytes, i); err != nil{
		return errors.NewInternalServerError("error when trying to parse item from bytes")
	}
	i.Id = itemId
	return nil
}

func (i *Item) Search(query queries.EsQuery) ([]Item, *errors.RestErr) {
	result, err := elasticsearch.Client.Search(indexItems, query.Build())
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("database error"))
	}

	items := make([]Item, result.TotalHits())
	for index, hit := range result.Hits.Hits {
		bytes, _ := hit.Source.MarshalJSON()
		var item Item
		if err := json.Unmarshal(bytes, &item); err != nil {
			return nil, errors.NewInternalServerError("database error")
		}
		item.Id = hit.Id
		items[index] = item
	}

	if len(items) == 0 {
		return nil, errors.NewNotFoundError("no items found matching given criteria")
	}
	return items, nil
}
