import java.util.*;

public class CourseSchedule {
    
    // 207. Course Schedule
    // Time: O(V+E), Space: O(V+E) where V is number of courses, E is prerequisites
    public static boolean canFinish(int numCourses, int[][] prerequisites) {
        // Build adjacency list and in-degree count
        Map<Integer, List<Integer>> adj = new HashMap<>();
        int[] inDegree = new int[numCourses];
        
        for (int[] prereq : prerequisites) {
            int course = prereq[0];
            int prereqCourse = prereq[1];
            
            adj.computeIfAbsent(prereqCourse, k -> new ArrayList<>()).add(course);
            inDegree[course]++;
        }
        
        // Find courses with no prerequisites (in-degree = 0)
        Queue<Integer> queue = new LinkedList<>();
        for (int i = 0; i < numCourses; i++) {
            if (inDegree[i] == 0) {
                queue.offer(i);
            }
        }
        
        // Process courses
        int processed = 0;
        while (!queue.isEmpty()) {
            int course = queue.poll();
            processed++;
            
            // Update in-degree for dependent courses
            for (int neighbor : adj.getOrDefault(course, new ArrayList<>())) {
                inDegree[neighbor]--;
                if (inDegree[neighbor] == 0) {
                    queue.offer(neighbor);
                }
            }
        }
        
        return processed == numCourses;
    }

    // DFS approach to detect cycles
    public static boolean canFinishDFS(int numCourses, int[][] prerequisites) {
        // Build adjacency list
        Map<Integer, List<Integer>> adj = new HashMap<>();
        for (int[] prereq : prerequisites) {
            int course = prereq[0];
            int prereqCourse = prereq[1];
            adj.computeIfAbsent(prereqCourse, k -> new ArrayList<>()).add(course);
        }
        
        // States: 0 = unvisited, 1 = visiting, 2 = visited
        int[] state = new int[numCourses];
        
        for (int i = 0; i < numCourses; i++) {
            if (state[i] == 0 && hasCycleDFS(i, adj, state)) {
                return false; // Cycle detected
            }
        }
        
        return true;
    }
    
    private static boolean hasCycleDFS(int course, Map<Integer, List<Integer>> adj, int[] state) {
        if (state[course] == 1) {
            return true; // Cycle detected
        }
        if (state[course] == 2) {
            return false; // Already processed
        }
        
        state[course] = 1; // Mark as visiting
        
        for (int neighbor : adj.getOrDefault(course, new ArrayList<>())) {
            if (hasCycleDFS(neighbor, adj, state)) {
                return true;
            }
        }
        
        state[course] = 2; // Mark as visited
        return false;
    }

    // Alternative approach using Union-Find to detect cycles
    public static boolean canFinishUnionFind(int numCourses, int[][] prerequisites) {
        int[] parent = new int[numCourses];
        for (int i = 0; i < numCourses; i++) {
            parent[i] = i;
        }
        
        for (int[] prereq : prerequisites) {
            int course = prereq[0];
            int prereqCourse = prereq[1];
            
            if (!union(prereqCourse, course, parent)) {
                return false; // Cycle detected
            }
        }
        
        return true;
    }
    
    private static int find(int x, int[] parent) {
        if (parent[x] != x) {
            parent[x] = find(parent[x], parent);
        }
        return parent[x];
    }
    
    private static boolean union(int x, int y, int[] parent) {
        int rootX = find(x, parent);
        int rootY = find(y, parent);
        
        if (rootX == rootY) {
            return false; // Cycle detected
        }
        
        parent[rootY] = rootX;
        return true;
    }

    public static void main(String[] args) {
        // Test cases
        Object[][] testCases = {
            {2, new int[][]{{1, 0}}, "Simple linear dependency"},
            {2, new int[][]{{1, 0}, {0, 1}}, "Direct cycle"},
            {4, new int[][]{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, "Complex DAG"},
            {1, new int[][]{}, "Single course no prerequisites"},
            {3, new int[][]{{0, 1}, {0, 2}, {1, 2}}, "Multiple dependencies"},
            {3, new int[][]{{0, 1}, {1, 2}, {2, 0}}, "3-course cycle"},
            {5, new int[][]{{1, 0}, {2, 0}, {3, 1}, {4, 2}, {4, 3}}, "Complex structure"},
            {0, new int[][]{}, "No courses"},
            {4, new int[][]{{2, 0}, {1, 0}, {3, 1}, {3, 2}, {1, 3}}, "Complex with cycle"},
            {4, new int[][]{{1, 0}, {2, 1}, {3, 2}}, "Linear chain"}
        };
        
        for (int i = 0; i < testCases.length; i++) {
            int numCourses = (int) testCases[i][0];
            int[][] prereqs = (int[][]) testCases[i][1];
            String description = (String) testCases[i][2];
            
            System.out.printf("Test Case %d: %s\n", i + 1, description);
            System.out.printf("  Courses: %d, Prerequisites: %s\n", numCourses, Arrays.deepToString(prereqs));
            
            boolean result1 = canFinish(numCourses, prereqs);
            boolean result2 = canFinishDFS(numCourses, prereqs);
            boolean result3 = canFinishUnionFind(numCourses, prereqs);
            
            System.out.printf("  BFS (Kahn's): %b\n", result1);
            System.out.printf("  DFS: %b\n", result2);
            System.out.printf("  Union-Find: %b\n\n", result3);
        }
    }
}
