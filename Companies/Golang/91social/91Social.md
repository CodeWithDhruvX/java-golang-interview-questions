# 91Social Golang Interview Questions

## 1. Jump Game (LeetCode 55)

**Problem:**
You are given an integer array `nums`. You are initially positioned at the array's **first index**, and each element in the array represents your maximum jump length at that position.

Return `true` if you can reach the last index, or `false` otherwise.

**Example 1:**
```
Input: nums = [2,3,1,1,4]
Output: true 
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
```

**Example 2:**
```
Input: nums = [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum jump length is 0, which makes it impossible to reach the last index.
```

**Constraints:**
- `1 <= nums.length <= 10^4`
- `0 <= nums[i] <= 10^5`

### Solution (Greedy Approach)

**Logic:**
We iterate backwards from the last index. The goal is to see if we can reach the current `lastPos` from the current index `i`.
- If `i + nums[i] >= lastPos`, it means we can jump from `i` to `lastPos`. So, `i` becomes the new target (`lastPos = i`).
- If `lastPos` becomes `0` by the end of the loop, it means we can reach the end from the start.

```go
func canJump(nums []int) bool {
	lastPos := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		// If we can reach the lastPos (or beyond) from index i
		if i+nums[i] >= lastPos {
			lastPos = i // Update lastPos to current index
		}
	}
	// If the lastPos reached 0, it means we can jump from start to end
	return lastPos == 0
}
```

**Complexity Analysis:**
- **Time Complexity:** O(N) — We iterate through the array once.
- **Space Complexity:** O(1) — We only use a single variable `lastPos`.