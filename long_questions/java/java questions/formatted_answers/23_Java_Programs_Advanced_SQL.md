# 23. Java Programs (Advanced APIs & SQL)

**Q: Date and Time API (Java 8)**
> "Stop using `Date` and `Calendar`. They are mutable and not thread-safe.
>
> 1.  **LocalDate**: `LocalDate.now()` (Date only, no time).
> 2.  **LocalDateTime**: `LocalDateTime.now()` (Date and Time).
> 3.  **ZonedDateTime**: `ZonedDateTime.now(ZoneId.of(\"America/New_York\"))`.
>
> To format:
> `DateTimeFormatter formatter = DateTimeFormatter.ofPattern(\"dd-MM-yyyy\");`
> `String formatted = date.format(formatter);`"

**Indepth:**
> **Immutability**: `LocalDate` is immutable (like `String`). Calling `date.plusDays(1)` does *not* change `date`; it returns a *new* object. This makes it thread-safe by default.


---

**Q: Reflection API**
> "Reflection allows code to inspect **itself** at runtime. You can look at a class and ask 'What methods do you have?' or 'What are your private fields?'.
>
> Example:
> `Class<?> clazz = obj.getClass();`
> `Method[] methods = clazz.getDeclaredMethods();`
>
> **Warning**: It breaks encapsulation (you can access private fields) and it is slower than normal method calls. Use it only for frameworks (like Spring) or generic libraries."

**Indepth:**
> **Security**: Reflection can bypass `private` access modifiers using `setAccessible(true)`. This is powerful but dangerous. Only the `SecurityManager` (if installed) can stop this.


---

**Q: ExecutorService (ThreadPool)**
> "Creating a new `Thread` for every task is expensive.
>
> **ExecutorService** manages a pool of threads for you.
> 1.  `ExecutorService executor = Executors.newFixedThreadPool(10);` (Creates 10 threads).
> 2.  `executor.submit(() -> { ... task ... });`
> 3.  `executor.shutdown();` (Stops accepting new tasks and shuts down safely).
>
> It reuses threads, keeping your application stable."

**Indepth:**
> **Types**: `CachedThreadPool` creates threads as needed (good for bursts of short tasks). `FixedThreadPool` has a limit (good for predictable load). `ScheduledThreadPool` is for repeating tasks (cron jobs).


---

**Q: Find Nth Highest Salary (SQL)**
> "The classic interview query.
>
> **Using Limit/Offset (MySQL/PostgreSQL)**:
> `SELECT salary FROM Employee ORDER BY salary DESC LIMIT 1 OFFSET N-1;`
> (To get the 3rd highest, you skip the first 2).
>
> **Generic Standard SQL (Subquery)**:
> `SELECT MAX(salary) FROM Employee WHERE salary < (SELECT MAX(salary) FROM Employee);` (For 2nd highest).
>
> **Modern Way (Window Functions)**:
> `SELECT * FROM (SELECT salary, DENSE_RANK() OVER (ORDER BY salary DESC) as rank FROM Employee) WHERE rank = N;`"

**Indepth:**
> **Dense Rank**: Why `DENSE_RANK`? If two people have the same top salary, `RANK()` skips the next number (1, 1, 3). `DENSE_RANK()` does not skip (1, 1, 2). Usually, "2nd highest" implies distinct values.


---

**Q: Duplicate Records (SQL)**
> "How to find duplicate emails?
> `SELECT email, COUNT(*) FROM Users GROUP BY email HAVING COUNT(*) > 1;`
>
> This groups rows by email and only keeps the groups where the count is greater than 1."

**Indepth:**
> **Logic**: `WHERE` filters rows *before* grouping. `HAVING` filters groups *after* aggregating. You cannot use `COUNT(*)` in a `WHERE` clause.


---

**Q: Count Employees per Department (SQL)**
> "You need to join execution data with department data (if normalized) or just group by department.
>
> `SELECT dept_name, COUNT(emp_id) FROM Employees GROUP BY dept_name;`
>
> If you have a separate `Departments` table:
> `SELECT d.name, COUNT(e.id) FROM Department d LEFT JOIN Employee e ON d.id = e.dept_id GROUP BY d.name;`"

**Indepth:**
> **Left Join**: Why `LEFT JOIN`? If a department has *zero* employees, an `INNER JOIN` would exclude that department from the result. A `LEFT JOIN` keeps the department and reports a count of 0.


---

**Q: Employees with Salary > Manager (SQL)**
> "This requires a **Self Join**. You treat the Employee table as two different tables: one for Employees (`e`) and one for Managers (`m`).
>
> `SELECT e.name FROM Employee e JOIN Employee m ON e.manager_id = m.id WHERE e.salary > m.salary;`"

**Indepth:**
> **Aliases**: In a self-join, aliases (`e` and `m`) are mandatory. The SQL engine needs to know if `salary` refers to the employee's salary or the manager's salary.


---

**Q: Max Salary per Department (SQL)**
> "Two steps:
> 1.  Group by Department and find Max Salary.
> 2.  (Optional) Join back to get the Employee Name.
>
> `SELECT dept_id, MAX(salary) FROM Employee GROUP BY dept_id;`
>
> If you need the *Name* of the person with max salary (Tricky!):
> `SELECT * FROM Employee e WHERE salary = (SELECT MAX(salary) FROM Employee WHERE dept_id = e.dept_id);`"

**Indepth:**
> **Correlated Subquery**: The second query is a correlated subquery. It runs once for *every single row* in the outer table. This is extremely slow O(n*m). Use a Window Function (`RANK()`) or a Join for better performance.


---

**Q: Common Records (Intersection) (SQL)**
> "To find records present in both Table A and Table B:
>
> 1.  **INNER JOIN**: `SELECT * FROM A JOIN B ON A.id = B.id;`
> 2.  **INTERSECT** (If databases supports it): `SELECT id FROM A INTERSECT SELECT id FROM B;`"

**Indepth:**
> **Performance**: `INTERSECT` removes duplicates by default. `INNER JOIN` does not (unless you look distinct). `INTERSECT` is often cleaner to read but sometimes slower depending on the DB optimizer.


---

**Q: Records in A but not B (SQL)**
> "How to find users who signed up but never ordered?
>
> 1.  **LEFT JOIN ... IS NULL**:
>     `SELECT u.name FROM Users u LEFT JOIN Orders o ON u.id = o.user_id WHERE o.id IS NULL;`
>
> 2.  **NOT EXISTS**:
>     `SELECT * FROM Users u WHERE NOT EXISTS (SELECT 1 FROM Orders o WHERE o.user_id = u.id);`"

**Indepth:**
> **Optimization**: `NOT EXISTS` is generally faster than `NOT IN` (especially if columns allow NULLs). A `LEFT JOIN / IS NULL` is often the most optimized by query planners for large datasets.


---

**Q: Copy Table Structure (SQL)**
> "To create a backup table with the same columns but no data:
>
> `CREATE TABLE Backup_Employee AS SELECT * FROM Employee WHERE 1=0;`
>
> The condition `1=0` is always false, so no rows are copied, but the schema (columns/types) is replicated."

**Indepth:**
> **Constraints**: Be careful—`CREATE TABLE AS SELECT` (CTAS) usually copies the column definitions but *skips* indexes, primary keys, and default values. You might need to add constraints manually afterwards.

