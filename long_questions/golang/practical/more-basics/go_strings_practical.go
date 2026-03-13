package main

import (
	"bufio"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

// Section 1: String Immutability & Pool (Q1–Q12)

// 1. String is Immutable
func demonstrateStringImmutability() {
	fmt.Println("=== 1. String Immutability ===")
	s := "hello"
	_ = strings.ToUpper(s) // returns a NEW string, doesn't modify s
	fmt.Println(s)
	upper := strings.ToUpper(s)
	fmt.Println(upper)
}

// 2. String Pool — Literals vs new (Go doesn't have explicit string pool like Java)
func demonstrateStringLiterals() {
	fmt.Println("\n=== 2. String Literals ===")
	s1 := "hello"
	s2 := "hello" // same underlying string
	s3 := string([]byte{'h', 'e', 'l', 'l', 'o'}) // creates new string
	fmt.Println(s1 == s2)
	fmt.Println(s1 == s3)
	fmt.Println(s1 == s3) // content comparison
}

// 3. String Concatenation Creates New Object
func demonstrateStringConcatenation() {
	fmt.Println("\n=== 3. String Concatenation ===")
	s := "Hello"
	s2 := s + " World" // creates a new String object
	fmt.Printf("%p\n", &s)
	fmt.Printf("%p\n", &s2)
	fmt.Println(s2)
}

// 4. Go doesn't have intern() like Java, but strings are immutable
func demonstrateStringInterning() {
	fmt.Println("\n=== 4. String Interning (Go equivalent) ===")
	s1 := "hello"
	s2 := "hello"
	fmt.Println(s1 == s2) // Go automatically interns string literals
}

// 5. Compile-Time Constant Folding
func demonstrateCompileTimeFolding() {
	fmt.Println("\n=== 5. Compile-Time Constant Folding ===")
	s1 := "hello" + " " + "world" // compile-time folding
	s2 := "hello world"
	fmt.Println(s1 == s2)
}

// 6. String Comparison — Always use == for content comparison
func demonstrateStringComparison() {
	fmt.Println("\n=== 6. String Comparison ===")
	s := "HELLO"
	fmt.Println(s == "hello")
	fmt.Println(strings.EqualFold(s, "hello")) // case-insensitive comparison
	fmt.Println(strings.Compare("hello", "HELLO")) // case-sensitive comparison
}

// 7. String is a sequence of runes
func demonstrateStringAsRuneSequence() {
	fmt.Println("\n=== 7. String as Rune Sequence ===")
	printCharSequence := func(cs string) { fmt.Println(len(cs)) }
	printCharSequence("hello")               // String implements sequence
	printCharSequence(strings.Builder{})      // Builder implements sequence
}

// 8. String.isEmpty() vs String.isBlank() equivalent
func demonstrateEmptyBlank() {
	fmt.Println("\n=== 8. Empty vs Blank ===")
	empty := ""
	blank := "   " // only whitespace
	fmt.Println(len(empty) == 0)
	fmt.Println(len(blank) == 0)
	fmt.Println(strings.TrimSpace(blank) == "") // equivalent to isBlank()
}

// 9. String Null Safety (Go doesn't have null, but has zero values)
func demonstrateNullSafety() {
	fmt.Println("\n=== 9. Null Safety ===")
	var s string // zero value is "", not null
	fmt.Println("Value:", s)    // empty string concatenation is safe
	fmt.Println(len(s))         // length of empty string is 0
}

// 10. String.valueOf() equivalent
func demonstrateStringValueOf() {
	fmt.Println("\n=== 10. String ValueOf ===")
	var obj interface{} = nil
	fmt.Println(fmt.Sprintf("%v", obj)) // safe nil → "<nil>"
	// obj.(string) would panic if nil
}

// 11. Character Case Methods
func demonstrateCharacterCase() {
	fmt.Println("\n=== 11. Character Case Methods ===")
	c := 'a'
	fmt.Println(unicode.IsLetter(c))
	fmt.Println(unicode.IsDigit(c))
	fmt.Println(unicode.ToUpper(c))
	fmt.Println(unicode.IsSpace(' '))
}

// 12. String Constant Pool — Impact on Memory
func demonstrateStringPoolImpact() {
	fmt.Println("\n=== 12. String Pool Impact ===")
	a := "abc"
	b := "ab" + "c"   // compile-time folding → same underlying string
	c := "ab"
	d := c + "c"      // runtime concatenation → new string
	fmt.Println(a == b)
	fmt.Println(a == d)
	fmt.Println(a == d) // content comparison
}

// Section 2: String Methods (Q13–Q28)

// 13. substring() — Indices
func demonstrateSubstring() {
	fmt.Println("\n=== 13. Substring ===")
	s := "Hello World"
	fmt.Println(s[6:])      // from index 6 to end
	fmt.Println(s[0:5])     // from 0, exclusive end 5
	fmt.Println(s[6:11])    // "World"
}

// 14. indexOf and lastIndexOf
func demonstrateIndexOf() {
	fmt.Println("\n=== 14. IndexOf ===")
	s := "abcabc"
	fmt.Println(strings.Index(s, "c"))      // first 'c'
	fmt.Println(strings.LastIndex(s, "c"))  // last 'c'
	fmt.Println(strings.Index(s[1:], "abc") + 1) // search from index 1
	fmt.Println(strings.Index(s, "xyz"))    // not found
}

// 15. replace() vs replaceAll()
func demonstrateReplace() {
	fmt.Println("\n=== 15. Replace ===")
	s := "aaa.bbb.ccc"
	fmt.Println(strings.ReplaceAll(s, ".", "-"))       // literal replace
	fmt.Println(strings.ReplaceAll(s, ".", "!"))       // literal string replace
	re := regexp.MustCompile(`\.`)
	fmt.Println(re.ReplaceAllString(s, "#"))          // regex replace
	re2 := regexp.MustCompile(`[0-9]`)
	fmt.Println(re2.ReplaceAllString("abc123", "*"))  // regex
}

// 16. split()
func demonstrateSplit() {
	fmt.Println("\n=== 16. Split ===")
	csv := "a,b,c,d,"
	parts := strings.Split(csv, ",")
	fmt.Println(len(parts))
	// Go's Split keeps trailing empty strings
	parts2 := strings.Split(csv, ",")
	fmt.Println(len(parts2))
}

// 17. join()
func demonstrateJoin() {
	fmt.Println("\n=== 17. Join ===")
	result := strings.Join([]string{"apple", "banana", "cherry"}, ", ")
	fmt.Println(result)
	
	list := []string{"x", "y", "z"}
	fmt.Println(strings.Join(list, "-"))
}

// 18. trim() vs strip()
func demonstrateTrimStrip() {
	fmt.Println("\n=== 18. Trim vs Strip ===")
	s := "  hello  "
	fmt.Println("[" + strings.TrimSpace(s) + "]")
	fmt.Println("[" + strings.TrimSpace(s) + "]")      // Go's TrimSpace handles Unicode
	fmt.Println("[" + strings.TrimLeft(s, " ") + "]")
	fmt.Println("[" + strings.TrimRight(s, " ") + "]")
}

// 19. startsWith() and endsWith()
func demonstrateStartsEnds() {
	fmt.Println("\n=== 19. Starts/Ends With ===")
	url := "https://www.example.com"
	fmt.Println(strings.HasPrefix(url, "https"))
	fmt.Println(strings.HasSuffix(url, ".com"))
	fmt.Println(strings.HasPrefix(url[8:], "www")) // start checking at index 8
}

// 20. contains() and matches()
func demonstrateContainsMatches() {
	fmt.Println("\n=== 20. Contains and Matches ===")
	email := "user@example.com"
	fmt.Println(strings.Contains(email, "@"))
	re := regexp.MustCompile(`^[a-z]+@[a-z]+\.[a-z]+$`)
	fmt.Println(re.MatchString(email)) // full match
	re2 := regexp.MustCompile(`^[a-z]+$`)
	fmt.Println(re2.MatchString("abc123")) // must match ENTIRE string
}

// 21. chars() Stream equivalent
func demonstrateCharsStream() {
	fmt.Println("\n=== 21. Chars Stream ===")
	count := 0
	for _, r := range "Hello World" {
		if unicode.IsUpper(r) {
			count++
		}
	}
	fmt.Println(count)
}

// 22. String.format() equivalent
func demonstrateStringFormat() {
	fmt.Println("\n=== 22. String Format ===")
	name := "Alice"
	age := 30
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Pi = %.4f\n", 3.1415926535)
}

// 23. String to rune Array and Back
func demonstrateStringToRuneArray() {
	fmt.Println("\n=== 23. String to Rune Array ===")
	s := "hello"
	runes := []rune(s)
	runes[0] = 'H' // modify the rune array
	fmt.Println(s)             // original unchanged
	fmt.Println(string(runes)) // new string from runes
}

// 24. String.valueOf() for Primitives
func demonstrateStringValueOfPrimitives() {
	fmt.Println("\n=== 24. String ValueOf Primitives ===")
	fmt.Println(fmt.Sprintf("%v", 42))
	fmt.Println(fmt.Sprintf("%v", 3.14))
	fmt.Println(fmt.Sprintf("%v", true))
	fmt.Println(fmt.Sprintf("%v", []rune{'h', 'i'}))
}

// 25. repeat()
func demonstrateRepeat() {
	fmt.Println("\n=== 25. Repeat ===")
	fmt.Println(strings.Repeat("ab", 3))
	fmt.Println(strings.Repeat("-", 10))
}

// 26. lines() equivalent
func demonstrateLines() {
	fmt.Println("\n=== 26. Lines ===")
	text := "line1\nline2\nline3"
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		fmt.Println(line)
	}
	fmt.Println(len(lines))
}

// 27. Text Blocks equivalent (Go has raw string literals)
func demonstrateTextBlocks() {
	fmt.Println("\n=== 27. Text Blocks ===")
	json := `{
		"name": "Alice",
		"age": 30
	}`
	fmt.Println(strings.TrimSpace(json))
}

// 28. String Interning — When to Use
func demonstrateStringInterningUsage() {
	fmt.Println("\n=== 28. String Interning Usage ===")
	// Go automatically interns string literals
	s1 := "common"
	s2 := "common"
	fmt.Println(s1 == s2) // same underlying string
}

// Section 3: StringBuilder & StringBuffer (Q29–Q42)

// 29. StringBuilder — Mutable String
func demonstrateStringBuilder() {
	fmt.Println("\n=== 29. StringBuilder ===")
	var sb strings.Builder
	sb.WriteString("Hello")
	sb.WriteString(" World")
	sb.WriteString("!")
	fmt.Println(sb.String())
	fmt.Println(sb.Len())
}

// 30. StringBuilder vs Concatenation in Loop
func demonstrateStringBuilderVsConcat() {
	fmt.Println("\n=== 30. StringBuilder vs Concatenation ===")
	// SLOW: creates a new String object on each iteration
	result := ""
	for i := 0; i < 5; i++ {
		result += strconv.Itoa(i)
	}
	fmt.Println(result)

	// FAST: O(n) amortized, single mutable buffer
	var sb strings.Builder
	for i := 0; i < 5; i++ {
		sb.WriteString(strconv.Itoa(i))
	}
	fmt.Println(sb.String())
}

// 31. StringBuilder reverse()
func demonstrateStringBuilderReverse() {
	fmt.Println("\n=== 31. StringBuilder Reverse ===")
	var sb strings.Builder
	sb.WriteString("hello")
	reversed := reverseString(sb.String())
	fmt.Println(reversed)
}

// Helper function to reverse string
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 32. StringBuilder insert() and delete()
func demonstrateStringBuilderInsertDelete() {
	fmt.Println("\n=== 32. StringBuilder Insert/Delete ===")
	var sb strings.Builder
	sb.WriteString("Hello World")
	
	// Insert at index 5
	s := sb.String()
	s = s[:5] + "," + s[5:]
	sb.Reset()
	sb.WriteString(s)
	fmt.Println(sb.String())
	
	// Delete from 0 to 6
	s = sb.String()
	s = s[6:]
	sb.Reset()
	sb.WriteString(s)
	fmt.Println(sb.String())
	
	// Delete char at index 0
	s = sb.String()
	s = s[1:]
	sb.Reset()
	sb.WriteString(s)
	fmt.Println(sb.String())
}

// 33. StringBuilder indexOf and replace
func demonstrateStringBuilderIndexOfReplace() {
	fmt.Println("\n=== 33. StringBuilder IndexOf/Replace ===")
	var sb strings.Builder
	sb.WriteString("Hello World")
	
	s := sb.String()
	idx := strings.Index(s, "World")
	s = s[:idx] + "Java" + s[idx+5:]
	sb.Reset()
	sb.WriteString(s)
	fmt.Println(sb.String())
}

// 34. StringBuffer vs StringBuilder (Go only has strings.Builder)
func demonstrateStringBuilderOnly() {
	fmt.Println("\n=== 34. StringBuilder Only ===")
	// Go only has strings.Builder, which is not thread-safe
	var sb strings.Builder
	sb.WriteString("hello ")
	sb.WriteString("world")
	
	var sb2 strings.Builder
	sb2.WriteString("hello ")
	sb2.WriteString("world")
	
	fmt.Println(sb.String() == sb2.String())
}

// 35. StringJoiner equivalent
func demonstrateStringJoiner() {
	fmt.Println("\n=== 35. StringJoiner ===")
	elements := []string{"apple", "banana", "cherry"}
	result := "[" + strings.Join(elements, ", ") + "]"
	fmt.Println(result)
	fmt.Println(len(result))
}

// 36. StringJoiner with Empty Check
func demonstrateStringJoinerEmpty() {
	fmt.Println("\n=== 36. StringJoiner Empty ===")
	var elements []string
	if len(elements) == 0 {
		fmt.Println("EMPTY")
	} else {
		fmt.Println("[" + strings.Join(elements, ", ") + "]")
	}
	
	elements = []string{"x"}
	fmt.Println("[" + strings.Join(elements, ", ") + "]")
}

// 37. StringBuilder — Chaining Methods
func demonstrateStringBuilderChaining() {
	fmt.Println("\n=== 37. StringBuilder Chaining ===")
	var sb strings.Builder
	sb.WriteString("Hello")
	sb.WriteString(", ")
	sb.WriteString("World")
	sb.WriteString("!")
	result := reverseString(sb.String())
	fmt.Println(result)
}

// 38. charAt() and setCharAt()
func demonstrateCharAtSetCharAt() {
	fmt.Println("\n=== 38. CharAt/SetCharAt ===")
	var sb strings.Builder
	sb.WriteString("Hello")
	runes := []rune(sb.String())
	fmt.Println(string(runes[1]))
	runes[0] = 'h'
	fmt.Println(string(runes))
}

// 39. StringBuilder capacity vs length
func demonstrateStringBuilderCapacity() {
	fmt.Println("\n=== 39. StringBuilder Capacity ===")
	var sb strings.Builder
	fmt.Println(sb.Cap()) // initial capacity
	sb.WriteString("hello")
	fmt.Println(sb.Len()) // actual characters
	fmt.Println(sb.Cap()) // still has capacity
}

// 40. Palindrome Check with StringBuilder
func demonstratePalindromeCheck() {
	fmt.Println("\n=== 40. Palindrome Check ===")
	isPalindrome := func(s string) bool {
		clean := regexp.MustCompile(`[^a-zA-Z0-9]`).ReplaceAllString(strings.ToLower(s), "")
		reversed := reverseString(clean)
		return clean == reversed
	}
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
}

// 41. String.format With Various Formats
func demonstrateStringFormatVarious() {
	fmt.Println("\n=== 41. String Format Various ===")
	fmt.Printf("%-10s | %5d | %.2f\n", "Alice", 30, 95.678)
	fmt.Printf("%-10s | %5d | %.2f\n", "Bob", 25, 87.5)
}

// 42. StringBuilder as Char Stack — Interview Pattern
func demonstrateStringBuilderAsStack() {
	fmt.Println("\n=== 42. StringBuilder as Stack ===")
	removeDuplicates := func(s string) string {
		var sb strings.Builder
		for _, c := range s {
			if sb.Len() > 0 {
				last := sb.String()[sb.Len()-1]
				if last == byte(c) {
					// pop
					current := sb.String()
					sb.Reset()
					sb.WriteString(current[:sb.Len()-1])
					continue
				}
			}
			sb.WriteRune(c) // push
		}
		return sb.String()
	}
	fmt.Println(removeDuplicates("abbaca"))
	fmt.Println(removeDuplicates("azxxzy"))
}

// Section 4: String Formatting & Parsing (Q43–Q52)

// 43. Integer.parseInt Edge Cases
func demonstrateIntegerParseInt() {
	fmt.Println("\n=== 43. Integer ParseInt ===")
	fmt.Println(strconv.Atoi("42"))
	fmt.Println(strconv.Atoi("-100"))
	fmt.Println(strconv.ParseInt("FF", 16, 64))
	fmt.Println(strconv.ParseInt("1010", 2, 64))
	_, err := strconv.Atoi("abc")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// 44. Double.parseDouble and Number Formatting
func demonstrateDoubleParse() {
	fmt.Println("\n=== 44. Double Parse ===")
	d, _ := strconv.ParseFloat("3.14159", 64)
	fmt.Printf("%.2f\n", d)
	fmt.Printf("%e\n", d)   // scientific notation
	fmt.Printf("%10.3f\n", d) // width 10, 3 decimals
}

// 45. String to boolean
func demonstrateStringToBool() {
	fmt.Println("\n=== 45. String to Bool ===")
	fmt.Println(strings.ToLower("true") == "true")
	fmt.Println(strings.ToLower("TRUE") == "true")
	fmt.Println(strings.ToLower("yes") == "true")  // false!
	fmt.Println(strings.ToLower("1") == "true")    // false!
}

// 46. String.format() with Padding
func demonstrateStringFormatPadding() {
	fmt.Println("\n=== 46. String Format Padding ===")
	fmt.Printf("%05d\n", 42)   // zero-padded
	fmt.Printf("%+d\n", 42)    // always show sign
	fmt.Printf("%x\n", 255)    // hex lowercase
	fmt.Printf("%X\n", 255)    // hex uppercase
	fmt.Printf("%o\n", 8)      // octal
}

// 47. Regex — Pattern and Matcher
func demonstrateRegexPatternMatcher() {
	fmt.Println("\n=== 47. Regex Pattern/Matcher ===")
	email := "user123@example.com"
	pattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	fmt.Println(pattern.MatchString(email))

	text := "Call 123-456-7890 or 987-654-3210"
	phones := regexp.MustCompile(`\d{3}-\d{3}-\d{4}`)
	fmt.Println(phones.FindAllString(text, -1))
}

// 48. Regex Groups
func demonstrateRegexGroups() {
	fmt.Println("\n=== 48. Regex Groups ===")
	pattern := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
	matches := pattern.FindStringSubmatch("Date: 2024-01-15")
	if len(matches) > 0 {
		fmt.Println("Year:", matches[1])
		fmt.Println("Month:", matches[2])
		fmt.Println("Day:", matches[3])
	}
}

// 49. String.chars() to Count Characters
func demonstrateCharsCount() {
	fmt.Println("\n=== 49. Chars Count ===")
	s := "hello world"
	freq := make(map[rune]int)
	for _, c := range s {
		if c != ' ' {
			freq[c]++
		}
	}
	
	// Sort keys for consistent output
	var keys []rune
	for k := range freq {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	
	for _, k := range keys {
		fmt.Printf("%c=%d ", k, freq[k])
	}
	fmt.Println()
}

// 50. Scanner for Parsing
func demonstrateScannerParsing() {
	fmt.Println("\n=== 50. Scanner Parsing ===")
	scanner := bufio.NewScanner(strings.NewReader("10 3.14 hello"))
	scanner.Split(bufio.ScanWords)
	
	i, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	d, _ := strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	s := scanner.Text()
	
	fmt.Println(i, d, s)
}

// 51. String.chars() — Anagram Check
func demonstrateAnagramCheck() {
	fmt.Println("\n=== 51. Anagram Check ===")
	isAnagram := func(s1, s2 string) bool {
		a := []rune(s1)
		b := []rune(s2)
		sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		return string(a) == string(b)
	}
	fmt.Println(isAnagram("listen", "silent"))
	fmt.Println(isAnagram("hello", "world"))
}

// 52. toUpperCase / toLowerCase with Locale
func demonstrateUpperCaseLocale() {
	fmt.Println("\n=== 52. UpperCase Locale ===")
	s := "istanbul"
	// Go's strings package handles Unicode correctly
	fmt.Println(strings.ToUpper(s))
	// Turkish locale handling would require external libraries
	fmt.Println(strings.ToUpper(s))
}

// Section 5: String Common Interview Patterns (Q53–Q60)

// 53. Find First Non-Repeating Character
func demonstrateFirstUnique() {
	fmt.Println("\n=== 53. First Non-Repeating Character ===")
	firstUnique := func(s string) rune {
		freq := make(map[rune]int)
		for _, c := range s {
			freq[c]++
		}
		for _, c := range s {
			if freq[c] == 1 {
				return c
			}
		}
		return '_'
	}
	fmt.Println(firstUnique("leetcode"))
	fmt.Println(firstUnique("aabb"))
}

// 54. Reverse Words in a String
func demonstrateReverseWords() {
	fmt.Println("\n=== 54. Reverse Words ===")
	reverseWords := func(s string) string {
		words := strings.Fields(strings.TrimSpace(s)) // splits on whitespace
		for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
			words[i], words[j] = words[j], words[i]
		}
		return strings.Join(words, " ")
	}
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  hello world  "))
}

// 55. Count Vowels
func demonstrateCountVowels() {
	fmt.Println("\n=== 55. Count Vowels ===")
	countVowels := func(s string) int {
		vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
			'A': true, 'E': true, 'I': true, 'O': true, 'U': true}
		count := 0
		for _, c := range s {
			if vowels[c] {
				count++
			}
		}
		return count
	}
	fmt.Println(countVowels("Hello World"))
	fmt.Println(countVowels("aeiou"))
}

// 56. Longest Common Prefix
func demonstrateLongestCommonPrefix() {
	fmt.Println("\n=== 56. Longest Common Prefix ===")
	longestCommonPrefix := func(strs []string) string {
		if len(strs) == 0 {
			return ""
		}
		prefix := strs[0]
		for _, s := range strs {
			for !strings.HasPrefix(s, prefix) {
				if len(prefix) == 0 {
					return ""
				}
				prefix = prefix[:len(prefix)-1]
			}
		}
		return prefix
	}
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

// 57. String Compression
func demonstrateStringCompression() {
	fmt.Println("\n=== 57. String Compression ===")
	compress := func(s string) string {
		var sb strings.Builder
		i := 0
		for i < len(s) {
			c := s[i]
			count := 0
			for i < len(s) && s[i] == c {
				i++
				count++
			}
			sb.WriteByte(c)
			sb.WriteString(strconv.Itoa(count))
		}
		compressed := sb.String()
		if len(compressed) < len(s) {
			return compressed
		}
		return s
	}
	fmt.Println(compress("aabcccccaaa"))
	fmt.Println(compress("abc")) // no compression benefit
}

// 58. Check if String is Numeric
func demonstrateIsNumeric() {
	fmt.Println("\n=== 58. Is Numeric ===")
	isNumeric := func(s string) bool {
		if len(s) == 0 {
			return false
		}
		for _, c := range s {
			if !unicode.IsDigit(c) {
				return false
			}
		}
		return true
	}
	fmt.Println(isNumeric("12345"))
	fmt.Println(isNumeric("123.45"))
	fmt.Println(isNumeric("123abc"))
}

// 59. Roman to Integer (Common Interview Q)
func demonstrateRomanToInt() {
	fmt.Println("\n=== 59. Roman to Integer ===")
	romanToInt := func(s string) int {
		romanMap := map[rune]int{
			'I': 1, 'V': 5, 'X': 10, 'L': 50,
			'C': 100, 'D': 500, 'M': 1000,
		}
		result := 0
		for i := 0; i < len(s); i++ {
			curr := romanMap[rune(s[i])]
			next := 0
			if i+1 < len(s) {
				next = romanMap[rune(s[i+1])]
			}
			if curr < next {
				result -= curr
			} else {
				result += curr
			}
		}
		return result
	}
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

// 60. Valid Parentheses (Stack + String)
func demonstrateValidParentheses() {
	fmt.Println("\n=== 60. Valid Parentheses ===")
	isValid := func(s string) bool {
		stack := []rune{}
		pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
		
		for _, c := range s {
			if c == '(' || c == '[' || c == '{' {
				stack = append(stack, c)
			} else {
				if len(stack) == 0 {
					return false
				}
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if pairs[c] != top {
					return false
				}
			}
		}
		return len(stack) == 0
	}
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("{[]}"))
}

func main() {
	// Section 1: String Immutability & Pool (Q1–Q12)
	demonstrateStringImmutability()
	demonstrateStringLiterals()
	demonstrateStringConcatenation()
	demonstrateStringInterning()
	demonstrateCompileTimeFolding()
	demonstrateStringComparison()
	demonstrateStringAsRuneSequence()
	demonstrateEmptyBlank()
	demonstrateNullSafety()
	demonstrateStringValueOf()
	demonstrateCharacterCase()
	demonstrateStringPoolImpact()

	// Section 2: String Methods (Q13–Q28)
	demonstrateSubstring()
	demonstrateIndexOf()
	demonstrateReplace()
	demonstrateSplit()
	demonstrateJoin()
	demonstrateTrimStrip()
	demonstrateStartsEnds()
	demonstrateContainsMatches()
	demonstrateCharsStream()
	demonstrateStringFormat()
	demonstrateStringToRuneArray()
	demonstrateStringValueOfPrimitives()
	demonstrateRepeat()
	demonstrateLines()
	demonstrateTextBlocks()
	demonstrateStringInterningUsage()

	// Section 3: StringBuilder & StringBuffer (Q29–Q42)
	demonstrateStringBuilder()
	demonstrateStringBuilderVsConcat()
	demonstrateStringBuilderReverse()
	demonstrateStringBuilderInsertDelete()
	demonstrateStringBuilderIndexOfReplace()
	demonstrateStringBuilderOnly()
	demonstrateStringJoiner()
	demonstrateStringJoinerEmpty()
	demonstrateStringBuilderChaining()
	demonstrateCharAtSetCharAt()
	demonstrateStringBuilderCapacity()
	demonstratePalindromeCheck()
	demonstrateStringFormatVarious()
	demonstrateStringBuilderAsStack()

	// Section 4: String Formatting & Parsing (Q43–Q52)
	demonstrateIntegerParseInt()
	demonstrateDoubleParse()
	demonstrateStringToBool()
	demonstrateStringFormatPadding()
	demonstrateRegexPatternMatcher()
	demonstrateRegexGroups()
	demonstrateCharsCount()
	demonstrateScannerParsing()
	demonstrateAnagramCheck()
	demonstrateUpperCaseLocale()

	// Section 5: String Common Interview Patterns (Q53–Q60)
	demonstrateFirstUnique()
	demonstrateReverseWords()
	demonstrateCountVowels()
	demonstrateLongestCommonPrefix()
	demonstrateStringCompression()
	demonstrateIsNumeric()
	demonstrateRomanToInt()
	demonstrateValidParentheses()
}
