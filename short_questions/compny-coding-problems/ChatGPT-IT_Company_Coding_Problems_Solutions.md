# IT Company Coding Problems - Pseudocode Solutions

This document contains pseudocode solutions for the top coding problems asked by Indian service-based IT companies (TCS, Infosys, Wipro, Accenture, Cognizant, etc.).

## 1. STRING PROBLEMS

### 1. Reverse a string
```text
FUNCTION reverseString(str):
    Initialize empty string reversedStr
    FOR i FROM length of str - 1 DOWN TO 0:
        Append str[i] to reversedStr
    RETURN reversedStr

// TEST CASES
// Input: "hello" -> Output: "olleh"
// Input: "Java"  -> Output: "avaJ"
// Input: ""      -> Output: ""
```

### 2. Reverse words in a sentence
```text
FUNCTION reverseWords(sentence):
    Split sentence by space into array words
    Initialize empty string result
    FOR i FROM length of words - 1 DOWN TO 0:
        Append words[i] + " " to result
    RETURN result trimmed

// TEST CASES
// Input: "Hello World" -> Output: "World Hello"
// Input: "I love Java" -> Output: "Java love I"
```

### 3. Check if a string is palindrome
```text
FUNCTION isPalindrome(str):
    Initialize start = 0, end = length of str - 1
    WHILE start < end:
        IF str[start] NOT EQUAL str[end]:
            RETURN False
        Increment start, Decrement end
    RETURN True

// TEST CASES
// Input: "madam" -> Output: True
// Input: "hello" -> Output: False
// Input: "racecar" -> Output: True
```

### 4. Check if two strings are anagrams
```text
FUNCTION areAnagrams(str1, str2):
    IF length of str1 NOT EQUAL length of str2:
        RETURN False
    Sort str1
    Sort str2
    IF str1 EQUALS str2:
        RETURN True
    RETURN False

// TEST CASES
// Input: "listen", "silent" -> Output: True
// Input: "hello", "world"   -> Output: False
```

### 5. Count vowels and consonants
```text
FUNCTION countVowelsConsonants(str):
    Initialize vowels = 0, consonants = 0
    Convert str to lowercase
    FOR EACH char IN str:
        IF char is a letter:
            IF char IN ['a', 'e', 'i', 'o', 'u']:
                Increment vowels
            ELSE:
                Increment consonants
    PRINT vowels, consonants

// TEST CASES
// Input: "Hello" -> Output: Vowels: 2, Consonants: 3
// Input: "Apple" -> Output: Vowels: 2, Consonants: 3
```

### 6. Count frequency of characters
```text
FUNCTION charFrequency(str):
    Initialize empty Map freqMap
    FOR EACH char IN str:
        IF char in freqMap:
            freqMap[char] += 1
        ELSE:
            freqMap[char] = 1
    PRINT freqMap

// TEST CASES
// Input: "banana" -> Output: {b:1, a:3, n:2}
// Input: "hello"  -> Output: {h:1, e:1, l:2, o:1}
```

### 7. Find first non-repeating character
```text
FUNCTION firstNonRepeating(str):
    Initialize empty Map freqMap
    FOR EACH char IN str:
        freqMap[char] = (freqMap[char] OR 0) + 1
    FOR EACH char IN str:
        IF freqMap[char] EQUAL 1:
            RETURN char
    RETURN NULL

// TEST CASES
// Input: "swiss" -> Output: "w"
// Input: "aabbc" -> Output: "c"
// Input: "aabb"  -> Output: NULL
```

### 8. Remove duplicate characters
```text
FUNCTION removeDuplicates(str):
    Initialize empty Set seen, empty string result
    FOR EACH char IN str:
        IF char NOT IN seen:
            Add char to seen
            Append char to result
    RETURN result

// TEST CASES
// Input: "banana" -> Output: "ban"
// Input: "hello"  -> Output: "helo"
```

### 9. Replace spaces with special character
```text
FUNCTION replaceSpaces(str, specialChar):
    Initialize empty string result
    FOR EACH char IN str:
        IF char EQUAL ' ':
            Append specialChar to result
        ELSE:
            Append char to result
    RETURN result

// TEST CASES
// Input: "Hello World", "-" -> Output: "Hello-World"
// Input: "Java Code", "%20" -> Output: "Java%20Code"
```

### 10. Convert lowercase to uppercase (without built-in)
```text
FUNCTION toUpperCase(str):
    Initialize empty string result
    FOR EACH char IN str:
        IF ascii(char) >= 97 AND ascii(char) <= 122:
            Append char from ascii(ascii(char) - 32)
        ELSE:
            Append char
    RETURN result

// TEST CASES
// Input: "java" -> Output: "JAVA"
// Input: "Hello" -> Output: "HELLO"
```

### 11. Find longest word in a string
```text
FUNCTION longestWord(sentence):
    Split sentence by space into words
    Initialize maxLen = 0, longest = ""
    FOR EACH word IN words:
        IF length of word > maxLen:
            maxLen = length of word
            longest = word
    RETURN longest

// TEST CASES
// Input: "I love programming" -> Output: "programming"
// Input: "Hello world" -> Output: "Hello" (or "world")
```

### 12. Count number of words
```text
FUNCTION countWords(sentence):
    IF sentence is empty: RETURN 0
    Split sentence by space into words
    RETURN length of words

// TEST CASES
// Input: "Hello World" -> Output: 2
// Input: " One Two Three " -> Output: 3
```

### 13. Check substring present or not
```text
FUNCTION isSubstring(mainStr, subStr):
    FOR i FROM 0 TO length of mainStr - length of subStr:
        IF substring of mainStr starting at i with length of subStr EQUALS subStr:
            RETURN True
    RETURN False

// TEST CASES
// Input: "Hello World", "World" -> Output: True
// Input: "Hello", "Java"        -> Output: False
```

### 14. Remove vowels from string
```text
FUNCTION removeVowels(str):
    Initialize empty string result
    FOR EACH char IN str:
        IF char NOT IN ['a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U']:
            Append char to result
    RETURN result

// TEST CASES
// Input: "Hello World" -> Output: "Hll Wrld"
// Input: "Apple"       -> Output: "ppl"
```

### 15. Sort characters in a string
```text
FUNCTION sortString(str):
    Convert str to character array arr
    Sort arr using Bubble Sort or Merge Sort
    Return string from arr

// TEST CASES
// Input: "edcba" -> Output: "abcde"
// Input: "python" -> Output: "hnopty"
```

### 16. Find duplicate characters
```text
FUNCTION findDuplicates(str):
    Initialize Map freqMap
    FOR EACH char IN str:
        freqMap[char] = (freqMap[char] OR 0) + 1
    FOR EACH key, value IN freqMap:
        IF value > 1:
            PRINT key

// TEST CASES
// Input: "banana" -> Output: "a", "n"
// Input: "hello"  -> Output: "l"
```

### 17. Reverse string using recursion
```text
FUNCTION reverseRecursive(str):
    IF str is empty:
        RETURN ""
    RETURN reverseRecursive(substring of str from 1) + str[0]

// TEST CASES
// Input: "recursion" -> Output: "noisrucer"
// Input: "abc"       -> Output: "cba"
```

### 18. Print string in zig-zag format
```text
FUNCTION printZigZag(str, k):
   // Print string characters in a zig-zag pattern of k rows
   Create k empty strings (rows)
   Initialize row = 0, down = True
   FOR EACH char IN str:
       Append char to rows[row]
       IF row == 0 -> down = True
       IF row == k-1 -> down = False
       IF down -> row++ ELSE row--
   Print all rows concatenated

// TEST CASES
// Input: "PAYPALISHIRING", 3 -> Output: "PAHNAPLSIIGYIR"
// Input: "ABC", 1            -> Output: "ABC"
```

### 19. Check string rotation
```text
FUNCTION isRotation(str1, str2):
    IF length of str1 NOT EQUAL length of str2:
        RETURN False
    temp = str1 + str1
    IF str2 is substring of temp:
        RETURN True
    RETURN False

// TEST CASES
// Input: "ABCD", "CDAB" -> Output: True
// Input: "ABCD", "ACBD" -> Output: False
```

### 20. Compare two strings without using equals()
```text
FUNCTION compareStrings(str1, str2):
    IF length of str1 NOT EQUAL length of str2:
        RETURN False
    FOR i FROM 0 TO length of str1 - 1:
        IF str1[i] NOT EQUAL str2[i]:
            RETURN False
    RETURN True

// TEST CASES
// Input: "abc", "abc" -> Output: True
// Input: "abc", "abd" -> Output: False
```

## 2. ARRAY PROBLEMS

### 21. Find largest element in array
```text
FUNCTION findMax(arr):
    max = arr[0]
    FOR i FROM 1 TO length of arr - 1:
        IF arr[i] > max:
            max = arr[i]
    RETURN max

// TEST CASES
// Input: [1, 5, 3, 9, 2] -> Output: 9
// Input: [-1, -5, -2]    -> Output: -1
```

### 22. Find smallest element in array
```text
FUNCTION findMin(arr):
    min = arr[0]
    FOR i FROM 1 TO length of arr - 1:
        IF arr[i] < min:
            min = arr[i]
    RETURN min

// TEST CASES
// Input: [1, 5, 3, 9, 2] -> Output: 1
// Input: [-1, -5, -2]    -> Output: -5
```

### 23. Find second largest element
```text
FUNCTION secondLargest(arr):
    first = -Infinity, second = -Infinity
    FOR EACH num IN arr:
        IF num > first:
            second = first
            first = num
        ELSE IF num > second AND num NOT EQUAL first:
            second = num
    RETURN second

// TEST CASES
// Input: [12, 35, 1, 10, 34, 1] -> Output: 34
// Input: [10, 5, 10]            -> Output: 5
```

### 24. Reverse an array
```text
FUNCTION reverseArray(arr):
    start = 0, end = length of arr - 1
    WHILE start < end:
        Swap arr[start] with arr[end]
        Increment start, Decrement end
    RETURN arr

// TEST CASES
// Input: [1, 2, 3, 4] -> Output: [4, 3, 2, 1]
// Input: [1]          -> Output: [1]
```

### 25. Rotate array left by K positions
```text
FUNCTION rotateLeft(arr, k):
    k = k % length of arr
    Reverse(arr, 0, k-1)
    Reverse(arr, k, length-1)
    Reverse(arr, 0, length-1)
    RETURN arr

// TEST CASES
// Input: [1, 2, 3, 4, 5], k=2 -> Output: [3, 4, 5, 1, 2]
// Input: [1, 2], k=3          -> Output: [2, 1]
```

### 26. Rotate array right by K positions
```text
FUNCTION rotateRight(arr, k):
    k = k % length of arr
    Reverse(arr, 0, length-k-1)
    Reverse(arr, length-k, length-1)
    Reverse(arr, 0, length-1)
    RETURN arr

// TEST CASES
// Input: [1, 2, 3, 4, 5], k=2 -> Output: [4, 5, 1, 2, 3]
// Input: [1, 2], k=3          -> Output: [2, 1]
```

### 27. Remove duplicates from array
```text
FUNCTION removeArrayDuplicates(arr):
    Initialize Set seen
    Initialize List result
    FOR EACH num IN arr:
        IF num NOT IN seen:
            Add num to seen
            Add num to result
    RETURN result

// TEST CASES
// Input: [1, 2, 2, 3, 4, 4] -> Output: [1, 2, 3, 4]
// Input: [1, 1, 1]          -> Output: [1]
```

### 28. Find frequency of elements
```text
FUNCTION arrayFrequency(arr):
    Initialize Map freqMap
    FOR EACH num IN arr:
        freqMap[num] = (freqMap[num] OR 0) + 1
    PRINT freqMap

// TEST CASES
// Input: [1, 2, 2, 3] -> Output: {1:1, 2:2, 3:1}
// Input: [5, 5, 5]    -> Output: {5:3}
```

### 29. Find missing number in array (1 to N)
```text
FUNCTION findMissing(arr, N):
    expectedSum = N * (N + 1) / 2
    actualSum = Sum of elements in arr
    RETURN expectedSum - actualSum

// TEST CASES
// Input: [1, 2, 4, 5], N=5 -> Output: 3
// Input: [1, 3], N=3       -> Output: 2
```

### 30. Find duplicate number
```text
FUNCTION findDuplicate(arr):
    Initialize Set seen
    FOR EACH num IN arr:
        IF num IN seen:
            RETURN num
        Add num to seen
    RETURN -1

// TEST CASES
// Input: [1, 3, 4, 2, 2] -> Output: 2
// Input: [3, 1, 3, 4, 2] -> Output: 3
```

### 31. Sort array without built-in methods (Bubble Sort)
```text
FUNCTION bubbleSort(arr):
    n = length of arr
    FOR i FROM 0 TO n-1:
        FOR j FROM 0 TO n-i-1:
            IF arr[j] > arr[j+1]:
                Swap arr[j] with arr[j+1]
    RETURN arr

// TEST CASES
// Input: [5, 1, 4, 2, 8] -> Output: [1, 2, 4, 5, 8]
// Input: [3, 2, 1]       -> Output: [1, 2, 3]
```

### 32. Merge two arrays
```text
FUNCTION mergeArrays(arr1, arr2):
    Initialize result array of size len(arr1) + len(arr2)
    Copy arr1 to result
    Copy arr2 to result after arr1
    RETURN result

// TEST CASES
// Input: [1, 2], [3, 4] -> Output: [1, 2, 3, 4]
// Input: [], [1]        -> Output: [1]
```

### 33. Find common elements in two arrays
```text
FUNCTION findCommon(arr1, arr2):
    Initialize Set set1 with elements of arr1
    Initialize List common
    FOR EACH num IN arr2:
        IF num IN set1:
            Add num to common
            Remove num from set1 (to avoid duplicates)
    RETURN common

// TEST CASES
// Input: [1, 2, 3], [2, 3, 4] -> Output: [2, 3]
// Input: [1, 2], [3, 4]       -> Output: []
```

### 34. Move all zeros to end
```text
FUNCTION moveZeros(arr):
    count = 0
    FOR i FROM 0 TO length of arr - 1:
        IF arr[i] NOT EQUAL 0:
            arr[count] = arr[i]
            Increment count
    WHILE count < length of arr:
        arr[count] = 0
        Increment count
    RETURN arr

// TEST CASES
// Input: [0, 1, 0, 3, 12] -> Output: [1, 3, 12, 0, 0]
// Input: [0, 0, 1]        -> Output: [1, 0, 0]
```

### 35. Find sum of array elements
```text
FUNCTION sumArray(arr):
    total = 0
    FOR EACH num IN arr:
        total += num
    RETURN total

// TEST CASES
// Input: [1, 2, 3, 4] -> Output: 10
// Input: [-1, 1]      -> Output: 0
```

### 36. Find pair with given sum
```text
FUNCTION findPairWithSum(arr, target):
    Initialize Set seen
    FOR EACH num IN arr:
        complement = target - num
        IF complement IN seen:
            RETURN {complement, num}
        Add num to seen
    RETURN NULL

// TEST CASES
// Input: [2, 7, 11, 15], 9 -> Output: {2, 7}
// Input: [3, 2, 4], 6      -> Output: {2, 4}
```

### 37. Find max & min in single loop
```text
FUNCTION findMaxMin(arr):
    max = arr[0], min = arr[0]
    FOR i FROM 1 TO length of arr - 1:
        IF arr[i] > max: max = arr[i]
        IF arr[i] < min: min = arr[i]
    RETURN {max, min}

// TEST CASES
// Input: [3, 5, 1, 2, 4, 8] -> Output: {8, 1}
// Input: [1]                -> Output: {1, 1}
```

### 38. Print array in reverse order
```text
FUNCTION printReverse(arr):
    FOR i FROM length of arr - 1 DOWN TO 0:
        PRINT arr[i]

// TEST CASES
// Input: [1, 2, 3, 4] -> Output: 4 3 2 1
// Input: [100]        -> Output: 100
```

### 39. Check array is sorted or not
```text
FUNCTION isSorted(arr):
    FOR i FROM 0 TO length of arr - 2:
        IF arr[i] > arr[i+1]:
            RETURN False
    RETURN True

// TEST CASES
// Input: [1, 2, 3, 4, 5] -> Output: True
// Input: [1, 3, 2]       -> Output: False
```

### 40. Count even & odd numbers
```text
FUNCTION countEvenOdd(arr):
    even = 0, odd = 0
    FOR EACH num IN arr:
        IF num % 2 EQUAL 0:
            Increment even
        ELSE:
            Increment odd
    PRINT even, odd

// TEST CASES
// Input: [1, 2, 3, 4, 5] -> Output: Even: 2, Odd: 3
// Input: [2, 4, 6]       -> Output: Even: 3, Odd: 0
```

## 3. PATTERN PRINTING

### 41. Right triangle star pattern
```text
FUNCTION rightTriangle(N):
    FOR i FROM 1 TO N:
        PRINT "*" repeated i times

// TEST CASES
// Input: 3
// Output:
// *
// **
// ***
```

### 42. Left triangle star pattern
```text
FUNCTION leftTriangle(N):
    FOR i FROM 1 TO N:
        PRINT " " repeated (N-i) times + "*" repeated i times

// TEST CASES
// Input: 3
// Output:
//   *
//  **
// ***
```

### 43. Pyramid star pattern
```text
FUNCTION pyramid(N):
    FOR i FROM 1 TO N:
        PRINT " " repeated (N-i) times + "* " repeated i times

// TEST CASES
// Input: 3
// Output:
//   * 
//  * * 
// * * * 
```

### 44. Inverted pyramid
```text
FUNCTION invertedPyramid(N):
    FOR i FROM N DOWN TO 1:
        PRINT " " repeated (N-i) times + "* " repeated i times

// TEST CASES
// Input: 3
// Output:
// * * * 
//  * * 
//   * 
```

### 45. Diamond pattern
```text
FUNCTION diamond(N):
    pyramid(N)
    invertedPyramid(N) // Handle middle row overlap if needed

// TEST CASES
// Input: 3 (Assuming simplified diamond logic)
// Output:
//   *
//  ***
// *****
//  ***
//   *
```

### 46. Number pyramid
```text
FUNCTION numberPyramid(N):
    FOR i FROM 1 TO N:
        PRINT " " repeated (N-i) times
        FOR j FROM 1 TO i:
            PRINT j + " "
        PRINT NEWLINE

// TEST CASES
// Input: 3
// Output:
//   1 
//  1 2 
// 1 2 3 
```

### 47. Floyd’s triangle
```text
FUNCTION floydsTriangle(N):
    num = 1
    FOR i FROM 1 TO N:
        FOR j FROM 1 TO i:
            PRINT num + " "
            Increment num
        PRINT NEWLINE

// TEST CASES
// Input: 3
// Output:
// 1 
// 2 3 
// 4 5 6 
```

### 48. Hollow square pattern
```text
FUNCTION hollowSquare(N):
    FOR i FROM 1 TO N:
        IF i EQUAL 1 OR i EQUAL N:
            PRINT "*" repeated N times
        ELSE:
            PRINT "*" + " " repeated (N-2) times + "*"

// TEST CASES
// Input: 3
// Output:
// ***
// * *
// ***
```

### 49. Hollow pyramid
```text
FUNCTION hollowPyramid(N):
   // Complex Logic: Print stars at boundaries, spaces inside.
   FOR i FROM 1 TO N:
      FOR j FROM 1 TO N-i: PRINT " "
      FOR k FROM 1 TO (2*i - 1):
         IF k==1 OR k==(2*i - 1) OR i==N: PRINT "*"
         ELSE: PRINT " "
      PRINT NEWLINE

// TEST CASES
// Input: 3
// Output:
//   *
//  * *
// *****
```

### 50. Zig-zag star pattern
```text
FUNCTION zigZag(N):
    // Generally requires specific row/column logic
    // Example for 3 rows
    FOR i FROM 1 TO 3:
        FOR j FROM 1 TO N:
            IF ((i+j) % 4 == 0) OR (i == 2 AND j % 4 == 0):
                PRINT "*"
            ELSE:
                PRINT " "
        PRINT NEWLINE

// TEST CASES
// Input: 9
// Output:
//   *   *
//  * * * *
// *   *   *
```

### 51. Butterfly pattern
```text
FUNCTION butterfly(N):
    FOR i FROM 1 TO N:
        PRINT "*" * i + spaces + "*" * i
    FOR i FROM N DOWN TO 1:
        PRINT "*" * i + spaces + "*" * i

// TEST CASES
// Input: 2
// Output:
// *  *
// ****
// ****
// *  *
```

### 52. Pascal triangle
```text
FUNCTION pascalTriangle(N):
    FOR i FROM 0 TO N-1:
        val = 1
        FOR j FROM 0 TO i:
            PRINT val + " "
            val = val * (i - j) / (j + 1)
        PRINT NEWLINE

// TEST CASES
// Input: 3
// Output:
// 1 
// 1 1 
// 1 2 1 
```

### 53. Number increasing pattern
```text
FUNCTION numberIncreasing(N):
    FOR i FROM 1 TO N:
        FOR j FROM 1 TO i:
            PRINT j
        PRINT NEWLINE

// TEST CASES
// Input: 3
// Output:
// 1
// 12
// 123
```

### 54. Character pattern
```text
FUNCTION charPattern(N):
    FOR i FROM 1 TO N:
        char = 'A'
        FOR j FROM 1 TO i:
            PRINT char
            char++
        PRINT NEWLINE

// TEST CASES
// Input: 3
// Output:
// A
// BC
// DEF
```

### 55. X pattern using stars
```text
FUNCTION xPattern(N):
    FOR i FROM 0 TO N-1:
        FOR j FROM 0 TO N-1:
            IF i EQUAL j OR i+j EQUAL N-1:
                PRINT "*"
            ELSE:
                PRINT " "
        PRINT NEWLINE

// TEST CASES
// Input: 3
// Output:
// * *
//  * 
// * *
```

## 4. MATHEMATICAL / NUMBER LOGIC

### 56. Check prime number
```text
FUNCTION isPrime(N):
    IF N <= 1: RETURN False
    FOR i FROM 2 TO sqrt(N):
        IF N % i EQUAL 0:
            RETURN False
    RETURN True

// TEST CASES
// Input: 7 -> Output: True
// Input: 4 -> Output: False
```

### 57. Print prime numbers in range
```text
FUNCTION printPrimes(start, end):
    FOR i FROM start TO end:
        IF isPrime(i):
            PRINT i

// TEST CASES
// Input: 1, 10 -> Output: 2, 3, 5, 7
// Input: 10, 20 -> Output: 11, 13, 17, 19
```

### 58. Fibonacci series
```text
FUNCTION fibonacci(N):
    a = 0, b = 1
    PRINT a, b
    FOR i FROM 2 TO N-1:
        c = a + b
        PRINT c
        a = b, b = c

// TEST CASES
// Input: 5 -> Output: 0, 1, 1, 2, 3
// Input: 1 -> Output: 0
```

### 59. Factorial of number
```text
FUNCTION factorial(N):
    fact = 1
    FOR i FROM 1 TO N:
        fact = fact * i
    RETURN fact

// TEST CASES
// Input: 5 -> Output: 120
// Input: 0 -> Output: 1
```

### 60. Armstrong number
```text
FUNCTION isArmstrong(N):
    original = N, sum = 0, digits = length of N
    WHILE N > 0:
        digit = N % 10
        sum += power(digit, digits)
        N = N / 10
    RETURN sum EQUAL original

// TEST CASES
// Input: 153 -> Output: True (1^3 + 5^3 + 3^3 = 153)
// Input: 123 -> Output: False
```

### 61. Perfect number
```text
FUNCTION isPerfect(N):
    sum = 0
    FOR i FROM 1 TO N/2:
        IF N % i EQUAL 0:
            sum += i
    RETURN sum EQUAL N

// TEST CASES
// Input: 6  -> Output: True (1+2+3 = 6)
// Input: 10 -> Output: False
```

### 62. Palindrome number
```text
FUNCTION isPalindromeNum(N):
    original = N, reverse = 0
    WHILE N > 0:
        reverse = reverse * 10 + (N % 10)
        N = N / 10
    RETURN original EQUAL reverse

// TEST CASES
// Input: 121 -> Output: True
// Input: 123 -> Output: False
```

### 63. Sum of digits
```text
FUNCTION sumDigits(N):
    sum = 0
    WHILE N > 0:
        sum += N % 10
        N = N / 10
    RETURN sum

// TEST CASES
// Input: 123 -> Output: 6
// Input: 901 -> Output: 10
```

### 64. Reverse a number
```text
FUNCTION reverseNum(N):
    reverse = 0
    WHILE N > 0:
        reverse = reverse * 10 + (N % 10)
        N = N / 10
    RETURN reverse

// TEST CASES
// Input: 123 -> Output: 321
// Input: 100 -> Output: 1
```

### 65. Count digits
```text
FUNCTION countDigits(N):
    count = 0
    IF N == 0: RETURN 1
    WHILE N > 0:
        N = N / 10
        Increment count
    RETURN count

// TEST CASES
// Input: 123 -> Output: 3
// Input: 0   -> Output: 1
```

### 66. Power of number
```text
FUNCTION power(base, exp):
    result = 1
    FOR i FROM 1 TO exp:
        result *= base
    RETURN result

// TEST CASES
// Input: 2, 3 -> Output: 8
// Input: 5, 0 -> Output: 1
```

### 67. GCD of two numbers
```text
FUNCTION gcd(a, b):
    WHILE b NOT EQUAL 0:
        temp = b
        b = a % b
        a = temp
    RETURN a

// TEST CASES
// Input: 12, 18 -> Output: 6
// Input: 7, 5   -> Output: 1
```

### 68. LCM of two numbers
```text
FUNCTION lcm(a, b):
    RETURN (a * b) / gcd(a, b)

// TEST CASES
// Input: 4, 6 -> Output: 12
// Input: 3, 5 -> Output: 15
```

### 69. Strong number
```text
FUNCTION isStrong(N):
    original = N, sum = 0
    WHILE N > 0:
        digit = N % 10
        sum += factorial(digit)
        N = N / 10
    RETURN sum EQUAL original

// TEST CASES
// Input: 145 -> Output: True (1! + 4! + 5! = 145)
// Input: 123 -> Output: False
```

### 70. Automorphic number
```text
FUNCTION isAutomorphic(N):
    square = N * N
    RETURN str(square) endsWith str(N)

// TEST CASES
// Input: 5 -> Output: True (25 ends with 5)
// Input: 6 -> Output: True (36 ends with 6)
// Input: 7 -> Output: False (49 does not end with 7)
```

## 5. SEARCHING & SORTING

### 71. Linear search
```text
FUNCTION linearSearch(arr, target):
    FOR i FROM 0 TO length of arr - 1:
        IF arr[i] EQUAL target:
            RETURN i
    RETURN -1

// TEST CASES
// Input: [10, 20, 30, 40], 30 -> Output: 2
// Input: [10, 20, 30], 50     -> Output: -1
```

### 72. Binary search (Sorted Array)
```text
FUNCTION binarySearch(arr, target):
    low = 0, high = length of arr - 1
    WHILE low <= high:
        mid = (low + high) / 2
        IF arr[mid] EQUAL target: RETURN mid
        ELSE IF arr[mid] < target: low = mid + 1
        ELSE: high = mid - 1
    RETURN -1

// TEST CASES
// Input: [10, 20, 30, 40], 30 -> Output: 2
// Input: [10, 20, 30], 25     -> Output: -1
```

### 73. Bubble sort
```text
FUNCTION bubbleSort(arr):
    FOR i FROM 0 TO len(arr)-1:
       FOR j FROM 0 TO len(arr)-i-1:
          IF arr[j] > arr[j+1]: Swap(arr[j], arr[j+1])

// TEST CASES
// Input: [5, 1, 4, 2, 8] -> Output: [1, 2, 4, 5, 8]
// Input: [3, 2, 1]       -> Output: [1, 2, 3]
```

### 74. Selection sort
```text
FUNCTION selectionSort(arr):
    FOR i FROM 0 TO len(arr)-1:
        minIndex = i
        FOR j FROM i+1 TO len(arr):
            IF arr[j] < arr[minIndex]: minIndex = j
        Swap(arr[i], arr[minIndex])

// TEST CASES
// Input: [64, 25, 12, 22, 11] -> Output: [11, 12, 22, 25, 64]
```

### 75. Insertion sort
```text
FUNCTION insertionSort(arr):
    FOR i FROM 1 TO len(arr)-1:
        key = arr[i]
        j = i - 1
        WHILE j >= 0 AND arr[j] > key:
            arr[j+1] = arr[j]
            j = j - 1
        arr[j+1] = key

// TEST CASES
// Input: [12, 11, 13, 5, 6] -> Output: [5, 6, 11, 12, 13]
```

### 76. Sort array of 0s, 1s, 2s
```text
FUNCTION sort012(arr):
    low = 0, mid = 0, high = len(arr)-1
    WHILE mid <= high:
        IF arr[mid] == 0: Swap(arr[low], arr[mid]), low++, mid++
        ELSE IF arr[mid] == 1: mid++
        ELSE: Swap(arr[mid], arr[high]), high--

// TEST CASES
// Input: [0, 1, 2, 0, 1, 2] -> Output: [0, 0, 1, 1, 2, 2]
// Input: [0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1] -> Output: [0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2]
```

### 77. Find kth largest element
```text
FUNCTION findKthLargest(arr, k):
    Sort(arr)
    RETURN arr[len(arr) - k]

// TEST CASES
// Input: [3, 2, 1, 5, 6, 4], 2 -> Output: 5
// Input: [3, 2, 3, 1, 2, 4, 5, 5, 6], 4 -> Output: 4
```

### 78. Find median of array
```text
FUNCTION findMedian(arr):
    Sort(arr)
    n = len(arr)
    IF n % 2 != 0: RETURN arr[n/2]
    ELSE: RETURN (arr[n/2] + arr[n/2 - 1]) / 2

// TEST CASES
// Input: [1, 3, 2]    -> Output: 2
// Input: [1, 2, 3, 4] -> Output: 2.5
```

### 79. Count sort (basic idea)
```text
FUNCTION countSort(arr):
    maxVal = findMax(arr)
    countArr = Array of size maxVal+1 initialized to 0
    FOR num IN arr: countArr[num]++
    index = 0
    FOR i FROM 0 TO maxVal:
        WHILE countArr[i] > 0:
            arr[index] = i
            index++, countArr[i]--

// TEST CASES
// Input: [1, 4, 1, 2, 7, 5, 2] -> Output: [1, 1, 2, 2, 4, 5, 7]
```

### 80. Sort string alphabetically
```text
FUNCTION sortString(str):
    arr = Convert str to char array
    Sort(arr)
    RETURN Convert arr to string

// TEST CASES
// Input: "edcba" -> Output: "abcde"
// Input: "geeks" -> Output: "eegks"
```

## 6. RECURSION & BASIC LOGIC

### 81. Reverse string using recursion
```text
(See Problem 17)

// TEST CASES
// Input: "hello" -> Output: "olleh"
```

### 82. Factorial using recursion
```text
FUNCTION factorialRec(N):
    IF N <= 1: RETURN 1
    RETURN N * factorialRec(N - 1)

// TEST CASES
// Input: 5 -> Output: 120
// Input: 1 -> Output: 1
```

### 83. Fibonacci using recursion
```text
FUNCTION fibonacciRec(N):
    IF N <= 1: RETURN N
    RETURN fibonacciRec(N - 1) + fibonacciRec(N - 2)

// TEST CASES
// Input: 6 -> Output: 8
// Input: 0 -> Output: 0
```

### 84. Sum of array using recursion
```text
FUNCTION sumRec(arr, n):
    IF n <= 0: RETURN 0
    RETURN sumRec(arr, n-1) + arr[n-1]

// TEST CASES
// Input: [1, 2, 3, 4, 5] -> Output: 15
```

### 85. Print numbers 1 to N using recursion
```text
FUNCTION printN(N):
    IF N == 0: RETURN
    printN(N - 1)
    PRINT N

// TEST CASES
// Input: 5 -> Output: 1 2 3 4 5
```

### 86. Check palindrome using recursion
```text
FUNCTION isPalindromeRec(str, start, end):
    IF start >= end: RETURN True
    IF str[start] != str[end]: RETURN False
    RETURN isPalindromeRec(str, start+1, end-1)

// TEST CASES
// Input: "madam" -> Output: True
// Input: "racecar" -> Output: True
// Input: "hello" -> Output: False
```

### 87. Count digits using recursion
```text
FUNCTION countDigitsRec(N):
    IF N == 0: RETURN 0
    RETURN 1 + countDigitsRec(N / 10)

// TEST CASES
// Input: 12345 -> Output: 5
// Input: 7     -> Output: 1
```

### 88. Power of number using recursion
```text
FUNCTION powerRec(base, exp):
    IF exp == 0: RETURN 1
    RETURN base * powerRec(base, exp - 1)

// TEST CASES
// Input: 2, 4 -> Output: 16
// Input: 5, 0 -> Output: 1
```

### 89. GCD using recursion
```text
FUNCTION gcdRec(a, b):
    IF b == 0: RETURN a
    RETURN gcdRec(b, a % b)

// TEST CASES
// Input: 48, 18 -> Output: 6
// Input: 10, 5  -> Output: 5
```

### 90. Remove duplicates recursively
```text
FUNCTION removeDupRec(str):
    IF len(str) <= 1: RETURN str
    IF str[0] == str[1]: RETURN removeDupRec(str[1:])
    RETURN str[0] + removeDupRec(str[1:])

// TEST CASES
// Input: "aabccba" -> Output: "abcba"
// Input: "xxxx"    -> Output: "x"
```

## 7. SCENARIO / OUTPUT-BASED QUESTIONS

### 91. Predict output of loop
```text
// Depends on specific code snippet.
// Strategy: Dry run step-by-step with variables.

// EXAMPLE
// for(i=0; i<3; i++) print(i); -> Output: 0 1 2
```

### 92. Predict output of string operations
```text
// Example: String concatenation vs addition
// "5" + 2 -> "52"
// 5 + 2 -> 7

// TEST CASES
// Input: "10" + 20 -> Output: "1020"
// Input: 10 + 20   -> Output: 30
```

### 93. Find error in given code
```text
// Look for:
// - Syntax errors (missing semicolons, braces)
// - Logical errors (off-by-one loops, uninitialized variables)
// - Runtime errors (division by zero, null pointer)
```

### 94. What will be printed (array index logic)
```text
// Watch for index out of bounds.
// 0-based vs 1-based indexing confusion.

// EXAMPLE
// arr = [1, 2, 3]; print(arr[3]); -> Error (IndexOutOfBounds)
```

### 95. Time complexity of simple loop
```text
// Single loop 1 to N -> O(N)
// Nested loop -> O(N^2)
// Loop doubling/halving -> O(log N)
```

### 96. Swap two numbers without temp
```text
FUNCTION swapNoTemp(a, b):
    a = a + b
    b = a - b
    a = a - b
    PRINT a, b

// TEST CASES
// Input: 5, 10 -> Output: 10, 5
```

### 97. Difference between == and equals()
```text
// == : Compares references (memory address)
// .equals() : Compares values (content)

// EXAMPLE
// s1="A", s2=new String("A"); s1==s2 -> False; s1.equals(s2) -> True
```

### 98. Difference between array and arraylist
```text
// Array: Fixed size, faster primitive access.
// ArrayList: Dynamic size, wrapper classes, slower.
```

### 99. Call by value vs reference
```text
// Call by Value: Copy of variable passed (primitives in Java).
// Call by Reference: Reference to memory passed (objects in Java - technically pass-by-value of reference).
```

### 100. Static vs non-static example
```text
// Static: Belongs to class, shared by all instances.
// Non-static: Belongs to specific object instance.
```

## ADDITIONAL 100 QUESTIONS

### 1. STRINGS – ADVANCED BASICS

#### 1. Find longest palindrome substring
```text
FUNCTION longestPalindrome(s):
    IF length of s < 1: RETURN ""
    start = 0, end = 0
    FOR i FROM 0 TO length of s:
        len1 = expandAroundCenter(s, i, i)
        len2 = expandAroundCenter(s, i, i+1)
        len = MAX(len1, len2)
        IF len > end - start:
            start = i - (len - 1) / 2
            end = i + len / 2
    RETURN substring of s from start to end

FUNCTION expandAroundCenter(s, left, right):
    WHILE left >= 0 AND right < length of s AND s[left] == s[right]:
        left--, right++
    RETURN right - left - 1

// TEST CASES
// Input: "babad" -> Output: "bab" (or "aba")
// Input: "cbbd"  -> Output: "bb"
```

#### 2. Find longest common prefix
```text
FUNCTION longestCommonPrefix(strs):
    IF strs is empty: RETURN ""
    prefix = strs[0]
    FOR i FROM 1 TO length of strs - 1:
        WHILE strs[i] does not start with prefix:
            prefix = substring of prefix from 0 to length-1
            IF prefix is empty: RETURN ""
    RETURN prefix

// TEST CASES
// Input: ["flower", "flow", "flight"] -> Output: "fl"
// Input: ["dog", "racecar", "car"]    -> Output: ""
```

#### 3. Check if string contains only digits
```text
FUNCTION isDigitsOnly(str):
    FOR EACH char IN str:
        IF char < '0' OR char > '9':
            RETURN False
    RETURN True

// TEST CASES
// Input: "12345" -> Output: True
// Input: "123a5" -> Output: False
```

#### 4. Count uppercase & lowercase letters
```text
FUNCTION countCase(str):
    upper = 0, lower = 0
    FOR EACH char IN str:
        IF char >= 'A' AND char <= 'Z': upper++
        IF char >= 'a' AND char <= 'z': lower++
    PRINT upper, lower

// TEST CASES
// Input: "Hello World" -> Output: Upper: 2, Lower: 8
```

#### 5. Remove special characters
```text
FUNCTION removeSpecialChars(str):
    result = ""
    FOR EACH char IN str:
        IF (char >= 'A' AND char <= 'Z') OR (char >= 'a' AND char <= 'z') OR (char >= '0' AND char <= '9'):
            Append char to result
    RETURN result

// TEST CASES
// Input: "$Hem$lo_World" -> Output: "HemloWorld"
```

#### 6. Find all permutations of string (basic)
```text
FUNCTION permute(str, l, r):
    IF l == r: PRINT str
    ELSE:
        FOR i FROM l TO r:
            Swap(str[l], str[i])
            permute(str, l+1, r)
            Swap(str[l], str[i]) // backtrack

// TEST CASES
// Input: "ABC" -> Output: ABC, ACB, BAC, BCA, CBA, CAB
```

#### 7. Check valid parentheses (basic stack)
```text
FUNCTION isValidParentheses(s):
    Stack stack
    FOR EACH char IN s:
        IF char IN ['(', '{', '[']:
            Push char to stack
        ELSE:
            IF stack is empty: RETURN False
            top = Pop stack
            IF (char == ')' AND top != '(') OR ... : RETURN False
    RETURN stack is empty

// TEST CASES
// Input: "()[]{}" -> Output: True
// Input: "(]"     -> Output: False
```

#### 8. Find duplicate words in sentence
```text
FUNCTION findDuplicateWords(sentence):
    words = Split sentence by space
    Map freqMap
    FOR EACH word IN words:
        freqMap[word]++
    FOR EACH word, count IN freqMap:
        IF count > 1: PRINT word

// TEST CASES
// Input: "big black bug bit a big black dog" -> Output: big, black
```

#### 9. Reverse each word in place
```text
FUNCTION reverseEachWord(sentence):
    words = Split sentence by space
    result = ""
    FOR EACH word IN words:
        result += reverseString(word) + " "
    RETURN result trimmed

// TEST CASES
// Input: "Hello World" -> Output: "olleH dlroW"
```

#### 10. Check isomorphic strings
```text
FUNCTION isIsomorphic(s, t):
    IF len(s) != len(t): RETURN False
    Map mapS, Map mapT
    FOR i FROM 0 TO len(s)-1:
        c1 = s[i], c2 = t[i]
        IF (c1 in mapS AND mapS[c1] != c2) OR (c2 in mapT AND mapT[c2] != c1):
            RETURN False
        mapS[c1] = c2
        mapT[c2] = c1
    RETURN True

// TEST CASES
// Input: "egg", "add" -> Output: True
// Input: "foo", "bar" -> Output: False
```

#### 11. Check pangram
```text
FUNCTION isPangram(sentence):
    Set letters
    FOR EACH char IN sentence:
        IF char is letter:
            Add lowercase(char) to letters
    RETURN size of letters == 26

// TEST CASES
// Input: "The quick brown fox jumps over the lazy dog" -> Output: True
// Input: "Hello World" -> Output: False
```

#### 12. Print all substrings
```text
FUNCTION printSubstrings(str):
    n = length of str
    FOR i FROM 0 TO n:
        FOR j FROM i+1 TO n:
            PRINT substring of str from i to j

// TEST CASES
// Input: "abc" -> Output: a, ab, abc, b, bc, c
```

#### 13. Remove consecutive duplicates
```text
FUNCTION removeConsecutiveDeep(str):
    IF len(str) < 2: RETURN str
    result = "" + str[0]
    FOR i FROM 1 TO len(str)-1:
        IF str[i] != str[i-1]:
            Append str[i] to result
    RETURN result

// TEST CASES
// Input: "aaabbc" -> Output: "abc"
// Input: "aabbcc" -> Output: "abc"
```

#### 14. Check if strings differ by one character
```text
FUNCTION differByOne(s1, s2):
    IF abs(len(s1) - len(s2)) > 1: RETURN False
    countDiff = 0
    i = 0, j = 0
    WHILE i < len(s1) AND j < len(s2):
        IF s1[i] != s2[j]:
            countDiff++
            IF countDiff > 1: RETURN False
            // Logic differs for same length vs different length
        ELSE:
            i++, j++
    RETURN True

// TEST CASES
// Input: "abc", "abd" -> Output: True
// Input: "abc", "xyz" -> Output: False
```

#### 15. Find smallest & largest word
```text
FUNCTION minMaxWords(sentence):
    words = Split sentence
    minWord = words[0], maxWord = words[0]
    FOR EACH word IN words:
        IF len(word) < len(minWord): minWord = word
        IF len(word) > len(maxWord): maxWord = word
    PRINT minWord, maxWord

// TEST CASES
// Input: "This is a test string" -> Output: Min: a, Max: string
```

### 2. ARRAYS – EXTENDED LOGIC

#### 16. Find leaders in array
```text
FUNCTION findLeaders(arr):
    n = len(arr)
    maxRight = arr[n-1]
    PRINT maxRight
    FOR i FROM n-2 DOWN TO 0:
        IF arr[i] > maxRight:
            maxRight = arr[i]
            PRINT maxRight

// TEST CASES
// Input: [16, 17, 4, 3, 5, 2] -> Output: 2, 5, 17
// Input: [1, 2, 3]            -> Output: 3
```

#### 17. Find equilibrium index
```text
FUNCTION equilibriumPoint(arr):
    totalSum = Sum(arr)
    leftSum = 0
    FOR i FROM 0 TO len(arr)-1:
        totalSum -= arr[i]
        IF leftSum == totalSum: RETURN i
        leftSum += arr[i]
    RETURN -1

// TEST CASES
// Input: [1, 3, 5, 2, 2] -> Output: 2 (Sum left: 4, Sum right: 4)
// Input: [1, 2, 3]       -> Output: -1
```

#### 18. Subarray with given sum
```text
FUNCTION subArraySum(arr, sum):
    currSum = arr[0], start = 0
    FOR i FROM 1 TO len(arr):
        WHILE currSum > sum AND start < i-1:
            currSum -= arr[start]
            start++
        IF currSum == sum:
            PRINT start, i-1
            RETURN
        IF i < len(arr): currSum += arr[i]

// TEST CASES
// Input: [1, 2, 3, 7, 5], 12 -> Output: 1, 3
// Input: [1, 2, 3], 10       -> Output: (Prints nothing or not found)
```

#### 19. Kadane’s algorithm (Maximum Subarray Sum)
```text
FUNCTION kadanes(arr):
    maxSoFar = -Infinity
    currMax = 0
    FOR EACH num IN arr:
        currMax += num
        IF maxSoFar < currMax: maxSoFar = currMax
        IF currMax < 0: currMax = 0
    RETURN maxSoFar

// TEST CASES
// Input: [-2, 1, -3, 4, -1, 2, 1, -5, 4] -> Output: 6
// Input: [1]                             -> Output: 1
```

#### 20. Find majority element
```text
FUNCTION majorityElement(arr):
    candidate = NULL, count = 0
    FOR EACH num IN arr:
        IF count == 0: candidate = num
        IF num == candidate: count++
        ELSE: count--
    RETURN candidate

// TEST CASES
// Input: [3, 2, 3]       -> Output: 3
// Input: [2, 2, 1, 1, 1, 2, 2] -> Output: 2
```

#### 21. Rearrange array alternately
```text
FUNCTION rearrangeAlternate(arr):
    Sort(arr)
    i = 0, j = len(arr)-1
    WHILE i < j:
        PRINT arr[j--]
        PRINT arr[i++]
    IF i == j: PRINT arr[i]

// TEST CASES
// Input: [1, 2, 3, 4, 5, 6] -> Output: 6 1 5 2 4 3
```

#### 22. Rotate array using reversal algorithm
```text
// Same as Q25/Q26

// TEST CASES
// Input: [1, 2, 3, 4, 5], k=2 -> Output: [3, 4, 5, 1, 2]
```

#### 23. Find union of two arrays
```text
FUNCTION unionArrays(arr1, arr2):
    Set s
    Add all elements of arr1 to s
    Add all elements of arr2 to s
    RETURN s

// TEST CASES
// Input: [1, 2, 3], [2, 3, 4] -> Output: {1, 2, 3, 4}
```

#### 24. Find intersection of two arrays
```text
FUNCTION intersectionArrays(arr1, arr2):
    Set s1
    Add all elements of arr1 to s1
    result = []
    FOR EACH num IN arr2:
        IF num IN s1:
            Add num to result
            Remove num from s1
    RETURN result

// TEST CASES
// Input: [1, 2, 3], [2, 3, 4] -> Output: [2, 3]
```

#### 25. Count pairs with given difference
```text
FUNCTION countDiffPairs(arr, k):
    count = 0
    FOR i FROM 0 TO len(arr):
        FOR j FROM i+1 TO len(arr):
            IF abs(arr[i] - arr[j]) == k: count++
    RETURN count

// TEST CASES
// Input: [1, 5, 3, 4, 2], 3 -> Output: 2 (pairs {1,4}, {2,5})
```

#### 26. Find peak element
```text
FUNCTION findPeak(arr):
    n = len(arr)
    IF n == 1: RETURN 0
    IF arr[0] >= arr[1]: RETURN 0
    IF arr[n-1] >= arr[n-2]: RETURN n-1
    FOR i FROM 1 TO n-2:
        IF arr[i] >= arr[i-1] AND arr[i] >= arr[i+1]:
            RETURN i

// TEST CASES
// Input: [1, 2, 3, 1] -> Output: 2 (Index of element 3)
// Input: [1, 2, 1, 3, 5, 6, 4] -> Output: 1 or 5
```

#### 27. Left rotate by 1
```text
FUNCTION leftRotateOne(arr):
    temp = arr[0]
    FOR i FROM 0 TO len(arr)-2:
        arr[i] = arr[i+1]
    arr[len(arr)-1] = temp

// TEST CASES
// Input: [1, 2, 3, 4, 5] -> Output: [2, 3, 4, 5, 1]
```

#### 28. Find minimum difference pair
```text
FUNCTION minDiffPair(arr):
    Sort(arr)
    minDiff = Infinity
    FOR i FROM 0 TO len(arr)-2:
        IF arr[i+1] - arr[i] < minDiff:
            minDiff = arr[i+1] - arr[i]
    RETURN minDiff

// TEST CASES
// Input: [2, 4, 5, 9, 7] -> Output: 1 (Pair {4,5})
```

#### 29. Product of array except self
```text
FUNCTION productExceptSelf(arr):
    n = len(arr)
    left = Array(n), right = Array(n), prod = Array(n)
    left[0] = 1, right[n-1] = 1
    FOR i FROM 1 TO n-1: left[i] = arr[i-1] * left[i-1]
    FOR i FROM n-2 DOWN TO 0: right[i] = arr[i+1] * right[i+1]
    FOR i FROM 0 TO n-1: prod[i] = left[i] * right[i]
    RETURN prod

// TEST CASES
// Input: [1, 2, 3, 4] -> Output: [24, 12, 8, 6]
```

#### 30. Find subarray with max product
```text
FUNCTION maxProductSubarray(arr):
    maxSoFar = arr[0], minSoFar = arr[0], result = maxSoFar
    FOR i FROM 1 TO len(arr)-1:
        curr = arr[i]
        tempMax = MAX(curr, MAX(curr*maxSoFar, curr*minSoFar))
        minSoFar = MIN(curr, MIN(curr*maxSoFar, curr*minSoFar))
        maxSoFar = tempMax
        result = MAX(maxSoFar, result)
    RETURN result

// TEST CASES
// Input: [2, 3, -2, 4] -> Output: 6
// Input: [-2, 0, -1]   -> Output: 0
```

#### 31. Check circular array rotation
```text
// Check if one array is a rotated version of another (Same as Q19 String Rotation)
```

#### 32. Separate positive and negative numbers
```text
FUNCTION separatePosNeg(arr):
    j = 0
    FOR i FROM 0 TO len(arr)-1:
        IF arr[i] < 0:
            IF i != j: Swap(arr[i], arr[j])
            j++

// TEST CASES
// Input: [-1, 2, -3, 4, 5, 6, -7, 8, 9] -> Output: [-1, -3, -7, 4, 5, 6, 2, 8, 9] (Relative order not preserved by this algo)
```

#### 33. Count distinct elements
```text
FUNCTION countDistinct(arr):
    Set s
    FOR EACH num IN arr: Add num to s
    RETURN size of s

// TEST CASES
// Input: [10, 20, 20, 10, 30] -> Output: 3
```

#### 34. Replace element with next greatest
```text
FUNCTION replaceNextGreatest(arr):
    maxFromRight = -1
    FOR i FROM len(arr)-1 DOWN TO 0:
        temp = arr[i]
        arr[i] = maxFromRight
        maxFromRight = MAX(maxFromRight, temp)

// TEST CASES
// Input: [16, 17, 4, 3, 5, 2] -> Output: [17, 5, 5, 5, 2, -1]
```

#### 35. Find smallest subarray with sum > X
```text
FUNCTION smallestSubWithSum(arr, x):
    minLen = len(arr) + 1
    currSum = 0, start = 0, end = 0
    WHILE end < len(arr):
        currSum += arr[end++]
        WHILE currSum > x:
            minLen = MIN(minLen, end - start)
            currSum -= arr[start++]
    RETURN minLen

// TEST CASES
// Input: [1, 4, 45, 6, 0, 19], 51 -> Output: 3
```

### 3. MATRIX / 2D ARRAY

#### 36. Matrix addition
```text
FUNCTION addMatrices(A, B):
    C = new Matrix
    FOR i FROM 0 TO rows:
        FOR j FROM 0 TO cols:
            C[i][j] = A[i][j] + B[i][j]
    RETURN C

// TEST CASES
// Input: [[1, 2], [3, 4]], [[5, 6], [7, 8]] -> Output: [[6, 8], [10, 12]]
```

#### 37. Matrix multiplication
```text
FUNCTION multiplyMatrices(A, B):
    C = new Matrix
    FOR i FROM 0 TO rowsA:
        FOR j FROM 0 TO colsB:
            FOR k FROM 0 TO colsA:
                C[i][j] += A[i][k] * B[k][j]
    RETURN C

// TEST CASES
// Input: [[1, 2], [3, 4]], [[1, 0], [0, 1]] -> Output: [[1, 2], [3, 4]]
```

#### 38. Transpose of matrix
```text
FUNCTION transpose(A):
    T = new Matrix
    FOR i FROM 0 TO rows:
        FOR j FROM 0 TO cols:
            T[j][i] = A[i][j]
    RETURN T

// TEST CASES
// Input: [[1, 2], [3, 4]] -> Output: [[1, 3], [2, 4]]
```

#### 39. Rotate matrix 90 degrees
```text
FUNCTION rotateMatrix(A):
    Transpose(A)
    Reverse each row of A

// TEST CASES
// Input: [[1, 2, 3], [4, 5, 6], [7, 8, 9]] 
// Output: [[7, 4, 1], [8, 5, 2], [9, 6, 3]]
```

#### 40. Print matrix in spiral order
```text
FUNCTION spiralOrder(matrix):
    top = 0, bottom = rows-1, left = 0, right = cols-1
    WHILE top <= bottom AND left <= right:
        FOR i FROM left TO right: PAINT matrix[top][i]
        top++
        FOR i FROM top TO bottom: PAINT matrix[i][right]
        right--
        IF top <= bottom:
            FOR i FROM right DOWN TO left: PAINT matrix[bottom][i]
            bottom--
        IF left <= right:
            FOR i FROM bottom DOWN TO top: PAINT matrix[i][left]
            left++

// TEST CASES
// Input: [[1, 2, 3], [4, 5, 6], [7, 8, 9]] 
// Output: 1 2 3 6 9 8 7 4 5
```

#### 41. Search element in sorted matrix
```text
FUNCTION searchMatrix(matrix, target):
    row = 0, col = cols-1
    WHILE row < rows AND col >= 0:
        IF matrix[row][col] == target: RETURN True
        ELSE IF matrix[row][col] > target: col--
        ELSE: row++
    RETURN False

// TEST CASES
// Input: [[1, 3, 5], [7, 9, 11], [13, 15, 17]], 9 -> Output: True
```

#### 42. Diagonal sum (Primary & Secondary)
```text
FUNCTION diagonalSum(matrix):
    sum = 0
    FOR i FROM 0 TO n-1:
        sum += matrix[i][i]
        sum += matrix[i][n-i-1]
    IF n is odd: sum -= matrix[n/2][n/2] // Remove center added twice
    RETURN sum

// TEST CASES
// Input: [[1, 2, 3], [4, 5, 6], [7, 8, 9]] -> Output: 25 (1+5+9 + 3+5+7 - 5 = 25)
```

#### 43. Print boundary elements
```text
FUNCTION printBoundary(matrix):
    FOR col FROM 0 TO cols-1: PRINT matrix[0][col]
    FOR row FROM 1 TO rows-1: PRINT matrix[row][cols-1]
    FOR col FROM cols-2 DOWN TO 0: PRINT matrix[rows-1][col]
    FOR row FROM rows-2 DOWN TO 1: PRINT matrix[row][0]

// TEST CASES
// Input: [[1, 2, 3], [4, 5, 6], [7, 8, 9]] -> Output: 1 2 3 6 9 8 7 4
```

#### 44. Check symmetric matrix
```text
FUNCTION isSymmetric(matrix):
    RETURN matrix == Transpose(matrix)

// TEST CASES
// Input: [[1, 2], [2, 1]] -> Output: True
// Input: [[1, 2], [3, 4]] -> Output: False
```

#### 45. Interchange rows and columns
```text
// Same as Transpose
```

#### 46. Count zeros and ones
```text
FUNCTION countZeroOne(matrix):
    zeros = 0, ones = 0
    FOR EACH cell IN matrix:
        IF cell == 0: zeros++
        IF cell == 1: ones++
    PRINT zeros, ones

// TEST CASES
// Input: [[0, 1], [1, 0]] -> Output: Zeros: 2, Ones: 2
```

#### 47. Row with maximum 1s
```text
FUNCTION rowMaxOnes(matrix):
    maxOnes = 0, rowIndex = -1
    FOR i FROM 0 TO rows-1:
        count = count(matrix[i], 1)
        IF count > maxOnes:
            maxOnes = count
            rowIndex = i
    RETURN rowIndex

// TEST CASES
// Input: [[0, 1, 1], [0, 0, 1], [1, 1, 1]] -> Output: 2
```

#### 48. Matrix palindrome check
```text
// Check if matrix is symmetric or if rows/cols form palindromes
```

#### 49. Snake pattern printing
```text
FUNCTION snakePattern(matrix):
    FOR i FROM 0 TO rows-1:
        IF i % 2 == 0:
            FOR j FROM 0 TO cols-1: PRINT matrix[i][j]
        ELSE:
            FOR j FROM cols-1 DOWN TO 0: PRINT matrix[i][j]

// TEST CASES
// Input: [[1, 2], [3, 4]] -> Output: 1 2 4 3
```

#### 50. Identity matrix check
```text
FUNCTION isIdentity(matrix):
    FOR i FROM 0 TO rows-1:
        FOR j FROM 0 TO cols-1:
            IF i == j AND matrix[i][j] != 1: RETURN False
            IF i != j AND matrix[i][j] != 0: RETURN False
    RETURN True

// TEST CASES
// Input: [[1, 0], [0, 1]] -> Output: True
// Input: [[1, 1], [0, 1]] -> Output: False
```

### 4. LINKED LIST (BASIC ONLY)

#### 51-60. Linked List Operations
```text
// Structure Node: { data, next }

// 51. Create LL: new Node(data)
// 52. Traverse: while head != null: print head.data, head = head.next
// 53. Insert Begin: newNode.next = head, head = newNode
// 54. Insert End: traverse to last, last.next = newNode
// 55. Delete Node: prev.next = curr.next
// 56. Reverse LL:
FUNCTION reverseLL(head):
    prev = NULL, curr = head
    WHILE curr != NULL:
        nextTemp = curr.next
        curr.next = prev
        prev = curr
        curr = nextTemp
    RETURN prev

// 57. Middle Element: slow = head, fast = head. While fast && fast.next: slow = slow.next, fast = fast.next.next. Return slow.
// 58. Detect Loop: Floyd’s Cycle (Slow/Fast). If slow == fast -> Loop.
// 59. Count Nodes: count = 0. While head: count++.
// 60. Merge 2 Lists: Dummy node, compare heads, attach smaller, move pointers.

// TEST CASES (Applied to general logic)
// Input: 1 -> 2 -> 3 -> 4 -> 5 -> NULL (Reverse)
// Output: 5 -> 4 -> 3 -> 2 -> 1 -> NULL
// Input: 1 -> 2 -> 3 -> NULL (Middle)
// Output: 2
```

### 5. STACK & QUEUE (BASIC)

#### 61-70. Stack/Queue Operations
```text
// 61. Stack (Array): push(arr[++top] = x), pop(return arr[top--])
// 62. Queue (Array): enqueue(arr[rear++] = x), dequeue(return arr[front++])
// 63. Reverse String (Stack): Push all chars, Pop all chars.
// 64. Balanced Parentheses: (See Q7 in Additional)
// 65. Stack using Queue: Use 2 queues or push + rotate.
// 66. Queue using Stack: Use 2 stacks (input/output).
// 67. Next Greater Element: Use Monotonic Stack.
// 68. Evaluate Postfix: Push operands, Pop 2 on operator, calc, push result.
// 69. Reverse Stack: Recursion (insertAtBottom).
// 70. Circular Queue: (rear + 1) % size.

// TEST CASES (Examples)
// Balanced Parentheses: Input: "((()))" -> Output: True
// Evaluate Postfix: Input: "231*+9-" -> Output: -4 (2 + (3*1) - 9)
```

### 6. HASHING / MAP LOGIC (VERY IMPORTANT)

#### 71-80. Hashing Problems
```text
// 71. Freq elements: Map<Element, Count> (See Q6/28 in Top 100)
// 72. First Repeating: Iterate, check if in Set -> Return.
// 73. First Non-Repeating: Map counts, iterate again -> check count == 1.
// 74. Two Sum: Map<NeededVal, Index>. If curr in Map -> Found.
// 75. Group Anagrams: Map<SortedStr, List<Original>>.
// 76. Count Distinct Chars: Set.size()
// 77. Majority Element: Map counts -> if count > n/2.
// 78. Check Subset: Put larger array in Map/Freq. Check if smaller array elements exist.
// 79. Common Elements: Intersection (See Q33).
// 80. Longest Substring No Repeats: Sliding Window with Set/Map of indices.

// TEST CASES
// Two Sum: [2, 7, 11, 15], Target 9 -> Output: [0, 1]
// Group Anagrams: ["eat", "tea", "tan", "ate", "nat", "bat"]
// Output: [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]]
```

### 7. OOPS & LANGUAGE FUNDAMENTALS

#### 81-90. OOPS Concepts
```text
// 81. Method Overloading: Same method name, different params.
// Example: void print(int a), void print(String s)

// 82. Method Overriding: Subclass same method signature as Superclass.
// Example: Class Dog extends Animal { void sound() { print "Bark" } }

// 83. Encapsulation: private vars, public getters/setters.
// Example: Class Person { private name; public getName() { return name; } }

// 84. Inheritance: class A extends B.
// 85. Polymorphism: runtime (override), compile-time (overload).
// 86. Abstract Class vs Interface: abstract can have state/impl, interface (pre-Java 8) only contract.
// 87. Final: var (const), method (no override), class (no inherit).
// 88. Exception Handling: try-catch-finally blocks.
// 89. Singleton: private constructor, static instance.
// 90. Immutable: final class, final fields, no setters.
```

### 8. SQL & BASIC CS

#### 91-100. SQL & CS
```sql
-- 91. Second Highest Salary
SELECT MAX(Salary) FROM Employee WHERE Salary < (SELECT MAX(Salary) FROM Employee);
-- Example Input Table: Salaries [100, 200, 300]
-- Output: 200

-- 92. Delete (DML, can rollback) vs Truncate (DDL, fast, no rollback) vs Drop (Remove table).

-- 93. Inner Join (Matches) vs Left Join (All left + Matches right).

-- 94. Find Duplicate Rows
SELECT col, COUNT(*) FROM table GROUP BY col HAVING COUNT(*) > 1;
-- Example Input: ['A', 'B', 'A', 'C']
-- Output: 'A', 2

-- 95. PK (Unique + Not Null) vs FK (Refers to PK).
-- 96. Normalization: 1NF, 2NF, 3NF (Remove redundancy).
-- 97. Index: Data structure (B-Tree) to speed up retrieval.
-- 98. ACID: Atomicity, Consistency, Isolation, Durability.
-- 99. Deadlock: Two processes waiting for each other.
-- 100. GET (Read data) vs POST (Create data).
```
