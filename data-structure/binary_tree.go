package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Tree is not balanced
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	max := 2 * n

	for i := 0; i < max; i++ {
		t = insert(t, rand.Intn(max)) // insert a rand value
	}

	return t
}

func traverse(t *Tree) {

	if t == nil {
		return
	}

	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)

}

func insert(t *Tree, v int) *Tree {

	if t == nil {
		return &Tree{nil, v, nil}
	}

	if v == t.Value {
		return t
	}

	if v < t.Value {
		t.Left = insert(t.Left, v)
		return t
	} else {
		t.Right = insert(t.Right, v)
		return t
	}

}

func printAtLevel(t *Tree, level int) {
	if t == nil || level < 0 {
		return
	}

	if level == 0 {
		fmt.Print(t.Value, " ")
		return
	}

	printAtLevel(t.Left, level-1)
	printAtLevel(t.Right, level-1)
}

func height(t *Tree) int {
	if t == nil {
		return 0
	}

	max := func(a, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	return max(height(t.Left), height(t.Right)) + 1
}

func main() {
	t := create(10)
	fmt.Println("The value of the root is", t.Value)

	traverse(t)
	fmt.Println()
	t = insert(t, -10)
	t = insert(t, 8)
	traverse(t)
	fmt.Println()

	fmt.Println("The value of the root is", t.Value)
	printAtLevel(t, 0)

	fmt.Println()
	fmt.Printf("the height is %d\n", height(t))
}
