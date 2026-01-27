package java_solutions;

public class PatternProblems {

    // 41. Right triangle star pattern
    public static void rightTriangle(int N) {
        for (int i = 1; i <= N; i++) {
            for (int k = 1; k <= i; k++)
                System.out.print("*");
            System.out.println();
        }
    }

    // 42. Left triangle star pattern
    public static void leftTriangle(int N) {
        for (int i = 1; i <= N; i++) {
            for (int k = 1; k <= N - i; k++)
                System.out.print(" ");
            for (int k = 1; k <= i; k++)
                System.out.print("*");
            System.out.println();
        }
    }

    // 43. Pyramid star pattern
    public static void pyramid(int N) {
        for (int i = 1; i <= N; i++) {
            for (int k = 1; k <= N - i; k++)
                System.out.print(" ");
            for (int k = 1; k <= i; k++)
                System.out.print("* ");
            System.out.println();
        }
    }

    // 44. Inverted pyramid
    public static void invertedPyramid(int N) {
        for (int i = N; i >= 1; i--) {
            for (int k = 1; k <= N - i; k++)
                System.out.print(" ");
            for (int k = 1; k <= i; k++)
                System.out.print("* ");
            System.out.println();
        }
    }

    // 45. Diamond pattern
    public static void diamond(int N) {
        pyramid(N);
        for (int i = N - 1; i >= 1; i--) {
            for (int k = 1; k <= N - i; k++)
                System.out.print(" ");
            for (int k = 1; k <= i; k++)
                System.out.print("* ");
            System.out.println();
        }
    }

    // 46. Number pyramid
    public static void numberPyramid(int N) {
        for (int i = 1; i <= N; i++) {
            for (int k = 1; k <= N - i; k++)
                System.out.print(" ");
            for (int j = 1; j <= i; j++) {
                System.out.print(j + " ");
            }
            System.out.println();
        }
    }

    // 47. Floyd’s triangle
    public static void floydsTriangle(int N) {
        int num = 1;
        for (int i = 1; i <= N; i++) {
            for (int j = 1; j <= i; j++) {
                System.out.print(num + " ");
                num++;
            }
            System.out.println();
        }
    }

    // 48. Hollow square pattern
    public static void hollowSquare(int N) {
        for (int i = 1; i <= N; i++) {
            if (i == 1 || i == N) {
                for (int k = 1; k <= N; k++)
                    System.out.print("*");
            } else {
                System.out.print("*");
                for (int k = 1; k <= N - 2; k++)
                    System.out.print(" ");
                System.out.print("*");
            }
            System.out.println();
        }
    }

    // 49. Hollow pyramid
    public static void hollowPyramid(int N) {
        for (int i = 1; i <= N; i++) {
            for (int j = 1; j <= N - i; j++)
                System.out.print(" ");
            for (int k = 1; k <= (2 * i - 1); k++) {
                if (k == 1 || k == (2 * i - 1) || i == N)
                    System.out.print("*");
                else
                    System.out.print(" ");
            }
            System.out.println();
        }
    }

    // 50. Zig-zag star pattern
    public static void zigZag(int N) {
        for (int i = 1; i <= 3; i++) {
            for (int j = 1; j <= N; j++) {
                if (((i + j) % 4 == 0) || (i == 2 && j % 4 == 0))
                    System.out.print("*");
                else
                    System.out.print(" ");
            }
            System.out.println();
        }
    }

    // 51. Butterfly pattern
    public static void butterfly(int N) {
        for (int i = 1; i <= N; i++) {
            for (int j = 1; j <= i; j++)
                System.out.print("*");
            for (int j = 1; j <= 2 * (N - i); j++)
                System.out.print(" ");
            for (int j = 1; j <= i; j++)
                System.out.print("*");
            System.out.println();
        }
        for (int i = N; i >= 1; i--) {
            for (int j = 1; j <= i; j++)
                System.out.print("*");
            for (int j = 1; j <= 2 * (N - i); j++)
                System.out.print(" ");
            for (int j = 1; j <= i; j++)
                System.out.print("*");
            System.out.println();
        }
    }

    // 52. Pascal triangle
    public static void pascalTriangle(int N) {
        for (int i = 0; i < N; i++) {
            int val = 1;
            for (int j = 0; j <= i; j++) {
                System.out.print(val + " ");
                val = val * (i - j) / (j + 1);
            }
            System.out.println();
        }
    }

    // 53. Number increasing pattern
    public static void numberIncreasing(int N) {
        for (int i = 1; i <= N; i++) {
            for (int j = 1; j <= i; j++)
                System.out.print(j + " ");
            System.out.println();
        }
    }

    // 54. Number increasing reverse
    public static void numberIncreasingReverse(int N) {
        for (int i = N; i >= 1; i--) {
            for (int j = 1; j <= i; j++)
                System.out.print(j + " ");
            System.out.println();
        }
    }

    // 55. Number changing pyramid
    public static void numberChangingPyramid(int N) {
        int num = 1;
        for (int i = 1; i <= N; i++) {
            for (int j = 1; j <= i; j++) {
                System.out.print(num + " ");
                num++;
            }
            System.out.println();
        }
    }

    public static void main(String[] args) {
        System.out.println("41. Right Triangle (N=3):");
        rightTriangle(3);
        System.out.println("\n42. Left Triangle (N=3):");
        leftTriangle(3);
        System.out.println("\n43. Pyramid (N=3):");
        pyramid(3);
        System.out.println("\n44. Inverted Pyramid (N=3):");
        invertedPyramid(3);
        System.out.println("\n45. Diamond (N=3):");
        diamond(3);
        System.out.println("\n46. Number Pyramid (N=3):");
        numberPyramid(3);
        System.out.println("\n47. Floyd's Triangle (N=3):");
        floydsTriangle(3);
        System.out.println("\n48. Hollow Square (N=3):");
        hollowSquare(3);
        System.out.println("\n49. Hollow Pyramid (N=3):");
        hollowPyramid(3);
        System.out.println("\n50. Zig-Zag (N=9):");
        zigZag(9);
        System.out.println("\n51. Butterfly (N=2):");
        butterfly(2);
        System.out.println("\n52. Pascal Triangle (N=3):");
        pascalTriangle(3);
        System.out.println("\n53. Number Increasing (N=3):");
        numberIncreasing(3);
        System.out.println("\n54. Number Increasing Reverse (N=3):");
        numberIncreasingReverse(3);
        System.out.println("\n55. Number Changing Pyramid (N=3):");
        numberChangingPyramid(3);
    }
}
