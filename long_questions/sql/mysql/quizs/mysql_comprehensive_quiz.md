# MySQL Comprehensive Analytical Quiz

## 1. Schema Setup

Execute the following SQL commands to set up the environment for the quiz.

```sql
-- Create Tables
CREATE TABLE customers (
    customer_id INT PRIMARY KEY,
    customer_name VARCHAR(50),
    country VARCHAR(50)
);

CREATE TABLE orders (
    order_id INT PRIMARY KEY,
    customer_id INT,
    order_date DATE,
    total_amount DECIMAL(10,2)
);

CREATE TABLE order_items (
    order_id INT,
    product_category VARCHAR(50),
    quantity INT,
    price DECIMAL(10,2)
);

-- Insert Data
INSERT INTO customers VALUES
(1, 'John Smith', 'USA'), (2, 'Maria Garcia', 'Spain'), 
(3, 'Chen Wei', 'China'), (4, 'Sarah Johnson', 'USA'), 
(5, 'David Brown', 'UK');

INSERT INTO orders VALUES
(101, 1, '2024-01-15', 1200.00), (102, 1, '2024-02-20', 800.00),
(103, 2, '2024-01-10', 300.00), (104, 2, '2024-03-05', 1500.00),
(105, 3, '2024-02-14', 2500.00), (106, 4, '2024-01-22', 450.00),
(107, 5, '2024-03-18', 1800.00), (108, 1, '2024-03-25', 950.00);

INSERT INTO order_items VALUES
(101, 'Electronics', 1, 1000.00), (101, 'Accessories', 2, 100.00),
(102, 'Home Appliances', 1, 800.00), (103, 'Clothing', 3, 100.00),
(104, 'Electronics', 1, 1500.00), (105, 'Electronics', 2, 1250.00),
(106, 'Books', 5, 90.00), (107, 'Home Appliances', 1, 1800.00),
(108, 'Electronics', 1, 950.00);
```

---

## 2. Analytical Interview Questions

### Question 1: Premium Customer Analysis
**Objective:** Identify "Premium Customers" defined as those who have an **average order value greater than $500** during **Q1 2024** (Jan-Mar). For these customers, analyze their spending habits by identifying which **product category** contributed most to their spending.

**Output Columns:** `customer_name`, `country`, `avg_order_value`, `top_category`

---

### Question 2: Monthly Revenue Growth
**Objective:** Calculate the **month-over-month percentage growth** in total revenue for 2024.

**Output Columns:** `month`, `total_revenue`, `previous_month_revenue`, `growth_percentage`

---

### Question 3: Regional Category Preferences
**Objective:** For each country, determine the **best-selling product category** based on total sales amount.

**Output Columns:** `country`, `top_category`, `total_sales`

---

### Question 4: Customer Retention Risk (Gap Analysis)
**Objective:** Identify customers who haven't placed an order in **February 2024**, despite having placed an order in January 2024. These might be at risk of churning.

**Output Columns:** `customer_name`, `jan_order_date`

---

### Question 5: Order Value Segmentation
**Objective:** Categorize all orders into three segments:
- **Low**: < $500
- **Medium**: $500 - $1500
- **High**: > $1500
Then count the number of orders in each segment.

**Output Columns:** `order_segment`, `order_count`

---

### Question 6: Running Total Earnings
**Objective:** Calculate the cumulative running total of revenue day by day for the entire dataset, ordered by date.

**Output Columns:** `order_date`, `daily_revenue`, `running_total`

---

## 3. Answer Key

### Answer 1: Premium Customer Analysis
```sql
WITH PremiumCustomers AS (
    SELECT 
        o.customer_id,
        AVG(o.total_amount) as avg_val
    FROM orders o
    WHERE o.order_date BETWEEN '2024-01-01' AND '2024-03-31'
    GROUP BY o.customer_id
    HAVING AVG(o.total_amount) > 500
),
CategorySpending AS (
    SELECT 
        c.customer_id,
        c.customer_name,
        c.country,
        oi.product_category,
        SUM(oi.price * oi.quantity) as category_total,
        pc.avg_val,
        RANK() OVER (PARTITION BY c.customer_id ORDER BY SUM(oi.price * oi.quantity) DESC) as rn
    FROM customers c
    JOIN PremiumCustomers pc ON c.customer_id = pc.customer_id
    JOIN orders o ON c.customer_id = o.customer_id
    JOIN order_items oi ON o.order_id = oi.order_id
    GROUP BY c.customer_id, c.customer_name, c.country, oi.product_category, pc.avg_val
)
SELECT 
    customer_name,
    country,
    ROUND(avg_val, 2) as avg_order_value,
    product_category as top_category
FROM CategorySpending
WHERE rn = 1;
```

**Explanation:**
1.  **CTE `PremiumCustomers`**: Filters for orders in Q1 2024 and groups by `customer_id` to find those with `AVG(total_amount) > 500`.
2.  **CTE `CategorySpending`**: Joins the premium customers back to orders and items.
3.  **`RANK()`**: Ranks product categories for each customer based on total spending (`SUM(price * quantity)`). The category with the highest spend gets rank 1.
4.  **Final Select**: Filters for `rn = 1` to get the top category for each premium customer.


### Answer 2: Monthly Revenue Growth
```sql
WITH MonthlyStats AS (
    SELECT 
        DATE_FORMAT(order_date, '%Y-%m') as order_month,
        SUM(total_amount) as revenue
    FROM orders
    GROUP BY DATE_FORMAT(order_date, '%Y-%m')
)
SELECT 
    order_month,
    revenue,
    LAG(revenue) OVER (ORDER BY order_month) as prev_revenue,
    ROUND(((revenue - LAG(revenue) OVER (ORDER BY order_month)) / LAG(revenue) OVER (ORDER BY order_month)) * 100, 2) as growth_pct
FROM MonthlyStats;
```

**Explanation:**
1.  **CTE `MonthlyStats`**: Aggregates total revenue by month using `DATE_FORMAT(order_date, '%Y-%m')`.
2.  **`LAG(revenue)`**: This window function retrieves the revenue from the *previous* row (month) to allow comparison.
3.  **Growth Calculation**: The formula `((Current - Previous) / Previous) * 100` computes the percentage growth.


### Answer 3: Regional Category Preferences
```sql
WITH RegionalSales AS (
    SELECT 
        c.country,
        oi.product_category,
        SUM(oi.price * oi.quantity) as total_sales,
        RANK() OVER (PARTITION BY c.country ORDER BY SUM(oi.price * oi.quantity) DESC) as rn
    FROM customers c
    JOIN orders o ON c.customer_id = o.customer_id
    JOIN order_items oi ON o.order_id = oi.order_id
    GROUP BY c.country, oi.product_category
)
SELECT 
    country, 
    product_category, 
    total_sales
FROM RegionalSales
WHERE rn = 1;
```

**Explanation:**
1.  **CTE `RegionalSales`**: Aggregates sales by `country` and `product_category`.
2.  **`RANK()` Window Function**: Partitions by `country` and orders by sales descending. This assigns a rank of 1 to the best-selling category in each country.
3.  **Filtering**: The final query selects only the rows where `rn = 1`.


### Answer 4: Customer Retention Risk
```sql
SELECT c.customer_name
FROM customers c
JOIN orders o1 ON c.customer_id = o1.customer_id
WHERE DATE_FORMAT(o1.order_date, '%Y-%m') = '2024-01'
AND c.customer_id NOT IN (
    SELECT o2.customer_id 
    FROM orders o2 
    WHERE DATE_FORMAT(o2.order_date, '%Y-%m') = '2024-02'
);
```

**Explanation:**
1.  **Identify Jan Shoppers**: The main query selects customers who made orders in '2024-01'.
2.  **`NOT IN` Subquery**: It filters out customers who *also* appear in the list of customers who made orders in '2024-02'.
3.  **Result**: Returns customers active in Jan but inactive in Feb. *Note: A `LEFT JOIN` ... `WHERE o2.customer_id IS NULL` approach is also valid and often more performant.*


### Answer 5: Order Value Segmentation
```sql
SELECT 
    CASE 
        WHEN total_amount < 500 THEN 'Low'
        WHEN total_amount BETWEEN 500 AND 1500 THEN 'Medium'
        ELSE 'High'
    END as order_segment,
    COUNT(*) as order_count
FROM orders
GROUP BY order_segment;
```

**Explanation:**
1.  **`CASE` Statement**: Iterates through each order's `total_amount` to assign a label ('Low', 'Medium', 'High').
2.  **`GROUP BY`**: Groups the result set by these newly created labels.
3.  **`COUNT(*)`**: Counts the number of orders falling into each segment.


### Answer 6: Running Total Earnings
```sql
SELECT 
    order_date,
    SUM(total_amount) as daily_revenue,
    SUM(SUM(total_amount)) OVER (ORDER BY order_date) as running_total
FROM orders
GROUP BY order_date;
```

**Explanation:**
1.  **Daily Aggregation**: First, we group by `order_date` to get `daily_revenue`.
2.  **`SUM() OVER()`**: This window function calculates a running total.
3.  **`ORDER BY order_date`**: Inside the window function, this ensures the sum accumulates chronologically from the first date to the last.

