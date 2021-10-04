package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	callback interface{}
}
type Message struct {
	number int
	name   string
}

func main() {
	t := &Test{
		callback: func(a interface{}) {
			fmt.Printf("callback:%v\n", a)
		},
	}
	number := &Message{
		number: 290,
		name:   "mat ebal",
	}
	v := reflect.ValueOf(t.callback)
	vargs := []reflect.Value{reflect.ValueOf(number)}
	v.Call(vargs)
	fmt.Printf("v: %v\n", v)
	fmt.Printf("vargs: %v\n", vargs)
}
