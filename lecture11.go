package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("hello,World")
}

func main() {

	u := User{1, "OK", 12}
	info(u)
}

func info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("type: ", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("fields: ")

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", f.Name, f.Type, val)
	}
}
