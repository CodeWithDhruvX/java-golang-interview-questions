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
}
