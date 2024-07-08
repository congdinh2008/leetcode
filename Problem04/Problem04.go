package main

import (
	"fmt"
)

func main() {
	nums1 := []int{3, 4}
	nums2 := []int{} // 1,2,3,4,5,9

	fmt.Println(findMedianSortedArraysTwoPointer(nums1, nums2))
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	resultLen := len1 + len2
	result := make([]int, resultLen)

	i := 0
	j := 0
	k := 0
	for i < len1 || j < len2 {
		if i < len1 && j < len2 && nums1[i] <= nums2[j] {
			result[k] = nums1[i]
			k++
			i++
			continue
		}

		if i < len1 && j < len2 && nums1[i] >= nums2[j] {
			result[k] = nums2[j]
			k++
			j++
			continue
		}

		if i == len1 || j == len2 {
			break
		}
	}

	for i < len1 {
		result[k] = nums1[i]
		i++
	}

	for j < len2 {
		result[k] = nums2[j]
		k++
		j++
	}

	if resultLen%2 == 1 {
		return float64(result[resultLen/2])
	} else {
		return (float64(result[resultLen/2]) + float64(result[resultLen/2-1])) / 2
	}
}

func findMedianSortedArraysTwoPointer(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	total := len1 + len2

	median1 := 0
	median2 := 0

	i := 0
	j := 0
	for counter := 0; counter <= total/2; counter++ {
		median2 = median1
		if i == len1 || (j < len2 && nums1[i] >= nums2[j]) {
			median1 = nums2[j]
			j++
		} else {
			median1 = nums1[i]
			i++
		}
	}
	if total%2 == 1 {
		return float64(median1)
	} else {
		return (float64(median1) + float64(median2)) / 2
	}
}
