package main

import "fmt"

func Dem() {

	//arrays
	var age23 [3]int = [3]int{10, 54, 12}
	fmt.Print(age23)

	var cost = [4]int{23, 43, 56, 76}
	fmt.Println(cost)
	//shorthand
	name := [5]string{"harsha", "sri", "boyina", "more", "onemore"}

	fmt.Println(name)
	// no limit

	no := []int{10, 22, 2, 3, 2, 34, 3, 4, 5, 234}
	fmt.Println(no)
	no[3] = 3333333333
	no = append(no, 444444)
	fmt.Println(no)
	fmt.Println(no[1:3])

}
