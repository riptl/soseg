// Package soseg implements a sorted sum tree
package soseg

import "fmt"

// Tree describes a list of weights sorted by unique keys.
// The tree also keeps track of the running total/sum of weights preceding each entry.
// Effectively, it's a specialized segment tree whose range entries all touch but don't overlap.
// The length of each range is equal to the entry weight.
type Tree struct {
	Root *Node
	size int
}

// A Node can be either a branch with two children or a leaf.
// The Value of branches is the sum of the children Values.
// Laeves carry a single weight and are marked Terminal.
// Each node points to its parent except the topmost (root).
type Node struct {
	Key      int
	Value    int
	Parent   *Node
	Children [2]*Node
	Terminal bool
}

// Put inserts a node by key with a positive size,
// or updates the size if a node with this key already exists.
func (t *Tree) Put(key int, size int) (created bool) {
	if t.Root == nil {
		t.Root = &Node{
			Key:      key,
			Value:    size,
			Terminal: true,
		}
		t.size++
		return
	}

	np := &t.Root
	for {
		n := *np
		if n.Terminal {
			break
		}
		if key < n.Key {
			np = &n.Children[0]
		} else {
			np = &n.Children[1]
		}
	}

	n := *np

	// Leaf reached
	if key == n.Key {
		old := n.Value
		n.Value = size
		n.Parent.addBranch(size - old)
		return false
	}

	branch := &Node{
		Value:  n.Value,
		Parent: n.Parent,
	}
	*np = branch
	newNode := &Node{
		Key:      key,
		Value:    size,
		Parent:   branch,
		Terminal: true,
	}
	n.Parent = branch

	if key < n.Key {
		branch.Key = n.Key
		branch.Children[0] = newNode
		branch.Children[1] = n
	} else {
		branch.Key = key
		branch.Children[0] = n
		branch.Children[1] = newNode
	}

	branch.addBranch(size)
	t.size++
	return true
}

// Get searches for the node with the specified key.
// It returns the size of the node and its offset (sum of preceding nodes).
func (t *Tree) Get(key int) (size int, offset int, ok bool) {
	if t.Root == nil {
		return 0, 0, false
	}

	n := t.Root
	for !n.Terminal {
		if key < n.Key {
			n = n.Children[0]
		} else {
			offset += n.Children[0].Value
			n = n.Children[1]
		}
	}
	if key == n.Key {
		return n.Value, offset, true
	} else {
		return 0, 0, false
	}
}

// Remove removes the node with the specified key.
func (t *Tree) Remove(key int) (ok bool) {
	if t.Root == nil {
		return false
	}

	var offset int
	var side, side2 int
	n := t.Root
	for {
		if n.Terminal {
			break
		}
		side2 = side
		if key < n.Key {
			side = 0
		} else {
			side = 1
			offset += n.Children[0].Value
		}
		n = n.Children[side]
	}

	if key != n.Key {
		return false
	}
	if n.Parent == nil {
		t.Root = nil
		t.size = 0
		return true
	}

	var parent *Node
	var pivot **Node
	if n.Parent.Parent == nil {
		parent = nil
		pivot = &t.Root
	} else {
		parent = n.Parent.Parent
		pivot = &n.Parent.Parent.Children[side2]
	}

	// Replace parent with neighbor
	neighbor := n.Parent.Children[1-side]
	*pivot = neighbor
	neighbor.Parent = parent
	neighbor.Parent.addBranch(-n.Value)
	t.size--
	return true
}

// Find returns the key with the range containing the specified point in O(log n).
func (t *Tree) Find(point int) (key int, ok bool) {
	if t.Root == nil {
		return 0, false
	}

	if point < 0 {
		return 0, false
	}

	var offset int
	n := t.Root
	for !n.Terminal {
		// Point outside the total tree range
		if point > offset+n.Value {
			return 0, false
		}

		mid := offset + n.Children[0].Value
		if point < mid {
			n = n.Children[0]
		} else {
			offset = mid
			n = n.Children[1]
		}
	}
	if point >= offset+n.Value {
		return 0, false
	}

	return n.Key, true
}

// Total returns the sum of all weights in O(1).
func (t *Tree) Total() int {
	if t.Root == nil {
		return 0
	}
	return t.Root.Value
}

func (n *Node) addBranch(delta int) {
	x := n
	for x != nil {
		x.Value += delta
		x = x.Parent
	}
}

// Empty returns true if tree does not contain any nodes.
func (t *Tree) Empty() bool {
	return t.size == 0
}

// Size returns the number of elements stored in the tree in O(1).
func (t *Tree) Size() int {
	return t.size
}

// Clear removes all nodes from the tree.
func (t *Tree) Clear() {
	t.Root = nil
	t.size = 0
}

func (t *Tree) Print() {
	fmt.Println("SoSeg Tree")
	if t.Root != nil {
		t.Root.print(0)
	}
}

func (n *Node) print(indent int) {
	if n.Terminal {
		printIndent(indent)
		fmt.Printf("- '%d/%d\n", n.Key, n.Value)
	} else {
		printIndent(indent)
		fmt.Printf("+ '%d/%d\n", n.Key, n.Value)
		n.Children[0].print(indent + 2)
		n.Children[1].print(indent + 2)
	}
}

func printIndent(n int) {
	for i := 0; i < n; i++ {
		fmt.Print(" ")
	}
}
