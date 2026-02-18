# 18. Java Programs (Basic Numbers Logic)

**Q: Fibonacci Series**
> "The Fibonacci series is a sequence where each number is the sum of the two preceding ones: 0, 1, 1, 2, 3, 5, 8...
>
> To print it, you start with `a=0` and `b=1`.
> Inside a loop, you print `a`. Then calculate `next = a + b`.
> Shift everything: `a = b` and `b = next`. Repeat `n` times.
>
> If you need the *nth* number specifically, you can use Recursion (`fib(n-1) + fib(n-2)`), but that's very slow (O(2^n)). Iteration (O(n)) is always better for production code."

**Indepth:**
> **Optimization**: A recursive Fibonacci implementation `fib(n-1) + fib(n-2)` has exponential time complexity O(2^n). It recomputes the same values over and over. Always use Memoization (storing results in a map) or simple Iteration for O(n).


---

**Q: Check Prime Number**
> "A prime number is divisible only by 1 and itself.
>
> To check if `n` is prime:
> 1.  Handle edge cases: if `n <= 1`, return false.
> 2.  Loop from `i = 2` up to `Math.sqrt(n)`.
> 3.  If `n % i == 0`, it's not prime. Return false immediately.
> 4.  If the loop finishes without finding a divisor, return true.
>
> Why `sqrt(n)`? Because if a number has a factor larger than its square root, the *other* factor must be smaller than the square root, so we would have already found it."

**Indepth:**
> **Sieve of Eratosthenes**: If you need to find *all* primes up to N, checking them one by one is slow. The Sieve algorithm creates a boolean array and eliminates multiples of each prime found, which is significantly faster.


---

**Q: Factorial of a Number**
> "Factorial of 5 (written 5!) is `5 * 4 * 3 * 2 * 1 = 120`.
>
> You can do this recursively: `return n * factorial(n-1);` (Base case: if n is 0 or 1, return 1).
> Or iteratively: `result = 1;` loop `i` from 2 to `n`, `result *= i`.
>
> **Watch out**: Factorials grow incredibly fast. `13!` allows overflows a standard `int`. `21!` overflows a `long`. For anything bigger, you **must** use `BigInteger`."

**Indepth:**
> **Recursion Depth**: Recursive factorial is prone to `StackOverflowError` for large inputs (thousands), whereas the iterative version can run until memory runs out (assuming you use `BigInteger`).


---

**Q: Palindrome Number**
> "To check if a number like `121` is a palindrome, you need to reverse it mathematically.
>
> 1.  Store original number in `temp`. Initialize `reversed = 0`.
> 2.  While `temp > 0`:
>     *   Get last digit: `digit = temp % 10`.
>     *   Append to reversed: `reversed = (reversed * 10) + digit`.
>     *   Remove last digit: `temp = temp / 10`.
> 3.  Check if `original == reversed`."

**Indepth:**
> **Strings**: You *could* convert the number to a String (`Integer.toString(n)`) and reverse it using `StringBuilder`. This is easier to write but slower due to memory allocation and parsing overhead.


---

**Q: Armstrong Number**
> "An Armstrong number (like 153) is equal to the sum of its digits raised to the power of the number of digits.
> For 153 (3 digits): `1^3 + 5^3 + 3^3 = 1 + 125 + 27 = 153`.
>
> The logic is similar to Palindrome:
> 1.  Count digits first.
> 2.  Loop through the number, extract each digit.
> 3.  Add `Math.pow(digit, count)` to a running sum.
> 4.  Compare sum with original."

**Indepth:**
> **Hardcoding**: Many candidates forget to count the digits first and just assume `Math.pow(digit, 3)`. Armstrong numbers are defined by the power of *number of digits* (Example: 1634 is `1^4 + 6^4 + 3^4 + 4^4`).


---

**Q: Swap Two Numbers without Third Variable**
> "This is a classic 'cool trick' question.
>
> Assume `a = 10`, `b = 20`.
> 1.  `a = a + b;` (a becomes 30)
> 2.  `b = a - b;` (30 - 20 = 10, so b is now original a)
> 3.  `a = a - b;` (30 - 10 = 20, so a is now original b)
>
> It works, but in real life, just use a temporary variable. It's more readable and avoids potential integer overflow issues."

**Indepth:**
> **XOR Swap**: A safer way (avoiding overflow) is using XOR: `a = a ^ b; b = a ^ b; a = a ^ b;`. It works because XOR is its own inverse. However, it's less readable and not necessarily faster on modern CPUs.


---

**Q: Check Leap Year**
> "A year is a leap year if:
> 1.  It is divisible by 4.
> 2.  **EXCEPT** if it's divisible by 100, then it is NOT a leap year.
> 3.  **UNLESS** it is also divisible by 400, then it IS a leap year.
>
> Logic: `(year % 4 == 0 && year % 100 != 0) || (year % 400 == 0)`."

**Indepth:**
> **Why 100/400?**: The Earth takes 365.2425 days to orbit. Adding a day every 4 years (365.25) is slightly too much. Skipping 100 years corrects it, but skipping 400 adds it back key to keep the calendar accurate over centuries.


---

**Q: GCD and LCM**
> "**GCD (Greatest Common Divisor)**: Use Euclid's algorithm.
> Recursive: `gcd(a, b)` -> if `b == 0` return `a`, else return `gcd(b, a % b)`.
>
> **LCM (Least Common Multiple)**: Once you have GCD, LCM is easy.
> Formula: `(a * b) / GCD(a, b)`."

**Indepth:**
> **Euclid**: The Euclidean algorithm (`gcd(b, a % b)`) is one of the oldest known algorithms. It works because the GCD of two numbers also divides their difference.


---

**Q: Perfect Number**
> "A number is Perfect if the sum of its proper divisors equals the number itself.
> Example: 6. Divisors are 1, 2, 3. Sum = 1 + 2 + 3 = 6.
>
> Logic: Loop from 1 to `n/2`. If `n % i == 0`, add `i` to sum. Compare sum to `n`."

**Indepth:**
> **Rarity**: Perfect numbers are extremely rare. The first few are 6, 28, 496, 8128. Don't try to find them by brute-force for large ranges.


---

**Q: Sum of Digits**
> "Very similar to reversing a number.
> Loop while `n > 0`:
> 1.  `sum += n % 10;` (Add last digit)
> 2.  `n /= 10;` (Remove last digit)
>
> If you need the 'recursive sum' (sum digits until you get a single digit), use the modulo 9 trick: `return (n == 0) ? 0 : (n % 9 == 0) ? 9 : n % 9;`."

**Indepth:**
> **Digital Root**: The recursive sum of digits until 1 digit remains is called the "Digital Root". The `n % 9` trick works because of congruences in base-10 arithmetic.

