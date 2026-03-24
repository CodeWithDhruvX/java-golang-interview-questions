package main

import (
    "fmt"
    "sort"
    "strings"
)

/**
 * Demonstrates Go slice operations equivalent to Java List Interface.
 * 
 * Implementations: slice (dynamic array), linked list
 * 
 * Methods covered:
 * - insert at index, get at index, set at index, remove at index
 * - indexOf, lastIndexOf
 * - sub-slice operations
 * - sort, replaceAll
 */

func main() {
    fmt.Println("=== 1. Slice Implementation ===")
    arrayList := []string{"Java", "Python", "C++", "Java"} // Duplicate
    
    // Basic List Operations
    fmt.Printf("Original List: %v\n", arrayList)
    
    // insert at index equivalent to add(int index, E element)
    arrayList = insertAt(arrayList, 1, "JavaScript")
    fmt.Printf("After add at index 1: %v\n", arrayList)
    
    // get at index
    fmt.Printf("Element at index 2: %s\n", getAt(arrayList, 2))
    
    // set at index
    arrayList = setAt(arrayList, 0, "Golang")
    fmt.Printf("After set at index 0: %v\n", arrayList)
    
    // indexOf / lastIndexOf
    fmt.Printf("IndexOf 'Java': %d\n", indexOf(arrayList, "Java"))
    fmt.Printf("LastIndexOf 'Java': %d\n", lastIndexOf(arrayList, "Java"))
    
    // remove at index
    arrayList = removeAt(arrayList, 3) // Removes C++
    fmt.Printf("After remove at index 3: %v\n", arrayList)

    // subList equivalent
    sub := subSlice(arrayList, 0, 2)
    fmt.Printf("SubSlice (0-2): %v\n", sub)
    
    // sort
    sort.Strings(arrayList)
    fmt.Printf("After sort: %v\n", arrayList)
    
    // replaceAll equivalent
    arrayList = replaceAll(arrayList, func(s string) string {
        return strings.ToUpper(s)
    })
    fmt.Printf("After replaceAll (uppercase): %v\n", arrayList)
    
    fmt.Println("\n=== 2. LinkedList Implementation ===")
    linkedList := NewLinkedList()
    linkedList.Add("Node1")
    linkedList.Add("Node2")
    linkedList.Add("Node3")
    linkedList.Insert(1, "InsertedNode")
    
    fmt.Printf("LinkedList: ")
    linkedList.Print()
    
    fmt.Printf("Get at index 1: %s\n", linkedList.Get(1))
    linkedList.Set(0, "FirstNode")
    fmt.Printf("After set at index 0: ")
    linkedList.Print()
    
    linkedList.Remove(2)
    fmt.Printf("After remove at index 2: ")
    linkedList.Print()
}

// Helper functions for slice operations
func insertAt(slice []string, index int, value string) []string {
    if index < 0 || index > len(slice) {
        return slice
    }
    slice = append(slice, "")
    copy(slice[index+1:], slice[index:])
    slice[index] = value
    return slice
}

func getAt(slice []string, index int) string {
    if index < 0 || index >= len(slice) {
        return ""
    }
    return slice[index]
}

func setAt(slice []string, index int, value string) []string {
    if index < 0 || index >= len(slice) {
        return slice
    }
    slice[index] = value
    return slice
}

func removeAt(slice []string, index int) []string {
    if index < 0 || index >= len(slice) {
        return slice
    }
    return append(slice[:index], slice[index+1:]...)
}

func indexOf(slice []string, value string) int {
    for i, item := range slice {
        if item == value {
            return i
        }
    }
    return -1
}

func lastIndexOf(slice []string, value string) int {
    for i := len(slice) - 1; i >= 0; i-- {
        if slice[i] == value {
            return i
        }
    }
    return -1
}

func subSlice(slice []string, start, end int) []string {
    if start < 0 || end > len(slice) || start > end {
        return nil
    }
    result := make([]string, end-start)
    copy(result, slice[start:end])
    return result
}

func replaceAll(slice []string, replacer func(string) string) []string {
    result := make([]string, len(slice))
    for i, item := range slice {
        result[i] = replacer(item)
    }
    return result
}

// LinkedList implementation
type Node struct {
    Value string
    Next  *Node
}

type LinkedList struct {
    Head *Node
    Size int
}

func NewLinkedList() *LinkedList {
    return &LinkedList{}
}

func (ll *LinkedList) Add(value string) {
    newNode := &Node{Value: value}
    if ll.Head == nil {
        ll.Head = newNode
    } else {
        current := ll.Head
        for current.Next != nil {
            current = current.Next
        }
        current.Next = newNode
    }
    ll.Size++
}

func (ll *LinkedList) Insert(index int, value string) {
    if index < 0 || index > ll.Size {
        return
    }
    
    newNode := &Node{Value: value}
    
    if index == 0 {
        newNode.Next = ll.Head
        ll.Head = newNode
    } else {
        current := ll.Head
        for i := 0; i < index-1; i++ {
            current = current.Next
        }
        newNode.Next = current.Next
        current.Next = newNode
    }
    ll.Size++
}

func (ll *LinkedList) Get(index int) string {
    if index < 0 || index >= ll.Size {
        return ""
    }
    
    current := ll.Head
    for i := 0; i < index; i++ {
        current = current.Next
    }
    return current.Value
}

func (ll *LinkedList) Set(index int, value string) {
    if index < 0 || index >= ll.Size {
        return
    }
    
    current := ll.Head
    for i := 0; i < index; i++ {
        current = current.Next
    }
    current.Value = value
}

func (ll *LinkedList) Remove(index int) {
    if index < 0 || index >= ll.Size {
        return
    }
    
    if index == 0 {
        ll.Head = ll.Head.Next
    } else {
        current := ll.Head
        for i := 0; i < index-1; i++ {
            current = current.Next
        }
        current.Next = current.Next.Next
    }
    ll.Size--
}

func (ll *LinkedList) Print() {
    current := ll.Head
    for current != nil {
        fmt.Printf("%s -> ", current.Value)
        current = current.Next
    }
    fmt.Println("nil")
}
