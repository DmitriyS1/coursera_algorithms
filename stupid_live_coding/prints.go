package stupid_live_coding

import "fmt"

// printDeferInRange prints the numbers in reverse order
// because of the defer statement
// It works because the defer statement is executed in LIFO order
// and the range statement is executed in FIFO order
// So the defer statement is executed after the range statement
// and the last element is printed first
// and the first element is printed last
// So the output is:
// start
// end
// 5
// 4
// 3
// 2
// 1
func printDeferInRange() {
	fmt.Println("start")

	for _, v := range []int{1, 2, 3, 4, 5} {
		defer fmt.Println(v)
	}

	fmt.Println("end")
}

// printIntAsStr prints the ASCII character represented by the integer 65
// The ASCII character represented by the integer 65 is 'A'
// So the output is:
// A
// But in reality it's a compilation error because the string function
// expects a slice of bytes or rune and not an integer
func printIntAsStr() {
	i := 65
	s := string(rune(i))
	fmt.Println(s)
}

type Person struct {
	Name string
}

func changeName(person *Person) {
	person = &Person{Name: "Alice"}
}

func changeNameCorrect(person *Person) {
	person.Name = "Alice"
}
