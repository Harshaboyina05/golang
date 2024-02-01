package main

import "fmt" //formatting string and controls

func main() {
	fmt.Println("hi")
	//strings

	var name string = "hello"
	var name1 = "second" //till considers as string since it is in double quotes and will not allow other type in future also
	var name2 string     //just declaration so to use in future

	name4 := "444" // just short hand for first time cannot be declared outside

	fmt.Println(name, name1, name2, name4)
	var age int = 12
	age2 := 34
	fmt.Println(age, age2)

	var number float32 = 14000000000000000000000.54564121222222224554854514

	fmt.Println(number)

	fmt.Printf("december is %v month of the year %v ", age, name)

	//Dem()
}
