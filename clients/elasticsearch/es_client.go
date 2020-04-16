package elasticsearch

import (
	"context"
	"github.com/olivere/elastic"
	"time"
)

const(
	envEsHost = "envEsHost"
)

var(
	Client esClientInterface = &esClient{}
)

type esClientInterface interface {
	SetClient(client *elastic.Client)
	Index(interface{}) (*elastic.IndexResponse, error)
}

type esClient struct {
	client *elastic.Client
}

func Init() {
	//log := logger.GetLogger()

	client, err := elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetHealthcheckInterval(10*time.Second),
		//elastic.SetErrorLog(log),
		//elastic.SetInfoLog(log),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	Client.SetClient(client)
}

func(c *esClient) Index(interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	return c.client.Index().Do(ctx)
}

func(c *esClient) SetClient(client *elastic.Client) {
	c.client = client
}