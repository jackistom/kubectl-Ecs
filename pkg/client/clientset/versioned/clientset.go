package versioned

import (
	"fmt"
	jackistomv1 "github.com/jackistom/kubectl-Ecs/pkg/client/clientset/versioned/typed/jackistom/v1"

	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	JackistomV1() jackistomv1.JackistomV1Interface
}

//ClientSet为一个自定义的对象的group创建一个client,用法访问他的资源
// 版本集包括了 server-supported API 组的rest client(就是加了个string) 和restclient
type Clientset struct {
	*discovery.DiscoveryClient
	jackistomV1 *jackistomv1.JackistomV1Client
}

// 获取到client接口
func (c *Clientset) JackistomV1() jackistomv1.JackistomV1Interface {
	return c.jackistomV1
}

// server-supported API 组的rest client(就是加了个string
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

//通过新的config实现一个rest client,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("未设置请求速率限制")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.jackistomV1, err = jackistomv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// /通过新的config实现一个rest client
// 报错警告
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.jackistomV1 = jackistomv1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

//新建定义类型的client
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.jackistomV1 = jackistomv1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
