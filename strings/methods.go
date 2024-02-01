package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	var greet = "Hi welcome to code lab"
	fmt.Println(strings.Index(greet, "we"))

	fmt.Println(strings.Split(greet, "we"))

	fmt.Println(strings.Split(greet, " "))

	age := []int{10, 23, 4, 3, 32, 5, 4, 6, 7, 43, 65, 5}
	strs := []string{"b", "ab", "abc", "abxd", "adce"}

	sort.Ints(age)
	sort.Strings(strs)
	index := sort.SearchInts(age, 32)

	fmt.Println(age, index, strs)
	fmt.Println(sort.SearchStrings(strs, "abxd"))

}
