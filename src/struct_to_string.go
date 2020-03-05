package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

type demo struct {
	Id   string
	Name string
}

func main() {
	demos := []demo{
		{
			Id:   "1",
			Name: "2s",
		},
		{
			Id:   "2",
			Name: "ll",
		},
		{
			Id:   "3",
			Name: "name",
		},
	}
	for _, v := range demos {
		tmpdata := Struct2Map(v)
		str, err := json.Marshal(tmpdata)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(str))
	}

	demo1 := demo{
		Id:   "4",
		Name: "123",
	}
	demoJson, _ := json.Marshal(demo1)
	fmt.Println(string(demoJson))
}
