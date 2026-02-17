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
