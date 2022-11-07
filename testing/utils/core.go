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

func ContainsAllIntSlice(xs, ys []int) bool {
	if len(xs) != len(ys) {
		return false
	}
	sort.Sort(IntSlice(xs))
	sort.Sort(IntSlice(ys))
	return reflect.DeepEqual(xs, ys)
}

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func ContainsAllStringSlice(xs, ys []string) bool {
	if len(xs) != len(ys) {
		return false
	}
	sort.Sort(StringSlice(xs))
	sort.Sort(StringSlice(ys))
	return reflect.DeepEqual(xs, ys)
}

func Contains(a []int, x int) bool {
	for _, elem := range a {
		if elem == x {
			return true
		}
	}
	return false
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type String2DSlice [][]string

func (xs String2DSlice) Len() int { return len(xs) }
func (xs String2DSlice) Less(i, j int) bool {
	if len(xs[i]) != len(xs[j]) {
		return len(xs[i]) < len(xs[j])
	}
	for m := 0; m < len(xs[i]); m++ {
		if xs[i][m] != xs[j][m] {
			return xs[i][m] < xs[j][m]
		}
	}
	return false
}
func (xs String2DSlice) Swap(i, j int) { xs[i], xs[j] = xs[j], xs[i] }

func Equals(xs, ys [][]string) bool {
	sort.Sort(String2DSlice(xs))
	sort.Sort(String2DSlice(ys))
	for _, x := range xs {
		sort.Slice(x, func(i, j int) bool {
			return x[i] < x[j]
		})
	}
	for _, y := range ys {
		sort.Slice(y, func(i, j int) bool {
			return y[i] < y[j]
		})
	}
	return reflect.DeepEqual(xs, ys)
}
