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

type Manager struct {
	User
	title string
}

// func (u User) Hello() {
// 	fmt.Println("hello,World")
// }

func (u User) Hello(name string) {

	fmt.Println("Hello ", name, "my name is ", u.Name)
}

func main() {

	u := User{1, "OK", 12}
	info(&u)

	//初始化匿名字段
	m := Manager{User: User{1, "OK", 12}, title: "123"}
	t := reflect.TypeOf(m)
	fmt.Printf("%#v\n", t.Field(0))

	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))

	//如何从指针获取值
	x := 123
	v := reflect.ValueOf(&x)
	v.Elem().SetInt(999)
	fmt.Println(x) //x的值从123变为999

	//如何修改struct
	Set(&u)
	fmt.Println(u)

	u.Hello("Joel")
	user := reflect.ValueOf(u)
	mv := user.MethodByName("Hello")
	args := []reflect.Value{reflect.ValueOf("joel")}
	mv.Call(args)
}

func info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("type: ", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("fields: ")
	//若 输入为&u，则会报错
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("转换错误")
		return
	}

	//取得字段信息
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", f.Name, f.Type, val)
	}
	//取得方法信息
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}

}

func Set(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("修改错误")
		return
	} else {

		v = v.Elem()
	}

	f := v.FieldByName("Name")

	if !f.IsValid() {
		fmt.Println("没有找到字段")
		return
	}
	if f.Kind() == reflect.String {
		f.SetString("ByeBye")
	}
}
