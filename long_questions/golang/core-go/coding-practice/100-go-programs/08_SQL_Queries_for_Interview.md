# Bonus: Top 20 SQL Queries for Interviews

**Principle**: Go Developers are expected to know basic SQL (especially Joins and Aggregation).

## 1. Find Nth Highest Salary
**Question**: Find the 2nd highest salary from `Employee` table.
**Query**:
```sql
-- Method 1: Limit
SELECT salary FROM Employee ORDER BY salary DESC LIMIT 1 OFFSET 1;

-- Method 2: Subquery
SELECT MAX(salary) FROM Employee WHERE salary < (SELECT MAX(salary) FROM Employee);
```

**Go Implementation**:
```go
package main

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq" // PostgreSQL driver
)

func getSecondHighestSalary(db *sql.DB) (int, error) {
    var salary int
    query := `SELECT MAX(salary) FROM Employee WHERE salary < (SELECT MAX(salary) FROM Employee)`
    err := db.QueryRow(query).Scan(&salary)
    return salary, err
}

func main() {
    db, err := sql.Open("postgres", "user=postgres dbname=test sslmode=disable")
    if err != nil {
        return
    }
    defer db.Close()
    
    salary, err := getSecondHighestSalary(db)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Second highest salary:", salary)
}
```

## 2. Duplicate Records
**Question**: Find duplicate names in `Employee` table.
**Query**:
```sql
SELECT name, COUNT(*) FROM Employee GROUP BY name HAVING COUNT(*) > 1;
```

**Go Implementation**:
```go
package main

import (
    "database/sql"
    "fmt"
)

type DuplicateRecord struct {
    Name  string
    Count int
}

func findDuplicates(db *sql.DB) ([]DuplicateRecord, error) {
    query := `SELECT name, COUNT(*) FROM Employee GROUP BY name HAVING COUNT(*) > 1`
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var duplicates []DuplicateRecord
    for rows.Next() {
        var dr DuplicateRecord
        if err := rows.Scan(&dr.Name, &dr.Count); err != nil {
            return nil, err
        }
        duplicates = append(duplicates, dr)
    }
    
    return duplicates, nil
}
```

## 3. Delete Duplicates
**Question**: Delete duplicate records keep the one with lowest ID.
**Query**:
```sql
DELETE e1 FROM Employee e1, Employee e2 
WHERE e1.salary = e2.salary AND e1.id > e2.id;
```

**Go Implementation**:
```go
package main

import (
    "database/sql"
    "fmt"
)

func deleteDuplicates(db *sql.DB) error {
    query := `DELETE e1 FROM Employee e1, Employee e2 WHERE e1.salary = e2.salary AND e1.id > e2.id`
    result, err := db.Exec(query)
    if err != nil {
        return err
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    
    fmt.Printf("Deleted %d duplicate rows\n", rowsAffected)
    return nil
}
```

## 4. Count Employees per Department
**Question**: Show Department Name and Count of Employees.
**Query**:
```sql
SELECT d.dept_name, COUNT(e.id) 
FROM Department d LEFT JOIN Employee e ON d.id = e.dept_id 
GROUP BY d.dept_name;
```

**Go Implementation**:
```go
package main

import (
    "database/sql"
    "fmt"
)

type DeptCount struct {
    DeptName string
    Count    int
}

func countEmployeesPerDept(db *sql.DB) ([]DeptCount, error) {
    query := `SELECT d.dept_name, COUNT(e.id) FROM Department d LEFT JOIN Employee e ON d.id = e.dept_id GROUP BY d.dept_name`
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var counts []DeptCount
    for rows.Next() {
        var dc DeptCount
        if err := rows.Scan(&dc.DeptName, &dc.Count); err != nil {
            return nil, err
        }
        counts = append(counts, dc)
    }
    
    return counts, nil
}
```

## 5. Employees with Salary > Manager
**Question**: Find employees earning more than their managers.
**Query**:
```sql
SELECT e.name 
FROM Employee e JOIN Employee m ON e.manager_id = m.id 
WHERE e.salary > m.salary;
```

**Go Implementation**:
```go
package main

import (
    "database/sql"
    "fmt"
)

func findEmployeesWithHigherSalaryThanManager(db *sql.DB) ([]string, error) {
    query := `SELECT e.name FROM Employee e JOIN Employee m ON e.manager_id = m.id WHERE e.salary > m.salary`
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var names []string
    for rows.Next() {
        var name string
        if err := rows.Scan(&name); err != nil {
            return nil, err
        }
        names = append(names, name)
    }
    
    return names, nil
}
```

## 6. Max Salary per Department
**Question**: Find the highest salary in each department.
**Query**:
```sql
SELECT dept_id, MAX(salary) FROM Employee GROUP BY dept_id;
```

**Go Implementation**:
```go
package main

import (
    "database/sql"
    "fmt"
)

type MaxSalaryPerDept struct {
    DeptID  int
    MaxSalary int
}

func getMaxSalaryPerDept(db *sql.DB) ([]MaxSalaryPerDept, error) {
    query := `SELECT dept_id, MAX(salary) FROM Employee GROUP BY dept_id`
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var results []MaxSalaryPerDept
    for rows.Next() {
        var ms MaxSalaryPerDept
        if err := rows.Scan(&ms.DeptID, &ms.MaxSalary); err != nil {
            return nil, err
        }
        results = append(results, ms)
    }
    
    return results, nil
}
```

## 7. Pattern Matching
**Question**: Find names starting with 'A'.
**Query**:
```sql
SELECT * FROM Employee WHERE name LIKE 'A%';
```

**Go Implementation**:
```go
package main

import (
    "database/sql"
    "fmt"
)

type Employee struct {
    ID     int
    Name   string
    Salary int
}

func findEmployeesStartingWithA(db *sql.DB) ([]Employee, error) {
    query := `SELECT id, name, salary FROM Employee WHERE name LIKE 'A%'`
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var employees []Employee
    for rows.Next() {
        var emp Employee
        if err := rows.Scan(&emp.ID, &emp.Name, &emp.Salary); err != nil {
            return nil, err
        }
        employees = append(employees, emp)
    }
    
    return employees, nil
}
```

## 8. Join 3 Tables
**Question**: Student, Course, Enrollment.
**Query**:
```sql
SELECT s.name, c.course_name 
FROM Student s 
JOIN Enrollment e ON s.id = e.student_id 
JOIN Course c ON e.course_id = c.id;
```

**Go Implementation**:
```go
package main

import (
    "database/sql"
    "fmt"
)

type StudentCourse struct {
    StudentName string
    CourseName  string
}

func getStudentCourses(db *sql.DB) ([]StudentCourse, error) {
    query := `SELECT s.name, c.course_name FROM Student s JOIN Enrollment e ON s.id = e.student_id JOIN Course c ON e.course_id = c.id`
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var results []StudentCourse
    for rows.Next() {
        var sc StudentCourse
        if err := rows.Scan(&sc.StudentName, &sc.CourseName); err != nil {
            return nil, err
        }
        results = append(results, sc)
    }
    
    return results, nil
}
```

## 9. Odd/Even Records
**Question**: Select employee with odd IDs.
**Query**:
```sql
SELECT * FROM Employee WHERE MOD(id, 2) = 1;
```

**Go Implementation**:
```go
package main

import (
    "database/sql"
    "fmt"
)

func getEmployeesWithOddID(db *sql.DB) ([]Employee, error) {
    query := `SELECT id, name, salary FROM Employee WHERE id % 2 = 1`
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var employees []Employee
    for rows.Next() {
        var emp Employee
        if err := rows.Scan(&emp.ID, &emp.Name, &emp.Salary); err != nil {
            return nil, err
        }
        employees = append(employees, emp)
    }
    
    return employees, nil
}
```

## 10. Copy Table Structure
**Question**: Create table with same structure but no data.
**Query**:
```sql
CREATE TABLE Emp_Copy AS SELECT * FROM Employee WHERE 1=0;
```

**Go Implementation**:
```go
package main

import (
    "database/sql"
    "fmt"
)

func copyTableStructure(db *sql.DB) error {
    query := `CREATE TABLE Emp_Copy AS SELECT * FROM Employee WHERE 1=0`
    _, err := db.Exec(query)
    return err
}
```

## 11. Common Records
**Question**: Find records common in Table A and Table B.
**Query**:
```sql
SELECT * FROM TableA INTERSECT SELECT * FROM TableB;
```

**Go Implementation**:
```go
package main

import (
    "database/sql"
    "fmt"
)

func findCommonRecords(db *sql.DB) ([]map[string]interface{}, error) {
    query := `SELECT * FROM TableA INTERSECT SELECT * FROM TableB`
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    // Get column names
    columns, err := rows.Columns()
    if err != nil {
        return nil, err
    }
    
    var results []map[string]interface{}
    for rows.Next() {
        // Create a slice of interface{} to hold the values
        values := make([]interface{}, len(columns))
        valuePtrs := make([]interface{}, len(columns))
        
        for i := range columns {
            valuePtrs[i] = &values[i]
        }
        
        if err := rows.Scan(valuePtrs...); err != nil {
            return nil, err
        }
        
        // Create a map for this row
        row := make(map[string]interface{})
        for i, col := range columns {
            row[col] = values[i]
        }
        
        results = append(results, row)
    }
    
    return results, nil
}
```

## 12. Records in A but not B
**Question**: Find IDs in A not in B.
**Query**:
```sql
SELECT id FROM TableA MINUS SELECT id FROM TableB;
-- OR
SELECT id FROM TableA WHERE id NOT IN (SELECT id FROM TableB);
```

**Go Implementation**:
```go
package main

import (
    "database/sql"
    "fmt"
)

func findIDsInANotInB(db *sql.DB) ([]int, error) {
    query := `SELECT id FROM TableA WHERE id NOT IN (SELECT id FROM TableB)`
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var ids []int
    for rows.Next() {
        var id int
        if err := rows.Scan(&id); err != nil {
            return nil, err
        }
        ids = append(ids, id)
    }
    
    return ids, nil
}
```

## 13. Current Date/Time
**Question**: Select current date.
**Query**:
```sql
SELECT NOW(); -- MySQL
SELECT SYSDATE FROM DUAL; -- Oracle
```

**Go Implementation**:
```go
package main

import (
    "database/sql"
    "fmt"
    "time"
)

func getCurrentDateTime(db *sql.DB) (time.Time, error) {
    var now time.Time
    query := `SELECT NOW()`
    err := db.QueryRow(query).Scan(&now)
    return now, err
}
```

## 14. Update with Join
**Question**: Increase salary by 10% for IT department.
**Query**:
```sql
UPDATE Employee e 
JOIN Department d ON e.dept_id = d.id 
SET e.salary = e.salary * 1.1 
WHERE d.name = 'IT';
```

**Go Implementation**:
```go
package main

import (
    "database/sql"
    "fmt"
)

func increaseITSalary(db *sql.DB) error {
    query := `UPDATE Employee e JOIN Department d ON e.dept_id = d.id SET e.salary = e.salary * 1.1 WHERE d.name = 'IT'`
    result, err := db.Exec(query)
    if err != nil {
        return err
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    
    fmt.Printf("Updated %d employee salaries\n", rowsAffected)
    return nil
}
```

## 15. Last 5 Records
**Question**: Fetch last 5 inserted records.
**Query**:
```sql
SELECT * FROM Employee ORDER BY id DESC LIMIT 5;
```

**Go Implementation**:
```go
package main

import (
    "database/sql"
    "fmt"
)

func getLast5Records(db *sql.DB) ([]Employee, error) {
    query := `SELECT id, name, salary FROM Employee ORDER BY id DESC LIMIT 5`
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var employees []Employee
    for rows.Next() {
        var emp Employee
        if err := rows.Scan(&emp.ID, &emp.Name, &emp.Salary); err != nil {
            return nil, err
        }
        employees = append(employees, emp)
    }
    
    return employees, nil
}
```

## 16. Replace Null
**Question**: Select name, if null display 'No Name'.
**Query**:
```sql
SELECT COALESCE(name, 'No Name') FROM Employee;
```

**Go Implementation**:
```go
package main

import (
    "database/sql"
    "fmt"
)

func getNamesWithDefault(db *sql.DB) ([]string, error) {
    query := `SELECT COALESCE(name, 'No Name') FROM Employee`
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var names []string
    for rows.Next() {
        var name string
        if err := rows.Scan(&name); err != nil {
            return nil, err
        }
        names = append(names, name)
    }
    
    return names, nil
}
```

## 17. Union vs Union All
**Question**: Difference?
**Answer**: `UNION` removes duplicates, `UNION ALL` keeps them (faster).

**Go Implementation**:
```go
package main

import (
    "database/sql"
    "fmt"
)

// UNION - removes duplicates
func getUnionRecords(db *sql.DB) ([]string, error) {
    query := `SELECT name FROM TableA UNION SELECT name FROM TableB`
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var names []string
    for rows.Next() {
        var name string
        if err := rows.Scan(&name); err != nil {
            return nil, err
        }
        names = append(names, name)
    }
    
    return names, nil
}

// UNION ALL - keeps duplicates
func getUnionAllRecords(db *sql.DB) ([]string, error) {
    query := `SELECT name FROM TableA UNION ALL SELECT name FROM TableB`
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var names []string
    for rows.Next() {
        var name string
        if err := rows.Scan(&name); err != nil {
            return nil, err
        }
        names = append(names, name)
    }
    
    return names, nil
}
```

## 18. Self Join
**Question**: Find pairs of employees with same salary.
**Query**:
```sql
SELECT e1.name, e2.name 
FROM Employee e1, Employee e2 
WHERE e1.salary = e2.salary AND e1.id != e2.id;
```

**Go Implementation**:
```go
package main

import (
    "database/sql"
    "fmt"
)

type EmployeePair struct {
    Name1 string
    Name2 string
}

func findEmployeesWithSameSalary(db *sql.DB) ([]EmployeePair, error) {
    query := `SELECT e1.name, e2.name FROM Employee e1, Employee e2 WHERE e1.salary = e2.salary AND e1.id != e2.id`
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var pairs []EmployeePair
    for rows.Next() {
        var pair EmployeePair
        if err := rows.Scan(&pair.Name1, &pair.Name2); err != nil {
            return nil, err
        }
        pairs = append(pairs, pair)
    }
    
    return pairs, nil
}
```

## 19. Rank Records
**Question**: Rank employees by salary.
**Query**:
```sql
SELECT name, salary, RANK() OVER (ORDER BY salary DESC) FROM Employee;
```

**Go Implementation**:
```go
package main

import (
    "database/sql"
    "fmt"
)

type RankedEmployee struct {
    Name   string
    Salary int
    Rank   int
}

func getRankedEmployees(db *sql.DB) ([]RankedEmployee, error) {
    query := `SELECT name, salary, RANK() OVER (ORDER BY salary DESC) FROM Employee`
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    
    var employees []RankedEmployee
    for rows.Next() {
        var emp RankedEmployee
        if err := rows.Scan(&emp.Name, &emp.Salary, &emp.Rank); err != nil {
            return nil, err
        }
        employees = append(employees, emp)
    }
    
    return employees, nil
}
```

## 20. Switch Rows (M/F)
**Question**: Swap gender 'm' to 'f' and vice versa.
**Query**:
```sql
UPDATE Employee 
SET gender = CASE gender 
    WHEN 'm' THEN 'f' 
    ELSE 'm' 
END;
```

**Go Implementation**:
```go
package main

import (
    "database/sql"
    "fmt"
)

func swapGender(db *sql.DB) error {
    query := `UPDATE Employee SET gender = CASE gender WHEN 'm' THEN 'f' ELSE 'm' END`
    result, err := db.Exec(query)
    if err != nil {
        return err
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    
    fmt.Printf("Swapped gender for %d employees\n", rowsAffected)
    return nil
}
```
