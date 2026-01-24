package main

import "fmt"

// Pattern: Linked List Pointer Manipulation
// Difficulty: Easy
// Key Concept: Manipulating 'Next' pointers to change the structure of the list without creating new nodes.

/*
INTUITION:
A Linked List is just a chain.
Node A -> Node B -> Node C -> Node D -> nil
To reverse it, we want:
Node A <- Node B <- Node C <- Node D
        \
         nil

Think of it as 3 people standing in a line holding hands.
You (Curr) are holding B's hand.
You let go of B, and instead grab the hand of the person BEHIND you (Prev).
Then you step forward.

We need 3 pointers to do this safely:
1. `prev` (The node behind me, initially nil)
2. `curr` (Me, initially Head)
3. `next` (The guy ahead of me. I need to remember him before I let go of his hand!)

PROBLEM:
Reverse a Singly Linked List.

ALGORITHM:
Loop while `curr` is not nil:
1. Save `next = curr.Next` (Save the future!)
2. `curr.Next = prev` (Do the reversal!)
3. `prev = curr` (Advance prev to where I am)
4. `curr = next` (Advance me to where the future was)
Return `prev` (Because `curr` will be nil at the end).
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head

	// DRY RUN:
	// List: 1 -> 2 -> 3 -> nil
	//
	// Iteration 1:
	// Curr = 1.
	// Next = 2 (Saved).
	// 1.Next = nil (Reversed: nil <- 1).
	// Prev = 1.
	// Curr = 2.
	//
	// Iteration 2:
	// Curr = 2.
	// Next = 3 (Saved).
	// 2.Next = 1 (Reversed: nil <- 1 <- 2).
	// Prev = 2.
	// Curr = 3.
	//
	// Iteration 3:
	// Curr = 3.
	// Next = nil.
	// 3.Next = 2 (Reversed: nil <- 1 <- 2 <- 3).
	// Prev = 3.
	// Curr = nil.
	//
	// Loop ends. Return Prev (3).

	for curr != nil {
		nextTemp := curr.Next // Save next
		curr.Next = prev      // Reverse pointer
		prev = curr           // Move pointers forward
		curr = nextTemp
	}

	return prev
}

// Helper to print list
func printList(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%d -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println("nil")
}

func main() {
	// Construct 1->2->3->4->5
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	fmt.Println("Original:")
	printList(head)

	newHead := reverseList(head)

	fmt.Println("Reversed:")
	printList(newHead)
}
