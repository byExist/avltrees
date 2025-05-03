

# avltrees

## What is "avltrees"?

`avltrees` is a generic AVL Tree implementation in Go. It supports efficient insertion, deletion, and search operations, all with O(log n) time complexity. The implementation also includes advanced operations such as rank, k-th smallest element, range queries, and predecessor/successor navigation.

## Features

- Generic support using Go generics
- Self-balancing AVL Tree with O(log n) insert/delete/search
- Rank and k-th element queries
- Range iteration with inclusive/exclusive bounds
- In-order traversal using `iter.Seq`
- Tree size maintained at each node

## Installation

```bash
go get github.com/yourusername/avltrees
```

## Quick Start

```go
package main

import (
	"fmt"
	avlt "github.com/yourusername/avltrees"
)

func main() {
	tree := avlt.New[int, string]()
	avlt.Insert(tree, 10, "ten")
	avlt.Insert(tree, 5, "five")
	avlt.Insert(tree, 20, "twenty")

	node, found := avlt.Search(tree, 10)
	if found {
		fmt.Println("Found:", node.Value())
	}

	avlt.Delete(tree, 5)

	for n := range avlt.InOrder(tree) {
		fmt.Println(n.Key(), "->", n.Value())
	}
}
```

## API Overview

### Core Types

- `Tree[K, V]`: The main AVL tree structure
- `Node[K, V]`: A node in the tree

### Key Functions

- `New[K cmp.Ordered, V any]() *Tree[K, V]`
- `Insert(t *Tree[K, V], key K, value V) bool`
- `Delete(t *Tree[K, V], key K) bool`
- `Search(t *Tree[K, V], key K) (*Node[K, V], bool)`
- `Min(t *Tree[K, V]) (*Node[K, V], bool)`
- `Max(t *Tree[K, V]) (*Node[K, V], bool)`
- `Rank(t *Tree[K, V], key K) int`
- `Kth(t *Tree[K, V], k int) (*Node[K, V], bool)`
- `InOrder(t *Tree[K, V]) iter.Seq[Node[K, V]]`
- `Range(t *Tree[K, V], from, to K) iter.Seq[Node[K, V]]`

## Performance

Benchmark results measured on an Apple M1 Pro (macOS/arm64):

| Benchmark               | Iterations (N) | Time per op (ns/op) | Memory (B/op) | Allocations (allocs/op) |
|------------------------|----------------|----------------------|----------------|--------------------------|
| Insert (Random)        | 2,123,395      | 669.3                | 26 B           | 0                        |
| Insert (Sequential)    | 8,211,963      | 151.6                | 64 B           | 1                        |
| Search (Hit)           | 98,068,260     | 11.52                | 0 B            | 0                        |
| Search (Miss)          | 158,177,852    | 7.582                | 0 B            | 0                        |
| Delete (Random)        | 388,123,057    | 2.901                | 0 B            | 0                        |

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.