package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 206. Reverse Linked List
// Time: O(N), Space: O(1)
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	
	return prev
}

// Helper function to create a linked list from slice
func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	
	head := &ListNode{Val: nums[0]}
	current := head
	
	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{Val: nums[i]}
		current = current.Next
	}
	
	return head
}

// Helper function to convert linked list to slice
func linkedListToSlice(head *ListNode) []int {
	var result []int
	current := head
	
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	
	return result
}

func main() {
	// Test cases
	testCases := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2},
		{},
		{1},
		{1, 2, 3, 4},
		{5, 4, 3, 2, 1},
	}
	
	for i, nums := range testCases {
		head := createLinkedList(nums)
		reversedHead := reverseList(head)
		result := linkedListToSlice(reversedHead)
		fmt.Printf("Test Case %d: %v -> Reversed: %v\n", i+1, nums, result)
	}
}
