package main

import (
	"fmt"
	"math/rand"
)

const (
	MAX = 10000000
)

func main() {
	fmt.Println("vim-go")

	arr := [MAX]int{}
	// create an array with MAX values
	for i := 0; i < MAX; i++ {
		arr[i] = i + 1
	}

	rs := rand.NewSource(MAX)

	r := rand.New(rs)
	fmt.Printf("The length of arr is: %d\n", len(arr))
	for i := 0; i < 300; i++ {
		//fmt.Printf("A random number %v\n ", r.Intn(MAX))
		binary_search(&arr, r.Intn(MAX))
	}
}

func binary_search(list *[MAX]int, search int) {

	fmt.Printf("Lenof list: %d, item to search: %d\n", len(list), search)
}
