
package v1

import (
	"context"
	v1 "github.com/jackistom/kubectl-Ecs/pkg/apis/jackistom/v1"
	scheme "github.com/jackistom/kubectl-Ecs/pkg/client/clientset/versioned/scheme"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)


type SylixosesGetter interface {
	Sylixoses(namespace string) SylixosInterface  //给相应namespace的参数返回相应的管理接口
}

// SylixosInterface 实现了rest api风格的接口实现对sylixos上自定义Ecs资源的管理
type SylixosInterface interface {
	Create(ctx context.Context, sylixos *v1.Sylixos, opts metav1.CreateOptions) (*v1.Sylixos, error)
	Update(ctx context.Context, sylixos *v1.Sylixos, opts metav1.UpdateOptions) (*v1.Sylixos, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Sylixos, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SylixosList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Sylixos, err error)
	SylixosExpansion
}


type sylixoses struct {
	client rest.Interface
	ns     string
}

// 传入一个类型的Client以及namespace 返回sylixoses类型的client接口和namespace
func newSylixoses(c *JackistomV1Client, namespace string) *sylixoses {
	return &sylixoses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// 使用sylixos资源的名称返回相应的对象 and an error if there is any.
func (c *sylixoses) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Sylixos, err error) {
	result = &v1.Sylixos{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sylixoses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return result ,err
}

// 列出sylixos类型的资源
func (c *sylixoses) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SylixosList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SylixosList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sylixoses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return result,err
}

// 监视相应sylixos
func (c *sylixoses) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sylixoses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// 创建
func (c *sylixoses) Create(ctx context.Context, sylixos *v1.Sylixos, opts metav1.CreateOptions) (result *v1.Sylixos, err error) {
	result = &v1.Sylixos{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sylixoses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sylixos).
		Do().
		Into(result)
	return result,err
}

// update更新操作 有错报错
func (c *sylixoses) Update(ctx context.Context, sylixos *v1.Sylixos, opts metav1.UpdateOptions) (result *v1.Sylixos, err error) {
	result = &v1.Sylixos{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sylixoses").
		Name(sylixos.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sylixos).
		Do().
		Into(result)
	return result ,err
}

// 有错误返回错误 没错就删除
func (c *sylixoses) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sylixoses").
		Name(name).
		Body(&opts).
		Do().
		Error()
}

//删除连接
func (c *sylixoses) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sylixoses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do().
		Error()
}

// 对sylixos资源进行patch并 返回
func (c *sylixoses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Sylixos, err error) {
	result = &v1.Sylixos{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sylixoses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do().
		Into(result)
	return  result ,err
}
