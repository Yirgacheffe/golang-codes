package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func hasCycle(t *Node) bool {
	if t == nil {
		return false
	}

	fast := t.Next
	slow := t

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if fast == slow {
			return true
		}
	}

	return false // not found ?
}

func detectCycle(t *Node) *Node {
	if t == nil || t.Next == nil {
		return t
	}

	cycled := false
	slow := t
	fast := t.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			cycled = true
			break
		}
	}

	if !cycled {
		return nil
	}

	slow = t
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}

func oddEvenList(t *Node) *Node {
	if t == nil || t.Next == nil {
		return t
	}

	evenHead := t.Next
	odd := t
	even := evenHead

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead
	return t
}

func main() {
	fmt.Println("Hello, cycle!")
}
