# Basic Number Programs (1-15)

## 1. Fibonacci Series
**Principle**: Each number is the sum of the two preceding ones, starting from 0 and 1.
**Question**: Write a program to print the Fibonacci series up to `n` terms.
**Code**:
```java
import java.util.Scanner;

public class Fibonacci {
    public static void main(String[] args) {
        int n = 10, first = 0, second = 1;
        System.out.println("Fibonacci Series up to " + n + " terms:");

        for (int i = 0; i < n; i++) {
            System.out.print(first + " ");
            int next = first + second;
            first = second;
            second = next;
        }
    }
}
```

## 2. Check Prime Number
**Principle**: A prime number is a number greater than 1 that has no positive divisors other than 1 and itself.
**Question**: Write a program to check if a given number is prime.
**Code**:
```java
public class PrimeCheck {
    public static void main(String[] args) {
        int num = 29;
        boolean isPrime = true;

        if (num <= 1) isPrime = false;
        else {
            for (int i = 2; i <= Math.sqrt(num); i++) {
                if (num % i == 0) {
                    isPrime = false;
                    break;
                }
            }
        }
        System.out.println(num + " is prime? " + isPrime);
    }
}
```

## 3. Factorial of a Number
**Principle**: Factorial of `n` is the product of all positive integers less than or equal to `n`.
**Question**: Write a program to find the factorial of a number.
**Code**:
```java
public class Factorial {
    public static void main(String[] args) {
        int num = 5;
        long factorial = 1;
        for(int i = 1; i <= num; ++i) {
            factorial *= i;
        }
        System.out.println("Factorial of " + num + " = " + factorial);
    }
}
```

## 4. Palindrome Number
**Principle**: A number that remains the same when its digits are reversed.
**Question**: Check if a number is a Palindrome.
**Code**:
```java
public class PalindromeNumber {
    public static void main(String[] args) {
        int num = 121, reversed = 0, original = num;
        
        while(num != 0) {
            int digit = num % 10;
            reversed = reversed * 10 + digit;
            num /= 10;
        }
        
        System.out.println(original + (original == reversed ? " is" : " is not") + " a palindrome.");
    }
}
```

## 5. Armstrong Number
**Principle**: An integer is an Armstrong number if the sum of its digits raised to the power of the number of digits equals the number itself. (e.g., 153 = 1^3 + 5^3 + 3^3).
**Question**: Check if a number is an Armstrong number.
**Code**:
```java
public class Armstrong {
    public static void main(String[] args) {
        int num = 153, original = num, result = 0, n = 0;
        
        for (;original != 0; original /= 10, ++n); // Count digits
        
        original = num;
        while (original != 0) {
            int remainder = original % 10;
            result += Math.pow(remainder, n);
            original /= 10;
        }
        
        System.out.println(num + (result == num ? " is" : " is not") + " an Armstrong number.");
    }
}
```

## 6. Reverse a Number
**Principle**: Extract last digit check, add to new number check.
**Question**: Reverse a given integer.
**Code**:
```java
public class ReverseNumber {
    public static void main(String[] args) {
        int num = 1234, reversed = 0;
        while(num != 0) {
            int digit = num % 10;
            reversed = reversed * 10 + digit;
            num /= 10;
        }
        System.out.println("Reversed: " + reversed);
    }
}
```

## 7. Swap Two Numbers without Third Variable
**Principle**: Use arithmetic operations (+/- or */%) or bitwise XOR.
**Question**: Swap two numbers without using a temporary variable.
**Code**:
```java
public class SwapNoTemp {
    public static void main(String[] args) {
        int a = 10, b = 20;
        a = a + b; // 30
        b = a - b; // 10
        a = a - b; // 20
        System.out.println("a: " + a + ", b: " + b);
    }
}
```

## 8. Check Even or Odd
**Principle**: Divisible by 2 (remainder 0) or use bitwise AND (`num & 1 == 0`).
**Question**: Check if a number is even or odd.
**Code**:
```java
public class EvenOdd {
    public static void main(String[] args) {
        int num = 11;
        System.out.println(num + " is " + ((num % 2 == 0) ? "Even" : "Odd"));
    }
}
```

## 9. Check Leap Year
**Principle**: Divisible by 4 and (not divisible by 100 OR divisible by 400).
**Question**: Check if a year is a leap year.
**Code**:
```java
public class LeapYear {
    public static void main(String[] args) {
        int year = 2024;
        boolean leap = false;
        
        if (year % 4 == 0) {
            if (year % 100 == 0) {
                if (year % 400 == 0) leap = true;
                else leap = false;
            } else leap = true;
        } else leap = false;
        
        System.out.println(year + (leap ? " is" : " is not") + " a leap year.");
    }
}
```

## 10. Greatest of Three Numbers
**Principle**: Conditional checks.
**Question**: Find the largest among three numbers.
**Code**:
```java
public class LargestOfThree {
    public static void main(String[] args) {
        int a = 10, b = 20, c = 15;
        int max = (a > b) ? (a > c ? a : c) : (b > c ? b : c);
        System.out.println("Largest: " + max);
    }
}
```

## 11. Sum of Digits
**Principle**: Extract digits modulo 10 and add to sum.
**Question**: Calculate sum of digits of a number.
**Code**:
```java
public class SumDigits {
    public static void main(String[] args) {
        int num = 1234, sum = 0;
        while(num > 0) {
            sum += num % 10;
            num /= 10;
        }
        System.out.println("Sum: " + sum);
    }
}
```

## 12. GCD of Two Numbers
**Principle**: Euclidean algorithm: GCD(a, b) = GCD(b, a%b).
**Question**: Find GCD (HCF) of two numbers.
**Code**:
```java
public class GCD {
    public static void main(String[] args) {
        int n1 = 81, n2 = 153;
        while(n2 != 0) {
            int temp = n2;
            n2 = n1 % n2;
            n1 = temp;
        }
        System.out.println("GCD: " + n1);
    }
}
```

## 13. LCM of Two Numbers
**Principle**: LCM * GCD = n1 * n2.
**Question**: Find LCM of two numbers.
**Code**:
```java
public class LCM {
    public static void main(String[] args) {
        int n1 = 72, n2 = 120;
        int gcd = 1;
        for(int i = 1; i <= n1 && i <= n2; ++i) {
            if(n1 % i == 0 && n2 % i == 0) gcd = i;
        }
        int lcm = (n1 * n2) / gcd;
        System.out.println("LCM: " + lcm);
    }
}
```

## 14. Perfect Number
**Principle**: Sum of positive divisors excluding the number itself equals the number.
**Question**: Check if a number is a Perfect Number.
**Code**:
```java
public class PerfectNumber {
    public static void main(String[] args) {
        int num = 28, sum = 0;
        for(int i = 1; i < num; i++) {
            if(num % i == 0) sum += i;
        }
        System.out.println(num + (sum == num ? " is" : " is not") + " a Perfect Number.");
    }
}
```

## 15. Print ASCII Value
**Principle**: Cast char to int.
**Question**: Print ASCII value of a character.
**Code**:
```java
public class ASCII {
    public static void main(String[] args) {
        char ch = 'a';
        System.out.println("ASCII of " + ch + " is " + (int)ch);
    }
}
```
