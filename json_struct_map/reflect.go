package main

import (
	"fmt"
	"reflect"
)

type Persion struct {
	Id       int
	Name     string
	Address  string
	Email    string
	School   string
	City     string
	Company  string
	Age      int
	Sex      string
	Proviece string
	Com      string
	PostTo   string
	Buys     string
	Hos      string
}



/**
todo
todo 反射时 ： Elem() 重要性
todo 反射时 ： Elem() 重要性
todo 反射时 ： Elem() 重要性

todo ValueOf(&person)
elem := reflect.ValueOf(&person).Elem()

todo  content {Id: *int}  结构体内部  v.Field(k).Elem()
t := reflect.TypeOf(content)
v := reflect.ValueOf(content)
for k := 0; k < t.NumField(); k++ {
	if !v.Field(k).IsNil() { // Check it is not nil
		fieldName := t.Field(k).Tag.Get("sql")
		valueName := v.Field(k).Elem().Interface()
		m[fieldName] = valueName
	}
}
 */
func main() {
	//StructToMapViaReflect()
	//StructToMapViaReflect2()
	StructToMapViaReflect3()
}


func StructToMapViaReflect() {
	m := make(map[string]interface{})
	person := Persion{
		Id:       98439,
		Name:     "zhaondifnei",
		Address:  "大沙地",
		Email:    "dashdisnin@126.com",
		School:   "广州第十五中学",
		City:     "zhongguoguanzhou",
		Company:  "sndifneinsifnienisn",
		Age:      23,
		Sex:      "F",
		Proviece: "jianxi",
		Com:      "广州兰博基尼",
		PostTo:   "蓝鲸XXXXXXXX",
		Buys:     "shensinfienisnfieni",
		Hos:      "zhonsndifneisnidnfie",
	}
	elem := reflect.ValueOf(&person).Elem()
	relType := elem.Type()
	for i := 0; i < relType.NumField(); i++ {
		m[relType.Field(i).Name] = elem.Field(i).Interface()
	}
	fmt.Println(m)
}
func StructToMapViaReflect2() {
	m := make(map[string]interface{})
	person := Persion{
		Id:       98439,
		Name:     "zhaondifnei",
		Address:  "大沙地",
		Email:    "dashdisnin@126.com",
		School:   "广州第十五中学",
		City:     "zhongguoguanzhou",
		Company:  "sndifneinsifnienisn",
		Age:      23,
		Sex:      "F",
		Proviece: "jianxi",
		Com:      "广州兰博基尼",
		PostTo:   "蓝鲸XXXXXXXX",
		Buys:     "shensinfienisnfieni",
		Hos:      "zhonsndifneisnidnfie",
	}
	elem := reflect.ValueOf(&person)
	relType := elem.Elem().Type()
	for i := 0; i < relType.NumField(); i++ {
		m[relType.Field(i).Name] = elem.Elem().Field(i).Interface()
	}
	fmt.Println(m)
}


type Content struct {
	ID             *int64   `sql:"id" json:"id"`
	Source         *string  `sql:"source" json:"source"`
	SourceId       *string  `sql:"source_id" json:"source_id"`
}


func StructToMapViaReflect3() {
	m := make(map[string]interface{})
	Id, Source, SourceId := int64(1), "sss", "ssID"
	content := Content{
		&Id, &Source, &SourceId,
	}
	t := reflect.TypeOf(content)
	v := reflect.ValueOf(content)
	for k := 0; k < t.NumField(); k++ {
		if !v.Field(k).IsNil() { // Check it is not nil
			fieldName := t.Field(k).Tag.Get("sql")
			valueName := v.Field(k).Elem().Interface()
			m[fieldName] = valueName
		}
	}
	fmt.Println(m)
}
