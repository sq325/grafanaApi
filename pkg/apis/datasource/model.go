package datasource

type DataSource struct {
	Id              int64          `json:"id,omitempty"`    // omitempty
	UID             string         `json:"uid,omitempty"`   // unique,omitempty
	OrgId           int64          `json:"orgId,omitempty"` // required if not token
	Name            string         `json:"name"`            // unique,required
	Type            DatesourceType `json:"type"`            // required
	TypeLogoUrl     string         `json:"typeLogoUrl,omitempty"`
	Access          DsAccess       `json:"access"` // proxy
	Url             string         `json:"url"`    // required
	User            string         `json:"user,omitempty"`
	Database        string         `json:"database,omitempty"`
	BasicAuth       bool           `json:"basicAuth"`
	BasicAuthUser   string         `json:"basicAuthUser,omitempty"`
	WithCredentials bool           `json:"withCredentials,omitempty"`
	IsDefault       bool           `json:"isDefault,omitempty"`
	// JsonData         *simplejson.Json       `json:"jsonData,omitempty"`
	SecureJsonFields map[string]bool `json:"secureJsonFields,omitempty"`
	Version          int             `json:"version,omitempty"`
	ReadOnly         bool            `json:"readOnly,omitempty"`
	AccessControl    Metadata        `json:"accessControl,omitempty"`
}

type DsAccess string
type Metadata map[string]bool

type DataSources []DataSource

type DatesourceType string

const (
	MySQL         DatesourceType = "mysql"
	Elasticsearch DatesourceType = "elasticsearch"
	Graphite      DatesourceType = "graphite"
	Prometheus    DatesourceType = "prometheus"
	Alertmanager  DatesourceType = "alertmanager"
	Jaeger        DatesourceType = "jaeger"
)

func GetDSType(_type string) DatesourceType {
	switch _type {
	case "mysql":
		return MySQL
	case "elasticsearch":
		return Elasticsearch
	case "graphite":
		return Graphite
	case "prometheus":
		return Prometheus
	case "alertmanager":
		return Alertmanager
	case "jaeger":
		return Jaeger
	default:
		return ""
	}
}
