package main

import "fmt"

type Student struct {
	id   string
	name string
}

func main() {
	classA := []Student{{"1", "tom"}, {"2", "jk"}, {"3", "john"}}
	classB := []Student{{"1", "jerry"}, {"2", "lucy"}}
	for index, BStudent := range classB {
		for _, AStudent := range classA {
			if BStudent.id == AStudent.id {
				classB[index].name = AStudent.name
			}
		}
	}
	fmt.Println(classB)
}
