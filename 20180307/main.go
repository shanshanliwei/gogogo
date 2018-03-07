package main

import "fmt"

func main() {

	liwei := people{}

	fmt.Println("var liwei age value ", liwei.age)
	fmt.Println("var liwei name value ", liwei.name)
	fmt.Println("var liwei value ", liwei)

	liwei.setAge(20)
	fmt.Println("14 set age: liwei:", liwei)
	liwei.setName("wade")
	fmt.Println("16 set name: liwei:", liwei)

}

type people struct {
	name string
	age  int
}

func (p people) setName(name string) {
	p.name = name
	fmt.Println("25  in set name", p)
}

func (p people) setAge(age int) {
	p.age = age
	fmt.Println("in set age", p)
}
