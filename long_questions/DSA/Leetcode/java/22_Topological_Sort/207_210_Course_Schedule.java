import java.util.*;

public class CourseSchedule {
    
    // 207. Course Schedule
    // Time: O(V + E), Space: O(V + E)
    public static boolean canFinish(int numCourses, int[][] prerequisites) {
        if (numCourses <= 0) {
            return true;
        }
        
        // Build adjacency list
        List<List<Integer>> graph = new ArrayList<>();
        for (int i = 0; i < numCourses; i++) {
            graph.add(new ArrayList<>());
        }
        
        int[] inDegree = new int[numCourses];
        
        for (int[] prereq : prerequisites) {
            int course = prereq[0];
            int prerequisite = prereq[1];
            graph.get(prerequisite).add(course);
            inDegree[course]++;
        }
        
        // Topological sort using Kahn's algorithm
        Queue<Integer> queue = new LinkedList<>();
        for (int i = 0; i < numCourses; i++) {
            if (inDegree[i] == 0) {
                queue.offer(i);
            }
        }
        
        int processed = 0;
        while (!queue.isEmpty()) {
            int current = queue.poll();
            processed++;
            
            for (int neighbor : graph.get(current)) {
                inDegree[neighbor]--;
                if (inDegree[neighbor] == 0) {
                    queue.offer(neighbor);
                }
            }
        }
        
        return processed == numCourses;
    }

    // 210. Course Schedule II
    // Time: O(V + E), Space: O(V + E)
    public static int[] findOrder(int numCourses, int[][] prerequisites) {
        if (numCourses <= 0) {
            return new int[0];
        }
        
        // Build adjacency list
        List<List<Integer>> graph = new ArrayList<>();
        for (int i = 0; i < numCourses; i++) {
            graph.add(new ArrayList<>());
        }
        
        int[] inDegree = new int[numCourses];
        
        for (int[] prereq : prerequisites) {
            int course = prereq[0];
            int prerequisite = prereq[1];
            graph.get(prerequisite).add(course);
            inDegree[course]++;
        }
        
        // Topological sort using Kahn's algorithm
        Queue<Integer> queue = new LinkedList<>();
        for (int i = 0; i < numCourses; i++) {
            if (inDegree[i] == 0) {
                queue.offer(i);
            }
        }
        
        int[] result = new int[numCourses];
        int index = 0;
        
        while (!queue.isEmpty()) {
            int current = queue.poll();
            result[index++] = current;
            
            for (int neighbor : graph.get(current)) {
                inDegree[neighbor]--;
                if (inDegree[neighbor] == 0) {
                    queue.offer(neighbor);
                }
            }
        }
        
        return index == numCourses ? result : new int[0];
    }

    public static void main(String[] args) {
        // Test cases for canFinish
        Object[][] testCases1 = {
            {2, new int[][]{{1, 0}}},
            {2, new int[][]{{1, 0}, {0, 1}}},
            {4, new int[][]{{1, 0}, {2, 0}, {3, 1}, {3, 2}}},
            {1, new int[][]{}},
            {3, new int[][]{{0, 1}, {0, 2}, {1, 2}}},
            {3, new int[][]{{0, 1}, {1, 2}, {2, 0}}},
            {5, new int[][]{{1, 0}, {2, 0}, {3, 1}, {4, 2}, {4, 3}}},
            {2, new int[][]{}},
            {4, new int[][]{{2, 0}, {1, 0}, {3, 1}, {3, 2}}},
            {3, new int[][]{{1, 0}}},
            {4, new int[][]{{1, 0}, {2, 1}, {3, 2}}},
            {2, new int[][]{{0, 1}}},
            {3, new int[][]{{2, 0}, {2, 1}}},
            {4, new int[][]{{1, 0}, {2, 1}, {3, 2}, {0, 3}}},
            {5, new int[][]{{1, 0}, {2, 1}, {3, 2}, {4, 3}}}
        };
        
        // Test cases for findOrder
        Object[][] testCases2 = {
            {2, new int[][]{{1, 0}}},
            {4, new int[][]{{1, 0}, {2, 0}, {3, 1}, {3, 2}}},
            {1, new int[][]{}},
            {2, new int[][]{{0, 1}, {1, 0}}},
            {4, new int[][]{{1, 0}, {2, 0}, {3, 1}, {3, 2}}},
            {3, new int[][]{{0, 1}, {0, 2}, {1, 2}}},
            {3, new int[][]{{2, 0}, {2, 1}}},
            {4, new int[][]{{1, 0}, {2, 0}, {3, 1}, {3, 2}}},
            {2, new int[][]{}},
            {3, new int[][]{{1, 0}}},
            {4, new int[][]{{1, 0}, {2, 1}, {3, 2}, {0, 3}}},
            {5, new int[][]{{1, 0}, {2, 1}, {3, 2}, {4, 3}}},
            {3, new int[][]{{1, 0}, {1, 2}}},
            {4, new int[][]{{2, 0}, {1, 0}, {3, 1}, {3, 2}}},
            {2, new int[][]{{0, 1}}}
        };
        
        System.out.println("Course Schedule I:");
        for (int i = 0; i < testCases1.length; i++) {
            int numCourses = (int) testCases1[i][0];
            int[][] prerequisites = (int[][]) testCases1[i][1];
            boolean result = canFinish(numCourses, prerequisites);
            System.out.printf("Test Case %d: numCourses=%d, prerequisites=%s -> %b\n", 
                i + 1, numCourses, Arrays.deepToString(prerequisites), result);
        }
        
        System.out.println("\nCourse Schedule II:");
        for (int i = 0; i < testCases2.length; i++) {
            int numCourses = (int) testCases2[i][0];
            int[][] prerequisites = (int[][]) testCases2[i][1];
            int[] result = findOrder(numCourses, prerequisites);
            System.out.printf("Test Case %d: numCourses=%d, prerequisites=%s -> %s\n", 
                i + 1, numCourses, Arrays.deepToString(prerequisites), 
                result.length == 0 ? "[]" : Arrays.toString(result));
        }
    }
}
