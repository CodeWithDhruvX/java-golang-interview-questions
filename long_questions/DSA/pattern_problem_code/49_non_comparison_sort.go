package main

import (
	"fmt"
	"math"
)

// Pattern: Non-Comparison Sort (Bucket Sort)
// Difficulty: Hard
// Key Concept: Sorting in O(N) by exploiting value ranges. Pigeonhole Principle usually guarantees gaps.

/*
INTUITION:
We have N numbers. Range is [Min, Max].
If distributed evenly, gap is `(Max - Min) / (N - 1)`.
If not even, some gap must be larger than this average gap.
We create buckets of size `gap`.
All elements in the same bucket are within `gap` distance, so the Max Gap cannot be inside a bucket.
It must be between the MAX of Bucket `i` and MIN of Bucket `j` (where j > i and buckets between are empty).

PROBLEM:
LeetCode 164. Maximum Gap.
Given an integer array nums, return the maximum difference between two successive elements in its sorted form.
You must write an algorithm that runs in linear time and uses linear extra space.

ALGORITHM:
1. Find `minVal` and `maxVal`.
2. `bucketSize = max(1, (maxVal - minVal) / (n - 1))`.
3. `bucketCount = (maxVal - minVal) / bucketSize + 1`.
4. Create buckets. Each stores `min` and `max` observed in it. Init with infinity/-infinity.
5. Place numbers into buckets: `idx = (num - minVal) / bucketSize`. Update bucket's min/max.
6. Iterate buckets. `maxGap = max(maxGap, currBucket.min - prevBucket.max)`.
*/

type Bucket struct {
	minVal int
	maxVal int
	used   bool
}

func maximumGap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	minVal, maxVal := nums[0], nums[0]
	for _, v := range nums {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}

	if minVal == maxVal {
		return 0
	}

	bucketSize := (maxVal - minVal) / (n - 1)
	if bucketSize == 0 {
		bucketSize = 1
	}
	bucketCount := (maxVal-minVal)/bucketSize + 1

	buckets := make([]Bucket, bucketCount)
	for i := range buckets {
		buckets[i] = Bucket{minVal: math.MaxInt32, maxVal: math.MinInt32, used: false}
	}

	for _, v := range nums {
		idx := (v - minVal) / bucketSize
		buckets[idx].used = true
		if v < buckets[idx].minVal {
			buckets[idx].minVal = v
		}
		if v > buckets[idx].maxVal {
			buckets[idx].maxVal = v
		}
	}

	maxGap := 0
	prevMax := minVal

	for _, b := range buckets {
		if !b.used {
			continue
		}
		// Gap is between current bucket's min and previous bucket's max
		gap := b.minVal - prevMax
		if gap > maxGap {
			maxGap = gap
		}
		prevMax = b.maxVal
	}

	return maxGap
}

func main() {
	// [3, 6, 9, 1]
	// Sorted: 1, 3, 6, 9. Gaps: 2, 3, 3. Max Gap 3.
	// Min 1, Max 9. N=4. Avg Gap (8/3) = 2.
	// Bucket Size 2.
	// Bucket 0 [1, 2]: Finds 1. Min 1, Max 1.
	// Bucket 1 [3, 4]: Finds 3. Min 3, Max 3.
	// Bucket 2 [5, 6]: Finds 6. Min 6, Max 6.
	// Bucket 3 [7, 8]: Empty.
	// Bucket 4 [9, 9]: Finds 9. Min 9, Max 9.

	// Pass 1: Bucket 0 (Max 1). PrevMax=1.
	// Pass 2: Bucket 1 (Min 3). Gap 3-1=2. PrevMax=3. MaxGap=2.
	// Pass 3: Bucket 2 (Min 6). Gap 6-3=3. PrevMax=6. MaxGap=3.
	// Pass 4: Bucket 3 (Used=false). Skip.
	// Pass 5: Bucket 4 (Min 9). Gap 9-6=3. PrevMax=9. MaxGap=3.

	nums := []int{3, 6, 9, 1}
	fmt.Printf("Nums: %v, MaxGap: %d\n", nums, maximumGap(nums))

	// [10]
	// 0.
	fmt.Printf("Nums: %v, MaxGap: %d\n", []int{10}, maximumGap([]int{10}))
}
