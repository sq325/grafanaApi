package datasource

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/sq325/grafanaApi/pkg/common"
	"github.com/sq325/grafanaApi/pkg/httpclient"
)

type API interface {
	GetAll() DataSources
	Create(name, datasourceType, url, access string, basicAuth bool) // token only
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

func (a *api) GetAll() DataSources {
	req, err := common.Request(http.MethodGet, a.u, apiAllDatasources, a.token, "", "", nil, nil)
	if err != nil {
		return nil
	}

	resp, err := a.client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		var msg json.RawMessage
		if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
			slog.Error("client.Do Failed", "err", err, "msg", msg, "url", req.URL)
		} else {
			slog.Error("client.Do Failed", "err", err, "url", req.URL)
		}
		return nil
	}
	defer resp.Body.Close()

	var datasources DataSources
	if err := json.NewDecoder(resp.Body).Decode(&datasources); err != nil {
		slog.Error("decode response body to datasources failed", "err", err, "url", req.URL)
		return nil
	}

	return datasources
}

// TODO
func (a *api) Create(name, datasourceType, url, access string, basicAuth bool) {
	datasource := &DataSource{
		Name:      name,
		Type:      datasourceType,
		Url:       url,
		Access:    DsAccess(access),
		BasicAuth: basicAuth,
	}
	body, err := json.Marshal(datasource)
	if err != nil {
		return
	}
	req, err := common.Request(http.MethodPost, a.u, apiAllDatasources, a.token, "", "", body, nil)
	if err != nil {
		return
	}

	resp, err := a.client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		var msg json.RawMessage
		if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
			slog.Error("client.Do Failed", "err", err, "msg", msg, "url", req.URL)
		} else {
			slog.Error("client.Do Failed", "err", err, "url", req.URL)
		}
		return
	}
	defer resp.Body.Close()
}
