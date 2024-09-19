package org

import (
	"encoding/base64"
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

	req, err := http.NewRequest(http.MethodGet, a.u.JoinPath(allOrgs).String(), nil)
	if err != nil {
	}
	auth := a.user + ":" + a.passwd
	encodedAuth := base64.StdEncoding.EncodeToString([]byte(auth))
	req.Header.Add("Authorization", "Basic "+encodedAuth)

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
	req, err := http.NewRequest(http.MethodGet, a.u.JoinPath(currentOrg).String(), nil)
	if err != nil {
	}
	req.Header.Set("Authorization", "Bearer "+a.token)

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
