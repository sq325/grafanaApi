package datasource

import (
	"encoding/json"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	"github.com/sq325/grafanaApi/pkg/apis/datasource"
	"github.com/sq325/grafanaApi/pkg/common"
)

var CreateCmd = &cobra.Command{
	Use:   "create [-f filename] [--datasource.name name] [--datasource.type type] [--datasource.url url] [--datasource.access access]",
	Short: "create datasource",
	Long:  `create datasource with given token.`,
	Run: func(cmd *cobra.Command, args []string) {
		ip, port, token, err := common.Ep(cmd)
		if err != nil {
			return
		}
		api := datasource.NewApi(ip+":"+port, token)
		if api == nil {
			return
		}

		switch {
		case file != "":
			f, err := os.Open(file)
			if err != nil {
				slog.Error("open file failed", "file", file, "err", err)
				return
			}
			defer f.Close()

			// file 中有多个 datesources
			dsList, ok := func(f *os.File) (datasource.DataSources, bool) {
				var dsList datasource.DataSources
				err = json.NewDecoder(f).Decode(&dsList)
				if err != nil {
					slog.Error("decode file to datasources failed", "file", file, "err", err)
					return nil, false
				}
				return dsList, true
			}(f)

			if ok {
				for _, ds := range dsList {
					if err := api.Create(ds); err != nil {
						return
					}
				}
				break
			}
			slog.Error("decode file to datasources failed, trying decode file to datesource", "file", file, "err", err)
			// file 中有一个 datasource
			ds, ok := func(f *os.File) (datasource.DataSource, bool) {
				var ds datasource.DataSource
				err = json.NewDecoder(f).Decode(&ds)
				if err != nil {
					return datasource.DataSource{}, false
				}
				return ds, true
			}(f)
			if ok {
				if err := api.Create(ds); err != nil {
					slog.Error("create datasource failed", "err", err, "datasource", ds)
					return
				}
			}
			slog.Error("decode file to datesource and datesources failed, exit", "file", file, "err", err)

		case name != "" && _type != "" && _url != "" && access != "":
			ds := datasource.DataSource{
				Name:   name,
				Type:   datasource.GetDSType(_type),
				Url:    _url,
				Access: datasource.DsAccess(access),
			}

			if err := api.Create(ds); err != nil {
				slog.Error("create datasource from name,type,url failed", "err", err, "datasource", ds)
				return
			}
		default:
			slog.Error("file or name type url access must be set")
		}

	},
}

var (
	file   string
	name   string
	_type  string
	_url   string
	access string
)

func init() {
	CreateCmd.Flags().StringVarP(&file, "file", "f", "", "datasource file")
	CreateCmd.Flags().StringVar(&name, "datasource.name", "", "datasource name")
	CreateCmd.Flags().StringVar(&_type, "datasource.type", "", "datasource type")
	CreateCmd.Flags().StringVar(&_url, "datasource.url", "", "datasource url")
	CreateCmd.Flags().StringVar(&access, "datasource.access", "", "datasource access")
}
