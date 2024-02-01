package main

import "fmt"

func main() {
	a := 10
	// for a < 20 {
	// 	fmt.Println("the value of a got incresed by 2:", a)
	// 	a = a + 2
	// }
	num := []string{"2", "3", "54", "675", "harsha", "jhsajs", "hfureihf", "547", "67687", "89743", "4"}
	for i := 1; i < len(num); i++ {
		fmt.Println(num[i])

	}
	for a <= 200 {
		fmt.Println("the value a:", a)
		a = a * 5
	}

	for i := 1; i < 100000; i++ {
		fmt.Println("*", i)
		i = i * 17
	}
}
