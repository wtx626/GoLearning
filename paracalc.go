package main

/**
 * @Author: tianxiongwu
 * @Date: 2019-08-09 11:43
 */

import (
	"fmt"
)
func sum(values [] int, result chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}
	result <- sum
}
func main() {
	value := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := make(chan int, 2)
	go sum(value[:len(value)/2], result)
	go sum(value[len(value)/2:], result)
	sum1, sum2 := <-result, <-result
	fmt.Print(sum1, sum2, sum1+sum2)
}
