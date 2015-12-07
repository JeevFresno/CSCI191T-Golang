package main

import "fmt"

func dog_year(x int) int{
	return x*7
}

func main(){
	x:=dog_year(5)
	fmt.Print(x)
}