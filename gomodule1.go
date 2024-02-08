package gomodule1

import "fmt"

func Parentfunc() {
	fmt.Println("This is exposed in Module1")
}

func parentfunc1() {
	fmt.Println("This is not exposed in Module1")
}
