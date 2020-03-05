package main

/**
 * @Author: tianxiongwu
 * @Date: 2019-08-09 11:43
 */

import (
	"fmt"
	"sync"
)

type Animal interface {
	run()
}

type dog struct {
	id string
}

func (dog) run() {
	println("dog is running")
}

type cat struct {
	id string
}

func (cat) run() {
	println("cat is running")
}

func sum(values []int, result chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	result <- sum
}

func add(x, y int) {
	fmt.Println(x + y)
}

var counter = 0

func count(lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println(counter)
	lock.Unlock()
}

func x(a Animal) {
	a.run()
}

func main() {
	//value := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//result := make(chan int, 2)
	//go sum(value[:len(value)/2], result)
	//go sum(value[len(value)/2:], result)
	//sum1, sum2 := <-result, <-result
	//fmt.Print(sum1, sum2, sum1+sum2)

	//lock :=&sync.Mutex{}
	//for i:=0;i<10 ;i++  {
	//	go count(lock)
	//}
	//for  {
	//	lock.Lock()
	//	c:=counter
	//	lock.Unlock()
	//	runtime.Gosched()
	//	if c>=10 {
	//		break
	//	}
	//}
	testMap := map[string][]string{"dfs": {"dfksf", "df"}}
	fmt.Print(testMap["df"])
}
