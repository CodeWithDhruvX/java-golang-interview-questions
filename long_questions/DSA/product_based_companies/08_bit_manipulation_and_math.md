# Bit Manipulation & Math (Product-Based Companies)

Bit manipulation and Math logic test a candidate's grasp of low-level data representation and algorithmic optimization. FAANG companies use these to separate average coders from excellent engineers.

## Question 1: Single Number
**Problem Statement:** Given a non-empty array of integers `nums`, every element appears twice except for one. Find that single one. You must implement a solution with a linear runtime complexity and use only constant extra space.

### Answer:
This is the most famous bit manipulation question. We use the XOR (`^`) operator. 
Properties of XOR:
1. `A ^ 0 = A`
2. `A ^ A = 0`
By XOR-ing all elements in the array, all duplicate numbers will cancel each other out (become 0), leaving only the single number.

**Code Implementation (Java):**
```java
public class SingleNumber {
    public int singleNumber(int[] nums) {
        int result = 0;
        for (int num : nums) {
            result ^= num; // XOR assignment
        }
        return result;
    }
}
```
**Time Complexity:** O(N)
**Space Complexity:** O(1)

---

## Question 2: Counting Bits
**Problem Statement:** Given an integer `n`, return an array `ans` of length `n + 1` such that for each `i` (`0 <= i <= n`), `ans[i]` is the number of `1`'s in the binary representation of `i`. It must be O(N) time.

### Answer:
This combines Dynamic Programming with Bit Manipulation. We can compute the number of 1s in `i` by looking at `i / 2` (which is `i >> 1`). 
- If `i` is even, the binary representation ends in `0`. So `ans[i] = ans[i / 2]`. 
- If `i` is odd, it ends in `1`. So `ans[i] = ans[i / 2] + 1`;
We can simplify this to: `ans[i] = ans[i >> 1] + (i & 1)`.

**Code Implementation (Java):**
```java
public class CountingBits {
    public int[] countBits(int n) {
        int[] ans = new int[n + 1];
        
        for (int i = 1; i <= n; i++) {
            // ans[i >> 1] -> Number of 1 bits in i/2
            // i & 1 -> Tests if the last bit is 1 (if i is odd)
            ans[i] = ans[i >> 1] + (i & 1);
        }
        
        return ans;
    }
}
```
**Time Complexity:** O(N)
**Space Complexity:** O(N) for output array.

---

## Question 3: Pow(x, n) (Fast Exponentiation)
**Problem Statement:** Implement `pow(x, n)`, which calculates `x` raised to the power `n` (i.e., `x^n`). 

### Answer:
A naive `O(N)` loop is too slow for large `n`. Instead, we use **Binary Exponentiation** (Divide and Conquer). For example, `x^10 = (x^5)^2`. And `x^5 = x * (x^2)^2`.
We check if `n` is even or odd, halving `n` continuously.
Be careful: `n` can be negative, and `-Integer.MIN_VALUE` causes overflow, so we must use a `long` for `n`.

**Code Implementation (Java):**
```java
public class Pow {
    public double myPow(double x, int n) {
        long N = n; // Use long to avoid overflow when converting -Integer.MIN_VALUE
        if (N < 0) {
            x = 1 / x;
            N = -N;
        }
        
        double ans = 1;
        double currentProduct = x;
        
        for (long i = N; i > 0; i /= 2) {
            if ((i % 2) == 1) { // If bit is 1 (odd)
                ans = ans * currentProduct;
            }
            // Square the base for the next bit
            currentProduct = currentProduct * currentProduct;
        }
        
        return ans;
    }
}
```
**Time Complexity:** O(log N)
**Space Complexity:** O(1)

---

## Question 4: Missing Number
**Problem Statement:** Given an array `nums` containing `n` distinct numbers in the range `[0, n]`, return the only number in the range that is missing from the array. Can you do it in `O(1)` extra space and `O(N)` runtime?

### Answer:
There are two elegant ways to solve this. 
1. **Math (Gauss Formula):** The sum of the first `N` numbers is `n * (n + 1) / 2`. The missing number is `ExpectedSum - ActualArraySum`.
2. **Bit Manipulation (XOR):** Since `A ^ A = 0`, if we XOR all indexes `0` to `n` and XOR all elements in the array, everything will cancel out except the missing number.

Here is the Math solution (often preferred for readability unless asked explicitly for XOR):

**Code Implementation (Java):**
```java
public class MissingNumber {
    // Math Solution
    public int missingNumber(int[] nums) {
        int n = nums.length;
        // Expected sum from 0 to n
        int expectedSum = n * (n + 1) / 2;
        
        int actualSum = 0;
        for (int num : nums) {
            actualSum += num;
        }
        
        return expectedSum - actualSum;
    }
    
    // Alternative: Bit Manipulation Solution
    public int missingNumberXOR(int[] nums) {
        int res = nums.length; // Start with n
        for (int i = 0; i < nums.length; i++) {
            res ^= i;        // XOR index
            res ^= nums[i];  // XOR value
        }
        return res;
    }
}
```
**Time Complexity:** O(N)
**Space Complexity:** O(1)
