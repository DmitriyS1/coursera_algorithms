package stupid_live_coding

import (
	"testing"
)

func TestLinkedList_Append(t *testing.T) {
	ll := NewLinkedList[int]()
	ll.Append(1)
	ll.Append(2)
	ll.Append(3)

	if ll.Size() != 3 {
		t.Fatalf("expected: 3, got: %v", ll.Size())
		return
	}

	currentEl := ll.head
	for currentEl != nil {
		//t.Log(strconv.Itoa(currentEl.value))
		t.Log(currentEl.value)
		currentEl = currentEl.next
	}
}

func TestLinkedList_Reverse(t *testing.T) {
	ll := NewLinkedList[int]()

	for i := 0; i < 10; i++ {
		ll.Append(i)
	}

	ll1 := ll.Reverse()

	if ll1.Size() != 10 {
		t.Fatalf("expected: 10, got: %v", ll1.Size())
		return
	}

	currentEl := ll1.head
	for currentEl != nil {
		t.Log(currentEl.value)
		currentEl = currentEl.next
	}
}
