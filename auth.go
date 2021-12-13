// Copyright 2017 EasyStack, Inc.

package keystone

import (
	"github.com/Treetreewu/golang-keystoneclient/model"

	"github.com/pkg/errors"

	"github.com/imroc/req"
)

const (
	MethodPassword = "password"
	MethodToken    = "token"
	DefaultDomain  = "default"
)

type AuthManager struct {
	BaseManager
}

func NewAuthManager(client *Client) *AuthManager {
	return &AuthManager{
		BaseManager{
			client:  client,
			URLPath: "auth",
		},
	}
}

func (m *AuthManager) Authenticate(scope interface{}, credential interface{}) (*req.Resp, error) {
	data := model.TokenRequestData{}
	switch cred := credential.(type) {
	case *model.PasswordCredential:
		data.Auth.Identity.Methods = []string{MethodPassword}
		data.Auth.Identity.Password = cred
	case model.PasswordCredential:
		data.Auth.Identity.Methods = []string{MethodPassword}
		data.Auth.Identity.Password = &cred
	case *model.TokenCredential:
		data.Auth.Identity.Methods = []string{MethodToken}
		data.Auth.Identity.Token = cred
	case model.TokenCredential:
		data.Auth.Identity.Methods = []string{MethodToken}
		data.Auth.Identity.Token = &cred
	default:
		return nil, errors.New("Invalid credential.")
	}
	data.Auth.Scope = scope
	r, err := m.Post("tokens", req.BodyJSON(&data))
	if err == nil {
		m.client.SetToken(r.Response().Header.Get("X-Subject-Token"))
	}
	return r, err
}

func (m *AuthManager) AlterScope(scope model.Scope) (*req.Resp, error) {
	if m.client.Token == "" {
		return nil, errors.New("Pleas login first.")
	}
	return m.Authenticate(scope, model.NewTokenCredential(m.client.Token))
}

func (m *AuthManager) CheckToken() error {
	_, err := m.Head("tokens")
	return err
}
