package main

import "fmt"

func firstNegInt(nums []int, k int) []int {
	result := []int{}
	low := 0
	high := 0
	n := len(nums)
	ans := []int{}
	for high < n {
		if nums[high] < 0 {
			ans = append(ans, nums[high])
		}

		if high-low+1 == k {
			if len(ans) != 0 {
				result = append(result, ans[0])
				if ans[0] == nums[low] {
					ans = ans[1:]
				}
			} else {
				result = append(result, 0)
			}
			low++
		}
		high++

	}
	return result
}

func main() {
	nums := []int{-8, 2, 3, -6, 10}
	k := 2
	result := firstNegInt(nums, k)
	for _, val := range result {
		fmt.Println(val)
	}
}
