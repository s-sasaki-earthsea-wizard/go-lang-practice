package intersection02

// Finf the intersection of two arrays
func Intersect(nums1 []int, nums2 []int) []int {

	// Set nums1 as the shorter array
	if len(nums1) > len(nums2) {
		return Intersect(nums2, nums1)
	}

	// Count the number of occurrences of each element in nums1
	// key: num, value: count
	count := make(map[int]int) // key is int, value is also int
	for _, num := range nums1 {
		count[num]++
	}

	// Find the intersection
	k := 0
	result := make([]int, len(nums1))
	for _, num := range nums2 {
		if count[num] > 0 {
			result[k] = num
			k++
			count[num]--
		}
	}

	return result[:k] // return the first k elements, because the length of the result may be longer than k as: [4, 9, 0]
}
