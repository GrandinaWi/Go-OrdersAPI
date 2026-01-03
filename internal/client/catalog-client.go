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
func (c *CatalogClient) CheckProduct(ctx context.Context, productId string) error {
	resp, err := http.Get(fmt.Sprintf("%s/products/%s", c.baseURL, productId))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("product not found")
	}
	return nil
}
