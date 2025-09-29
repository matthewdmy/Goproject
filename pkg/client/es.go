package client

import (
	"fmt"

	"github.com/olivere/elastic/v7"
)

func NewEsClient() (*elastic.Client, error) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
	)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return client, nil
}
