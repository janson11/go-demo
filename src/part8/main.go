package main

import (
	"fmt"
	"time"
)

func main() {
	//moreDefer()
	//goroutineTest()
	//chanTest()
	//cacheChanTest()
	selectChannelTest()
}

func selectChannelTest() {
	// 声明三个存放结果的channel
	firstCh := make(chan string)
	secondCh := make(chan string)
	threeeCh := make(chan string)
	go func() {
		firstCh <- downloadFile("firstCh")
	}()

	go func() {
		secondCh <- downloadFile("secondCh")
	}()

	go func() {
		threeeCh <- downloadFile("threeeCh")
	}()

	// 开始select 多路复用，哪个channel能获取到值就说明哪个最先下载好，就用哪个
	select {
	case filePath := <-firstCh:
		fmt.Println(filePath)

	case filePath := <-secondCh:
		fmt.Println(filePath)

	case filePath := <-threeeCh:
		fmt.Println(filePath)

	}
}

func downloadFile(chanName string) string {
	time.Sleep(time.Second)
	return chanName + ":filePath"
}

func cacheChanTest() {
	cacheCh := make(chan int, 5)
	cacheCh <- 2
	cacheCh <- 3
	fmt.Println("cacheCh容量为：", cap(cacheCh), "元素个数为：", len(cacheCh))
}

func chanTest() {
	ch := make(chan string)

	go func() {
		fmt.Println("janson1")
		ch <- "gorotine 完成"
	}()

	fmt.Println("我是main goroutine")
	v := <-ch
	fmt.Println("接收到的chan中的值为：", v)
}

func goroutineTest() {
	go fmt.Println("janson")
	fmt.Println("我是main goroutine")
	time.Sleep(time.Second)
}

func moreDefer() {
	defer fmt.Println("First defer")
	defer fmt.Println("Second defer")
	defer fmt.Println("Three defer")
	defer fmt.Println("函数自身代码")
}
