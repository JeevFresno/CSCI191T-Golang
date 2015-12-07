package main

import "fmt"

func f_a(x string, str string) string{
	s:=str+" is "+x+" years old"
	return s
}

func main(){
	x:=f_a("22","JJ")
	fmt.Println(x)
}