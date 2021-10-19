package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type TreeNode struct {
	Left  *TreeNode
	Value int
	Right *TreeNode
}

//----------------------------- ### ------------------
func invert(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	right := root.Right
	root.Right = invert(root.Left)
	root.Left = right

	return root // .......................... # sole...
}

//----------------------------- ### ------------------
func sameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Value == q.Value && sameTree(p.Left, q.Left) && sameTree(p.Right, q.Right)
}

//----------------------------- ### ------------------
func create(n int) *TreeNode {
	var t *TreeNode
	max := 2 * n

	for i := 0; i < max; i++ {
		t = insert(t, rand.Intn(max)) // insert a rand value
	}

	return t
}

func insert(t *TreeNode, v int) *TreeNode {
	if t == nil {
		return &TreeNode{nil, v, nil}
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

func traverse(root *TreeNode) {
	if root == nil {
		return
	}
	traverse(root.Left)
	fmt.Print(root.Value, " ")
	traverse(root.Right)
}

//----------------------------- ### ------------------
type Iterator struct {
	stack []*TreeNode
	curr  *TreeNode
}

func NewIterator(root *TreeNode) Iterator {
	return Iterator{curr: root}
}

func (it *Iterator) next() int {
	for node := it.curr; node != nil; node = node.Left {
		it.stack = append(it.stack, node)
	}

	l := len(it.stack)

	it.curr, it.stack = it.stack[l-1], it.stack[:l-1]
	result := it.curr.Value
	it.curr = it.curr.Right
	return result
}

func (it *Iterator) hasNext() bool {
	return it.curr != nil || len(it.stack) > 0
}

//----------------------------- ------------------------
func lowestAncestor(root, p, q *TreeNode) *TreeNode {

	if !checkExist(root, p) || !checkExist(root, q) {
		fmt.Println("Not exist!")
		return nil
	}

	for root != nil {
		if p.Value < root.Value && q.Value < root.Value {
			root = root.Left
		} else if (p.Value > root.Value) && (q.Value > root.Value) {
			root = root.Right
		} else {
			break
		}
	}

	return root // found~~!
}

func checkExist(root, node *TreeNode) bool {
	if root == nil || node == nil {
		return false
	}

	curr := root

	for curr != nil {
		if node.Value == curr.Value {
			return true
		}

		if node.Value < curr.Value {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}

	return false
}

//----------------------------- ------------------------
func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	abs := func(x int) int {
		if x > 0 {
			return x
		} else {
			return -1 * x
		}
	}

	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	lHeight := height(root.Left)
	rHeight := height(root.Right)

	if lHeight == -1 || rHeight == -1 || abs(lHeight-rHeight) > 1 {
		return -1
	}

	return max(lHeight, rHeight) + 1
}

//----------------------------- ------------------------
func maxSumPath(root *TreeNode) {
	maxGain(root)
	fmt.Println(maxSum)
}

var maxSum = math.MinInt32

func maxGain(node *TreeNode) int {
	if node == nil {
		return 0
	}

	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	lGain := maxGain(node.Left)
	rGain := maxGain(node.Right)

	// current gain of the node
	innerMax := lGain + node.Value + rGain
	maxSum = max(maxSum, innerMax)

	outMaxSum := node.Value + max(lGain, rGain)
	return max(outMaxSum, 0)
}

//----------------------------- ------------------------
func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	t := create(10)
	fmt.Println("The value of the root is", t.Value)

	traverse(t)

	fmt.Println()
	x := invert(t) // after invert
	traverse(x)
}

//----------------------------- ------------------------
