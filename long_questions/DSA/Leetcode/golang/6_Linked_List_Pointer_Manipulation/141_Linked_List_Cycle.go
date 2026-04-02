package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 141. Linked List Cycle
// Time: O(N), Space: O(1) - Floyd's Cycle Detection Algorithm
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	
	slow := head
	fast := head.Next
	
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	
	return false
}

// Helper function to create a linked list with optional cycle
func createLinkedListWithCycle(nums []int, cyclePos int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	
	head := &ListNode{Val: nums[0]}
	current := head
	var cycleNode *ListNode
	
	if cyclePos == 0 {
		cycleNode = head
	}
	
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
		
		if i == cyclePos {
			cycleNode = current
		}
	}
	
	// Create cycle if cyclePos is valid
	if cyclePos >= 0 && cyclePos < len(nums) {
		current.Next = cycleNode
	}
	
	return head
}

func main() {
	// Test cases
	testCases := []struct {
		nums     []int
		cyclePos int // -1 means no cycle
	}{
		{[]int{3, 2, 0, -4}, 1}, // Cycle at position 1 (value 2)
		{[]int{1, 2}, 0},        // Cycle at position 0 (value 1)
		{[]int{1}, -1},          // No cycle
		{[]int{}, -1},           // Empty list
		{[]int{1, 2, 3, 4}, 2},  // Cycle at position 2 (value 3)
		{[]int{1, 2, 3, 4, 5}, -1}, // No cycle
	}
	
	for i, tc := range testCases {
		head := createLinkedListWithCycle(tc.nums, tc.cyclePos)
		hasCycleResult := hasCycle(head)
		
		cycleInfo := "No cycle"
		if tc.cyclePos >= 0 {
			cycleInfo = fmt.Sprintf("Cycle at position %d", tc.cyclePos)
		}
		
		fmt.Printf("Test Case %d: %v (%s) -> Has cycle: %t\n", 
			i+1, tc.nums, cycleInfo, hasCycleResult)
	}
}
