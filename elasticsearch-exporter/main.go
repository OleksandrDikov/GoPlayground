package main

import (
	"context"
	"os"

	"github.com/elastic/go-elasticsearch/v9"
)

func main() {

	cert, _ := os.ReadFile("./elastic_ca.crt")

	client, err := elasticsearch.NewTyped(
		elasticsearch.WithAddresses("https://192.168.98.28:9200"),
		elasticsearch.WithBasicAuth("elastic", os.Getenv("ELASTIC_PASSWORD")),
		elasticsearch.WithCACert(cert),
	)
	if err != nil {
		panic(err)
	}

	document := struct {
		Name string `json:"name"`
	}{
		"go-elasticsearch",
	}
	client.Index("logstash").
		Id("1").
		Document(document).
		Do(context.TODO())
}
