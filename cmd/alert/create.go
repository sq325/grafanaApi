package alert

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"slices"

	"github.com/spf13/cobra"
	alertApi "github.com/sq325/grafanaApi/pkg/apis/alert"
)

// create a alert 的前置条件是datasource已经存在
var CreateCmd = &cobra.Command{
	Use:   "create -f alerts.json",
	Short: "create alerts from file",
	Long: `create alerts from file.
Alerts in files must has same org`,
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Open(file)
		if err != nil {
			slog.Error("open file failed", "file", file, "err", err)
			return
		}

		var alerts alertApi.ProvisionedAlertRules
		err = json.NewDecoder(f).Decode(&alerts)
		if err != nil {
			slog.Error("decode file to alerts failed", "file", file, "err", err)
			return
		}

		if len(alerts) == 0 {
			slog.Info("no alerts found", "file", file)
			return
		}

		// TODO: check datasource exist
		dsUids := make([]string, 0, 1)
		for _, alert := range alerts {
			dss, err := datasourcesFromAlert(&alert)
			if err != nil {
			}
			if len(dss) > 0 {
				for _, ds := range dss {
					if !slices.Contains(dsUids, ds) {
						dsUids = append(dsUids, ds)
					}
				}
			}
		}

		// TODO:
		// 1. 如果 datasource 不存在，创建datasource

		// TODO:
		// 1. 剔除id uid
		// 2. 替换 orgid
		// 3. 替换 folderid
		// 4. 处理 rule gourp
		// 5. 验证 alerttitle 是否唯一
		// 6. X-Disable-Provenance
	},
}

// type datesource struct {
// 	UID  string `json:"uid"`
// 	Name string `json:"name"`
// 	Type string `json:"type"`
// }

func datasourcesFromAlert(alert *alertApi.ProvisionedAlertRule) ([]string, error) {
	uids := make([]string, 0, 1)
	for _, d := range alert.Data {
		if d.DatasourceUID != "__expr__" {
			if len(uids) > 0 && !slices.Contains(uids, d.DatasourceUID) {
				uids = append(uids, d.DatasourceUID)
			}
		}
	}

	if len(uids) == 0 {
		return nil, fmt.Errorf("no datasource found in alert")
	}
	return uids, nil
}

func init() {
	CreateCmd.Flags().StringVarP(&file, "file", "f", "", "alerts file with json format") // file require
	CreateCmd.Flags().Bool("Provenance", false, "enable editing these alerts in the Grafana UI")
	CreateCmd.MarkFlagRequired("file")
}
