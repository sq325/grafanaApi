package alert

import (
	"encoding/json"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	alertApi "github.com/sq325/grafanaApi/pkg/apis/alert"
	"github.com/sq325/grafanaApi/pkg/common"
)

// create a alert 的前置条件是datasource已经存在
var CreateCmd = &cobra.Command{
	Use:   "create -f alerts.json [--replace.orgid source:target] [--replace.folderuid source:target] [--replace.grouptitle source:target] [--replace.datasourceuid source:target]",
	Short: "create alerts from file",
	Long:  `create alerts from file`,
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

		ip, port, token, err := common.Ep(cmd)
		if err != nil {
			slog.Error("get grafana endpoint failed", "err", err)
			return
		}
		api := alertApi.NewApi(ip+":"+port, token)
		if api == nil {
			slog.Error("create alert api failed", "ip", ip, "port", port, "token", token)
			return
		}

		for _, alert := range alerts {
			// TODO:
			// 剔除id uid

			// 验证 alerttitle 是否唯一

			// replace orgid

			// replace folderuid, 验证存在，不存在创建

			// replace grouptitle，验证存在，不存在创建

			// replace datasourceuid, 验证存在

			if err := api.Create(alert, provenance); err != nil {
				return
			}
		}
	},
}

var (
	replaceOrgIds         []string
	replaceFolderUids     []string
	replaceGroupTitles    []string
	replaceDatasourceUids []string
	provenance            bool
)

func init() {
	CreateCmd.Flags().StringVarP(&file, "file", "f", "", "alerts file with json format") // file require
	CreateCmd.MarkFlagRequired("file")
	CreateCmd.Flags().BoolVar(&provenance, "provenance", false, "enable editing these alerts in the Grafana UI")

	// replace: source -> target
	CreateCmd.Flags().StringSliceVar(&replaceOrgIds, "replace.orgid", nil, "source:target org id")
	CreateCmd.Flags().StringSliceVar(&replaceFolderUids, "replace.folderuid", nil, "source:target folder uid")
	CreateCmd.Flags().StringSliceVar(&replaceGroupTitles, "replace.grouptitle", nil, "source:target group title")
	CreateCmd.Flags().StringSliceVar(&replaceDatasourceUids, "replace.datasourceuid", nil, "source:target datasource uid")

}
