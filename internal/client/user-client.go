package client

import (
	"context"
	"errors"
	"net/http"
)

type UserClient struct {
	baseURL string
}

func NewUserClient(baseURL string) *UserClient {
	return &UserClient{baseURL: baseURL}
}
func (c *UserClient) Check(ctx context.Context, token string) error {
	req, _ := http.NewRequestWithContext(ctx, "GET", c.baseURL+"/user", nil)
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("user not found")
	}
	return nil
}
