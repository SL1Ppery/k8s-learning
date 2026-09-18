package main

import (
	"k8s-learning/client"
	"k8s-learning/controller"
	"k8s-learning/factory"
	myparser "k8s-learning/parser"
	"k8s-learning/resource"
)

//"time"
// "k8s-learning/input"
// "k8s-learning/processor"
// "os"
// "strconv"
// "strings"

// 匿名函数+闭包
// func ResourceIdGenerator(restype string) func() string {
// 	id := 0
// 	var ResourceName string
// 	switch restype {
// 	case "pod":
// 		ResourceName = "Pod"
// 	case "deployment":
// 		ResourceName = "Deployment"
// 	default:
// 		ResourceName = "UnKnown"
// 	}

// 	return func() string {
// 		id++
// 		return fmt.Sprintf("%s-%d", ResourceName, id)
// 	}
// }

// channel part
// func Controller(Podch <-chan string, DepCh <-chan string, SvcCh chan string) {
// 	counts := 0
// 	for counts < 3 {
// 		select {
// 		case pod, ok := <-Podch:
// 			if !ok {
// 				Podch = nil
// 				counts++
// 				continue
// 			}
// 			fmt.Println(pod)
// 		case dep, ok := <-DepCh:
// 			if !ok {
// 				DepCh = nil
// 				counts++
// 				continue
// 			}
// 			fmt.Println(dep)
// 		case svc, ok := <-SvcCh:
// 			if !ok {
// 				SvcCh = nil
// 				counts++
// 				continue
// 			}
// 			fmt.Println(svc)
// 		}
// 	}
// 	fmt.Printf("所有资源处理完成")
// }
// //channel part
// func CreatePod(ch chan string) {
// 	ch <- "Pod nginx created"
// 	close(ch)
// }
// // channel part
// func CreateDeployment(ch chan<- string) {
// 	ch <- "Deployment ruoyi-backend"
// 	close(ch)
// }

// // channel part
// func CreateService(ch chan string) {
// 	ch <- "Service nginx"
// 	close(ch)
// }

// select+stopCh
//
//	func ResourceCreator(ResourceCh chan<- string, ctx context.Context, wg *sync.WaitGroup) {
//		defer wg.Done()
//		for i := 1; i < 10; i++ {
//			select {
//			case ResourceCh <- fmt.Sprintf("Pod nginx-%d", i):
//				fmt.Printf("Pod-%d发送成功,namespace=%v\n", i, ctx.Value("namespace"))
//				time.Sleep(time.Second)
//			case <-ctx.Done():
//				fmt.Printf("Pod创建停止:%v\n", ctx.Err())
//				return
//			}
//		}
//	}
//
//	func Controller(ResourceCh <-chan string, ctx context.Context, wg *sync.WaitGroup) {
//		defer wg.Done()
//		for {
//			select {
//			case message, ok := <-ResourceCh:
//				if !ok {
//					fmt.Println("Creator-channel已关闭")
//					return
//				}
//				fmt.Printf("namespace:%v  %s has been created\n", ctx.Value("namespace"), message)
//			case <-ctx.Done():
//				fmt.Printf("收到停止信号,Controller停止：%v", ctx.Err())
//				return
//			}
//		}
//	}

func main() {
	registry := factory.NewRegistry()
	registry.Register("Pod", resource.NewPod)
	registry.Register("Deployment", resource.NewDeployment)
	controllerManager := controller.NewControllerManager()
	fakeClient := &client.FakeClient{}
	podController := &controller.PodController{Client: fakeClient}
	deploymentController := &controller.DeploymentController{Client: fakeClient}
	controllerManager.Registry("Pod", podController)
	controllerManager.Registry("Deployment", deploymentController)
	resources := myparser.Parser("test.yaml", registry)
	for _, r := range resources {
		controllerManager.Handle(r)
	}
	// pm := resource.PodManage{}
	// NewPod1 := resource.Pod{
	// 	Name:      "nginx",
	// 	NameSpace: "default",
	// 	Status:    "running",
	// }
	// NewPod2 := resource.Pod{
	// 	Name:      "MYSQL",
	// 	NameSpace: "default",
	// 	Status:    "running",
	// }
	// NewPod3 := resource.Pod{
	// 	Name:      "ruoyi",
	// 	NameSpace: "ruoyi",
	// 	Status:    "pending",
	// }
	// pm.Pods = append(pm.Pods, NewPod1)
	// pm.Pods = append(pm.Pods, NewPod2)
	// pm.Pods = append(pm.Pods, NewPod3)
	// pm.Wg.Add(6)
	// for i := 1; i < 4; i++ {
	// 	go func(id int) {
	// 		defer pm.Wg.Done()
	// 		fmt.Printf("reader%d启动\n", id)
	// 		pm.GetInfo()
	// 		fmt.Printf("reader%d结束\n", id)
	// 	}(i)
	// 	go func(id int) {
	// 		defer pm.Wg.Done()
	// 		fmt.Printf("changer%d启动\n", id)
	// 		pm.ChangeStatus("nginx", "pending")
	// 		fmt.Printf("changer%d结束\n", id)
	// 	}(i)
	// }
	// pm.Wg.Wait()

	// var wg sync.WaitGroup
	// CreatorCh := make(chan string, 10)
	// parentCtx := context.WithValue(context.Background(), "namespace", "ruoyi")
	// CreatorCtx, CreatorCancel := context.WithTimeout(parentCtx, time.Second*5)
	// deadline := time.Now().Add(time.Second * 7)
	// ControllerCtx, ControllerCancel := context.WithDeadline(parentCtx, deadline)
	// defer CreatorCancel()
	// defer ControllerCancel()
	// wg.Add(2)
	// go ResourceCreator(CreatorCh, CreatorCtx, &wg)
	// go Controller(CreatorCh, ControllerCtx, &wg)
	// wg.Wait()

	// PodCh := make(chan string)
	// DepCh := make(chan string)
	// SvcCh := make(chan string)
	// go CreatePod(PodCh)
	// go CreateDeployment(DepCh)
	// go CreateService(SvcCh)
	// // fmt.Printf("当前channel缓冲数：%d\n", len(ch))
	// // fmt.Printf("当前channel缓冲总数：%d\n", cap(ch))
	// Controller(PodCh, DepCh, SvcCh)

	// var wg sync.WaitGroup
	// wg.Add(3)
	// go resource.CreateDeployment(&wg)
	// go resource.CreatePod(&wg)
	// go resource.CreateService(&wg)
	// fmt.Printf("Main Running\n")
	// wg.Wait()
	// fmt.Printf("Main end\n")

	//匿名函数
	/*
		nums, _ := input.Reader.ReadString('\n')
		nums = strings.TrimSpace(nums)
		parts := strings.Fields(nums)
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])
		add := func(a int, b int) int {
			return a + b
		}
		result := add(a, b)
		fmt.Printf("%d", result)
	*/

	// fmt.Printf("请输入要生成ID的资源名：")
	// restype, _ := input.Reader.ReadString('\n')
	// restype = strings.TrimSpace(restype)
	// ResourceID := ResourceIdGenerator(restype)
	// fmt.Println(ResourceID())
	// fmt.Println(ResourceID())

	// var r k8sinterface.Resource
	// var pods resource.Pods
	// r = &pods
	// r.AddRes()
	// r.GetInfo()
	// // r = &dep
	// // r.AddRes()
	// // r.GetInfo()
	// //var dep resource.Deployments

	// processor.ProcessorRescource(&pods)

	// for {
	// 	fmt.Print("选择你要进行的操作：\n1.添加Pod\n2.查看pod\n3.退出(exit)\n")
	// 	input, _ := input.Reader.ReadString('\n')
	// 	input = strings.TrimSpace(input)
	// 	choice, _ := strconv.Atoi(input)
	// 	switch choice {
	// 	case 1:
	// 		r.AddRes()
	// 	case 2:
	// 		r.GetInfo()
	// 	case 3:
	// 		os.Exit(0)
	// 	}
	// }
}
