# Bonus: Top 20 SQL Queries for Interviews

**Principle**: Java Developers are expected to know basic SQL (especially Joins and Aggregation).

## 1. Find Nth Highest Salary
**Question**: Find the 2nd highest salary from `Employee` table.
**Query**:
```sql
-- Method 1: Limit
SELECT salary FROM Employee ORDER BY salary DESC LIMIT 1 OFFSET 1;

-- Method 2: Subquery
SELECT MAX(salary) FROM Employee WHERE salary < (SELECT MAX(salary) FROM Employee);
```

## 2. Duplicate Records
**Question**: Find duplicate names in `Employee` table.
**Query**:
```sql
SELECT name, COUNT(*) FROM Employee GROUP BY name HAVING COUNT(*) > 1;
```

## 3. Delete Duplicates
**Question**: Delete duplicate records keep the one with lowest ID.
**Query**:
```sql
DELETE e1 FROM Employee e1, Employee e2 
WHERE e1.salary = e2.salary AND e1.id > e2.id;
```

## 4. Count Employees per Department
**Question**: Show Department Name and Count of Employees.
**Query**:
```sql
SELECT d.dept_name, COUNT(e.id) 
FROM Department d LEFT JOIN Employee e ON d.id = e.dept_id 
GROUP BY d.dept_name;
```

## 5. Employees with Salary > Manager
**Question**: Find employees earning more than their managers.
**Query**:
```sql
SELECT e.name 
FROM Employee e JOIN Employee m ON e.manager_id = m.id 
WHERE e.salary > m.salary;
```

## 6. Max Salary per Department
**Question**: Find the highest salary in each department.
**Query**:
```sql
SELECT dept_id, MAX(salary) FROM Employee GROUP BY dept_id;
```

## 7. Pattern Matching
**Question**: Find names starting with 'A'.
**Query**:
```sql
SELECT * FROM Employee WHERE name LIKE 'A%';
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

## 9. Odd/Even Records
**Question**: Select employee with odd IDs.
**Query**:
```sql
SELECT * FROM Employee WHERE MOD(id, 2) = 1;
```

## 10. Copy Table Structure
**Question**: Create table with same structure but no data.
**Query**:
```sql
CREATE TABLE Emp_Copy AS SELECT * FROM Employee WHERE 1=0;
```

## 11. Common Records
**Question**: Find records common in Table A and Table B.
**Query**:
```sql
SELECT * FROM TableA INTERSECT SELECT * FROM TableB;
```

## 12. Records in A but not B
**Question**: Find IDs in A not in B.
**Query**:
```sql
SELECT id FROM TableA MINUS SELECT id FROM TableB;
-- OR
SELECT id FROM TableA WHERE id NOT IN (SELECT id FROM TableB);
```

## 13. Current Date/Time
**Question**: Select current date.
**Query**:
```sql
SELECT NOW(); -- MySQL
SELECT SYSDATE FROM DUAL; -- Oracle
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

## 15. Last 5 Records
**Question**: Fetch last 5 inserted records.
**Query**:
```sql
SELECT * FROM Employee ORDER BY id DESC LIMIT 5;
```

## 16. Replace Null
**Question**: Select name, if null display 'No Name'.
**Query**:
```sql
SELECT COALESCE(name, 'No Name') FROM Employee;
```

## 17. Union vs Union All
**Question**: Difference?
**Answer**: `UNION` removes duplicates, `UNION ALL` keeps them (faster).

## 18. Self Join
**Question**: Find pairs of employees with same salary.
**Query**:
```sql
SELECT e1.name, e2.name 
FROM Employee e1, Employee e2 
WHERE e1.salary = e2.salary AND e1.id != e2.id;
```

## 19. Rank Records
**Question**: Rank employees by salary.
**Query**:
```sql
SELECT name, salary, RANK() OVER (ORDER BY salary DESC) FROM Employee;
```

## 20. Switch Rows (M/F)
**Question**: Swap gender 'm' to 'f' and vice versa.
**Query**:
```sql
UPDATE Salary 
SET sex = CASE sex 
    WHEN 'm' THEN 'f' 
    ELSE 'm' 
END;
```
