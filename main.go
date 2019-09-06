package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	lvmv1 "k8s-crd-lvmnodestorage/pkg/apis/control/v1"
	lvmclientset "k8s-crd-lvmnodestorage/pkg/client/clientset/versioned"

	"k8s.io/klog"

	// apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/sample-controller/pkg/signals"
)

/* 控制器 */
type Controller struct {
	// 此控制器使用的客户端
	clientset lvmclientset.Interface
	// 此控制器使用的工作队列
	queue workqueue.RateLimitingInterface
	// 此控制器使用的共享Informer，SharedIndexInformer可以维护缓存中对象的索引
	informer cache.SharedIndexInformer
}

/* 主函数 */
var (
	// 参数变量
	masterURL  string
	kubeconfig string
)

// 启动控制器
func (c *Controller) Run(stopCh <-chan struct{}) {
	// 捕获应用程序崩溃并打印日志
	defer utilruntime.HandleCrash()
	// 关闭队列，从而导致Worker结束
	defer c.queue.ShutDown()

	klog.Info("启动控制器……")

	// 运行Informer
	go c.informer.Run(stopCh)

	// 在启动Worker之前，等待缓存同步完成
	if !cache.WaitForCacheSync(stopCh, c.informer.HasSynced) {
		utilruntime.HandleError(fmt.Errorf("同步缓存超时"))
		return
	}

	klog.Info("缓存已经同步，准备启动Worker")
	// 循环执行Worker，直到TERM
	wait.Until(c.runWorker, time.Second, stopCh)
}

// 启动Worker
func (c *Controller) runWorker() {
	for c.processNextItem() {
	}
}

// Worker的逻辑框架
func (c *Controller) processNextItem() bool {
	// 最大重试次数
	maxRetries := 3

	// 获取下一个元素，第2个出参提示队列是否已经关闭
	key, quit := c.queue.Get()
	if quit {
		return false
	}

	// 总是移除Key
	defer c.queue.Done(key)

	// 处理Key
	err := c.processItem(key.(string))

	if err == nil {
		// 处理成功，提示队列不再跟踪事件历史
		c.queue.Forget(key)
	} else if c.queue.NumRequeues(key) < maxRetries {
		klog.Errorf("处理%s事件失败，准备重试： %v", key, err)
		c.queue.AddRateLimited(key)
	} else {
		klog.Errorf("处理%s事件失败，放弃： %v", key, err)
		c.queue.Forget(key)
		utilruntime.HandleError(err)
	}
	return true
}

// Worker核心逻辑
func (c *Controller) processItem(key string) error {
	klog.Infof("开始处理事件%s", key)
	// 根据Key获取对象
	obj, exists, err := c.informer.GetIndexer().GetByKey(key)
	if err != nil {
		return fmt.Errorf("获取对象%s失败: %v", key, err)
	}
	fmt.Print(obj)
	if !exists {
		// 在这里处理对象删除事件
	} else {
		// 在这里处理对象创建事件
	}
	// 因为不进行Resync，不会有更新事件
	return nil
}

func main() {
	// 解析参数，存入上述变量
	flag.Parse()
	// cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	// if err != nil {
	// 	klog.Fatalf("构建kubeconfig失败: %s", err.Error())
	// }
	// 创建客户端，Clientset是一系列K8S API的集合
	// clientset, err := kubernetes.NewForConfig(cfg)
	lvmclientset, err := newNodeLVMKubeClient()
	if err != nil {
		klog.Fatalf("构建clientset失败: %s", err.Error())
	}
	// 信号处理通道，当进程接收到信号后，此通道可读
	stopCh := signals.SetupSignalHandler()

	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())

	informer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				// 仅仅列出所有命名空间的Pod
				return lvmclientset.LvmV1().NodeLVMStorages().List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				return lvmclientset.LvmV1().NodeLVMStorages().Watch(options)
			},
		},
		&lvmv1.NodeLVMStorage{},
		0,                // 不进行relist
		cache.Indexers{}, // map[string]IndexFunc
	)

	// 添加事件处理回调，仅仅是简单的入队
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{ // 此结构实现ResourceEventHandler
		AddFunc: func(obj interface{}) {
			// 从对象中抽取Key
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
		DeleteFunc: func(obj interface{}) {
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
	})

	// 构建控制器对象
	ctrl := Controller{
		lvmclientset,
		queue,
		informer,
	}

	// 启动
	ctrl.Run(stopCh)
}

func newNodeLVMKubeClient() (lvmclientset.Interface, error) {
	kubeConfigPath := os.Getenv("HOME") + "/.kube/config"
	// kubeConfigPath := "/etc/kubernetes/kubelet.conf"

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create out-cluster kube cli configuration: %v", err)
	}

	cli, err := lvmclientset.NewForConfig(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create custom kube client: %v", err)
	}
	return cli, nil
}
