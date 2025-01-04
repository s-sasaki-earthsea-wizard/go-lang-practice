package main

import "fmt"

// Plus one
func plusOne(digits []int) []int {
	n := len(digits)

	// Iterate from the end of the array
	for i := n - 1; i >= 0; i-- {
		// If the digit is less than 9, increment it and return the array
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		// If the digit is 9, set it to 0
		digits[i] = 0
	}

	// If all the digits are 9, create a new array with the first element as 1 and the rest as 0
	result := make([]int, n+1)
	result[0] = 1
	return result

}

func main() {
	// Arrays for testing
	nums1 := []int{1, 2, 3}
	result1 := plusOne(nums1)

	nums2 := []int{9, 9, 9}
	result2 := plusOne(nums2)

	fmt.Println(result1)
	fmt.Println(result2)
}
