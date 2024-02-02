package main

import "fmt"

func sayhi(x string) {

	fmt.Printf("say Goodmorning %v \n", x)
	fmt.Printf("say GoodAfternoon %v \n", x)
}

func main() {
	sayhi("Harsha")
	fmt.Printf("Good Bye!")
}
