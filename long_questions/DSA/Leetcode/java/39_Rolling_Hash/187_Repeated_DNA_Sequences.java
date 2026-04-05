import java.util.*;

public class RepeatedDNASequences {
    
    // 187. Repeated DNA Sequences - Rabin-Karp with Rolling Hash
    // Time: O(N*L) where N is string length, L is pattern length (10), Space: O(N)
    public List<String> findRepeatedDnaSequences(String s) {
        if (s.length() < 10) {
            return new ArrayList<>();
        }
        
        // Hash set to store seen sequences
        Set<String> seen = new HashSet<>();
        Set<String> repeated = new HashSet<>();
        
        for (int i = 0; i <= s.length() - 10; i++) {
            String sequence = s.substring(i, i + 10);
            
            if (seen.contains(sequence)) {
                repeated.add(sequence);
            } else {
                seen.add(sequence);
            }
        }
        
        return new ArrayList<>(repeated);
    }
    
    // Optimized version using rolling hash
    public List<String> findRepeatedDnaSequencesRollingHash(String s) {
        if (s.length() < 10) {
            return new ArrayList<>();
        }
        
        // Map for character to integer (A=0, C=1, G=2, T=3)
        Map<Character, Integer> charMap = new HashMap<>();
        charMap.put('A', 0);
        charMap.put('C', 1);
        charMap.put('G', 2);
        charMap.put('T', 3);
        
        Set<Integer> seen = new HashSet<>();
        Map<Integer, String> repeated = new HashMap<>();
        
        // Base for hash calculation
        int base = 4;
        int windowSize = 10;
        
        // Calculate initial hash
        int hash = 0;
        for (int i = 0; i < windowSize; i++) {
            hash = hash * base + charMap.get(s.charAt(i));
        }
        
        seen.add(hash);
        
        // Rolling hash
        for (int i = 1; i <= s.length() - windowSize; i++) {
            // Remove leftmost character contribution
            int leftChar = charMap.get(s.charAt(i - 1));
            hash -= leftChar * (int) Math.pow(base, windowSize - 1);
            
            // Add new character
            hash = hash * base + charMap.get(s.charAt(i + windowSize - 1));
            
            if (seen.contains(hash)) {
                String sequence = s.substring(i, i + windowSize);
                repeated.put(hash, sequence);
            } else {
                seen.add(hash);
            }
        }
        
        return new ArrayList<>(repeated.values());
    }
    
    // More efficient rolling hash with bit manipulation
    public List<String> findRepeatedDnaSequencesOptimized(String s) {
        if (s.length() < 10) {
            return new ArrayList<>();
        }
        
        Set<String> seen = new HashSet<>();
        Set<String> repeated = new HashSet<>();
        
        // Use bit manipulation for more efficient encoding
        int[] charToBits = new int[26];
        charToBits['A' - 'A'] = 0; // 00
        charToBits['C' - 'A'] = 1; // 01
        charToBits['G' - 'A'] = 2; // 10
        charToBits['T' - 'A'] = 3; // 11
        
        int hash = 0;
        int mask = (1 << 20) - 1; // 10 characters * 2 bits each
        
        // Calculate initial hash
        for (int i = 0; i < 10; i++) {
            hash = (hash << 2) | charToBits[s.charAt(i) - 'A'];
        }
        
        seen.add(s.substring(0, 10));
        
        // Rolling hash
        for (int i = 10; i < s.length(); i++) {
            // Shift left by 2 bits and add new character
            hash = ((hash << 2) & mask) | charToBits[s.charAt(i) - 'A'];
            String sequence = s.substring(i - 9, i + 1);
            
            if (seen.contains(sequence)) {
                repeated.add(sequence);
            } else {
                seen.add(sequence);
            }
        }
        
        return new ArrayList<>(repeated);
    }
    
    // Standard approach with HashMap
    public List<String> findRepeatedDnaSequencesStandard(String s) {
        if (s.length() < 10) {
            return new ArrayList<>();
        }
        
        Map<String, Integer> frequency = new HashMap<>();
        
        // Count all 10-character sequences
        for (int i = 0; i <= s.length() - 10; i++) {
            String sequence = s.substring(i, i + 10);
            frequency.put(sequence, frequency.getOrDefault(sequence, 0) + 1);
        }
        
        // Collect sequences that appear more than once
        List<String> result = new ArrayList<>();
        for (Map.Entry<String, Integer> entry : frequency.entrySet()) {
            if (entry.getValue() > 1) {
                result.add(entry.getKey());
            }
        }
        
        return result;
    }
    
    // Version with detailed explanation
    public class DNAAnalysisResult {
        List<String> repeatedSequences;
        Map<String, Integer> frequency;
        List<String> explanation;
        
        DNAAnalysisResult(List<String> repeatedSequences, Map<String, Integer> frequency, List<String> explanation) {
            this.repeatedSequences = repeatedSequences;
            this.frequency = frequency;
            this.explanation = explanation;
        }
    }
    
    public DNAAnalysisResult analyzeDNASequences(String s) {
        List<String> explanation = new ArrayList<>();
        explanation.add("=== DNA Sequence Analysis ===");
        explanation.add("Input string: " + s);
        explanation.add("Length: " + s.length());
        
        if (s.length() < 10) {
            explanation.add("String too short for 10-character sequences");
            return new DNAAnalysisResult(new ArrayList<>(), new HashMap<>(), explanation);
        }
        
        Map<String, Integer> frequency = new HashMap<>();
        
        explanation.add("Analyzing all 10-character sequences:");
        
        for (int i = 0; i <= s.length() - 10; i++) {
            String sequence = s.substring(i, i + 10);
            int count = frequency.getOrDefault(sequence, 0) + 1;
            frequency.put(sequence, count);
            
            if (count == 1) {
                explanation.add(String.format("  Position %d: %s (first occurrence)", i, sequence));
            } else if (count == 2) {
                explanation.add(String.format("  Position %d: %s (REPEATED!)", i, sequence));
            }
        }
        
        List<String> repeated = new ArrayList<>();
        for (Map.Entry<String, Integer> entry : frequency.entrySet()) {
            if (entry.getValue() > 1) {
                repeated.add(entry.getKey());
                explanation.add(String.format("Repeated sequence: %s (appears %d times)", 
                    entry.getKey(), entry.getValue()));
            }
        }
        
        explanation.add(String.format("Found %d repeated sequences", repeated.size()));
        
        return new DNAAnalysisResult(repeated, frequency, explanation);
    }
    
    // Performance comparison
    public void comparePerformance(String s) {
        System.out.println("=== Performance Comparison ===");
        System.out.println("Input length: " + s.length());
        
        // Standard approach
        long startTime = System.nanoTime();
        List<String> result1 = findRepeatedDnaSequencesStandard(s);
        long endTime = System.nanoTime();
        System.out.printf("Standard approach: %d results (took %d ns)\n", 
            result1.size(), endTime - startTime);
        
        // Optimized approach
        startTime = System.nanoTime();
        List<String> result2 = findRepeatedDnaSequencesOptimized(s);
        endTime = System.nanoTime();
        System.out.printf("Optimized approach: %d results (took %d ns)\n", 
            result2.size(), endTime - startTime);
        
        // Rolling hash approach
        startTime = System.nanoTime();
        List<String> result3 = findRepeatedDnaSequencesRollingHash(s);
        endTime = System.nanoTime();
        System.out.printf("Rolling hash: %d results (took %d ns)\n", 
            result3.size(), endTime - startTime);
    }
    
    // Find all repeated sequences of any length
    public Map<Integer, List<String>> findAllRepeatedSequences(String s, int minLength, int maxLength) {
        Map<Integer, List<String>> allRepeated = new HashMap<>();
        
        for (int length = minLength; length <= maxLength; length++) {
            if (s.length() < length) {
                continue;
            }
            
            Map<String, Integer> frequency = new HashMap<>();
            
            for (int i = 0; i <= s.length() - length; i++) {
                String sequence = s.substring(i, i + length);
                frequency.put(sequence, frequency.getOrDefault(sequence, 0) + 1);
            }
            
            List<String> repeated = new ArrayList<>();
            for (Map.Entry<String, Integer> entry : frequency.entrySet()) {
                if (entry.getValue() > 1) {
                    repeated.add(entry.getKey());
                }
            }
            
            if (!repeated.isEmpty()) {
                allRepeated.put(length, repeated);
            }
        }
        
        return allRepeated;
    }
    
    public static void main(String[] args) {
        RepeatedDNASequences rna = new RepeatedDNASequences();
        
        // Test cases
        String[] testCases = {
            "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT",
            "AAAAAAAAAAAAA",
            "ACGTACGTACGTACGT",
            "CCCCCCCCCCCCCCCCCCCCCCCC",
            "AGTCAGTCAGTCAGTCAGTC",
            "AAAAAAAAAA",
            "ACGTACGTAC",
            "",
            "A",
            "ACGTACGTACGTACGTACGTACGT"
        };
        
        String[] descriptions = {
            "Standard case with repeats",
            "All A's",
            "Repeating pattern",
            "All C's",
            "Alternating pattern",
            "Exactly 10 A's",
            "Shorter than 10",
            "Empty string",
            "Single character",
            "Long repeating pattern"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, descriptions[i]);
            System.out.printf("Input: \"%s\"\n", testCases[i]);
            
            List<String> result1 = rna.findRepeatedDnaSequences(testCases[i]);
            List<String> result2 = rna.findRepeatedDnaSequencesOptimized(testCases[i]);
            List<String> result3 = rna.findRepeatedDnaSequencesRollingHash(testCases[i]);
            
            System.out.printf("Standard: %s\n", result1);
            System.out.printf("Optimized: %s\n", result2);
            System.out.printf("Rolling Hash: %s\n", result3);
            System.out.println();
        }
        
        // Detailed analysis
        System.out.println("=== Detailed Analysis ===");
        String testString = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT";
        DNAAnalysisResult analysis = rna.analyzeDNASequences(testString);
        
        System.out.printf("Repeated sequences: %s\n", analysis.repeatedSequences);
        System.out.println("Explanation:");
        for (String step : analysis.explanation) {
            System.out.println("  " + step);
        }
        
        // Performance comparison
        System.out.println("\n=== Performance Comparison ===");
        String performanceTest = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT".repeat(100);
        rna.comparePerformance(performanceTest);
        
        // Find all repeated sequences
        System.out.println("\n=== All Repeated Sequences ===");
        Map<Integer, List<String>> allRepeated = rna.findAllRepeatedSequences(testString, 5, 15);
        for (Map.Entry<Integer, List<String>> entry : allRepeated.entrySet()) {
            System.out.printf("Length %d: %s\n", entry.getKey(), entry.getValue());
        }
        
        // Edge cases
        System.out.println("\n=== Edge Cases ===");
        System.out.printf("All same character: %s\n", 
            rna.findRepeatedDnaSequences("AAAAAAAAAA"));
        System.out.printf("No repeats: %s\n", 
            rna.findRepeatedDnaSequences("ACGTACGTACGT"));
        System.out.printf("Exactly one repeat: %s\n", 
            rna.findRepeatedDnaSequences("AAAAAAAAAACGTACGT"));
    }

/*
=======================================
PATTERN RECOGNITION & INSIGHTS
=======================================

## 1. ALGORITHM PATTERN: Rolling Hash (Rabin-Karp)
- **Sliding Window**: Efficient hash computation for substrings
- **Rolling Update**: Update hash in O(1) when window slides
- **Collision Avoidance**: Use large base and modulo
- **Bit Manipulation**: Encode characters efficiently

## 2. PROBLEM CHARACTERISTICS
- **DNA Sequences**: Find repeated 10-character sequences
- **Fixed Length**: All sequences are exactly 10 characters
- **Alphabet**: Only 4 characters (A, C, G, T)
- **Efficiency**: Need faster than O(N*L) substring approach

## 3. SIMILAR PROBLEMS
- Find All Anagrams in String
- Longest Duplicate Substring
- Repeated Substring Patterns
- String Matching with Wildcards

## 4. KEY OBSERVATIONS
- Rolling hash enables O(N) time vs O(N*L) substring
- Each character can be encoded with 2 bits (4 possibilities)
- 10-character sequence needs 20 bits (10 * 2)
- Hash collisions possible but unlikely with good base
- Time complexity: O(N) with rolling hash vs O(N*L) naive

## 5. VARIATIONS & EXTENSIONS
- Different sequence lengths
- Different character sets
- Multiple pattern matching
- Hash collision handling

## 6. INTERVIEW INSIGHTS
- Clarify: "Are sequences always 10 characters?"
- Edge cases: short strings, no repeats, all same characters
- Time complexity: O(N) vs O(N*L) naive
- Space complexity: O(N) for hash sets

## 7. COMMON MISTAKES
- Incorrect hash calculation (overflow issues)
- Wrong modulo operation
- Improper bit encoding
- Not handling collisions
- Off-by-one errors in window sliding

## 8. OPTIMIZATION STRATEGIES
- Use bit manipulation for character encoding
- Use large base to minimize collisions
- Use efficient hash set operations
- Precompute powers for rolling updates

## 9. EXECUTION VISUALIZATION

## 10. HUMAN LOGIC PHASE

### Mental Model & Intuition
**Think of it like fingerprinting DNA sequences:**
- You have a long DNA string and need to find repeated patterns
- Each 10-character sequence is like a fingerprint
- Instead of storing all fingerprints, use rolling hash for efficiency
- Rolling hash is like sliding a window and updating fingerprint incrementally
- This avoids recomputing the entire fingerprint each time

### Step-by-Step Human Reasoning

#### Phase 1: Problem Understanding
1. **Input**: DNA string, find repeated 10-character sequences
2. **Goal**: Return all sequences that appear more than once
3. **Output**: List of repeated sequences

#### Phase 2: Key Insight Recognition
- **"What's the bottleneck?"** → O(N*L) substring creation
- **"How to optimize?"** → Use rolling hash for O(1) updates
- **"Why rolling hash?"** → Avoid recomputing hash for each window
- **"How to encode characters?"** → Use 2 bits for 4 DNA characters

#### Phase 3: Strategy Development
```
Human thought process:
"I'll use rolling hash:
1. Encode A=0, C=1, G=2, T=3 (2 bits each)
2. Calculate initial hash for first 10 characters
3. For each position:
   - Remove leftmost character contribution
   - Shift hash left by 2 bits
   - Add new character
   - Check if hash seen before
4. Use bit mask to keep only last 20 bits
5. Store sequences that appear multiple times"
```

#### Phase 4: Edge Case Handling
- **Short strings**: Return empty list
- **No repeats**: Return empty list
- **All same characters**: Handle efficiently
- **Empty string**: Return empty list

#### Phase 5: Algorithm Walkthrough (Human Perspective)
```
DNA: "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"

Human thinking:
"Let's apply rolling hash:

Character encoding:
A=0 (00), C=1 (01), G=2 (10), T=3 (11)

Initial hash for first 10 chars "AAAAACCCCC":
hash = 0*4^9 + 0*4^8 + ... + 1*4^0
hash = 0 + 0 + 0 + 0 + 0*4^5 + 1*4^4 + 1*4^3 + 1*4^2 + 1*4^1 + 1*4^0
hash = 0 + 0 + 0 + 0 + 0 + 1024 + 256 + 64 + 16 + 4 + 1 = 1365

Seen: {1365}, Repeated: {}

Slide window to position 1:
Remove 'A' (0) contribution: 0 * 4^9 = 0
hash = 1365 - 0 = 1365
Shift left: hash = (1365 << 2) & mask
Add 'A' (0): hash = hash | 0
New hash for "AAAACCCCCA" ✓

Continue sliding...
When we encounter same hash again:
Add sequence to repeated set

Final repeated sequences: ["AAAAACCCCC", "CCCCCAAAAA"] ✓

Manual check:
"AAAAACCCCC" appears at positions 0 and 10 ✓
"CCCCCAAAAA" appears at positions 5 and 15 ✓"
```

#### Phase 6: Intuition Validation
- **Why it works**: Rolling hash maintains correct fingerprint
- **Why it's efficient**: O(N) vs O(N*L) substring approach
- **Why it's correct**: Each sequence processed exactly once

### Common Human Pitfalls & How to Avoid Them
1. **"Why not just use substrings?"** → O(N*L) too slow
2. **"What about hash collisions?"** → Use large base to minimize
3. **"How to handle encoding?"** → Use 2 bits for DNA characters
4. **"What about overflow?"** → Use appropriate data types

### Real-World Analogy
**Like finding repeated patterns in genomic data:**
- You have a long DNA sequence and need to find repeated patterns
- Each 10-character sequence is like a genetic marker
- Rolling hash is like scanning the DNA with a sliding window
- Each window generates a fingerprint (hash) of the pattern
- When you see the same fingerprint again, you've found a repeat
- This is used in bioinformatics for pattern discovery
- Useful in DNA analysis, plagiarism detection, data mining

### Human-Readable Pseudocode
```
function findRepeatedDNASequences(s):
    if s.length < 10: return []
    
    // Character encoding
    charMap = {'A':0, 'C':1, 'G':2, 'T':3}
    base = 4
    windowSize = 10
    mask = (1 << 20) - 1 // 20 bits for 10 chars
    
    seen = {}
    repeated = {}
    
    // Calculate initial hash
    hash = 0
    for i from 0 to windowSize-1:
        hash = hash * base + charMap[s[i]]
    
    seen.add(hash)
    
    // Rolling hash
    for i from 1 to s.length - windowSize:
        // Remove leftmost character
        leftChar = charMap[s[i-1]]
        hash -= leftChar * (base ^ (windowSize-1))
        
        // Add new character
        hash = hash * base + charMap[s[i + windowSize - 1]]
        
        // Keep only last 20 bits
        hash = hash & mask
        
        // Check for repetition
        if hash in seen:
            sequence = s.substring(i, i + windowSize)
            repeated.add(sequence)
        else:
            seen.add(hash)
    
    return list(repeated)
```

### Execution Visualization

### Example: DNA="AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
```
Rolling Hash Process:

Character encoding: A=0, C=1, G=2, T=3
Base: 4, Window: 10, Mask: 2^20-1

Initial window [0,9]: "AAAAACCCCC"
Hash calculation:
hash = 0*4^9 + 0*4^8 + 0*4^7 + 0*4^6 + 0*4^5
hash += 1*4^4 + 1*4^3 + 1*4^2 + 1*4^1 + 1*4^0
hash = 0 + 0 + 0 + 0 + 0 + 1024 + 256 + 64 + 16 + 4 + 1 = 1365

Seen: {1365}, Repeated: {}

Slide to position 1, window [1,10]: "AAAACCCCCA"
Remove 'A' (0) from position 0: hash -= 0 * 4^9 = 1365
Shift left: hash = (1365 << 2) & mask
Add 'A' (0): hash = hash | 0
New hash = 5460

Seen: {1365, 5460}, Repeated: {}

Continue sliding...
When we reach position 10, window [10,19]: "CCCCCAAAAA"
Hash = 1365 (same as initial)
Seen already contains 1365 → Add "CCCCCAAAAA" to repeated

Final repeated: ["AAAAACCCCC", "CCCCCAAAAA"] ✓

Visualization:
Window slides one position at a time
Hash updates in O(1) using rolling technique
Repetitions detected when hash seen before
```

### Key Visualization Points:
- **Rolling Update**: Remove leftmost, shift, add new character
- **Bit Encoding**: 2 bits per character for DNA
- **Hash Collision**: Unlikely with good base selection
- **Efficiency**: O(N) time vs O(N*L) substring approach

### Memory Layout Visualization:
```
DNA: "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
Window size: 10

Hash evolution:
Pos 0-9: "AAAAACCCCC" → hash=1365
Pos 1-10: "AAAACCCCCA" → hash=5460
Pos 2-11: "AAACCCCCAA" → hash=21841
...
Pos 10-19: "CCCCCAAAAA" → hash=1365 (REPEAT!)
Pos 15-24: "AAAAACCCCC" → hash=1365 (REPEAT!)

Seen hashes: {1365, 5460, 21841, ...}
Repeated sequences: ["AAAAACCCCC", "CCCCCAAAAA"]

Rolling hash enables O(1) hash updates
Bit encoding minimizes memory usage
```

### Time Complexity Breakdown:
- **Initial Hash**: O(L) time for first window
- **Rolling Updates**: O(N-L+1) * O(1) = O(N) time
- **Hash Operations**: O(1) per slide
- **Space**: O(N) for hash sets
- **Total**: O(N) time, O(N) space
- **Optimal**: Cannot do better than O(N) for this problem
- **vs Naive**: O(N*L) vs O(N) with rolling hash
*/
