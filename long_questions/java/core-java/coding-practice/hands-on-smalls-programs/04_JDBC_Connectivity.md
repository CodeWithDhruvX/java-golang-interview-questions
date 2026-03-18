# Specialized Topic 1: JDBC (Java Database Connectivity)

**Goal**: Learn how to connect Java applications to a Relational Database (MySQL/PostgreSQL/Oracle).

**Prerequisites**:
1.  Database installed (e.g., MySQL).
2.  JDBC Driver jar in classpath (e.g., `mysql-connector-java.jar`).

## 1. 5 Steps of JDBC
1.  Load Driver (`Class.forName`)
2.  Create Connection (`DriverManager.getConnection`)
3.  Create Statement (`conn.createStatement`)
4.  Execute Query (`stmt.executeQuery` / `executeUpdate`)
5.  Close Connection (`conn.close`)

## 2. Basic Connection & Fetch Example

```java
import java.sql.*;

public class JdbcDemo {
    public static void main(String[] args) {
        String url = "jdbc:mysql://localhost:3306/testdb";
        String user = "root";
        String password = "password";

        try {
            // 2. Establish Connection
            Connection conn = DriverManager.getConnection(url, user, password);
            System.out.println("Connected to Database!");

            // 3. Create Statement
            Statement stmt = conn.createStatement();

            // 4. Execute Query (Select)
            ResultSet rs = stmt.executeQuery("SELECT id, name, salary FROM employees");

            // Process Results
            while (rs.next()) {
                int id = rs.getInt("id");
                String name = rs.getString("name");
                double salary = rs.getDouble("salary");
                System.out.println(id + " | " + name + " | " + salary);
            }

            // 5. Close Connection
            conn.close();

        } catch (SQLException e) {
            e.printStackTrace();
        }
    }
}
```

## 3. Prepared Statement (Insert/Update)
**Why?** Prevents SQL Injection and improves performance.

```java
import java.sql.*;
import java.util.Scanner;

public class JdbcInsert {
    public static void main(String[] args) {
        String url = "jdbc:mysql://localhost:3306/testdb";
        String user = "root";
        String pass = "password";
        
        String query = "INSERT INTO employees (name, salary) VALUES (?, ?)";

        try (Connection conn = DriverManager.getConnection(url, user, pass);
             PreparedStatement pstmt = conn.prepareStatement(query)) {
            
            Scanner sc = new Scanner(System.in);
            System.out.print("Enter Name: ");
            String name = sc.next();
            System.out.print("Enter Salary: ");
            double salary = sc.nextDouble();

            // Set parameters (1-based index)
            pstmt.setString(1, name);
            pstmt.setDouble(2, salary);

            int rowsAffected = pstmt.executeUpdate();
            System.out.println(rowsAffected + " row(s) inserted.");

        } catch (SQLException e) {
            e.printStackTrace();
        }
    }
}
```

## Interview Questions on JDBC
1.  **Statement vs PreparedStatement?**
    *   Statement: Compiles SQL every time. Vulnerable to SQL Injection.
    *   PreparedStatement: Pre-compiled. Faster. Safe from SQL Injection.
2.  **What is ResultSet?**
    *   It represents a table of data representing a database result set.
3.  **Class.forName() purpose?**
    *   Used to load the driver class dynamically (Not needed in newer JDBC versions that support SPI).

---

## 📋 Comprehensive Interview Questions

### **JDBC Fundamentals & Architecture Questions**

**Q1: Explain the 5 steps of JDBC in detail.**
**A**: "The 5 steps of JDBC are: 1) Load the driver using `Class.forName()` (though modern JDBC uses SPI), 2) Establish connection with `DriverManager.getConnection()` using database URL, credentials, and properties, 3) Create Statement object with `conn.createStatement()` for SQL execution, 4) Execute queries using `executeQuery()` for SELECT or `executeUpdate()` for INSERT/UPDATE/DELETE, 5) Close the connection and other resources in reverse order of creation to prevent resource leaks. Each step is crucial for proper database interaction."

**Q2: What is the role of DriverManager in JDBC?**
**A**: "DriverManager is the central management class in JDBC that handles driver registration and connection creation. It maintains a list of registered drivers and when I call `getConnection()`, it attempts to connect to the database using each registered driver until one succeeds. It's the factory class that creates actual database connections. In modern applications, connection pools often replace DriverManager for better performance and resource management."

**Q3: What's the difference between `execute()`, `executeQuery()`, and `executeUpdate()`?**
**A**: "`executeQuery()` is used for SELECT statements and returns a ResultSet with the query results. `executeUpdate()` is for INSERT, UPDATE, DELETE operations and returns an integer representing the number of affected rows. `execute()` is more general - it can execute any SQL statement and returns a boolean indicating if the first result is a ResultSet. If it returns true, I can get the ResultSet with `getResultSet()`, if false, I can get the update count with `getUpdateCount()`."

### **PreparedStatement & Security Questions**

**Q4: How does PreparedStatement prevent SQL injection?**
**A**: "PreparedStatement prevents SQL injection by separating the SQL query logic from the parameter values. When I use parameter placeholders like '?', the database driver compiles the SQL template first, then safely binds the parameter values. Even if someone passes malicious input like 'OR 1=1', it's treated as literal data, not executable SQL. The driver automatically escapes special characters, ensuring that parameters can never break out of their intended context. This is much safer than string concatenation."

**Q5: What are the performance benefits of PreparedStatement?**
**A**: "PreparedStatement offers significant performance benefits because the SQL is compiled only once by the database and then reused. When I execute the same PreparedStatement multiple times with different parameters, the database doesn't need to re-parse and re-optimize the SQL each time. The execution plan is cached, making subsequent executions much faster. This is especially beneficial in applications that perform similar queries repeatedly, like inserting multiple records or running parameterized reports."

**Q6: When would you use CallableStatement?**
**A**: "I use CallableStatement when I need to execute stored procedures in the database. Stored procedures are pre-compiled SQL programs that reside on the database server. CallableStatement allows me to call these procedures and handle both input and output parameters. It's useful for complex business logic that's better handled on the database side, for security (encapsulating logic), or when I need to return multiple result sets or output parameters from a single database call."

### **ResultSet & Data Handling Questions**

**Q7: Explain ResultSet types and their characteristics.**
**A**: "ResultSet has different types: `TYPE_FORWARD_ONLY` is the default and only allows forward movement, which is most efficient. `TYPE_SCROLL_INSENSITIVE` allows bidirectional movement but doesn't reflect changes made by others. `TYPE_SCROLL_SENSITIVE` allows scrolling and shows changes made by other users. For concurrency, `CONCUR_READ_ONLY` is default and can't update the database, while `CONCUR_UPDATABLE` allows modifying the database through the ResultSet. I choose based on whether I need scrolling, real-time updates, or update capabilities."

**Q8: What is ResultSetMetaData and when would you use it?**
**A**: "ResultSetMetaData provides information about the ResultSet structure - column names, types, sizes, whether columns are nullable, etc. I use it when writing generic database tools or when I need to dynamically process query results without knowing the column structure beforehand. For example, when building a data export utility or a generic reporting tool that works with any SQL query. It allows me to write flexible code that can adapt to different result sets."

**Q9: How do you handle large result sets efficiently?**
**A**: "For large result sets, I use several techniques: 1) Set appropriate fetch size using `setFetchSize()` to control how many rows are retrieved from the database at once, 2) Use streaming with `TYPE_FORWARD_ONLY` ResultSet and specific fetch size for very large datasets, 3) Process rows in batches rather than loading everything into memory, 4) Consider using LIMIT/OFFSET or pagination for web applications, 5) Use server-side cursors if the database supports them. This prevents memory exhaustion and improves performance."

### **Connection Management & Performance Questions**

**Q10: What are connection pools and why are they important?**
**A**: "Connection pools are caches of database connections that are reused rather than created and destroyed for each database operation. Creating connections is expensive - it involves network overhead, authentication, and database resource allocation. Connection pools maintain a set of ready-to-use connections, dramatically improving performance. They also help manage resource limits by controlling the maximum number of concurrent connections. Popular pools include HikariCP, Apache DBCP, and C3P0. In production applications, connection pooling is essential for good performance."

**Q11: Explain JDBC transaction management.**
**A**: "JDBC transactions allow me to group multiple SQL operations into a single atomic unit. By default, each SQL statement is auto-committed. For transactions, I use `conn.setAutoCommit(false)` to start a transaction, execute multiple statements, then either `conn.commit()` to make changes permanent or `conn.rollback()` to undo them. Transactions ensure ACID properties - Atomicity (all or nothing), Consistency (database remains valid), Isolation (transactions don't interfere), and Durability (committed changes persist)."

**Q12: What are savepoints in JDBC transactions?**
**A**: "Savepoints allow me to create intermediate rollback points within a transaction. I can create a savepoint with `conn.setSavepoint('name')`, and later rollback to that specific point with `conn.rollback(savepoint)` instead of rolling back the entire transaction. This is useful for complex operations where I want to undo only part of the work while keeping other changes. For example, in a batch processing operation where some records fail but I want to keep the successful ones."

### **Error Handling & Best Practices Questions**

**Q13: How do you handle SQLException properly?**
**A**: "For proper SQLException handling, I use specific exception handling strategies: 1) Use try-with-resources to automatically close connections, statements, and result sets, 2) Catch specific SQLException subtypes if available, 3) Use `getNextException()` to handle chained exceptions, 4) Log detailed error information including SQL state and error codes, 5) Implement retry logic for transient failures, 6) Provide meaningful error messages to users while hiding sensitive database details. Proper error handling prevents resource leaks and helps with debugging."

**Q14: What is the JDBC Batch Update feature?**
**A**: "Batch updates allow me to execute multiple SQL statements as a single batch, which is much more efficient than executing statements individually. I use `addBatch()` to add statements to the batch and `executeBatch()` to execute them all at once. This reduces network round trips and database overhead. It's particularly useful for bulk operations like inserting multiple records or updating many rows. The database can optimize the batch execution, resulting in significantly better performance compared to individual statement execution."

**Q15: How do you implement proper resource cleanup in JDBC?**
**A**: "Proper resource cleanup in JDBC follows the reverse order of creation: ResultSet first, then Statement, then Connection. The best approach is using try-with-resources (Java 7+) which automatically closes resources even if exceptions occur. For older Java versions, I use finally blocks to ensure `close()` is called. It's critical to close all resources to prevent memory leaks and connection exhaustion. I also check for null before closing to avoid NullPointerException."

### **Advanced JDBC Features Questions**

**Q16: What is JDBC RowSet and what are its advantages?**
**A**: "RowSet is an enhanced version of ResultSet that adds features like scrollability, updatability, and the ability to operate without a continuous database connection. CachedRowSet can disconnect from the database, work with data offline, then reconnect to sync changes. WebRowSet can serialize to XML. RowSet is JavaBean component, making it easier to use in GUI applications and web services. It's particularly useful for disconnected operations and when I need to pass data between application tiers."

**Q17: Explain the difference between Type 1, Type 2, Type 3, and Type 4 JDBC drivers.**
**A**: "Type 1 drivers are JDBC-ODBC bridge drivers - inefficient and platform-dependent. Type 2 are native API drivers - partially Java, partially native code. Type 3 are network protocol drivers that translate JDBC calls to database-independent network protocol. Type 4 are pure Java drivers that convert JDBC calls directly to database-specific protocol. Type 4 drivers are most common today because they're platform-independent, efficient, and don't require native libraries. Most modern databases provide Type 4 drivers."

**Q18: How do you handle BLOB and CLOB data types in JDBC?**
**A**: "For BLOB (Binary Large Object) and CLOB (Character Large Object) data, I use specific methods: `getBlob()`, `getClob()` for reading, and `setBlob()`, `setClob()` for writing. For large files, I use streams with `getBinaryStream()` and `setBinaryStream()` to avoid loading everything into memory. I can also work with `Blob` and `Clob` objects directly, which provide methods like `length()`, `getBytes()`, and `truncate()`. Proper handling is important to prevent memory issues with large files."

### **Database Integration & Design Questions**

**Q19: How would you implement DAO pattern with JDBC?**
**A**: "I'd implement DAO (Data Access Object) pattern by creating separate DAO classes for each entity, with interfaces defining the contract. Each DAO would handle all database operations for its entity, hiding JDBC complexity from business logic. For example, EmployeeDAO with methods like `findById()`, `save()`, `delete()`, `findAll()`. The DAO would manage connections, prepared statements, and result set mapping. This separation makes code more maintainable, testable, and allows me to switch persistence implementations easily."

**Q20: What are best practices for JDBC in enterprise applications?**
**A**: "Enterprise JDBC best practices include: 1) Always use connection pools, 2) Use PreparedStatement for all parameterized queries, 3) Implement proper transaction management, 4) Use try-with-resources for cleanup, 5) Implement proper exception handling and logging, 6) Use appropriate fetch sizes for large result sets, 7) Consider ORM frameworks like Hibernate for complex applications, 8) Monitor connection usage and performance, 9) Implement connection timeout settings, 10) Use database-specific features judiciously to maintain portability."

**Q21: How do you handle database schema migrations with JDBC?**
**A**: "For schema migrations, I implement version-controlled migration scripts that track which version of the schema the database is currently at. I create a migrations table to track applied migrations, then execute migration scripts in order using JDBC. Each migration script contains the SQL statements needed to upgrade from one version to the next. I handle transactions carefully - each migration should be atomic. Tools like Flyway or Liquibase automate this process, but I can implement a basic version using JDBC directly."

**Q22: What strategies do you use for optimizing JDBC performance?**
**A**: "JDBC performance optimization strategies include: 1) Using connection pools to avoid connection overhead, 2) Using PreparedStatement for query reuse, 3) Setting appropriate fetch sizes for large result sets, 4) Using batch updates for bulk operations, 5) Choosing appropriate isolation levels for transactions, 6) Using stored procedures for complex operations, 7) Implementing proper indexing on the database side, 8) Monitoring and tuning database-specific settings, 9) Using connection validation in pools, 10) Implementing caching for frequently accessed, rarely changing data."
