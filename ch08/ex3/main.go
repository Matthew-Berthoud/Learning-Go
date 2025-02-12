package main

import "fmt"

// Write a generic singly linked list data type.
// Each element can hold a comparable value and has a pointer to the next element in the list.
// The methods to implement are as follows:

// adds a new element to the end of the linked list
// Add(T)

// adds an element at the specified position in the linked list
// Insert(T, int)

// returns the position of the supplied value, -1 if it's not present
// Index (T) int

type List[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

type Node[T comparable] struct {
	Value T
	Next *Node[T]
}

func (l *List[T]) Add(value T) {
	n := &Node[T]{Value: value}
	if l.Head == nil {
		l.Head = n
		l.Tail = n
		return
	}
	l.Tail.Next = n
	l.Tail = l.Tail.Next
}

func (l *List[T]) Insert(value T, position int) {
	n := &Node[T]{Value: value}
	if position == 0 {
		n.Next = l.Head
		l.Head = n
		return
	}
	c := l.Head
	for i := 1; i < position; i++ {
		c = c.Next
	}
	n.Next = c.Next
	c.Next = n
}

func (l *List[T]) Index(value T) int {
	c := l.Head
	i := 0
	for ; c.Value != value; i++ {
		if c.Next == nil {
			panic("value not found")
		}
		c = c.Next
	}
	return i
}

func (l *List[T]) Print() {
	fmt.Println("     HEAD   ")
	i := 0
	for n := l.Head; n != nil; n = n.Next {
		fmt.Println("       |    ")
		fmt.Println("   +-------+")
		fmt.Printf("%3d|  %3d  |\n", i, n.Value)
		fmt.Println("   +-------+")
		i++
	}
	fmt.Println("       |    ")
	fmt.Println("     TAIL   ")
}

func main() {
	l := List[int]{}
	for i := 0; i < 5; i++ {
		l.Add(i)
	}
	for i := 0; i < 5; i += 2 {
		l.Insert(66, i)
	}
	l.Print()

	fmt.Println(l.Index(3))
	fmt.Println(l.Index(2))
}
