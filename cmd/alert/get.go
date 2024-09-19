package alert

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	alertApi "github.com/sq325/grafanaApi/pkg/apis/alert"
	"github.com/sq325/grafanaApi/pkg/common"
)

var GetCmd = &cobra.Command{
	Use:   "get [-o alerts.json] [alertUid...]",
	Short: "get alert",
	Long:  `Get alerts from grafana. If alertUid is empty, get all alerts. Support multiple alertUid`,
	Run: func(cmd *cobra.Command, args []string) {
		ip, port, token, err := common.Ep(cmd)
		if err != nil {
			slog.Error("get grafana endpoint failed", "err", err)
			return
		}
		api := alertApi.NewApi(ip+":"+port, token)
		if api == nil {
			return
		}

		var alerts alertApi.ProvisionedAlertRules
		if len(args) == 0 {
			alerts = api.GetAll()
		} else {
			alerts = api.Get(args...)
		}

		if len(alerts) == 0 {
			slog.Error("no alerts found")
			return
		}

		jsonBys, err := json.MarshalIndent(alerts, "", "  ")
		if err != nil {
			slog.Error("marshal alerts failed", "err", err)
			return
		}

		if outputfile != "" {
			f, err := os.Create(outputfile)
			if err != nil {
				slog.Error("create output file failed", "outputfile", outputfile, "err", err)
				return
			}
			defer f.Close()
			f.Write(jsonBys)
			return
		}
		fmt.Println(string(jsonBys))
	},
}

// flag
var (
	outputfile string
	file       string
)

func init() {
	GetCmd.Flags().StringVarP(&outputfile, "output", "o", "", "output file")

	GetCmd.Flags().Bool("remove.id", false, "ignore id")   // remove id
	GetCmd.Flags().Bool("remove.uid", false, "ignore uid") // remove uid
	GetCmd.Flags().String("replace.orgid", "", "replace orgid")
	GetCmd.Flags().String("replace.folderid", "", "replace folderid")
	GetCmd.Flags().String("replace.group", "", "replace group")

}
