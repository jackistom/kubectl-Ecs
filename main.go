package main

import "C"
import (
	"flag"
	"github.com/golang/glog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"time"

	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	cclient "github.com/jackistom/kubectl-Ecs/pkg/CClient"
	clientset "github.com/jackistom/kubectl-Ecs/pkg/client/clientset/versioned"
	informers "github.com/jackistom/kubectl-Ecs/pkg/client/informers/externalversions"
	"github.com/jackistom/kubectl-Ecs/pkg/signals"
)

var (
	masterURL  string
	kubeconfig string
	//args   string
	//inter  interface{} = os.Interrupt

)

func main() {

	cclient.Init()
	flag.Parse()
	//if os.Args[1]!=""{
	//	args=os.Args[1]
	//	GoSendAndRecv(args);
	//}

	// 处理信号量
	stopCh := signals.SetupSignalHandler()

	// 处理入参
	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	sylixosClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building example clientset: %s", err.Error())
	}

	sylixosInformerFactory := informers.NewSharedInformerFactory(sylixosClient, time.Second*30)

	//得到controller
	controller := NewController(kubeClient, sylixosClient,
		sylixosInformerFactory.Jackistom().V1().Sylixoses())

	//启动informer
	go sylixosInformerFactory.Start(stopCh)

	//controller开始处理消息
	if err = controller.Run(2, stopCh); err != nil {
		glog.Fatalf("Error running controller: %s", err.Error())
	}
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. .")

}
