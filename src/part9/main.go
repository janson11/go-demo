package main

import (
	"fmt"
	"sync"
	"time"
)

//var sum = 0

var (
	sum   int
	mutex sync.Mutex
)

func main() {
	//shareTest()
	//shareTest1()
	//run()
	//doOnce()
	race()
}

/**
  10个人赛跑，1个裁判发号施令
 */
func race()  {
	// & 代表引用
	cond:=sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(11)
	for i := 0; i < 10; i++ {
		go func(num int) {
			defer wg.Done()
			fmt.Println(num,"号已经就位")
			cond.L.Lock()
			//等待发令枪响
			cond.Wait()
			fmt.Println(num,"号开始跑......")
			cond.L.Unlock()
		}(i)
	}

	//等待所有goroutine都进入wait状态
	time.Sleep(2*time.Second)
	go func() {
		defer wg.Done()
		fmt.Println("裁判已经就位，准备发令枪")
		fmt.Println("比赛开始，大家准备跑")
		// 发令枪响
		cond.Broadcast()
	}()

	// 防止函数提前返回退出
	wg.Wait()

}


/**
代码只执行一次
sync.Once 适用于创建某个对象的单例、只加载一次的资源等只执行一次的场景。
*/
func doOnce() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}

	// 用于等待携程执行完毕
	done := make(chan bool)
	//启动10个携程执行once.Do（onceBody）
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++{
		<-done
	}

}

func run() {
	var wg sync.WaitGroup
	// 因为要监控110个携程，所以设置计数为110
	wg.Add(110)
	for i := 0; i < 100; i++ {
		go func() {
			// 计数值减1
			defer wg.Done()
			add(10)
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			// 计数值减1
			defer wg.Done()
			fmt.Println("和为：", readSum())
		}()
	}

	// 一直的等待，直到计数值为0
	wg.Wait()
}

func shareTest1() {
	for i := 0; i < 100; i++ {
		go add(10)
	}

	for i := 0; i < 10; i++ {
		go fmt.Println("和为：", readSum())
	}
	// 防止提前退出
	time.Sleep(2 * time.Second)
}

func readSum() int {
	mutex.Lock()
	defer mutex.Unlock()
	b := sum
	return b
}

func shareTest() {
	for i := 0; i < 100; i++ {
		go add(10)
	}
	// 防止提前退出
	time.Sleep(2 * time.Second)
	fmt.Println("和为：", sum)
}

func add(i int) {
	mutex.Lock()
	sum += i
	mutex.Unlock()
}
