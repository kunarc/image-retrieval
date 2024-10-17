package es

import (
	"context"
	"fmt"

	"github.com/olivere/elastic/v7"
)

func DocInsert(model EsModel) error {
	indexRespones, err := ESClient.Index().Index(model.Index()).BodyJson(model).Do(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(indexRespones)
	return nil
}
func DocInsertBatch(model []EsModel) error {
	if len(model) == 0 {
		return nil
	}
	bulk := ESClient.Bulk().Index(model[0].Index()).Refresh("true")
	for _, m := range model {
		bulk.Add(elastic.NewBulkIndexRequest().Doc(m))
	}
	bulkResponse, err := bulk.Do(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(bulkResponse.Succeeded())
	return nil
}
