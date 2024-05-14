package main

import "fmt"


func main() {
	a := 1
	b := &a

	*b =  30
	fmt.Println(a)
}