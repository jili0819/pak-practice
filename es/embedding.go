package main

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
	"log"
)

func main() {
	indexName := "os_test_pipeline_resource"
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://es-cn-pvh45ggap000h3ju6.elasticsearch.aliyuncs.com:9200",
		},
		Username: "elastic",
		Password: "Buzhongyao123",
	}
	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	embeddings := make([]float64, 0)
	for i := 1; i <= 1024; i++ {
		embeddings = append(embeddings, 0.1)
	}

	_, err = es.Update(indexName, "YPl5rJUBjP9y-Av4Jpgs").Doc(map[string]interface{}{
		"content": "test02",
		"chunk": []map[string]interface{}{
			{
				"content": "test02.embedding01",
				"meta": map[string]interface{}{
					"id":        "1",
					"parent_id": "0",
				},
				"embedding": embeddings,
				"sparse_embedding": map[string]float64{
					"1": 0.1,
					"2": 0.2,
				},
			},
		},
	}).Refresh(refresh.True).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
