package elk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

// ElasticClient wraps the elasticsearch client and index name
type ElasticClient struct {
	es    *elasticsearch.Client
	index string
}

// NewElasticClient initializes an elasticsearch connection
func NewElasticClient(url, index string) (*ElasticClient, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{url},
		// Username:  os.Getenv("PROJECT_LOGGER_ELASTIC_USER"),
		// Password:  os.Getenv("PROJECT_LOGGER_ELASTIC_PASS"),
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create elasticsearch client: %w", err)
	}

	// Check connection
	_, err = es.Info()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to elasticsearch: %w", err)
	}

	log.Printf("connected to elasticsearch at %s", url)

	return &ElasticClient{
		es:    es,
		index: index,
	}, nil
}

// IndexLog sends a log entry to elasticsearch
func (ec *ElasticClient) IndexLog(doc any) error {
	data, err := json.Marshal(doc)
	if err != nil {
		return fmt.Errorf("failed to marshal log entry: %w", err)
	}

	req := bytes.NewReader(data)
	res, err := ec.es.Index(
		ec.index,
		req,
		ec.es.Index.WithContext(context.Background()),
		ec.es.Index.WithRefresh("true"),
	)
	if err != nil {
		return fmt.Errorf("failed to index log: %w", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error indexing log: %s", res.String())
	}

	return nil
}
