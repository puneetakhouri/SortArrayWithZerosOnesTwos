package main

import (
	"fmt"
	"strconv"
)

func main() {

	arr := []int{0, 2, 1, 2, 0}
	zeroCount, oneCount, twoCount := 0, 0, 0
	for i := 0; i < len(arr); i++ {
		switch arr[i] {
		case 0:
			zeroCount = zeroCount + 1
		case 1:
			oneCount = oneCount + 1
		case 2:
			twoCount = twoCount + 1
		}
	}

	print(0, zeroCount)
	print(1, oneCount)
	print(2, twoCount)
	fmt.Println()
}

func print(num int, count int) {
	for count > 0 {
		fmt.Printf(strconv.Itoa(num) + " ")
		count = count - 1
	}
}
