package cmd

import "github.com/spf13/cobra"

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate commands",
	Long:  `migrate commands`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
