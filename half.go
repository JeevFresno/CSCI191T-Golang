package main

import "fmt"

func half(x int) (int,bool){
	return x / 2, x%2 == 0
}

func main() {
	h, even := half(6)
	fmt.Println(h, even)
}