package main

import "math"

func minSubarrayLen(target int, nums []int) int {
	result := math.MaxInt
	i := 0
	j := 0
	n := len(nums)
	ans := 0
	for j < n {
		ans += nums[j]

		for ans >= target {
			result = int(math.Min(float64(result), float64(j-i+1)))
			ans -= nums[i]
			i++
		}
		j++
	}
	return result
}
