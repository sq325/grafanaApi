package alert

import (
	"encoding/json"
	"log/slog"
	"os"
	"strconv"
	"strings"

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
			alert.ID = 0
			alert.UID = ""
			alert.OrgID = 0

			// 验证 alerttitle 是否唯一

			// replace orgid
			if replaceOrgIds != nil {
				for _, orgid := range replaceOrgIds {
					orgids := strings.Split(orgid, ":")
					if len(orgids) != 2 {
						slog.Error("replace.orgid format error", "replace.orgid", orgid)
						return
					}
					source, err := strconv.Atoi(orgids[0])
					if err != nil {
						slog.Error("replace.orgid format error", "replace.orgid", orgid)
						continue
					}
					target, err := strconv.Atoi(orgids[1])
					if err != nil {
						slog.Error("replace.orgid format error", "replace.orgid", orgid)
						continue
					}
					if alert.OrgID == int64(source) {
						alert.OrgID = int64(target)
						break
					}
				}
			}

			// replace folderuid, 验证存在，不存在创建
			if replaceFolderUids != nil {
				for _, folderuid := range replaceFolderUids {
					folderuids := strings.Split(folderuid, ":")
					if len(folderuids) != 2 {
						slog.Error("replace.folderuid format error", "replace.folderuid", folderuid)
						continue
					}
					source := folderuids[0]
					target := folderuids[1]
					if alert.FolderUID == source {
						alert.FolderUID = target
						break
					}
				}
				// TODO: 创建folder
			}

			// replace grouptitle，验证存在，不存在创建
			if replaceGroupTitles != nil {
				for _, grouptitle := range replaceGroupTitles {
					grouptitles := strings.Split(grouptitle, ":")
					if len(grouptitles) != 2 {
						slog.Error("replace.grouptitle format error", "replace.grouptitle", grouptitle)
						continue
					}
					source := grouptitles[0]
					target := grouptitles[1]
					if alert.RuleGroup == source {
						alert.RuleGroup = target
						break
					}
				}
				// TODO: 创建group
			}

			// replace datasourceuid, 验证存在
		OUTLOOP:
			for _, datasourceuid := range replaceDatasourceUids {
				datasourceuids := strings.Split(datasourceuid, ":")
				if len(datasourceuids) != 2 {
					slog.Error("replace.datasourceuid format error", "replace.datasourceuid", datasourceuid)
					continue
				}
				source := datasourceuids[0]
				target := datasourceuids[1]
				for i, d := range alert.Data {
					if d.DatasourceUID == source {
						d.DatasourceUID = target
						// 防止 d.DatasourceUID 和 d.model.datasource.uid 冲突，剔除 d.model.datasource
						var modelMap map[string]any
						var newModel json.RawMessage
						if err := json.Unmarshal(d.Model, &modelMap); err != nil {
							slog.Error("unmarshal model failed", "err", err)
							continue
						}
						delete(modelMap, "datasource")
						newModel, err = json.Marshal(modelMap)
						if err != nil {
							slog.Error("marshal model failed", "err", err)
							continue
						}
						d.Model = newModel
						alert.Data[i] = d
						break OUTLOOP
					}
				}
			}

			if err := api.Create(alert, provenance); err != nil {
				slog.Error("create alert failed", "alert", alert, "err", err)
				return
			}
			slog.Info("create alert success", "alert", alert.Title)
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
	CreateCmd.Flags().BoolVar(&provenance, "provenance", true, "enable editing these alerts in the Grafana UI")

	// replace: source -> target
	CreateCmd.Flags().StringSliceVar(&replaceOrgIds, "replace.orgid", nil, "source:target org id")
	CreateCmd.Flags().StringSliceVar(&replaceFolderUids, "replace.folderuid", nil, "source:target folder uid")
	CreateCmd.Flags().StringSliceVar(&replaceGroupTitles, "replace.grouptitle", nil, "source:target group title")
	CreateCmd.Flags().StringSliceVar(&replaceDatasourceUids, "replace.datasourceuid", nil, "source:target datasource uid")
}

type checkAlert func(alert alertApi.ProvisionedAlertRule, api alertApi.API) bool

// 检查orgid是否存在
func checkOrgId(alert alertApi.ProvisionedAlertRule, api alertApi.API) bool {
	return false
}

// 检查folderuid是否存在
func checkFolderUid(alert alertApi.ProvisionedAlertRule, api alertApi.API) bool {
	return false
}

// 检查grouptitle是否存在
func checkGroupTitle(alert alertApi.ProvisionedAlertRule, api alertApi.API) bool {
	return false
}

// 检查标题是否唯一
func checkAlertTitle(alert alertApi.ProvisionedAlertRule, api alertApi.API) bool {
	return false
}

// 检查datasourceuid是否存在
func checkDatasourceUid(alert alertApi.ProvisionedAlertRule, api alertApi.API) bool {
	return false
}

// check all conditions
func checkAll(fns ...checkAlert) checkAlert {
	return func(alert alertApi.ProvisionedAlertRule, api alertApi.API) bool {
		for _, fn := range fns {
			if !fn(alert, api) {
				return false
			}
		}
		return true
	}
}
