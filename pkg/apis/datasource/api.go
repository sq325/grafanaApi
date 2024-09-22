package datasource

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"net/url"
	"strings"

	"github.com/sq325/grafanaApi/pkg/common"
	"github.com/sq325/grafanaApi/pkg/httpclient"
)

type API interface {
	GetAll() DataSources
	Get(uid string) DataSource
	// Create(name, datasourceType, url, access string, basicAuth bool) // token only
	Create(DataSource) error
	CreateFromArgs(name, datasourceType, url, access string, basicAuth bool) error
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

// ep: ip:port
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

func (a *api) CreateFromArgs(name, datasourceType, url, access string, basicAuth bool) error {
	datasource := &DataSource{
		Name:      name,
		Type:      datasourceType,
		Url:       url,
		Access:    DsAccess(access),
		BasicAuth: basicAuth,
	}
	body, err := json.Marshal(datasource)
	if err != nil {
		return err
	}
	req, err := common.Request(http.MethodPost, a.u, apiAllDatasources, a.token, "", "", body, nil)
	if err != nil {
		return err
	}

	resp, err := a.client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		var msg json.RawMessage
		if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
			slog.Error("client.Do Failed", "err", err, "msg", msg, "url", req.URL)
		} else {
			slog.Error("client.Do Failed", "err", err, "url", req.URL)
		}
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (a *api) Create(ds DataSource) error {
	body, err := json.Marshal(ds)
	if err != nil {
		return err
	}
	req, err := common.Request(http.MethodPost, a.u, apiAllDatasources, a.token, "", "", body, nil)
	if err != nil {
		return err
	}

	resp, err := a.client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		var msg json.RawMessage
		if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
			slog.Error("client.Do Failed", "err", err, "msg", msg, "url", req.URL)
		} else {
			slog.Error("client.Do Failed", "err", err, "url", req.URL)
		}
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (a *api) Get(uid string) DataSource {
	req, err := common.Request(http.MethodGet, a.u, strings.Replace(apiDatasource, "{uid}", uid, 1), a.token, "", "", nil, nil)
	if err != nil {
		slog.Error("new request failed", "err", err)
		return DataSource{}
	}

	resp, err := a.client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		var msg json.RawMessage
		if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
			slog.Error("client.Do Failed", "err", err, "msg", msg, "url", req.URL)
		} else {
			slog.Error("client.Do Failed", "err", err, "url", req.URL)
		}
		return DataSource{}
	}
	defer resp.Body.Close()

	var datasource DataSource
	if err := json.NewDecoder(resp.Body).Decode(&datasource); err != nil {
		slog.Error("decode response body to datasource failed", "err", err, "url", req.URL)
		return DataSource{}
	}

	return datasource
}
