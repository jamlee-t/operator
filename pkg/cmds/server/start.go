// start 就是系统的入口
package server

import (
	"fmt"
	"io"
	"net"

	"github.com/spf13/pflag"
	admissionv1beta1 "k8s.io/api/admission/v1beta1"
	genericapiserver "k8s.io/apiserver/pkg/server"
	genericoptions "k8s.io/apiserver/pkg/server/options"
	"kmodules.xyz/client-go/meta"
	"kmodules.xyz/client-go/tools/clientcmd"
	"kubedb.dev/operator/pkg/controller"
	"kubedb.dev/operator/pkg/server"
)

const defaultEtcdPathPrefix = "/registry/kubedb.com"

// DBServerOption 是命令行完成首次使用的对象，一切由它而起。
// run 命令的参数也靠这个类提供方法添加 flagset
type KubeDBServerOptions struct {
	RecommendedOptions *genericoptions.RecommendedOptions
	ExtraOptions       *ExtraOptions

	StdOut io.Writer
	StdErr io.Writer
}

// 初始化一个 option 对象
// 1. 这个对象可以给 run 命令添加 flag set
// 2. 这个对象可以启动 server
func NewKubeDBServerOptions(out, errOut io.Writer) *KubeDBServerOptions {
	o := &KubeDBServerOptions{
		// TODO we will nil out the etcd storage options.  This requires a later level of k8s.io/apiserver
		RecommendedOptions: genericoptions.NewRecommendedOptions(
			defaultEtcdPathPrefix,
			server.Codecs.LegacyCodec(admissionv1beta1.SchemeGroupVersion),
			genericoptions.NewProcessInfo("kubedb-operator", meta.Namespace()),
		),
		ExtraOptions: NewExtraOptions(),
		StdOut:       out,
		StdErr:       errOut,
	}
	o.RecommendedOptions.Etcd = nil
	o.RecommendedOptions.Admission = nil

	return o
}

// 为 run命令的 flagset 添砖加瓦, 可以这个函数测试会发现 run 一个命令也就没有了
func (o KubeDBServerOptions) AddFlags(fs *pflag.FlagSet) {
	o.RecommendedOptions.AddFlags(fs)
	o.ExtraOptions.AddFlags(fs)
}

func (o KubeDBServerOptions) Validate(args []string) error {
	return nil
}

func (o *KubeDBServerOptions) Complete() error {
	return nil
}

// 获取 KubeDBServerConfig 对象，KubeDBServerConfig 对象中含有配置如下, 根据命令行参数来的:
// 	GenericConfig  *genericapiserver.RecommendedConfig
//	ExtraConfig    ExtraConfig
//	OperatorConfig *controller.OperatorConfig
//
func (o KubeDBServerOptions) Config() (*server.KubeDBServerConfig, error) {
	// TODO have a "real" external address
	if err := o.RecommendedOptions.SecureServing.MaybeDefaultWithSelfSignedCerts("localhost", nil, []net.IP{net.ParseIP("127.0.0.1")}); err != nil {
		return nil, fmt.Errorf("error creating self-signed certificates: %v", err)
	}

	// 把命令行参数应用到默认的 generic
	serverConfig := genericapiserver.NewRecommendedConfig(server.Codecs)
	if err := o.RecommendedOptions.ApplyTo(serverConfig); err != nil {
		return nil, err
	}
	clientcmd.Fix(serverConfig.ClientConfig)

	// ExtraOptions 往 controller 上拷贝对象
	controllerConfig := controller.NewOperatorConfig(serverConfig.ClientConfig)
	if err := o.ExtraOptions.ApplyTo(controllerConfig); err != nil {
		return nil, err
	}

	// 系统的真正的配置对象初始化完毕
	config := &server.KubeDBServerConfig{
		GenericConfig:  serverConfig,
		ExtraConfig:    server.ExtraConfig{},
		OperatorConfig: controllerConfig,
	}
	return config, nil
}

// kubedb 启动的入口
func (o KubeDBServerOptions) Run(stopCh <-chan struct{}) error {
	// 得到系统真正的配置对象
	config, err := o.Config()
	if err != nil {
		return err
	}

	// config  转为 completeConfig，然后调用 New 方法
	s, err := config.Complete().New()
	if err != nil {
		return err
	}

	return s.Run(stopCh)
}
