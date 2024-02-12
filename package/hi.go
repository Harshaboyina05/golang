package main

import "fmt"

var points = []int{1, 3, 5, 3, 5, 675, 432, 654, 876, 89, 798, 2132}
var hello string = "hihihihihih" //declared at global so accesible inother function used in main.go

func sayhi(x string) {
	fmt.Printf("Hello %v how do yo do?\n", x)

}
