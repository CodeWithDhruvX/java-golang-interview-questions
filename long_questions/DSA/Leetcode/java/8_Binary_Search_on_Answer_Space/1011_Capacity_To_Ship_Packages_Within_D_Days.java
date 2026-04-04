public class CapacityToShipPackagesWithinDDays {
    
    // 1011. Capacity To Ship Packages Within D Days
    // Time: O(N log M), Space: O(1) where M is sum(weights)
    public int shipWithinDays(int[] weights, int days) {
        int left = maxWeight(weights);
        int right = sumWeights(weights);
        int result = right;
        
        while (left <= right) {
            int mid = left + (right - left) / 2;
            
            if (canShip(weights, days, mid)) {
                result = mid;
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }
        
        return result;
    }
    
    private boolean canShip(int[] weights, int days, int capacity) {
        int daysNeeded = 1;
        int currentLoad = 0;
        
        for (int weight : weights) {
            if (currentLoad + weight <= capacity) {
                currentLoad += weight;
            } else {
                daysNeeded++;
                currentLoad = weight;
                if (daysNeeded > days) {
                    return false;
                }
            }
        }
        
        return daysNeeded <= days;
    }
    
    private int maxWeight(int[] weights) {
        int max = 0;
        for (int weight : weights) {
            if (weight > max) {
                max = weight;
            }
        }
        return max;
    }
    
    private int sumWeights(int[] weights) {
        int sum = 0;
        for (int weight : weights) {
            sum += weight;
        }
        return sum;
    }
    
    // Alternative approach with detailed explanation
    public int shipWithinDaysDetailed(int[] weights, int days) {
        System.out.println("=== Binary Search on Answer Space ===");
        System.out.printf("Weights: %s, Days: %d\n", java.util.Arrays.toString(weights), days);
        
        int left = maxWeight(weights);
        int right = sumWeights(weights);
        int result = right;
        
        System.out.printf("Search space: [%d, %d]\n", left, right);
        
        while (left <= right) {
            int mid = left + (right - left) / 2;
            System.out.printf("Trying capacity: %d\n", mid);
            
            if (canShipDetailed(weights, days, mid)) {
                result = mid;
                right = mid - 1;
                System.out.printf("  ✓ Can ship in %d days, searching lower\n", days);
            } else {
                left = mid + 1;
                System.out.printf("  ✗ Cannot ship, need higher capacity\n");
            }
        }
        
        System.out.printf("Final result: %d\n", result);
        return result;
    }
    
    private boolean canShipDetailed(int[] weights, int days, int capacity) {
        int daysNeeded = 1;
        int currentLoad = 0;
        StringBuilder dayLoads = new StringBuilder();
        
        for (int i = 0; i < weights.length; i++) {
            int weight = weights[i];
            
            if (currentLoad + weight <= capacity) {
                currentLoad += weight;
                dayLoads.append(weight).append(" ");
            } else {
                System.out.printf("    Day %d: %s (load: %d)\n", daysNeeded, dayLoads.toString().trim(), currentLoad);
                daysNeeded++;
                currentLoad = weight;
                dayLoads = new StringBuilder();
                dayLoads.append(weight).append(" ");
                
                if (daysNeeded > days) {
                    return false;
                }
            }
        }
        
        System.out.printf("    Day %d: %s (load: %d)\n", daysNeeded, dayLoads.toString().trim(), currentLoad);
        return daysNeeded <= days;
    }
    
    // Brute force approach for comparison
    public int shipWithinDaysBruteForce(int[] weights, int days) {
        int minCapacity = maxWeight(weights);
        int maxCapacity = sumWeights(weights);
        
        for (int capacity = minCapacity; capacity <= maxCapacity; capacity++) {
            if (canShip(weights, days, capacity)) {
                return capacity;
            }
        }
        
        return maxCapacity;
    }
    
    // Optimized version with early termination
    public int shipWithinDaysOptimized(int[] weights, int days) {
        int left = maxWeight(weights);
        int right = sumWeights(weights);
        int result = right;
        
        // Early optimization: if days == weights.length, answer is maxWeight
        if (days == weights.length) {
            return left;
        }
        
        // Early optimization: if days == 1, answer is sumWeights
        if (days == 1) {
            return right;
        }
        
        while (left <= right) {
            int mid = left + (right - left) / 2;
            
            if (canShipOptimized(weights, days, mid)) {
                result = mid;
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }
        
        return result;
    }
    
    private boolean canShipOptimized(int[] weights, int days, int capacity) {
        int daysNeeded = 1;
        int currentLoad = 0;
        
        for (int weight : weights) {
            if (weight > capacity) {
                return false; // Early termination
            }
            
            if (currentLoad + weight <= capacity) {
                currentLoad += weight;
            } else {
                daysNeeded++;
                currentLoad = weight;
                
                if (daysNeeded > days) {
                    return false;
                }
            }
        }
        
        return true;
    }
    
    // Version that returns the actual shipping arrangement
    public class ShippingPlan {
        int minCapacity;
        java.util.List<java.util.List<Integer>> dailyLoads;
        
        ShippingPlan(int minCapacity, java.util.List<java.util.List<Integer>> dailyLoads) {
            this.minCapacity = minCapacity;
            this.dailyLoads = dailyLoads;
        }
    }
    
    public ShippingPlan getShippingPlan(int[] weights, int days) {
        int minCapacity = shipWithinDays(weights, days);
        java.util.List<java.util.List<Integer>> dailyLoads = new java.util.ArrayList<>();
        
        int currentLoad = 0;
        java.util.List<Integer> currentDay = new java.util.ArrayList<>();
        
        for (int weight : weights) {
            if (currentLoad + weight <= minCapacity) {
                currentLoad += weight;
                currentDay.add(weight);
            } else {
                dailyLoads.add(new java.util.ArrayList<>(currentDay));
                currentDay.clear();
                currentDay.add(weight);
                currentLoad = weight;
            }
        }
        
        if (!currentDay.isEmpty()) {
            dailyLoads.add(currentDay);
        }
        
        return new ShippingPlan(minCapacity, dailyLoads);
    }
    
    public static void main(String[] args) {
        CapacityToShipPackagesWithinDDays solver = new CapacityToShipPackagesWithinDDays();
        
        // Test cases
        int[][][] testCases = {
            {{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, {5}},
            {{3, 2, 2, 4, 1, 4}, {3}},
            {{1, 2, 3, 1, 1}, {4}},
            {{10, 50, 100, 100, 50, 10}, {5}},
            {{1, 2, 3, 4, 5}, {5}},
            {{1, 2, 3, 4, 5}, {1}},
            {{100}, {1}},
            {{1, 1, 1, 1, 1, 1, 1, 1, 1}, {9}},
            {{5, 5, 5, 5, 5}, {3}}
        };
        
        String[] descriptions = {
            "Standard case",
            "Mixed weights",
            "Small weights",
            "Large weights",
            "Equal weights and days",
            "Single day",
            "Single package",
            "Many small packages",
            "Uniform weights"
        };
        
        for (int i = 0; i < testCases.length; i++) {
            System.out.printf("Test Case %d: %s\n", i + 1, descriptions[i]);
            int[] weights = testCases[i][0];
            int days = testCases[i][1][0];
            
            int result1 = solver.shipWithinDays(weights, days);
            int result2 = solver.shipWithinDaysBruteForce(weights, days);
            int result3 = solver.shipWithinDaysOptimized(weights, days);
            
            System.out.printf("  Binary Search: %d\n", result1);
            System.out.printf("  Brute Force: %d\n", result2);
            System.out.printf("  Optimized: %d\n", result3);
            
            // Get shipping plan
            ShippingPlan plan = solver.getShippingPlan(weights, days);
            System.out.printf("  Shipping Plan (capacity %d):\n", plan.minCapacity);
            for (int j = 0; j < plan.dailyLoads.size(); j++) {
                System.out.printf("    Day %d: %s\n", j + 1, plan.dailyLoads.get(j));
            }
            System.out.println();
        }
        
        // Detailed explanation for one case
        System.out.println("=== Detailed Explanation ===");
        int[] detailedWeights = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
        int detailedDays = 5;
        solver.shipWithinDaysDetailed(detailedWeights, detailedDays);
        
        // Performance comparison
        System.out.println("\n=== Performance Comparison ===");
        
        // Large test case
        int[] largeWeights = new int[1000];
        for (int i = 0; i < 1000; i++) {
            largeWeights[i] = (i % 100) + 1;
        }
        
        long startTime = System.nanoTime();
        int largeResult = solver.shipWithinDays(largeWeights, 50);
        long endTime = System.nanoTime();
        
        System.out.printf("Large test (1000 weights, 50 days): %d (took %d ns)\n", 
            largeResult, endTime - startTime);
        
        startTime = System.nanoTime();
        int bruteResult = solver.shipWithinDaysBruteForce(largeWeights, 50);
        endTime = System.nanoTime();
        
        System.out.printf("Brute force result: %d (took %d ns)\n", 
            bruteResult, endTime - startTime);
    }
}
