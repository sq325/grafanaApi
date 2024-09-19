package org

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/sq325/grafanaApi/pkg/common"
	"github.com/sq325/grafanaApi/pkg/httpclient"
)

const (
	allOrgs = "/api/orgs" // must use basic auth

	currentOrg = "/api/org" // can use api key
)

type API interface {
	GetAll() Orgs
	GetCurrent() Org
	Create(name string) error
}

type api struct {
	u            *url.URL
	token        string
	user, passwd string
	client       *http.Client
}

var _ API = (*api)(nil)

func NewApi(ep, token string, user, passwd string) API {

	u, err := common.Url(ep)
	if err != nil {
	}

	return &api{
		u:      u,
		token:  token,
		user:   user,
		passwd: passwd,
		client: httpclient.New(),
	}
}

func (a *api) GetAll() Orgs {
	if a.user == "" || a.passwd == "" {
		fmt.Println("must provide user and passwd")
		return nil
	}

	req, err := common.Request(http.MethodGet, a.u, allOrgs, "", a.user, a.passwd, nil, nil)
	if err != nil {
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

	var orgs Orgs
	if err := json.NewDecoder(resp.Body).Decode(&orgs); err != nil {
		slog.Error("decode response body to Orgs failed", "err", err, "url", req.URL)
		return nil
	}

	return orgs
}

func (a *api) GetCurrent() Org {
	req, err := common.Request(http.MethodGet, a.u, currentOrg, a.token, "", "", nil, nil)
	if err != nil {
	}

	resp, err := a.client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		var msg json.RawMessage
		if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
			slog.Error("client.Do Failed", "err", err, "msg", msg, "url", req.URL)
		} else {
			slog.Error("client.Do Failed", "err", err, "url", req.URL)
		}
		return Org{}
	}
	defer resp.Body.Close()

	var org Org
	if err := json.NewDecoder(resp.Body).Decode(&org); err != nil {
		slog.Error("decode response body to Org failed", "err", err, "url", req.URL)
		return Org{}
	}

	return org

}

func (a *api) Create(name string) error {
	st := struct {
		Name string `json:"name"`
	}{name}
	reqBodyBys, _ := json.Marshal(st)

	req, err := common.Request(http.MethodPost, a.u, allOrgs, "", a.user, a.passwd, reqBodyBys, nil)
	if err != nil {
	}

	resp, err := a.client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		var msg json.RawMessage
		if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
			slog.Error("client.Do Failed", "err", err, "msg", msg, "url", req.URL, "reqbody", string(reqBodyBys))
		} else {
			slog.Error("client.Do Failed", "err", err, "url", req.URL, "reqbody", string(reqBodyBys))
		}
		return fmt.Errorf("create org failed")
	}
	defer resp.Body.Close()

	fmt.Println("create org success")
	return nil
}
