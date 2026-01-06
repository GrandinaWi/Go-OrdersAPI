package client

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type CatalogClient struct {
	baseURL string
}

func NewCatalogClient(baseURL string) *CatalogClient {
	return &CatalogClient{baseURL: baseURL}
}
func (c *CatalogClient) CheckProduct(ctx context.Context, productId int64) error {
	url := fmt.Sprintf("%s/products/%d", c.baseURL, productId)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("product not found")
	}
	return nil
}
