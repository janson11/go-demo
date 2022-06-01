package main

import "fmt"

func main() {
	personInit()

}

type WalkRun interface {
	Walk()
	Run()
}

func (p *person) Walk()  {
	fmt.Printf("%s能走\n",p.name)
}

func (p *person) Run()  {
	fmt.Printf("%s能跑\n",p.name)
}

func NewPerson(name string) *person {
	return &person{name:name}
}

func (addr address) String() string {
	return fmt.Sprintf("the addr is %s%s", addr.province, addr.city)
}

func printString(s fmt.Stringer) {
	fmt.Println(s.String())
}

func (p person) String() string {
	return fmt.Sprintf("the name is %s,age is %d", p.name, p.age)
}

func personInit() {
	p := person{
		age:  31,
		name: "janson",
		addr: address{
			province: "广东省",
			city:     "深圳市",
		},
	}
	fmt.Println(p.addr)
	printString(p)
}

type person struct {
	name string
	age  uint
	addr address
}

type address struct {
	province string
	city     string
}


