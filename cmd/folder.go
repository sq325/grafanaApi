package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sq325/grafanaApi/cmd/folder"
)

var folderCmd = &cobra.Command{
	Use:   "folder",
	Short: "folder commands",
	Long:  `folder commands`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	RootCmd.AddCommand(folderCmd)
	folderCmd.AddCommand(folder.GetCmd)
}
