package datasource

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	"github.com/sq325/grafanaApi/pkg/apis/datasource"
	"github.com/sq325/grafanaApi/pkg/common"
)

var GetCmd = &cobra.Command{
	Use:   "get [all]",
	Short: "get datasource from grafana",
	Long:  `Get current datasource with given token.`,
	Run: func(cmd *cobra.Command, args []string) {
		ip, port, token, err := common.Ep(cmd)
		if err != nil {
			return
		}
		api := datasource.NewApi(ip+":"+port, token)
		if api == nil {
			return
		}

		datasources := api.GetAll()
		if datasources == nil || len(datasources) == 0 {
			return
		}

		jsonBys, err := json.MarshalIndent(datasources, "", "  ")
		if err != nil {
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

var (
	outputfile string
)

func init() {
	GetCmd.Flags().StringVarP(&outputfile, "outputfile", "o", "", "output file")
}
