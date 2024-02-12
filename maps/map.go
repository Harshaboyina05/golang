package main

import "fmt"

func main() {
	menu := map[string]float64{
		"1.pani-Poori": 20,
		"2.pavbhaji":   50,
		"3.samosa":     15,
		"4.tea":        12.23,
		"5.coffee":     67,
		"6.pastry":     75.50,
	}
	fmt.Print("what is the cost of tea ", menu["4.tea"], "\n")
	fmt.Printf("customer bill for 2 items %v", menu["4.tea"]+menu["6.pastry"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}
	menu["4.tea"] = 15
	fmt.Println(menu["4.tea"])
}
