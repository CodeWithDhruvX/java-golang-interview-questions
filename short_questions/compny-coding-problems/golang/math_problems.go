package main

import (
	"fmt"
	"math"
	"strconv"
)

// 56. Check prime number
func isPrime(N int) bool {
	if N <= 1 {
		return false
	}
	for i := 2; i*i <= N; i++ {
		if N%i == 0 {
			return false
		}
	}
	return true
}

// 57. Print prime numbers in range
func printPrimes(start, end int) {
	for i := start; i <= end; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}

// 58. Fibonacci series
func fibonacci(N int) {
	a, b := 0, 1
	fmt.Printf("%d %d ", a, b)
	for i := 2; i < N; i++ {
		c := a + b
		fmt.Printf("%d ", c)
		a, b = b, c
	}
	fmt.Println()
}

// 59. Factorial of a number
func factorial(N int) int {
	fact := 1
	for i := 1; i <= N; i++ {
		fact *= i
	}
	return fact
}

// 60. Check Armstrong number
func isArmstrong(N int) bool {
	temp := N
	sum := 0
	digits := len(strconv.Itoa(N))
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}
	return sum == N
}

// 61. Sum of digits
func sumDigits(N int) int {
	sum := 0
	for N > 0 {
		sum += N % 10
		N /= 10
	}
	return sum
}

// 62. Reverse a number
func reverseNumber(N int) int {
	rev := 0
	for N > 0 {
		rev = rev*10 + N%10
		N /= 10
	}
	return rev
}

// 63. Check Palindrome Number
func isPalindromeNum(N int) bool {
	return N == reverseNumber(N)
}

// 64. Swap two numbers without temp
func swapNumbers(a, b int) {
	fmt.Printf("Before: a=%d, b=%d\n", a, b)
	a = a + b
	b = a - b
	a = a - b
	fmt.Printf("After: a=%d, b=%d\n", a, b)
}

// 65. LCM of two numbers
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

// 66. GCD of two numbers
func calculateGCD(a, b int) int {
	return gcd(a, b)
}

// 67. Check if number is power of 2
func isPowerOfTwo(N int) bool {
	return N > 0 && (N&(N-1)) == 0
}

// 68. Find square root without built-in
func sqrt(N float64) float64 {
	x := N
	y := 1.0
	e := 0.000001
	for x-y > e {
		x = (x + y) / 2
		y = N / x
	}
	return x
}

// 69. Check perfect number
func isPerfect(N int) bool {
	sum := 1
	for i := 2; i*i <= N; i++ {
		if N%i == 0 {
			if i*i != N {
				sum = sum + i + N/i
			} else {
				sum = sum + i
			}
		}
	}
	return sum == N && N != 1
}

// 70. Count digits in a number
func countDigits(N int) int {
	return len(strconv.Itoa(N))
}

func main() {
	fmt.Println("56. Is Prime:")
	fmt.Println("7 ->", isPrime(7))
	fmt.Println("4 ->", isPrime(4))

	fmt.Println("\n57. Print Primes (1-20):")
	printPrimes(1, 20)

	fmt.Println("\n58. Fibonacci (N=5):")
	fibonacci(5)

	fmt.Println("\n59. Factorial:")
	fmt.Println("5 ->", factorial(5))

	fmt.Println("\n60. Is Armstrong:")
	fmt.Println("153 ->", isArmstrong(153))
	fmt.Println("123 ->", isArmstrong(123))

	fmt.Println("\n61. Sum Digits:")
	fmt.Println("123 ->", sumDigits(123))

	fmt.Println("\n62. Reverse Number:")
	fmt.Println("123 ->", reverseNumber(123))

	fmt.Println("\n63. Is Palindrome Num:")
	fmt.Println("121 ->", isPalindromeNum(121))
	fmt.Println("123 ->", isPalindromeNum(123))

	fmt.Println("\n64. Swap Numbers (5, 10):")
	swapNumbers(5, 10)

	fmt.Println("\n65. LCM:")
	fmt.Println("4, 6 ->", lcm(4, 6))

	fmt.Println("\n66. GCD:")
	fmt.Println("12, 18 ->", calculateGCD(12, 18))

	fmt.Println("\n67. Is Power of Two:")
	fmt.Println("8 ->", isPowerOfTwo(8))
	fmt.Println("6 ->", isPowerOfTwo(6))

	fmt.Println("\n68. Sqrt:")
	fmt.Printf("16 -> %.2f\n", sqrt(16))

	fmt.Println("\n69. Is Perfect:")
	fmt.Println("6 ->", isPerfect(6))
	fmt.Println("28 ->", isPerfect(28))
	fmt.Println("12 ->", isPerfect(12))

	fmt.Println("\n70. Count Digits:")
	fmt.Println("1234 ->", countDigits(1234))
}
