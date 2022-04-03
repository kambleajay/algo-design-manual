package utils

import (
	"reflect"
	"sort"
)

type Int2DSlice [][]int

func (xs Int2DSlice) Len() int { return len(xs) }
func (xs Int2DSlice) Less(i, j int) bool {
	for a, b := 0, 0; a < len(xs[i]) && b < len(xs[j]); a, b = a+1, b+1 {
		if xs[i][a] != xs[j][b] {
			return xs[i][a] < xs[j][b]
		}
	}
	return false
}
func (xs Int2DSlice) Swap(i, j int) { xs[i], xs[j] = xs[j], xs[i] }

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func ContainsAll(xs, ys [][]int) bool {
	if len(xs) != len(ys) {
		return false
	}
	for _, x := range xs {
		sort.Sort(IntSlice(x))
	}
	for _, y := range ys {
		sort.Sort(IntSlice(y))
	}
	sort.Sort(Int2DSlice(xs))
	sort.Sort(Int2DSlice(ys))
	return reflect.DeepEqual(xs, ys)
}
