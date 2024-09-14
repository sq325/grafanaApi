package alert

import (
	"encoding/json"
	"time"

	"github.com/prometheus/common/model"
)

type ProvisionedAlertRules []ProvisionedAlertRule
type ProvisionedAlertRule struct {
	ID int64 `json:"id"` // unique,omitempty
	// required: false
	// minLength: 1
	// maxLength: 40
	// pattern: ^[a-zA-Z0-9-_]+$
	UID string `json:"uid"` // unique,omitempty
	// required: true
	OrgID int64 `json:"orgID"` // replace
	// required: true
	// example: project_x
	FolderUID string `json:"folderUID"` // replace
	// required: true
	// minLength: 1
	// maxLength: 190
	// example: eval_group_1
	RuleGroup string `json:"ruleGroup"` // replace
	// required: true
	// minLength: 1
	// maxLength: 190
	// example: Always firing
	Title string `json:"title"` // unique,required
	// required: true
	// example: A
	Condition string `json:"condition"`
	// required: true
	// example: [{"refId":"A","queryType":"","relativeTimeRange":{"from":0,"to":0},"datasourceUid":"__expr__","model":{"conditions":[{"evaluator":{"params":[0,0],"type":"gt"},"operator":{"type":"and"},"query":{"params":[]},"reducer":{"params":[],"type":"avg"},"type":"query"}],"datasource":{"type":"__expr__","uid":"__expr__"},"expression":"1 == 1","hide":false,"intervalMs":1000,"maxDataPoints":43200,"refId":"A","type":"math"}}]
	Data []AlertQuery `json:"data"`
	// readonly: true
	Updated time.Time `json:"updated,omitempty"`
	// required: true
	NoDataState NoDataState `json:"noDataState"`
	// required: true
	ExecErrState ExecutionErrorState `json:"execErrState"`
	// required: true
	For model.Duration `json:"for"`
	// example: {"runbook_url": "https://supercoolrunbook.com/page/13"}
	Annotations map[string]string `json:"annotations,omitempty"`
	// example: {"team": "sre-team-1"}
	Labels map[string]string `json:"labels,omitempty"`
	// readonly: true
	Provenance Provenance `json:"provenance,omitempty"`
	// example: false
	IsPaused bool `json:"isPaused"`
}

type AlertQuery struct {
	// RefID is the unique identifier of the query, set by the frontend call.
	RefID string `json:"refId"`
	// QueryType is an optional identifier for the type of query.
	// It can be used to distinguish different types of queries.
	QueryType string `json:"queryType"`
	// RelativeTimeRange is the relative Start and End of the query as sent by the frontend.
	RelativeTimeRange RelativeTimeRange `json:"relativeTimeRange"`

	// Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.
	DatasourceUID string `json:"datasourceUid"` // replace

	// JSON is the raw JSON query and includes the above properties as well as custom properties.
	Model json.RawMessage `json:"model"`
}

type RelativeTimeRange struct {
	From Duration `json:"from" yaml:"from"`
	To   Duration `json:"to" yaml:"to"`
}

type Duration time.Duration

type NoDataState string

const (
	Alerting NoDataState = "Alerting"
	NoData   NoDataState = "NoData"
	OK       NoDataState = "OK"
)

type ExecutionErrorState string

const (
	OkErrState       ExecutionErrorState = "OK"
	AlertingErrState ExecutionErrorState = "Alerting"
	ErrorErrState    ExecutionErrorState = "Error"
)

type Provenance string
