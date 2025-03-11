package stupid_live_coding

type Node[T any] struct {
	value T
	next  *Node[T]
}

// LinkedList is a singly linked list
type LinkedList[T any] struct {
	head *Node[T]
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

// Append adds a new element to the end of the linked list
func (ll *LinkedList[T]) Append(elem T) *LinkedList[T] {
	current := ll.head
	newNode := &Node[T]{value: elem}

	if current == nil {
		ll.head = newNode
		return ll
	}

	for current.next != nil {
		current = current.next
	}

	current.next = newNode
	return ll
}

// Size returns the number of elements in the linked list
func (ll *LinkedList[T]) Size() int {
	current := ll.head
	if current == nil {
		return 0
	}

	size := 1
	for current.next != nil {
		size++
		current = current.next
	}

	return size
}

// Reverse reverses the linked list
func (ll *LinkedList[T]) Reverse() *LinkedList[T] {
	size := ll.Size()
	nodes := make([]*Node[T], size)

	tail := ll.head
	for i := 0; i < size; i++ {
		nodes[i] = tail
		if tail.next != nil {
			tail = tail.next
		}
	}

	e := nodes[size-1]
	for i := size - 1; i >= 0; i-- {
		if i == size-1 {
			ll.head = nodes[i]
			e = ll.head
			continue
		}

		e.next = nodes[i]
		e = nodes[i]
	}

	e.next = nil
	return ll
}
