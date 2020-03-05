package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	path := "/Users/wutianxiong/Downloads/test/"
	file := "test.log"
	// 检查文件是否存在
	err := os.MkdirAll(path, 0755)
	if err != nil {
		fmt.Println(err.Error())
	}
	file1, err1 := os.Create(path + file)
	if err1 != nil {
		log.Fatalln(err)
	}
	fmt.Println(file1)
}
