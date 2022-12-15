package main

import (
	"errors"
	"sort"
)

// SolveV1 uses a map and formula y = 10 - x
// Time Complexity: O(n); Space Complexity: O(n)
func SolveV1(a []int, target int) (int, int, error) {
	if len(a) == 0 {
		return 0, 0, errors.New("No pairs found")
	}
	m := map[int]bool{a[0]: true}
	for _, x := range a[1:] {
		y := target - x
		ok, _ := m[y]
		if ok {
			return x, y, nil
		}
		m[x] = true
	}
	return 0, 0, errors.New("No pairs found")
}

// SolveV2 uses a left and right pointer inside a sorted array
// Time Complexity: O(nolan); Space Complexity: O(1)
func SolveV2(a []int, target int) (int, int, error) {
	sorted := make([]int, len(a))
	copy(sorted, a)
	sort.Ints(sorted)

	lp := 0
	rp := len(sorted) - 1

	for rp != lp {
		sum := sorted[rp] + sorted[lp]
		if sum == target {
			return sorted[rp], sorted[lp], nil
		} else if sum > target {
			rp--
		} else {
			lp++
		}
	}
	return 0, 0, errors.New("No pairs found")
}
