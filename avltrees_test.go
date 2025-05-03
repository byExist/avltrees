package avltrees_test

import (
	"fmt"
	"math/rand"
	"testing"

	avlts "github.com/byExist/avltrees"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	tree := avlts.New[int, string]()
	assert.Nil(t, tree.Root, "New tree should have nil Root")
	assert.Equal(t, 0, avlts.Len(tree), "New tree should have size 0")
}

func TestInsert(t *testing.T) {
	tree := avlts.New[int, string]()

	inserted := avlts.Insert(tree, 10, "TEN")
	assert.True(t, inserted, "Expected first insert of 10 to return true")
	inserted = avlts.Insert(tree, 10, "ten")
	assert.False(t, inserted, "Expected second insert of 10 to return false (overwrite)")
	avlts.Insert(tree, 20, "twenty")
	avlts.Insert(tree, 5, "five")

	assert.Equal(t, 3, avlts.Len(tree), "Expected size 3")

	node, found := avlts.Search(tree, 10)
	require.True(t, found, "Insert failed for key 10")
	assert.Equal(t, "ten", node.Value())
}

func TestDelete(t *testing.T) {
	tree := avlts.New[int, string]()
	avlts.Insert(tree, 10, "ten")
	avlts.Insert(tree, 20, "twenty")
	avlts.Insert(tree, 5, "five")

	avlts.Delete(tree, 10)
	assert.Equal(t, 2, avlts.Len(tree), "Expected size 2 after deletion")

	_, found := avlts.Search(tree, 10)
	assert.False(t, found, "Key 10 should have been deleted")
}

func TestSearch(t *testing.T) {
	tree := avlts.New[int, string]()
	avlts.Insert(tree, 10, "ten")
	avlts.Insert(tree, 20, "twenty")

	node, found := avlts.Search(tree, 10)
	require.True(t, found, "Search failed for existing key 10")
	assert.Equal(t, "ten", node.Value())

	_, found = avlts.Search(tree, 30)
	assert.False(t, found, "Search should fail for non-existent key 30")
}

func TestInOrder(t *testing.T) {
	tree := avlts.New[int, string]()
	values := []int{20, 10, 30, 5, 15, 25, 35}
	for _, v := range values {
		avlts.Insert(tree, v, "")
	}

	prev := -1
	for n := range avlts.InOrder(tree) {
		if prev != -1 {
			assert.Less(t, prev, n.Key(), "InOrder traversal is not sorted")
		}
		prev = n.Key()
	}
}

func TestCeiling(t *testing.T) {
	tree := avlts.New[int, string]()
	for _, v := range []int{10, 20, 30} {
		avlts.Insert(tree, v, "")
	}

	n, ok := avlts.Ceiling(tree, 5)
	require.True(t, ok)
	assert.Equal(t, 10, n.Key())

	n, ok = avlts.Ceiling(tree, 20)
	require.True(t, ok)
	assert.Equal(t, 20, n.Key())

	_, ok = avlts.Ceiling(tree, 40)
	assert.False(t, ok)
}

func TestFloor(t *testing.T) {
	tree := avlts.New[int, string]()
	for _, v := range []int{10, 20, 30} {
		avlts.Insert(tree, v, "")
	}

	_, ok := avlts.Floor(tree, 5)
	assert.False(t, ok)

	n, ok := avlts.Floor(tree, 15)
	require.True(t, ok)
	assert.Equal(t, 10, n.Key())

	n, ok = avlts.Floor(tree, 20)
	require.True(t, ok)
	assert.Equal(t, 20, n.Key())
}

func TestHigher(t *testing.T) {
	tree := avlts.New[int, string]()
	for _, v := range []int{10, 20, 30} {
		avlts.Insert(tree, v, "")
	}

	n, ok := avlts.Higher(tree, 10)
	require.True(t, ok)
	assert.Equal(t, 20, n.Key())

	_, ok = avlts.Higher(tree, 35)
	assert.False(t, ok)
}

func TestLower(t *testing.T) {
	tree := avlts.New[int, string]()
	for _, v := range []int{10, 20, 30} {
		avlts.Insert(tree, v, "")
	}

	_, ok := avlts.Lower(tree, 5)
	assert.False(t, ok)

	n, ok := avlts.Lower(tree, 15)
	require.True(t, ok)
	assert.Equal(t, 10, n.Key())

	n, ok = avlts.Lower(tree, 30)
	require.True(t, ok)
	assert.Equal(t, 20, n.Key())
}

func TestRange(t *testing.T) {
	tree := avlts.New[int, string]()
	for _, v := range []int{10, 20, 30, 40, 50} {
		avlts.Insert(tree, v, "")
	}

	var collected []int
	for n := range avlts.Range(tree, 15, 45) {
		collected = append(collected, n.Key())
	}

	expected := []int{20, 30, 40}
	assert.Equal(t, len(expected), len(collected), "Expected range length")

	for i, v := range expected {
		assert.Equal(t, v, collected[i], "Expected value at position")
	}
}

func TestRank(t *testing.T) {
	tree := avlts.New[int, string]()
	values := []int{10, 20, 30, 40, 50}
	for _, v := range values {
		avlts.Insert(tree, v, "")
	}

	assert.Equal(t, 2, avlts.Rank(tree, 25))
	assert.Equal(t, 0, avlts.Rank(tree, 10))
	assert.Equal(t, 5, avlts.Rank(tree, 60))
}

func TestKth(t *testing.T) {
	tree := avlts.New[int, string]()
	values := []int{10, 20, 30, 40, 50}
	for _, v := range values {
		avlts.Insert(tree, v, "")
	}

	n, ok := avlts.Kth(tree, 0)
	require.True(t, ok)
	assert.Equal(t, 10, n.Key())

	n, ok = avlts.Kth(tree, 3)
	require.True(t, ok)
	assert.Equal(t, 40, n.Key())

	_, ok = avlts.Kth(tree, 5)
	assert.False(t, ok)
}

func TestPredecessor(t *testing.T) {
	tree := avlts.New[int, string]()
	for _, v := range []int{10, 20, 30, 40, 50} {
		avlts.Insert(tree, v, "")
	}

	n, found := avlts.Search(tree, 30)
	require.True(t, found)
	pred, ok := avlts.Predecessor(n)
	require.True(t, ok)
	assert.Equal(t, 20, pred.Key())
}

func TestSuccessor(t *testing.T) {
	tree := avlts.New[int, string]()
	for _, v := range []int{10, 20, 30, 40, 50} {
		avlts.Insert(tree, v, "")
	}

	n, found := avlts.Search(tree, 30)
	require.True(t, found)
	succ, ok := avlts.Successor(n)
	require.True(t, ok)
	assert.Equal(t, 40, succ.Key())
}

func TestMin(t *testing.T) {
	tree := avlts.New[int, string]()
	for _, v := range []int{20, 10, 30} {
		avlts.Insert(tree, v, "")
	}

	m, ok := avlts.Min(tree)
	require.True(t, ok)
	assert.Equal(t, 10, m.Key())
}

func TestMax(t *testing.T) {
	tree := avlts.New[int, string]()
	for _, v := range []int{20, 10, 30} {
		avlts.Insert(tree, v, "")
	}

	m, ok := avlts.Max(tree)
	require.True(t, ok)
	assert.Equal(t, 30, m.Key())
}

func TestLen(t *testing.T) {
	tree := avlts.New[int, string]()
	assert.Equal(t, 0, avlts.Len(tree))
	avlts.Insert(tree, 1, "")
	avlts.Insert(tree, 2, "")
	avlts.Insert(tree, 3, "")
	assert.Equal(t, 3, avlts.Len(tree))
}

func ExampleNew() {
	tree := avlts.New[int, string]()
	fmt.Println(avlts.Len(tree))
	// Output: 0
}

func ExampleInsert() {
	tree := avlts.New[int, string]()
	avlts.Insert(tree, 10, "ten")
	avlts.Insert(tree, 5, "five")
	avlts.Insert(tree, 15, "fifteen")
	fmt.Println(avlts.Len(tree))
	// Output: 3
}

func ExampleDelete() {
	tree := avlts.New[int, string]()
	avlts.Insert(tree, 10, "ten")
	avlts.Delete(tree, 10)
	fmt.Println(avlts.Len(tree))
	// Output: 0
}

func ExampleSearch() {
	tree := avlts.New[int, string]()
	avlts.Insert(tree, 20, "twenty")
	node, found := avlts.Search(tree, 20)
	fmt.Println(found, node.Value())
	// Output: true twenty
}

func ExampleMin() {
	tree := avlts.New[int, string]()
	avlts.Insert(tree, 20, "")
	avlts.Insert(tree, 10, "")
	min, ok := avlts.Min(tree)
	if ok {
		fmt.Println(min.Key())
	}
	// Output: 10
}

func ExampleMax() {
	tree := avlts.New[int, string]()
	avlts.Insert(tree, 20, "")
	avlts.Insert(tree, 30, "")
	max, ok := avlts.Max(tree)
	if ok {
		fmt.Println(max.Key())
	}
	// Output: 30
}

func ExampleCeiling() {
	tree := avlts.New[int, string]()
	for _, v := range []int{10, 20, 30} {
		avlts.Insert(tree, v, "")
	}

	n, _ := avlts.Ceiling(tree, 15)
	fmt.Println(n.Key())
	n, _ = avlts.Ceiling(tree, 20)
	fmt.Println(n.Key())
	// Output:
	// 20
	// 20
}

func ExampleFloor() {
	tree := avlts.New[int, string]()
	for _, v := range []int{10, 20, 30} {
		avlts.Insert(tree, v, "")
	}

	n, _ := avlts.Floor(tree, 25)
	fmt.Println(n.Key())
	n, _ = avlts.Floor(tree, 20)
	fmt.Println(n.Key())
	// Output:
	// 20
	// 20
}

func ExampleHigher() {
	tree := avlts.New[int, string]()
	for _, v := range []int{10, 20, 30} {
		avlts.Insert(tree, v, "")
	}

	n, _ := avlts.Higher(tree, 15)
	fmt.Println(n.Key())
	n, _ = avlts.Higher(tree, 20)
	fmt.Println(n.Key())
	// Output:
	// 20
	// 30
}

func ExampleLower() {
	tree := avlts.New[int, string]()
	for _, v := range []int{10, 20, 30} {
		avlts.Insert(tree, v, "")
	}

	n, _ := avlts.Lower(tree, 15)
	fmt.Println(n.Key())
	n, _ = avlts.Lower(tree, 20)
	fmt.Println(n.Key())
	// Output:
	// 10
	// 10
}

func ExampleInOrder() {
	tree := avlts.New[int, string]()
	avlts.Insert(tree, 20, "")
	avlts.Insert(tree, 10, "")
	avlts.Insert(tree, 30, "")
	for n := range avlts.InOrder(tree) {
		fmt.Print(n.Key(), " ")
	}
	fmt.Println()
	// Output: 10 20 30
}

func ExampleRange() {
	tree := avlts.New[int, string]()
	avlts.Insert(tree, 10, "")
	avlts.Insert(tree, 20, "")
	avlts.Insert(tree, 30, "")
	for n := range avlts.Range(tree, 15, 25) {
		fmt.Print(n.Key(), " ")
	}
	fmt.Println()
	// Output: 20
}

func ExamplePredecessor() {
	tree := avlts.New[int, string]()
	avlts.Insert(tree, 20, "")
	avlts.Insert(tree, 10, "")
	node, _ := avlts.Search(tree, 20)
	pred, ok := avlts.Predecessor(node)
	if ok {
		fmt.Println(pred.Key())
	}
	// Output: 10
}

func ExampleSuccessor() {
	tree := avlts.New[int, string]()
	avlts.Insert(tree, 20, "")
	avlts.Insert(tree, 30, "")
	node, _ := avlts.Search(tree, 20)
	succ, ok := avlts.Successor(node)
	if ok {
		fmt.Println(succ.Key())
	}
	// Output: 30
}

func ExampleRank() {
	tree := avlts.New[int, string]()
	avlts.Insert(tree, 10, "")
	avlts.Insert(tree, 20, "")
	rank := avlts.Rank(tree, 15)
	fmt.Println(rank)
	// Output: 1
}

func ExampleKth() {
	tree := avlts.New[int, string]()
	avlts.Insert(tree, 10, "")
	avlts.Insert(tree, 20, "")
	n, _ := avlts.Kth(tree, 1)
	fmt.Println(n.Key())
	// Output: 20
}

func ExampleLen() {
	tree := avlts.New[int, string]()
	fmt.Println(avlts.Len(tree))
	// Output: 0
}

func BenchmarkInsertRandom(b *testing.B) {
	r := rand.New(rand.NewSource(42))
	tree := avlts.New[int, string]()
	keys := make([]int, b.N)
	for i := range keys {
		keys[i] = r.Intn(1_000_000)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		avlts.Insert(tree, keys[i], "value")
	}
}

func BenchmarkInsertSequential(b *testing.B) {
	tree := avlts.New[int, string]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		avlts.Insert(tree, i, "value")
	}
}

func BenchmarkSearchHit(b *testing.B) {
	tree := avlts.New[int, string]()
	for i := 0; i < 1000; i++ {
		avlts.Insert(tree, i, "value")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		avlts.Search(tree, i%1000)
	}
}

func BenchmarkSearchMiss(b *testing.B) {
	tree := avlts.New[int, string]()
	for i := range 1000 {
		avlts.Insert(tree, i, "value")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		avlts.Search(tree, 1_000_000+i)
	}
}

func BenchmarkDeleteRandom(b *testing.B) {
	r := rand.New(rand.NewSource(42))
	tree := avlts.New[int, string]()
	keys := make([]int, 1000)
	for i := range 1000 {
		keys[i] = r.Intn(1_000_000)
		avlts.Insert(tree, keys[i], "value")
	}
	perm := r.Perm(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		avlts.Delete(tree, keys[perm[i%1000]])
	}
}
