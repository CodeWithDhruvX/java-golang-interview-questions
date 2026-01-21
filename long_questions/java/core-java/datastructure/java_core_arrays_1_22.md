## 1️⃣ Arrays (Java Array & Array Manipulation)

### Question 1: How do you declare, initialize, and copy an array in Java?

**Answer:**
**Declaration & Initialization:**
```java
// Declaration
int[] arr; 

// Initialization (Size must be fixed)
arr = new int[5]; // Default values (0)

// Declaration + Initialization + Assignment
int[] numbers = {1, 2, 3, 4, 5};
```
**Copying:**
Use `System.arraycopy()`, `Arrays.copyOf()`, or `clone()`.

---

### Question 2: Difference between `Arrays.copyOf()` and `System.arraycopy()`.

**Answer:**
- **`System.arraycopy(src, srcPos, dest, destPos, length)`**:
  - Native method (fastest).
  - Copies data into an **existing** destination array.
  - Throws exception if destination is too small.

- **`Arrays.copyOf(original, newLength)`**:
  - Internally uses `System.arraycopy`.
  - Creates and **returns a NEW array**.
  - Can be used to resize an array (pad with defaults if larger).

---

### Question 3: Difference between shallow copy and deep copy of arrays.

**Answer:**
- **Shallow Copy:**
  - Copies references.
  - 1D Primitive arrays: Effectively deep (values copied).
  - Object arrays (or 2D arrays): Only references copied. Modifying an object in the copy affects the original.
  - `clone()`, `System.arraycopy()` perform shallow copies.

- **Deep Copy:**
  - Recursively copies objects.
  - Must be implemented manually (looping and strictly cloning elements) or using Serialization.

---

### Question 4: How do you find the maximum or minimum in an array?

**Answer:**
Iterate via loop.
```java
int[] arr = {5, 2, 9, 1};
int max = arr[0];
for (int i : arr) {
    if (i > max) max = i;
}
```
Or use Streams: `Arrays.stream(arr).max().getAsInt()`.

---

### Question 5: How do you reverse an array in place?

**Answer:**
Use Two-Pointer approach. Swap elements at `start` and `end`, then move pointers inward.
Time: O(N), Space: O(1).
```java
int i = 0, j = arr.length - 1;
while (i < j) {
    int temp = arr[i];
    arr[i] = arr[j];
    arr[j] = temp;
    i++; j--;
}
```

---

### Question 6: How do you rotate an array by k positions?

**Answer:**
**Algorithm (Reversal Method):**
To rotate Right by K:
1.  Reverse whole array.
2.  Reverse first K elements.
3.  Reverse remaining N-K elements.
Time: O(N), Space: O(1).

---

### Question 7: How do you remove duplicates from an array?

**Answer:**
Arrays are fixed size, so you cannot physically "remove" an element (length doesn't change).
1.  **Using Set:** Convert to `HashSet` (O(N) space).
2.  **Sort + Scan (In-place):** Sort (O(N log N)), then overwrite unique elements to the front. Return new "logical" length.

---

### Question 8: How do you find the second largest element in an array?

**Answer:**
One pass O(N).
Maintain two variables: `largest` and `secondLargest`.
```java
if (num > largest) {
    secondLargest = largest;
    largest = num;
} else if (num > secondLargest && num != largest) {
    secondLargest = num;
}
```

---

### Question 9: How do you find a missing number in a sequence of integers?

**Answer:**
If sequence is 1 to N:
1.  Calculate expected sum: `N * (N + 1) / 2`.
2.  Calculate actual sum of array elements.
3.  Difference is the missing number.
(Avoid overflow by subtracting as you go, or use XOR method).

---

### Question 10: How do you find duplicate elements in an array?

**Answer:**
1.  **Brute Force:** O(N^2).
2.  **Sorting:** O(N log N) - check neighbors.
3.  **HashSet:** O(N) time, O(N) space. Add elements; if `set.add()` returns false, it's duplicate.
4.  **Negation (if 1 <= x <= N):** Visit index `abs(x)-1`. If value is negative, it's duplicate. Else flip sign.

---

### Question 11: `Arrays.sort()` vs `Collections.sort()`?

**Answer:**
- **`Arrays.sort(primitive[])`**: Uses **Dual-Pivot Quicksort**. O(N log N). Not stable.
- **`Arrays.sort(Object[])`** / **`Collections.sort()`**: Uses **Timsort** (MergeSort + InsertionSort). Stable. O(N log N).

---

### Question 12: `Arrays.binarySearch()` — how does it work? What are the prerequisites?

**Answer:**
Performs Binary Search O(log N).
**Prerequisite:** The array **MUST be sorted**.
**Returns:**
- Index of element if found.
- `-(insertion_point) - 1` if not found.

---

### Question 13: `Arrays.equals()` vs `==` for arrays?

**Answer:**
- **`==`**: Checks reference equality (Are they the same object in memory?).
- **`Arrays.equals(a, b)`**: Checks content equality (Are lengths same and elements equal?).
- **`Arrays.deepEquals(a, b)`**: Required for multidimensional arrays.

---

### Question 14: `Arrays.fill()` — practical use cases?

**Answer:**
Fills an array with a specific value.
`Arrays.fill(arr, -1);` (Useful for initializing DP tables or memoization arrays).
Time: O(N).

---

### Question 15: `Arrays.asList()` — caveats when using with arrays

**Answer:**
Returns a **fixed-size** list backed by the array.
- You **cannot** add/remove elements (Throws `UnsupportedOperationException`).
- Modifying list elements modifies the original array.
- Pass usage: `new ArrayList<>(Arrays.asList(arr))` to get a modifiable list.

---

### Question 16: `Arrays.stream()` — creating streams from arrays

**Answer:**
`IntStream s = Arrays.stream(intArr);`
Supports slicing: `Arrays.stream(arr, 0, 5)`.
Useful for functional operations: `Arrays.stream(arr).filter(x -> x > 10).sum()`.

---

### Question 17: How to convert a primitive array to a list or stream?

**Answer:**
`int[]` cannot be directly preserved in `Arrays.asList()` (it becomes `List<int[]>`).
**Correct way:**
1.  **Stream:** `Arrays.stream(arr).boxed().collect(Collectors.toList())`.
2.  **Apache Commons:** `ArrayUtils.toObject(arr)`.

---

### Question 18: How do you declare, initialize, and traverse a 2D array?

**Answer:**
**Decl:** `int[][] matrix = new int[3][4];` (3 rows, 4 cols).
**Init:** `int[][] m = { {1,2}, {3,4} };`
**Traverse:** Nested loops.
```java
for (int i = 0; i < matrix.length; i++) {
    for (int j = 0; j < matrix[i].length; j++) {
        print(matrix[i][j]);
    }
}
```

---

### Question 19: How do you find the sum of all elements in a 2D array?

**Answer:**
Iterate all elements and accumulate.
Or `Arrays.stream(matrix).flatMapToInt(Arrays::stream).sum()`.

---

### Question 20: How do you rotate a 2D matrix (like image rotation)?

**Answer:**
**Rotate 90 deg Cokwise:**
1.  **Transpose:** Swap `matrix[i][j]` with `matrix[j][i]`.
2.  **Reverse Rows:** Reverse each row individually.

---

### Question 21: How do you check if a matrix is symmetric?

**Answer:**
A matrix is symmetric if `matrix[i][j] == matrix[j][i]` for all i, j.
Must be a square matrix (rows == cols).

---

### Question 22: How to efficiently search in a row-wise & column-wise sorted matrix?

**Answer:**
Start from **Top-Right** corner (or Bottom-Left).
Search `Target`:
- If `current == target`: Found.
- If `current < target`: Move Down (Row++).
- If `current > target`: Move Left (Col--).
Time: O(N + M).
