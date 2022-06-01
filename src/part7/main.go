package main

import (
	"fmt"
	"strconv"
)

func main() {
	//string2int()
	/*sum,err:=add(-1,2)
	if err!=nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum)
	}*/

	defer func(){
		if p:=recover();p!=nil {
			fmt.Println(p)
		}
	}()
	connectMySQL("","root","123456")
}




func connectMySQL(ip,username,password string)  {
	if ip==""{
		panic("ip不能为空")
	}

}


func (e *MyError) Error() string  {
	return e.err.Error() +e.msg
}

type MyError struct {
	err error
	msg string
}

type commonError struct {
	errorCode int
	errorMsg  string
}

func (ce commonError) Error() string {
	return ce.errorMsg
}

func add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		//return 0, errors.New("a或b不能为负数")
		return 0, &commonError{
			errorCode: 1,
			errorMsg:  "a或者不能为负数",
		}
	} else {
		return a + b, nil
	}
}

func string2int() {
	i, err := strconv.Atoi("a")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

}
