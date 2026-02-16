# Pattern Programs (56-65)

## 56. Right Triangle Star Pattern
**Principle**: Nested loops. Outer loop for rows, inner for columns.
**Question**: Print a right triangle of stars.
**Code**:
```java
/*
*
**
***
****
*/
public class RightTriangle {
    public static void main(String[] args) {
        int n = 4;
        for(int i=1; i<=n; i++) {
            for(int j=1; j<=i; j++) System.out.print("*");
            System.out.println();
        }
    }
}
```

## 57. Left Triangle Star Pattern
**Principle**: Print spaces then stars.
**Question**: Print a left-aligned triangle.
**Code**:
```java
/*
   *
  **
 ***
****
*/
public class LeftTriangle {
    public static void main(String[] args) {
        int n = 4;
        for(int i=1; i<=n; i++) {
            for(int j=1; j<=n-i; j++) System.out.print(" ");
            for(int k=1; k<=i; k++) System.out.print("*");
            System.out.println();
        }
    }
}
```

## 58. Pyramid Star Pattern
**Principle**: Spaces decreasing, Stars increasing (odd numbers 1, 3, 5...).
**Question**: Print a pyramid.
**Code**:
```java
/*
  *
 ***
*****
*/
public class Pyramid {
    public static void main(String[] args) {
        int n = 3;
        for(int i=1; i<=n; i++) {
            for(int j=1; j<=n-i; j++) System.out.print(" ");
            for(int k=1; k<=2*i-1; k++) System.out.print("*");
            System.out.println();
        }
    }
}
```

## 59. Diamond Pattern
**Principle**: Combine normal pyramid and inverted pyramid.
**Question**: Print a diamond shape.
**Code**:
```java
public class Diamond {
    public static void main(String[] args) {
        int n = 3;
        // Upper
        for(int i=1; i<=n; i++) {
            for(int j=1; j<=n-i; j++) System.out.print(" ");
            for(int k=1; k<=2*i-1; k++) System.out.print("*");
            System.out.println();
        }
        // Lower
        for(int i=n-1; i>=1; i--) {
            for(int j=1; j<=n-i; j++) System.out.print(" ");
            for(int k=1; k<=2*i-1; k++) System.out.print("*");
            System.out.println();
        }
    }
}
```

## 60. Number Triangle Pattern
**Principle**: Print `j` (column index).
**Question**: Print number triangle.
**Code**:
```java
/*
1
12
123
*/
public class NumTriangle {
    public static void main(String[] args) {
        int n = 3;
        for(int i=1; i<=n; i++) {
            for(int j=1; j<=i; j++) System.out.print(j);
            System.out.println();
        }
    }
}
```

## 61. Checkered/Floyd's Triangle (0-1)
**Principle**: If `(i+j)` is even print 1, else 0.
**Question**: Print 0-1 triangle.
**Code**:
```java
/*
1
0 1
1 0 1
*/
public class BinaryTriangle {
    public static void main(String[] args) {
        int n = 3;
        for(int i=1; i<=n; i++) {
            for(int j=1; j<=i; j++) {
                if((i+j)%2 == 0) System.out.print("1 ");
                else System.out.print("0 ");
            }
            System.out.println();
        }
    }
}
```

## 62. Pascal's Triangle
**Principle**: `Val = Val * (i-j) / j`.
**Question**: Print Pascal's Triangle.
**Code**:
```java
/*
 1
 1 1
 1 2 1
*/
public class Pascal {
    public static void main(String[] args) {
        int n = 3;
        for (int i=0; i<n; i++) {
            for (int s=0; s<n-i; s++) System.out.print(" ");
            int val = 1;
            for (int j=0; j<=i; j++) {
                System.out.print(val + " ");
                val = val * (i - j) / (j + 1);
            }
            System.out.println();
        }
    }
}
```

## 63. Rhombus Pattern
**Principle**: Shifted square.
**Question**: Print a solid rhombus.
**Code**:
```java
/*
  ****
 ****
****
*/
public class Rhombus {
    public static void main(String[] args) {
        int n = 4;
        for(int i=1; i<=n; i++) {
            for(int j=1; j<=n-i; j++) System.out.print(" ");
            for(int j=1; j<=n; j++) System.out.print("*");
            System.out.println();
        }
    }
}
```

## 64. Hollow Square
**Principle**: Print star if boundary index, else space.
**Question**: Print a hollow square pattern.
**Code**:
```java
public class HollowSquare {
    public static void main(String[] args) {
        int n = 4;
        for(int i=1; i<=n; i++) {
            for(int j=1; j<=n; j++) {
                if(i==1 || i==n || j==1 || j==n) System.out.print("*");
                else System.out.print(" ");
            }
            System.out.println();
        }
    }
}
```

## 65. Spiral Pattern (Number Grid)
**Principle**: Min distance to edge.
**Question**: Print concentric number layers.
**Code**:
```java
public class SpiralNum {
    public static void main(String[] args) {
        int  n = 4; // Size
        int len = 2*n - 1;
        for(int i=0; i<len; i++){
            for(int j=0; j<len; j++){
                int min = i < j ? i : j;
                min = min < len-i ? min : len-i-1;
                min = min < len-j-1 ? min : len-j-1;
                System.out.print((n-min) + " ");
            }
            System.out.println();
        }
    }
}
```
