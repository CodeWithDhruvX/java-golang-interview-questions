import java.util.*;

public class FrequencyCount {

    // Brute Force: O(N^2) nested loops
    public static Map<Integer, Integer> frequencyBruteForce(int[] arr) {
        Map<Integer, Integer> freq = new HashMap<>();
        
        for (int i = 0; i < arr.length; i++) {
            int count = 0;
            for (int j = 0; j < arr.length; j++) {
                if (arr[i] == arr[j]) {
                    count++;
                }
            }
            freq.put(arr[i], count);
        }
        return freq;
    }
    
    // Optimized: HashMap single pass
    public static Map<Integer, Integer> frequencyOptimized(int[] arr) {
        Map<Integer, Integer> freq = new HashMap<>();
        
        for (int num : arr) {
            freq.put(num, freq.getOrDefault(num, 0) + 1);
        }
        return freq;
    }
    
    private static void printMap(Map<Integer, Integer> map) {
        for (Map.Entry<Integer, Integer> entry : map.entrySet()) {
            System.out.println(entry.getKey() + ": " + entry.getValue());
        }
    }

    public static void main(String[] args) {
        int[] testArray = {1, 2, 2, 3, 1, 4, 2, 3, 1};
        
        System.out.println("Find Frequency of Each Element");
        System.out.println("Array: " + java.util.Arrays.toString(testArray));
        
        System.out.println("\nBrute Force:");
        printMap(frequencyBruteForce(testArray));
        
        System.out.println("\nOptimized:");
        printMap(frequencyOptimized(testArray));
    }
}
