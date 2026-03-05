# Trees and Graphs (Product-Based Companies)

Graph traversals (BFS/DFS), finding shortest paths, Connected Components, and complex Binary Tree permutations are staple questions in Amazon, Google, and Meta interviews.

## Question 1: Serialize and Deserialize Binary Tree
**Problem Statement:** Design an algorithm to serialize and deserialize a binary tree. Serialization is the process of converting a data structure or object into a sequence of bits/strings. Deserialization reconstructs the object.

### Answer:
A Pre-order traversal works well here. We can serialize the tree into a string delimited by commas, adding an `X` or `null` for empty nodes. When deserializing, we use a Queue or Array iterator to reconstruct the tree recursively in pre-order fashion.

**Code Implementation (Java):**
```java
import java.util.Arrays;
import java.util.LinkedList;
import java.util.Queue;

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;
    TreeNode(int x) { val = x; }
}

public class Codec {
    // Encodes a tree to a single string.
    public String serialize(TreeNode root) {
        StringBuilder sb = new StringBuilder();
        buildString(root, sb);
        return sb.toString();
    }
    
    private void buildString(TreeNode node, StringBuilder sb) {
        if (node == null) {
            sb.append("X").append(",");
        } else {
            sb.append(node.val).append(",");
            buildString(node.left, sb);
            buildString(node.right, sb);
        }
    }

    // Decodes your encoded data to tree.
    public TreeNode deserialize(String data) {
        Queue<String> nodes = new LinkedList<>(Arrays.asList(data.split(",")));
        return buildTree(nodes);
    }
    
    private TreeNode buildTree(Queue<String> nodes) {
        String val = nodes.poll();
        if (val.equals("X")) return null;
        
        TreeNode node = new TreeNode(Integer.parseInt(val));
        node.left = buildTree(nodes);
        node.right = buildTree(nodes);
        return node;
    }
}
```
**Time Complexity:** O(N) for both serialize and deserialize.
**Space Complexity:** O(N) for the resulting string and queue.

---

## Question 2: Lowest Common Ancestor of a Binary Tree
**Problem Statement:** Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.

### Answer:
We traverse the tree using DFS. If the current node is `null`, `p`, or `q`, we return the current node. Then we recursively search the left and right subtrees. If both left and right return non-null, the current node is the LCA. If only one side returns non-null, that side contains the LCA.

**Code Implementation (Java):**
```java
public class LowestCommonAncestor {
    public TreeNode lowestCommonAncestor(TreeNode root, TreeNode p, TreeNode q) {
        if (root == null || root == p || root == q) {
            return root;
        }
        
        TreeNode left = lowestCommonAncestor(root.left, p, q);
        TreeNode right = lowestCommonAncestor(root.right, p, q);
        
        if (left != null && right != null) {
            return root; // Both found, root is the LCA
        }
        
        return left != null ? left : right; // Either left or right is not null
    }
}
```
**Time Complexity:** O(N)
**Space Complexity:** O(H) where H is the height of the tree (for the recursion stack).

---

## Question 3: Course Schedule (Topological Sort)
**Problem Statement:** There are a total of `numCourses` courses you have to take, labeled from `0` to `numCourses - 1`. You are given an array `prerequisites` where `prerequisites[i] = [ai, bi]` indicates that you must take course `bi` first if you want to take course `ai`. Return `true` if you can finish all courses. Otherwise, return `false`.

### Answer:
This is a classic cycle detection problem in a directed graph. It can be solved using Kahn's Algorithm (BFS for Topological Sort) by keeping an in-degree array, or by using DFS to look for back edges. If there is a cycle, we cannot finish all courses.

**Code Implementation (Java):**
```java
import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

public class CourseSchedule {
    public boolean canFinish(int numCourses, int[][] prerequisites) {
        List<List<Integer>> graph = new ArrayList<>();
        int[] inDegree = new int[numCourses];
        
        for (int i = 0; i < numCourses; i++) {
            graph.add(new ArrayList<>());
        }
        
        for (int[] pre : prerequisites) {
            graph.get(pre[1]).add(pre[0]);
            inDegree[pre[0]]++;
        }
        
        Queue<Integer> queue = new LinkedList<>();
        for (int i = 0; i < numCourses; i++) {
            if (inDegree[i] == 0) {
                queue.add(i);
            }
        }
        
        int count = 0;
        while (!queue.isEmpty()) {
            int current = queue.poll();
            count++;
            for (int neighbor : graph.get(current)) {
                inDegree[neighbor]--;
                if (inDegree[neighbor] == 0) {
                    queue.add(neighbor);
                }
            }
        }
        
        return count == numCourses;
    }
}
```
**Time Complexity:** O(V + E) where V is `numCourses` and E is the length of `prerequisites`.
**Space Complexity:** O(V + E)

---

## Question 4: Word Ladder
**Problem Statement:** A transformation sequence from word `beginWord` to word `endWord` using a dictionary `wordList` is a sequence of words such that every adjacent pair of words differs by a single letter. Return the number of words in the shortest transformation sequence from `beginWord` to `endWord`, or 0 if no such sequence exists.

### Answer:
The problem asks for the *shortest* path, which signals we should use Breadth-First Search (BFS). We start from the `beginWord`, change one letter at a time to all `a-z`, check if the new word is in the `wordList` (converted to a HashSet for O(1) lookup), and add it to our Queue. The layer level implies the path length.

**Code Implementation (Java):**
```java
import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;
import java.util.Set;

public class WordLadder {
    public int ladderLength(String beginWord, String endWord, List<String> wordList) {
        Set<String> set = new HashSet<>(wordList);
        if (!set.contains(endWord)) return 0;
        
        Queue<String> queue = new LinkedList<>();
        queue.add(beginWord);
        int level = 1;
        
        while (!queue.isEmpty()) {
            int size = queue.size();
            for (int i = 0; i < size; i++) {
                String currentWord = queue.poll();
                char[] wordChars = currentWord.toCharArray();
                
                for (int j = 0; j < wordChars.length; j++) {
                    char originalChar = wordChars[j];
                    for (char c = 'a'; c <= 'z'; c++) {
                        if (wordChars[j] == c) continue;
                        wordChars[j] = c;
                        String newWord = String.valueOf(wordChars);
                        
                        if (newWord.equals(endWord)) return level + 1;
                        if (set.contains(newWord)) {
                            queue.add(newWord);
                            set.remove(newWord); // Visited
                        }
                    }
                    wordChars[j] = originalChar;
                }
            }
            level++;
        }
        return 0;
    }
}
```
**Time Complexity:** O(M^2 * N), where M is word length, and N is the total number of words in `wordList`.
**Space Complexity:** O(M * N)
