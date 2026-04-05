import java.util.*;

public class ImplementTriePrefixTree {
    
    // TrieNode represents a node in trie
    public static class TrieNode {
        Map<Character, TrieNode> children;
        boolean isEnd;
        
        public TrieNode() {
            children = new HashMap<>();
            isEnd = false;
        }
    }

    // Trie represents trie data structure
    public static class Trie {
        private TrieNode root;

        // Constructor creates a new Trie
        public Trie() {
            root = new TrieNode();
        }

        // Insert inserts a word into trie
        // Time: O(L), Space: O(L) where L is word length
        public void insert(String word) {
            TrieNode node = root;
            
            for (char c : word.toCharArray()) {
                if (!node.children.containsKey(c)) {
                    node.children.put(c, new TrieNode());
                }
                node = node.children.get(c);
            }
            
            node.isEnd = true;
        }

        // Search returns true if word is in trie
        // Time: O(L), Space: O(1)
        public boolean search(String word) {
            TrieNode node = root;
            
            for (char c : word.toCharArray()) {
                if (!node.children.containsKey(c)) {
                    return false;
                }
                node = node.children.get(c);
            }
            
            return node.isEnd;
        }

        // StartsWith returns true if there is any word in trie that starts with given prefix
        // Time: O(L), Space: O(1)
        public boolean startsWith(String prefix) {
            TrieNode node = root;
            
            for (char c : prefix.toCharArray()) {
                if (!node.children.containsKey(c)) {
                    return false;
                }
                node = node.children.get(c);
            }
            
            return true;
        }

        // Helper function to visualize trie structure
        public void visualize() {
            System.out.println("Trie Structure:");
            visualizeHelper(root, "", true);
        }

        private void visualizeHelper(TrieNode node, String prefix, boolean isLast) {
            if (node == null) {
                return;
            }
            
            // Print children
            int count = 0;
            for (Map.Entry<Character, TrieNode> entry : node.children.entrySet()) {
                count++;
                char c = entry.getKey();
                TrieNode child = entry.getValue();
                
                String newPrefix = prefix;
                if (isLast) {
                    newPrefix += "    ";
                } else {
                    newPrefix += "│   ";
                }
                
                System.out.printf("%s└── %c\n", newPrefix, c);
                visualizeHelper(child, newPrefix + "    ", count == node.children.size());
            }
            
            // Print current node
            if (node.isEnd) {
                System.out.printf("%s└── (END)\n", prefix);
            }
        }

        // Additional useful methods

        // GetAllWords returns all words in trie
        public List<String> getAllWords() {
            List<String> words = new ArrayList<>();
            getAllWordsHelper(root, "", words);
            return words;
        }

        private void getAllWordsHelper(TrieNode node, String current, List<String> words) {
            if (node.isEnd) {
                words.add(current);
            }
            
            for (Map.Entry<Character, TrieNode> entry : node.children.entrySet()) {
                getAllWordsHelper(entry.getValue(), current + entry.getKey(), words);
            }
        }

        // CountWords returns number of words in trie
        public int countWords() {
            return countWordsHelper(root);
        }

        private int countWordsHelper(TrieNode node) {
            int count = 0;
            if (node.isEnd) {
                count++;
            }
            
            for (TrieNode child : node.children.values()) {
                count += countWordsHelper(child);
            }
            
            return count;
        }

        // Delete removes a word from trie
        public boolean delete(String word) {
            return deleteHelper(root, word, 0);
        }

        private boolean deleteHelper(TrieNode node, String word, int depth) {
            if (depth == word.length()) {
                if (!node.isEnd) {
                    return false; // Word doesn't exist
                }
                node.isEnd = false;
                return node.children.size() == 0;
            }
            
            char c = word.charAt(depth);
            TrieNode child = node.children.get(c);
            if (child == null) {
                return false; // Word doesn't exist
            }
            
            boolean shouldDeleteChild = deleteHelper(child, word, depth + 1);
            if (shouldDeleteChild) {
                node.children.remove(c);
                return node.children.size() == 0 && !node.isEnd;
            }
            
            return false;
        }
    }

    public static void main(String[] args) {
        // Test cases
        Trie trie = new Trie();
        
        // Test insertion and search
        System.out.println("=== Testing Insert and Search ===");
        String[] words = {"apple", "app", "application", "apt", "bat", "batch"};
        for (String word : words) {
            trie.insert(word);
            System.out.printf("Inserted: %s\n", word);
        }
        
        // Test search
        String[] searchWords = {"apple", "app", "appl", "bat", "batter"};
        for (String word : searchWords) {
            boolean found = trie.search(word);
            System.out.printf("Search %s: %b\n", word, found);
        }
        
        // Test starts with
        System.out.println("\n=== Testing StartsWith ===");
        String[] prefixes = {"app", "ap", "bat", "cat"};
        for (String prefix : prefixes) {
            boolean found = trie.startsWith(prefix);
            System.out.printf("StartsWith %s: %b\n", prefix, found);
        }
        
        // Test additional methods
        System.out.println("\n=== Additional Methods ===");
        System.out.printf("Total words: %d\n", trie.countWords());
        
        List<String> allWords = trie.getAllWords();
        System.out.printf("All words: %s\n", allWords);
        
        // Test deletion
        System.out.println("\n=== Testing Delete ===");
        trie.delete("app");
        System.out.printf("After deleting 'app', Search 'app': %b\n", trie.search("app"));
        System.out.printf("After deleting 'app', Search 'apple': %b\n", trie.search("apple"));
        System.out.printf("Total words after deletion: %d\n", trie.countWords());
        
        // Test edge cases
        System.out.println("\n=== Testing Edge Cases ===");
        Trie emptyTrie = new Trie();
        System.out.printf("Empty trie search 'test': %b\n", emptyTrie.search("test"));
        System.out.printf("Empty trie startsWith 'test': %b\n", emptyTrie.startsWith("test"));
        
        emptyTrie.insert("");
        System.out.printf("Insert empty string, count: %d\n", emptyTrie.countWords());
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Trie (Prefix Tree)
- **Trie Structure**: Tree-like data structure for efficient prefix operations
- **Character Nodes**: Each node contains children for next characters
- **Prefix Search**: Efficiently check if any word starts with prefix
- **Space Efficiency**: Shared prefixes save memory

## 2. PROBLEM CHARACTERISTICS
- **String Operations**: Insert, search, prefix matching
- **Character-based**: Each level represents one character
- **Prefix Tree**: Common prefixes are shared among words
- **Dynamic Operations**: Support insert, delete, search operations

## 3. SIMILAR PROBLEMS
- Implement Trie (Word Dictionary)
- Word Search II
- Design Add and Search Words
- Replace Words

## 4. KEY OBSERVATIONS
- Each node contains map of children characters
- isEnd flag marks complete words
- Prefix search only needs to traverse, not check isEnd
- Insert creates nodes as needed for new words
- Delete requires careful backtracking

## 5. VARIATIONS & EXTENSIONS
- Array-based children (26 letters for lowercase)
- Compressed trie (radix tree)
- Support for Unicode characters
- Different character sets (digits, uppercase)

## 6. INTERVIEW INSIGHTS
- Clarify: "What character set should I support?"
- Edge cases: empty string, duplicate words, single character
- Time complexity: O(L) vs O(N*L) naive search
- Space complexity: O(total characters) vs O(N*L) separate storage

## 7. COMMON MISTAKES
- Using array of size 26 instead of HashMap for sparse data
- Not handling isEnd flag correctly
- Memory leaks in delete operation
- Not considering null children

## 8. OPTIMIZATION STRATEGIES
- Use array for fixed character sets (faster)
- Use HashMap for sparse character sets (memory efficient)
- Lazy deletion (mark as deleted)
- Path compression for long common prefixes

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like organizing words in a dictionary:**
- You have words that share common prefixes
- Instead of storing each word separately, share prefixes
- Each level represents one character position
- Navigate down the tree to spell words
- Common prefixes are shared among multiple words

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: Collection of words to store
2. **Goal**: Support insert, search, prefix operations
3. **Output**: Efficient prefix tree implementation

#### Phase 2: Key Insight Recognition
- **"How to share prefixes?"** → Tree structure where common prefixes are shared
- **"What operations needed?"** → Insert, search, prefix check, delete
- **"How to navigate?"** → Follow character by character path
- **"When does word end?"** → isEnd flag marks complete words

#### Phase 3: Strategy Development
```
Human thought process:
"I'll implement a trie:
1. Root node has no character, just children
2. Each node has:
   - Map of children (character → node)
   - isEnd flag (marks complete word)
3. Insert: traverse/create nodes for each character
4. Search: follow path, check isEnd at end
5. Prefix: follow path, success if path exists"
```

#### Phase 4: Edge Case Handling
- **Empty string**: Handle specially (no characters to traverse)
- **Single character**: Simple one-level tree
- **Duplicate words**: Share existing nodes
- **Non-existent words**: Return false gracefully

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
Words: ["apple", "app", "application"]

Human thinking:
"Let's build the trie step by step:

Insert "apple":
- Root → 'a' → 'p' → 'p' → 'l' → 'e'
- Mark 'e' node asEnd = true

Insert "app":
- Root → 'a' → 'p' → 'p'
- 'p' node already exists, use it
- Mark 'p' node asEnd = true

Insert "application":
- Root → 'a' → 'p' → 'p' → 'l' → 'i' → 'c' → 'a' → 't' → 'i' → 'o' → 'n'
- Mark 'n' node asEnd = true

Search "app":
- Root → 'a' → 'p' → 'p'
- 'p' node has isEnd = true ✓

Search "appl":
- Root → 'a' → 'p' → 'p' → 'l'
- 'l' node exists but isEnd = false ✗

Prefix "app":
- Root → 'a' → 'p' → 'p'
- Path exists ✓ (don't need to check isEnd)"
```

#### Phase 6: Intuition Validation
- **Why it works**: Tree structure naturally shares common prefixes
- **Why it's efficient**: O(L) operations vs O(N*L) naive
- **Why it's correct**: Each word has unique path in tree

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just use array?"** → O(N*L) search vs O(L) trie
2. **"What about HashMap?"** → Can't efficiently handle prefix search
3. **"How to handle delete?"** → Complex backtracking needed
4. **"What about memory?"** → Shared prefixes save space

### Real-World Analogy
**Like organizing contacts in a phone:**
- You have contacts (words) with shared prefixes
- Instead of storing each contact separately, organize by first letter
- Within each letter, organize by second letter, etc.
- When searching for "John", you quickly find all "J" contacts
- Common prefixes like "App" are shared among "Apple", "Application"
- This is exactly how phone contacts are organized

### Human-Readable Pseudocode
```
class Trie:
    root = new TrieNode()
    
    function insert(word):
        node = root
        for character in word:
            if character not in node.children:
                node.children[character] = new TrieNode()
            node = node.children[character]
        node.isEnd = true
    
    function search(word):
        node = root
        for character in word:
            if character not in node.children:
                return false
            node = node.children[character]
        return node.isEnd
    
    function startsWith(prefix):
        node = root
        for character in prefix:
            if character not in node.children:
                return false
            node = node.children[character]
        return true
```

### Execution Visualization

### Example: Words = ["apple", "app", "application"]
```
Trie Structure:
Root
└── 'a'
    └── 'p'
        ├── 'p' (END) → "app"
        └── 'l'
            └── 'e' (END) → "apple"
            └── 'i'
                └── 'c'
                    └── 'a'
                        └── 't'
                            └── 'i'
                                └── 'o'
                                    └── 'n' (END) → "application"

Operations:
Search "app": Root→'a'→'p'→'p' ✓ (END found)
Search "appl": Root→'a'→'p'→'p'→'l' ✗ (no END)
Prefix "app": Root→'a'→'p'→'p' ✓ (path exists)
```

### Key Visualization Points:
- **Tree structure** naturally shares common prefixes
- **isEnd flags** mark complete words
- **Character navigation** follows word paths
- **Prefix search** only needs path existence

### Memory Layout Visualization:
```
Node Structure:
TrieNode {
    children: Map<Character, TrieNode>
    isEnd: boolean
}

Insert "apple":
Root.children['a'] = new Node()
Node['a'].children['p'] = new Node()
Node['a']['p'].children['p'] = new Node()
Node['a']['p']['p'].children['l'] = new Node()
Node['a']['p']['p']['l'].children['e'] = new Node()
Node['a']['p']['p']['l']['e'].isEnd = true

Shared prefixes save memory:
"app" and "apple" share: Root→'a'→'p'→'p'
Only "apple" continues: →'l'→'e'
```

### Time Complexity Breakdown:
- **Insert**: O(L) where L is word length
- **Search**: O(L) where L is word length
- **Prefix**: O(P) where P is prefix length
- **Space**: O(total characters) for all words
- **Optimal**: Much better than O(N*L) naive approaches
*/
