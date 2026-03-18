# Comprehensive E-Commerce SQL Interview Questions (Product & Service Based Companies)

## 📌 Rationale
This E-commerce schema is the "Gold Standard" for SQL interviews at product-based companies (Amazon, FAANG, Flipkart, Uber, etc.) and service-based companies alike. It closely mimics a real-world relational database. It perfectly tests Basic Joins, Advanced Window Functions, CTEs, Date Manipulations, and Complex Aggregations.

---

## 🟢 Schema & Sample Data

### 1. Tables Structure
```sql
CREATE TABLE Users (
    UserID INT PRIMARY KEY,
    UserName VARCHAR(50),
    Email VARCHAR(100),
    RegistrationDate DATE
);

CREATE TABLE Categories (
    CategoryID INT PRIMARY KEY,
    CategoryName VARCHAR(50),
    ParentCategoryID INT NULL, -- For testing Self Joins / Recursive CTEs
    FOREIGN KEY (ParentCategoryID) REFERENCES Categories(CategoryID)
);

CREATE TABLE Products (
    ProductID INT PRIMARY KEY,
    ProductName VARCHAR(100),
    CategoryID INT,
    Price DECIMAL(10, 2),
    StockQuantity INT,
    FOREIGN KEY (CategoryID) REFERENCES Categories(CategoryID)
);

CREATE TABLE Orders (
    OrderID INT PRIMARY KEY,
    UserID INT,
    OrderDate DATE,
    TotalAmount DECIMAL(15, 2),
    Status VARCHAR(20), -- 'Completed', 'Cancelled', 'Pending'
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);

CREATE TABLE OrderItems (
    OrderItemID INT PRIMARY KEY,
    OrderID INT,
    ProductID INT,
    Quantity INT,
    UnitPrice DECIMAL(10, 2),
    FOREIGN KEY (OrderID) REFERENCES Orders(OrderID),
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID)
);

CREATE TABLE Reviews (
    ReviewID INT PRIMARY KEY,
    ProductID INT,
    UserID INT,
    Rating INT CHECK(Rating BETWEEN 1 AND 5),
    ReviewDate DATE,
    FOREIGN KEY (ProductID) REFERENCES Products(ProductID),
    FOREIGN KEY (UserID) REFERENCES Users(UserID)
);
```

### 2. Sample Data
```sql
-- Users
INSERT INTO Users VALUES 
(1, 'Alice', 'alice@test.com', '2023-01-10'),
(2, 'Bob', 'bob@test.com', '2023-02-15'),
(3, 'Charlie', 'charlie@test.com', '2023-03-20'),
(4, 'David', 'david@test.com', '2023-04-05'),
(5, 'Eve', 'eve@test.com', '2023-05-10');

-- Categories (Electronics -> Mobile | Laptops)
INSERT INTO Categories VALUES 
(1, 'Electronics', NULL),
(2, 'Mobiles', 1),
(3, 'Laptops', 1),
(4, 'Clothing', NULL),
(5, 'Men Clothing', 4);

-- Products
INSERT INTO Products VALUES 
(101, 'iPhone 14', 2, 999.00, 50),
(102, 'Samsung S23', 2, 899.00, 40),
(103, 'MacBook Pro', 3, 1999.00, 20),
(104, 'Dell XPS', 3, 1499.00, 30),
(105, 'T-Shirt', 5, 20.00, 200),
(106, 'Jeans', 5, 50.00, 100);

-- Orders
INSERT INTO Orders VALUES 
(1001, 1, '2023-06-01', 1019.00, 'Completed'),
(1002, 2, '2023-06-02', 1999.00, 'Completed'),
(1003, 1, '2023-06-05', 50.00, 'Completed'),
(1004, 3, '2023-06-10', 0.00, 'Cancelled'),
(1005, 4, '2023-07-01', 949.00, 'Completed'),
(1006, 1, '2023-07-02', 20.00, 'Completed');

-- Order Items
INSERT INTO OrderItems VALUES 
(1, 1001, 101, 1, 999.00),
(2, 1001, 105, 1, 20.00),
(3, 1002, 103, 1, 1999.00),
(4, 1003, 106, 1, 50.00),
(5, 1005, 102, 1, 899.00),
(6, 1005, 106, 1, 50.00),
(7, 1006, 105, 1, 20.00);

-- Reviews
INSERT INTO Reviews VALUES 
(1, 101, 1, 5, '2023-06-10'),
(2, 103, 2, 4, '2023-06-15'),
(3, 105, 1, 3, '2023-06-20'),
(4, 101, 4, 1, '2023-07-05');
```

---

## � Part 1: Basic Questions (Data Extraction & Filtering)
*Focus: SELECT, WHERE, ORDER BY, LIMIT, and Basic Functions.*

### 1. Retrieve Recent User Registrations
**Question:** Retrieve the details of all users who registered in the year 2023, ordered by their registration date from newest to oldest.
**Logic:** Use `WHERE` clause with the `YEAR()` function and `ORDER BY DESC`.
```sql
SELECT UserID, UserName, Email, RegistrationDate
FROM Users
WHERE YEAR(RegistrationDate) = 2023
ORDER BY RegistrationDate DESC;
```

### 2. Products Out of Stock
**Question:** Find the names and prices of all products that are currently out of stock (StockQuantity is 0).
**Logic:** A simple `WHERE` condition on the `StockQuantity` column.
```sql
SELECT ProductName, Price
FROM Products
WHERE StockQuantity = 0;
```

### 3. High-Value Completed Orders
**Question:** List all completed orders where the total amount exceeds $1000.
**Logic:** Filter by `Status` and `TotalAmount`.
```sql
SELECT OrderID, UserID, OrderDate, TotalAmount
FROM Orders
WHERE Status = 'Completed' AND TotalAmount > 1000.00;
```

### 4. Find Users by Email Domain
**Question:** Retrieve all users who have registered with a `@test.com` email address.
**Logic:** Use `LIKE` operator to pattern match the email domain.
```sql
SELECT UserID, UserName, Email
FROM Users
WHERE Email LIKE '%@test.com';
```

### 5. Count Total Number of Products
**Question:** Find the total number of products available in the catalog.
**Logic:** Use the `COUNT()` aggregate function.
```sql
SELECT COUNT(ProductID) AS TotalProducts
FROM Products;
```

### 6. Find Unique Order Statuses
**Question:** List all the different statuses an order can have, without duplicates.
**Logic:** Use the `DISTINCT` keyword.
```sql
SELECT DISTINCT Status 
FROM Orders;
```

### 7. Products in a Price Range
**Question:** Find the names and prices of products priced between $50 and $1000.
**Logic:** Use the `BETWEEN` operator.
```sql
SELECT ProductName, Price
FROM Products
WHERE Price BETWEEN 50.00 AND 1000.00;
```

### 8. Find the Most Expensive Product
**Question:** Retrieve the details of the most expensive product in the store.
**Logic:** Order by Price descending and use `LIMIT 1`.
```sql
SELECT ProductID, ProductName, Price
FROM Products
ORDER BY Price DESC
LIMIT 1;
```

### 9. Top 3 Lowest Stock Products
**Question:** Find the names and stock quantities of the 3 products with the lowest stock remaining.
**Logic:** Order by StockQuantity ascending and apply `LIMIT 3`.
```sql
SELECT ProductName, StockQuantity
FROM Products
ORDER BY StockQuantity ASC
LIMIT 3;
```

### 10. Users Registered in a Specific Month
**Question:** Retrieve all users who registered in the month of March 2023.
**Logic:** Use `MONTH()` and `YEAR()` functions.
```sql
SELECT UserName, RegistrationDate
FROM Users
WHERE MONTH(RegistrationDate) = 3 AND YEAR(RegistrationDate) = 2023;
```

### 11. Total Revenue from Completed Orders
**Question:** Calculate the total overall revenue from all completed orders.
**Logic:** Use the `SUM()` function and a `WHERE` clause.
```sql
SELECT SUM(TotalAmount) AS TotalRevenue
FROM Orders
WHERE Status = 'Completed';
```

---

## �🟡 Part 2: Intermediate Questions (Joins & Group By)
*Focus: Data aggregation, filtering, and multi-table joins.*

### 12. Total Revenue by Category
**Question:** Calculate the total revenue generated by each category. Include categories that have generated zero revenue.
**Logic:** Need to join Categories -> Products -> OrderItems. Use `LEFT JOIN` to keep categories with no sales.
```sql
SELECT C.CategoryName, COALESCE(SUM(OI.Quantity * OI.UnitPrice), 0) AS TotalRevenue
FROM Categories C
LEFT JOIN Products P ON C.CategoryID = P.CategoryID
LEFT JOIN OrderItems OI ON P.ProductID = OI.ProductID
GROUP BY C.CategoryID, C.CategoryName
ORDER BY TotalRevenue DESC;
```

### 13. Users Who Bought from Multiple Categories
**Question:** Find users who have purchased products from 2 or more distinct categories.
**Logic:** Join Users -> Orders -> OrderItems -> Products -> Categories. Use `COUNT(DISTINCT CategoryID)`.
```sql
SELECT U.UserName, COUNT(DISTINCT P.CategoryID) AS UniqueCategories
FROM Users U
JOIN Orders O ON U.UserID = O.UserID
JOIN OrderItems OI ON O.OrderID = OI.OrderID
JOIN Products P ON OI.ProductID = P.ProductID
WHERE O.Status = 'Completed'
GROUP BY U.UserID, U.UserName
HAVING COUNT(DISTINCT P.CategoryID) >= 2;
```

### 14. Idle Users (Anti-Join)
**Question:** Find users who registered but never placed an order.
**Logic:** `LEFT JOIN` Users with Orders and filter where `OrderID IS NULL`.
```sql
SELECT U.UserName, U.Email
FROM Users U
LEFT JOIN Orders O ON U.UserID = O.UserID
WHERE O.OrderID IS NULL;
```

---

## 🟠 Part 3: Intermediate to Advanced Questions (Window Functions & CTEs)
*Focus: Frequently asked in FAANG/Top Tier tech companies.*

### 15. Nth Highest Purchase 
**Question:** Find the 2nd highest completed order amount for each user. If a user only has 1 order, don't show them.
**Logic:** Use `DENSE_RANK()` partitioning by UserID, ordering by TotalAmount descending.
```sql
WITH RankedOrders AS (
    SELECT UserID, OrderID, TotalAmount,
           DENSE_RANK() OVER(PARTITION BY UserID ORDER BY TotalAmount DESC) as Rnk
    FROM Orders
    WHERE Status = 'Completed'
)
SELECT U.UserName, R.TotalAmount
FROM RankedOrders R
JOIN Users U ON R.UserID = U.UserID
WHERE R.Rnk = 2;
```

### 16. Top 2 Selling Products Per Category
**Question:** Write a query to find the top 2 best-selling products (by total quantity sold) in each category.
**Logic:** Pre-aggregate quantity by product, then use `ROW_NUMBER()` or `RANK()` over Category.
```sql
WITH ProductSales AS (
    SELECT P.CategoryID, P.ProductID, P.ProductName, SUM(OI.Quantity) AS TotalQty
    FROM Products P
    JOIN OrderItems OI ON P.ProductID = OI.ProductID
    GROUP BY P.CategoryID, P.ProductID, P.ProductName
),
RankedSales AS (
    SELECT CategoryID, ProductName, TotalQty,
           DENSE_RANK() OVER(PARTITION BY CategoryID ORDER BY TotalQty DESC) as Rnk
    FROM ProductSales
)
SELECT C.CategoryName, R.ProductName, R.TotalQty
FROM RankedSales R
JOIN Categories C ON R.CategoryID = C.CategoryID
WHERE R.Rnk <= 2;
```

### 17. Cumulative/Running Total Revenue
**Question:** Calculate the daily running total of revenue across the platform for the month of June 2023.
**Logic:** Use `SUM()` as a Window Function over `OrderDate`.
```sql
SELECT OrderDate, 
       SUM(TotalAmount) AS DailyRevenue,
       SUM(SUM(TotalAmount)) OVER(ORDER BY OrderDate) AS RunningTotalRevenue
FROM Orders
WHERE Status = 'Completed' AND MONTH(OrderDate) = 6 AND YEAR(OrderDate) = 2023
GROUP BY OrderDate;
```

---

## 🔴 Part 4: Advanced to Expert Questions (Self Joins, Dates, Recursion)
*Focus: Very complex logic, measuring edge cases and deep SQL understanding.*

### 18. Recursive CTE: Category Hierarchy
**Question:** Given CategoryID = 1 ('Electronics'), list all sub-categories and their depth levels.
**Logic:** Use `WITH RECURSIVE` to traverse the parent-child relationship.
```sql
WITH RECURSIVE CategoryTree AS (
    -- Anchor member: Start with base category
    SELECT CategoryID, CategoryName, ParentCategoryID, 1 AS Level
    FROM Categories 
    WHERE CategoryID = 1
    
    UNION ALL
    
    -- Recursive member: Find children of the current level
    SELECT C.CategoryID, C.CategoryName, C.ParentCategoryID, CT.Level + 1
    FROM Categories C
    JOIN CategoryTree CT ON C.ParentCategoryID = CT.CategoryID
)
SELECT * FROM CategoryTree;
```

### 19. Month-over-Month (MoM) Growth
**Question:** Calculate the Month-over-Month percentage growth in revenue.
**Logic:** Extract Monthly Revenue, then use `LAG()` to get the previous month's revenue to calculate `(Current - Previous) / Previous * 100`.
```sql
WITH MonthlySales AS (
    SELECT DATE_FORMAT(OrderDate, '%Y-%m') AS Month, 
           SUM(TotalAmount) AS Revenue
    FROM Orders
    WHERE Status = 'Completed'
    GROUP BY DATE_FORMAT(OrderDate, '%Y-%m')
)
SELECT Month,
       Revenue,
       LAG(Revenue) OVER(ORDER BY Month) AS PrevMonthRev,
       ROUND(((Revenue - LAG(Revenue) OVER(ORDER BY Month)) / LAG(Revenue) OVER(ORDER BY Month)) * 100, 2) AS MoM_Growth_Percentage
FROM MonthlySales;
```

### 20. Consecutive Days Orders (Self Join / Window)
**Question:** Find users who have placed an order on 2 consecutive days.
**Logic:** Use `LEAD()` to look at the next order date for the user, and use `DATEDIFF()` to check if it's exactly 1 day.
```sql
WITH UserOrderDates AS (
    -- Get distinct dates a user ordered (ignoring multiple orders on same day)
    SELECT DISTINCT UserID, OrderDate
    FROM Orders
    WHERE Status = 'Completed'
),
NextOrderData AS (
    SELECT UserID, OrderDate,
           LEAD(OrderDate) OVER(PARTITION BY UserID ORDER BY OrderDate) AS NextOrderDate
    FROM UserOrderDates
)
SELECT DISTINCT U.UserName
FROM NextOrderData N
JOIN Users U ON N.UserID = U.UserID
WHERE DATEDIFF(N.NextOrderDate, N.OrderDate) = 1;
```

### 21. First Time vs Repeat Purchases (Cohort / Retention)
**Question:** What percentage of users who made their first purchase in June 2023 made a second purchase in July 2023?
**Logic:** Find the "First Purchase Date" for each user. Then check if they possess an order in the subsequent month.
```sql
WITH FirstPurchase AS (
    SELECT UserID, MIN(OrderDate) AS FirstOrderDate
    FROM Orders
    WHERE Status = 'Completed'
    GROUP BY UserID
),
JuneCohorts AS (
    -- Users whose FIRST ever purchase was in June
    SELECT UserID FROM FirstPurchase 
    WHERE MONTH(FirstOrderDate) = 6 AND YEAR(FirstOrderDate) = 2023
),
JulyRetained AS (
    -- Out of June Cohorts, who bought in July?
    SELECT DISTINCT O.UserID 
    FROM Orders O
    JOIN JuneCohorts JC ON O.UserID = JC.UserID
    WHERE MONTH(O.OrderDate) = 7 AND YEAR(O.OrderDate) = 2023 AND O.Status = 'Completed'
)
SELECT 
    (SELECT COUNT(*) FROM JuneCohorts) AS June_New_Users,
    (SELECT COUNT(*) FROM JulyRetained) AS Retained_in_July,
    ROUND((SELECT COUNT(*) FROM JulyRetained) / CAST((SELECT COUNT(*) FROM JuneCohorts) AS DECIMAL(10,2)) * 100, 2) AS Retention_Rate;
```
