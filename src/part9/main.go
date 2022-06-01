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
	shareTest1()
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
	b:=sum
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
