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
}
