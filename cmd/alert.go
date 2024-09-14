package cmd

import "github.com/spf13/cobra"

var AlertCmd = &cobra.Command{
	Use:   "alert",
	Short: "alert commands",
	Long:  `alert commands`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	RootCmd.AddCommand(AlertCmd)
}
