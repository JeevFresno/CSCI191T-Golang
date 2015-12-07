package main

import(
	"fmt"
	"strconv"
)

func max(numbers ...int) int {
	var largest int=0
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func main() {
	largest := max(4, 7, 9, 123, 543, 23, 435, 53, 125,99999)
	fmt.Println(largest)
	i,_:=strconv.Atoi("-42")
	fmt.Println(i)
}
