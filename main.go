package main

import "fmt"

type ListNode struct {
	val  int
	next *ListNode
}

type LinkedList struct {
	head *ListNode
}

func (ll *LinkedList) Append(vi ...int) {
	for _, v := range vi {
		newNode := &ListNode{val: v, next: nil}

		if ll.head == nil {
			ll.head = newNode
			continue
		}

		c := ll.head
		for c.next != nil {
			c = c.next
		}
		c.next = newNode
	}
}

func (ll *LinkedList) Print() {
	c := ll.head
	for c != nil {
		fmt.Printf("%d->", c.val)
		c = c.next
	}
	fmt.Println("nil")
}

func main() {
	list1 := &LinkedList{}
	list2 := &LinkedList{}

	//[2,5,4]
	list1.Append(2, 4, 3)
	// list2.head = m1

	list1.Print()
	fmt.Println("")

	list2.Append(5, 6, 4)
	list2.Print()
}
