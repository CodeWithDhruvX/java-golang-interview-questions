import java.util.*;

public class Candy {
    
    // 135. Candy
    // Time: O(N), Space: O(N)
    public static int candy(int[] ratings) {
        if (ratings == null || ratings.length == 0) {
            return 0;
        }
        
        int n = ratings.length;
        int[] candies = new int[n];
        Arrays.fill(candies, 1);
        
        // Left to right pass
        for (int i = 1; i < n; i++) {
            if (ratings[i] > ratings[i - 1]) {
                candies[i] = candies[i - 1] + 1;
            }
        }
        
        // Right to left pass
        for (int i = n - 2; i >= 0; i--) {
            if (ratings[i] > ratings[i + 1]) {
                candies[i] = Math.max(candies[i], candies[i + 1] + 1);
            }
        }
        
        int total = 0;
        for (int candy : candies) {
            total += candy;
        }
        
        return total;
    }

    public static void main(String[] args) {
        int[][] testCases = {
            {1, 0, 2},
            {1, 2, 2},
            {1, 2, 3, 4, 5},
            {5, 4, 3, 2, 1},
            {1, 3, 4, 5, 2},
            {2, 2, 2, 2, 2},
            {1},
            {1, 2},
            {2, 1},
            {1, 2, 2, 1},
            {1, 3, 2, 2, 1},
            {3, 2, 1, 2, 3},
            {1, 2, 3, 2, 1, 2, 3},
            {1, 2, 3, 4, 3, 2, 1},
            {1, 1, 2, 2, 3, 3, 2, 2, 1, 1}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int[] ratings = testCases[i];
            int result = candy(ratings);
            
            System.out.printf("Test Case %d: %s -> %d candies\n", 
                i + 1, Arrays.toString(ratings), result);
        }
    }
}
