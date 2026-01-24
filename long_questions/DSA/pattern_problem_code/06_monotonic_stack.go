package main

import "fmt"

// Pattern: Monotonic Stack
// Difficulty: Medium
// Key Concept: Using a stack to find the "Next Greater" or "Next Smaller" element efficiently.

/*
INTUITION:
Imagine people standing in a line. You want to look to your right and find the first person taller than you.
If a short person stands behind a tall person, the short person is "blocked".
We can maintain a stack of people who are "waiting" to find someone taller.

When a new person comes:
- If new person is SHORTER than the guy on stack top: Just add him to the stack. He is also waiting.
- If new person is TALLER than the guy on stack top:
  - This new person is the "Next Greater Element" for the stack top guy!
  - Pop the stack top, record the answer.
  - Repeat until the stack top is taller than the new guy in (or stack empty).
  - Push the new person.

This ensures the stack is always sorted (Monotonic Decreasing).

PROBLEM:
"Next Greater Element"
Given an array, for each element, find the next greater element to its right. If none, -1.

ALGORITHM:
1. Initialize `result` array with -1.
2. Initialize `stack` (stores INDICES, not values, so we know where to write the answer).
3. Loop `i` from 0 to N.
4. While `stack` is not empty AND `arr[i] > arr[stack.top]`:
   - `index = stack.pop()`
   - `result[index] = arr[i]` (Because arr[i] is the one who "popped" the stack element).
5. Push `i` to stack.
*/

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	// Initialize with -1 (default if no greater element found)
	for i := range result {
		result[i] = -1
	}

	// This slice will act as our stack, storing INDICES
	stack := []int{}

	// DRY RUN:
	// Input: [2, 1, 5]
	//
	// i=0 (Val 2): Stack empty. Push 0. Stack: [0] (Val 2).
	//
	// i=1 (Val 1): Is 1 > 2? No.
	//              Push 1. Stack: [0, 1] (Vals 2, 1). (Notice it's decreasing: 2, 1).
	//
	// i=2 (Val 5): Is 5 > 1 (Stack top)? YES!
	//              - Pop 1. Result[1] = 5. (Next greater for 1 is 5).
	//              - Stack: [0] (Val 2).
	//              Is 5 > 2 (Stack top)? YES!
	//              - Pop 0. Result[0] = 5. (Next greater for 2 is 5).
	//              - Stack: [].
	//              Push 2. Stack: [2] (Val 5).
	//
	// End Loop. Result: [5, 5, -1].

	for i := 0; i < n; i++ {
		currentVal := nums[i]

		// While stack has elements AND current element is taller than the element referenced by stack top
		for len(stack) > 0 && currentVal > nums[stack[len(stack)-1]] {
			// Pop index
			lastIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// Record the answer
			result[lastIndex] = currentVal
		}

		// Push current index to wait for its own greater neighbor
		stack = append(stack, i)
	}

	return result
}

func main() {
	input := []int{2, 1, 5, 6, 2, 3}
	fmt.Printf("Input: %v\n", input)

	output := nextGreaterElements(input)
	fmt.Printf("Next Greater Elements: %v\n", output)
	// Expected for [2, 1, 5, 6, 2, 3]:
	// 2 -> 5
	// 1 -> 5
	// 5 -> 6
	// 6 -> -1
	// 2 -> 3
	// 3 -> -1
	// Res: [5, 5, 6, -1, 3, -1]
}
