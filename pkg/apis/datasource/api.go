package datasource

import (
	"net/http"
	"net/url"

	"github.com/sq325/grafanaApi/pkg/common"
	"github.com/sq325/grafanaApi/pkg/httpclient"
)

type API interface {
	GetAll() *DataSources
	Create()
}

const (
	// GET RESPONSE.BODY:Datasources
	apiAllDatasources = "/api/datasources"
	// GET RESPONSE.BODY:Datasource
	apiDatasource = "/api/datasources/uid/{uid}"
)

type api struct {
	u      *url.URL
	token  string
	client *http.Client
}

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

var _ API = (*api)(nil)

func (a *api) GetAll() *DataSources {
	req, err := common.Request(http.MethodGet, a.u, apiAllDatasources, a.token, "", "", nil, nil)
	if err != nil {
	}

	resp, err := a.client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		// TODO
	}

	// TODO

	return nil
}

func (a *api) Create() {
}
