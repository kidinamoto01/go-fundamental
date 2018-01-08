package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name   string
	Number string
	Title  string
}

func (u User) SetName(nn string) User {
	result := User{nn, u.Number, u.Title}
	//fmt.Println(result)
	return result
}

func main() {
	u := User{"aa", "bb", "cc"}
	//c := u.SetName("test")

	user := reflect.ValueOf(u)
	mv := user.MethodByName("SetName")
	args := []reflect.Value{reflect.ValueOf("joel")}
	tt := mv.Call(args)
	fmt.Println(tt)
}
