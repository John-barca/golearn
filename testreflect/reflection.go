package main

import (
	"fmt"
	"reflect"
)

type Person interface {
	SayHello(name string)
	Run() string
}

type Hero struct {
	Name string
	Age int
	Speed int
}

func (hero *Hero) SayHello(name string) {
	fmt.Println("Hello" + name, ", I am " + hero.Name)
}

func (hero *Hero) Run() string {
	fmt.Println("I am running at spead " + string(hero.Speed))
	return "Running"
}

func main () {
	// 获取实例反射类型对象
	typeOfHero := reflect.TypeOf(Hero{})
	fmt.Printf("Hero's type is %s, kind is %s\n", typeOfHero, typeOfHero.Kind())
}