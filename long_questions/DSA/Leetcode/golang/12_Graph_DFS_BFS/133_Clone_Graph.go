package main

import "fmt"

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

// 133. Clone Graph
// Time: O(V+E), Space: O(V)
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	
	visited := make(map[*Node]*Node)
	return cloneNode(node, visited)
}

func cloneNode(node *Node, visited map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	
	// If node is already cloned, return the clone
	if clone, exists := visited[node]; exists {
		return clone
	}
	
	// Create a clone for the current node
	clone := &Node{Val: node.Val}
	visited[node] = clone
	
	// Clone all neighbors
	for _, neighbor := range node.Neighbors {
		clone.Neighbors = append(clone.Neighbors, cloneNode(neighbor, visited))
	}
	
	return clone
}

// Helper function to create a graph from adjacency list
func createGraph(adjList [][]int) *Node {
	if len(adjList) == 0 {
		return nil
	}
	
	nodes := make([]*Node, len(adjList))
	for i := range nodes {
		nodes[i] = &Node{Val: i + 1}
	}
	
	for i, neighbors := range adjList {
		for _, neighbor := range neighbors {
			nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[neighbor-1])
		}
	}
	
	return nodes[0]
}

// Helper function to convert graph to adjacency list for verification
func graphToAdjList(node *Node) [][]int {
	if node == nil {
		return [][]int{}
	}
	
	visited := make(map[*Node]bool)
	adjList := make(map[int][]int)
	
	var dfs func(*Node)
	dfs = func(n *Node) {
		if n == nil || visited[n] {
			return
		}
		
		visited[n] = true
		for _, neighbor := range n.Neighbors {
			adjList[n.Val] = append(adjList[n.Val], neighbor.Val)
			dfs(neighbor)
		}
	}
	
	dfs(node)
	
	// Convert to slice format
	result := make([][]int, len(visited))
	for val, neighbors := range adjList {
		result[val-1] = neighbors
	}
	
	return result
}

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Graph Traversal with Cloning
- **DFS Traversal**: Depth-first exploration of original graph
- **Node Mapping**: Maintain map from original to cloned nodes
- **Recursive Cloning**: Clone node and recursively clone neighbors
- **Cycle Handling**: Use visited map to handle cycles and shared nodes

## 2. PROBLEM CHARACTERISTICS
- **Undirected Graph**: Edges are bidirectional
- **Node Structure**: Each node has value and list of neighbors
- **Deep Copy**: Create completely independent copy of graph
- **Cycle Detection**: Graph may contain cycles

## 3. SIMILAR PROBLEMS
- Copy List with Random Pointer (LeetCode 138)
- Clone Binary Tree (LeetCode - various variants)
- Graph Serialization/Deserialization
- Network Clone problems

## 4. KEY OBSERVATIONS
- **Visited tracking**: Essential to handle cycles and shared references
- **Recursive structure**: Natural fit for DFS approach
- **Deep copy requirement**: Must create new nodes, not copy references
- **Bidirectional edges**: Need to maintain neighbor relationships

## 5. VARIATIONS & EXTENSIONS
- **BFS Approach**: Use queue for iterative cloning
- **Directed Graph**: Handle directed edges instead of undirected
- **Weighted Graph**: Clone with edge weights
- **Graph Validation**: Verify cloned graph structure

## 6. INTERVIEW INSIGHTS
- Always clarify: "Can the graph be empty? Can there be self-loops?"
- Edge cases: empty graph, single node, disconnected components
- Time complexity: O(V+E) - visit each node and edge once
- Space complexity: O(V) - for visited map and recursion stack

## 7. COMMON MISTAKES
- Not handling empty graph case
- Forgetting to mark nodes as visited when creating clones
- Creating shallow copies instead of deep copies
- Not handling cycles properly leading to infinite recursion
- Missing bidirectional edge relationships

## 8. OPTIMIZATION STRATEGIES
- **Current solution is optimal**: O(V+E) time, O(V) space
- **Iterative BFS**: Can avoid recursion stack limitations
- **Early termination**: Not applicable (need to clone entire graph)
- **Memory efficiency**: Use minimal additional data structures

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like photocopying a network of connected documents:**
- You have a network where each document references other documents
- You want to create an exact copy of this entire network
- Each copied document must reference other copied documents, not originals
- As you copy each document, you need to note which original maps to which copy
- If you encounter a document you've already copied, use the existing copy

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Reference to a node in an undirected connected graph
2. **Goal**: Create deep copy of the entire graph
3. **Output**: Reference to corresponding node in cloned graph
4. **Constraint**: Must handle cycles and shared references correctly

#### Phase 2: Key Insight Recognition
- **"Graph traversal"** → Need to visit all reachable nodes
- **"Deep copy requirement"** → Must create new node objects
- **"Cycle handling"** → Need visited tracking to avoid infinite loops
- **"Reference mapping"** → Need map from original to cloned nodes

#### Phase 3: Strategy Development
```
Human thought process:
"I need to create a complete copy of this graph.
I'll use DFS to traverse the graph.
For each node, I'll create a new clone and map the original to clone.
Then I'll recursively clone all neighbors.
If I encounter a node I've already cloned, I'll return the existing clone.
This ensures proper handling of cycles and shared references."
```

#### Phase 4: Edge Case Handling
- **Empty graph**: Return nil
- **Single node**: Create single cloned node with no neighbors
- **Self-loops**: Handle node that references itself
- **Disconnected components**: Only clone reachable component from given node

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Graph: 1 -- 2 -- 3
       |    |
       4 -- 5

Human thinking:
"I'll start cloning from node 1:

1. Clone node 1: create clone1, map[1]=clone1
2. Clone neighbors of 1: nodes 2 and 4
   - Clone node 2: create clone2, map[2]=clone2
   - Clone neighbors of 2: nodes 1, 3, 5
     * Node 1 already cloned -> use clone1
     * Clone node 3: create clone3, map[3]=clone3
     * Clone node 5: create clone5, map[5]=clone5
   - Clone neighbors of 3: node 2 already cloned
   - Clone neighbors of 5: nodes 2, 4
     * Node 2 already cloned -> use clone2
     * Clone node 4: create clone4, map[4]=clone4
   - Clone neighbors of 4: nodes 1, 5 already cloned
   
Final mapping:
1 -> clone1 (neighbors: clone2, clone4)
2 -> clone2 (neighbors: clone1, clone3, clone5)  
3 -> clone3 (neighbors: clone2)
4 -> clone4 (neighbors: clone1, clone5)
5 -> clone5 (neighbors: clone2, clone4)

Return clone1 as entry point to cloned graph"
```

#### Phase 6: Intuition Validation
- **Why DFS works**: Natural recursive structure for graph traversal
- **Why visited map needed**: Prevents infinite recursion and handles cycles
- **Why deep copy**: Must create independent node objects
- **Why O(V+E)**: Visit each node and edge exactly once

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just copy the node reference?"** → Need deep copy, not shallow copy
2. **"Should I use BFS instead?"** → Can, but DFS is more natural for recursive cloning
3. **"What about very deep graphs?"** → Consider iterative BFS to avoid stack overflow
4. **"Can I optimize further?"** → Already optimal for this problem

### Real-World Analogy
**Like creating a backup copy of a social network:**
- You have a social network where each person follows others
- You want to create an exact backup copy of this network
- Each copied person must follow other copied people, not originals
- As you copy each person, you note which original maps to which backup
- If you encounter a person already copied, use the existing backup copy
- This ensures the backup network is completely independent

### Human-Readable Pseudocode
```
function cloneGraph(node):
    if node is nil:
        return nil
    
    visited = map() // original node -> cloned node
    return cloneNode(node, visited)

function cloneNode(node, visited):
    if node is nil:
        return nil
    
    // If node is already cloned, return the clone
    if node in visited:
        return visited[node]
    
    // Create a clone for the current node
    clone = new Node(node.val)
    visited[node] = clone
    
    // Clone all neighbors
    for neighbor in node.neighbors:
        cloneNeighbor = cloneNode(neighbor, visited)
        clone.neighbors.append(cloneNeighbor)
    
    return clone
```

### Execution Visualization

### Example Graph:
```
1 -- 2
|    |
4 -- 3
```

### Cloning Process:
```
Start with node 1:
1. Create clone1, map[1]=clone1
2. Process neighbors of 1: nodes 2, 4

Process node 2:
1. Create clone2, map[2]=clone2  
2. Process neighbors of 2: nodes 1, 3
   - Node 1 already cloned -> use clone1
   - Create clone3, map[3]=clone3

Process node 3:
1. Node 3 already cloned -> return clone3
2. Process neighbors of 3: node 2 already cloned

Process node 4:
1. Create clone4, map[4]=clone4
2. Process neighbors of 4: nodes 1, 3 already cloned

Final cloned graph:
clone1 -- clone2
|        |
clone4 -- clone3
```

### Key Visualization Points:
- **Visited mapping**: Original nodes map to cloned nodes
- **Recursive cloning**: Each node clones its neighbors recursively
- **Cycle handling**: Already cloned nodes are reused
- **Deep copy**: All nodes are new objects

### Memory Layout Visualization:
```
Visited Map Evolution:
{} → {1:clone1} → {1:clone1, 2:clone2} → {1:clone1, 2:clone2, 3:clone3} → {1:clone1, 2:clone2, 3:clone3, 4:clone4}

Clone Graph Structure:
clone1.neighbors = [clone2, clone4]
clone2.neighbors = [clone1, clone3]  
clone3.neighbors = [clone2]
clone4.neighbors = [clone1, clone3]
```

### Time Complexity Breakdown:
- **Each node visited**: O(1) work per node for cloning
- **Each edge processed**: O(1) work per edge for neighbor cloning
- **Total nodes**: V, Total edges: E
- **Total time**: O(V+E)
- **Space**: O(V) for visited map and recursion stack

### Alternative Approaches:

#### 1. Iterative BFS Approach (O(V+E) time, O(V) space)
```go
func cloneGraphBFS(node *Node) *Node {
    if node == nil {
        return nil
    }
    
    visited := make(map[*Node]*Node)
    queue := []*Node{node}
    
    // Create clone for the starting node
    visited[node] = &Node{Val: node.Val}
    
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        
        for _, neighbor := range current.Neighbors {
            if _, exists := visited[neighbor]; !exists {
                visited[neighbor] = &Node{Val: neighbor.Val}
                queue = append(queue, neighbor)
            }
            
            visited[current].Neighbors = append(visited[current].Neighbors, visited[neighbor])
        }
    }
    
    return visited[node]
}
```
- **Pros**: No recursion stack, iterative approach
- **Cons**: Slightly more complex implementation

#### 2. Two-Phase Approach (O(V+E) time, O(V) space)
```go
func cloneGraphTwoPhase(node *Node) *Node {
    if node == nil {
        return nil
    }
    
    // Phase 1: Create all nodes
    visited := make(map[*Node]*Node)
    queue := []*Node{node}
    visited[node] = &Node{Val: node.Val}
    
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        
        for _, neighbor := range current.Neighbors {
            if _, exists := visited[neighbor]; !exists {
                visited[neighbor] = &Node{Val: neighbor.Val}
                queue = append(queue, neighbor)
            }
        }
    }
    
    // Phase 2: Connect all edges
    queue = []*Node{node}
    processed := make(map[*Node]bool)
    
    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]
        
        if processed[current] {
            continue
        }
        
        for _, neighbor := range current.Neighbors {
            visited[current].Neighbors = append(visited[current].Neighbors, visited[neighbor])
            if !processed[neighbor] {
                queue = append(queue, neighbor)
            }
        }
        
        processed[current] = true
    }
    
    return visited[node]
}
```
- **Pros**: Clear separation of node creation and edge connection
- **Cons**: Two passes over the graph

#### 3. Union Find Approach (O(V+E) time, O(V) space)
```go
func cloneGraphUnionFind(node *Node) *Node {
    if node == nil {
        return nil
    }
    
    // This is more complex and not typically used for this problem
    // but demonstrates another approach to graph cloning
    // Implementation would involve using Union-Find to track
    // equivalent nodes and build the cloned graph
    
    return cloneGraph(node) // Fall back to DFS approach
}
```
- **Pros**: Theoretical interest
- **Cons**: Overly complex for this problem

### Extensions for Interviews:
- **Directed Graph**: Handle one-way edges instead of bidirectional
- **Weighted Graph**: Clone with edge weights preserved
- **Graph Validation**: Verify cloned graph matches original structure
- **Multiple Components**: Handle graphs with multiple disconnected components
- **Serialization**: Convert graph to string representation and back
*/
func main() {
	// Test cases
	testCases := [][][]int{
		{{2, 4}, {1, 3}, {2, 4}, {1, 3}}, // 4-node graph
		{{}}, // Single node
		{}, // Empty graph
		{{2}, {1}}, // 2-node graph
		{{2, 3}, {1, 3}, {1, 2}}, // 3-node complete graph
		{{2}, {3}, {4}, {5}, {1}}, // 5-node cycle
	}
	
	for i, adjList := range testCases {
		node := createGraph(adjList)
		cloned := cloneGraph(node)
		clonedAdjList := graphToAdjList(cloned)
		fmt.Printf("Test Case %d: %v -> Cloned: %v\n", i+1, adjList, clonedAdjList)
	}
}
