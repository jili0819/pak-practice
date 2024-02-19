package main

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"log"
)

type Document struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	indexName := "my_index"
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
	}
	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Fatal(err)
	}
	//es7.Indices.Create("my_index")

	//document := struct {
	//	Name string `json:"name"`
	//}{
	//	"go-elasticsearch",
	//}
	//data, _ := json.Marshal(document)
	//res, err := es7.Index("my_index", bytes.NewReader(data))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(res)
	//res, err := es.Info()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(res)
	//res, err := es.Indices.Create("my_index")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(res)
	//document := Document{
	//	Id:    1,
	//	Name:  "Foo",
	//	Price: 10,
	//}
	/*	document := struct {
			Id    int    `json:"id"`
			Name  string `json:"name"`
			Price int    `json:"price"`
		}{
			Id:    1,
			Name:  "Foo",
			Price: 10,
		}

		res, err := es.Index("my_index").
			Id("1").
			Request(document).
			Do(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res)*/

	res, err := es.Search().
		Index(indexName).
		Query(&types.Query{
			Match: map[string]types.MatchQuery{
				"name": {Query: "Foo"},
			},
		}).
		Do(context.Background())

	if err != nil {
		log.Fatalf("error runnning search query: %s", err)
	}

	if res.Hits.Total.Value == 1 {
		doc := Document{}
		err = json.Unmarshal(res.Hits.Hits[0].Source_, &doc)
		if err != nil {
			log.Fatalf("cannot unmarshal document: %s", err)
		}
		if doc.Name != "Foo" {
			log.Fatalf("unexpected search result")
		}
	} else {
		log.Fatalf("unexpected search result")
	}
}
