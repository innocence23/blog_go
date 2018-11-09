package main

import (
	"fmt"
	"reflect"
)

func main() {
	type User struct {
		Id        int    `json:"id"`
		Username    string    `json:"username"`
		Password    string    `json:"password"`
	}
	user := User{5, "zhangsan", "password"}
	data := Struct2Map(user)
	fmt.Printf("%T --- %+v\n", data, data)
}


//当然 可以  struct -》 json -》 map   （marshal unmarshal）  效率低

// struct 转 map  用反射
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
