Question : identify premium customers (avg order value > $500 in Q1 2024) and analyze their spending by product category. 

Definition:

Average order value > $500


SQL Command:

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


Final Product:

answer : 

 with cte as (
    select 
        customer_id,
        avg(total_amount) as avg_order_value
    from orders
    where order_date between '2024-01-01' and '2024-03-31'
    group by customer_id
    having avg(total_amount) > 500
)
select 
    c.customer_name,
    c.country,
    cte.avg_order_value
from customers c
join cte on c.customer_id = cte.customer_id
order by cte.avg_order_value desc;


golang question : 


https://dbfiddle.uk/zzFS6zo2

we have a slice of numbers, we need to find the sum of all the numbers in the slice using concurrency
    numbers := []int{
        1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
        11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
        21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
        31, 32, 33, 34, 35, 36, 37, 38, 39, 40,
        41, 42, 43, 44, 45, 46, 47, 48, 49, 50,
        51, 52, 53, 54, 55, 56, 57, 58, 59, 60,
        61, 62, 63, 64, 65, 66, 67, 68, 69, 70,
        71, 72, 73, 74, 75, 76, 77, 78, 79, 80,
        81, 82, 83, 84, 85, 86, 87, 88, 89, 90,
        91, 92, 93, 94, 95, 96, 97, 98, 99, 100,
    }

answer : 

package main

import (
	"fmt"
	"sync"
)

type writeCounter struct {
	mx    sync.Mutex
	value int
}

func (w *writeCounter) add(val int) {
	w.mx.Lock()         // Lock BEFORE the operation
	defer w.mx.Unlock() // Unlock AFTER the operation
	w.value += val
}

func (w *writeCounter) getValue() int {
	w.mx.Lock() // It is good practice to lock reads if writes are concurrent
	defer w.mx.Unlock()
	return w.value
}

func main() {
	numbers := make([]int, 100)
	for i := 0; i < 100; i++ {
		numbers[i] = i + 1
	}

	wc := writeCounter{}
	var wg sync.WaitGroup

	for _, n := range numbers {
		wg.Add(1)
		// Pass 'n' into the goroutine to avoid closure race conditions
		go func(val int) {
			defer wg.Done() // Signal completion
			wc.add(val)
		}(n)
	}

	wg.Wait()
	fmt.Printf("Sum of all values: %d\n", wc.getValue())
}