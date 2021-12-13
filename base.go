// Copyright 2017 EasyStack, Inc.

package keystone

import (
	"path"

	"github.com/imroc/req"
)

type BaseManager struct {
	client  *Client
	URLPath string
}

func (m *BaseManager) buildURL(subPath string) string {
	url := *m.client.BaseURL
	url.Path = path.Join(url.Path, m.URLPath, subPath)
	return url.String()
}

func (m *BaseManager) Do(method, subPath string, vs ...interface{}) (resp *req.Resp, err error) {
	url := m.buildURL(subPath)
	return m.client.Do(method, url, vs...)
}

func (m *BaseManager) Post(subPath string, vs ...interface{}) (*req.Resp, error) {
	return m.Do("POST", subPath, vs...)
}
func (m *BaseManager) Get(subPath string, vs ...interface{}) (*req.Resp, error) {
	return m.Do("GET", subPath, vs...)
}
func (m *BaseManager) Delete(subPath string, vs ...interface{}) (*req.Resp, error) {
	return m.Do("DELETE", subPath, vs...)
}
func (m *BaseManager) Patch(subPath string, vs ...interface{}) (*req.Resp, error) {
	return m.Do("PATCH", subPath, vs...)
}
func (m *BaseManager) Put(subPath string, vs ...interface{}) (*req.Resp, error) {
	return m.Do("PUT", subPath, vs...)
}
func (m *BaseManager) Head(subPath string, vs ...interface{}) (*req.Resp, error) {
	return m.Do("HEAD", subPath, vs...)
}
