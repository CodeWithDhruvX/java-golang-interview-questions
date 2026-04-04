import java.util.*;

public class WordLadder {
    
    // 127. Word Ladder - Meet in the Middle Approach
    // Time: O(N * L^2), Space: O(N * L) where N is word count, L is word length
    public int ladderLength(String beginWord, String endWord, List<String> wordList) {
        // Check if endWord exists in wordList
        Set<String> wordSet = new HashSet<>(wordList);
        
        if (!wordSet.contains(endWord)) {
            return 0;
        }
        
        // Meet in the middle BFS
        Set<String> beginSet = new HashSet<>();
        beginSet.add(beginWord);
        
        Set<String> endSet = new HashSet<>();
        endSet.add(endWord);
        
        Set<String> visited = new HashSet<>();
        int level = 1;
        
        while (!beginSet.isEmpty() && !endSet.isEmpty()) {
            // Always expand the smaller set
            if (beginSet.size() > endSet.size()) {
                Set<String> temp = beginSet;
                beginSet = endSet;
                endSet = temp;
            }
            
            Set<String> nextSet = new HashSet<>();
            
            for (String word : beginSet) {
                // Generate all possible next words
                char[] wordChars = word.toCharArray();
                
                for (int i = 0; i < wordChars.length; i++) {
                    char originalChar = wordChars[i];
                    
                    for (char c = 'a'; c <= 'z'; c++) {
                        if (c == originalChar) {
                            continue;
                        }
                        
                        wordChars[i] = c;
                        String newWord = new String(wordChars);
                        
                        if (endSet.contains(newWord)) {
                            return level + 1;
                        }
                        
                        if (wordSet.contains(newWord) && !visited.contains(newWord)) {
                            visited.add(newWord);
                            nextSet.add(newWord);
                        }
                    }
                    
                    wordChars[i] = originalChar; // Restore original character
                }
            }
            
            beginSet = nextSet;
            level++;
        }
        
        return 0;
    }
    
    // Standard BFS approach for comparison
    public int ladderLengthBFS(String beginWord, String endWord, List<String> wordList) {
        Set<String> wordSet = new HashSet<>(wordList);
        
        if (!wordSet.contains(endWord)) {
            return 0;
        }
        
        Queue<String> queue = new LinkedList<>();
        queue.offer(beginWord);
        
        Set<String> visited = new HashSet<>();
        visited.add(beginWord);
        
        int level = 1;
        
        while (!queue.isEmpty()) {
            int levelSize = queue.size();
            
            for (int i = 0; i < levelSize; i++) {
                String word = queue.poll();
                
                if (word.equals(endWord)) {
                    return level;
                }
                
                // Generate all possible next words
                char[] wordChars = word.toCharArray();
                
                for (int j = 0; j < wordChars.length; j++) {
                    char originalChar = wordChars[j];
                    
                    for (char c = 'a'; c <= 'z'; c++) {
                        if (c == originalChar) {
                            continue;
                        }
                        
                        wordChars[j] = c;
                        String newWord = new String(wordChars);
                        
                        if (wordSet.contains(newWord) && !visited.contains(newWord)) {
                            visited.add(newWord);
                            queue.offer(newWord);
                        }
                    }
                    
                    wordChars[j] = originalChar; // Restore original character
                }
            }
            
            level++;
        }
        
        return 0;
    }
    
    // Bidirectional BFS with detailed explanation
    public class LadderResult {
        int length;
        List<String> path;
        List<String> explanation;
        
        LadderResult(int length, List<String> path, List<String> explanation) {
            this.length = length;
            this.path = path;
            this.explanation = explanation;
        }
    }
    
    public LadderResult ladderLengthDetailed(String beginWord, String endWord, List<String> wordList) {
        List<String> explanation = new ArrayList<>();
        
        Set<String> wordSet = new HashSet<>(wordList);
        
        if (!wordSet.contains(endWord)) {
            explanation.add("End word not found in word list");
            return new LadderResult(0, new ArrayList<>(), explanation);
        }
        
        explanation.add("Starting bidirectional BFS");
        explanation.add("Begin set: " + beginWord);
        explanation.add("End set: " + endWord);
        
        Set<String> beginSet = new HashSet<>();
        beginSet.add(beginWord);
        
        Set<String> endSet = new HashSet<>();
        endSet.add(endWord);
        
        Set<String> visited = new HashSet<>();
        int level = 1;
        
        while (!beginSet.isEmpty() && !endSet.isEmpty()) {
            explanation.add(String.format("Level %d: Begin set size: %d, End set size: %d", 
                level, beginSet.size(), endSet.size()));
            
            // Always expand the smaller set
            if (beginSet.size() > endSet.size()) {
                Set<String> temp = beginSet;
                beginSet = endSet;
                endSet = temp;
                explanation.add("Swapped sets - expanding smaller set");
            }
            
            Set<String> nextSet = new HashSet<>();
            
            for (String word : beginSet) {
                explanation.add("Expanding word: " + word);
                
                char[] wordChars = word.toCharArray();
                
                for (int i = 0; i < wordChars.length; i++) {
                    char originalChar = wordChars[i];
                    
                    for (char c = 'a'; c <= 'z'; c++) {
                        if (c == originalChar) {
                            continue;
                        }
                        
                        wordChars[i] = c;
                        String newWord = new String(wordChars);
                        
                        if (endSet.contains(newWord)) {
                            explanation.add(String.format("Found connection: %s -> %s", word, newWord));
                            return new LadderResult(level + 1, new ArrayList<>(), explanation);
                        }
                        
                        if (wordSet.contains(newWord) && !visited.contains(newWord)) {
                            visited.add(newWord);
                            nextSet.add(newWord);
                            explanation.add(String.format("Added to next set: %s", newWord));
                        }
                    }
                    
                    wordChars[i] = originalChar; // Restore original character
                }
            }
            
            beginSet = nextSet;
            level++;
        }
        
        explanation.add("No path found");
        return new LadderResult(0, new ArrayList<>(), explanation);
    }
    
    // Optimized version with preprocessing
    public int ladderLengthOptimized(String beginWord, String endWord, List<String> wordList) {
        Set<String> wordSet = new HashSet<>(wordList);
        
        if (!wordSet.contains(endWord)) {
            return 0;
        }
        
        // Preprocess: create adjacency map
        Map<String, List<String>> adjacencyMap = new HashMap<>();
        
        for (String word : wordList) {
            for (int i = 0; i < word.length(); i++) {
                String pattern = word.substring(0, i) + "*" + word.substring(i + 1);
                adjacencyMap.computeIfAbsent(pattern, k -> new ArrayList<>()).add(word);
            }
        }
        
        // Add beginWord to the map
        for (int i = 0; i < beginWord.length(); i++) {
            String pattern = beginWord.substring(0, i) + "*" + beginWord.substring(i + 1);
            adjacencyMap.computeIfAbsent(pattern, k -> new ArrayList<>()).add(beginWord);
        }
        
        // Bidirectional BFS
        Set<String> beginSet = new HashSet<>();
        beginSet.add(beginWord);
        
        Set<String> endSet = new HashSet<>();
        endSet.add(endWord);
        
        Set<String> visited = new HashSet<>();
        int level = 1;
        
        while (!beginSet.isEmpty() && !endSet.isEmpty()) {
            if (beginSet.size() > endSet.size()) {
                Set<String> temp = beginSet;
                beginSet = endSet;
                endSet = temp;
            }
            
            Set<String> nextSet = new HashSet<>();
            
            for (String word : beginSet) {
                // Generate neighbors using adjacency map
                for (int i = 0; i < word.length(); i++) {
                    String pattern = word.substring(0, i) + "*" + word.substring(i + 1);
                    
                    for (String neighbor : adjacencyMap.getOrDefault(pattern, new ArrayList<>())) {
                        if (endSet.contains(neighbor)) {
                            return level + 1;
                        }
                        
                        if (!visited.contains(neighbor)) {
                            visited.add(neighbor);
                            nextSet.add(neighbor);
                        }
                    }
                }
            }
            
            beginSet = nextSet;
            level++;
        }
        
        return 0;
    }
    
    // Version that returns the actual path
    public List<String> findLadderPath(String beginWord, String endWord, List<String> wordList) {
        Set<String> wordSet = new HashSet<>(wordList);
        
        if (!wordSet.contains(endWord)) {
            return new ArrayList<>();
        }
        
        // BFS from beginWord
        Map<String, List<String>> graph = new HashMap<>();
        Queue<String> queue = new LinkedList<>();
        queue.offer(beginWord);
        
        Set<String> visited = new HashSet<>();
        visited.add(beginWord);
        
        boolean found = false;
        
        while (!queue.isEmpty() && !found) {
            int levelSize = queue.size();
            Set<String> levelVisited = new HashSet<>();
            
            for (int i = 0; i < levelSize; i++) {
                String word = queue.poll();
                
                char[] wordChars = word.toCharArray();
                
                for (int j = 0; j < wordChars.length; j++) {
                    char originalChar = wordChars[j];
                    
                    for (char c = 'a'; c <= 'z'; c++) {
                        if (c == originalChar) {
                            continue;
                        }
                        
                        wordChars[j] = c;
                        String newWord = new String(wordChars);
                        
                        if (wordSet.contains(newWord)) {
                            if (!visited.contains(newWord)) {
                                levelVisited.add(newWord);
                                queue.offer(newWord);
                            }
                            
                            graph.computeIfAbsent(word, k -> new ArrayList<>()).add(newWord);
                            
                            if (newWord.equals(endWord)) {
                                found = true;
                            }
                        }
                    }
                    
                    wordChars[j] = originalChar;
                }
            }
            
            visited.addAll(levelVisited);
        }
        
        if (!found) {
            return new ArrayList<>();
        }
        
        // DFS to find path
        return dfsPath(beginWord, endWord, graph, new HashSet<>());
    }
    
    private List<String> dfsPath(String current, String target, 
                                 Map<String, List<String>> graph, Set<String> visited) {
        if (current.equals(target)) {
            List<String> path = new ArrayList<>();
            path.add(target);
            return path;
        }
        
        visited.add(current);
        
        for (String neighbor : graph.getOrDefault(current, new ArrayList<>())) {
            if (!visited.contains(neighbor)) {
                List<String> path = dfsPath(neighbor, target, graph, visited);
                if (!path.isEmpty()) {
                    path.add(0, current);
                    return path;
                }
            }
        }
        
        visited.remove(current);
        return new ArrayList<>();
    }
    
    public static void main(String[] args) {
        WordLadder wl = new WordLadder();
        
        // Test cases
        List<String> wordList1 = Arrays.asList("hot","dot","dog","lot","log","cog");
        List<String> wordList2 = Arrays.asList("hot","dot","dog","lot","log");
        List<String> wordList3 = Arrays.asList("a","b","c");
        List<String> wordList4 = Arrays.asList("abc","abd","acd","bcd","abce");
        
        Object[][] testCases = {
            {"hit", "cog", wordList1, "Standard case"},
            {"hit", "cog", wordList2, "No solution"},
            {"a", "c", wordList3, "Single character"},
            {"abc", "abce", wordList4, "Prefix change"},
            {"hit", "hot", Arrays.asList("hot"), "Direct neighbor"},
            {"hit", "hit", Arrays.asList("hit"), "Same word"},
            {"abc", "def", Arrays.asList("abd","abf","def"), "Multiple steps"}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, testCases[i][3]);
            String beginWord = (String) testCases[i][0];
            String endWord = (String) testCases[i][1];
            List<String> wordList = (List<String>) testCases[i][2];
            
            int result1 = wl.ladderLength(beginWord, endWord, wordList);
            int result2 = wl.ladderLengthBFS(beginWord, endWord, wordList);
            int result3 = wl.ladderLengthOptimized(beginWord, endWord, wordList);
            
            System.out.printf("  Bidirectional BFS: %d\n", result1);
            System.out.printf("  Standard BFS: %d\n", result2);
            System.out.printf("  Optimized: %d\n", result3);
            
            // Find actual path
            List<String> path = wl.findLadderPath(beginWord, endWord, wordList);
            if (!path.isEmpty()) {
                System.out.printf("  Path: %s\n", path);
            }
            System.out.println();
        }
        
        // Detailed explanation
        System.out.println("=== Detailed Explanation ===");
        LadderResult detailedResult = wl.ladderLengthDetailed("hit", "cog", wordList1);
        System.out.printf("Result length: %d\n", detailedResult.length);
        for (String step : detailedResult.explanation) {
            System.out.printf("  %s\n", step);
        }
        
        // Performance comparison
        System.out.println("\n=== Performance Comparison ===");
        
        // Large test case
        List<String> largeWordList = new ArrayList<>();
        for (int i = 0; i < 1000; i++) {
            largeWordList.add(String.format("%03d", i));
        }
        
        long startTime = System.nanoTime();
        int largeResult1 = wl.ladderLength("000", "999", largeWordList);
        long endTime = System.nanoTime();
        
        System.out.printf("Large test - Bidirectional: %d (took %d ns)\n", 
            largeResult1, endTime - startTime);
        
        startTime = System.nanoTime();
        int largeResult2 = wl.ladderLengthBFS("000", "999", largeWordList);
        endTime = System.nanoTime();
        
        System.out.printf("Large test - Standard BFS: %d (took %d ns)\n", 
            largeResult2, endTime - startTime);
    }
}
