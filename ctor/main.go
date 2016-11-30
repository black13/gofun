package main

import (
	"fmt"
)

type Thing struct {
	Name     string
	Num      int
	commands map[string]func() interface{}
}

func NewThing(someParameter string) *Thing {
	p := new(Thing)
	p.commands = make(map[string]func() interface{})
	p.commands["blah"] = p.blah
	/*
		this -> function
		map[string]int
		std::map<std::string,int>
		map[string] func () interface{}
	*/
	//commands["thing1"] = doit()
	p.Name = someParameter
	p.Num = 33 // <- a very sensible default value
	return p
}

func (v *Thing) blah() interface{} {
	return "his name is " + v.Name
}

func NewThingEx(someParameter string) *Thing {
	return &Thing{someParameter, 33, make(map[string]func() interface{})}
}

func main() {
	fmt.Println("Hello, playground")
	thing1 := NewThing("jimb0")

	fin, ok := thing1.commands["blah"]
	if ok {
		fmt.Println(fin())
	}

	fmt.Println(thing1)
	thing2 := NewThing("james")
	fmt.Println(thing2)

	fout, ok := thing2.commands["blah"]
	if ok {

		fmt.Println(fout())
	}
}
