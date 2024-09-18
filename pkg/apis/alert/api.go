package alert

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/sq325/grafanaApi/pkg/common"
	"github.com/sq325/grafanaApi/pkg/httpclient"
)

// crud
type API interface {
	Create(alert *ProvisionedAlertRule, datasourceUid string) error
	GetAll() ProvisionedAlertRules
	Get(uids ...string) ProvisionedAlertRules
	Update(uid string, alert *ProvisionedAlertRule) error
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
	}

	return &api{
		u:      u,
		token:  token,
		client: httpclient.New(),
	}
}

var _ API = (*api)(nil)

func (a *api) Create(alert *ProvisionedAlertRule, datasourceUid string) error {

	return nil
}

func (a *api) GetAll() ProvisionedAlertRules {
	req, err := http.NewRequest(http.MethodGet, a.u.JoinPath(apiAllAlertRules).String(), nil)
	if err != nil {
		fmt.Println("new request err:", err)
		return nil
	}

	resp, err := a.client.Do(req)
	if err != nil {
		var msg any
		if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
			fmt.Printf("decode response body err:%s message:%s\n", err, msg)
			return nil
		}
	}
	defer resp.Body.Close()

	var alerts ProvisionedAlertRules
	if err := json.NewDecoder(resp.Body).Decode(&alerts); err != nil {
		fmt.Println("decode response body err:", err)
		return nil
	}

	return alerts
}

func (a *api) Get(uids ...string) ProvisionedAlertRules {
	if len(uids) == 0 {
		fmt.Println("no alert uid")
		return nil
	}

	alerts := make(ProvisionedAlertRules, 0, len(uids))
	for _, uid := range uids {
		req, err := http.NewRequest(http.MethodGet, a.u.JoinPath(strings.Replace(apiAlertRule, "{uid}", uid, 1)).String(), nil)
		if err != nil {
			fmt.Println("new request err:", err)
			return nil
		}

		resp, err := a.client.Do(req)
		if err != nil {
			var msg any
			if err := json.NewDecoder(resp.Body).Decode(&msg); err != nil {
				fmt.Printf("decode response body err:%s message:%s\n", err, msg)
				return nil
			}
		}
		defer resp.Body.Close()

		var alert ProvisionedAlertRule
		if err := json.NewDecoder(resp.Body).Decode(&alert); err != nil {
			fmt.Println("decode response body err:", err)
			return nil
		}

		alerts = append(alerts, alert)
	}

	return alerts
}

func (a *api) Update(uid string, alert *ProvisionedAlertRule) error {
	return nil
}

func (a *api) Delete(uids ...string) error {
	return nil
}
