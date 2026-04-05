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

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Meet in the Middle
- **Bidirectional BFS**: Expand from both begin and end simultaneously
- **Level-by-Level**: Process words by transformation distance
- **Word Transformation**: Change one character at a time
- **Set Expansion**: Always expand smaller frontier first

## 2. PROBLEM CHARACTERISTICS
- **Word Ladder**: Find shortest transformation sequence
- **One-Character Changes**: Each step changes exactly one character
- **Dictionary Constraint**: All intermediate words must be valid
- **Shortest Path**: BFS guarantees minimum transformations

## 3. SIMILAR PROBLEMS
- Word Ladder II (return all paths)
- Minimum Genetic Mutation
- Edit Distance
- Shortest Word Transformation

## 4. KEY OBSERVATIONS
- Bidirectional search reduces search space significantly
- Level-by-level processing ensures shortest path
- Word generation requires checking all 26 possibilities
- Time complexity: O(N * L² * 26) worst case
- Space complexity: O(N * L) for word storage

## 5. VARIATIONS & EXTENSIONS
- Return all shortest paths
- Different character sets
- Weighted transformations
- Multiple word lengths

## 6. INTERVIEW INSIGHTS
- Clarify: "Are all words the same length?"
- Edge cases: begin equals end, no solution, single word
- Time complexity: O(N * L² * 26) vs O(N * L² * 26²) naive
- Space complexity: O(N * L) for sets and queues

## 7. COMMON MISTAKES
- Not checking if endWord exists in dictionary
- Incorrect word generation logic
- Not handling single character words
- Forgetting to mark visited words
- Not expanding smaller frontier first

## 8. OPTIMIZATION STRATEGIES
- Bidirectional search reduces search space
- Preprocess adjacency relationships
- Early termination when frontiers meet
- Use efficient word generation

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the shortest route between two cities:**
- You have cities (words) connected by one-character changes
- Start at beginWord, want to reach endWord
- Each move changes one character (like driving to adjacent city)
- Explore all possible routes simultaneously from both ends
- When routes meet in the middle, you've found the shortest path
- This is like two explorers meeting in the middle

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: beginWord, endWord, wordList (dictionary)
2. **Goal**: Find shortest transformation sequence
3. **Output**: Length of shortest path

#### Phase 2: Key Insight Recognition
- **"What defines a valid step?"** → Change exactly one character
- **"How to explore efficiently?"** → BFS from both ends
- **"Why bidirectional?"** → Reduces search space exponentially
- **"When to stop?"** → When frontiers meet

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use bidirectional BFS:
1. Initialize begin and end sets with start words
2. While both sets are non-empty:
   - Always expand smaller set
   - Generate all one-character transformations
   - Add valid, unvisited words to next set
   - Increment level counter
3. Return level when sets meet"
```

#### Phase 4: Edge Case Handling
- **Begin equals end**: Return 0 or 1 depending on definition
- **End not in dictionary**: Return 0 (no solution)
- **Single character words**: Handle specially
- **Empty word list**: Return 0

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]

Human thinking:
"Let's do bidirectional BFS:

Initialize:
Begin set: {"hit"}
End set: {"cog"}
Level: 1

Level 1:
- Expand smaller set (begin, size 1)
- Generate neighbors of "hit": "hot"
- "hot" is in dictionary and unvisited → Add to next set
- Next set: {"hot"}
- Level: 2

Level 2:
- Expand smaller set (end, size 1)
- Generate neighbors of "cog": "log"
- "log" is in dictionary and unvisited → Add to next set
- Next set: {"log"}
- Level: 3

Level 3:
- Expand smaller set (begin, size 1)
- Generate neighbors of "hot": "dot", "lot"
- Both are valid and unvisited → Add to next set
- Next set: {"dot", "lot"}
- Level: 4

Level 4:
- Expand smaller set (end, size 1)
- Generate neighbors of "log": "cog"
- "cog" is endWord → Return level 4 ✓

Result: 4 transformations (hit→hot→dot→log→cog) ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: BFS guarantees shortest path, bidirectional reduces space
- **Why it's efficient**: Much smaller search space than unidirectional
- **Why it's correct**: Level-by-level ensures optimal solution

### Common Human Pitfalls & How to Avoid Them
1. **"Why not unidirectional BFS?"** → O(26^L) vs O(2*26^(L/2)) bidirectional
2. **"What about DFS?"** → Doesn't guarantee shortest path
3. **"How to generate neighbors?"** → Must change exactly one character
4. **"What about visited tracking?"** → Essential to avoid cycles

### Real-World Analogy
**Like solving a word puzzle:**
- You have a starting word and target word
- Each move changes one letter (like rotating a puzzle piece)
- You can explore moves from both start and end simultaneously
- When your exploration paths meet, you've found the shortest solution
- This is like two people solving a maze from opposite ends
- When they meet, they've found the optimal path

### Human-Readable Pseudocode
```
function ladderLength(beginWord, endWord, wordList):
    if endWord not in wordList:
        return 0
    
    beginSet = {beginWord}
    endSet = {endWord}
    visited = {}
    level = 1
    
    while !beginSet.isEmpty() and !endSet.isEmpty():
        if beginSet.size() > endSet.size():
            nextSet = expandSet(endSet)
            endSet = beginSet
            beginSet = nextSet
        else:
            nextSet = expandSet(beginSet)
            beginSet = endSet
            endSet = nextSet
        
        level++
        visited.addAll(nextSet)
    
    return level
    
function expandSet(wordSet):
    nextSet = {}
    for word in wordSet:
        for each character position in word:
            for each letter from 'a' to 'z':
                if letter != original character:
                    newWord = word with character replaced
                    if newWord in dictionary and not visited:
                        nextSet.add(newWord)
    return nextSet
```

### Execution Visualization

### Example: beginWord="hit", endWord="cog"
```
Bidirectional BFS Process:

Level 1:
Begin: {"hit"}, End: {"cog"}
Expand begin (smaller): "hot" → Next: {"hot"}

Level 2:
Begin: {"hot"}, End: {"cog"}
Expand end (smaller): "log" → Next: {"log"}

Level 3:
Begin: {"hot"}, End: {"log"}
Expand begin (smaller): "dot", "lot" → Next: {"dot", "lot"}

Level 4:
Begin: {"dot", "lot"}, End: {"log"}
Expand end (smaller): "cog" → Found end!

Result: 4 (hit→hot→dot→log→cog) ✓

Visualization:
Level 1: hit → hot
Level 2: cog → log  
Level 3: hot → dot, lot
Level 4: log → cog (FOUND!)

Path: hit→hot→dot→log→cog
```

### Key Visualization Points:
- **Bidirectional expansion** reduces search space
- **Level-by-level** ensures shortest path
- **Word generation** requires checking all 26 letters
- **Set management** tracks visited and frontier words

### Memory Layout Visualization:
```
Search Space Evolution:
Level 1: Begin={hit}, End={cog}
Level 2: Begin={hot}, End={log}
Level 3: Begin={dot,lot}, End={log}
Level 4: Begin={dot,lot}, End={cog} → FOUND!

Word Generation:
For "hot": change each position to 'a'-'z'
For "cog": change each position to 'a'-'z'
Valid neighbors: those in dictionary

Complexity Reduction:
Unidirectional: O(26^L) search space
Bidirectional: O(2*26^(L/2)) search space
```

### Time Complexity Breakdown:
- **Word Generation**: O(L * 26) per word expanded
- **Set Operations**: O(1) average for hash operations
- **Levels**: At most L/2 levels in bidirectional search
- **Total**: O(N * L² * 26) worst case, much better in practice
- **Space**: O(N * L) for word storage and sets
- **Optimal**: Bidirectional significantly reduces search space
- **vs DFS**: Exponential time complexity
*/
