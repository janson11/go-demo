package main

import (
	"errors"
	"fmt"
)

func main() {
	/*result, er := sum2(1, 2)
	fmt.Println(result, er)
	result1, er1 := sum2(-1, 2)
	if er1 != nil {
		fmt.Println(er1)
	} else {
		fmt.Println(result1)
	}*/
	fmt.Println(sum4(1,2))
	fmt.Println(sum4(1,2,3))
	fmt.Println(sum4(1,2,3,4))

}

func sum(a int, b int) int {
	return a + b
}
func sum1(a, b int) int {
	return a + b
}

/**
多值返回
*/
func sum2(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能负数")
	}
	return a + b, nil
}

/**
命名返回参数,不常用
*/
func sum3(a, b int) (sum int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能负数")
	}
	sum = a + b
	err = nil
	return

}

/**
可变参数
 */
func sum4(params ...int) int {
	sum := 0
	for _, i := range params {
		sum += i
	}
	return sum
}
