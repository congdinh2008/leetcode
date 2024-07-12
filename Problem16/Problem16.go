package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	result := threeSumClosest(nums, target)

	fmt.Println(result)
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	sum := nums[0] + nums[1] + nums[2]
	minSum := sum
	min := target - sum

	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			currentMin := int(math.Abs(float64(target - sum)))

			if sum == target {
				return sum
			} else if sum < target {
				if currentMin < min {
					min = currentMin
					minSum = sum
				}
				left++
			} else {
				if currentMin < min {
					min = currentMin
					minSum = sum
				}
				right--
			}
		}
	}

	return minSum
}
