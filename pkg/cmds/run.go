package cmds

import (
	"io"

	"github.com/appscode/go/log"
	"github.com/spf13/cobra"
	"kmodules.xyz/client-go/meta"
	"kmodules.xyz/client-go/tools/cli"
	"kubedb.dev/operator/pkg/cmds/server"
)

func NewCmdRun(version string, out, errOut io.Writer, stopCh <-chan struct{}) *cobra.Command {
	// 服务的配置对象先初始化，这个在形成 command 阶段做的事情
	o := server.NewKubeDBServerOptions(out, errOut)

	cmd := &cobra.Command{
		Use:               "run",
		Short:             "Run kubedb operator in Kubernetes",
		DisableAutoGenTag: true,
		PreRun: func(c *cobra.Command, args []string) {
			cli.SendPeriodicAnalytics(c, version)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			// 此前源码中的全局变量等已经准备好了，此时开始执行启动 kubedb-server. kubedeb-server 中有两个内容，一个是 controller，一个
			// 是 generic server。generic server 和 admission 都有
			log.Infoln("Starting kubedb-server...")

			// 下面两个都是空方法，没有任何作用
			if err := o.Complete(); err != nil {
				return err
			}
			if err := o.Validate(args); err != nil {
				return err
			}

			// 使用 KubeDBServerOptions 启动 server。一般都是server启动时传入 option。option 启动 server ，阿三确实脑路清奇
			if err := o.Run(stopCh); err != nil {
				return err
			}
			return nil
		},
	}

	o.AddFlags(cmd.Flags())
	meta.AddLabelBlacklistFlag(cmd.Flags())

	return cmd
}
