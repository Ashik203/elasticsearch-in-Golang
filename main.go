package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch"
)

func indexBook() {

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://192.168.52.118:9200",
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	if es == nil {
		log.Fatalf("Elasticsearch client is nil!")
	}

	book := map[string]string{
		"title":       "golang book",
		"author":      "abcd",
		"description": "A deep dive into Elasticsearch",
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(book); err != nil {
		log.Fatalf("Error encoding book data: %s", err)
	}

	res, err := es.Index(
		"books", // Index name
		strings.NewReader(buf.String()),
		es.Index.WithContext(context.Background()),
		es.Index.WithRefresh("true"),
	)

	if err != nil {
		log.Fatalf("Error indexing book: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error response from Elasticsearch: %s", res.String())
	}

	fmt.Println("Book indexed successfully:", res.String())
}
func searchBooksByTitle(title string) {

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://192.168.52.118:9200",
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	query := fmt.Sprintf(`{
        "query": {
            "match": {
                "title": "%s"
            }
        }
    }`, title)

	res, err := es.Search(
		es.Search.WithIndex("books"),
		es.Search.WithBody(strings.NewReader(query)),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithContext(context.Background()),
	)

	if err != nil {
		log.Fatalf("Error searching for books: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error response from Elasticsearch: %s", res.String())
	}

	fmt.Println(res.String())
}

func main() {
	indexBook()

	searchBooksByTitle("golang book")
}
