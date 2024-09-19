package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	_versionInfo   string
	buildTime      string
	buildGoVersion string
	_version       string
	author         string
	projectName    string
)

var RootCmd = &cobra.Command{
	Use:   "grafanaApi",
	Short: "A CLI tool to manage grafana",
	Long:  `A CLI tool to manage grafana`,
	Run: func(cmd *cobra.Command, args []string) {
		if version {
			fmt.Println(projectName, _version)
			fmt.Println("build time:", buildTime)
			fmt.Println("go version:", buildGoVersion)
			fmt.Println("author:", author)
			fmt.Println("version info:", _versionInfo)
		}
		cmd.Help()
	},
}

func Execute() error {
	return RootCmd.Execute()
}

// flags
var (
	version bool
)

func init() {
	RootCmd.PersistentFlags().String("http.token", "", "auth bearer token, The token is bound to the user of the organization.")
	RootCmd.PersistentFlags().StringP("http.port", "p", "3000", "grafana port")
	RootCmd.PersistentFlags().String("http.ip", "localhost", "grafana ip")
	RootCmd.PersistentFlags().String("http.user", "", "grafana user")
	RootCmd.PersistentFlags().String("http.passwd", "", "grafana passwd")
	// RootCmd.MarkFlagRequired("http.token")

	RootCmd.Flags().BoolVarP(&version, "version", "v", false, "show version")
}
