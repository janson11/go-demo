package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//forTestUpdate()
	//arrayTest()
	//forRangeTest()
	//sliceTest()
	//mapTest()
	byteTest()
}

func byteTest()  {
	s:="Hello janson"
	bs:=[]byte(s)
	fmt.Println(bs)
	fmt.Println(s[0],s[1],s[2])
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i,r:=range s{
		fmt.Println(i,r)
	}
}

func mapTest() {
	nameAgeMap:=make(map[string]int)
	nameAgeMap["janson"]=20
	nameAgeMap["janson1"]=21
	nameAgeMap["janson2"]=22
	nameAgeMap["janson3"]=23
	fmt.Println(nameAgeMap)
	fmt.Println(len(nameAgeMap))
	age,ok:=nameAgeMap["janson"]
	if ok{
		fmt.Println(age)
	}
}

func sliceTest() {
	array := [...]string{"a", "b", "c", "d", "e","f"}
	// 注意：这里是包含索引 2，但是不包含索引 5 的元素，即在 : 右边的数字不会被包含。
	slice:= array[2:5]
	fmt.Println(slice)
	slice[1]="f"
	fmt.Println("修改之后的array",array)
	slice1:=make([]string,6,8)
	//slice1:=[]string{"a", "b", "c", "d", "e"}
	fmt.Println(len(slice1),cap(slice1))
	slice2:=append(slice1,"g","h")
	slice3:=append(slice1,slice...)
	fmt.Println("slice2:",slice2)
	fmt.Println(slice3)


}

func forRangeTest() {
	array := [...]string{"a", "b", "c", "d", "e"}
	for i,v:=range array{
		fmt.Printf("数组索引：%d,对应值：%s\n", i, v)
	}

	for _,v:=range array{
		fmt.Printf("数组对应值：%s\n", v)
	}
}

func arrayTest() {
	//array:=[5]string{"a","b","c","d","e"}
	array := [...]string{"a", "b", "c", "d", "e"}
	fmt.Println(array[2])

	for i := 0; i < 5; i++ {
		fmt.Printf("数组索引：%d,对应值：%s\n", i, array[i])
	}
}

func forTestUpdate() {
	sum := 0
	for i := 1; i <= 100; i++ {
		if i%2 != 0 {
			continue
		}
		sum += i
	}
	fmt.Println("the sum is ", sum)
}
