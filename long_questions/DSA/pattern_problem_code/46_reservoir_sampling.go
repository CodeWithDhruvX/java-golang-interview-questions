package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Pattern: Reservoir Sampling
// Difficulty: Medium
// Key Concept: Selecting k random elements from a stream of unknown size N in O(N) time and O(k) space, ensuring equal probability.

/*
INTUITION:
You are fishing in a river. You don't know how many fish are there. You want to catch ONE random fish.
- Fish 1 comes. Keep it. (Prob 1/1)
- Fish 2 comes. 50/50 chance to swap. (Prob 1/2).
- Fish 3 comes. 1/3 chance to swap.
  - Prob(Fish 3 kept) = 1/3.
  - Prob(Fish 2 kept) = (Kept at step 2 [1/2]) * (Not swapped at step 3 [2/3]) = 1/3.
  - Prob(Fish 1 kept) = (Kept at step 2 [1/2]) * (Not swapped at step 3 [2/3]) = 1/3.
All have equal probability.

PROBLEM:
LeetCode 382. Linked List Random Node.
Given a singly linked list, return a random node's value from the linked list. Each node must have the same probability of being chosen.
What if the linked list is extremely large and its length is unknown to you? Could you solve this efficiently without using extra space?

ALGORITHM:
1. `Solution(head)`: Initialize.
2. `getRandom()`:
   - `scope = 1`, `chosen = head.val`.
   - `curr = head.next`.
   - While `curr != nil`:
     - `scope++`.
     - `rand_idx = rand.Intn(scope)`.
     - If `rand_idx == 0` (Prob 1/scope), `chosen = curr.val`.
     - `curr = curr.next`.
   - Return `chosen`.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}

func ConstructorRS(head *ListNode) Solution {
	rand.Seed(time.Now().UnixNano())
	return Solution{head: head}
}

func (this *Solution) GetRandom() int {
	scope := 1
	chosenValue := 0
	curr := this.head

	// Safety check
	if curr != nil {
		chosenValue = curr.Val
	}

	for curr != nil {
		// Decision: Pick current node with probability 1/scope
		if rand.Intn(scope) == 0 {
			chosenValue = curr.Val
		}
		scope++
		curr = curr.Next
	}
	return chosenValue
}

func main() {
	// List: 1 -> 2 -> 3
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}

	obj := ConstructorRS(head)

	// Test randomness (should be roughly equal distribution)
	counts := make(map[int]int)
	for i := 0; i < 3000; i++ {
		val := obj.GetRandom()
		counts[val]++
	}
	fmt.Printf("Counts (approx 1000 each): %v\n", counts)
}
