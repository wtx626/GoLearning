package main

import (
	"fmt"
	"gopkg.in/fatih/set.v0"
)

func main() {
	oldKey := set.New(set.ThreadSafe)
	newKey := set.New(set.ThreadSafe)
	oldKey.Add(1, 2, 3, 4, 5)
	newKey.Add(7)
	offKey := set.Difference(oldKey, newKey)
	fmt.Print(set.Difference(oldKey, newKey))
	fmt.Print(set.Difference(oldKey, set.Intersection(oldKey, newKey)))

	cod, ok := 1, 2
	cod, df := 3, 9
	fmt.Print(cod, ok, df)
	fmt.Print(offKey.Has(1))

	arr := []int64{12, 74, 86}
	fmt.Print(arr[len(arr)-1], "\n")

	for _,set:=range oldKey.List(){
		fmt.Print(set)
	}
}
