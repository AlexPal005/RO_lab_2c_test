package main

import (
	"fmt"
	"time"
)

func main() {
	arr_in := []int{3, 5, 7, 1, 10, 9, 2, 5}
	fight(&arr_in)
	time.Sleep(time.Second * 1 / 2)
	for len(arr_in) != 1 {
		fight(&arr_in)
	}
	fmt.Println(arr_in)
}
func run(c chan []int, number_first int, number_second int) { //перевіряє
	arr := <-c
	if arr[number_first] > arr[number_second] {
		arr[number_second] = 0
	} else {
		arr[number_first] = 0
	}
	fmt.Println(arr)
	c <- arr
}
func fight(arr *[]int) {
	c := make(chan []int)
	j := len(*arr) / 2
	for i := 0; i < len(*arr)/2; i++ {
		go run(c, i, j)
		c <- *arr
		j++
	}

	i := 0
	for i < len(*arr) {
		if (*arr)[i] == 0 {
			*arr = removeByIndex(*arr, i)
		} else {
			i++
		}
	}
}
func removeByIndex(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}
