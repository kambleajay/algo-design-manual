package temp

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestDailyTemperatures(t *testing.T) {
	var tests = []struct {
		temperatures []int
		want         []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{[]int{30, 60, 90}, []int{1, 1, 0}},
		{[]int{50, 49, 48, 47, 46}, []int{0, 0, 0, 0, 0}},
		{[]int{}, []int{}},
		{[]int{70, 69}, []int{0, 0}},
		{[]int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}, []int{8, 1, 5, 4, 3, 2, 1, 1, 0, 0}},
	}
	for _, test := range tests {
		if got := dailyTemperatures3(test.temperatures); !reflect.DeepEqual(got, test.want) {
			t.Errorf("dailyTemperatures(%v) = %v (should be %v)\n", test.temperatures, got, test.want)
		}
	}
}

func BenchmarkDailyTemperatures(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < b.N; i++ {
		noOfTemps := rng.Intn(100001)
		temps := make([]int, noOfTemps)
		for i := 0; i < len(temps); i++ {
			temps[i] = rng.Intn(100-30) + 30
		}
		dailyTemperatures(temps)
	}
}

func TestNode(t *testing.T) {
	tree := RedBlackBST{}
	tree.Put(10, 10)
	tree.Put(23, 23)
	tree.Put(47, 47)
	tree.Put(5, 5)
	tree.Put(39, 39)
	fmt.Printf("tree = %v\n", tree)
	a, ok := tree.Get(47)
	if !ok || !reflect.DeepEqual(a, []int{47}) {
		t.Errorf("got: %d, want: %d\n", a, 47)
	}
	b, ok := tree.Get(22)
	if ok {
		t.Errorf("got: %d, it should be 0\n", b)
	}
}

func TestPut(t *testing.T) {
	tree := RedBlackBST{}
	tree.Put(10, 10)
	tree.Put(23, 23)
	tree.Put(57, 57)
	fmt.Printf("root: %s\n", tree.root)
	fmt.Printf("left: %s\n", tree.root.left)
	fmt.Printf("right: %s\n", tree.root.right)
	fmt.Printf("tree: %s\n", &tree)
}

func TestNodes(t *testing.T) {
	tree := RedBlackBST{}
	tree.Put(10, 10)
	tree.Put(23, 23)
	tree.Put(57, 57)
	tree.Put(25, 25)
	tree.Put(33, 33)
	tree.Put(42, 42)
	fmt.Printf("tree: %s\n", &tree)
	fmt.Printf("nodes: %v\n", tree.nodes(20, 30))
}
