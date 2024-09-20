package datasource

type DataSource struct {
	Id              int64    `json:"id,omitempty"`  // omitempty
	UID             string   `json:"uid,omitempty"` // unique,omitempty
	OrgId           int64    `json:"orgId,omitempty"`
	Name            string   `json:"name"` // unique,required
	Type            string   `json:"type"` // required
	TypeLogoUrl     string   `json:"typeLogoUrl,omitempty"`
	Access          DsAccess `json:"access"` // proxy
	Url             string   `json:"url"`    // required
	User            string   `json:"user,omitempty"`
	Database        string   `json:"database,omitempty"`
	BasicAuth       bool     `json:"basicAuth"`
	BasicAuthUser   string   `json:"basicAuthUser,omitempty"`
	WithCredentials bool     `json:"withCredentials,omitempty"`
	IsDefault       bool     `json:"isDefault,omitempty"`
	// JsonData         *simplejson.Json       `json:"jsonData,omitempty"`
	SecureJsonFields map[string]bool `json:"secureJsonFields,omitempty"`
	Version          int             `json:"version,omitempty"`
	ReadOnly         bool            `json:"readOnly,omitempty"`
	AccessControl    Metadata        `json:"accessControl,omitempty"`
}

type DsAccess string
type Metadata map[string]bool

type DataSources []DataSource

type datesourceType struct{}

func (dt datesourceType) Type(name string) string {
	switch name {
	case mysql:
		return mysql
	case elasticsearch:
		return elasticsearch
	case graphite:
		return graphite
	case prometheus:
		return prometheus
	case alertmanager:
		return alertmanager
	case jaeger:
		return jaeger
	default:
		return ""
	}
}

// datasource type
const (
	mysql         = "mysql"
	elasticsearch = "elasticsearch"
	graphite      = "graphite"
	prometheus    = "prometheus"
	alertmanager  = "alertmanager"
	jaeger        = "jaeger"
)
