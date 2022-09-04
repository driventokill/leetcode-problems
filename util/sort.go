package util

type IntSlices [][]int

func (slices IntSlices) Len() int {
	return len(slices)
}

func (slices IntSlices) Swap(i, j int) {
	slices[i], slices[j] = slices[j], slices[i]
}

func (slices IntSlices) Less(i, j int) bool {
	if len(slices[i]) != len(slices[j]) {
		panic("slice lens not equal.")
	}

	for idx := range slices[i] {
		if slices[i][idx] < slices[j][idx] {
			return true
		}
	}
	return false
}
