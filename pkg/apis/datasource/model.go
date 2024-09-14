package datasource

type DataSource struct {
	Id              int64    `json:"id"`  // omitempty
	UID             string   `json:"uid"` // unique,omitempty
	OrgId           int64    `json:"orgId"`
	Name            string   `json:"name"` // unique,required
	Type            string   `json:"type"`
	TypeLogoUrl     string   `json:"typeLogoUrl"`
	Access          DsAccess `json:"access"`
	Url             string   `json:"url"`
	User            string   `json:"user"`
	Database        string   `json:"database"`
	BasicAuth       bool     `json:"basicAuth"`
	BasicAuthUser   string   `json:"basicAuthUser"`
	WithCredentials bool     `json:"withCredentials"`
	IsDefault       bool     `json:"isDefault"`
	// JsonData         *simplejson.Json       `json:"jsonData,omitempty"`
	SecureJsonFields map[string]bool `json:"secureJsonFields"`
	Version          int             `json:"version"`
	ReadOnly         bool            `json:"readOnly"`
	AccessControl    Metadata        `json:"accessControl,omitempty"`
}

type DsAccess string
type Metadata map[string]bool

type DataSources []*DataSource
