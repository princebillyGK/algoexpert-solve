package main

func SolveV1(a []int64, sub []int64) bool {
	si := 0
	sli := len(sub) - 1
	for _, v := range a {
		if sub[si] == v {
			if si == sli {
				return true
			}
			si++
			continue
		}
	}
	return false
}
