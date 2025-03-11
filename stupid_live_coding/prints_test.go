package stupid_live_coding

import "testing"

func TestPrintDeferInRange(t *testing.T) {
	printDeferInRange()
}

func TestPrintIntAsStr(t *testing.T) {
	printIntAsStr()
}

func TestPrintChangeName(t *testing.T) {
	person := &Person{Name: "Bob"}
	changeName(person)
	if person.Name != "Alice" {
		t.Fatalf("expected: Alice, got: %v", person.Name)
	}
}

func TestChangeNameCorrect(t *testing.T) {
	person := &Person{Name: "Bob"}
	changeNameCorrect(person)
	if person.Name != "Alice" {
		t.Fatalf("expected: Alice, got: %v", person.Name)
	}
}
