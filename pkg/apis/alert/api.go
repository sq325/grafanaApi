package alert

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/sq325/grafanaApi/pkg/common"
	"github.com/sq325/grafanaApi/pkg/httpclient"
)

// crud
type API interface {
	// provenance Header:X-Disable-Provenance
	Create(alert ProvisionedAlertRule, provenance bool) error
	GetAll() ProvisionedAlertRules
	Get(uids ...string) ProvisionedAlertRules
	Update(uid string, alert ProvisionedAlertRule) error
	Delete(uids ...string) error
}

// Grafana Alerting Provisioning Http Api
const (
	// GET RESPONSE.BODY:ProvisionedAlertRules
	// POST REQUEST.BODY:ProvisionedAlertRule Header:X-Disable-Provenance:true
	apiAllAlertRules       = "/api/v1/provisioning/alert-rules"
	apiAllAlertRulesExport = "/api/v1/provisioning/alert-rules/export" // GET:AlertingFileExport

	// GET RESPONSE.BODY:ProvisionedAlertRule
	// PUT Header:X-Disable-Provenance:true RESPONSE.BODY:ProvisionedAlertRule
	// DELETE Header:X-Disable-Provenance:true
	apiAlertRule = "/api/v1/provisioning/alert-rules/{uid}" // GET:ProvisionedAlertRule PUT:ProvisionedAlertRule DELETE
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
		slog.Error("url parse failed", "err", err, "ep", ep, "token", token)
		return nil
	}

	return &api{
		u:      u,
		token:  token,
		client: httpclient.New(),
	}
}

var _ API = (*api)(nil)

func (a *api) Create(alert ProvisionedAlertRule, Provenance bool) error {
	body, err := json.Marshal(alert)
	if err != nil {
		slog.Error("marshal alert failed", "err", err)
		return err
	}
	req, err := common.Request(http.MethodPost, a.u, apiAllAlertRules, a.token, "", "", body, nil)
	if err != nil {
		slog.Error("new request failed", "err", err)
		return err
	}
	if Provenance {
		req.Header.Set("X-Disable-Provenance", "true")
	}

	resp, err := a.client.Do(req)
	if err != nil || resp.StatusCode >= 400 {
		var msg json.RawMessage
		alertBys, _ := json.Marshal(alert)
		if decodeErr := json.NewDecoder(resp.Body).Decode(&msg); decodeErr != nil {
			slog.Error("client.Do Failed", "err", err, "url", req.URL, "respCode", resp.StatusCode, "alert", string(alertBys))
		} else {
			slog.Error("client.Do Failed", "err", err, "msg", string(msg), "url", req.URL, "respCode", resp.StatusCode, "alert", string(alertBys))
		}
		if err == nil {
			return errors.New("respCode is larger than 400, code=" + strconv.Itoa(resp.StatusCode))
		}
		return err
	}
	defer resp.Body.Close()
	return nil
}

func (a *api) GetAll() ProvisionedAlertRules {
	req, err := common.Request(http.MethodGet, a.u, apiAllAlertRules, a.token, "", "", nil, nil)
	if err != nil {
		slog.Error("new request failed", "err", err)
		return nil
	}

	resp, err := a.client.Do(req)
	if err != nil || resp.StatusCode >= 400 {
		var msg json.RawMessage
		if decodeErr := json.NewDecoder(resp.Body).Decode(&msg); decodeErr != nil {
			slog.Error("client.Do Failed", "err", err, "url", req.URL)
		} else {
			slog.Error("client.Do Failed", "err", err, "msg", string(msg), "url", req.URL)
		}
	}
	defer resp.Body.Close()

	var alerts ProvisionedAlertRules
	if err := json.NewDecoder(resp.Body).Decode(&alerts); err != nil {
		slog.Error("decode response body to alerts failed", "err", err, "url", req.URL)
		return nil
	}

	return alerts
}

func (a *api) Get(uids ...string) ProvisionedAlertRules {

	if len(uids) == 0 {
		slog.Info("no alert uid")
		return nil
	}

	alerts := make(ProvisionedAlertRules, 0, len(uids))
	for _, uid := range uids {
		req, err := common.Request(http.MethodGet, a.u, strings.Replace(apiAlertRule, "{uid}", uid, 1), a.token, "", "", nil, nil)
		if err != nil {
			slog.Error("new request failed", "err", err)
			return nil
		}

		resp, err := a.client.Do(req)
		if err != nil || resp.StatusCode >= 400 {
			var msg json.RawMessage
			if decodeErr := json.NewDecoder(resp.Body).Decode(&msg); decodeErr != nil {
				slog.Error("client.Do Failed", "err", err, "url", req.URL)
			} else {
				slog.Error("client.Do Failed", "err", err, "msg", string(msg), "url", req.URL)
			}
		}
		defer resp.Body.Close()

		var alert ProvisionedAlertRule
		if err := json.NewDecoder(resp.Body).Decode(&alert); err != nil {
			slog.Error("decode response body to alert failed", "err", err, "url", req.URL)
			return nil
		}

		alerts = append(alerts, alert)
	}

	return alerts
}

func (a *api) Update(uid string, alert ProvisionedAlertRule) error {
	// TODO
	return nil
}

func (a *api) Delete(uids ...string) error {
	// TODO
	return nil
}
