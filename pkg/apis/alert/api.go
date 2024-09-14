package alert

import (
	"net/url"

	"github.com/sq325/grafanaApi/pkg/common"
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
	ApiAllAlertRules       = "/api/v1/provisioning/alert-rules"
	ApiAllAlertRulesExport = "/api/v1/provisioning/alert-rules/export" // GET:AlertingFileExport

	// GET RESPONSE.BODY:ProvisionedAlertRule
	// PUT Header:X-Disable-Provenance:true RESPONSE.BODY:ProvisionedAlertRule
	// DELETE Header:X-Disable-Provenance:true
	ApiAlertRule = "/api/v1/provisioning/alert-rules/{uid}" // GET:ProvisionedAlertRule PUT:ProvisionedAlertRule DELETE
)

type api struct {
	u     *url.URL
	token string
}

// ep: ip:port
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

func (a *api) Create(alert *ProvisionedAlertRule, datasourceUid string) error {
	return nil
}

func (a *api) GetAll() ProvisionedAlertRules {
	return nil
}

func (a *api) Get(uids ...string) ProvisionedAlertRules {
	return nil
}

func (a *api) Update(uid string, alert *ProvisionedAlertRule) error {
	return nil
}

func (a *api) Delete(uids ...string) error {
	return nil
}
