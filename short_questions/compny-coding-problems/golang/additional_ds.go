package main

import (
	"fmt"
	"sort"
	"strings"
)

// --- LINKED LIST (BASIC ONLY) ---

type ListNode struct {
	Val  int
	Next *ListNode
}

// 51. Create LL
func createLL(val int) *ListNode {
	return &ListNode{Val: val}
}

// 52. Traverse
func traverseLL(head *ListNode) {
	curr := head
	for curr != nil {
		fmt.Printf("%d -> ", curr.Val)
		curr = curr.Next
	}
	fmt.Println("NULL")
}

// 53. Insert Begin
func insertBegin(head *ListNode, val int) *ListNode {
	newNode := &ListNode{Val: val, Next: head}
	return newNode
}

// 54. Insert End
func insertEnd(head *ListNode, val int) *ListNode {
	newNode := &ListNode{Val: val}
	if head == nil {
		return newNode
	}
	curr := head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = newNode
	return head
}

// 55. Delete Node (by value)
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	curr := head
	for curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
			return head
		}
		curr = curr.Next
	}
	return head
}

// 56. Reverse LL
func reverseLL(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}
	return prev
}

// 57. Middle Element
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 58. Detect Loop
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// 59. Count Nodes
func countNodes(head *ListNode) int {
	count := 0
	curr := head
	for curr != nil {
		count++
		curr = curr.Next
	}
	return count
}

// 60. Merge 2 Lists (Sorted)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}
	if l1 != nil {
		curr.Next = l1
	} else {
		curr.Next = l2
	}
	return dummy.Next
}

// --- STACK & QUEUE (BASIC) ---

// 61. Stack (Array based)
type Stack struct {
	items []int
}

func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

// 62. Queue (Array based)
type Queue struct {
	items []int
}

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}
func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// 63. Reverse String (Stack)
func reverseStringStack(str string) string {
	stack := []rune{}
	for _, c := range str {
		stack = append(stack, c)
	}
	var res strings.Builder
	for i := len(stack) - 1; i >= 0; i-- {
		res.WriteRune(stack[i])
	}
	return res.String()
}

// 64. Balanced Parentheses (Covered in Additional Strings Q7)

// 65. Stack using Queue (Not implemented fully here, logic described)
// 66. Queue using Stack (Not implemented fully here, logic described)

// 67. Next Greater Element
func nextGreaterElement(nums []int) []int {
	res := make([]int, len(nums))
	for i := range res {
		res[i] = -1
	}
	stack := []int{} // stores indices
	for i, num := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < num {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[idx] = num
		}
		stack = append(stack, i)
	}
	return res
}

// 68. Evaluate Postfix
func evalPostfix(exp string) int {
	stack := []int{}
	for _, char := range exp {
		if char >= '0' && char <= '9' {
			stack = append(stack, int(char-'0'))
		} else {
			if len(stack) < 2 {
				return 0
			}
			val1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			val2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			switch char {
			case '+':
				stack = append(stack, val2+val1)
			case '-':
				stack = append(stack, val2-val1)
			case '*':
				stack = append(stack, val2*val1)
			case '/':
				stack = append(stack, val2/val1)
			}
		}
	}
	return stack[0]
}

// 69. Reverse Stack (Recursion) - Simulated
// 70. Circular Queue - Logic Only

// --- HASHING / MAP LOGIC ---

// 71. Freq elements
func freqElements(arr []int) {
	m := make(map[int]int)
	for _, v := range arr {
		m[v]++
	}
	fmt.Println(m)
}

// 72. First Repeating
func firstRepeating(arr []int) int {
	seen := make(map[int]bool)
	for _, v := range arr {
		if seen[v] {
			return v
		}
		seen[v] = true
	}
	return -1
}

// 73. First Non-Repeating (Covered in String Q7, Array logic similar)
func firstNonRepeatingArr(arr []int) int {
	counts := make(map[int]int)
	for _, v := range arr {
		counts[v]++
	}
	for _, v := range arr {
		if counts[v] == 1 {
			return v
		}
	}
	return -1
}

// 74. Two Sum
func twoSum(arr []int, target int) []int {
	m := make(map[int]int)
	for i, v := range arr {
		needed := target - v
		if idx, ok := m[needed]; ok {
			return []int{idx, i}
		}
		m[v] = i
	}
	return nil
}

// 75. Group Anagrams
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string)
	for _, s := range strs {
		split := strings.Split(s, "")
		sort.Strings(split)
		key := strings.Join(split, "")
		m[key] = append(m[key], s)
	}
	var res [][]string
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

// 76. Count Distinct Chars
func countDistinctChars(s string) int {
	m := make(map[rune]bool)
	for _, c := range s {
		m[c] = true
	}
	return len(m)
}

// 77. Majority Element (Covered in Arrays Q20, Map version)
func majorityElementMap(arr []int) int {
	m := make(map[int]int)
	n := len(arr)
	for _, v := range arr {
		m[v]++
		if m[v] > n/2 {
			return v
		}
	}
	return -1
}

// 78. Check Subset
func isSubset(arr1, arr2 []int) bool {
	m := make(map[int]bool)
	for _, v := range arr1 {
		m[v] = true
	}
	for _, v := range arr2 {
		if !m[v] {
			return false
		}
	}
	return true
}

// 79. Common Elements (Intersection - Covered in Arrays Q24)

// 80. Longest Substring No Repeats
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	maxLen, start := 0, 0
	for i := 0; i < len(s); i++ {
		if idx, ok := m[s[i]]; ok && idx >= start {
			start = idx + 1
		}
		m[s[i]] = i
		if i-start+1 > maxLen {
			maxLen = i - start + 1
		}
	}
	return maxLen
}

func main() {
	fmt.Println("--- LINKED LIST ---")
	head := createLL(1)
	head = insertEnd(head, 2)
	head = insertEnd(head, 3)
	traverseLL(head)
	fmt.Println("Middle:", middleNode(head).Val)
	head = reverseLL(head)
	traverseLL(head)

	fmt.Println("\n--- STACK/QUEUE (Basic) ---")
	st := Stack{}
	st.Push(10)
	st.Push(20)
	fmt.Println("Pop:", st.Pop())

	fmt.Println("Reverse String Stack: hello ->", reverseStringStack("hello"))
	fmt.Println("Next Greater: [4, 5, 2, 25] ->", nextGreaterElement([]int{4, 5, 2, 25}))
	fmt.Println("Postfix: 231*+9- ->", evalPostfix("231*+9-"))

	fmt.Println("\n--- HASHING ---")
	freqElements([]int{1, 2, 2, 3})
	fmt.Println("First Repeating: [1, 2, 3, 2] ->", firstRepeating([]int{1, 2, 3, 2}))
	fmt.Println("Two Sum: [2, 7, 11, 15], 9 ->", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println("Group Anagrams: [eat, tea, tan, ate, nat, bat] ->", groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println("Longest Substr No Repeat: abcabcbb ->", lengthOfLongestSubstring("abcabcbb"))
}
