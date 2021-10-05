package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

var root = new(Node)

func addNode(t *Node, v int) int {
	if t == nil {
		t = &Node{v, nil}
		root = t
		return 0
	}

	if t.Value == v {
		fmt.Println("Node already exists:", v)
		return -1
	}

	if t.Next == nil {
		next := &Node{v, nil}
		t.Next = next
		return -2
	}

	return addNode(t.Next, v)
}

func addSortedNode(t *Node, v int) int {
	if t == nil {
		t = &Node{v, nil}
		root = t
		return 0
	}

	if t.Value == v {
		fmt.Println("Node already exists:", v)
		return -1
	}

	if v < t.Value {
		t := &Node{v, nil}
		t.Next = root
		root = t
		return -3
	}

	if t.Next == nil {
		next := &Node{v, nil}
		t.Next = next
		return -2
	}

	if t.Value < v && v < t.Next.Value {
		newNode := &Node{v, t.Next}
		t.Next = newNode
		return -4
	}

	return addSortedNode(t.Next, v)
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}

	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}

	fmt.Println() // Give it a new line
}

func lookupNode(t *Node, v int) bool {

	if root == nil {
		t = &Node{v, nil}
		root = t
		return false
	}

	if v == t.Value {
		return true
	}

	if t.Next == nil {
		return false
	}

	return lookupNode(t.Next, v)

}

func size(t *Node) int {
	if t == nil {
		fmt.Println("-> Empty list!")
		return 0
	}

	i := 0
	for t != nil {
		i++
		t = t.Next
	}
	return i
}

// reverse - another solution is using stack
func reverse(t *Node) *Node {
	if t == nil {
		fmt.Println("-> Empty list!")
		return nil
	}

	var prev, next *Node
	curr := t

	for curr != nil {
		next = curr.Next
		curr.Next = prev // reconnect last loop node
		prev = curr      // the last node
		curr = next
	}
	return prev // head_ref of node to outside
}

func main() {

	fmt.Println(root)
	root = nil

	traverse(root)

	addNode(root, 1)
	addNode(root, -1)
	traverse(root)

	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 45)
	addNode(root, 5)
	addNode(root, 5)
	traverse(root)

	addNode(root, 100)
	traverse(root)

	if lookupNode(root, 100) {
		fmt.Println("Node exist!")
	} else {
		fmt.Println("Node does not exist!")
	}

	if lookupNode(root, -100) {
		fmt.Println("Node exists!")
	} else {
		fmt.Println("Node does not exists!")
	}

	traverse(reverse(root))

}
