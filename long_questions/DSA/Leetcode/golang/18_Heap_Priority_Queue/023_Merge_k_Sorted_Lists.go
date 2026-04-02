package main

import (
	"container/heap"
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 23. Merge k Sorted Lists
// Time: O(N log K), Space: O(K) where N is total nodes, K is number of lists
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	
	minHeap := &NodeHeap{}
	heap.Init(minHeap)
	
	// Push the head of each list into the heap
	for _, list := range lists {
		if list != nil {
			heap.Push(minHeap, list)
		}
	}
	
	dummy := &ListNode{}
	current := dummy
	
	for minHeap.Len() > 0 {
		// Get the smallest node
		smallest := heap.Pop(minHeap).(*ListNode)
		current.Next = smallest
		current = current.Next
		
		// Push the next node from the same list
		if smallest.Next != nil {
			heap.Push(minHeap, smallest.Next)
		}
	}
	
	return dummy.Next
}

// NodeHeap implementation for ListNode
type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
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
	testCases := [][][]int{
		{{1, 4, 5}, {1, 3, 4}, {2, 6}},
		{},
		{{}},
		{{1, 2, 3}},
		{{1}, {1}, {1}},
		{{1, 2}, {3, 4}, {5, 6}},
		{{-5, -3, -1}, {-4, -2, 0}, {-6, -4, -2}},
		{{1, 3, 5}, {2, 4, 6}, {7, 8, 9}},
		{{1, 100}, {2, 99}, {3, 98}},
	}
	
	for i, tc := range testCases {
		// Convert to linked lists
		lists := make([]*ListNode, len(tc))
		for j, nums := range tc {
			lists[j] = createLinkedList(nums)
		}
		
		merged := mergeKLists(lists)
		result := linkedListToSlice(merged)
		
		fmt.Printf("Test Case %d: %v -> Merged: %v\n", i+1, tc, result)
	}
}
