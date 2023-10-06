package main

import "fmt"

func initQuickFind(n int) []int {
	var id []int
	for i := 0; i < n; i++ {
		id = append(id, i)
	}

	return id
}

func main() {
	var unions = initQuickFind(10)

	Unite(1, 2, &unions)
	fmt.Println(unions)
}
