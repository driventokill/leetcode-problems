package util

import (
	"sort"

	"golang.org/x/exp/constraints"
)

func SortSlices[T constraints.Ordered](s [][]T) {
	sort.Slice(s, func(i, j int) bool {
		if len(s[i]) != len(s[j]) {
			panic("length of slice not euqal.")
		}
		for idx := range s[i] {
			if s[i][idx] < s[j][idx] {
				return true
			}
		}

		return false
	})
}