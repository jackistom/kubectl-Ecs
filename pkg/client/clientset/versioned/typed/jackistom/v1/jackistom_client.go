

package v1

import (
	v1 "github.com/jackistom/kubectl-Ecs/pkg/apis/jackistom/v1"
	"github.com/jackistom/kubectl-Ecs/client/clientset/versioned/scheme"

	rest "k8s.io/client-go/rest"
)


type JackistomV1Interface interface {
	RESTClient() rest.Interface
	SylixosesGetter
}

type JackistomV1Client struct { // JackistomV1Client 实现与 jackistom.k8s.io group.交互接口
	restClient rest.Interface
}

func (c *JackistomV1Client) Sylixoses(namespace string) SylixosInterface {
	return newSylixoses(c, namespace)
}


func NewForConfig(c *rest.Config) (*JackistomV1Client, error) {  //通过所给的config创建一个JackistomV1client
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &JackistomV1Client{client}, err
}


func NewForConfigOrDie(c *rest.Config) *JackistomV1Client { //通过新的config实现一个rest client
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

//新建定义类型的client
func New(c rest.Interface) *JackistomV1Client {
	return &JackistomV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}//重置一个client的config


//就是返回一个rest 风格的client 然后验证一下
func (c *JackistomV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
