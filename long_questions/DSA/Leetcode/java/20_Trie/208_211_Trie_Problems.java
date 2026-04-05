import java.util.*;

public class ImplementTrie {
    
    // 208. Implement Trie (Prefix Tree)
    // Time: O(L) for insert, search, startsWith, Space: O(N * L) where N is number of words, L is average length
    public static class Trie {
        private TrieNode root;
        
        public Trie() {
            root = new TrieNode();
        }
        
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
        
        public boolean search(String word) {
            TrieNode node = searchNode(word);
            return node != null && node.isEnd;
        }
        
        public boolean startsWith(String prefix) {
            return searchNode(prefix) != null;
        }
        
        private TrieNode searchNode(String word) {
            TrieNode node = root;
            for (char c : word.toCharArray()) {
                if (!node.children.containsKey(c)) {
                    return null;
                }
                node = node.children.get(c);
            }
            return node;
        }
    }
    
    private static class TrieNode {
        Map<Character, TrieNode> children;
        boolean isEnd;
        
        public TrieNode() {
            children = new HashMap<>();
            isEnd = false;
        }
    }

    // 211. Design Add and Search Words Data Structure
    public static class WordDictionary {
        private TrieNode root;
        
        public WordDictionary() {
            root = new TrieNode();
        }
        
        public void addWord(String word) {
            TrieNode node = root;
            for (char c : word.toCharArray()) {
                if (!node.children.containsKey(c)) {
                    node.children.put(c, new TrieNode());
                }
                node = node.children.get(c);
            }
            node.isEnd = true;
        }
        
        public boolean search(String word) {
            return searchHelper(word, 0, root);
        }
        
        private boolean searchHelper(String word, int index, TrieNode node) {
            if (index == word.length()) {
                return node.isEnd;
            }
            
            char c = word.charAt(index);
            
            if (c == '.') {
                for (TrieNode child : node.children.values()) {
                    if (searchHelper(word, index + 1, child)) {
                        return true;
                    }
                }
                return false;
            } else {
                if (!node.children.containsKey(c)) {
                    return false;
                }
                return searchHelper(word, index + 1, node.children.get(c));
            }
        }
    }

    public static void main(String[] args) {
        // Test cases for Trie
        System.out.println("Trie Operations:");
        Trie trie = new Trie();
        trie.insert("apple");
        System.out.printf("Search \"apple\": %b\n", trie.search("apple")); // true
        System.out.printf("Search \"app\": %b\n", trie.search("app")); // false
        System.out.printf("StartsWith \"app\": %b\n", trie.startsWith("app")); // true
        trie.insert("app");
        System.out.printf("Search \"app\" after insert: %b\n", trie.search("app")); // true
        
        // Test cases for WordDictionary
        System.out.println("\nWordDictionary Operations:");
        WordDictionary wordDict = new WordDictionary();
        wordDict.addWord("bad");
        wordDict.addWord("dad");
        wordDict.addWord("mad");
        System.out.printf("Search \"pad\": %b\n", wordDict.search("pad")); // false
        System.out.printf("Search \"bad\": %b\n", wordDict.search("bad")); // true
        System.out.printf("Search \".ad\": %b\n", wordDict.search(".ad")); // true
        System.out.printf("Search \"b..\": %b\n", wordDict.search("b..")); // true
        
        // Additional test cases
        System.out.println("\nAdditional Trie Tests:");
        Trie trie2 = new Trie();
        String[] words = {"hello", "help", "helps", "helmet", "hero", "her"};
        for (String word : words) {
            trie2.insert(word);
        }
        
        System.out.printf("Search \"help\": %b\n", trie2.search("help"));
        System.out.printf("Search \"helps\": %b\n", trie2.search("helps"));
        System.out.printf("Search \"hel\": %b\n", trie2.search("hel"));
        System.out.printf("StartsWith \"he\": %b\n", trie2.startsWith("he"));
        System.out.printf("StartsWith \"hel\": %b\n", trie2.startsWith("hel"));
        System.out.printf("StartsWith \"hero\": %b\n", trie2.startsWith("hero"));
        
        System.out.println("\nAdditional WordDictionary Tests:");
        WordDictionary wordDict2 = new WordDictionary();
        String[] dictWords = {"a", "ab", "abc", "abcd", "abcde"};
        for (String word : dictWords) {
            wordDict2.addWord(word);
        }
        
        System.out.printf("Search \"a\": %b\n", wordDict2.search("a"));
        System.out.printf("Search \"ab\": %b\n", wordDict2.search("ab"));
        System.out.printf("Search \"abc\": %b\n", wordDict2.search("abc"));
        System.out.printf("Search \".\": %b\n", wordDict2.search("."));
        System.out.printf("Search \"..\": %b\n", wordDict2.search(".."));
        System.out.printf("Search \"...\": %b\n", wordDict2.search("..."));
        System.out.printf("Search \"....\": %b\n", wordDict2.search("...."));
        System.out.printf("Search \".....\": %b\n", wordDict2.search("....."));
        System.out.printf("Search \"......\": %b\n", wordDict2.search("......"));
    }
}
