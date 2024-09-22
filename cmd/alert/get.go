package alert

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"strings"

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

		// filter
		filteredAlerts := make(alertApi.ProvisionedAlertRules, 0, len(alerts))
		for _, alert := range alerts {
			if filterOrgId != 0 && alert.OrgID != filterOrgId {
				continue
			}
			if filterFolderUid != "" && alert.FolderUID != filterFolderUid {
				continue
			}
			if filterGroup != "" && alert.RuleGroup != filterGroup {
				continue
			}
			filteredAlerts = append(filteredAlerts, alert)
		}

		// output
		if printDatasource {
			type modelDatasource struct {
				DatasourceUid  string `json:"datasourceUid"`
				DatasourceType string `json:"type"`
			}
			var result []modelDatasource
			var dsList []string
			for _, alert := range filteredAlerts {
				dss, err := datasourcesFromAlert(alert)
				if err != nil {
					slog.Error("get datasources from alert failed", "err", err, "alert", alert)
					return
				}
				for _, ds := range dss {
					if !slices.Contains(dsList, ds) {
						dsList = append(dsList, ds)
					}
				}
			}

			for _, ds := range dsList {
				dstype := strings.Split(ds, ";")
				uid, _type := dstype[0], dstype[1]
				if uid != "" {
					result = append(result, modelDatasource{
						DatasourceUid:  uid,
						DatasourceType: _type,
					})
				}
			}

			jsonBys, err := json.MarshalIndent(result, "", "  ")
			if err != nil {
				slog.Error("marshal alerts datasources failed", "err", err)
				return
			}
			fmt.Println(string(jsonBys))
			return
		}

		jsonBys, err := json.MarshalIndent(filteredAlerts, "", "  ")
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

var (
	filterOrgId     int64
	filterFolderUid string
	filterGroup     string
)

var (
	printDatasource bool
)

func init() {
	GetCmd.Flags().StringVarP(&outputfile, "output", "o", "", "output file")
	GetCmd.Flags().Int64Var(&filterOrgId, "filter.orgid", 0, "org id")
	GetCmd.Flags().StringVar(&filterFolderUid, "filter.folderuid", "", "folder uid")
	GetCmd.Flags().StringVar(&filterGroup, "filter.group", "", "rule group")
	// GetCmd.Flags().Bool("remove.id", false, "ignore id")   // remove id
	// GetCmd.Flags().Bool("remove.uid", false, "ignore uid") // remove uid

	GetCmd.Flags().BoolVar(&printDatasource, "datasource", false, "print datasources from alerts")
}

// return datesources: [uid;type, ...]
func datasourcesFromAlert(alert alertApi.ProvisionedAlertRule) ([]string, error) {
	uids := make([]string, 0, 1)
	for _, d := range alert.Data {
		if d.DatasourceUID != "__expr__" {
			if !slices.ContainsFunc(uids, func(e string) bool { return strings.HasPrefix(e, d.DatasourceUID) }) {
				// has type
				if d.Model != nil {
					var model struct {
						Datasource struct {
							Uid            string `json:"uid"`
							DatasourceType string `json:"type"`
						} `json:"datasource"`
					}
					if err := json.Unmarshal(d.Model, &model); err != nil {
						return nil, err
					}
					uids = append(uids, fmt.Sprintf("%s;%s", model.Datasource.Uid, model.Datasource.DatasourceType))
					continue
				}

				// no type
				uids = append(uids, fmt.Sprintf("%s;%s", d.DatasourceUID, ""))
			}
		}
	}

	if len(uids) == 0 {
		return nil, fmt.Errorf("no datasource found in alert")
	}
	return uids, nil
}
