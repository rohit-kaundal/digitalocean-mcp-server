package client

import (
	"context"
	"fmt"
	"os"

	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

type DOClient struct {
	client *godo.Client
}

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func NewDOClient() (*DOClient, error) {
	token := os.Getenv("DIGITALOCEAN_ACCESS_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("DIGITALOCEAN_ACCESS_TOKEN environment variable is required")
	}

	tokenSource := &TokenSource{
		AccessToken: token,
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := godo.NewClient(oauthClient)

	return &DOClient{
		client: client,
	}, nil
}

func (d *DOClient) GetClient() *godo.Client {
	return d.client
}

func (d *DOClient) TestConnection() error {
	_, _, err := d.client.Account.Get(context.Background())
	return err
}