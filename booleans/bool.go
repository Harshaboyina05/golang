package main

import "fmt"

func main() {
	// if elseif and else
	age := 10
	if age < 10 || age == 10 {
		fmt.Println("executed only for 1!") //or operator
	}
	fmt.Println(age == 24)
	fmt.Println(!(age > 26)) //not operator
	fmt.Println(age < 24)

	if age < 25 && age > 15 { //and operator
		fmt.Println("not yet silver jubliee")
	} else if age <= 24 && age > 23 {
		fmt.Println("1 more to go!")

	} else {
		fmt.Println("Grow up!")
	}
}
