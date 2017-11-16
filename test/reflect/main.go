package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name   string "user name"
	Passwd string "user password"
}

func main() {
	user := &User{"zjq", "pass"}
	s := reflect.TypeOf(user).Elem()
	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Tag)
	}
}
