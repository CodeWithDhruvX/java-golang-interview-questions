# Mini-Project 2: Student Management System (File I/O)

**Goal**: Demonstrate File Handling (Reading/Writing data), JSON Encoding/Decoding, and basic CRUD operations.

## Features
1.  Add Student
2.  Display All Students (Read from file)
3.  Save Data Persistence (Data survives app restart)

## Code Implementation

```go
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// Student Struct (JSON tags for serialization)
type Student struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Marks float64 `json:"marks"`
}

func NewStudent(id int, name string, marks float64) *Student {
	return &Student{
		ID:    id,
		Name:  name,
		Marks: marks,
	}
}

func (s *Student) String() string {
	return fmt.Sprintf("ID: %d, Name: %s, Marks: %.2f", s.ID, s.Name, s.Marks)
}

const FileName = "students.json"

func main() {
	students := loadData()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n1. Add Student\n2. Show All\n3. Save & Exit")
		fmt.Print("Choice: ")
		
		if !scanner.Scan() {
			break
		}
		
		choice, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid input!")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("ID: ")
			scanner.Scan()
			id, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Invalid ID!")
				continue
			}

			fmt.Print("Name: ")
			scanner.Scan()
			name := scanner.Text()

			fmt.Print("Marks: ")
			scanner.Scan()
			marks, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				fmt.Println("Invalid marks!")
				continue
			}

			// Check for duplicate ID
			duplicate := false
			for _, student := range students {
				if student.ID == id {
					duplicate = true
					break
				}
			}

			if duplicate {
				fmt.Println("Student with this ID already exists!")
				continue
			}

			students = append(students, NewStudent(id, name, marks))
			fmt.Println("Student added successfully!")

		case 2:
			if len(students) == 0 {
				fmt.Println("No records found.")
			} else {
				fmt.Println("Student Records:")
				for _, student := range students {
					fmt.Println(student)
				}
			}

		case 3:
			err := saveData(students)
			if err != nil {
				fmt.Printf("Error saving data: %v\n", err)
			} else {
				fmt.Println("Data saved. Exiting...")
			}
			return

		default:
			fmt.Println("Invalid choice!")
		}
	}
}

// Save students to JSON file
func saveData(students []*Student) error {
	data, err := json.MarshalIndent(students, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal students: %w", err)
	}

	err = os.WriteFile(FileName, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// Load students from JSON file
func loadData() []*Student {
	var students []*Student

	data, err := os.ReadFile(FileName)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist, return empty slice
			return students
		}
		fmt.Printf("Error reading file: %v\n", err)
		return students
	}

	err = json.Unmarshal(data, &students)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return []*Student{}
	}

	return students
}
```

## Key Code Concepts Used
*   **JSON Serialization**: Struct tags (`json:"field"`) for automatic JSON encoding/decoding.
*   **File I/O**: `os.ReadFile` and `os.WriteFile` for file operations.
*   **Slices**: `[]*Student` to hold data in memory.
*   **Error Handling**: Proper error checking and wrapping with `fmt.Errorf`.
*   **Persistence**: Data remains available even after terminating the program.

---

## 📋 Interview Questions

### **Serialization & File I/O Questions**

**Q1: Why use JSON instead of Go's gob encoding for serialization?**
**A**: "I chose JSON because it's human-readable, language-agnostic, and widely supported. JSON files can be easily inspected, edited manually, or used by other programming languages. While gob encoding is more efficient and Go-specific, JSON provides better interoperability and debugging capabilities. For a learning example, JSON's readability makes it easier to understand what's being stored and helps with troubleshooting data issues."

**Q2: What is the purpose of JSON struct tags like `json:"id"`?**
**A**: "JSON struct tags control how the struct fields are encoded and decoded in JSON. The tag `json:"id"` tells the JSON encoder to use "id" as the field name in the JSON, even if the Go struct field is named "ID" (uppercase). This allows for different naming conventions between Go (which prefers PascalCase) and JSON (which typically uses camelCase or snake_case). It also provides flexibility to change Go field names without affecting the JSON format."

**Q3: Why use `json.MarshalIndent` instead of `json.Marshal`?**
**A**: "I used `MarshalIndent` because it produces pretty-printed JSON with proper indentation, making the output file human-readable. This is helpful for debugging and manual inspection of the data. `Marshal` would produce compact JSON without formatting, which is more efficient for transmission but harder to read. For a learning application where humans might look at the file, the indented format is more educational and user-friendly."

**Q4: What happens if the JSON file is corrupted or contains invalid data?**
**A**: "If the file is corrupted or contains invalid JSON, the `json.Unmarshal` call will return an error. In the current code, I catch this error and print an error message, then return an empty slice. This graceful degradation allows the program to continue running even if the data file is corrupted. In a production system, I might implement backup mechanisms - like keeping multiple backup files or attempting to recover partial data from valid portions of the JSON."

**Q5: Why use `os.WriteFile` instead of manually creating and writing to a file?**
**A**: "`os.WriteFile` is a convenience function that handles file creation, writing, and closing in a single operation. It's safer because it automatically handles the file closing, even if an error occurs during writing. This prevents resource leaks and is much cleaner than manually managing file handles with `os.Create` and `defer file.Close()`. It's the recommended approach for simple file writing operations."

### **Data Structure & Design Questions**

**Q6: Why use `[]*Student` (slice of pointers) instead of `[]Student` (slice of values)?**
**A**: "I used a slice of pointers because it's more memory-efficient when dealing with large collections. When you have a slice of structs, each operation like appending or passing elements copies the entire struct. With pointers, you only copy the pointer (8 bytes on 64-bit systems), regardless of how large the struct is. This is especially important if the Student struct grows to include more fields. Pointers also allow for easier modification of individual students without needing to reassign them."

**Q7: What's the risk of using a global slice for student storage?**
**A**: "Using a slice that's recreated on each program startup from the file is fine for this simple console application. But in a larger application with concurrent access, this could cause race conditions. Multiple goroutines trying to modify the slice simultaneously could lead to data corruption. For a multi-user system, I'd need proper synchronization using mutexes or move to a database that handles concurrent access safely."

**Q8: Why are the Student fields public instead of private?**
**A**: "In this simple example, I made the fields public (uppercase) to allow JSON encoding/decoding to work. The `encoding/json` package can only access public fields. For better encapsulation, I could make the fields private and implement the `Marshaler` and `Unmarshaler` interfaces to control serialization. However, for this learning example, public fields keep the code simpler and more readable while still demonstrating the core concepts."

### **Error Handling & Edge Cases Questions**

**Q9: How does Go's error handling differ from Java's exception handling?**
**A**: "Go uses explicit error returns instead of exceptions. Functions that can fail return an error as their last return value, and callers must check this error. This approach makes error handling more visible and explicit in the code. Unlike Java's try-catch blocks where errors can be caught far from where they occur, Go's approach forces you to handle errors immediately where they happen, leading to more robust code."

**Q10: What's the purpose of `fmt.Errorf` with `%w` verb?**
**A**: "`fmt.Errorf` with the `%w` verb creates wrapped errors that preserve the original error information while adding context. This allows for error chaining and better debugging. When I use `fmt.Errorf("failed to write file: %w", err)`, the new error contains both my custom message and the original error. This is Go's equivalent of exception chaining in Java, allowing you to trace the full error path while providing meaningful context at each level."

**Q11: How would you handle duplicate student IDs more efficiently?**
**A**: "Currently, I check for duplicates by iterating through the entire slice, which is O(n). For better performance with many students, I'd maintain a separate map[int]bool to track used IDs. This would give O(1) duplicate checking. I could also implement a unique ID generator that automatically assigns IDs, eliminating the possibility of duplicates entirely. For database-backed systems, I'd use database constraints to enforce uniqueness at the storage level."

**Q12: What happens if the program crashes while writing to the file?**
**A**: "If the program crashes while writing, the file could be left in a corrupted state. To handle this, I'd implement a safer approach: write to a temporary file first, then rename it to the target filename. This atomic operation ensures that either the old file remains intact or the new file is completely written. I could also implement a backup mechanism where I keep the previous version as a backup until the new write is successful."

### **Performance & Scalability Questions**

**Q13: How would this perform with 100,000 students?**
**A**: "With 100,000 students, the current approach might face performance challenges. The slice would consume significant memory, and loading/saving the entire collection at once could be slow. For large datasets, I'd consider pagination - loading only subsets of data at a time, or moving to a database with proper indexing. The JSON approach also creates large files that might be slow to parse. I might switch to a more efficient binary format or use streaming JSON for large files."

**Q14: Why not use a database instead of JSON file storage?**
**A**: "For this simple learning example, JSON file storage is sufficient and doesn't require external dependencies. But for a production system, a database would be better because it offers concurrent access, query capabilities, data integrity constraints, and better scalability. Database also provides backup and recovery tools, and can handle much larger datasets efficiently. The JSON file approach is good for learning and small applications but doesn't scale well."

**Q15: How would you implement search functionality efficiently?**
**A**: "For efficient searching, I'd maintain index maps. For example, `map[int]*Student` for ID-based lookups and `map[string][]*Student` for name-based searches. This would give O(1) and O(k) search times respectively, where k is the number of matches. For complex queries, I'd move to a database with proper indexing. I could also implement a simple indexing system using sorted slices and binary search for range queries."

### **Security & Data Integrity Questions**

**Q16: What are the security risks of JSON serialization?**
**A**: "JSON serialization is generally safer than Java's serialization, but still has risks. The JSON files could contain sensitive data in plain text. There's no built-in encryption, so anyone with file access can read student data. There's also risk of JSON injection if I were to parse untrusted JSON. For production systems, I'd implement file encryption, access controls, and input validation when parsing external JSON data."

**Q17: How would you add data validation for student marks?**
**A**: "I'd add validation in the `NewStudent` constructor or in the input handling code. For example, check that marks are between 0 and 100, and return an error if they're not. I could also add validation for the name field to ensure it's not empty and contains only valid characters. This prevents invalid data from entering the system and maintains data integrity. In Go, I'd typically return errors from validation functions rather than throwing exceptions."

**Q18: How would you implement data backup and recovery?**
**A**: "I'd implement automatic backup by creating timestamped backup files before saving new data. For recovery, I'd add a menu option to restore from a specific backup file. I could also implement a journaling system where every change is logged, allowing me to replay changes if the main file gets corrupted. Regular automated backups to different locations would protect against data loss and file corruption."

### **Extension & Enhancement Questions**

**Q19: How would you modify this to support multiple courses or departments?**
**A**: "I'd create a more complex data model with separate structs for Course, Department, and Student, with proper relationships between them. I might use a map[string][]*Student where the key is the department name. For course management, I'd add enrollment tracking. The JSON serialization would need to handle these nested objects, and I'd need to ensure all related structs have proper JSON tags."

**Q20: How would you add CSV export functionality?**
**A**: "I'd add an export function that uses the `encoding/csv` package to write student data to CSV format. This would involve creating a CSV writer, writing a header row, then iterating through the students slice and writing each student as a row. I'd also add error handling for file operations and maybe provide options for different CSV formats or field selections. This would make the data more portable and usable in spreadsheet applications."

**Q21: What improvements would you make for a multi-user environment?**
**A**: "For multi-user, I'd need to add concurrency control using mutexes or channels to protect the student slice during modifications. I'd implement proper authentication and authorization to control who can access or modify student data. I'd also move from file-based storage to a database that can handle concurrent transactions. And I'd add audit logging to track who made what changes and when."

**Q22: How would you implement a REST API for this student management system?**
**A**: "I'd use Go's `net/http` package to create HTTP endpoints for CRUD operations. I'd implement handlers for GET /students (list), POST /students (create), GET /students/{id} (get one), PUT /students/{id} (update), and DELETE /students/{id} (delete). I'd use JSON for request/response bodies and implement proper HTTP status codes. For a production API, I'd also add authentication, rate limiting, and proper error handling with meaningful HTTP error responses."

---
