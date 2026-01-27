package main

import (
	"fmt"
	"strings"
)

// 41. Right triangle star pattern
func rightTriangle(N int) {
	for i := 1; i <= N; i++ {
		fmt.Println(strings.Repeat("*", i))
	}
}

// 42. Left triangle star pattern
func leftTriangle(N int) {
	for i := 1; i <= N; i++ {
		fmt.Print(strings.Repeat(" ", N-i))
		fmt.Println(strings.Repeat("*", i))
	}
}

// 43. Pyramid star pattern
func pyramid(N int) {
	for i := 1; i <= N; i++ {
		fmt.Print(strings.Repeat(" ", N-i))
		fmt.Println(strings.Repeat("* ", i))
	}
}

// 44. Inverted pyramid
func invertedPyramid(N int) {
	for i := N; i >= 1; i-- {
		fmt.Print(strings.Repeat(" ", N-i))
		fmt.Println(strings.Repeat("* ", i))
	}
}

// 45. Diamond pattern
func diamond(N int) {
	pyramid(N)
	// To strictly follow standard diamond, we remove the top row of inverted to avoid duplicating the middle line
	// However, the pseudocode says "Handle middle row overlap if needed", let's replicate pyramid + inverted for simplicity as per pseudocode structure
	// Or better, let's implement standard diamond logic to match test cases
	for i := N - 1; i >= 1; i-- {
		fmt.Print(strings.Repeat(" ", N-i))
		fmt.Println(strings.Repeat("* ", i))
	}
}

// 46. Number pyramid
func numberPyramid(N int) {
	for i := 1; i <= N; i++ {
		fmt.Print(strings.Repeat(" ", N-i))
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}

// 47. Floyd’s triangle
func floydsTriangle(N int) {
	num := 1
	for i := 1; i <= N; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", num)
			num++
		}
		fmt.Println()
	}
}

// 48. Hollow square pattern
func hollowSquare(N int) {
	for i := 1; i <= N; i++ {
		if i == 1 || i == N {
			fmt.Println(strings.Repeat("*", N))
		} else {
			fmt.Print("*")
			fmt.Print(strings.Repeat(" ", N-2))
			fmt.Println("*")
		}
	}
}

// 49. Hollow pyramid
func hollowPyramid(N int) {
	for i := 1; i <= N; i++ {
		fmt.Print(strings.Repeat(" ", N-i))
		for k := 1; k <= (2*i - 1); k++ {
			if k == 1 || k == (2*i-1) || i == N {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// 50. Zig-zag star pattern
func zigZag(N int) {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= N; j++ {
			if ((i+j)%4 == 0) || (i == 2 && j%4 == 0) {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// 51. Butterfly pattern
func butterfly(N int) {
	// Upper half
	for i := 1; i <= N; i++ {
		fmt.Print(strings.Repeat("*", i))
		fmt.Print(strings.Repeat(" ", 2*(N-i)))
		fmt.Println(strings.Repeat("*", i))
	}
	// Lower half
	for i := N; i >= 1; i-- {
		fmt.Print(strings.Repeat("*", i))
		fmt.Print(strings.Repeat(" ", 2*(N-i)))
		fmt.Println(strings.Repeat("*", i))
	}
}

// 52. Pascal triangle
func pascalTriangle(N int) {
	for i := 0; i < N; i++ {
		val := 1
		for j := 0; j <= i; j++ {
			fmt.Printf("%d ", val)
			val = val * (i - j) / (j + 1)
		}
		fmt.Println()
	}
}

// 53. Number increasing pattern
func numberIncreasing(N int) {
	for i := 1; i <= N; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}

// 54. Number increasing reverse
func numberIncreasingReverse(N int) {
	for i := N; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
	}
}

// 55. Number changing pyramid
func numberChangingPyramid(N int) {
	num := 1
	for i := 1; i <= N; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", num)
			num++
		}
		fmt.Println()
	}
}

func main() {
	fmt.Println("41. Right Triangle (N=3):")
	rightTriangle(3)

	fmt.Println("\n42. Left Triangle (N=3):")
	leftTriangle(3)

	fmt.Println("\n43. Pyramid (N=3):")
	pyramid(3)

	fmt.Println("\n44. Inverted Pyramid (N=3):")
	invertedPyramid(3)

	fmt.Println("\n45. Diamond (N=3):")
	diamond(3)

	fmt.Println("\n46. Number Pyramid (N=3):")
	numberPyramid(3)

	fmt.Println("\n47. Floyd's Triangle (N=3):")
	floydsTriangle(3)

	fmt.Println("\n48. Hollow Square (N=3):")
	hollowSquare(3)

	fmt.Println("\n49. Hollow Pyramid (N=3):")
	hollowPyramid(3)

	fmt.Println("\n50. Zig-Zag (N=9):")
	zigZag(9)

	fmt.Println("\n51. Butterfly (N=2):")
	butterfly(2)

	fmt.Println("\n52. Pascal Triangle (N=3):")
	pascalTriangle(3)

	fmt.Println("\n53. Number Increasing (N=3):")
	numberIncreasing(3)

	fmt.Println("\n54. Number Increasing Reverse (N=3):")
	numberIncreasingReverse(3)

	fmt.Println("\n55. Number Changing Pyramid (N=3):")
	numberChangingPyramid(3)
}
