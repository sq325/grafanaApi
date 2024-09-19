package cmd

import (
	"github.com/spf13/cobra"
	alertcmd "github.com/sq325/grafanaApi/cmd/alert"
)

var AlertCmd = &cobra.Command{
	Use:   "alert",
	Short: "alert commands",
	Long:  `alert commands`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var (
	orgId     string
	folderUid string
	group     string
)

func init() {
	RootCmd.AddCommand(AlertCmd)
	AlertCmd.AddCommand(alertcmd.GetCmd)

	AlertCmd.PersistentFlags().StringVar(&orgId, "orgid", "1", "org id")
	AlertCmd.PersistentFlags().StringVar(&folderUid, "folderuid", "", "folder id")
	AlertCmd.PersistentFlags().StringVar(&group, "group", "", "rule group")
}
