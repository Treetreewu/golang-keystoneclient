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
	DefaultDomainID = "default"
	TokenHeader     = "X-Auth-Token"
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

func NewClient(config *openapi.Configuration) (*Client, error) {
	return NewClientWithServerIndex(config, 0)
}

func NewClientWithServerIndex(config *openapi.Configuration, serverIndex int) (*Client, error) {
	client := openapi.NewAPIClient(config)
	rawURL, _ := config.Servers.URL(serverIndex, nil)

	c := &Client{
		Req:       req.New(),
		APIClient: client,
	}
	c.Req.SetClient(&http.Client{})
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, errors.New("Invalid URL:")
	}
	u.Path = path.Join(u.Path, "v3")
	c.BaseURL = u
	c.Auth = NewAuthManager(c)
	return c, nil
}

func (c *Client) SetToken(token string) {
	c.Token = token
	c.APIClient.GetConfig().AddDefaultHeader(TokenHeader, token)
}

func (c *Client) Do(method, url string, vs ...interface{}) (resp *req.Resp, err error) {
	//Override function sending request, add Keystone token header.
	vs = append(vs, &req.Header{TokenHeader: c.Token})
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
	return c.Auth.CheckToken()
}
