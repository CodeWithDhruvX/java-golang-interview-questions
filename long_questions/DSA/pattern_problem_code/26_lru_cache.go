package main

import (
	"container/list"
	"fmt"
)

// Pattern: Data Structure Design
// Difficulty: Medium/Hard
// Key Concept: Combining Hash Map (O(1) access) and Doubly Linked List (O(1) reordering) to enforce LRU policy.

/*
INTUITION:
"LRU Cache" (Least Recently Used)
Imagine you have a desk with space for only 3 books.
1. You read Book A. Put it on desk.
2. Read Book B. Put it on desk.
3. Read Book C. Put it on desk. (Desk Full: C, B, A).
4. You want to read Book D. No space!
   - Principle: "If I haven't used it in a while, I probably won't use it soon."
   - Throw away the book you read *longest* ago (Book A).
   - Put D on desk.

To do this efficienty (O(1)):
- **Hash Map**: Lets us find if a book is on the desk instantly. Key=BookID, Val=PointerToNode.
- **Doubly Linked List**: Keeps order.
  - Head = Most Recently Used.
  - Tail = Least Recently Used.
  - When we read a book, we move it to the Head.
  - When we evict, we remove the Tail.

PROBLEM:
Design LRUCache with `get(key)` and `put(key, value)`. Both O(1).

ALGORITHM:
1. `get(key)`:
   - If key in Map:
     - Move its Node to Front of List (Make it "Recent").
     - Return Value.
   - Else: Return -1.
2. `put(key, value)`:
   - If key in Map:
     - Update value. Move Node to Front.
   - If key NOT in Map:
     - If Capacity Full: Remove Back Node (Tail). Delete from Map.
     - Add New Node to Front. Add to Map.
*/

// Use Go's built-in container/list for Doubly Linked List
type LRUCache struct {
	capacity int
	cache    map[int]*list.Element // Key -> List Element
	lruList  *list.List            // Doubly Linked List
}

type Pair struct {
	key   int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		lruList:  list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	if elem, found := this.cache[key]; found {
		// Key exists. Move it to the front (Most Recently Used).
		this.lruList.MoveToFront(elem)
		return elem.Value.(Pair).value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if elem, found := this.cache[key]; found {
		// Update existing
		this.lruList.MoveToFront(elem)
		// Update the value inside the element
		elem.Value = Pair{key, value}
	} else {
		// Insert new
		if this.lruList.Len() == this.capacity {
			// Evict LRU (Back of list)
			lastElem := this.lruList.Back()
			if lastElem != nil {
				// Remove from List
				this.lruList.Remove(lastElem)
				// Remove from Map
				delete(this.cache, lastElem.Value.(Pair).key)
			}
		}
		// Add to Front
		newElem := this.lruList.PushFront(Pair{key, value})
		this.cache[key] = newElem
	}
}

func main() {
	obj := Constructor(2) // Capacity 2

	obj.Put(1, 1) // Cache: {1:1}
	obj.Put(2, 2) // Cache: {2:2, 1:1} (2 is recent)

	fmt.Printf("Get 1: %d\n", obj.Get(1)) // Returns 1. Cache: {1:1, 2:2} (1 is recent)

	obj.Put(3, 3) // Capacity full! Evicts LRU (2). Cache: {3:3, 1:1}

	fmt.Printf("Get 2: %d\n", obj.Get(2)) // Returns -1 (Evicted)

	obj.Put(4, 4) // Capacity full! Evicts LRU (1). Cache: {4:4, 3:3}

	fmt.Printf("Get 1: %d\n", obj.Get(1)) // Returns -1 (Evicted)
	fmt.Printf("Get 3: %d\n", obj.Get(3)) // Returns 3
	fmt.Printf("Get 4: %d\n", obj.Get(4)) // Returns 4
}
