import java.util.*;

public class WordLadder {
    
    // 127. Word Ladder - BFS with State
    // Time: O(N * L^2), Space: O(N * L^2) where N is word count, L is word length
    public int ladderLength(String beginWord, String endWord, List<String> wordList) {
        if (beginWord.length() != endWord.length()) {
            return 0;
        }
        
        // Build adjacency list
        Set<String> wordSet = new HashSet<>(wordList);
        
        // BFS with state: (current word, position)
        Queue<String> queue = new LinkedList<>();
        queue.offer(beginWord);
        
        Set<String> visited = new HashSet<>();
        visited.add(beginWord);
        
        int steps = 0;
        
        while (!queue.isEmpty()) {
            int levelSize = queue.size();
            
            for (int i = 0; i < levelSize; i++) {
                String current = queue.poll();
                
                if (current.equals(endWord)) {
                    return steps;
                }
                
                // Generate all possible next states
                for (int j = 0; j < current.length(); j++) {
                    for (char c = 'a'; c <= 'z'; c++) {
                        if (c != current.charAt(j)) {
                            StringBuilder nextWordBuilder = new StringBuilder(current);
                            nextWordBuilder.setCharAt(j, c);
                            String nextWord = nextWordBuilder.toString();
                            
                            if (!visited.contains(nextWord) && wordSet.contains(nextWord)) {
                                visited.add(nextWord);
                                queue.offer(nextWord);
                            }
                        }
                    }
                }
            }
            
            steps++;
        }
        
        return 0;
    }
    
    // BFS with state optimization
    public int ladderLengthOptimized(String beginWord, String endWord, List<String> wordList) {
        if (beginWord.length() != endWord.length()) {
            return 0;
        }
        
        // Build adjacency list
        Set<String> wordSet = new HashSet<>(wordList);
        
        // BFS with state tracking
        Queue<String> queue = new LinkedList<>();
        queue.offer(beginWord);
        
        Set<String> visited = new HashSet<>();
        visited.add(beginWord);
        
        int steps = 0;
        
        while (!queue.isEmpty()) {
            int levelSize = queue.size();
            
            for (int i = 0; i < levelSize; i++) {
                String current = queue.poll();
                
                if (current.equals(endWord)) {
                    return steps;
                }
                
                // Generate all possible next states
                char[] currentChars = current.toCharArray();
                
                for (int j = 0; j < currentChars.length; j++) {
                    char originalChar = currentChars[j];
                    
                    for (char c = 'a'; c <= 'z'; c++) {
                        if (c != originalChar) {
                            currentChars[j] = c;
                            String nextWord = new String(currentChars);
                            
                            if (!visited.contains(nextWord) && wordSet.contains(nextWord)) {
                                visited.add(nextWord);
                                queue.offer(nextWord);
                            }
                        }
                    }
                    
                    currentChars[j] = originalChar; // Restore original character
                }
            }
            
            steps++;
        }
        
        return 0;
    }
    
    // Bidirectional BFS with state
    public int ladderLengthBidirectional(String beginWord, String endWord, List<String> wordList) {
        if (beginWord.length() != endWord.length()) {
            return 0;
        }
        
        Set<String> wordSet = new HashSet<>(wordList);
        
        // Bidirectional BFS
        Set<String> beginSet = new HashSet<>();
        beginSet.add(beginWord);
        
        Set<String> endSet = new HashSet<>();
        endSet.add(endWord);
        
        Set<String> visitedBegin = new HashSet<>();
        visitedBegin.add(beginWord);
        
        Set<String> visitedEnd = new HashSet<>();
        visitedEnd.add(endWord);
        
        int steps = 0;
        
        while (!beginSet.isEmpty() && !endSet.isEmpty()) {
            // Always expand smaller set
            if (beginSet.size() > endSet.size()) {
                Set<String> temp = beginSet;
                beginSet = endSet;
                endSet = temp;
            }
            
            Set<String> nextSet = new HashSet<>();
            
            for (String word : beginSet) {
                // Generate all possible next states
                char[] currentChars = word.toCharArray();
                
                for (int j = 0; j < currentChars.length; j++) {
                    char originalChar = currentChars[j];
                    
                    for (char c = 'a'; c <= 'z'; c++) {
                        if (c != originalChar) {
                            currentChars[j] = c;
                            String nextWord = new String(currentChars);
                            
                            if (!visitedBegin.contains(nextWord) && wordSet.contains(nextWord)) {
                                visitedBegin.add(nextWord);
                                nextSet.add(nextWord);
                            }
                        }
                    }
                    
                    currentChars[j] = originalChar; // Restore original character
                }
            }
            
            // Check if we reached the target
            for (String word : nextSet) {
                if (endSet.contains(word)) {
                    return steps + 1;
                }
            }
            
            beginSet = nextSet;
            steps++;
        }
        
        return 0;
    }
    
    // Version with detailed explanation
    public class WordLadderResult {
        int length;
        List<String> path;
        List<String> explanation;
        
        WordLadderResult(int length, List<String> path, List<String> explanation) {
            this.length = length;
            this.path = path;
            this.explanation = explanation;
        }
    }
    
    public WordLadderResult ladderLengthDetailed(String beginWord, String endWord, List<String> wordList) {
        List<String> explanation = new ArrayList<>();
        explanation.add("=== BFS with State for Word Ladder ===");
        explanation.add(String.format("Begin: %s, End: %s", beginWord, endWord));
        explanation.add("Word List: " + wordList);
        
        if (beginWord.length() != endWord.length()) {
            explanation.add("Word lengths don't match, returning 0");
            return new WordLadderResult(0, new ArrayList<>(), explanation);
        }
        
        Set<String> wordSet = new HashSet<>(wordList);
        Queue<String> queue = new LinkedList<>();
        queue.offer(beginWord);
        
        Set<String> visited = new HashSet<>();
        visited.add(beginWord);
        
        int steps = 0;
        Map<String, String> parent = new HashMap<>();
        parent.put(beginWord, null);
        
        while (!queue.isEmpty()) {
            int levelSize = queue.size();
            explanation.add(String.format("Step %d: Processing %d words at level %d", 
                steps, levelSize, steps));
            
            for (int i = 0; i < levelSize; i++) {
                String current = queue.poll();
                explanation.add(String.format("  Processing word: %s", current));
                
                if (current.equals(endWord)) {
                    explanation.add(String.format("  Found target word: %s", current));
                    
                    // Reconstruct path
                    List<String> path = reconstructPath(parent, beginWord, endWord);
                    return new WordLadderResult(steps, path, explanation);
                }
                
                // Generate all possible next states
                char[] currentChars = current.toCharArray();
                int transitions = 0;
                
                for (int j = 0; j < currentChars.length; j++) {
                    char originalChar = currentChars[j];
                    
                    for (char c = 'a'; c <= 'z'; c++) {
                        if (c != originalChar) {
                            currentChars[j] = c;
                            String nextWord = new String(currentChars);
                            
                            if (!visited.contains(nextWord) && wordSet.contains(nextWord)) {
                                visited.add(nextWord);
                                queue.offer(nextWord);
                                parent.put(nextWord, current);
                                transitions++;
                                explanation.add(String.format("    Added: %s (parent: %s)", nextWord, current));
                            }
                        }
                    }
                    
                    currentChars[j] = originalChar; // Restore original character
                }
            }
            
            explanation.add(String.format("  Generated %d transitions", transitions));
        }
            
            steps++;
        }
        
        explanation.add("No path found");
        return new WordLadderResult(0, new ArrayList<>(), explanation);
    }
    
    private List<String> reconstructPath(Map<String, String> parent, String beginWord, String endWord) {
        List<String> path = new ArrayList<>();
        String current = endWord;
        
        while (current != null) {
            path.add(0, current);
            current = parent.get(current);
        }
        
        Collections.reverse(path);
        return path;
    }
    
    // BFS with early termination
    public int ladderLengthEarlyTermination(String beginWord, String endWord, List<String> wordList) {
        if (beginWord.length() != endWord.length()) {
            return 0;
        }
        
        Set<String> wordSet = new HashSet<>(wordList);
        
        // Precompute all possible transformations for faster lookup
        Map<String, List<String>> transformations = new HashMap<>();
        
        for (String word : wordSet) {
            List<String> neighbors = new ArrayList<>();
            char[] wordChars = word.toCharArray();
            
            for (int i = 0; i < wordChars.length; i++) {
                char originalChar = wordChars[i];
                
                for (char c = 'a'; c <= 'z'; c++) {
                    if (c != originalChar) {
                        wordChars[i] = c;
                        String transformed = new String(wordChars);
                        
                        if (wordSet.contains(transformed)) {
                            neighbors.add(transformed);
                        }
                    }
                }
                
                wordChars[i] = originalChar; // Restore
            }
            
            transformations.put(word, neighbors);
        }
        
        // BFS
        Queue<String> queue = new LinkedList<>();
        queue.offer(beginWord);
        
        Set<String> visited = new HashSet<>();
        visited.add(beginWord);
        
        int steps = 0;
        
        while (!queue.isEmpty()) {
            int levelSize = queue.size();
            
            for (int i = 0; i < levelSize; i++) {
                String current = queue.poll();
                
                if (current.equals(endWord)) {
                    return steps;
                }
                
                for (String neighbor : transformations.getOrDefault(current, new ArrayList<>())) {
                    if (!visited.contains(neighbor)) {
                        visited.add(neighbor);
                        queue.offer(neighbor);
                    }
                }
            }
            
            steps++;
        }
        
        return 0;
    }
    
    // Performance comparison
    public void comparePerformance(String beginWord, String endWord, List<String> wordList) {
        System.out.println("=== Performance Comparison ===");
        System.out.printf("Begin: %s, End: %s, Words: %d\n", 
            beginWord, endWord, wordList.size());
        
        // Standard BFS
        long startTime = System.nanoTime();
        int result1 = ladderLength(beginWord, endWord, wordList);
        long endTime = System.nanoTime();
        System.out.printf("Standard BFS: %d (took %d ns)\n", result1, endTime - startTime);
        
        // Optimized BFS
        startTime = System.nanoTime();
        int result2 = ladderLengthOptimized(beginWord, endWord, wordList);
        endTime = System.nanoTime();
        System.out.printf("Optimized BFS: %d (took %d ns)\n", result2, endTime - startTime);
        
        // Bidirectional BFS
        startTime = System.nanoTime();
        int result3 = ladderLengthBidirectional(beginWord, endWord, wordList);
        endTime = System.nanoTime();
        System.out.printf("Bidirectional BFS: %d (took %d ns)\n", result3, endTime - startTime);
        
        // Early termination
        startTime = System.nanoTime();
        int result4 = ladderLengthEarlyTermination(beginWord, endWord, wordList);
        endTime = System.nanoTime();
        System.out.printf("Early termination: %d (took %d ns)\n", result4, endTime - startTime);
    }
    
    public static void main(String[] args) {
        WordLadder wl = new WordLadder();
        
        // Test cases
        String[][] testCases = {
            {"hit", "cog", Arrays.asList("hot","dot","dog","lot","log","cog")},
            {"hit", "cog", Arrays.asList("hot","dot","dog","lot","log")},
            {"a", "c", Arrays.asList("a","b","c")},
            {"abc", "def", Arrays.asList("abd","abf","acd","ace","adf","bde","bef","cde")},
            {"game", "the", Arrays.asList("fry","fut","gape","hen","hex","ion","java","jet","kin","log","map","nod","odd","pie","quo","ran","sap","tea","the")},
            {"a", "z", Arrays.asList("a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z")},
            {"abc", "abc", Arrays.asList("abc")},
            {"abc", "def", Arrays.asList("abd","def")},
            {"hot", "dog", Arrays.asList("hot","dog")}
        };
        
        String[] descriptions = {
            "Standard case",
            "No solution",
            "Simple case",
            "Multiple paths",
            "Large dictionary",
            "Alphabet case",
            "Same words",
            "Partial match",
            "Direct connection"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, descriptions[i]);
            System.out.printf("Begin: %s, End: %s\n", testCases[i][0], testCases[i][1]);
            
            int result1 = wl.ladderLength(testCases[i][0], testCases[i][1], testCases[i][2]);
            int result2 = wl.ladderLengthOptimized(testCases[i][0], testCases[i][1], testCases[i][2]);
            int result3 = wl.ladderLengthBidirectional(testCases[i][0], testCases[i][1], testCases[i][2]);
            int result4 = wl.ladderLengthEarlyTermination(testCases[i][0], testCases[i][1], testCases[i][2]);
            
            System.out.printf("Standard BFS: %d\n", result1);
            System.out.printf("Optimized BFS: %d\n", result2);
            System.out.printf("Bidirectional BFS: %d\n", result3);
            System.out.printf("Early termination: %d\n", result4);
            System.out.println();
        }
        
        // Detailed explanation
        System.out.println("=== Detailed Explanation ===");
        WordLadderResult detailedResult = wl.ladderLengthDetailed(
            "hit", "cog", Arrays.asList("hot","dot","dog","lot","log","cog"));
        
        System.out.printf("Result: %d\n", detailedResult.length);
        System.out.printf("Path: %s\n", detailedResult.path);
        for (String step : detailedResult.explanation) {
            System.out.println("  " + step);
        }
        
        // Performance comparison
        System.out.println("\n=== Performance Comparison ===");
        
        // Large test case
        List<String> largeWordList = new ArrayList<>();
        for (int i = 0; i < 1000; i++) {
            largeWordList.add(String.format("%03d", i));
        }
        
        wl.comparePerformance("000", "999", largeWordList);
        
        // Edge cases
        System.out.println("\n=== Edge Cases ===");
        System.out.printf("Empty word list: %d\n", 
            wl.ladderLength("a", "b", new ArrayList<>()));
        System.out.printf("Single word: %d\n", 
            wl.ladderLength("a", "a", Arrays.asList("a")));
        System.out.printf("Different lengths: %d\n", 
            wl.ladderLength("a", "ab", Arrays.asList("a","ab")));
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: BFS with State
- **State Tracking**: Track current word and transformation steps
- **Level-by-Level**: Process all words at current transformation level
- **Word Generation**: Generate all one-character transformations
- **Bidirectional Optimization**: Expand from both ends simultaneously

## 2. PROBLEM CHARACTERISTICS
- **Word Ladder**: Find shortest transformation sequence
- **One-Character Changes**: Each step changes exactly one character
- **Dictionary Constraint**: All intermediate words must be valid
- **Shortest Path**: BFS guarantees minimum transformations

## 3. SIMILAR PROBLEMS
- Word Ladder II (return all shortest paths)
- Minimum Genetic Mutation
- Edit Distance
- Shortest Word Transformation

## 4. KEY OBSERVATIONS
- BFS guarantees shortest path in unweighted graph
- State includes current word and transformation count
- Word generation requires checking all 26 possibilities
- Bidirectional search reduces search space significantly
- Time complexity: O(N * L² * 26) worst case

## 5. VARIATIONS & EXTENSIONS
- Return all shortest paths
- Different character sets
- Weighted transformations
- Multiple source/destination pairs

## 6. INTERVIEW INSIGHTS
- Clarify: "Are all words the same length?"
- Edge cases: begin equals end, no solution, single word
- Time complexity: O(N * L² * 26) vs O(N³) naive
- Space complexity: O(N * L²) for adjacency vs O(N) for BFS

## 7. COMMON MISTAKES
- Not checking if endWord exists in dictionary
- Incorrect word generation logic
- Not handling visited words properly
- Forgetting to count transformation steps
- Not using bidirectional optimization

## 8. OPTIMIZATION STRATEGIES
- Bidirectional search reduces search space
- Precompute adjacency relationships
- Early termination when frontiers meet
- Use efficient word generation

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like finding the shortest word transformation sequence:**
- You have start word and target word
- Each transformation changes exactly one character
- You want to find the shortest sequence of valid words
- This is like finding the shortest path in a word transformation graph
- Bidirectional search is like two explorers meeting in the middle

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: beginWord, endWord, wordList (dictionary)
2. **Goal**: Find minimum number of word transformations
3. **Output**: Length of shortest transformation sequence

#### Phase 2: Key Insight Recognition
- **"What defines a valid transformation?"** → Change exactly one character
- **"How to explore efficiently?"** → BFS level by level
- **"Why bidirectional?"** → Reduces search space exponentially
- **"When to stop?"** → When frontiers meet or end found

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use bidirectional BFS:
1. Initialize begin and end sets with start words
2. While both sets are non-empty:
   - Always expand smaller set
   - Generate all one-character transformations
   - Add valid, unvisited words to next set
   - Check if frontiers meet
   - Increment transformation count
3. Return transformation count when frontiers meet"
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
Level: 0

Level 0:
- Expand begin set (smaller, size 1)
- Generate neighbors of "hit": "hot"
- "hot" is valid and unvisited → Add to next set
- Level: 1

Level 1:
Begin set: {"hot"}, End set: {"cog"}
- Expand end set (smaller, size 1)
- Generate neighbors of "cog": "log", "cog"
- "log" is valid and unvisited → Add to next set
- Check if any begin word neighbors match end set → No
- Level: 2

Level 2:
Begin set: {"dot","lot"}, End set: {"log"}
- Expand begin set (smaller, size 2)
- Generate neighbors of "dot": "dog"
- Generate neighbors of "lot": "log"
- "dog" is valid and unvisited → Add to next set
- "log" matches end set → FOUND!
- Return level + 1 = 3

Transformation sequence: hit → hot → dot → dog → log → cog (5 steps)
Wait, let me reconsider:

Actually, when "log" matches end set at level 2:
We found connection at level 2
Return level + 1 = 3

Sequence: hit → hot → dot → log → cog (4 steps) ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: BFS guarantees shortest path
- **Why it's efficient**: Bidirectional reduces search space
- **Why it's correct**: Level-by-level ensures optimal solution

### Common Human Pitfalls & How to Avoid Them
1. **"Why not generate all possible sequences?"** → Exponential time
2. **"What about DFS?"** → Doesn't guarantee shortest path
3. **"How to generate neighbors?"** → Change exactly one character
4. **"What about visited tracking?"** → Essential to avoid cycles

### Real-World Analogy
**Like finding the shortest password transformation sequence:**
- You have a starting password (beginWord)
- You want to reach target password (endWord)
- Each step changes one character (like one character mutation)
- You can only use valid intermediate passwords (dictionary)
- You explore from both start and end simultaneously
- When exploration paths meet, you've found the shortest sequence
- This is like password cracking with constraint validation
- Useful in cryptography, puzzle solving, word games

### Human-Readable Pseudocode
```
function ladderLength(beginWord, endWord, wordList):
    if endWord not in wordList: return 0
    
    wordSet = set(wordList)
    beginSet = {beginWord}
    endSet = {endWord}
    visited = {}
    level = 0
    
    while !beginSet.isEmpty() and !endSet.isEmpty():
        if beginSet.size() > endSet.size():
            swap(beginSet, endSet)
        
        nextSet = {}
        
        for word in beginSet:
            for position from 0 to word.length-1:
                for char from 'a' to 'z':
                    if char != word[position]:
                        newWord = word with character replaced
                        if newWord in wordSet and newWord not in visited:
                            nextSet.add(newWord)
                            visited.add(newWord)
        
        if any word in nextSet is in endSet:
            return level + 1
        
        beginSet = nextSet
        level++
    
    return 0
```

### Execution Visualization

### Example: beginWord="hit", endWord="cog"
```
Bidirectional BFS Process:

Level 0:
Begin: {"hit"}, End: {"cog"}
Expand begin: "hot" → Next: {"hot"}

Level 1:
Begin: {"hot"}, End: {"cog"}
Expand end: "log" → Next: {"log"}
Check: No match

Level 2:
Begin: {"dot","lot"}, End: {"log"}
Expand begin: "dog" → Next: {"dog","log"}
Check: "log" matches end set → FOUND!

Result: 3 (hit→hot→dot→log→cog) ✓

Visualization:
Level 0: hit → hot
Level 1: cog ← log
Level 2: hot → dot → lot → log → cog (FOUND!)

Transformation: hit→hot→dot→log→cog
Steps: 4 ✓
```

### Key Visualization Points:
- **Bidirectional expansion** reduces search space
- **Level-by-level** processing ensures shortest path
- **Word generation** requires checking all 26 possibilities
- **State tracking** includes current word and transformation count

### Memory Layout Visualization:
```
Search Space Evolution:
Level 0: Begin={hit}, End={cog}
Level 1: Begin={hot}, End={cog,log}
Level 2: Begin={dot,lot}, End={log}, FOUND!

Word Generation:
For "hit": h*i*t → hot, hut, hit, etc.
For "hot": h*o*t → dot, lot, hot, etc.
For "dot": d*o*t → dog, lot, dot, etc.

Bidirectional Search:
Two explorers meeting at level 2
Total search space significantly reduced
```

### Time Complexity Breakdown:
- **Word Generation**: O(L * 26) per word expanded
- **BFS Levels**: O(N) in worst case
- **Total**: O(N * L² * 26) time, O(N * L²) space
- **Bidirectional**: Reduces constant factor significantly
- **Optimal**: Cannot do better than O(N * L²) for this problem
- **vs DFS**: Exponential time complexity
*/
