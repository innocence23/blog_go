package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//todo
	/**
	1. json 相关用 marshal unmarshal
	2， json 转出去 用 unmarshal
	3， 其他转json  用 marshal
	 */
	//todo
	json2map()
	json2struct()
	json2struct2()

	Struct2Json()
	Map2Json()
}

// 用 unmarshal
func json2map() {
	b := []byte(`{"IP": "192.168.11.22", "name": "SKY"}`)

	m := make(map[string]string)

	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println("Umarshal failed:", err)
		return
	}
	fmt.Printf("%T --- %+v\n", m, m)
}

// 用 unmarshal
func json2struct() {
	type Person struct {
		name string
		age  int
	}
	jsonStr := ` {"name":"liang", "age":12} `
	var person Person
	json.Unmarshal([]byte(jsonStr), &person)
	fmt.Printf("%T --- %+v\n", person, person) //{name1: age:0}
}
// 用 unmarshal todo 必须大写 必须大写 必须大写
func json2struct2() {
	type Person struct {
		Name string
		Age  int
	}
	jsonStr := ` {"name":"liang", "age":12} `
	var person Person
	json.Unmarshal([]byte(jsonStr), &person)
	fmt.Printf("%T --- %+v\n", person, person) // {Name:liang Age:12}
}

// 转成json 用 marshal
func Struct2Json() {
	type Person struct {
		Name string
		Age  int
	}
	p := Person{
		Name: "liangyongxing",
		Age: 29,
	}

	//Person 结构体转换为对应的 Json
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println("failed:", err)
		return
	}
	fmt.Printf("%T --- %+v\n", jsonBytes, string(jsonBytes))
}


// 转成json 也用 marshal
func Map2Json() {
	mapInstance := make(map[string]interface{})
	mapInstance["Name"] = "liang637210"
	mapInstance["Age"] = 28
	mapInstance["Address"] = "北京昌平区"

	jsonStr, err := json.Marshal(mapInstance)

	if err != nil {
		fmt.Println("failed:", err)
		return
	}

	fmt.Printf("%T --- %+v\n", jsonStr, string(jsonStr))
}
