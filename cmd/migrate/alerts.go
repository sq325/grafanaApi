package migrate

import "github.com/spf13/cobra"

// create order: org -> datesource -> folder -> group -> dashboard -> alert

var AlertCmd = &cobra.Command{
	Use:   "alert",
	Short: "alert commands",
	Long:  `alert commands`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		checkConflicts()
		// TODO:
		// Prepare
		// 1. org已经建好
		// DO:
		// 1. 从 source grafana 按 token 导出 alerts 和 datasources 并保存成两个json文件
		// 2. 解析 alerts.json, 逐个创建 alert
		//  - 替换 OrgID 为 currentOrgID
		//  - 替换 FolderID，如果不存在则创建
		//  - 替换 datasource，如果不存在则创建

		// NOTE:
		// 1. 导入前后的 alert 的 folder 和 group 保持一致
		// 2. 如何处理 source datasource 和 target datasource 的对应关系？
	},
}

var (
	targetOrgId string
)

func init() {}

// TODO
func checkConflicts() {
	// 1. 检查目标grafana是否有冲突的alert
	// 2. 检查目标grafana是否有冲突的datasource
}
