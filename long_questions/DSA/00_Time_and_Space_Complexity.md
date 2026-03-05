# Time and Space Complexity

Understanding **Time and Space Complexity** is crucial for writing efficient code and passing technical interviews. They are typically expressed using **Big O Notation**, which describes how an algorithm's runtime or memory usage scales as the input size ($n$) grows.

Here is a comprehensive guide on how to calculate them and how to identify the best and worst complexities.

---

## 1. How to Know Which is Best and Worst
Before calculating, it's helpful to know the hierarchy of complexities from best (most efficient) to worst (least efficient). 

When we talk about "best" or "worst", we are looking at how the function behaves when $n$ (the input size) becomes infinitely large.

| Big O Notation | Name | Is it Good? | Example Algorithm |
| :--- | :--- | :--- | :--- |
| **$O(1)$** | Constant | **Best (Excellent)** | Accessing an array element by index (`arr[5]`). |
| **$O(\log n)$** | Logarithmic | **Excellent** | Binary search. |
| **$O(n)$** | Linear | **Good** | Looping through an array once. |
| **$O(n \log n)$** | Linearithmic | **Fair / Acceptable** | Efficient sorting (Merge Sort, Quick Sort). |
| **$O(n^2)$** | Quadratic | **Bad (Slow)** | Nested loops (Bubble Sort, checking all pairs). |
| **$O(n^3)$** | Cubic | **Very Bad** | Three nested loops. |
| **$O(2^n)$** | Exponential | **Terrible** | Recursive Fibonacci without memoization. |
| **$O(n!)$** | Factorial | **Worst** | Generating all permutations of a string. |

*(Note: In most interviews, if you write an $O(n^2)$ solution, the interviewer will ask, "Can we do better?" and want you to optimize it to $O(n \log n)$ or $O(n)$.)*

---

## 2. How to Calculate Time Complexity
Time complexity measures **how many operations** an algorithm performs relative to the input size $n$.

**Rules for Calculating Time Complexity:**
1. **Drop Constants:** $O(2n)$ becomes $O(n)$. We only care about how it grows, not exact numbers.
2. **Drop Non-Dominant Terms:** $O(n^2 + n + 5)$ becomes $O(n^2)$ because as $n$ gets huge, $n^2$ completely overshadows $n$ and $5$.

### Examples of Calculating Time Complexity:

**Example A: $O(1)$ - Constant Time**
No matter how big the array is, it takes the same amount of time to return the first element.
```java
int getFirstNumber(int[] numbers) {
    return numbers[0]; // 1 operation
}
```

**Example B: $O(n)$ - Linear Time**
The loop runs $n$ times where $n$ is the length of the array.
```java
void printAll(int[] numbers) {
    for (int i = 0; i < numbers.length; i++) { // runs n times
        System.out.println(numbers[i]);
    }
}
```

**Example C: $O(n^2)$ - Quadratic Time**
For every element in $n$, we loop over $n$ again. $n \times n = n^2$.
```java
void printPairs(int[] numbers) {
    for (int i = 0; i < numbers.length; i++) {       // n times
        for (int j = 0; j < numbers.length; j++) {   // n times
            System.out.println(numbers[i] + ", " + numbers[j]);
        }
    }
}
```

**Example D: $O(\log n)$ - Logarithmic Time**
If the input size is halved at every step (like searching a phone book), it's $\log n$.
```java
// Binary Search cuts the search space in half each iteration
int binarySearch(int[] arr, int target) {
    int left = 0, right = arr.length - 1;
    while (left <= right) { // Search space halves every time
        int mid = left + (right - left) / 2;
        if (arr[mid] == target) return mid;
        if (arr[mid] < target) left = mid + 1;
        else right = mid - 1;
    }
    return -1;
}
```

---

## 3. How to Calculate Space Complexity
Space complexity measures **how much extra memory (RAM)** an algorithm uses relative to the input size $n$. **We do not count the memory taken by the inputs themselves**, only the *auxiliary* (extra) space created inside the function.

### Examples of Calculating Space Complexity:

**Example A: $O(1)$ - Constant Space (Best)**
We only created one integer (`sum`), regardless of how large the input array is.
```java
int calculateSum(int[] numbers) {
    int sum = 0; // O(1) space
    for (int num : numbers) {
        sum += num;
    }
    return sum;
}
```

**Example B: $O(n)$ - Linear Space**
We create a brand new list and put $n$ items into it. The extra memory scales directly with the input size.
```java
List<Integer> doubleNumbers(int[] numbers) {
    List<Integer> doubled = new ArrayList<>(); // O(n) space
    for (int num : numbers) {
        doubled.add(num * 2);
    }
    return doubled;
}
```

**Example C: $O(n)$ - Recursive Stack Space**
When using recursion, every function call is added to the Call Stack in memory. If we recurse $n$ times, space complexity is $O(n)$, even if no variables are created!
```java
void printNumbersRecursively(int n) {
    if (n == 0) return;
    printNumbersRecursively(n - 1); // Depth of recursion tree is n = O(n) space
}
```

---

## Summary Checklist for Interviews:
1. **Are there single loops?** Usually $O(n)$ time.
2. **Are there nested loops?** Usually $O(n^2)$ time.
3. **Does the problem size halve every step?** $O(\log n)$ time.
4. **Did you create a new array, list, or hash map as big as the input?** $O(n)$ space.
5. **Did you sort the data?** Sorting is usually $O(n \log n)$ time and $O(1)$ to $O(n)$ space depending on the language's sorting algorithm.
6. **Did you use recursion?** Remember to count the depth of the recursion tree as Space Complexity!
