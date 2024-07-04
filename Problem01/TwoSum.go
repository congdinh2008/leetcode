package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 1, 1, 4, 1, 1, 1, 1, 1, 7, 1, 1, 1, 1, 1}
	target := 11

	result := TwoSumDict(nums, target)

	fmt.Printf("%d - %d\n", result[0], result[1])
}

func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i != j && nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func TwoSumDict(nums []int, target int) []int {
	myMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		temp := target - nums[i]
		if j, ok := myMap[temp]; ok {
			return []int{j, i}
		}
		myMap[nums[i]] = i
	}
	return nil
}
