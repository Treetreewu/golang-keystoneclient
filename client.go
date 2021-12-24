// Copyright 2017 EasyStack, Inc.

package keystone

import (
	"github.com/Treetreewu/golang-keystoneclient/openapi"
	"net/http"
	"net/url"
	"path"

	"github.com/imroc/req"
	"github.com/pkg/errors"
)

const (
	TokenHeader = "X-Auth-Token"
)

type Client struct {
	*req.Req
	Token   string
	BaseURL *url.URL
	// Managers
	Auth *AuthManager
	// OpenAPI
	*openapi.APIClient
}

func NewClientByToken(config *openapi.Configuration, token string) (*Client, error) {
	client, err := NewClient(config)
	if err != nil {
		return nil, err
	}
	client.SetToken(token)
	return client, nil
}

func NewDefaultClientByToken(token string) (*Client, error) {
	return NewClientByToken(openapi.NewConfiguration(), token)
}

func NewDefaultClient() (*Client, error) {
	return NewClient(openapi.NewConfiguration())
}

func NewClient(config *openapi.Configuration) (*Client, error) {
	rawURL, _ := config.Servers.URL(0, nil)

	client := &Client{
		Req:       req.New(),
		APIClient: openapi.NewAPIClient(config),
	}
	client.Req.SetClient(&http.Client{})
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, errors.New("Invalid URL:")
	}
	u.Path = path.Join(u.Path, "v3")
	client.BaseURL = u

	client.Auth = NewAuthManager(client)
	return client, nil
}

func (c *Client) SetToken(token string) {
	c.Token = token
	c.APIClient.GetConfig().AddDefaultHeader(TokenHeader, token)
}

func (c *Client) Do(method, url string, vs ...interface{}) (resp *req.Resp, err error) {
	vs = append(
		vs,
		// Mapping OpenAPI client configurations.
		req.Header{"User-Agent": c.APIClient.GetConfig().UserAgent},
		c.APIClient.GetConfig().DefaultHeader,
		// Add Keystone token header.
		req.Header{TokenHeader: c.Token},
	)

	resp, err = c.Req.Do(method, url, vs...)
	if err == nil && resp.Response().StatusCode >= 400 {
		err = errors.New("Request failed")
	}
	return resp, errors.WithStack(err)
}

func (c *Client) AuthRequired() error {
	if c.Token == "" {
		return errors.New("Not authorized.")
	}
	_, err := c.Auth.TokenDetail()
	return err
}
