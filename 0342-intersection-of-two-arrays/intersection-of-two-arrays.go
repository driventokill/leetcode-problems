package intersectionOfTwoArrays

func intersection(nums1 []int, nums2 []int) []int {
	var result []int
	m1 := toMap(nums1)
	m2 := toMap(nums2)
	for k := range m1 {
		if _, ok := m2[k]; ok {
			result = append(result, k)
		}
	}
	return result
}

func toMap(nums []int) map[int]bool {
	m := make(map[int]bool)
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			m[v] = true
		}
	}
	return m
}
