package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 876. Middle of the Linked List
// Time: O(N), Space: O(1) - Tortoise and Hare Algorithm
func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	
	return slow
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
		{1, 2, 3, 4, 5},     // Odd length
		{1, 2, 3, 4, 5, 6},  // Even length
		{1},                  // Single node
		{},                   // Empty list
		{1, 2},               // Two nodes
		{1, 2, 3},            // Three nodes
		{1, 2, 3, 4},         // Four nodes
		{10, 20, 30, 40, 50, 60, 70}, // Seven nodes
	}
	
	for i, nums := range testCases {
		head := createLinkedList(nums)
		middle := middleNode(head)
		result := linkedListToSlice(middle)
		
		fmt.Printf("Test Case %d: %v -> Middle node(s): %v\n", i+1, nums, result)
	}
}
