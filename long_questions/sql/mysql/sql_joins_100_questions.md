# 100 SQL Join Questions - From Basic to Advanced

## ðŸŸ¢ Schema & Sample Data
**Description:** Use the following schema to answer the questions. The data includes edge cases like NULLs, duplicates, and unmatched records to test different join behaviors.

### 1. Tables Structure
```sql
CREATE TABLE Departments (
    DeptID INT PRIMARY KEY,
    DeptName VARCHAR(50)
);

CREATE TABLE Employees (
    EmpID INT PRIMARY KEY,
    EmpName VARCHAR(50),
    DeptID INT,
    ManagerID INT,
    Salary DECIMAL(10, 2),
    JoinDate DATE
);

CREATE TABLE Projects (
    ProjectID INT PRIMARY KEY,
    ProjectName VARCHAR(50),
    Budget DECIMAL(15, 2)
);

CREATE TABLE Employee_Projects (
    EmpID INT,
    ProjectID INT,
    HoursWorked INT,
    PRIMARY KEY (EmpID, ProjectID)
);
```

### 2. Sample Data
```sql
-- Departments
INSERT INTO Departments VALUES (1, 'HR'), (2, 'IT'), (3, 'Finance'), (4, 'Marketing');

-- Employees (Note: EmpID 105 has NULL DeptID, EmpID 106 has invalid DeptID 99)
INSERT INTO Employees VALUES 
(101, 'Alice', 1, NULL, 60000, '2020-01-15'),
(102, 'Bob', 2, 101, 80000, '2020-06-01'),
(103, 'Charlie', 2, 102, 75000, '2021-03-10'),
(104, 'David', 3, 101, 70000, '2021-05-20'),
(105, 'Eve', NULL, 101, 55000, '2022-01-05'),
(106, 'Frank', 99, 102, 50000, '2022-08-15');

-- Projects
INSERT INTO Projects VALUES (10, 'Alpha'), (20, 'Beta'), (30, 'Gamma'), (40, 'Delta');

-- Employee_Projects
INSERT INTO Employee_Projects VALUES 
(101, 10, 20), (101, 20, 15),
(102, 10, 30), (102, 30, 10),
(103, 20, 40), 
(105, 10, 25);
```

---

## ðŸŸ¢ Part 1: Basic Joins (Questions 1-20)

### Scenario 1: Basic Employee-Department list
**Question 1:** Retrieve a list of all employees and their department names. Only include employees who are assigned to a valid existing department.
**Logical Reasoning:** We need matches in both tables. Employees with NULL or invalid department IDs should be excluded.
**Join Type:** `INNER JOIN` matches rows present in both tables.
```sql
SELECT E.EmpName, D.DeptName
FROM Employees E
INNER JOIN Departments D ON E.DeptID = D.DeptID;
```

### Scenario 2: Finding all employees including those without departments
**Question 2:** Retrieve a list of all employees and their department names, even if they explicitly do not have a department assigned.
**Logical Reasoning:** We want *all* rows from the Employees table (Left table), regardless of whether a match exists in Departments.
**Join Type:** `LEFT JOIN` (or LEFT OUTER JOIN).
```sql
SELECT E.EmpName, D.DeptName
FROM Employees E
LEFT JOIN Departments D ON E.DeptID = D.DeptID;
```

### Scenario 3: Departments with no employees
**Question 3:** List all departments and the employees working in them. Include departments that have no employees currently assigned.
**Logical Reasoning:** We want to ensure every Department (Right table) is listed, even if there is no matching Employee.
**Join Type:** `RIGHT JOIN` (or `LEFT JOIN` if tables are swapped).
```sql
SELECT D.DeptName, E.EmpName
FROM Employees E
RIGHT JOIN Departments D ON E.DeptID = D.DeptID;
```

### Scenario 4: Full list of Employees and Departments
**Question 4:** Retrieve a complete list of all employees and all departments, matching them where possible, but including unmatched records from both sides.
**Logical Reasoning:** We need the union of both sets: Employees without Departments, and Departments without Employees.
**Join Type:** `FULL JOIN` (Note: MySQL does not support `FULL JOIN` directly; use `UNION`).
```sql
SELECT E.EmpName, D.DeptName
FROM Employees E
LEFT JOIN Departments D ON E.DeptID = D.DeptID
UNION
SELECT E.EmpName, D.DeptName
FROM Employees E
RIGHT JOIN Departments D ON E.DeptID = D.DeptID;
```

### Scenario 5: Finding Orphan Records (Invalid Departments)
**Question 5:** Find employees who have a `DeptID` that does not exist in the `Departments` table.
**Logical Reasoning:** We need to find rows in the left table (Employees) that fail to match the right table (Departments).
**Join Type:** `LEFT JOIN` with `WHERE ... IS NULL`.
```sql
SELECT E.EmpName, E.DeptID
FROM Employees E
LEFT JOIN Departments D ON E.DeptID = D.DeptID
WHERE D.DeptID IS NULL;
```

### Scenario 6: Employees working on specific projects
**Question 6:** List each employee and the names of the projects they are working on.
**Logical Reasoning:** This is a many-to-many relationship involving an intermediary table (`Employee_Projects`). We need to join three tables.
**Join Type:** Multiple `INNER JOIN`s.
```sql
SELECT E.EmpName, P.ProjectName
FROM Employees E
INNER JOIN Employee_Projects EP ON E.EmpID = EP.EmpID
INNER JOIN Projects P ON EP.ProjectID = P.ProjectID;
```

### Scenario 7: Project Budgets and Hours
**Question 7:** Show the project name, its budget, and the total hours worked on it by all employees.
**Logical Reasoning:** We need to aggregate header data from `Projects` with transactional data from `Employee_Projects`.
**Join Type:** `INNER JOIN` with `GROUP BY`.
```sql
SELECT P.ProjectName, P.Budget, SUM(EP.HoursWorked) as TotalHours
FROM Projects P
JOIN Employee_Projects EP ON P.ProjectID = EP.ProjectID
GROUP BY P.ProjectName, P.Budget;
```

### Scenario 8: Self Join - Manager Names
**Question 8:** Retrieve a list of employees along with their manager's name. Use aliases to distinguish between the employee and the manager.
**Logical Reasoning:** The manager is also an employee. We must join the `Employees` table with itself.
**Join Type:** `INNER JOIN` (Self Join).
```sql
SELECT E.EmpName as Employee, M.EmpName as Manager
FROM Employees E
INNER JOIN Employees M ON E.ManagerID = M.EmpID;
```

### Scenario 9: Managers including top-level (No Manager)
**Question 9:** Retrieve all employees and their manager's name. If an employee has no manager (e.g., the CEO), display 'Top Level' or NULL.
**Logical Reasoning:** An `INNER JOIN` would filter out the CEO (NULL ManagerID). We need a `LEFT JOIN` to keep employees with NULL managers.
**Join Type:** `LEFT JOIN` (Self Join).
```sql
SELECT E.EmpName, COALESCE(M.EmpName, 'Top Level') as Manager
FROM Employees E
LEFT JOIN Employees M ON E.ManagerID = M.EmpID;
```

### Scenario 10: Cross Join - Possible Assignments
**Question 10:** Generate a list of every possible combination of Employee and Project to evaluate potential assignments.
**Logical Reasoning:** We need a Cartesian product of all employees and all projects.
**Join Type:** `CROSS JOIN`.
```sql
SELECT E.EmpName, P.ProjectName
FROM Employees E
CROSS JOIN Projects P;
```

### Scenario 11: Employees who have not worked on any project
**Question 11:** Find employees who are not assigned to any project.
**Logical Reasoning:** Similar to finding invalid departments, we look for non-matches in the `Employee_Projects` table.
**Join Type:** `LEFT JOIN` with `IS NULL` check.
```sql
SELECT E.EmpName
FROM Employees E
LEFT JOIN Employee_Projects EP ON E.EmpID = EP.EmpID
WHERE EP.ProjectID IS NULL;
```

### Scenario 12: Projects with no assigned employees
**Question 12:** List projects that have zero hours logged (no employees assigned).
**Logical Reasoning:** Find projects that do not exist in the assignment table.
**Join Type:** `LEFT JOIN` (from Project perspective).
```sql
SELECT P.ProjectName
FROM Projects P
LEFT JOIN Employee_Projects EP ON P.ProjectID = EP.ProjectID
WHERE EP.EmpID IS NULL;
```

### Scenario 13: Department Salary Stats
**Question 13:** Calculate the average salary for each department, including the department name.
**Logical Reasoning:** Join Employees to Departments and aggregate salaries.
**Join Type:** `INNER JOIN` with `GROUP BY`.
```sql
SELECT D.DeptName, AVG(E.Salary) as AvgSalary
FROM Departments D
JOIN Employees E ON D.DeptID = E.DeptID
GROUP BY D.DeptName;
```

### Scenario 14: Employees in IT or HR
**Question 14:** List employees who work in either 'IT' or 'HR' departments.
**Logical Reasoning:** Filter the joined results by specific department names.
**Join Type:** `INNER JOIN` with `WHERE IN`.
```sql
SELECT E.EmpName, D.DeptName
FROM Employees E
JOIN Departments D ON E.DeptID = D.DeptID
WHERE D.DeptName IN ('IT', 'HR');
```

### Scenario 15: Salary comparison with Manager
**Question 15:** Find employees who earn more than their direct manager.
**Logical Reasoning:** Self-join to align Employee salary with Manager salary, then filter.
**Join Type:** `INNER JOIN` (Self Join).
```sql
SELECT E.EmpName as Employee, E.Salary as EmpSalary, M.EmpName as Manager, M.Salary as MgrSalary
FROM Employees E
JOIN Employees M ON E.ManagerID = M.EmpID
WHERE E.Salary > M.Salary;
```

### Scenario 16: Multiple Join conditions
**Question 16:** Join Employees and Departments, but only for employees who joined after 2021.
**Logical Reasoning:** We can place the filter condition inside the `ON` clause or in the `WHERE` clause.
**Join Type:** `INNER JOIN` with complex condition.
```sql
SELECT E.EmpName, D.DeptName
FROM Employees E
JOIN Departments D ON E.DeptID = D.DeptID
WHERE E.JoinDate > '2021-01-01';
```

### Scenario 17: Non-Equi Join (Salary Ranges)
**Question 17:** Assuming a `SalaryGrades` table exists (ID, MinSal, MaxSal), find the grade of each employee.
**Logical Reasoning:** The join condition is not Equality (`=`), but a range (`BETWEEN`).
**Join Type:** `INNER JOIN` with `BETWEEN`.
```sql
-- Conceptual Query (assuming Table SalaryGrades)
SELECT E.EmpName, S.Grade
FROM Employees E
JOIN SalaryGrades S ON E.Salary BETWEEN S.MinSal AND S.MaxSal;
```

### Scenario 18: Join and distinct counts
**Question 18:** Count how many unique projects each department is working on (via their employees).
**Logical Reasoning:** Join Dept -> Emp -> EmpProj -> Proj. Count Distinct ProjectIDs.
**Join Type:** Multiple `INNER JOIN`s.
```sql
SELECT D.DeptName, COUNT(DISTINCT EP.ProjectID) as UniqueProjects
FROM Departments D
JOIN Employees E ON D.DeptID = E.DeptID
JOIN Employee_Projects EP ON E.EmpID = EP.EmpID
GROUP BY D.DeptName;
```

### Scenario 19: Finding specific string patterns in Joins
**Question 19:** List employees in the 'Finance' department whose name starts with 'D'.
**Logical Reasoning:** Join for context, `LIKE` for filtering.
**Join Type:** `INNER JOIN`.
```sql
SELECT E.EmpName
FROM Employees E
JOIN Departments D ON E.DeptID = D.DeptID
WHERE D.DeptName = 'Finance' AND E.EmpName LIKE 'D%';
```

### Scenario 20: Natural Join (Theoretical)
**Question 20:** Using `NATURAL JOIN` to link Employees and Departments. What is the risk?
**Logical Reasoning:** `NATURAL JOIN` joins on all columns with the same name. If both tables have a `Name` column (DeptName vs EmpName), it might fail or give wrong results.
**Join Type:** `NATURAL JOIN`.
```sql
-- Works if the only common column is DeptID
SELECT * 
FROM Employees 
NATURAL JOIN Departments;
-- Risk: If you add 'CreatedDate' to both tables later, the existing query changes logic silently.
```

---

## ðŸŸ¡ Part 2: Advanced Join Concepts (Questions 21-40)

### Scenario 21: Filtering in ON vs WHERE clause
**Question 21:** Explain the difference between putting a filter condition in the `ON` clause vs the `WHERE` clause for a `LEFT JOIN`.
**Logical Reasoning:** In `LEFT JOIN`, `ON` filters rows *before* joining (preserving left rows), while `WHERE` filters rows *after* the join (potentially removing unmatched left rows).
**Join Type:** `LEFT JOIN`.
```sql
-- Keeps all employees, only joins department if Name is 'HR'
SELECT E.EmpName, D.DeptName
FROM Employees E
LEFT JOIN Departments D ON E.DeptID = D.DeptID AND D.DeptName = 'HR';

-- Removes employees who are NOT in HR (effectively becomes Inner Join)
SELECT E.EmpName, D.DeptName
FROM Employees E
LEFT JOIN Departments D ON E.DeptID = D.DeptID
WHERE D.DeptName = 'HR';
```

### Scenario 22: Count(*) on Left Join
**Question 22:** You want to count how many employees are in each department. Why gives `count(*)` vs `count(E.EmpID)` different results on a Right Join or List of Depts?
**Logical Reasoning:** `COUNT(*)` counts the row even if the joined side is NULL (effectively counting the Department itself as 1 employee). `COUNT(Column)` ignores NULLs.
**Join Type:** `LEFT JOIN` (Departments -> Employees).
```sql
SELECT D.DeptName, COUNT(E.EmpID) as RealCount
FROM Departments D
LEFT JOIN Employees E ON D.DeptID = E.DeptID
GROUP BY D.DeptName;
-- Using COUNT(*) would incorrectly show 1 for empty departments because the row (Dept, NULL) exists.
```

### Scenario 23: Handling Nulls in Joins with COALESCE
**Question 23:** Retrieve instances where the Project Name is missing in an assignment list, replacing it with 'No Project Assigned'.
**Logical Reasoning:** When a `LEFT JOIN` results in a NULL, use `COALESCE` to provide a default value.
**Join Type:** `LEFT JOIN` + `COALESCE`.
```sql
SELECT E.EmpName, COALESCE(P.ProjectName, 'No Project Assigned')
FROM Employees E
LEFT JOIN Employee_Projects EP ON E.EmpID = EP.EmpID
LEFT JOIN Projects P ON EP.ProjectID = P.ProjectID;
```

### Scenario 24: Finding Duplicates via Self Join
**Question 24:** Find employees who have the exact same salary as another employee.
**Logical Reasoning:** Join the table to itself on Salary, but ensure IDs are different to avoid self-matching.
**Join Type:** `INNER JOIN` (Self).
```sql
SELECT A.EmpName, B.EmpName, A.Salary
FROM Employees A
JOIN Employees B ON A.Salary = B.Salary
WHERE A.EmpID < B.EmpID; -- Ensures pairs like (Alice, Bob) aren't repeated as (Bob, Alice)
```

### Scenario 25: Identifying Missing Relations (Reverse)
**Question 25:** Find Departments that have no projects associated with them (assuming a Dept-Project link existed).
**Logical Reasoning:** (Hypothetical schema extension) Use `LEFT JOIN` and check for NULL.
**Join Type:** `LEFT JOIN` / `IS NULL`.
```sql
-- Concept:
SELECT D.DeptName 
FROM Departments D
LEFT JOIN Projects P ON D.DeptID = P.DeptID -- (Assuming Key exists)
WHERE P.ProjectID IS NULL;
```

### Scenario 26: 3-Way Left Join
**Question 26:** List all Employees, their Departments, and their Projects. Ensure Employees appear even if they have no Dept or no Project.
**Logical Reasoning:** Chain multiple `LEFT JOIN`s starting from the main entity (Employee).
**Join Type:** Multi `LEFT JOIN`.
```sql
SELECT E.EmpName, D.DeptName, P.ProjectName
FROM Employees E
LEFT JOIN Departments D ON E.DeptID = D.DeptID
LEFT JOIN Employee_Projects EP ON E.EmpID = EP.EmpID
LEFT JOIN Projects P ON EP.ProjectID = P.ProjectID;
```

### Scenario 27: Updates using Joins
**Question 27:** Increase the salary of all employees in the 'IT' Department by 10%.
**Logical Reasoning:** Use `JOIN` in the `UPDATE` statement to target specific rows based on related table data.
**Join Type:** `UPDATE` with `JOIN`.
```sql
UPDATE Employees E
JOIN Departments D ON E.DeptID = D.DeptID
SET E.Salary = E.Salary * 1.10
WHERE D.DeptName = 'IT';
```

### Scenario 28: Deletes using Joins
**Question 28:** Delete all employees who are working in a defined 'Temporary' department (ID 99).
**Logical Reasoning:** Delete from the target table based on a join condition.
**Join Type:** `DELETE` with `JOIN`.
```sql
DELETE E
FROM Employees E
JOIN Departments D ON E.DeptID = D.DeptID
WHERE D.DeptName = 'Temporary'; 
-- Or simply WHERE DeptID = 99 if known.
```

### Scenario 29: Cartesian Product Trap
**Question 29:** What happens if you forget the `ON` clause in a Join?
**Logical Reasoning:** It becomes a `CROSS JOIN`, producing rows = TableA_Count * TableB_Count.
**Join Type:** Implicit `CROSS JOIN`.
```sql
-- Accidentally returning 6 * 4 = 24 rows for 6 employees and 4 depts
SELECT * FROM Employees, Departments;
```

### Scenario 30: Aggregate Filter with Join
**Question 30:** Find Departments where the average salary is greater than 70,000.
**Logical Reasoning:** Join, Group By, then Filter the group using `HAVING`.
**Join Type:** `INNER JOIN` + `HAVING`.
```sql
SELECT D.DeptName, AVG(E.Salary)
FROM Departments D
JOIN Employees E ON D.DeptID = E.DeptID
GROUP BY D.DeptName
HAVING AVG(E.Salary) > 70000;
```

### Scenario 31: Greatest Value across matches
**Question 31:** For each employee, find the project with the highest budget they are working on.
**Logical Reasoning:** Join tables, then use `MAX` or Window Functions.
**Join Type:** `INNER JOIN` + Subquery/Window.
```sql
SELECT E.EmpName, MAX(P.Budget)
FROM Employees E
JOIN Employee_Projects EP ON E.EmpID = EP.EmpID
JOIN Projects P ON EP.ProjectID = P.ProjectID
GROUP BY E.EmpName;
```

### Scenario 32: Join on non-key columns
**Question 32:** Can you join two tables on a column that is not a Primary or Foreign Key?
**Logical Reasoning:** Yes, any compatible data type can be joined (e.g., joining on City names).
**Join Type:** `INNER JOIN`.
```sql
-- Example: Joining Employees and Suppliers if they are in the same City
SELECT E.EmpName, S.SupplierName
FROM Employees E
JOIN Suppliers S ON E.City = S.City;
```

### Scenario 33: Using aliases for readability
**Question 33:** Why is using aliases (AS) recommended in complex joins?
**Logical Reasoning:** Reduces typing, improves readability, and is *mandatory* in Self Joins to distinguish instances.
**Join Type:** General Best Practice.
```sql
SELECT T1.Col, T2.Col
FROM VeryLongTableName T1
JOIN AnotherTable T2 ON T1.ID = T2.ID;
```

### Scenario 34: Join Optimization (Indexing)
**Question 34:** Which column should be indexed to optimize `SELECT * FROM Orders O JOIN Customers C ON O.CustomerID = C.ID`?
**Logical Reasoning:** The Foreign Key `O.CustomerID` and Primary Key `C.ID` should be indexed.
**Join Type:** Performance Tuning.
```sql
-- Ensure Index exists on Orders(CustomerID)
CREATE INDEX idx_orders_cust ON Orders(CustomerID);
```

### Scenario 35: Exclusion Join (Using NOT EXISTS)
**Question 35:** Alternative to `LEFT JOIN ... WHERE IS NULL`?
**Logical Reasoning:** `NOT EXISTS` is semantically the same for finding non-matches and can be more performant in some DBs.
**Join Type:** Semijoin / Anti-Join.
```sql
SELECT E.EmpName 
FROM Employees E
WHERE NOT EXISTS (
    SELECT 1 FROM Departments D WHERE E.DeptID = D.DeptID
);
```

### Scenario 36: Join vs Subquery Performance
**Question 36:** Is a JOIN always faster than a Subquery?
**Logical Reasoning:** Not always. Modern optimizers often rewrite subqueries as joins. However, `IN` lists with thousands of items can be slower than a straight `JOIN`.
**Join Type:** Theoretical.

### Scenario 37: Handling Duplicate records in Join
**Question 37:** If Table A has 2 matches in Table B, does the result have 1 or 2 rows?
**Logical Reasoning:** It yields 2 rows. Joins multiply rows based on matches.
**Join Type:** One-to-Many Join.
```sql
-- If Emp A works on 3 projects, joining Emp to Project yields 3 rows for Emp A.
```

### Scenario 38: Joining with a Derived Table
**Question 38:** Join Employees with a subquery that calculates Dept average salary.
**Logical Reasoning:** Join a table with a temporary result set (Derived Table).
**Join Type:** `INNER JOIN` with Subquery.
```sql
SELECT E.EmpName, E.Salary, DAvg.AvgSal
FROM Employees E
JOIN (
    SELECT DeptID, AVG(Salary) as AvgSal 
    FROM Employees 
    GROUP BY DeptID
) DAvg ON E.DeptID = DAvg.DeptID;
```

### Scenario 39: Lateral Join (Postgres/Oracle specific but conceptual)
**Question 39:** What is a lateral join?
**Logical Reasoning:** It allows a subquery in the `FROM` clause to reference columns from preceding tables in the same `FROM` list. In MySQL 8.0+, similar to `LATERAL DERIVED TABLES`.
**Join Type:** `LATERAL JOIN`.
```sql
SELECT D.DeptName, HighEarner.EmpName
FROM Departments D,
LATERAL (SELECT * FROM Employees E WHERE E.DeptID = D.DeptID ORDER BY Salary DESC LIMIT 1) HighEarner;
```

### Scenario 40: Pivot with Join
**Question 40:** Create a report showing Department Names as columns and total salary as values.
**Logical Reasoning:** Requires Aggregation + Case logic (Pivoting).
**Join Type:** `crosstab` (Conceptual).
```sql
SELECT 
    SUM(CASE WHEN D.DeptName = 'HR' THEN E.Salary ELSE 0 END) as HR_Total,
    SUM(CASE WHEN D.DeptName = 'IT' THEN E.Salary ELSE 0 END) as IT_Total
FROM Employees E
JOIN Departments D ON E.DeptID = D.DeptID;
```

---

## ðŸŸ¢ Part 3: Deep Dive & Performance (Questions 41-60)

### Scenario 41: Grand-Manager Hierarchy (Self Join x2)
**Question 41:** Retrieve Employee Name, Manager Name, and the Manager's Manager Name (Senior Manager).
**Logical Reasoning:** We need to traverse the hierarchy two levels up. This requires joining `Employees` to itself twice.
**Join Type:** Double `LEFT JOIN` (Self).
```sql
SELECT E.EmpName, M.EmpName as Manager, SM.EmpName as SeniorManager
FROM Employees E
LEFT JOIN Employees M ON E.ManagerID = M.EmpID
LEFT JOIN Employees SM ON M.ManagerID = SM.EmpID;
```

### Scenario 42: Generating a Calendar (Cross Join)
**Question 42:** Generate a list of all dates for the current month combined with all employees (to track daily attendance, even if no entry exists).
**Logical Reasoning:** Use a recursive CTE (or a calendar table) to generate dates, then `CROSS JOIN` with Employees.
**Join Type:** `CROSS JOIN`.
```sql
WITH RECURSIVE Dates AS (
    SELECT '2023-10-01' as DateVal
    UNION ALL
    SELECT DATE_ADD(DateVal, INTERVAL 1 DAY) FROM Dates WHERE DateVal < '2023-10-31'
)
SELECT E.EmpName, D.DateVal
FROM Employees E
CROSS JOIN Dates D;
```

### Scenario 43: Finding Missing Numbers (Gap Analysis)
**Question 43:** Find which Project IDs between 1 and 50 are missing from the `Projects` table.
**Logical Reasoning:** Generate a sequence 1-50 (Virtual Table) and `LEFT JOIN` Projects to find NULLs.
**Join Type:** `LEFT JOIN` of Sequence vs Table.
```sql
-- Assuming a Numbers table or CTE exists
SELECT N.Number
FROM Numbers N
LEFT JOIN Projects P ON N.Number = P.ProjectID
WHERE P.ProjectID IS NULL AND N.Number <= 50;
```

### Scenario 44: Comparing Join vs Window Function for Running Total
**Question 44:** Self-joining to calculate running totals is possible (`T1.Date >= T2.Date`). Why is it discouraged?
**Logical Reasoning:** Self-joining for running totals creates an O(N^2) complexity triangle join. Window functions (`SUM() OVER`) are O(N) or O(N log N) and much faster.
**Join Type:** Performance comparison.
```sql
-- Inefficient Join method
SELECT T1.Date, SUM(T2.Val)
FROM Sales T1 JOIN Sales T2 ON T1.Date >= T2.Date
GROUP BY T1.Date;

-- Efficient Window method
SELECT Date, SUM(Val) OVER (ORDER BY Date) FROM Sales;
```

### Scenario 45: Complex Multi-Join with Aggregation
**Question 45:** List Departments, the number of employees, and the number of projects being worked on by that department.
**Logical Reasoning:** Join Dept -> Emp -> EmpProj. Aggegrate carefully to avoid "fanning out" counts (using `COUNT(DISTINCT)`).
**Join Type:** Multi-Join + Aggregation.
```sql
SELECT D.DeptName, 
       COUNT(DISTINCT E.EmpID) as EmpCount, 
       COUNT(DISTINCT EP.ProjectID) as ProjCount
FROM Departments D
LEFT JOIN Employees E ON D.DeptID = E.DeptID
LEFT JOIN Employee_Projects EP ON E.EmpID = EP.EmpID
GROUP BY D.DeptName;
```

### Scenario 46: JOIN with OR Condition
**Question 46:** Join Employees to an "events" table if the employee is the organizer OR the participant.
**Logical Reasoning:** Join condition usually checks equality, but can check logical OR. Note: This often disables index usage.
**Join Type:** `INNER JOIN` on `OR`.
```sql
SELECT E.EmpName, Ev.EventName
FROM Employees E
JOIN Events Ev ON E.EmpID = Ev.OrganizerID OR E.EmpID = Ev.ParticipantID;
```

### Scenario 47: JOIN with LIKE (Fuzzy Matching)
**Question 47:** Join an Imported Data table to the Employees table where the names might be misspelled or slightly different (substring match).
**Logical Reasoning:** Using `LIKE` in a join condition. Very slow, but functional for data cleaning.
**Join Type:** Non-Equi Join.
```sql
SELECT I.RawName, E.EmpName
FROM ImportedData I
JOIN Employees E ON I.RawName LIKE CONCAT('%', E.EmpName, '%');
```

### Scenario 48: USING clause
**Question 48:** Simplify `ON E.DeptID = D.DeptID` using the `USING` clause.
**Logical Reasoning:** If column names are identical, `USING(ColName)` is cleaner.
**Join Type:** Syntax variation.
```sql
SELECT E.EmpName, D.DeptName
FROM Employees E
JOIN Departments D USING (DeptID);
```

### Scenario 49: NATURAL LEFT JOIN
**Question 49:** What does `NATURAL LEFT JOIN` do?
**Logical Reasoning:** It performs a Left Join using all columns with matching names as keys, automatically coalescing the columns.
**Join Type:** Implicit Key Join.
```sql
SELECT * 
FROM Employees 
NATURAL LEFT JOIN Departments;
```

### Scenario 50: Simulating FULL OUTER JOIN in MySQL
**Question 50:** MySQL doesn't have `FULL OUTER JOIN`. How do you simulate it?
**Logical Reasoning:** `LEFT JOIN` UNION `RIGHT JOIN`.
**Join Type:** Union of Joins.
```sql
SELECT * FROM T1 LEFT JOIN T2 ON T1.ID = T2.ID
UNION
SELECT * FROM T1 RIGHT JOIN T2 ON T1.ID = T2.ID;
```

### Scenario 51: Join Algorithms - Nested Loop
**Question 51:** Explain the "Nested Loop Join" algorithm.
**Logical Reasoning:** The DB iterates through every row of the Outer table, and for each row, scans the Inner table for a match. Good for small datasets or when the inner table is indexed.
**Join Type:** Theoretical.

### Scenario 52: Join Algorithms - Hash Join
**Question 52:** Explain the "Hash Join" algorithm.
**Logical Reasoning:** The DB builds a Hash Map of the smaller table in memory, then scans the larger table and probes the hash map. Very fast for large, unsorted joins (Equi-joins only).
**Join Type:** Theoretical.

### Scenario 53: Join Explosion (Fan-out) problem
**Question 53:** You join `Orders` (100 rows) to `OrderItems` (300 rows). Why does `SUM(Orders.Total)` give a wrong result?
**Logical Reasoning:** The Join duplicates `Orders` rows for each Item. Summing the Order total includes duplicates.
**Fix:** Aggregate `OrderItems` *before* joining, or use `SUM` on `Orders` table separately.
**Join Type:** One-to-Many Pitfall.

### Scenario 54: Solving Join Explosion with CTEs
**Question 54:** Write a safe query to sum Order Totals and Item Counts without duplication risk.
**Logical Reasoning:** Pre-aggregate Item counts in a CTE, then join to Orders.
**Join Type:** `JOIN` to Pre-Aggregated CTE.
```sql
WITH ItemStats AS (
    SELECT OrderID, COUNT(*) as ItemCount 
    FROM OrderLines 
    GROUP BY OrderID
)
SELECT O.OrderID, O.TotalAmt, I.ItemCount
FROM Orders O
JOIN ItemStats I ON O.OrderID = I.OrderID;
```

### Scenario 55: Force Index in Join
**Question 55:** The optimizer is choosing a full table scan instead of an index. How do you force it?
**Logical Reasoning:** Use hints like `FORCE INDEX`.
**Join Type:** Performance Hint.
```sql
SELECT * 
FROM Employees E FORCE INDEX (idx_lastname)
JOIN Departments D ON E.DeptID = D.DeptID;
```

### Scenario 56: Maximum tables in a Join
**Question 56:** Is there a limit to how many tables you can join?
**Logical Reasoning:** MySQL has a limit (61 tables). SQL Server/Postgres have higher/practical limits. Performance degrades rapidly after 5-10 tables due to optimizer complexity (N! join orders).
**Join Type:** System interaction.

### Scenario 57: Joining Partitioned Tables
**Question 57:** Does joining partitioned tables require special syntax?
**Logical Reasoning:** No, the optimizer handles it ("Partition Pruning"). However, including the Partition Key in the `ON` clause helps performance significantly.
**Join Type:** Performance.

### Scenario 58: Anti-Join with NULLs (The trap)
**Question 58:** Why does `WHERE ID NOT IN (SELECT ID ...)` return 0 rows if the subquery contains a single NULL?
**Logical Reasoning:** `1 != NULL` is Unknown. If ANY value in the list is NULL, `NOT IN` returns Unknown (False) for everything. Always use `WHERE ID IS NOT NULL` in the subquery.
**Join Type:** `NOT IN` Logic.

### Scenario 59: Dynamic Joins (Polymorphic Associations)
**Question 59:** You have a `Comments` table that links to `Posts` OR `Photos` based on a `Type` column. How do you join?
**Logical Reasoning:** You need multiple Left Joins with logic in the ON clause.
**Join Type:** Polymorphic Join.
```sql
SELECT C.Comment, COALESCE(P.Title, Ph.Caption)
FROM Comments C
LEFT JOIN Posts P ON C.TargetID = P.ID AND C.Type = 'Post'
LEFT JOIN Photos Ph ON C.TargetID = Ph.ID AND C.Type = 'Photo';
```

### Scenario 60: Joining tables from different databases
**Question 60:** How do you join tables from two different databases on the same server?
**Logical Reasoning:** Qualify the table names with the database name.
**Join Type:** Cross-Database Join.
```sql
SELECT *
FROM Db1.TableA A
JOIN Db2.TableB B ON A.ID = B.ID;
```

---


## 🔵 Part 4: Complex Logic & Optimization Scenarios (Questions 61-80)

### Scenario 61: Generating Hierarchy Paths
**Question 61:** Create a string representing the full path from CEO to Employee (e.g., "CEO/Manager/Employee").
**Logical Reasoning:** Requires Recursive CTE to build the string level by level.
**Join Type:** Recursive Join.
```sql
WITH RECURSIVE PathCTE AS (
    SELECT EmpID, EmpName, CAST(EmpName AS CHAR(200)) as Path
    FROM Employees WHERE ManagerID IS NULL
    UNION ALL
    SELECT E.EmpID, E.EmpName, CONCAT(P.Path, '/', E.EmpName)
    FROM Employees E
    JOIN PathCTE P ON E.ManagerID = P.EmpID
)
SELECT * FROM PathCTE;
```

### Scenario 62: Finding "Islands" (Consecutive IDs)
**Question 62:** Find ranges of consecutive IDs in a table (e.g., 1,2,3... 10,11).
**Logical Reasoning:** Self Join to find where `Current.ID = Previous.ID + 1` fails (start of island) or usage of `ROW_NUMBER()`.
**Join Type:** Self Join (Offset).
```sql
-- Using Window Functions (More modern than Join method)
SELECT Min(ID) as StartRange, Max(ID) as EndRange
FROM (
    SELECT ID, ID - ROW_NUMBER() OVER (ORDER BY ID) as Grp
    FROM Logs
) T
GROUP BY Grp;
```

### Scenario 63: Finding "Gaps" (Missing IDs)
**Question 63:** Find missing ID numbers in a sequence (e.g., 1, 2, 4 -> Missing 3).
**Logical Reasoning:** Join `T1` to `T1` where `T2.ID = T1.ID + 1` and check where T2 is NULL.
**Join Type:** Self Left Join.
```sql
SELECT T1.ID + 1 as MissingID
FROM Tickets T1
LEFT JOIN Tickets T2 ON T1.ID + 1 = T2.ID
WHERE T2.ID IS NULL;
```

### Scenario 64: Overlapping Time Ranges (Scheduling)
**Question 64:** Find Booking requests that overlap with existing Bookings.
**Logical Reasoning:** Join table to itself (Requests vs Existing) on Time Overlap logic.
**Join Type:** Non-Equi Join (`StartA < EndB AND EndA > StartB`).
```sql
SELECT B1.RoomID, B1.Timerange
FROM Bookings B1
JOIN Bookings B2 ON B1.RoomID = B2.RoomID
    AND B1.ID <> B2.ID
    AND B1.StartTime < B2.EndTime 
    AND B1.EndTime > B2.StartTime;
```

### Scenario 65: Keyset Pagination (Optimized Scrolling)
**Question 65:** Why is `OFFSET` bad for pagination, and how can Joins help?
**Logical Reasoning:** `OFFSET` scans and discards rows. "Seek Method" uses a `WHERE` clause on the last seen ID. If needing to join for details, join *after* pagination.
**Join Type:** Late Join (Deferred Join).
```sql
SELECT E.*, D.DeptName
FROM (
    SELECT EmpID FROM Employees ORDER BY EmpID LIMIT 10 OFFSET 100000 
    -- ^ Slow
) Sub
JOIN Employees E ON Sub.EmpID = E.EmpID -- Join details later
JOIN Departments D ON E.DeptID = D.DeptID;
```

### Scenario 66: Count Distinct in Joined Subset
**Question 66:** How to optimize `COUNT(DISTINCT UserID)` in a massive join?
**Logical Reasoning:** `COUNT(DISTINCT)` is expensive. Pre-aggregate unique IDs in a subquery/CTE before joining.
**Join Type:** Join to Aggregate.

### Scenario 67: UPDATE based on Aggregate
**Question 67:** Set a `DepartmentSize` column in `Departments` table using data from `Employees`.
**Logical Reasoning:** Join `Departments` to an aggregate subquery of `Employees`.
**Join Type:** `UPDATE` with `JOIN` + `GROUP BY`.
```sql
UPDATE Departments D
JOIN (
    SELECT DeptID, COUNT(*) as Cnt FROM Employees GROUP BY DeptID
) E ON D.DeptID = E.DeptID
SET D.DepartmentSize = E.Cnt;
```

### Scenario 68: Delete Duplicates keeping Newest
**Question 68:** Delete duplicate email entries, keeping only the one with the latest JoinDate.
**Logical Reasoning:** Self Join. Delete `T1` if there exists `T2` with same Email but newer Date/ID.
**Join Type:** `DELETE` via Self Join.
```sql
DELETE T1
FROM Users T1
JOIN Users T2 ON T1.Email = T2.Email
WHERE T1.JoinDate < T2.JoinDate;
```

### Scenario 69: EAV Model Joins (Entity-Attribute-Value)
**Question 69:** You have `ProductAttributes` (ProdID, AttrName, AttrValue). How to get distinct columns for 'Color' and 'Size'?
**Logical Reasoning:** Multiple joins to the same table, one for each attribute needed.
**Join Type:** Multi Self Join.
```sql
SELECT P.ProdName, A1.AttrValue as Color, A2.AttrValue as Size
FROM Products P
LEFT JOIN ProductAttributes A1 ON P.ID = A1.ProdID AND A1.AttrName = 'Color'
LEFT JOIN ProductAttributes A2 ON P.ID = A2.ProdID AND A2.AttrName = 'Size';
```

### Scenario 70: JSON_TABLE Join (MySQL 8.0)
**Question 70:** You have a JSON column `Tags` (['A', 'B']). How to join this with a `Tags` table?
**Logical Reasoning:** Use `JSON_TABLE` to transform JSON array into rows, then join.
**Join Type:** JSON Table Function Join.
```sql
SELECT P.Name, T.TagName
FROM Products P,
     JSON_TABLE(P.Tags, '$[*]' COLUMNS(TagName VARCHAR(50) PATH '$')) AS T;
```

### Scenario 71: Lateral Derived Table (Top N per Group)
**Question 71:** Get the top 3 highest paid employees for *every* department.
**Logical Reasoning:** Use `LATERAL` (or correlated subquery in older MySQL) to limiting inside the join loop.
**Join Type:** Lateral Join.
```sql
SELECT D.DeptName, TopEmp.EmpName, TopEmp.Salary
FROM Departments D,
LATERAL (
    SELECT * FROM Employees E 
    WHERE E.DeptID = D.DeptID 
    ORDER BY Salary DESC LIMIT 3
) TopEmp;
```

### Scenario 72: STRAIGHT_JOIN hint
**Question 72:** The optimizer is joining the wrong table first (driving table). How to fix?
**Logical Reasoning:** `STRAIGHT_JOIN` forces the optimizer to join tables in the exact order listed in the `SUB` Query.
**Join Type:** Forced Order Join.
```sql
SELECT STRAIGHT_JOIN * 
FROM BigTable B 
JOIN SmallTable S ON B.ID = S.ID; 
-- Forces B to be scanned first.
```

### Scenario 73: Understanding EXPLAIN 'ref' vs 'eq_ref'
**Question 73:** In `EXPLAIN` output for a join, what is the difference between `ref` and `eq_ref`?
**Logical Reasoning:**
*   `eq_ref`: One row matches (Primary/Unique Key join). Fastest.
*   `ref`: Multiple rows may match (Non-Unique Index or Left Prefix). Fast.
**Join Type:** Performance Analysis.

### Scenario 74: Block Nested Loop (BNL)
**Question 74:** What does "Block Nested Loop" in EXPLAIN mean?
**Logical Reasoning:** It means MySQL is loading a "Block" of the outer table into a buffer (Join Buffer) and scanning the inner table against the buffer, reducing disk scans. Common when joining on non-indexed columns.
**Join Type:** Algorithm.

### Scenario 75: Batched Delete with Join
**Question 75:** You need to delete 1 million unmatched rows. `DELETE ... WHERE NOT EXISTS ...` is locking the table. Solution?
**Logical Reasoning:** Iterate in chunks. Join on Primary Key range.
**Join Type:** Batched Operation.
```sql
-- Application Loop:
DELETE E 
FROM Logs E 
LEFT JOIN KeepList K ON E.ID = K.ID 
WHERE K.ID IS NULL 
LIMIT 1000;
-- Repeat until affected rows = 0
```

### Scenario 76: Refreshing Materialized View (Simulation)
**Question 76:** How strictly to update a reporting table using `INSERT ... ON DUPLICATE KEY UPDATE` with a complex join?
**Logical Reasoning:** Aggregate source, join with destination, upsert results.
**Join Type:** Upsert Join.

### Scenario 77: SCD Type 2 Join (Historical Data)
**Question 77:** Join `Orders` (OrderDate) to `TaxRates` (EffectiveStart, EffectiveEnd) to find the tax rate at that moment.
**Logical Reasoning:** Joining on Date ranges.
**Join Type:** Range Join.
```sql
SELECT O.OrderID, T.Rate
FROM Orders O
JOIN TaxRates T ON O.OrderDate >= T.StartDate AND O.OrderDate < T.EndDate;
```

### Scenario 78: Case Sensitivity in Joins
**Question 78:** Joining `TableA.Code` ('abc') to `TableB.Code` ('ABC'). Matches?
**Logical Reasoning:** Depends on Collation (`_ci` = Case Insensitive, `_bin` = Binary/Sensitive).
**Join Type:** Collation Aware Join.

### Scenario 79: Skewed Data Join (Salting)
**Question 79:** One Key has 1M rows (Data Skew), causing one Reducer/Thread to hang in a distributed/parallel join. Fix?
**Logical Reasoning:** "Salt" the key. Add a random suffix (0-9) to the key in table A, and replicate rows 10 times in table B with suffixes 0-9. (Advanced/Big Data concept, relevant if using MySQL Cluster or mapped to Spark).
**Join Type:** Skew Handling.

### Scenario 80: Map-Reduce via SQL
**Question 80:** How to implement a "Word Count" style logic on 2 joined tables?
**Logical Reasoning:** Explosion Join (Splitting string to rows) -> Group By.
**Join Type:** Transformation Join.



## 🟣 Part 5: Master Class & Theoretical Edge Cases (Questions 81-100)

### Scenario 81: Partition Pruning in Joins
**Question 81:** You join `Sales` (Partitioned by Year) to `DateDim`. How do you ensure the DB only scans the relevant partition?
**Logical Reasoning:** Include the Partition Key in the `ON` clause or `WHERE` clause explicitly.
**Join Type:** Partition Wise Join.
```sql
SELECT * 
FROM Sales S
JOIN DateDim D ON S.SaleDate = D.DateVal
WHERE S.SaleDate BETWEEN '2023-01-01' AND '2023-12-31'; 
-- Optimizer prunes partitions outside 2023.
```

### Scenario 82: Merge Join prerequisites
**Question 82:** When will the optimizer choose a "Merge Join" over a Hash Join?
**Logical Reasoning:** Merge Joins require *both* inputs to be sorted on the join key. This happens if columns are indexed or explicitly sorted. RAM usage is low, but sorting is expensive if not pre-indexed.
**Join Type:** Algorithm.

### Scenario 83: Join Order (Left-Deep vs Bushy Trees)
**Question 83:** Does the order of tables in `FROM` matter?
**Logical Reasoning:** Logically NO (Inner Joins are commutative). Physically YES (Optimizer builds a tree).
*   **Left-Deep**: ((A+B)+C)+D. Standard.
*   **Bushy**: (A+B) + (C+D). Parallelizable.
**Join Type:** Optimizer Logic.

### Scenario 84: Null-Safe Equality (<=>)
**Question 84:** Join two tables allowing `NULL = NULL` to be True.
**Logical Reasoning:** Standard `=` fails on NULL. Use `<=>` (Spaceship operator in MySQL) or `IS NOT DISTINCT FROM` (Postgres).
**Join Type:** Equi-Join with Nulls.
```sql
SELECT A.ID, B.ID
FROM TableA A
JOIN TableB B ON A.Category <=> B.Category;
```

### Scenario 85: Distributed Join (Broadcast vs Shuffle)
**Question 85:** (Big Data/Sharding context) What is a "Broadcast Join"?
**Logical Reasoning:** Copying the entire small table to every node where the large table resides to avoid moving the large table across the network (Shuffling).
**Join Type:** Distributed System Strategy.

### Scenario 86: Joining on Calculated Columns
**Question 86:** `JOIN on YEAR(T1.Date) = T2.Year`. Performance impact?
**Logical Reasoning:** terrible. Calculated columns usually prevent Index usage.
**Fix:** Add a generated column `Year` and index it.
**Join Type:** Anti-pattern.

### Scenario 87: The "Fan Trap"
**Question 87:** You join Master -> Detail_A and Master -> Detail_B. Aggregating A and B gives wrong results. Why?
**Logical Reasoning:** (Same as Join Explosion/Scenario 53). One `Master` row expands to `Count(A) * Count(B)` rows.
**Join Type:** Data Modeling Trap.

### Scenario 88: The "Chasm Trap"
**Question 88:** You join Detail_A -> Master -> Detail_B, but the path from A to B doesn't exist for some rows.
**Logical Reasoning:** Careful usage of Outer Joins is needed. If you Inner Join, you lose partial data from A if it doesn't link all the way to B.
**Join Type:** Data Modeling Trap.

### Scenario 89: Median Calculation via Self Join
**Question 89:** Calculate the Median salary without using `PERCENTILE_CONT`.
**Logical Reasoning:** Self Join to count how many rows are above and below a value.
**Join Type:** Analytical Self Join.
```sql
SELECT AVG(T1.Salary)
FROM Employees T1, Employees T2
GROUP BY T1.Salary
HAVING SUM(CASE WHEN T2.Salary >= T1.Salary THEN 1 ELSE 0 END) >= COUNT(*)/2
   AND SUM(CASE WHEN T2.Salary <= T1.Salary THEN 1 ELSE 0 END) >= COUNT(*)/2;
```

### Scenario 90: Mode Selection (Most Frequent)
**Question 90:** Find the most frequent salary (Mode).
**Logical Reasoning:** Aggregation.
**Join Type:** Aggregation (No Join needed usually, but can be done via Self Join to find ties).
```sql
SELECT Salary 
FROM Employees 
GROUP BY Salary 
ORDER BY COUNT(*) DESC 
LIMIT 1;
```

### Scenario 91: Finding Symmetric Pairs
**Question 91:** Find pairs (X, Y) such that a row (Y, X) also exists, where X != Y.
**Logical Reasoning:** Self Join `ON T1.X = T2.Y AND T1.Y = T2.X`.
**Join Type:** Self Join.
```sql
SELECT T1.X, T1.Y
FROM Coordinates T1
JOIN Coordinates T2 ON T1.X = T2.Y AND T1.Y = T2.X
WHERE T1.X < T1.Y; -- Order to avoid duplicates
```

### Scenario 92: Rolling 3-Month Average
**Question 92:** Calculate average sales for the past 3 months for every month.
**Logical Reasoning:** Self Join on Date Range.
**Join Type:** Range Self Join.
```sql
SELECT T1.Month, AVG(T2.Sales)
FROM MonthlySales T1
JOIN MonthlySales T2 
    ON T2.Month BETWEEN T1.Month - INTERVAL 2 MONTH AND T1.Month
GROUP BY T1.Month;
```

### Scenario 93: Consecutive Wins (Gaps & Islands)
**Question 93:** Find teams that have won 3 or more games in a row.
**Logical Reasoning:** Join `T1`, `T2` (T1+1), `T3` (T1+2).
**Join Type:** Multi Self Join.
```sql
SELECT DISTINCT T1.Team
FROM Games T1
JOIN Games T2 ON T1.Team = T2.Team AND T2.ID = T1.ID + 1
JOIN Games T3 ON T1.Team = T3.Team AND T3.ID = T1.ID + 2
WHERE T1.Result = 'Win' AND T2.Result = 'Win' AND T3.Result = 'Win';
```

### Scenario 94: Nested Set Model Join
**Question 94:** How to find all descendants of a Node in Nested Set Model (Left, Right)?
**Logical Reasoning:** Join where `Descendant.Left` is between `Ancestor.Left` and `Ancestor.Right`.
**Join Type:** Range Join (Hierarchy).
```sql
SELECT Descendant.Name
FROM Categories Ancestor
JOIN Categories Descendant 
    ON Descendant.Lft BETWEEN Ancestor.Lft AND Ancestor.Rgt
WHERE Ancestor.Name = 'Electronics';
```

### Scenario 95: Path Enumeration
**Question 95:** Find all ancestors of 'IPhone' if stored as Path 'Electronics/Phones/IPhone'.
**Logical Reasoning:** Join on `LIKE`.
**Join Type:** String Pattern Join.
```sql
SELECT Ancestor.*
FROM Products P
JOIN Products Ancestor ON P.Path LIKE CONCAT(Ancestor.Path, '%')
WHERE P.Name = 'IPhone';
```

### Scenario 96: Joining Text/Blob Columns
**Question 96:** Why is joining on a `TEXT` column (Description) dangerous?
**Logical Reasoning:** 
1.  **Performance:** No full-width index (usually prefix only).
2.  **Temp Tables:** TEXT columns force implicit temp tables to Disk (vs Memory/RAM) in older MySQL versions, killing performance.
**Join Type:** Performance Pitfall.

### Scenario 97: EXISTS vs IN vs JOIN (2024 Update)
**Question 97:** Is `EXISTS` still faster than `IN`?
**Logical Reasoning:** In modern MySQL (8.0+), `IN` is often optimized as a "Semijoin" just like `EXISTS`. However, `EXISTS` is safer against NULLs. `JOIN` helps if you need columns from the other table.
**Join Type:** Optimization Evolution.

### Scenario 98: One-to-None Join (Anti-Join intention)
**Question 98:** Retrieve Customers who bought product A but NEVER bought product B.
**Logical Reasoning:** Join A (Positive) and Join B (Negative).
**Join Type:** Compound Logic Join.
```sql
SELECT C.Name
FROM Customers C
JOIN Orders O1 ON C.ID = O1.CustID AND O1.ProdID = 'A'
LEFT JOIN Orders O2 ON C.ID = O2.CustID AND O2.ProdID = 'B'
WHERE O2.ID IS NULL;
```

### Scenario 99: "The Monster Query" (10+ Joins)
**Question 99:** You inherit a query with 15 JOINs that takes 10 seconds. First step?
**Logical Reasoning:**
1.  Check `EXPLAIN`.
2.  Identify the "Driving Table" (Start of the chain).
3.  Check if 15 joins are needed (Denormalization opportunity?).
4.  Break into Temp Tables (Divide and conquer).
**Join Type:** Debugging Strategy.

### Scenario 100: Final Exam - The "Super User"
**Question 100:** Select Users who:
1.  Are in 'IT'.
2.  Have > 5 Projects.
3.  Manager is 'Alice'.
4.  Never logged a ticket in 'Q1 2023'.
**Reasoning:** Putting it all together.
```sql
SELECT E.EmpName
FROM Employees E
JOIN Departments D ON E.DeptID = D.DeptID
JOIN Employees M ON E.ManagerID = M.EmpID
-- 2. Projects Aggregation (>5)
JOIN (
    SELECT EmpID FROM Employee_Projects GROUP BY EmpID HAVING COUNT(*) > 5
) EP ON E.EmpID = EP.EmpID
-- 4. Anti-Join Ticket Check
LEFT JOIN Tickets T ON E.EmpID = T.EmpID 
    AND T.Date BETWEEN '2023-01-01' AND '2023-03-31'
WHERE D.DeptName = 'IT'
  AND M.EmpName = 'Alice'
  AND T.TicketID IS NULL;
```

---
**End of 100 SQL Join Questions**
