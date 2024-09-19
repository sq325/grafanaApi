package org

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"reflect"

	"github.com/spf13/cobra"
	orgApi "github.com/sq325/grafanaApi/pkg/apis/org"
	"github.com/sq325/grafanaApi/pkg/common"
)

var GetCmd = &cobra.Command{
	Use:   "get [all]",
	Short: "get org from grafana",
	Long: `Get current org with given token.
Get all need user and passwd permission`,
	Run: func(cmd *cobra.Command, args []string) {
		ip, port, token, err := common.Ep(cmd)
		if err != nil {
			slog.Error("get grafana endpoint failed", "err", err)
			return
		}
		user, err := cmd.Flags().GetString("http.user")
		if err != nil {
			slog.Error("get http user failed", "err", err)
			return
		}
		passwd, err := cmd.Flags().GetString("http.passwd")
		if err != nil {
			slog.Error("get http passwd failed", "err", err)
			return
		}

		api := orgApi.NewApi(ip+":"+port, token, user, passwd)
		if api == nil {
			slog.Error("new org api failed")
			return
		}

		// current org
		if len(args) == 0 {
			org := api.GetCurrent()
			if reflect.DeepEqual(org, orgApi.Org{}) {
				slog.Error("get current org failed")
				return
			}
			jsonBys, err := json.MarshalIndent(org, "", "  ")
			if err != nil {
			}
			fmt.Println(string(jsonBys))
			return
		}

		// all orgs
		if len(args) == 1 && args[0] == "all" {
			orgs := api.GetAll()
			if len(orgs) == 0 {
				slog.Error("get all orgs failed")
				return
			}
			jsonBys, err := json.MarshalIndent(orgs, "", "  ")
			if err != nil {
			}
			fmt.Println(string(jsonBys))
			return
		}

	},
}
