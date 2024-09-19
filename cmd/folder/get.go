package folder

import "github.com/spf13/cobra"

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "get folder",
	Long:  `get folder`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
