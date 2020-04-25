package elasticsearch

import (
	"context"
	"fmt"
	"github.com/fvukojevic/bookstore_util-go/utils/logger"
	"github.com/olivere/elastic"
	"time"
)

const (
	envEsHost = "envEsHost"
)

var (
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	SetClient(client *elastic.Client)
	Index(index string, doc interface{}) (*elastic.IndexResponse, error)
	Get(index string, id string) (*elastic.GetResult, error)
	Search(index string, query elastic.Query) (*elastic.SearchResult, error)
}

type esClient struct {
	client *elastic.Client
}

func (c *esClient) SetClient(client *elastic.Client) {
	c.client = client
}

func Init() {
	log := logger.GetLogger()

	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log),
		elastic.SetInfoLog(log),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	Client.SetClient(client)
}

func (c *esClient) Index(index string, doc interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	res, err := c.client.Index().
		Index(index).
		BodyJson(doc).
		Do(ctx)

	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to index the document in index %s", index), err)
		return nil, err
	}

	return res, nil
}

func (c *esClient) Get(index string, id string) (*elastic.GetResult, error) {
	ctx := context.Background()

	result, err := c.client.Get().
		Index(index).
		Id(id).
		Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to get item by id %s", id), err)
		return nil, err
	}
	return result, nil
}

func (c *esClient) Search(index string, query elastic.Query) (*elastic.SearchResult, error) {
	ctx := context.Background()
	result, err := c.client.Search(index).
		Query(query).
		Do(ctx)
	if err != nil {
		logger.Error(fmt.Sprintf("error when trying to search documents in index %s", index), err)
		return nil, err
	}
	return result, nil
}