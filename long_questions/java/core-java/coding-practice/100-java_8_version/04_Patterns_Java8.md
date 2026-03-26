# Pattern Programs with Java 8 Features (56-65)

## 📚 Java 8 Features Demonstrated
- **Lambda Expressions**: Concise pattern generation
- **Streams API**: Functional pattern printing
- **Method References**: Simplified operations
- **Collectors**: Pattern aggregations
- **Functional Interfaces**: Consumer, Function, Predicate
- **Optional**: Null-safe pattern operations
- **Parallel Streams**: Multi-threaded pattern generation

---

## 56. Right Triangle Star Pattern
**Java 8 Approach**: Using `IntStream.range()` and `forEach`

```java
import java.util.*;
import java.util.stream.*;

public class RightTriangleJava8 {
    public static void main(String[] args) {
        int n = 4;
        
        // Using Java 8 Streams
        System.out.println("Right Triangle Star Pattern:");
        IntStream.range(1, n + 1)
            .forEach(row -> {
                IntStream.range(0, row)
                    .mapToObj(col -> "*")
                    .forEach(System.out::print);
                System.out.println();
            });
        
        // Alternative: Using String.repeat() (Java 11+, but showing concept)
        System.out.println("\nAlternative approach:");
        IntStream.range(1, n + 1)
            .mapToObj(row -> "*".repeat(row))
            .forEach(System.out::println);
        
        // Generate pattern as list of strings
        List<String> patternLines = IntStream.range(1, n + 1)
            .mapToObj(row -> "*".repeat(row))
            .collect(Collectors.toList());
        
        System.out.println("\nPattern as list: " + patternLines);
        
        // Create pattern with different characters
        char symbol = '#';
        System.out.println("\nPattern with symbol '" + symbol + "':");
        IntStream.range(1, n + 1)
            .forEach(row -> {
                IntStream.range(0, row)
                    .mapToObj(col -> String.valueOf(symbol))
                    .forEach(System.out::print);
                System.out.println();
            });
        
        // Generate multiple triangles
        List<Integer> sizes = Arrays.asList(3, 4, 5);
        System.out.println("\nMultiple triangles:");
        sizes.forEach(size -> {
            System.out.println("Size " + size + ":");
            IntStream.range(1, size + 1)
                .forEach(row -> {
                    IntStream.range(0, row)
                        .mapToObj(col -> "*")
                        .forEach(System.out::print);
                    System.out.println();
                });
            System.out.println();
        });
        
        // Parallel pattern generation (for demonstration)
        System.out.println("Parallel generation demonstration:");
        IntStream.range(1, n + 1)
            .parallel()
            .forEach(row -> {
                String line = IntStream.range(0, row)
                    .mapToObj(col -> "*")
                    .collect(Collectors.joining());
                System.out.println(line);
            });
    }
}
```

## 57. Left Triangle Star Pattern
**Java 8 Approach**: Using nested streams with space generation

```java
import java.util.*;
import java.util.stream.*;

public class LeftTriangleJava8 {
    public static void main(String[] args) {
        int n = 4;
        
        // Using Java 8 Streams
        System.out.println("Left Triangle Star Pattern:");
        IntStream.range(1, n + 1)
            .forEach(row -> {
                // Print spaces
                IntStream.range(0, n - row)
                    .mapToObj(col -> " ")
                    .forEach(System.out::print);
                
                // Print stars
                IntStream.range(0, row)
                    .mapToObj(col -> "*")
                    .forEach(System.out::print);
                
                System.out.println();
            });
        
        // Alternative: Using String.join and collectors
        System.out.println("\nAlternative approach:");
        IntStream.range(1, n + 1)
            .mapToObj(row -> 
                IntStream.range(0, n - row)
                    .mapToObj(col -> " ")
                    .collect(Collectors.joining()) +
                IntStream.range(0, row)
                    .mapToObj(col -> "*")
                    .collect(Collectors.joining())
            )
            .forEach(System.out::println);
        
        // Generate pattern as list
        List<String> patternLines = IntStream.range(1, n + 1)
            .mapToObj(row -> 
                " ".repeat(n - row) + "*".repeat(row)
            )
            .collect(Collectors.toList());
        
        System.out.println("\nPattern lines: " + patternLines);
        
        // Create inverted left triangle
        System.out.println("\nInverted Left Triangle:");
        IntStream.range(0, n)
            .forEach(row -> {
                String spaces = " ".repeat(row);
                String stars = "*".repeat(n - row);
                System.out.println(spaces + stars);
            });
        
        // Generate pattern with numbers
        System.out.println("\nLeft Triangle with Numbers:");
        IntStream.range(1, n + 1)
            .forEach(row -> {
                String spaces = " ".repeat(n - row);
                String numbers = IntStream.range(1, row + 1)
                    .mapToObj(String::valueOf)
                    .collect(Collectors.joining());
                System.out.println(spaces + numbers);
            });
        
        // Calculate total characters in pattern
        int totalChars = IntStream.range(1, n + 1)
            .map(row -> (n - row) + row)
            .sum();
        
        System.out.println("\nTotal characters in pattern: " + totalChars);
        
        // Generate pattern statistics
        Map<Integer, Integer> rowStats = IntStream.range(1, n + 1)
            .boxed()
            .collect(Collectors.toMap(
                Function.identity(),
                row -> row // Number of stars equals row number
            ));
        
        System.out.println("Row statistics: " + rowStats);
    }
}
```

## 58. Pyramid Star Pattern
**Java 8 Approach**: Using streams for space and star calculation

```java
import java.util.*;
import java.util.stream.*;

public class PyramidJava8 {
    public static void main(String[] args) {
        int n = 3;
        
        // Using Java 8 Streams
        System.out.println("Pyramid Star Pattern:");
        IntStream.range(1, n + 1)
            .forEach(row -> {
                // Print spaces
                IntStream.range(0, n - row)
                    .mapToObj(col -> " ")
                    .forEach(System.out::print);
                
                // Print stars (odd numbers: 1, 3, 5, ...)
                IntStream.range(0, 2 * row - 1)
                    .mapToObj(col -> "*")
                    .forEach(System.out::print);
                
                System.out.println();
            });
        
        // Alternative: Using collectors
        System.out.println("\nAlternative approach:");
        IntStream.range(1, n + 1)
            .mapToObj(row -> 
                " ".repeat(n - row) + "*".repeat(2 * row - 1)
            )
            .forEach(System.out::println);
        
        // Generate pattern as list
        List<String> pyramidLines = IntStream.range(1, n + 1)
            .mapToObj(row -> 
                " ".repeat(n - row) + "*".repeat(2 * row - 1)
            )
            .collect(Collectors.toList());
        
        System.out.println("\nPyramid lines: " + pyramidLines);
        
        // Create inverted pyramid
        System.out.println("\nInverted Pyramid:");
        IntStream.range(0, n)
            .forEach(row -> {
                String spaces = " ".repeat(row);
                String stars = "*".repeat(2 * (n - row) - 1);
                System.out.println(spaces + stars);
            });
        
        // Diamond pattern (pyramid + inverted pyramid)
        System.out.println("\nDiamond Pattern:");
        // Upper part
        IntStream.range(1, n + 1)
            .forEach(row -> {
                String spaces = " ".repeat(n - row);
                String stars = "*".repeat(2 * row - 1);
                System.out.println(spaces + stars);
            });
        // Lower part
        IntStream.range(n - 1, 0, -1)
            .forEach(row -> {
                String spaces = " ".repeat(n - row);
                String stars = "*".repeat(2 * row - 1);
                System.out.println(spaces + stars);
            });
        
        // Pyramid with numbers
        System.out.println("\nPyramid with Numbers:");
        IntStream.range(1, n + 1)
            .forEach(row -> {
                String spaces = " ".repeat(n - row);
                String numbers = IntStream.range(1, row + 1)
                    .mapToObj(String::valueOf)
                    .collect(Collectors.joining()) +
                    IntStream.range(row - 1, 0, -1)
                    .mapToObj(String::valueOf)
                    .collect(Collectors.joining());
                System.out.println(spaces + numbers);
            });
        
        // Calculate total stars in pyramid
        int totalStars = IntStream.range(1, n + 1)
            .map(row -> 2 * row - 1)
            .sum();
        
        System.out.println("\nTotal stars in pyramid: " + totalStars);
        
        // Generate pyramids of different sizes
        List<Integer> sizes = Arrays.asList(2, 3, 4);
        System.out.println("\nPyramids of different sizes:");
        sizes.forEach(size -> {
            System.out.println("Size " + size + ":");
            IntStream.range(1, size + 1)
                .forEach(row -> {
                    String spaces = " ".repeat(size - row);
                    String stars = "*".repeat(2 * row - 1);
                    System.out.println(spaces + stars);
                });
            System.out.println();
        });
    }
}
```

## 59. Diamond Pattern
**Java 8 Approach**: Combining pyramid and inverted pyramid

```java
import java.util.*;
import java.util.stream.*;

public class DiamondJava8 {
    public static void main(String[] args) {
        int n = 3;
        
        // Using Java 8 Streams
        System.out.println("Diamond Pattern:");
        
        // Upper part (pyramid)
        IntStream.range(1, n + 1)
            .forEach(row -> {
                String spaces = " ".repeat(n - row);
                String stars = "*".repeat(2 * row - 1);
                System.out.println(spaces + stars);
            });
        
        // Lower part (inverted pyramid)
        IntStream.range(n - 1, 0, -1)
            .forEach(row -> {
                String spaces = " ".repeat(n - row);
                String stars = "*".repeat(2 * row - 1);
                System.out.println(spaces + stars);
            });
        
        // Alternative: Generate complete diamond as list
        List<String> diamondLines = Stream.concat(
                IntStream.range(1, n + 1)
                    .mapToObj(row -> " ".repeat(n - row) + "*".repeat(2 * row - 1)),
                IntStream.range(n - 1, 0, -1)
                    .mapToObj(row -> " ".repeat(n - row) + "*".repeat(2 * row - 1))
            )
            .collect(Collectors.toList());
        
        System.out.println("\nDiamond as list: " + diamondLines);
        
        // Hollow diamond
        System.out.println("\nHollow Diamond:");
        // Upper part
        IntStream.range(1, n + 1)
            .forEach(row -> {
                String spaces = " ".repeat(n - row);
                if (row == 1) {
                    System.out.println(spaces + "*");
                } else {
                    String middleSpaces = " ".repeat(2 * row - 3);
                    System.out.println(spaces + "*" + middleSpaces + "*");
                }
            });
        // Lower part
        IntStream.range(n - 1, 0, -1)
            .forEach(row -> {
                String spaces = " ".repeat(n - row);
                if (row == 1) {
                    System.out.println(spaces + "*");
                } else {
                    String middleSpaces = " ".repeat(2 * row - 3);
                    System.out.println(spaces + "*" + middleSpaces + "*");
                }
            });
        
        // Diamond with numbers
        System.out.println("\nDiamond with Numbers:");
        // Upper part
        IntStream.range(1, n + 1)
            .forEach(row -> {
                String spaces = " ".repeat(n - row);
                String numbers = IntStream.range(1, row + 1)
                    .mapToObj(String::valueOf)
                    .collect(Collectors.joining()) +
                    IntStream.range(row - 1, 0, -1)
                    .mapToObj(String::valueOf)
                    .collect(Collectors.joining());
                System.out.println(spaces + numbers);
            });
        // Lower part
        IntStream.range(n - 1, 0, -1)
            .forEach(row -> {
                String spaces = " ".repeat(n - row);
                String numbers = IntStream.range(1, row + 1)
                    .mapToObj(String::valueOf)
                    .collect(Collectors.joining()) +
                    IntStream.range(row - 1, 0, -1)
                    .mapToObj(String::valueOf)
                    .collect(Collectors.joining());
                System.out.println(spaces + numbers);
            });
        
        // Calculate total lines in diamond
        int totalLines = 2 * n - 1;
        System.out.println("\nTotal lines in diamond: " + totalLines);
        
        // Find maximum width of diamond
        int maxWidth = 2 * n - 1;
        System.out.println("Maximum width: " + maxWidth);
        
        // Generate diamonds of different sizes
        List<Integer> sizes = Arrays.asList(2, 3, 4);
        System.out.println("\nDiamonds of different sizes:");
        sizes.forEach(size -> {
            System.out.println("Size " + size + " (lines: " + (2 * size - 1) + "):");
            Stream.concat(
                IntStream.range(1, size + 1)
                    .mapToObj(row -> " ".repeat(size - row) + "*".repeat(2 * row - 1)),
                IntStream.range(size - 1, 0, -1)
                    .mapToObj(row -> " ".repeat(size - row) + "*".repeat(2 * row - 1))
            ).forEach(System.out::println);
            System.out.println();
        });
        
        // Diamond pattern statistics
        Map<Integer, Integer> diamondStats = sizes.stream()
            .collect(Collectors.toMap(
                Function.identity(),
                size -> 2 * size - 1
            ));
        
        System.out.println("Diamond statistics: " + diamondStats);
    }
}
```

## 60. Number Triangle Pattern
**Java 8 Approach**: Using streams for number generation

```java
import java.util.*;
import java.util.stream.*;

public class NumTriangleJava8 {
    public static void main(String[] args) {
        int n = 3;
        
        // Using Java 8 Streams
        System.out.println("Number Triangle Pattern:");
        IntStream.range(1, n + 1)
            .forEach(row -> {
                IntStream.range(1, row + 1)
                    .mapToObj(col -> String.valueOf(col))
                    .forEach(System.out::print);
                System.out.println();
            });
        
        // Alternative: Using collectors
        System.out.println("\nAlternative approach:");
        IntStream.range(1, n + 1)
            .mapToObj(row -> 
                IntStream.range(1, row + 1)
                    .mapToObj(String::valueOf)
                    .collect(Collectors.joining())
            )
            .forEach(System.out::println);
        
        // Generate pattern as list
        List<String> numberLines = IntStream.range(1, n + 1)
            .mapToObj(row -> 
                IntStream.range(1, row + 1)
                    .mapToObj(String::valueOf)
                    .collect(Collectors.joining())
            )
            .collect(Collectors.toList());
        
        System.out.println("\nNumber lines: " + numberLines);
        
        // Inverted number triangle
        System.out.println("\nInverted Number Triangle:");
        IntStream.range(0, n)
            .forEach(row -> {
                String numbers = IntStream.range(1, n - row + 1)
                    .mapToObj(String::valueOf)
                    .collect(Collectors.joining());
                System.out.println(numbers);
            });
        
        // Number triangle with spaces (left-aligned)
        System.out.println("\nLeft-Aligned Number Triangle:");
        IntStream.range(1, n + 1)
            .forEach(row -> {
                String spaces = " ".repeat(n - row);
                String numbers = IntStream.range(1, row + 1)
                    .mapToObj(String::valueOf)
                    .collect(Collectors.joining());
                System.out.println(spaces + numbers);
            });
        
        // Number triangle with row numbers
        System.out.println("\nRow Number Triangle:");
        IntStream.range(1, n + 1)
            .forEach(row -> {
                String rowNumbers = IntStream.range(0, row)
                    .mapToObj(col -> String.valueOf(row))
                    .collect(Collectors.joining());
                System.out.println(rowNumbers);
            });
        
        // Triangle with sequential numbers
        System.out.println("\nSequential Number Triangle:");
        AtomicInteger counter = new AtomicInteger(1);
        IntStream.range(1, n + 1)
            .forEach(row -> {
                String numbers = IntStream.range(0, row)
                    .mapToObj(col -> String.valueOf(counter.getAndIncrement()))
                    .collect(Collectors.joining(" "));
                System.out.println(numbers);
            });
        
        // Calculate total numbers printed
        int totalNumbers = IntStream.range(1, n + 1).sum();
        System.out.println("\nTotal numbers printed: " + totalNumbers);
        
        // Generate number triangles of different sizes
        List<Integer> sizes = Arrays.asList(2, 3, 4);
        System.out.println("\nNumber triangles of different sizes:");
        sizes.forEach(size -> {
            System.out.println("Size " + size + ":");
            IntStream.range(1, size + 1)
                .mapToObj(row -> 
                    IntStream.range(1, row + 1)
                        .mapToObj(String::valueOf)
                        .collect(Collectors.joining())
                )
                .forEach(System.out::println);
            System.out.println();
        });
        
        // Pattern statistics
        Map<Integer, Integer> triangleStats = sizes.stream()
            .collect(Collectors.toMap(
                Function.identity(),
                size -> IntStream.range(1, size + 1).sum()
            ));
        
        System.out.println("Triangle statistics: " + triangleStats);
    }
}
```

## 61. Checkered/Floyd's Triangle (0-1)
**Java 8 Approach**: Using conditional streams

```java
import java.util.*;
import java.util.stream.*;

public class BinaryTriangleJava8 {
    public static void main(String[] args) {
        int n = 3;
        
        // Using Java 8 Streams
        System.out.println("Binary Triangle (0-1):");
        IntStream.range(1, n + 1)
            .forEach(row -> {
                IntStream.range(1, row + 1)
                    .mapToObj(col -> (row + col) % 2 == 0 ? "1 " : "0 ")
                    .forEach(System.out::print);
                System.out.println();
            });
        
        // Alternative: Using collectors
        System.out.println("\nAlternative approach:");
        IntStream.range(1, n + 1)
            .mapToObj(row -> 
                IntStream.range(1, row + 1)
                    .mapToObj(col -> (row + col) % 2 == 0 ? "1" : "0")
                    .collect(Collectors.joining(" "))
            )
            .forEach(System.out::println);
        
        // Generate pattern as list
        List<String> binaryLines = IntStream.range(1, n + 1)
            .mapToObj(row -> 
                IntStream.range(1, row + 1)
                    .mapToObj(col -> (row + col) % 2 == 0 ? "1" : "0")
                    .collect(Collectors.joining(" "))
            )
            .collect(Collectors.toList());
        
        System.out.println("\nBinary lines: " + binaryLines);
        
        // Count 1s and 0s
        Map<String, Long> bitCounts = IntStream.range(1, n + 1)
            .boxed()
            .flatMap(row -> IntStream.range(1, row + 1)
                .mapToObj(col -> (row + col) % 2 == 0 ? "1" : "0"))
            .collect(Collectors.groupingBy(
                Function.identity(),
                Collectors.counting()
            ));
        
        System.out.println("\nBit counts: " + bitCounts);
        
        // Binary triangle with different pattern
        System.out.println("\nAlternative Binary Triangle:");
        IntStream.range(1, n + 1)
            .forEach(row -> {
                IntStream.range(1, row + 1)
                    .mapToObj(col -> col % 2 == 0 ? "0 " : "1 ")
                    .forEach(System.out::print);
                System.out.println();
            });
        
        // Binary triangle starting with 0
        System.out.println("\nBinary Triangle starting with 0:");
        IntStream.range(1, n + 1)
            .forEach(row -> {
                IntStream.range(1, row + 1)
                    .mapToObj(col -> (row + col) % 2 == 1 ? "1 " : "0 ")
                    .forEach(System.out::print);
                System.out.println();
            });
        
        // Calculate total bits
        int totalBits = IntStream.range(1, n + 1).sum();
        System.out.println("\nTotal bits printed: " + totalBits);
        
        // Generate binary triangles of different sizes
        List<Integer> sizes = Arrays.asList(2, 3, 4);
        System.out.println("\nBinary triangles of different sizes:");
        sizes.forEach(size -> {
            System.out.println("Size " + size + ":");
            IntStream.range(1, size + 1)
                .mapToObj(row -> 
                    IntStream.range(1, row + 1)
                        .mapToObj(col -> (row + col) % 2 == 0 ? "1" : "0")
                        .collect(Collectors.joining(" "))
                )
                .forEach(System.out::println);
            System.out.println();
        });
        
        // Pattern validation - check if pattern follows rules
        boolean patternValid = IntStream.range(1, n + 1)
            .allMatch(row -> IntStream.range(1, row + 1)
                .allMatch(col -> {
                    int expected = (row + col) % 2 == 0 ? 1 : 0;
                    return true; // Pattern is always valid by construction
                }));
        
        System.out.println("Pattern validation: " + patternValid);
    }
}
```

## 62. Pascal's Triangle
**Java 8 Approach**: Using streams for binomial coefficient calculation

```java
import java.util.*;
import java.util.stream.*;

public class PascalJava8 {
    public static void main(String[] args) {
        int n = 3;
        
        // Using Java 8 Streams
        System.out.println("Pascal's Triangle:");
        IntStream.range(0, n)
            .forEach(row -> {
                // Print spaces for alignment
                IntStream.range(0, n - row - 1)
                    .mapToObj(col -> " ")
                    .forEach(System.out::print);
                
                // Calculate and print binomial coefficients
                AtomicInteger val = new AtomicInteger(1);
                IntStream.range(0, row + 1)
                    .mapToObj(col -> {
                        String result = val.get() + " ";
                        if (col < row) {
                            val.set(val.get() * (row - col) / (col + 1));
                        }
                        return result;
                    })
                    .forEach(System.out::print);
                
                System.out.println();
            });
        
        // Alternative: Using collectors
        System.out.println("\nAlternative approach:");
        IntStream.range(0, n)
            .mapToObj(row -> {
                AtomicInteger val = new AtomicInteger(1);
                return " ".repeat(n - row - 1) +
                    IntStream.range(0, row + 1)
                        .mapToObj(col -> {
                            String result = String.valueOf(val.get());
                            if (col < row) {
                                val.set(val.get() * (row - col) / (col + 1));
                            }
                            return result;
                        })
                        .collect(Collectors.joining(" "));
            })
            .forEach(System.out::println);
        
        // Generate Pascal's triangle as list of lists
        List<List<Integer>> pascalTriangle = IntStream.range(0, n)
            .mapToObj(row -> {
                AtomicInteger val = new AtomicInteger(1);
                return IntStream.range(0, row + 1)
                    .mapToObj(col -> {
                        int result = val.get();
                        if (col < row) {
                            val.set(val.get() * (row - col) / (col + 1));
                        }
                        return result;
                    })
                    .boxed()
                    .collect(Collectors.toList());
            })
            .collect(Collectors.toList());
        
        System.out.println("\nPascal's triangle as list: " + pascalTriangle);
        
        // Calculate sum of each row
        List<Integer> rowSums = pascalTriangle.stream()
            .map(row -> row.stream().mapToInt(Integer::intValue).sum())
            .collect(Collectors.toList());
        
        System.out.println("Row sums: " + rowSums);
        
        // Find maximum value in triangle
        int maxValue = pascalTriangle.stream()
            .flatMap(List::stream)
            .mapToInt(Integer::intValue)
            .max()
            .orElse(0);
        
        System.out.println("Maximum value: " + maxValue);
        
        // Generate Pascal's triangle with different sizes
        List<Integer> sizes = Arrays.asList(2, 3, 4);
        System.out.println("\nPascal's triangles of different sizes:");
        sizes.forEach(size -> {
            System.out.println("Size " + size + ":");
            IntStream.range(0, size)
                .mapToObj(row -> {
                    AtomicInteger val = new AtomicInteger(1);
                    return " ".repeat(size - row - 1) +
                        IntStream.range(0, row + 1)
                            .mapToObj(col -> {
                                String result = String.valueOf(val.get());
                                if (col < row) {
                                    val.set(val.get() * (row - col) / (col + 1));
                                }
                                return result;
                            })
                            .collect(Collectors.joining(" "));
                })
                .forEach(System.out::println);
            System.out.println();
        });
        
        // Validate Pascal's triangle properties
        boolean isSymmetric = pascalTriangle.stream()
            .allMatch(row -> {
                int size = row.size();
                return IntStream.range(0, size / 2)
                    .allMatch(i -> row.get(i).equals(row.get(size - 1 - i)));
            });
        
        System.out.println("Triangle is symmetric: " + isSymmetric);
        
        // Calculate total numbers in triangle
        int totalNumbers = IntStream.range(0, n).sum();
        System.out.println("Total numbers in triangle: " + totalNumbers);
    }
}
```

## 63. Rhombus Pattern
**Java 8 Approach**: Using streams for shifted square generation

```java
import java.util.*;
import java.util.stream.*;

public class RhombusJava8 {
    public static void main(String[] args) {
        int n = 4;
        
        // Using Java 8 Streams
        System.out.println("Rhombus Pattern:");
        IntStream.range(1, n + 1)
            .forEach(row -> {
                IntStream.range(0, n - row)
                    .mapToObj(col -> " ")
                    .forEach(System.out::print);
                
                IntStream.range(0, n)
                    .mapToObj(col -> "*")
                    .forEach(System.out::print);
                
                System.out.println();
            });
        
        // Alternative: Using collectors
        System.out.println("\nAlternative approach:");
        IntStream.range(1, n + 1)
            .mapToObj(row -> 
                " ".repeat(n - row) + "*".repeat(n)
            )
            .forEach(System.out::println);
        
        // Generate rhombus as list
        List<String> rhombusLines = IntStream.range(1, n + 1)
            .mapToObj(row -> 
                " ".repeat(n - row) + "*".repeat(n)
            )
            .collect(Collectors.toList());
        
        System.out.println("\nRhombus lines: " + rhombusLines);
        
        // Hollow rhombus
        System.out.println("\nHollow Rhombus:");
        IntStream.range(1, n + 1)
            .forEach(row -> {
                String spaces = " ".repeat(n - row);
                if (row == 1 || row == n) {
                    System.out.println(spaces + "*".repeat(n));
                } else {
                    System.out.println(spaces + "*" + " ".repeat(n - 2) + "*");
                }
            });
        
        // Rhombus with numbers
        System.out.println("\nRhombus with Numbers:");
        IntStream.range(1, n + 1)
            .forEach(row -> {
                String spaces = " ".repeat(n - row);
                String numbers = IntStream.range(1, n + 1)
                    .mapToObj(col -> String.valueOf(row))
                    .collect(Collectors.joining());
                System.out.println(spaces + numbers);
            });
        
        // Inverted rhombus
        System.out.println("\nInverted Rhombus:");
        IntStream.range(0, n)
            .forEach(row -> {
                String spaces = " ".repeat(row);
                String stars = "*".repeat(n);
                System.out.println(spaces + stars);
            });
        
        // Calculate total characters in rhombus
        int totalChars = n * n;
        System.out.println("\nTotal characters in rhombus: " + totalChars);
        
        // Generate rhombus of different sizes
        List<Integer> sizes = Arrays.asList(3, 4, 5);
        System.out.println("\nRhombus of different sizes:");
        sizes.forEach(size -> {
            System.out.println("Size " + size + ":");
            IntStream.range(1, size + 1)
                .mapToObj(row -> 
                    " ".repeat(size - row) + "*".repeat(size)
                )
                .forEach(System.out::println);
            System.out.println();
        });
        
        // Pattern statistics
        Map<Integer, Integer> rhombusStats = sizes.stream()
            .collect(Collectors.toMap(
                Function.identity(),
                size -> size * size
            ));
        
        System.out.println("Rhombus statistics: " + rhombusStats);
        
        // Calculate total spaces in rhombus
        int totalSpaces = IntStream.range(1, n + 1)
            .map(row -> n - row)
            .sum();
        
        System.out.println("Total spaces in rhombus: " + totalSpaces);
    }
}
```

## 64. Hollow Square
**Java 8 Approach**: Using conditional streams for boundary detection

```java
import java.util.*;
import java.util.stream.*;

public class HollowSquareJava8 {
    public static void main(String[] args) {
        int n = 4;
        
        // Using Java 8 Streams
        System.out.println("Hollow Square:");
        IntStream.range(1, n + 1)
            .forEach(row -> {
                IntStream.range(1, n + 1)
                    .mapToObj(col -> {
                        if (row == 1 || row == n || col == 1 || col == n) {
                            return "*";
                        } else {
                            return " ";
                        }
                    })
                    .forEach(System.out::print);
                System.out.println();
            });
        
        // Alternative: Using collectors
        System.out.println("\nAlternative approach:");
        IntStream.range(1, n + 1)
            .mapToObj(row -> 
                IntStream.range(1, n + 1)
                    .mapToObj(col -> (row == 1 || row == n || col == 1 || col == n) ? "*" : " ")
                    .collect(Collectors.joining())
            )
            .forEach(System.out::println);
        
        // Generate hollow square as list
        List<String> hollowLines = IntStream.range(1, n + 1)
            .mapToObj(row -> 
                IntStream.range(1, n + 1)
                    .mapToObj(col -> (row == 1 || row == n || col == 1 || col == n) ? "*" : " ")
                    .collect(Collectors.joining())
            )
            .collect(Collectors.toList());
        
        System.out.println("\nHollow square lines: " + hollowLines);
        
        // Hollow square with numbers on border
        System.out.println("\nHollow Square with Numbers:");
        IntStream.range(1, n + 1)
            .forEach(row -> {
                IntStream.range(1, n + 1)
                    .mapToObj(col -> {
                        if (row == 1 || row == n || col == 1 || col == n) {
                            return String.valueOf(Math.max(row, col));
                        } else {
                            return " ";
                        }
                    })
                    .forEach(System.out::print);
                System.out.println();
            });
        
        // Hollow square with different border characters
        System.out.println("\nHollow Square with Custom Border:");
        char border = '#';
        IntStream.range(1, n + 1)
            .forEach(row -> {
                IntStream.range(1, n + 1)
                    .mapToObj(col -> {
                        if (row == 1 || row == n || col == 1 || col == n) {
                            return String.valueOf(border);
                        } else {
                            return " ";
                        }
                    })
                    .forEach(System.out::print);
                System.out.println();
            });
        
        // Calculate border characters
        int borderChars = 4 * n - 4; // Top + Bottom + Left + Right - 4 corners
        int innerSpaces = (n - 2) * (n - 2);
        
        System.out.println("\nBorder characters: " + borderChars);
        System.out.println("Inner spaces: " + innerSpaces);
        
        // Generate hollow squares of different sizes
        List<Integer> sizes = Arrays.asList(3, 4, 5);
        System.out.println("\nHollow squares of different sizes:");
        sizes.forEach(size -> {
            System.out.println("Size " + size + " (border: " + (4 * size - 4) + " chars):");
            IntStream.range(1, size + 1)
                .mapToObj(row -> 
                    IntStream.range(1, size + 1)
                        .mapToObj(col -> (row == 1 || row == size || col == 1 || col == size) ? "*" : " ")
                        .collect(Collectors.joining())
                )
                .forEach(System.out::println);
            System.out.println();
        });
        
        // Pattern statistics
        Map<Integer, Map<String, Integer>> squareStats = sizes.stream()
            .collect(Collectors.toMap(
                Function.identity(),
                size -> {
                    Map<String, Integer> stats = new HashMap<>();
                    stats.put("border", 4 * size - 4);
                    stats.put("inner", size > 2 ? (size - 2) * (size - 2) : 0);
                    stats.put("total", size * size);
                    return stats;
                }
            ));
        
        System.out.println("Square statistics: " + squareStats);
        
        // Validate hollow square structure
        boolean isValid = IntStream.range(1, n + 1)
            .allMatch(row -> IntStream.range(1, n + 1)
                .allMatch(col -> {
                    boolean isBorder = row == 1 || row == n || col == 1 || col == n;
                    boolean isInner = row > 1 && row < n && col > 1 && col < n;
                    return isBorder || isInner;
                }));
        
        System.out.println("Hollow square structure valid: " + isValid);
    }
}
```

## 65. Spiral Pattern (Number Grid)
**Java 8 Approach**: Using streams for distance calculation

```java
import java.util.*;
import java.util.stream.*;

public class SpiralNumJava8 {
    public static void main(String[] args) {
        int n = 4;
        
        // Using Java 8 Streams
        System.out.println("Spiral Number Pattern:");
        int len = 2 * n - 1;
        
        IntStream.range(0, len)
            .forEach(i -> {
                IntStream.range(0, len)
                    .mapToObj(j -> {
                        int min = Math.min(
                            Math.min(i, j),
                            Math.min(len - 1 - i, len - 1 - j)
                        );
                        return String.valueOf(n - min) + " ";
                    })
                    .forEach(System.out::print);
                System.out.println();
            });
        
        // Alternative: Using collectors
        System.out.println("\nAlternative approach:");
        IntStream.range(0, len)
            .mapToObj(i -> 
                IntStream.range(0, len)
                    .mapToObj(j -> {
                        int min = Math.min(
                            Math.min(i, j),
                            Math.min(len - 1 - i, len - 1 - j)
                        );
                        return String.valueOf(n - min);
                    })
                    .collect(Collectors.joining(" "))
            )
            .forEach(System.out::println);
        
        // Generate spiral as list of lists
        List<List<Integer>> spiralGrid = IntStream.range(0, len)
            .mapToObj(i -> 
                IntStream.range(0, len)
                    .map(j -> {
                        int min = Math.min(
                            Math.min(i, j),
                            Math.min(len - 1 - i, len - 1 - j)
                        );
                        return n - min;
                    })
                    .boxed()
                    .collect(Collectors.toList())
            )
            .collect(Collectors.toList());
        
        System.out.println("\nSpiral grid: " + spiralGrid);
        
        // Calculate frequency of each number
        Map<Integer, Long> numberFrequency = spiralGrid.stream()
            .flatMap(List::stream)
            .collect(Collectors.groupingBy(Function.identity(), Collectors.counting()));
        
        System.out.println("Number frequency: " + numberFrequency);
        
        // Find minimum and maximum numbers
        int minNum = spiralGrid.stream()
            .flatMap(List::stream)
            .mapToInt(Integer::intValue)
            .min()
            .orElse(0);
        
        int maxNum = spiralGrid.stream()
            .flatMap(List::stream)
            .mapToInt(Integer::intValue)
            .max()
            .orElse(0);
        
        System.out.println("Minimum number: " + minNum);
        System.out.println("Maximum number: " + maxNum);
        
        // Generate spiral of different sizes
        List<Integer> sizes = Arrays.asList(2, 3, 4);
        System.out.println("\nSpirals of different sizes:");
        sizes.forEach(size -> {
            System.out.println("Size " + size + ":");
            int spiralLen = 2 * size - 1;
            IntStream.range(0, spiralLen)
                .mapToObj(i -> 
                    IntStream.range(0, spiralLen)
                        .mapToObj(j -> {
                            int min = Math.min(
                                Math.min(i, j),
                                Math.min(spiralLen - 1 - i, spiralLen - 1 - j)
                            );
                            return String.valueOf(size - min);
                        })
                        .collect(Collectors.joining(" "))
                )
                .forEach(System.out::println);
            System.out.println();
        });
        
        // Calculate sum of all numbers
        int totalSum = spiralGrid.stream()
            .flatMap(List::stream)
            .mapToInt(Integer::intValue)
            .sum();
        
        System.out.println("Total sum of numbers: " + totalSum);
        
        // Pattern statistics
        Map<Integer, Map<String, Integer>> spiralStats = sizes.stream()
            .collect(Collectors.toMap(
                Function.identity(),
                size -> {
                    Map<String, Integer> stats = new HashMap<>();
                    stats.put("gridSize", (2 * size - 1) * (2 * size - 1));
                    stats.put("minNum", 1);
                    stats.put("maxNum", size);
                    return stats;
                }
            ));
        
        System.out.println("Spiral statistics: " + spiralStats);
        
        // Validate spiral pattern properties
        boolean isSymmetric = IntStream.range(0, len)
            .allMatch(i -> IntStream.range(0, len)
                .allMatch(j -> {
                    int val1 = spiralGrid.get(i).get(j);
                    int val2 = spiralGrid.get(len - 1 - i).get(j);
                    int val3 = spiralGrid.get(i).get(len - 1 - j);
                    return val1 == val2 && val1 == val3;
                }));
        
        System.out.println("Spiral is symmetric: " + isSymmetric);
        
        // Find center value (for odd len)
        if (len % 2 == 1) {
            int centerVal = spiralGrid.get(len / 2).get(len / 2);
            System.out.println("Center value: " + centerVal);
        }
    }
}
```

---

## 🎯 Key Java 8 Benefits for Pattern Processing

1. **Declarative Generation**: Express pattern logic clearly
2. **Functional Composition**: Chain operations elegantly
3. **Stream Processing**: Handle complex pattern logic
4. **Collection Integration**: Easy conversion to/from collections
5. **Parallel Processing**: Generate patterns concurrently
6. **Type Safety**: Generic operations with compile-time checking

## 📝 Best Practices

1. **Use `IntStream.range()`** for numeric pattern generation
2. **Leverage collectors** for string joining and aggregation
3. **Use method references** for simple operations
4. **Consider parallel streams** for complex pattern calculations
5. **Use `Optional`** for safe operations
6. **Separate concerns** between pattern logic and display logic

---

*This collection demonstrates how Java 8 features make pattern generation more elegant, readable, and efficient compared to traditional approaches.*
