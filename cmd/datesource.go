package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sq325/grafanaApi/cmd/datasource"
)

var DatasourceCmd = &cobra.Command{
	Use:   "datasource",
	Short: "datasource commands",
	Long:  `datasource commands`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	RootCmd.AddCommand(DatasourceCmd)
	DatasourceCmd.AddCommand(datasource.GetCmd)
	DatasourceCmd.AddCommand(datasource.CreateCmd)
}
