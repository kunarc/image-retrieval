package es

import (
	"fmt"
	"os"

	"github.com/olivere/elastic/v7"
)

var ESClient *elastic.Client

func ConnectEs() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://image-retrieval-es:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetBasicAuth("", ""),
	)
	if err != nil {
		fmt.Println("Failed to connect to elasticsearch!")
		os.Exit(1)
	}
	fmt.Println("Connected to elasticsearch!")
	ESClient = client
}
