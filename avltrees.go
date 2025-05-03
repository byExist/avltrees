package avltrees

import (
	"cmp"
	"iter"
)

// Node represents a node in the AVL tree.
type Node[K cmp.Ordered, V any] struct {
	key    K
	value  V
	height int
	size   int
	left   *Node[K, V]
	right  *Node[K, V]
	parent *Node[K, V]
}

// Key returns the key of the node.
func (n *Node[K, V]) Key() K {
	return n.key
}

// Value returns the value of the node.
func (n *Node[K, V]) Value() V {
	return n.value
}

// Tree represents an AVL tree.
type Tree[K cmp.Ordered, V any] struct {
	Root *Node[K, V]
}

// New returns a new empty AVL Tree.
func New[K cmp.Ordered, V any]() *Tree[K, V] {
	return &Tree[K, V]{}
}

// Clear removes all nodes from the AVL tree.
func Clear[K cmp.Ordered, V any](t *Tree[K, V]) {
	t.Root = nil
}

// Insert inserts a key-value pair into the AVL tree.
// Returns true if the key was inserted, or false if it replaced an existing key.
func Insert[K cmp.Ordered, V any](t *Tree[K, V], key K, value V) bool {
	var inserted bool
	t.Root, inserted = insertRec(t.Root, key, value, nil)
	return inserted
}

// Delete removes the node with the specified key from the AVL tree.
// Returns true if the key existed and was deleted.
func Delete[K cmp.Ordered, V any](t *Tree[K, V], key K) bool {
	var deleted bool
	t.Root, deleted = deleteRec(t.Root, key)
	if t.Root != nil {
		t.Root.parent = nil
	}
	return deleted
}

// Search finds and returns the node with the given key in the AVL tree.
// Returns the node and true if found, or nil and false otherwise.
func Search[K cmp.Ordered, V any](t *Tree[K, V], key K) (*Node[K, V], bool) {
	curr := t.Root
	for curr != nil {
		if key < curr.key {
			curr = curr.left
		} else if key > curr.key {
			curr = curr.right
		} else {
			return curr, true
		}
	}
	return nil, false
}

// Min returns the node with the smallest key in the AVL tree.
// Returns the node and true if the tree is not empty, or nil and false otherwise.
func Min[K cmp.Ordered, V any](t *Tree[K, V]) (*Node[K, V], bool) {
	if t.Root == nil {
		return nil, false
	}
	return minNode(t.Root), true
}

// Max returns the node with the largest key in the AVL tree.
// Returns the node and true if the tree is not empty, or nil and false otherwise.
func Max[K cmp.Ordered, V any](t *Tree[K, V]) (*Node[K, V], bool) {
	if t.Root == nil {
		return nil, false
	}
	return maxNode(t.Root), true
}

// Ceiling returns the node with the smallest key greater than or equal to the given key.
// Returns the node and true if such a key exists, or nil and false otherwise.
func Ceiling[K cmp.Ordered, V any](t *Tree[K, V], key K) (*Node[K, V], bool) {
	curr := t.Root
	var result *Node[K, V]
	for curr != nil {
		if key == curr.key {
			return curr, true
		} else if key < curr.key {
			result = curr
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return result, result != nil
}

// Floor returns the node with the largest key less than or equal to the given key.
// Returns the node and true if such a key exists, or nil and false otherwise.
func Floor[K cmp.Ordered, V any](t *Tree[K, V], key K) (*Node[K, V], bool) {
	curr := t.Root
	var result *Node[K, V]
	for curr != nil {
		if key == curr.key {
			return curr, true
		} else if key < curr.key {
			curr = curr.left
		} else {
			result = curr
			curr = curr.right
		}
	}
	return result, result != nil
}

// Higher returns the node with the smallest key greater than the given key.
// Returns the node and true if such a key exists, or nil and false otherwise.
func Higher[K cmp.Ordered, V any](t *Tree[K, V], key K) (*Node[K, V], bool) {
	curr := t.Root
	var result *Node[K, V]
	for curr != nil {
		if key < curr.key {
			result = curr
			curr = curr.left
		} else {
			curr = curr.right
		}
	}
	return result, result != nil
}

// Lower returns the node with the largest key less than the given key.
// Returns the node and true if such a key exists, or nil and false otherwise.
func Lower[K cmp.Ordered, V any](t *Tree[K, V], key K) (*Node[K, V], bool) {
	curr := t.Root
	var result *Node[K, V]
	for curr != nil {
		if key <= curr.key {
			curr = curr.left
		} else {
			result = curr
			curr = curr.right
		}
	}
	return result, result != nil
}

// Predecessor returns the in-order predecessor of the given node, if any.
func Predecessor[K cmp.Ordered, V any](n *Node[K, V]) (*Node[K, V], bool) {
	if n.left != nil {
		return maxNode(n.left), true
	}
	p := n.parent
	for p != nil && n == p.left {
		n = p
		p = p.parent
	}
	if p != nil {
		return p, true
	}
	return nil, false
}

// Successor returns the in-order successor of the given node, if any.
func Successor[K cmp.Ordered, V any](n *Node[K, V]) (*Node[K, V], bool) {
	if n.right != nil {
		return minNode(n.right), true
	}
	p := n.parent
	for p != nil && n == p.right {
		n = p
		p = p.parent
	}
	if p != nil {
		return p, true
	}
	return nil, false
}

// InOrder returns an iterator for in-order traversal of the AVL tree.
func InOrder[K cmp.Ordered, V any](t *Tree[K, V]) iter.Seq[Node[K, V]] {
	return func(yield func(Node[K, V]) bool) {
		stack := []*Node[K, V]{}
		curr := t.Root
		for curr != nil || len(stack) > 0 {
			for curr != nil {
				stack = append(stack, curr)
				curr = curr.left
			}
			n := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if !yield(*n) {
				return
			}
			curr = n.right
		}
	}
}

// Range returns an iterator for nodes with keys in the range [from, to).
func Range[K cmp.Ordered, V any](t *Tree[K, V], from, to K) iter.Seq[Node[K, V]] {
	return func(yield func(Node[K, V]) bool) {
		stack := []*Node[K, V]{}
		curr := t.Root
		for curr != nil || len(stack) > 0 {
			for curr != nil {
				stack = append(stack, curr)
				curr = curr.left
			}
			n := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if n.key >= from && n.key < to {
				if !yield(*n) {
					return
				}
			}
			if n.key >= to {
				curr = nil
			} else {
				curr = n.right
			}
		}
	}
}

// Rank returns the number of nodes with keys less than the given key.
func Rank[K cmp.Ordered, V any](t *Tree[K, V], key K) int {
	rank := 0
	curr := t.Root
	for curr != nil {
		if key < curr.key {
			curr = curr.left
		} else {
			leftSize := 0
			if curr.left != nil {
				leftSize = curr.left.size
			}
			if key == curr.key {
				rank += leftSize
				break
			}
			rank += leftSize + 1
			curr = curr.right
		}
	}
	return rank
}

// Kth returns the node with the given 0-based rank.
// Returns the node and true if such rank exists, or nil and false otherwise.
func Kth[K cmp.Ordered, V any](t *Tree[K, V], k int) (*Node[K, V], bool) {
	curr := t.Root
	for curr != nil {
		leftSize := 0
		if curr.left != nil {
			leftSize = curr.left.size
		}
		if k < leftSize {
			curr = curr.left
		} else if k > leftSize {
			k -= leftSize + 1
			curr = curr.right
		} else {
			return curr, true
		}
	}
	return nil, false
}

// Len returns the number of nodes in the AVL tree.
func Len[K cmp.Ordered, V any](t *Tree[K, V]) int {
	if t.Root == nil {
		return 0
	}
	return t.Root.size
}

func insertRec[K cmp.Ordered, V any](n *Node[K, V], key K, value V, parent *Node[K, V]) (*Node[K, V], bool) {
	if n == nil {
		return &Node[K, V]{key: key, value: value, height: 1, size: 1, parent: parent}, true
	}
	if key < n.key {
		var inserted bool
		n.left, inserted = insertRec(n.left, key, value, n)
		return rebalance(n), inserted
	} else if key > n.key {
		var inserted bool
		n.right, inserted = insertRec(n.right, key, value, n)
		return rebalance(n), inserted
	} else {
		n.value = value
		return n, false
	}
}

func deleteRec[K cmp.Ordered, V any](n *Node[K, V], key K) (*Node[K, V], bool) {
	if n == nil {
		return nil, false
	}
	var deleted bool
	if key < n.key {
		n.left, deleted = deleteRec(n.left, key)
	} else if key > n.key {
		n.right, deleted = deleteRec(n.right, key)
	} else {
		deleted = true
		if n.left == nil || n.right == nil {
			var child *Node[K, V]
			if n.left != nil {
				child = n.left
			} else {
				child = n.right
			}
			if child != nil {
				child.parent = n.parent
			}
			return child, true
		}
		successor := n.right
		for successor.left != nil {
			successor = successor.left
		}
		n.key, n.value = successor.key, successor.value
		n.right, _ = deleteRec(n.right, successor.key)
	}
	return rebalance(n), deleted
}

func height[K cmp.Ordered, V any](n *Node[K, V]) int {
	if n == nil {
		return 0
	}
	return n.height
}

func updateSize[K cmp.Ordered, V any](n *Node[K, V]) {
	n.height = max(height(n.left), height(n.right)) + 1
	n.size = 1
	if n.left != nil {
		n.size += n.left.size
	}
	if n.right != nil {
		n.size += n.right.size
	}
}

func balanceFactor[K cmp.Ordered, V any](n *Node[K, V]) int {
	return height(n.left) - height(n.right)
}

func rebalance[K cmp.Ordered, V any](n *Node[K, V]) *Node[K, V] {
	updateSize(n)
	balance := balanceFactor(n)

	if balance > 1 {
		if balanceFactor(n.left) < 0 {
			n.left = rotateLeft(n.left)
		}
		return rotateRight(n)
	} else if balance < -1 {
		if balanceFactor(n.right) > 0 {
			n.right = rotateRight(n.right)
		}
		return rotateLeft(n)
	}
	return n
}

func rotateLeft[K cmp.Ordered, V any](z *Node[K, V]) *Node[K, V] {
	y := z.right
	z.right = y.left
	if y.left != nil {
		y.left.parent = z
	}
	y.left = z
	y.parent = z.parent
	z.parent = y
	updateSize(z)
	updateSize(y)
	return y
}

func rotateRight[K cmp.Ordered, V any](z *Node[K, V]) *Node[K, V] {
	y := z.left
	z.left = y.right
	if y.right != nil {
		y.right.parent = z
	}
	y.right = z
	y.parent = z.parent
	z.parent = y
	updateSize(z)
	updateSize(y)
	return y
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minNode[K cmp.Ordered, V any](n *Node[K, V]) *Node[K, V] {
	for n.left != nil {
		n = n.left
	}
	return n
}

func maxNode[K cmp.Ordered, V any](n *Node[K, V]) *Node[K, V] {
	for n.right != nil {
		n = n.right
	}
	return n
}
