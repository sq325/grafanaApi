package alert

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"
)

func TestUrl(t *testing.T) {
	u, err := url.Parse("http://1.2.3.4:8080/")
	if err != nil {
		t.Log("err:", err)
	}
	t.Logf("url: %+v", u)
	fmt.Println(u.JoinPath(apiAllAlertRules))
}

func Test_api_Create(t *testing.T) {
	// init
	var (
		alert1 ProvisionedAlertRule
		// alert2 ProvisionedAlertRule
	)
	json.Unmarshal([]byte(alert1Json), &alert1)

	type args struct {
		alert      ProvisionedAlertRule
		Provenance bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"alert1",
			args{
				alert:      alert1,
				Provenance: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := NewApi("localhost:3000", "glsa_KJQIl68ZSu1UPIJDltd5DVVlJUHhpjYb_9b4b0fbe")
			if err := a.Create(tt.args.alert, tt.args.Provenance); err != nil {
				t.Errorf("api.Create() error = %v", err)
			}
		})
	}
}

const (
	alert1Json = `{
  "id": 0,
  "uid": "",
  "orgID": 0,
  "folderUID": "edyqh3afkavwgb",
  "ruleGroup": "1m",
  "title": "库级业务成功率告警",
  "condition": "C",
  "data": [
    {
      "refId": "A",
      "queryType": "",
      "relativeTimeRange": {
        "from": 600,
        "to": 0
      },
      "datasourceUid": "f3d3ba6d-1322-4af6-852c-bbce6b95ccf1",
      "model": {
        "editorMode": "code",
        "expr": "sum by (appid) (rate(arts_transaction_cost_ms_count{status=\"success\"}[1m])) / sum by (appid) (rate(arts_transaction_cost_ms_count[1m]))",
        "instant": true,
        "intervalMs": 1000,
        "legendFormat": "__auto",
        "maxDataPoints": 43200,
        "range": false,
        "refId": "A"
      }
    },
    {
      "refId": "B",
      "queryType": "",
      "relativeTimeRange": {
        "from": 600,
        "to": 0
      },
      "datasourceUid": "__expr__",
      "model": {
        "conditions": [
          {
            "evaluator": {
              "params": [],
              "type": "gt"
            },
            "operator": {
              "type": "and"
            },
            "query": {
              "params": [
                "B"
              ]
            },
            "reducer": {
              "params": [],
              "type": "last"
            },
            "type": "query"
          }
        ],
        "datasource": {
          "type": "__expr__",
          "uid": "__expr__"
        },
        "expression": "A",
        "intervalMs": 1000,
        "maxDataPoints": 43200,
        "reducer": "last",
        "refId": "B",
        "settings": {
          "mode": "dropNN"
        },
        "type": "reduce"
      }
    },
    {
      "refId": "C",
      "queryType": "",
      "relativeTimeRange": {
        "from": 600,
        "to": 0
      },
      "datasourceUid": "__expr__",
      "model": {
        "conditions": [
          {
            "evaluator": {
              "params": [
                0.8
              ],
              "type": "lt"
            },
            "operator": {
              "type": "and"
            },
            "query": {
              "params": [
                "C"
              ]
            },
            "reducer": {
              "params": [],
              "type": "last"
            },
            "type": "query"
          }
        ],
        "datasource": {
          "type": "__expr__",
          "uid": "__expr__"
        },
        "expression": "B",
        "intervalMs": 1000,
        "maxDataPoints": 43200,
        "refId": "C",
        "type": "threshold"
      }
    }
  ],
  "updated": "2024-09-13T10:25:39+08:00",
  "noDataState": "NoData",
  "execErrState": "Error",
  "for": "3m",
  "annotations": {
    "description": "库名：{{ index $labels \"appid\" }}\\n阈值：80%",
    "summary": "库级业务成功率连续3分钟低于阈值"
  },
  "labels": {
    "appNameLabel": "{{ $labels.appid }}",
    "severity": "critical"
  },
  "isPaused": false
}`
)
