package random_pick_blacklist

import (
	"algo/testing/utils"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSolution1(t *testing.T) {
	var tests = []struct {
		n         int
		blacklist []int
	}{
		{7, []int{2, 3, 5}},
		{1, []int{}},
	}
	for _, test := range tests {
		solution := Constructor(test.n, test.blacklist)
		results := make(map[int]int)
		for i := 0; i < 1000000; i++ {
			got := solution.Pick()
			if utils.Contains(test.blacklist, got) {
				t.Errorf("got a number from blacklist: %d, %v", got, test.blacklist)
			}
			_, ok := results[got]
			if !ok {
				results[got] = 1
			} else {
				results[got]++
			}
		}
		fmt.Printf("%v\n", results)
	}
}

func TestSolution2(t *testing.T) {
	var tests = []struct {
		n         int
		blacklist []int
	}{
		{7, []int{2, 3, 5}},
		{1, []int{}},
	}
	for _, test := range tests {
		solution := Constructor1(test.n, test.blacklist)
		results := make(map[int]int)
		for i := 0; i < 1000000; i++ {
			got := solution.Pick1()
			if utils.Contains(test.blacklist, got) {
				t.Errorf("got a number from blacklist: %d, %v", got, test.blacklist)
			}
			_, ok := results[got]
			if !ok {
				results[got] = 1
			} else {
				results[got]++
			}
		}
		fmt.Printf("%v\n", results)
	}
}

func BenchmarkSolution1(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	n := rng.Intn(1000000000-1) + 1
	blacklistMax := utils.Min(100000, n-1)
	blacklistLen := rng.Intn(blacklistMax)
	var blacklist []int
	for i := 0; i < blacklistLen; i++ {
		blacklist = append(blacklist, rng.Intn(n))
	}
	noOfCalls := rng.Intn((2*10000)-1) + 1
	for i := 0; i < b.N; i++ {
		solution := Constructor(n, blacklist)
		for j := 0; j < noOfCalls; j++ {
			solution.Pick()
		}
	}
}

func BenchmarkSolution2(b *testing.B) {
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))
	n := rng.Intn(1000000000-1) + 1
	blacklistMax := utils.Min(100000, n-1)
	blacklistLen := rng.Intn(blacklistMax)
	var blacklist []int
	for i := 0; i < blacklistLen; i++ {
		blacklist = append(blacklist, rng.Intn(n))
	}
	noOfCalls := rng.Intn((2*10000)-1) + 1
	for i := 0; i < b.N; i++ {
		solution := Constructor1(n, blacklist)
		for j := 0; j < noOfCalls; j++ {
			solution.Pick1()
		}
	}
}
