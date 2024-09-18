package org

import (
	"net/http"
	"net/url"

	"github.com/sq325/grafanaApi/pkg/common"
	"github.com/sq325/grafanaApi/pkg/httpclient"
)

const (
	// must use basic auth
	allOrgs    = "/api/orgs"
	currentOrg = "/api/org" // can use api key
)

type API interface {
	GetAll() Orgs
	GetCurrent() Org
}

type api struct {
	u      *url.URL
	token  string
	client *http.Client
}

var _ API = (*api)(nil)

func NewApi(ep, token string) API {

	u, err := common.Url(ep)
	if err != nil {
	}

	return &api{
		u:      u,
		token:  token,
		client: httpclient.New(),
	}
}

func (a *api) GetAll() Orgs {
	return nil
}

func (a *api) GetCurrent() Org {
	return Org{}
}
