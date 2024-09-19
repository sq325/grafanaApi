package cmd

import (
	"github.com/spf13/cobra"
	org "github.com/sq325/grafanaApi/cmd/org"
)

var orgCmd = &cobra.Command{
	Use:   "org",
	Short: "org commands",
	Long:  `org commands`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	RootCmd.AddCommand(orgCmd)
	orgCmd.AddCommand(org.GetCmd)
}
