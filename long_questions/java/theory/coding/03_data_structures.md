# Coding: Data Structures - Interview Answers

> 🎯 **Focus:** These questions check if you understand what's under the hood of `java.util`.

### 1. Reverse a Linked List
"I need to reverse the pointers. I’ll keep track of three nodes: `prev`, `current`, and `next`.
In each iteration, I point `current.next` backwards to `prev`, then shift everything forward. O(N) time."

```java
public ListNode reverseList(ListNode head) {
    ListNode prev = null;
    ListNode current = head;
    
    while (current != null) {
        ListNode nextTemp = current.next; // Save next
        current.next = prev;              // Reverse pointer
        prev = current;                   // Move prev
        current = nextTemp;               // Move current
    }
    return prev; // New head
}
```

---

### 2. Detect a Cycle in a Linked List
"I’ll use Floyd’s Cycle-Finding Algorithm (Tortoise and Hare).
I have a `slow` pointer moving 1 step and a `fast` pointer moving 2 steps.
If there is a loop, `fast` will eventually catch `slow` and they will overlap. If `fast` reaches null, there is no loop."

```java
public boolean hasCycle(ListNode head) {
    if (head == null) return false;
    ListNode slow = head;
    ListNode fast = head;
    
    while (fast != null && fast.next != null) {
        slow = slow.next;        // 1 step
        fast = fast.next.next;   // 2 steps
        
        if (slow == fast) return true; // Collision
    }
    return false;
}
```

---

### 3. Binary Tree Inorder Traversal (Left-Root-Right)
"Inorder traversal gives values in sorted order for a BST.
I’ll use a standard recursion. It’s elegant and only takes 3 lines."

```java
public void inorder(TreeNode root) {
    if (root == null) return;
    
    inorder(root.left);            // Left
    System.out.print(root.val);    // Root
    inorder(root.right);           // Right
}
```

---

### 4. Implement a Stack using Arrays
"I need an array and a `top` pointer.
`push`: increment top, insert value.
`pop`: return value, decrement top.
I added a check for stack overflow."

```java
class MyStack {
    int[] arr = new int[100];
    int top = -1;
    
    void push(int x) {
        if (top == 99) throw new StackOverflowError();
        arr[++top] = x;
    }
    
    int pop() {
        if (top == -1) throw new EmptyStackException();
        return arr[top--];
    }
}
```

---

### 5. Find Maximum Depth of Binary Tree
"I’ll use recursion. The depth of a tree is `1 + max(depth of left, depth of right)`.
The base case is when the node is null, depth is 0."

```java
public int maxDepth(TreeNode root) {
    if (root == null) return 0;
    
    int leftDepth = maxDepth(root.left);
    int rightDepth = maxDepth(root.right);
    
    return Math.max(leftDepth, rightDepth) + 1;
}
```

---

### 6. BFS (Level Order Traversal) of a Tree
"Recursion is hard for BFS. I’ll use a `Queue`.
I add the root. Then, while the queue isn't empty, I remove a node, process it, and add its children to the back of the queue."

```java
public void levelOrder(TreeNode root) {
    if (root == null) return;
    Queue<TreeNode> queue = new LinkedList<>();
    queue.add(root);
    
    while (!queue.isEmpty()) {
        TreeNode node = queue.poll();
        System.out.print(node.val + " ");
        
        if (node.left != null) queue.add(node.left);
        if (node.right != null) queue.add(node.right);
    }
}
```
