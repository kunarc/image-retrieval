package es

import (
	"context"
	"fmt"
)

type EsModel interface {
	Index() string
	Mapping() string
}

func CreateIndex(model EsModel) error {
	if ExistsIndex(model.Index()) {
		fmt.Println("索引已存在")
		return nil
	}
	_, err := ESClient.CreateIndex(model.Index()).BodyString(model.Mapping()).Do(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println("创建索引成功!")
	return nil
}
func ExistsIndex(index string) bool {
	exists, _ := ESClient.IndexExists(index).Do(context.Background())
	return exists
}
