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
	Use:   "get [uid...]",
	Short: "get all datasources or get datasources with given uids",
	Long:  `get all datasources with given token.`,
	Run: func(cmd *cobra.Command, args []string) {
		ip, port, token, err := common.Ep(cmd)
		if err != nil {
			return
		}
		api := datasource.NewApi(ip+":"+port, token)
		if api == nil {
			return
		}

		var dsList datasource.DataSources
		if len(args) == 0 {
			dsList = api.GetAll()
		} else {
			dsList = api.Get(args...)
		}

		if dsList == nil {
			slog.Error("Datasources not found")
			return
		}

		jsonBys, err := json.MarshalIndent(dsList, "", "  ")
		if err != nil {
			slog.Error("marshal datasources failed", "err", err)
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
