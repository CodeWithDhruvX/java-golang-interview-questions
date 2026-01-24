package main

import (
	"fmt"
)

// Pattern: Binary Search on Answer Space
// Difficulty: Medium/Hard
// Key Concept: We can't search the array directly, but we can search for the *Answer* (Result) because the answer is monotonic.
// "Can you do it in X hours?" -> If Yes, I can also do it in X+1. If No, I can't do it in X-1. This Monotonicity allows Binary Search.

/*
INTUITION:
Problem: "Koko Eating Bananas".
Koko loves bananas. There are N piles.
Koko must finish all piles within H hours.
What is the MINIMUM eating speed (K bananas/hour) she needs?

Think:
- Minimum possible speed? 1 banana/hour (Slowest).
- Maximum possible speed? Eating the biggest pile in 1 hour (Fastest).
- Range of possible answers: [1, Max(Piles)].
- Is this range sorted? YES. 1, 2, 3, 4, ...
- Can we "Check" if a speed works? Yes. Loop through piles and calculate hours.

ALGORITHM:
1. Define Search Space: `low = 1`, `high = max(piles)`.
2. Binary Search:
   - `mid = (low + high) / 2`. (Candidate Speed).
   - Call helper `canEatAll(speed, piles, H)`.
   - If `canEatAll` is True: This speed works! But maybe we can be slower? We want the *Minimum*. Record answer, try slower (`high = mid - 1`).
   - If `canEatAll` is False: Too slow! We failed. We must eat faster. (`low = mid + 1`).
3. Return best answer found.
*/

func minEatingSpeed(piles []int, h int) int {
	// Find max element for High bound
	maxPile := 0
	for _, p := range piles {
		if p > maxPile {
			maxPile = p
		}
	}

	low := 1
	high := maxPile
	result := maxPile

	// DRY RUN:
	// Piles: [3, 6, 7, 11], H=8
	// Low=1, High=11.
	//
	// mid = 6. Can we eat in speed 6?
	//   3 -> 1 hr
	//   6 -> 1 hr
	//   7 -> 2 hrs (6 + 1)
	//   11 -> 2 hrs (6 + 5)
	//   Total = 6 hrs.
	//   6 <= 8 (H). YES! Speed 6 works.
	//   Try slower? High = 5. Result = 6.
	//
	// mid = 3. Can we eat in speed 3?
	//   3 -> 1 hr
	//   6 -> 2 hrs
	//   7 -> 3 hrs
	//   11 -> 4 hrs
	//   Total = 10 hrs.
	//   10 > 8 (H). NO! Too slow.
	//   Must go faster. Low = 4.

	for low <= high {
		mid := low + (high-low)/2 // This is our candidate speed K

		if canFinish(piles, h, mid) {
			result = mid
			high = mid - 1 // Try smaller speed
		} else {
			low = mid + 1 // Need higher speed
		}
	}

	return result
}

// Helper: Can Koko finish all piles in H hours with speed K?
func canFinish(piles []int, h int, speed int) bool {
	hoursUsed := 0
	for _, pile := range piles {
		// Calculate ceiling division: (pile + speed - 1) / speed
		// E.g. 7 bananas, speed 3.  (7+2)/3 = 3 hours.
		hoursUsed += (pile + speed - 1) / speed
	}
	return hoursUsed <= h
}

func main() {
	piles := []int{3, 6, 7, 11}
	h := 8

	fmt.Printf("Piles: %v, Hours: %d\n", piles, h)
	k := minEatingSpeed(piles, h)
	fmt.Printf("Min Eating Speed: %d\n", k) // Expected: 4
}
