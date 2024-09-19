package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sq325/grafanaApi/cmd/migrate"
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate commands",
	Long:  `migrate commands`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	RootCmd.AddCommand(MigrateCmd)
	MigrateCmd.AddCommand(migrate.AlertCmd)

	
}
