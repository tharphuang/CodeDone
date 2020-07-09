package medium

import (
	"fmt"
	"sort"
	"strings"
)

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
func lengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		if index == -1 {
			if i+1 > end {
				end = i + 1
			}
		} else {
			start += index + 1
			end += index + 1
		}
	}
	return end - start

}

func TestLenOfLongestS() {
	fmt.Print(lengthOfLongestSubstring("pwwkew"))
}

//找出这两个正序数组的中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	s := append(nums1, nums2...)

	sort.Ints(s)

	mid := len(s) / 2
	if len(s)%2 == 0 {
		return float64(s[mid-1]+s[mid]) / 2
	} else {
		return float64(s[mid])
	}

}

func TestFindMedianSortedA() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	fmt.Print(findMedianSortedArrays(nums1, nums2))
}

func longestPalindrome(s string) string {
	max := 0
	var subS string
	for i := 0; i < len(s); i++ {
		for j := i; j <= len(s); j++ {
			if isPalindrome(s[i:j]) > max {
				max = isPalindrome(s[i:j])
				subS = s[i:j]
			}
		}

	}
	return subS

}

func isPalindrome(s string) int {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return 0
		}
	}
	return len(s)
}

func TestlongestP() {
	s := "addacdcdc"
	fmt.Print(longestPalindrome(s))
}
