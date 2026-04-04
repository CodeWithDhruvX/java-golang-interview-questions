public class KokoEatingBananas {
    
    // 875. Koko Eating Bananas
    // Time: O(N log M), Space: O(1) where M is max(piles)
    public int minEatingSpeed(int[] piles, int h) {
        int left = 1;
        int right = maxPiles(piles);
        int result = right;
        
        while (left <= right) {
            int mid = left + (right - left) / 2;
            
            if (canEatAll(piles, h, mid)) {
                result = mid;
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }
        
        return result;
    }
    
    private boolean canEatAll(int[] piles, int h, int speed) {
        int hours = 0;
        for (int pile : piles) {
            hours += (pile + speed - 1) / speed; // Equivalent to Math.ceil(pile / speed)
            if (hours > h) {
                return false;
            }
        }
        return hours <= h;
    }
    
    private int maxPiles(int[] piles) {
        int max = 0;
        for (int pile : piles) {
            if (pile > max) {
                max = pile;
            }
        }
        return max;
    }
    
    // Alternative approach with detailed explanation
    public int minEatingSpeedDetailed(int[] piles, int h) {
        System.out.println("=== Binary Search on Answer Space ===");
        System.out.printf("Piles: %s, Hours: %d\n", java.util.Arrays.toString(piles), h);
        
        int left = 1;
        int right = maxPiles(piles);
        int result = right;
        
        System.out.printf("Search space: [%d, %d]\n", left, right);
        
        while (left <= right) {
            int mid = left + (right - left) / 2;
            System.out.printf("Trying speed: %d bananas/hour\n", mid);
            
            int hoursNeeded = calculateHoursNeeded(piles, mid);
            System.out.printf("  Hours needed: %d\n", hoursNeeded);
            
            if (hoursNeeded <= h) {
                result = mid;
                right = mid - 1;
                System.out.printf("  ✓ Can finish in time, searching lower\n");
            } else {
                left = mid + 1;
                System.out.printf("  ✗ Cannot finish, need higher speed\n");
            }
        }
        
        System.out.printf("Final result: %d bananas/hour\n", result);
        return result;
    }
    
    private int calculateHoursNeeded(int[] piles, int speed) {
        int hours = 0;
        for (int pile : piles) {
            int hoursForPile = (pile + speed - 1) / speed;
            hours += hoursForPile;
            System.out.printf("    Pile %d: %d hours (%.1f hours)\n", 
                pile, hoursForPile, (double) pile / speed);
        }
        return hours;
    }
    
    // Brute force approach for comparison
    public int minEatingSpeedBruteForce(int[] piles, int h) {
        for (int speed = 1; speed <= maxPiles(piles); speed++) {
            if (canEatAll(piles, h, speed)) {
                return speed;
            }
        }
        return maxPiles(piles);
    }
    
    // Optimized version with early termination
    public int minEatingSpeedOptimized(int[] piles, int h) {
        int left = 1;
        int right = maxPiles(piles);
        int result = right;
        
        // Early optimization: if h == piles.length, answer is maxPiles
        if (h == piles.length) {
            return right;
        }
        
        // Early optimization: if h >= sum(piles), answer is 1
        int totalBananas = 0;
        for (int pile : piles) {
            totalBananas += pile;
        }
        if (h >= totalBananas) {
            return 1;
        }
        
        while (left <= right) {
            int mid = left + (right - left) / 2;
            
            if (canEatAllOptimized(piles, h, mid)) {
                result = mid;
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }
        
        return result;
    }
    
    private boolean canEatAllOptimized(int[] piles, int h, int speed) {
        int hours = 0;
        for (int pile : piles) {
            hours += (pile + speed - 1) / speed;
            
            // Early termination
            if (hours > h) {
                return false;
            }
        }
        return true;
    }
    
    // Version that returns the eating schedule
    public class EatingSchedule {
        int minSpeed;
        java.util.List<java.util.List<Integer>> hourlySchedule;
        int totalHours;
        
        EatingSchedule(int minSpeed, java.util.List<java.util.List<Integer>> hourlySchedule, int totalHours) {
            this.minSpeed = minSpeed;
            this.hourlySchedule = hourlySchedule;
            this.totalHours = totalHours;
        }
    }
    
    public EatingSchedule getEatingSchedule(int[] piles, int h) {
        int minSpeed = minEatingSpeed(piles, h);
        java.util.List<java.util.List<Integer>> schedule = new java.util.ArrayList<>();
        int totalHours = 0;
        
        for (int pile : piles) {
            int remaining = pile;
            java.util.List<Integer> hourlyIntake = new java.util.ArrayList<>();
            
            while (remaining > 0) {
                int eat = Math.min(minSpeed, remaining);
                hourlyIntake.add(eat);
                remaining -= eat;
                totalHours++;
            }
            
            schedule.add(hourlyIntake);
        }
        
        return new EatingSchedule(minSpeed, schedule, totalHours);
    }
    
    // Mathematical approach to find lower bound
    public int minEatingSpeedMathematical(int[] piles, int h) {
        int totalBananas = 0;
        int maxPile = 0;
        
        for (int pile : piles) {
            totalBananas += pile;
            maxPile = Math.max(maxPile, pile);
        }
        
        // Lower bound: ceil(totalBananas / h)
        int lowerBound = (totalBananas + h - 1) / h;
        int upperBound = maxPile;
        
        System.out.printf("Mathematical bounds: [%d, %d]\n", lowerBound, upperBound);
        
        int left = Math.max(1, lowerBound);
        int right = upperBound;
        int result = right;
        
        while (left <= right) {
            int mid = left + (right - left) / 2;
            
            if (canEatAll(piles, h, mid)) {
                result = mid;
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }
        
        return result;
    }
    
    public static void main(String[] args) {
        KokoEatingBananas koko = new KokoEatingBananas();
        
        // Test cases
        int[][][] testCases = {
            {{3, 6, 7, 11}, {8}},
            {{30, 11, 23, 4, 20}, {5}},
            {{30, 11, 23, 4, 20}, {6}},
            {{1}, {1}},
            {{100}, {1}},
            {{100}, {100}},
            {{312884470}, {312884469}},
            {{1, 1, 1, 1, 1, 1, 1, 1, 1}, {10}},
            {{5, 10, 15, 20, 25}, {15}}
        };
        
        String[] descriptions = {
            "Standard case",
            "Tight deadline",
            "Relaxed deadline",
            "Single banana",
            "Large pile, tight deadline",
            "Large pile, relaxed deadline",
            "Very large pile",
            "Many small piles",
            "Increasing piles"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, descriptions[i]);
            int[] piles = testCases[i][0];
            int h = testCases[i][1][0];
            
            int result1 = koko.minEatingSpeed(piles, h);
            int result2 = koko.minEatingSpeedBruteForce(piles, h);
            int result3 = koko.minEatingSpeedOptimized(piles, h);
            int result4 = koko.minEatingSpeedMathematical(piles, h);
            
            System.out.printf("  Binary Search: %d\n", result1);
            System.out.printf("  Brute Force: %d\n", result2);
            System.out.printf("  Optimized: %d\n", result3);
            System.out.printf("  Mathematical: %d\n", result4);
            
            // Get eating schedule
            EatingSchedule schedule = koko.getEatingSchedule(piles, h);
            System.out.printf("  Eating Schedule (speed %d):\n", schedule.minSpeed);
            for (int j = 0; j < schedule.hourlySchedule.size(); j++) {
                System.out.printf("    Pile %d: %s\n", j + 1, schedule.hourlySchedule.get(j));
            }
            System.out.printf("  Total hours: %d\n", schedule.totalHours);
            System.out.println();
        }
        
        // Detailed explanation for one case
        System.out.println("=== Detailed Explanation ===");
        int[] detailedPiles = {30, 11, 23, 4, 20};
        int detailedHours = 6;
        koko.minEatingSpeedDetailed(detailedPiles, detailedHours);
        
        // Performance comparison
        System.out.println("\n=== Performance Comparison ===");
        
        // Large test case
        int[] largePiles = new int[1000];
        for (int i = 0; i < 1000; i++) {
            largePiles[i] = (i % 100) + 1;
        }
        
        long startTime = System.nanoTime();
        int largeResult = koko.minEatingSpeed(largePiles, 500);
        long endTime = System.nanoTime();
        
        System.out.printf("Large test (1000 piles, 500 hours): %d (took %d ns)\n", 
            largeResult, endTime - startTime);
        
        startTime = System.nanoTime();
        int bruteResult = koko.minEatingSpeedBruteForce(largePiles, 500);
        endTime = System.nanoTime();
        
        System.out.printf("Brute force result: %d (took %d ns)\n", 
            bruteResult, endTime - startTime);
        
        // Edge cases
        System.out.println("\n=== Edge Cases ===");
        
        // Minimum possible speed
        int[] minSpeedCase = {1, 1, 1, 1, 1};
        System.out.printf("Minimum speed case: %d\n", koko.minEatingSpeed(minSpeedCase, 5));
        
        // Maximum possible speed
        int[] maxSpeedCase = {1000000, 1000000, 1000000};
        System.out.printf("Maximum speed case: %d\n", koko.minEatingSpeed(maxSpeedCase, 3));
        
        // Very tight deadline
        int[] tightDeadline = {100, 200, 300, 400, 500};
        System.out.printf("Tight deadline case: %d\n", koko.minEatingSpeed(tightDeadline, 5));
    }
}
