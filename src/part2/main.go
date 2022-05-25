package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//varTest()
	//floatTest()
	//boolTest()
	//stringTest()
	//zeroTest()
	//varShortTest()
	//piTest()
	//valueTest()
	//constTest()
	//iotaTest()
	//stringConvertInt()
	stringsPackageTest()
}

func stringsPackageTest() {
	s := "Hello"
	fmt.Println(strings.HasPrefix(s, "H"))
	fmt.Println(strings.Index(s, "o"))
	fmt.Println(strings.ToUpper(s))
}

func stringConvertInt() {
	i := 1
	i2s := strconv.Itoa(i)
	s2i, err := strconv.Atoi(i2s)
	fmt.Println(i2s, s2i, err)

	f64 := 11.2222
	i2f := float64(i)
	f2i := int(f64)
	fmt.Println(i2f, f2i)
}

func iotaTest() {
	//const (
	//	one=1
	//	two=2
	//	three =3
	//	four =4
	//)

	const (
		// one=(0)+1
		one = iota + 1
		// two=(0+1)+1
		two
		// three=(0+1+1)+1
		three
		// four=(0+1+1+1)+1
		four
	)
	fmt.Println(one, two, three, four)
}

func constTest() {
	const name = "闪电"
	fmt.Println(name)
}

func valueTest() {
	i := 10
	i = 20
	fmt.Println("i的新值是", i)

}

func piTest() {
	i := 10
	pi := &i
	fmt.Println("i ", i, "指针 pi:", pi)
}

func varShortTest() {
	i := 10
	bf := false
	s1 := "Hello"
	fmt.Println(i, bf, s1)

}

func zeroTest() {
	var zi int
	var zf float64
	var zb bool
	var zs string
	fmt.Println(zi, zf, zb, zs)
}

func stringTest() {
	var s1 string = "Hello"
	var s2 string = "世界"
	fmt.Println("s1 is ", s1, "s2 is ", s2)
	fmt.Println("s1+s2=", s1+s2)
	s1 += s2
	fmt.Println("s1", s1, "s2 ", s2)
}

func boolTest() {
	var bf bool = false
	var bt bool = true
	fmt.Println("bf is ", bf, "bt is", bt)
}

/**
* 变量使用
 */
func varTest() {
	//var i int = 10
	//var i  = 10
	var (
		i = 0
		j = 1
	)
	fmt.Println(i)
	fmt.Println(j)
}

/**
浮点数使用
*/
func floatTest() {
	var f32 float32 = 2.2
	var f64 float64 = 10.3456
	fmt.Println("f32 is ", f32, "f64 is", f64)
}
