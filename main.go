package main

import "fmt"

// List represents a singly-linked list item holds
// values of any type.
type List[T any] struct {
	next *List[T] // This is not a slice, this is a list of type T
	val  T
}

// LinkedList ...
type LinkedList[T any] struct {
	length          int
	listData        *List[T]
	currentPosition *List[T]
}

// NewLinkedList ...
func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{length: 0, listData: nil, currentPosition: nil}
}

// PutItem ...
func (ll *LinkedList[T]) PutItem(item T) {
	var location *List[T]
	location = new(List[T])
	location.val = item
	location.next = ll.listData
	ll.listData = location
	ll.length++
}

// GetNextItem ...
func (ll *LinkedList[T]) GetNextItem() T {
	if ll.currentPosition == nil {
		ll.currentPosition = ll.listData
	} else {
		ll.currentPosition = ll.currentPosition.next
	}
	item := ll.currentPosition.val
	return item
}

func main() {
	// Here is a lit of string
	list := NewLinkedList[string]()
	list.PutItem("Generic")
	list.PutItem("Go")
	list.PutItem("Start")

	for i := 0; i < list.length; i++ {
		v := list.GetNextItem()
		fmt.Println("Item string value: ", v)
	}

	// And here is a list of integer
	listInt := NewLinkedList[int]()
	listInt.PutItem(5)
	listInt.PutItem(10)
	listInt.PutItem(3)

	for i := 0; i < listInt.length; i++ {
		v := listInt.GetNextItem()
		fmt.Println("Item int value: ", v)
	}

	// listInt.GetNextItem() 	// IVALID - This will panic, as you try to fetch none existing item
}
