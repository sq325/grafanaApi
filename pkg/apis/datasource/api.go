package datasource

import (
	"net/url"

	"github.com/sq325/grafanaApi/pkg/common"
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
	u     *url.URL
	token string
}

func NewApi(ep, token string) API {
	u, err := common.Url(ep)
	if err != nil {
	}

	return &api{
		u:     u,
		token: token,
	}
}

var _ API = (*api)(nil)

func (a *api) GetAll() *DataSources {
	return nil
}

func (a *api) Create() {
}
