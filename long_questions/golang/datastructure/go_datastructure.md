🟤 1–27: Strings & Text Processing
1. How are strings represented in memory in Go?
2. What is the difference between string and []byte in Go?
3. Why are strings immutable in Go?
4. What is the most efficient way to concatenate strings?
5. What is strings.Builder and how does it work?
6. How do you iterate over a string containing multibyte characters?
7. What is a rune in Go?
8. What is the difference between len(str) and utf8.RuneCountInString(str)?
9. How does the strings package differ from the bytes package?
10. How do you efficiently replace a substring multiple times?
11. How do you check if a string contains a substring?
12. How do you split a string by whitespace or custom separators?
13. How do you convert a string to an integer?
14. What happens if you cast a negative integer to a string?
15. How do you use Raw String Literals?
16. How to compare two strings in a case-insensitive manner?
17. How do you reverse a string correctly in Go?
18. How do you check if two strings are anagrams?
19. How to check if a string is a palindrome?
20. What is string interning and does Go support it?
21. How to format strings using fmt.Sprintf?
22. How to trim whitespace from a string?
23. How do you convert a struct to a string?
24. What is the difference between Sprint, Sprintln, and Sprintf?
25. How to parse a URL string in Go?
26. How to use regular expressions (regexp) in Go?
27. How do you validate an email format using regex?

⚫ 28–47: Advanced Slices & Built-in Containers
28. How do you delete an element from a slice (at index i)?
29. How do you insert an element at a specific index in a slice?
30. How do you reverse a slice in place?
31. How do you remove duplicate elements from a slice?
32. How do you check if two slices are equal?
33. How do you perform a deep copy of a slice?
34. How does copy() work with overlapping slices?
35. How do you merge two sorted slices?
36. How do you implement a generic slice filter function?
37. Does Go have a built-in Set data structure?
38. How do you implement a Set efficiently in Go?
39. What is the container/list package used for?
40. What is container/ring?
41. How do you use container/heap to implement a Priority Queue?
42. How does Go's sort.Slice work under the hood?
43. How do you sort a slice of custom structs?
44. How do you search in a sorted slice?
45. What is the difference between sort.Sort and sort.Slice?
46. How do you handle slice out of bounds panics?
47. How efficient is appending to a nil slice?

⚪ 48–67: Classic Data Structures & Algorithms
48. How do you implement a Stack (LIFO) in Go?
49. How do you implement a Queue (FIFO) in Go?
50. How do you implement a Linked List in Go?
51. How do you reverse a Linked List?
52. How do you detect a cycle in a Linked List?
53. How do you implement a Binary Search Tree (BST) in Go?
54. How do you implement Tree traversals (Inorder, Preorder, Postorder)?
55. How do you find the max depth of a binary tree?
56. How do you implement a graph using an adjacency list?
57. How do you implement BFS (Breadth-First Search) in Go?
58. How do you implement DFS (Depth-First Search) in Go?
59. How do you implement an LRU Cache in Go?
60. How do you implement a Trie (Prefix Tree)?
61. How do you implement a Hash Map from scratch (conceptually)?
62. How do you handle hash collisions in Go maps?
63. How do you implement a Min-Stack (get min in O(1))?
64. How to find the 'nth' Fibonacci number efficiently?
65. How to validate balanced parentheses in a string?
66. How to find the first non-repeating character in a string?
67. How to implement Merge Sort in Go?

🟡 68–97: Basics of Arrays, Maps, Structs & Loops
68. What is the difference between an Array and a Slice in Go?
69. How do you declare an array of fixed size?
70. What is the zero value of a slice vs. an array?
71. How does the `append` function work internally?
72. What are the components of a slice header?
73. How do you initialize a map with values?
74. What happens if you read from a nil map?
75. What happens if you write to a nil map?
76. How do you check if a key exists in a map?
77. Is the iteration order of a map guaranteed?
78. How do you delete a key from a map?
79. Can you use a slice as a map key? Why or why not?
80. How do you define a struct in Go?
81. What represent anonymous structs?
82. How do you access fields of a struct?
83. What are promoted fields in embedded structs?
84. How to compare two structs?
85. What is the only loop construct in Go?
86. How do you simulate a while-loop in Go?
87. How do you create an infinite loop?
88. How does the `range` keyword work?
89. What are the two values returned by `range` for a slice?
90. What are the two values returned by `range` for a map?
91. How to ignore index or value in a range loop?
92. What is the problem with capturing loop variables in closures?
93. How do you break out of a nested loop with execution labels?
94. What is the `goto` statement and when should you use it?
95. How do you iterate over a channel?
96. What is the difference between `break` and `continue`?
97. How do you loop through a string (byte vs rune)?

🟣 98–117: Time & Date Handling
98. How do you get the current time in Go?
99. What does `time.Time` represent?
100. How do you format a date string in Go (e.g., YYYY-MM-DD)?
101. Why is the reference date "Mon Jan 2 15:04:05 MST 2006"?
102. How do you parse a string into a `time.Time` object?
103. How do you calculate the difference between two times?
104. What is `time.Duration`?
105. How do you add or subtract time from a date?
106. How do you convert a Unix timestamp to `time.Time`?
107. How do you get the Unix timestamp from a `time.Time` object?
108. What is the difference between `time.NewTicker` and `time.NewTimer`?
109. How do you implement a simple timeout using `time.After`?
110. How do you compare if one time is before or after another?
111. How do you handle time zones in Go?
112. How to load a specific location (Timezone)?
113. What is `time.Sleep` and how does it work?
114. How to measure execution time of a function?
115. How strictly does `time.Parse` validate input?
116. How do you reset a timer?
117. How do you serialize `time.Time` to JSON?

🟤 118–147: Methods, Pointers, Interfaces, Channels & Functions
118. What is the difference between passing by value and passing by pointer in Go?
119. How do you define a method on a struct type?
120. Can you define methods on non-struct types (e.g., `type MyInt int`)?
121. What is the difference between a Value Receiver and a Pointer Receiver?
122. When should you use a Pointer Receiver?
123. Can you call a pointer receiver method on a value variable?
124. Can you call a value receiver method on a pointer variable?
125. What is a "Method Set" in Go?
126. What methods belong to the method set of type `T` vs `*T`?
127. How does `new()` differ from `make()` exactly?
128. What types can be created using `make()`?
129. What is the return value of `new(T)`?
130. How are interfaces represented in memory (itab and data)?
131. What is a Type Assertion and how is it used?
132. What is a Type Switch?
133. How do you check if an interface value is `nil`?
134. Can an interface holding a nil concrete pointer be nil?
135. What are the methods required to implement `sort.Interface`?
136. How do you get the capacity (`cap`) and length (`len`) of a channel?
137. What happens if you send to a closed channel?
138. What happens if you receive from a closed channel?
139. How do you check if a channel is closed ensuring no panic?
140. What is the zero value of a function type?
141. Can functions be used as map keys?
142. How do anonymous functions (closures) capture variables?
143. What is a variadic function and how do you pass a slice to it?
144. How does `defer` work with method evaluation (arguments vs execution)?
145. What is `unsafe.Pointer` and when is it used?
146. How does `uintptr` differ from `*int`?
147. How do you manually manage memory alignment (padding)?

🟢 148–167: Generics & Concurrency Data Structures (Sync Package)
148. How do you define a generic Stack data structure in Go?
149. How do you use the `comparable` constraint in generic maps?
150. Can you use a generic type as a receiver for a method?
151. How to implement a generic Linked List?
152. What is `sync.Map` and when should you use it over a regular map?
153. How does `sync.Pool` work and what is it used for?
154. What is the downside of using `sync.Map` for all concurrent map needs?
155. How do you use `sync.WaitGroup` to wait for multiple goroutines?
156. What is `sync.Cond` and how do you use it for signaling?
157. How does `sync.Once` ensure a function is called exactly once?
158. What is the difference between `sync.Mutex` and `sync.RWMutex`?
159. What defines an "Atomic Operation" in Go (`sync/atomic`)?
160. How do you strictly type atomic values using `atomic.Pointer[T]` (Go 1.19+)?
161. How to implement a semaphore using a buffered channel?
162. What is `errgroup.Group` and how does it help with structured concurrency?
163. How to implement a Thread-Safe Queue using Mutex?
164. How to implement a Worker Pool using channels?
165. What is the `context.Context` structure used for?
166. How to implement a generic "Set" using a map [T]struct{}?
167. How to implement a generic "Option Pattern" for struct initialization?
