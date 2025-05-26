# avltrees [![GoDoc](https://pkg.go.dev/badge/github.com/byExist/avltrees)](https://pkg.go.dev/github.com/byExist/avltrees) [![Go Report Card](https://goreportcard.com/badge/github.com/byExist/avltrees)](https://goreportcard.com/report/github.com/byExist/avltrees)

A generic self-balancing AVL tree for Go with rank, k-th, and range queries.

---

## ✨ Features

- Self-balancing AVL Tree with O(log n) insert/delete/search
- Order-statistics support: rank, k-th element, and range queries
- Generic (Go generics)
- In-order iterator
- Memory-efficient (no extra allocations on lookup)

---

## ✅ Use When

- You need **fast and predictable search performance**
- You want **rank or k-th queries**
- Your workload is **read-heavy** or search-dominant

---

## 🚫 Avoid If

- You need **frequent insertions and deletions** with minimal balancing overhead → use [Red-Black Tree](https://github.com/byExist/redblacktrees)
- You need **concurrent** access (not thread-safe)

---

## 📦 Installation

```bash
go get github.com/byExist/avltrees
```

---

## 🚀 Quick Start

```go
package main

import (
	"fmt"
	avlt "github.com/byExist/avltrees"
)

func main() {
	tree := avlt.New[int, string]()
	avlt.Insert(tree, 10, "ten")
	avlt.Insert(tree, 5, "five")
	avlt.Insert(tree, 20, "twenty")

	if node, found := avlt.Search(tree, 10); found {
		fmt.Println("Found:", node.Value())
	}

	for n := range avlt.InOrder(tree) {
		fmt.Println(n.Key(), "->", n.Value())
	}
}
```

---

## 📊 Performance

Benchmarked on Apple M1 Pro:

| Operation            | Time (ns/op) | Memory (B/op) | Allocations |
|---------------------|--------------|----------------|-------------|
| Insert (Random)     | 669.3        | 26 B           | 0           |
| Insert (Sequential) | 151.6        | 64 B           | 1           |
| Search (Hit)        | 11.52        | 0 B            | 0           |
| Search (Miss)       | 7.58         | 0 B            | 0           |
| Delete (Random)     | 2.90         | 0 B            | 0           |

---

## 🪪 License

MIT License. See [LICENSE](LICENSE).