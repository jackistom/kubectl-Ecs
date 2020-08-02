package main


import (
	"fmt"
	"time"
	"encoding/json"
	"github.com/golang/glog"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"
	cclient "github.com/jackistom/kubectl-Ecs/pkg/CClient"

	v1 "github.com/jackistom/kubectl-Ecs/pkg/apis/jackistom/v1"
	clientset "github.com/jackistom/kubectl-Ecs/pkg/client/clientset/versioned"
	sylixosescheme "github.com/jackistom/kubectl-Ecs/pkg/client/clientset/versioned/scheme"
	informers "github.com/jackistom/kubectl-Ecs/pkg/client/informers/externalversions/jackistom/v1"
	listers "github.com/jackistom/kubectl-Ecs/pkg/client/listers/jackistom/v1"
	//"C"
)

const controllerAgentName = "sylixos-controller"

const (
	SuccessSynced = "Synced"

	MessageResourceSynced = "Sylixos synced successfully"
)

// Controller is the controller implementation for Sylixos resources
type Controller struct {
	// kubeclientset is a standard kubernetes clientset
	kubeclientset kubernetes.Interface
	// sylixosclientset is a clientset for our own API group
	sylixosclientset clientset.Interface

	sylixosesLister listers.SylixosLister
	sylixosesSynced cache.InformerSynced //sylixos 类型容器对象的缓存

	workqueue workqueue.RateLimitingInterface

	recorder record.EventRecorder
}


//func GoSendAndRecv(msg string) string{
//	cs := C.CString(msg)
//	C.SendAndRece(cs)
//	str :=C.GoString(C.chars)
//	return str
//}
//
//func GoCloseClient(){
//	C.Close()
//}


// NewController returns a new sylixos controller
func NewController(
	kubeclientset kubernetes.Interface,
	sylixosclientset clientset.Interface,
	sylixosInformer informers.SylixosInformer) *Controller {

	utilruntime.Must(sylixosescheme.AddToScheme(scheme.Scheme))
	glog.V(4).Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.Infof)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeclientset.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})

	controller := &Controller{
		kubeclientset:    kubeclientset,
		sylixosclientset: sylixosclientset,
		sylixosesLister:   sylixosInformer.Lister(),
		sylixosesSynced:   sylixosInformer.Informer().HasSynced,
		workqueue:        workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "Sylixos"),
		recorder:         recorder,
	}

	fmt.Println("Ecs资源初始化成功\n")
	// Set up an event handler for when Sylixos resources change
	sylixosInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueueSylixos,
		UpdateFunc: func(old, new interface{}) {
			var syspecold v1.SylixosSpec
			var syspecnew v1.SylixosSpec
			oldSylixos := old.(*v1.Sylixos)
			newSylixos := new.(*v1.Sylixos)
			if oldSylixos.ResourceVersion == newSylixos.ResourceVersion {
				//如果版本一致的话，就表示没有实际更新的操作，立即返回
				return
			}
			jsold,err:=json.Marshal(oldSylixos.Spec)
			if err!=nil{
				glog.Info(err)

			}
			json.Unmarshal(jsold,&syspecold)
			jsnew,err:=json.Marshal(oldSylixos.Spec)
			if err!=nil{
				glog.Info(err)

			}
			json.Unmarshal(jsnew,&syspecnew)

			fmt.Print("成功更新Ecs容器,现在该容器名字为:",syspecold.Sylixosname,"  成功更新Ecs容器,现在该ECS容器具体路径为:",syspecnew.Sylixosname,"\n")
			cclient.UpdateContainer(syspecold.Sylixosname,syspecold.Sylixospath,syspecnew.Sylixospath)


			//os.Exit(1)
		},
		DeleteFunc: controller.enqueueSylixosForDelete,
	})

	return controller
}

//在此处开始controller的业务
func (c *Controller) Run(threadiness int, stopCh <-chan struct{}) error {
	defer runtime.HandleCrash()
	defer c.workqueue.ShutDown()

	//glog.Info("开始一次缓存数据同步")
	if ok := cache.WaitForCacheSync(stopCh, c.sylixosesSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	//glog.Info("资源监听启动____")
	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	//glog.Info("资源监听已经启动")
	<-stopCh
	//glog.Info("资源监听已经结束")
	//glog.Info("------------")

	return nil
}

func (c *Controller) runWorker() {
	for c.processNextWorkItem() {
		cclient.Init()
	}
}

// 取数据处理
func (c *Controller) processNextWorkItem() bool {

	obj, shutdown := c.workqueue.Get()

	if shutdown {
		return false
	}

	// We wrap this block in a func so we can defer c.workqueue.Done.
	err := func(obj interface{}) error {
		defer c.workqueue.Done(obj)
		var key string
		var ok bool

		if key, ok = obj.(string); !ok {

			c.workqueue.Forget(obj)
			runtime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
			return nil
		}
		// 在syncHandler中处理业务
		if err := c.syncHandler(key); err != nil {
			return fmt.Errorf("error syncing '%s': %s", key, err.Error())
		}

		c.workqueue.Forget(obj)
		fmt.Printf("成功同步到'%s'资源\n", key)
		return nil
	}(obj)

	if err != nil {
		runtime.HandleError(err)
		return true
	}

	return true
}

// 处理
func (c *Controller) syncHandler(key string) error {
	// Convert the namespace/name string into a distinct namespace and name
	var  syspec v1.SylixosSpec
	//sylixosmap:=make(map[string]string)

	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		runtime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	// 从缓存中取对象
	sylixos, err := c.sylixosesLister.Sylixoses(namespace).Get(name)
	if err != nil {
		// 如果Sylixos对象被删除了，就会走到这里，所以应该在这里加入执行
		if errors.IsNotFound(err) {
			fmt.Printf("位于kubernetes的%s空间中的%s资源已经被删除了\n", namespace, name)

			return nil
		}

		runtime.HandleError(fmt.Errorf("failed to list sylixos by: %s/%s", namespace, name))

		return err
	}


	c.recorder.Event(sylixos, corev1.EventTypeNormal, SuccessSynced, MessageResourceSynced)
	js,err:=json.Marshal(sylixos.Spec)
	if err!=nil{
		glog.Info(err)

	}
	//fmt.Print("详细信息的json形式:",string(js))
	json.Unmarshal(js,&syspec)
	return nil
}

// 数据先放入缓存，再入队列
func (c *Controller) enqueueSylixos(obj interface{}) {
	var key string
	var err error
	var syspec v1.SylixosSpec
	// 将对象放入缓存
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		runtime.HandleError(err)
		return
	}

	// 将key放入队列
	c.workqueue.AddRateLimited(key)

	tampsy:=obj.(*v1.Sylixos)
	js,err:=json.Marshal(tampsy.Spec)
	if err!=nil{
		glog.Info(err)

	}
	json.Unmarshal(js,&syspec)
	fmt.Print("成功创建ECS容器:",syspec.Sylixosname,"  该ECS容器具体路径为:",syspec.Sylixospath,"\n")
	cclient.CreateContainer(syspec.Sylixosname,syspec.Sylixospath)
	//os.Exit(1)




}
//更新操作

func (c *Controller) enqueueSylixosUpdate(old, new interface{}) {
	var syspecold v1.SylixosSpec
	var syspecnew v1.SylixosSpec
	oldSylixos := old.(*v1.Sylixos)
	newSylixos := new.(*v1.Sylixos)
	if oldSylixos.ResourceVersion == newSylixos.ResourceVersion {
		//版本一致，就表示没有实际更新的操作，立即返回
		return
	}
	jsold,err:=json.Marshal(oldSylixos.Spec)
	if err!=nil{
		glog.Info(err)

	}
	json.Unmarshal(jsold,&syspecold)
	jsnew,err:=json.Marshal(oldSylixos.Spec)
	if err!=nil{
		glog.Info(err)

	}
	json.Unmarshal(jsnew,&syspecnew)

	fmt.Print("成功更新Ecs容器,现在该容器名字为:",syspecold.Sylixosname,"  成功更新Ecs容器,现在该ECS容器具体路径为:",syspecnew.Sylixosname,"\n")
	cclient.UpdateContainer(syspecold.Sylixosname,syspecold.Sylixospath,syspecnew.Sylixospath)

	c.enqueueSylixos(new)

	//os.Exit(1)
}

// 删除操作
func (c *Controller) enqueueSylixosForDelete(obj interface{}) {
	var key string
	var err error
	var syspec v1.SylixosSpec

	// 从缓存中删除指定对象
	key, err = cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
	if err != nil {
		runtime.HandleError(err)
		return
	}
	//再将key放入队列
	c.workqueue.AddRateLimited(key)
	tampsy:=obj.(*v1.Sylixos)
	js,err:=json.Marshal(tampsy.Spec)
	if err!=nil{
		glog.Info(err)

	}
	json.Unmarshal(js,&syspec)
	fmt.Print("成功删除ECS容器:",syspec.Sylixosname,"  该ECS容器具体路径为:",syspec.Sylixospath,"\n")
	cclient.RemoveContainer(syspec.Sylixosname,syspec.Sylixospath)
	//os.Exit(1)

}


