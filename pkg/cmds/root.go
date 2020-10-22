package cmds

import (
	"flag"
	"os"

	"github.com/appscode/go/flags"
	"github.com/appscode/go/log/golog"
	v "github.com/appscode/go/version"
	"github.com/spf13/cobra"
	genericapiserver "k8s.io/apiserver/pkg/server"
	clientsetscheme "k8s.io/client-go/kubernetes/scheme"
	"kmodules.xyz/client-go/logs"
	"kmodules.xyz/client-go/tools/cli"
	appcatscheme "kmodules.xyz/custom-resources/client/clientset/versioned/scheme"
	"kubedb.dev/apimachinery/client/clientset/versioned/scheme"
)

func NewRootCmd(version string) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:               "kubedb-operator [command]",
		Short:             `KubeDB operator by AppsCode`,
		DisableAutoGenTag: true,
		PersistentPreRun: func(c *cobra.Command, args []string) {
			// 启动时输出的那一堆参数解析
			flags.DumpAll(c.Flags())

			// 偷偷发送了分析报告
			cli.SendAnalytics(c, version)

			// kubedb 的 clientset 添加了 k8s 的所有原生 scheme
			scheme.AddToScheme(clientsetscheme.Scheme)

			// appcatalog 的客户端添加了 scheme。appcatalog 是另外一个模块中定义的自定义资源
			appcatscheme.AddToScheme(clientsetscheme.Scheme)

			// 设置日志配置
			cli.LoggerOptions = golog.ParseFlags(c.Flags())
		},
	}
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)

	// 日志读取配置
	// ref: https://github.com/kubernetes/kubernetes/issues/17162#issuecomment-225596212
	logs.ParseFlags()

	rootCmd.PersistentFlags().BoolVar(&cli.EnableAnalytics, "enable-analytics", cli.EnableAnalytics, "Send analytical events to Google Analytics")

	// v 所在的 github.com/appscode/go 包就是个 go 的工具类包
	rootCmd.AddCommand(v.NewCmdVersion())

	// 停止所有线程的信号
	stopCh := genericapiserver.SetupSignalHandler()

	// Root 命令和 Run 命令（作为 root 的子命令）
	rootCmd.AddCommand(NewCmdRun(version, os.Stdout, os.Stderr, stopCh))

	return rootCmd
}
