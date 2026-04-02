package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 21. Merge Two Sorted Lists
// Time: O(N+M), Space: O(1)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}
	
	// Attach the remaining elements
	if list1 != nil {
		current.Next = list1
	} else {
		current.Next = list2
	}
	
	return dummy.Next
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
	testCases := []struct {
		list1 []int
		list2 []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}},
		{[]int{}, []int{}},
		{[]int{}, []int{0}},
		{[]int{1, 2, 3}, []int{4, 5, 6}},
		{[]int{1, 3, 5}, []int{2, 4, 6}},
		{[]int{1, 1, 1}, []int{1, 1, 1}},
		{[]int{1, 2, 3}, []int{}},
		{[]int{-3, -1, 1}, []int{-2, 0, 2}},
	}
	
	for i, tc := range testCases {
		list1 := createLinkedList(tc.list1)
		list2 := createLinkedList(tc.list2)
		merged := mergeTwoLists(list1, list2)
		result := linkedListToSlice(merged)
		
		fmt.Printf("Test Case %d: list1=%v, list2=%v -> Merged: %v\n", 
			i+1, tc.list1, tc.list2, result)
	}
}
