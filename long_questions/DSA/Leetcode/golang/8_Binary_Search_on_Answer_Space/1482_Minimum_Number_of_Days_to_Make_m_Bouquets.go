package main

import "fmt"

// 1482. Minimum Number of Days to Make m Bouquets
// Time: O(N log (max-min)), Space: O(1)
func minDays(bloomDay []int, m int, k int) int {
	if m*k > len(bloomDay) {
		return -1
	}
	
	left, right := minDay(bloomDay), maxDay(bloomDay)
	result := -1
	
	for left <= right {
		mid := left + (right-left)/2
		
		if canMakeBouquets(bloomDay, m, k, mid) {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	
	return result
}

func canMakeBouquets(bloomDay []int, m, k, days int) bool {
	bouquets := 0
	flowers := 0
	
	for _, day := range bloomDay {
		if day <= days {
			flowers++
			if flowers == k {
				bouquets++
				flowers = 0
				if bouquets >= m {
					return true
				}
			}
		} else {
			flowers = 0
		}
	}
	
	return bouquets >= m
}

func minDay(bloomDay []int) int {
	min := bloomDay[0]
	for _, day := range bloomDay {
		if day < min {
			min = day
		}
	}
	return min
}

func maxDay(bloomDay []int) int {
	max := bloomDay[0]
	for _, day := range bloomDay {
		if day > max {
			max = day
		}
	}
	return max
}

func main() {
	// Test cases
	testCases := []struct {
		bloomDay []int
		m        int
		k        int
	}{
		{[]int{1, 10, 3, 10, 2}, 3, 1},
		{[]int{1, 10, 3, 10, 2}, 3, 2},
		{[]int{7, 7, 7, 7, 12, 7, 7}, 2, 3},
		{[]int{1, 1, 1, 1}, 1, 1},
		{[]int{1, 1, 1, 1}, 4, 1},
		{[]int{1000000000, 1000000000}, 1, 1},
		{[]int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}, 4, 2},
		{[]int{5, 5, 5, 5, 5}, 3, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 2, 3},
	}
	
	for i, tc := range testCases {
		result := minDays(tc.bloomDay, tc.m, tc.k)
		fmt.Printf("Test Case %d: bloomDay=%v, m=%d, k=%d -> Min days: %d\n", 
			i+1, tc.bloomDay, tc.m, tc.k, result)
	}
}
