package main

import "fmt"

func main() {
	//ifTest()
	//ifTestSimple()
	//switchTest()
	//fallthroughTest()
	//switchExpressTest()
	forTest()
}

func forTest() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("the sum is ", sum)
}

func switchExpressTest() {
	switch 2 > 1 {
	case true:
		fmt.Println("2>1")
	case false:
		fmt.Println("2<=1")
	}
}

func fallthroughTest() {
	switch j := 1; j {
	case 1:
		fallthrough
	case 2:
		fmt.Println("1")
	default:
		fmt.Println("没有匹配")
	}
}

func switchTest() {
	switch i := 6; {
	case i > 10:
		fmt.Println("i>10")
	case i > 5 && i <= 10:
		fmt.Println("5<i<=10")
	default:
		fmt.Println("i<=5")
	}
}

func ifTest() {
	i := 6
	if i > 10 {
		fmt.Println("i>10")
	} else if i > 5 && i <= 10 {
		fmt.Println("5<i<=10")
	} else {
		fmt.Println("i<=5")
	}
}

func ifTestSimple() {
	if i := 6; i > 10 {
		fmt.Println("i>10")
	} else if i > 5 && i <= 10 {
		fmt.Println("5<i<=10")
	} else {
		fmt.Println("i<=5")
	}
}
