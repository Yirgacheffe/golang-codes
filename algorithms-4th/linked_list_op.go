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

func insertionSort(t *Node) *Node {
	if t == nil || t.Next == nil {
		return t
	}

	var dummy = new(Node)
	dummy.Next = t

	lastSorted := t
	cur := lastSorted.Next

	for cur != nil {
		if lastSorted.Value <= cur.Value {
			lastSorted = lastSorted.Next
		} else {
			prev := dummy.Next
			for prev.Value <= cur.Value {
				prev = prev.Next
			}

			lastSorted.Next = cur.Next
			prev.Next = cur
			cur.Next = lastSorted
		}
		cur = lastSorted.Next
	}

	return dummy.Next
}

func swapPairs(t *Node) *Node {
	dummyHead := &Node{0, t}
	temp := dummyHead

	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next
		node2 := temp.Next.Next

		// swap
		temp.Next = node2
		node1.Next = node2.Next
		node2.Next = node1

		temp = node1 // move to node1
	}

	return dummyHead.Next
}

func removeValue(t *Node, target int) *Node {
	dummy := new(Node)
	node := dummy

	for t != nil {
		if t.Value != target {
			node.Next = t
			node = node.Next
		}
		t = t.Next
	}
	return dummy.Next
}

func removeDupInSorted(t *Node) *Node {
	dummy := new(Node)
	cur := dummy

	for t != nil {
		if cur == dummy || cur.Value != t.Value {
			cur.Next = t
			cur = cur.Next
		}
		t = t.Next
	}

	cur.Next = nil
	return dummy.Next
}

func main() {
	fmt.Println("Hello, cycle!")
}
