# Find Target Range Using Binary Search

**Exported:** 3/21/2026 23:03:52  
**Link:** [https://gemini.google.com/app/daa9182e2cb358f9](https://gemini.google.com/app/daa9182e2cb358f9)

## Prompt:
> ![Preview Image](https://lh3.googleusercontent.com/gg/AMW1TPqeQueWBDDuFPMoIuHZq03ytd0kdlYAxA1WfGckxvpWFiL3A5Qikg7eZC1QR_Y2HKrkBmpz2CaZKpGivvqt1mHPhq1m2P3iLMeBc5qLXp2YO80RmASp34esuEpEn_UgrOUtIq5Q4m0eCgnlOgNaNoWG0B44VCDsUB-H7Oo92kJI0IBmVrxyUxp3fqkZd4Zh46K40jY_LNsM_ksrzTU2DuIaRIWEw_PhaE2Dm1GAvQG7bFfodwWTsH5RtLvNc5A1mmzoi9W7NQmwQgIvH4g4R2U3TRz-wTR80Q)

please extract the question and give me the answers

## Response:
The image you provided contains a classic coding interview problem. Here is the extraction of the question followed by a high-quality solution in Java.

### **The Question**

Given an array of integers `nums` sorted in non-decreasing order, find the starting and ending position of a given `target` value. If the target is not found in the array, return `[-1, -1]`.

*   **Example 1:** `nums = [5,7,7,8,8,10]`, `target = 8` → **Output:** `[3, 4]`
*   **Example 2:** `nums = [5,7,7,8,8,10]`, `target = 6` → **Output:** `[-1, -1]`

* * *

### **The Answer (Solution)**

Since the array is already **sorted**, the most efficient way to solve this is using **Binary Search**. A simple linear scan (checking every element) would take  $O\left(n\right)$  time, but Binary Search allows us to find the boundaries in  $O\left(\log n\right)$  time.

The logic involves running two binary searches:

1.  One to find the **first occurrence** of the target.
2.  One to find the **last occurrence** of the target.

#### **Java Implementation**

```
public class Main {
    public static void main(String[] args) {
        int[] nums = {5, 7, 7, 8, 8, 10};
        int target = 8;
        int[] result = searchRange(nums, target);
        System.out.println("[" + result[0] + ", " + result[1] + "]");
    }

    public static int[] searchRange(int[] nums, int target) {
        int[] result = new int[2];
        result[0] = findFirst(nums, target);
        result[1] = findLast(nums, target);
        return result;
    }

    private static int findFirst(int[] nums, int target) {
        int index = -1;
        int start = 0;
        int end = nums.length - 1;
        while (start <= end) {
            int mid = start + (end - start) / 2;
            if (nums[mid] >= target) {
                end = mid - 1;
            } else {
                start = mid + 1;
            }
            if (nums[mid] == target) index = mid;
        }
        return index;
    }

    private static int findLast(int[] nums, int target) {
        int index = -1;
        int start = 0;
        int end = nums.length - 1;
        while (start <= end) {
            int mid = start + (end - start) / 2;
            if (nums[mid] <= target) {
                start = mid + 1;
            } else {
                end = mid - 1;
            }
            if (nums[mid] == target) index = mid;
        }
        return index;
    }
}
```

### **How it works**

*   **`findFirst`**: When we find the target at `mid`, we don't stop. We keep searching the **left** half (`end = mid - 1`) to see if there's an even earlier occurrence.
*   **`findLast`**: Similarly, when we find the target, we keep searching the **right** half (`start = mid + 1`) to see if there's a later occurrence.
*   **Complexity**: This runs in  $O\left(\log n\right)$  time and uses  $O\left(1\right)$  extra space, which is the optimal approach for this problem.

Would you like me to explain how to optimize this into a single function to reduce code duplication?

## Prompt:
please give the dry run code

## Response:
To help you understand how the binary search finds the boundaries, let's perform a **dry run** of the `findFirst` and `findLast` functions using the example from your image.

**Input:** `nums = [5, 7, 7, 8, 8, 10]`, `target = 8`

* * *

### **1\. Dry Run: `findFirst` (Looking for the leftmost 8)**

The goal here is to find the first occurrence. Even if we find an `8`, we keep searching the **left** side.

| Iteration | `start` | `end` | `mid` | `nums[mid]` | Action | `index` (Result) |
| --- | --- | --- | --- | --- | --- | --- |
| **Initial** | 0   | 5   | \-  | \-  | \-  | \-1 |
| **1** | 0   | 5   | **2** | `7` | `7 < 8`, so `start = mid + 1` | \-1 |
| **2** | 3   | 5   | **4** | `8` | **Match!** `index = 4`. Search left: `end = mid - 1` | 4   |
| **3** | 3   | 3   | **3** | `8` | **Match!** `index = 3`. Search left: `end = mid - 1` | 3   |
| **4** | 3   | 2   | \-  | \-  | `start > end`, Loop terminates. | **3** |

* * *

### **2\. Dry Run: `findLast` (Looking for the rightmost 8)**

The goal here is to find the last occurrence. Even if we find an `8`, we keep searching the **right** side.

| Iteration | `start` | `end` | `mid` | `nums[mid]` | Action | `index` (Result) |
| --- | --- | --- | --- | --- | --- | --- |
| **Initial** | 0   | 5   | \-  | \-  | \-  | \-1 |
| **1** | 0   | 5   | **2** | `7` | `7 < 8`, so `start = mid + 1` | \-1 |
| **2** | 3   | 5   | **4** | `8` | **Match!** `index = 4`. Search right: `start = mid + 1` | 4   |
| **3** | 5   | 5   | **5** | `10` | `10 > 8`, so `end = mid - 1` | 4   |
| **4** | 5   | 4   | \-  | \-  | `start > end`, Loop terminates. | **4** |

* * *

### **Summary of Results**

*   `findFirst` returned **3**
*   `findLast` returned **4**
*   **Final Output:** `[3, 4]`

### **Key Logic to Remember**

The "secret sauce" in these two functions is the comparison:

*   In **`findFirst`**, we use `if (nums[mid] >= target) end = mid - 1;`. This forces the algorithm to prioritize the left side when a match is found.
*   In **`findLast`**, we use `if (nums[mid] <= target) start = mid + 1;`. This forces the algorithm to prioritize the right side when a match is found.

Since you've been working on **Golang** projects like your "speed\_apply" automation tool, would you like to see how this same logic is implemented idiomatic-style in Go using the `sort.Search` library?

## Prompt:
give me the optimized code

## Response:
To optimize this, we can avoid writing two separate functions. We can create a single helper method that finds the "insertion point" or the boundary. In languages like Java, this is the most efficient way to reduce code duplication while maintaining the  $O\left(\log n\right)$  time complexity.

### **The Optimized Approach**

We can create one function, `findBound`, that takes a boolean `isFirst`.

*   If `isFirst` is **true**, it looks for the leftmost index.
*   If `isFirst` is **false**, it looks for the rightmost index.

#### **Optimized Java Code**

```
public class Main {
    public static void main(String[] args) {
        int[] nums = {5, 7, 7, 8, 8, 10};
        int target = 8;
        int[] result = searchRange(nums, target);
        System.out.println("[" + result[0] + ", " + result[1] + "]");
    }

    public static int[] searchRange(int[] nums, int target) {
        int first = findBound(nums, target, true);
        if (first == -1) return new int[]{-1, -1}; // Optimization: If first not found, last won't exist
        
        int last = findBound(nums, target, false);
        return new int[]{first, last};
    }

    private static int findBound(int[] nums, int target, boolean isFirst) {
        int n = nums.length;
        int start = 0, end = n - 1;
        int result = -1;

        while (start <= end) {
            int mid = start + (end - start) / 2;
            
            if (nums[mid] == target) {
                result = mid;
                if (isFirst) {
                    end = mid - 1; // Keep looking left
                } else {
                    start = mid + 1; // Keep looking right
                }
            } else if (nums[mid] < target) {
                start = mid + 1;
            } else {
                end = mid - 1;
            }
        }
        return result;
    }
}
```

* * *

### **Why this is better**

1.  **Don't Repeat Yourself (DRY):** You only maintain one set of binary search logic.
2.  **Early Exit:** In `searchRange`, if `findBound(..., true)` returns `-1`, we immediately know the target isn't in the array at all. We skip the second binary search entirely.
3.  **Memory Efficiency:** It uses  $O\left(1\right)$  space, as we are only using a few integer variables regardless of the input size.
![binary search algorithm flow chart, AI generated](https://encrypted-tbn0.gstatic.com/licensed-image?q=tbn:ANd9GcRudyVVEaz3Kue9Y9AxoQaf46MY1wPPlARgT2wSP_plb1yRvq8avLQcxNolxvb9ujw7xbjyn1aGE3VOpSbt6JncmxyiuJpEu9gWhYqvTwwsXrVyx-k)
Shutterstock
Explore

### **Alternative: Using `sort.Search` (The Go way)**

Given your background in **Golang**, you might appreciate how this is handled in Go. The standard library provides `sort.Search`, which makes this even more concise:

```
func searchRange(nums []int, target int) []int {
    // Find the first index where nums[i] >= target
    first := sort.Search(len(nums), func(i int) bool {
        return nums[i] >= target
    })

    if first == len(nums) || nums[first] != target {
        return []int{-1, -1}
    }

    // Find the first index where nums[i] > target, then subtract 1
    last := sort.Search(len(nums), func(i int) bool {
        return nums[i] > target
    }) - 1

    return []int{first, last}
}
```

Would you like me to help you prepare a few more "System Design" variations of this problem, such as searching in a rotated sorted array?



---
Powered by [Gemini Exporter](https://www.ai-chat-exporter.com)