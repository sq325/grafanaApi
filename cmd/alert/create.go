package alert

import (
	"encoding/json"
	"log/slog"
	"os"

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
		defer f.Close()

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
		// dsUids := make([]string, 0, 1)
		// for _, alert := range alerts {
		// 	dss, err := datasourcesFromAlert(&alert)
		// 	if err != nil {
		// 	}
		// 	if len(dss) > 0 {
		// 		for _, ds := range dss {
		// 			if !slices.Contains(dsUids, ds) {
		// 				dsUids = append(dsUids, ds)
		// 			}
		// 		}
		// 	}
		// }

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

func init() {
	CreateCmd.Flags().StringVarP(&file, "file", "f", "", "alerts file with json format") // file require
	CreateCmd.MarkFlagRequired("file")
	CreateCmd.Flags().Bool("Provenance", false, "enable editing these alerts in the Grafana UI")

	// replace: source -> target
	CreateCmd.Flags().String("source.orgid", "", "source org id")
	CreateCmd.Flags().String("source.folderuid", "", "source folder uid")
	CreateCmd.Flags().String("source.group", "", "source group")
	CreateCmd.Flags().String("target.orgid", "", "target org id")
	CreateCmd.Flags().String("target.folderuid", "", "target folder uid")
	CreateCmd.Flags().String("target.group", "", "target group")

	// datasource
	CreateCmd.Flags().String("source.datasourceuid", "", "datasource uid")

}
