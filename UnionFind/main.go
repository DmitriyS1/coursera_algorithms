package main

import "fmt"

func main() {
	var unions = InitQuickFind(10)

	Unite(1, 2, &unions)
	fmt.Println(unions)
}
