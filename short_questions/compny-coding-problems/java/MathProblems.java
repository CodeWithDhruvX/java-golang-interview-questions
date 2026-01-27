package java_solutions;

import java.util.*;

public class MathProblems {

    // 56. Check prime number
    public static boolean isPrime(int N) {
        if (N <= 1)
            return false;
        for (int i = 2; i * i <= N; i++) {
            if (N % i == 0)
                return false;
        }
        return true;
    }

    // 57. Print prime numbers in range
    public static void printPrimes(int start, int end) {
        for (int i = start; i <= end; i++) {
            if (isPrime(i))
                System.out.print(i + " ");
        }
        System.out.println();
    }

    // 58. Fibonacci series
    public static void fibonacci(int N) {
        int a = 0, b = 1;
        System.out.print(a + " " + b + " ");
        for (int i = 2; i < N; i++) {
            int c = a + b;
            System.out.print(c + " ");
            a = b;
            b = c;
        }
        System.out.println();
    }

    // 59. Factorial of a number
    public static int factorial(int N) {
        int fact = 1;
        for (int i = 1; i <= N; i++)
            fact *= i;
        return fact;
    }

    // 60. Check Armstrong number
    public static boolean isArmstrong(int N) {
        int temp = N, sum = 0;
        int digits = String.valueOf(N).length();
        while (temp > 0) {
            int digit = temp % 10;
            sum += Math.pow(digit, digits);
            temp /= 10;
        }
        return sum == N;
    }

    // 61. Sum of digits
    public static int sumDigits(int N) {
        int sum = 0;
        while (N > 0) {
            sum += N % 10;
            N /= 10;
        }
        return sum;
    }

    // 62. Reverse a number
    public static int reverseNumber(int N) {
        int rev = 0;
        while (N > 0) {
            rev = rev * 10 + N % 10;
            N /= 10;
        }
        return rev;
    }

    // 63. Check Palindrome Number
    public static boolean isPalindromeNum(int N) {
        return N == reverseNumber(N);
    }

    // 64. Swap two numbers without temp
    public static void swapNumbers(int a, int b) {
        System.out.println("Before: a=" + a + ", b=" + b);
        a = a + b;
        b = a - b;
        a = a - b;
        System.out.println("After: a=" + a + ", b=" + b);
    }

    // 65. LCM of two numbers
    public static int gcd(int a, int b) {
        if (b == 0)
            return a;
        return gcd(b, a % b);
    }

    public static int lcm(int a, int b) {
        return (a * b) / gcd(a, b);
    }

    // 66. GCD of two numbers
    public static int calculateGCD(int a, int b) {
        return gcd(a, b);
    }

    // 67. Check if number is power of 2
    public static boolean isPowerOfTwo(int N) {
        return N > 0 && (N & (N - 1)) == 0;
    }

    // 68. Find square root without built-in
    public static double sqrt(double N) {
        double x = N;
        double y = 1.0;
        double e = 0.000001;
        while (x - y > e) {
            x = (x + y) / 2;
            y = N / x;
        }
        return x;
    }

    // 69. Check perfect number
    public static boolean isPerfect(int N) {
        int sum = 1;
        for (int i = 2; i * i <= N; i++) {
            if (N % i == 0) {
                if (i * i != N)
                    sum = sum + i + N / i;
                else
                    sum = sum + i;
            }
        }
        return sum == N && N != 1;
    }

    // 70. Count digits in a number
    public static int countDigits(int N) {
        return String.valueOf(N).length();
    }

    public static void main(String[] args) {
        System.out.println("56. Is Prime: 7 -> " + isPrime(7));
        System.out.println("57. Print Primes (1-20):");
        printPrimes(1, 20);
        System.out.println("58. Fibonacci (N=5):");
        fibonacci(5);
        System.out.println("59. Factorial: 5 -> " + factorial(5));
        System.out.println("60. Is Armstrong: 153 -> " + isArmstrong(153));
        System.out.println("61. Sum Digits: 123 -> " + sumDigits(123));
        System.out.println("62. Reverse Number: 123 -> " + reverseNumber(123));
        System.out.println("63. Is Palindrome Num: 121 -> " + isPalindromeNum(121));
        System.out.println("64. Swap Numbers (5, 10):");
        swapNumbers(5, 10);
        System.out.println("65. LCM: 4, 6 -> " + lcm(4, 6));
        System.out.println("66. GCD: 12, 18 -> " + calculateGCD(12, 18));
        System.out.println("67. Is Power of Two: 8 -> " + isPowerOfTwo(8));
        System.out.printf("68. Sqrt: 16 -> %.2f\n", sqrt(16));
        System.out.println("69. Is Perfect: 6 -> " + isPerfect(6));
        System.out.println("70. Count Digits: 1234 -> " + countDigits(1234));
    }
}
