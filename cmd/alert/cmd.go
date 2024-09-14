package alert

import (
	"encoding/json"
	"fmt"
	"os"
	"slices"

	"github.com/spf13/cobra"
	"github.com/sq325/grafanaApi/cmd"
	alertApi "github.com/sq325/grafanaApi/pkg/apis/alert"
	"github.com/sq325/grafanaApi/pkg/common"
)

var GetCmd = &cobra.Command{
	Use:   "get [alertUid]...",
	Short: "get alert",
	Long:  `get alert`,
	Run: func(cmd *cobra.Command, args []string) {

		ip, port, token, err := common.Ep(cmd)
		if err != nil {
			fmt.Println("get endpoint err:", err)
			return
		}
		api := alertApi.NewApi(ip+":"+port, token)

		// all alerts
		if len(args) == 0 {
			api.GetAll()
		}

		// one alert
		if len(args) == 1 {
			uid := args[0]
			api.Get(uid)
		}
	},
}

// create a alert 的前置条件是datasource已经存在
var CreateCmd = &cobra.Command{
	Use:   "create -f alerts.json",
	Short: "create alerts from file",
	Long:  `create alert from file`,
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.Open(file)
		if err != nil {
			fmt.Println("open file err:", err)
			return
		}

		var alerts alertApi.ProvisionedAlertRules
		err = json.NewDecoder(f).Decode(&alerts)
		if err != nil {
		}

		if len(alerts) == 0 {
			fmt.Println("no alerts found in", file)
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

// flag
var (
	output string
	file   string
)

func init() {
	cmd.AlertCmd.AddCommand(GetCmd)
	GetCmd.Flags().StringVarP(&output, "output", "o", "output.json", "output file")
	GetCmd.Flags().Bool("ignore.id", false, "ignore id")   // remove id
	GetCmd.Flags().Bool("ignore.uid", false, "ignore uid") // remove uid

	CreateCmd.Flags().StringVarP(&file, "file", "f", "", "alerts file with json format") // file require

	CreateCmd.MarkFlagRequired("file")

}
