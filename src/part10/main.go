package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	//goQuit()
	//goQuit1()
	//goQuit2()
	contextValueTest()
}

func contextValueTest() {
	var wg sync.WaitGroup
	wg.Add(4)
	ctx,stop:= context.WithCancel(context.Background())
	valCtx:=context.WithValue(ctx,"userId",2)
	go func() {
		defer wg.Done()
		getUser(valCtx)
	}()
	stop()
}


func getUser(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("【获取用户】", "携程退出")
			return
		default:
			userId := ctx.Value("userId")
			fmt.Println("【获取用户】", "用户ID为：", userId)
			time.Sleep(1 * time.Second)
		}
	}
}

func goQuit2() {
	var wg sync.WaitGroup
	// 用来停止监控狗
	wg.Add(1)
	ctx, stop := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		watchDog2(ctx, "[监控狗3]")
	}()
	// 发送停止指令
	stop()
	wg.Wait()
}

func goQuit1() {
	var wg sync.WaitGroup
	// 用来停止监控狗
	wg.Add(1)
	stopCh := make(chan bool)
	go func() {
		defer wg.Done()
		watchDog1(stopCh, "[监控狗2]")
	}()
	// 发送停止指令
	stopCh <- true
	wg.Wait()
}

func goQuit() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		watchDog("[监控狗1]")
	}()
	wg.Wait()
}

func watchDog2(ctx context.Context, name string) {
	// 开启 for select 循环，一直后台监控
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "停止指令已收到，马上停止")
			return
		default:
			fmt.Println(name, "正在监控......")
		}
		time.Sleep(time.Second)
	}
}

func watchDog1(stopCh chan bool, name string) {
	// 开启 for select 循环，一直后台监控
	for {
		select {
		case <-stopCh:
			fmt.Println(name, "停止指令已收到，马上停止")
			return
		default:
			fmt.Println(name, "正在监控......")
		}
		time.Sleep(time.Second)
	}
}

func watchDog(name string) {
	// 开启 for select 循环，一直后台监控
	for {
		select {
		default:
			fmt.Println(name, "正在监控......")
		}
		time.Sleep(time.Second)
	}
}
