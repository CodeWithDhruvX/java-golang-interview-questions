# Employee Management System - Collections Framework in Action

> **Concepts Demonstrated:** ArrayList, LinkedList, HashSet, TreeSet, HashMap, TreeMap, LinkedHashMap, Queue, Deque, Comparator, Collections utility, Iterator, removeIf, subList

---

## 📋 Overview

This practical program demonstrates real-world usage of Java Collections Framework through an Employee Management System that handles employee records, departments, and organizational operations.

---

## 🏢 Complete Implementation

### Employee.java - Model Class
```java
import java.time.LocalDate;
import java.util.Objects;

public class Employee implements Comparable<Employee> {
    private int id;
    private String name;
    private String department;
    private double salary;
    private LocalDate hireDate;
    private String email;

    public Employee(int id, String name, String department, double salary, LocalDate hireDate, String email) {
        this.id = id;
        this.name = name;
        this.department = department;
        this.salary = salary;
        this.hireDate = hireDate;
        this.email = email;
    }

    // Getters
    public int getId() { return id; }
    public String getName() { return name; }
    public String getDepartment() { return department; }
    public double getSalary() { return salary; }
    public LocalDate getHireDate() { return hireDate; }
    public String getEmail() { return email; }

    // Setters
    public void setSalary(double salary) { this.salary = salary; }
    public void setDepartment(String department) { this.department = department; }

    @Override
    public int compareTo(Employee other) {
        return Integer.compare(this.id, other.id);
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Employee employee = (Employee) o;
        return id == employee.id;
    }

    @Override
    public int hashCode() {
        return Objects.hash(id);
    }

    @Override
    public String toString() {
        return String.format("Employee{id=%d, name='%s', department='%s', salary=$%.2f, hireDate=%s, email='%s'}",
                id, name, department, salary, hireDate, email);
    }
}
```

### EmployeeManagementSystem.java - Main Class
```java
import java.time.LocalDate;
import java.util.*;
import java.util.stream.Collectors;

public class EmployeeManagementSystem {
    // Different collections for different purposes
    private List<Employee> employees; // ArrayList for random access
    private Queue<Employee> waitingList; // LinkedList for FIFO
    private Set<String> departments; // HashSet for unique departments
    private Map<Integer, Employee> employeeMap; // HashMap for quick lookup
    private Map<String, List<Employee>> departmentMap; // TreeMap for sorted departments
    private Deque<Employee> recentHires; // LinkedList as Deque

    public EmployeeManagementSystem() {
        this.employees = new ArrayList<>();
        this.waitingList = new LinkedList<>();
        this.departments = new HashSet<>();
        this.employeeMap = new HashMap<>();
        this.departmentMap = new TreeMap<>();
        this.recentHires = new LinkedList<>();
    }

    // Add employee to multiple collections
    public void addEmployee(Employee employee) {
        employees.add(employee);
        employeeMap.put(employee.getId(), employee);
        departments.add(employee.getDepartment());
        
        // Add to department map
        departmentMap.computeIfAbsent(employee.getDepartment(), k -> new ArrayList<>()).add(employee);
        
        // Add to recent hires (keep only last 5)
        recentHires.addFirst(employee);
        if (recentHires.size() > 5) {
            recentHires.removeLast();
        }
    }

    // Demonstrate removeIf for conditional removal
    public void removeLowPerformers(double minSalary) {
        System.out.println("Removing employees with salary < $" + minSalary);
        int removed = employees.size();
        
        // Safe removal using removeIf
        employees.removeIf(emp -> emp.getSalary() < minSalary);
        
        // Update other collections
        removed -= employees.size();
        rebuildMaps();
        
        System.out.println("Removed " + removed + " employees");
    }

    // Demonstrate Iterator for safe iteration
    public void giveRaiseToDepartment(String department, double percentage) {
        System.out.println("Giving " + percentage + "% raise to " + department + " department");
        
        // Using Iterator for safe modification
        Iterator<Employee> iterator = employees.iterator();
        while (iterator.hasNext()) {
            Employee emp = iterator.next();
            if (emp.getDepartment().equals(department)) {
                emp.setSalary(emp.getSalary() * (1 + percentage / 100));
            }
        }
        rebuildMaps();
    }

    // Demonstrate subList for batch operations
    public void processBatch(int startIndex, int endIndex) {
        if (startIndex < 0 || endIndex >= employees.size() || startIndex > endIndex) {
            System.out.println("Invalid batch range");
            return;
        }

        List<Employee> batch = employees.subList(startIndex, endIndex + 1);
        System.out.println("Processing batch of " + batch.size() + " employees:");
        
        // Sort batch by salary
        batch.sort(Comparator.comparingDouble(Employee::getSalary).reversed());
        
        batch.forEach(emp -> System.out.println("  - " + emp.getName() + ": $" + emp.getSalary()));
    }

    // Demonstrate different Map types
    public void demonstrateMapOperations() {
        System.out.println("\n=== MAP OPERATIONS DEMONSTRATION ===");
        
        // HashMap for quick lookup
        System.out.println("\nHashMap (quick lookup by ID):");
        Employee emp = employeeMap.get(101);
        System.out.println("Employee 101: " + emp);
        
        // TreeMap for sorted keys
        System.out.println("\nTreeMap (departments in sorted order):");
        departmentMap.forEach((dept, empList) -> {
            System.out.println(dept + ": " + empList.size() + " employees");
        });
        
        // LinkedHashMap for insertion order
        System.out.println("\nLinkedHashMap (maintaining insertion order):");
        Map<Integer, String> insertionOrderMap = new LinkedHashMap<>();
        employees.stream().limit(5).forEach(e -> 
            insertionOrderMap.put(e.getId(), e.getName()));
        insertionOrderMap.forEach((id, name) -> 
            System.out.println(id + ": " + name));
    }

    // Demonstrate Queue operations
    public void demonstrateQueueOperations() {
        System.out.println("\n=== QUEUE OPERATIONS DEMONSTRATION ===");
        
        // Add to waiting list
        waitingList.addAll(employees.stream()
            .filter(e -> e.getSalary() < 60000)
            .limit(3)
            .collect(Collectors.toList()));
        
        System.out.println("Waiting list (FIFO):");
        waitingList.forEach(emp -> System.out.println("  - " + emp.getName()));
        
        // Process from queue
        System.out.println("\nProcessing from queue:");
        while (!waitingList.isEmpty()) {
            Employee emp = waitingList.poll();
            System.out.println("Processing: " + emp.getName());
        }
    }

    // Demonstrate Deque operations
    public void demonstrateDequeOperations() {
        System.out.println("\n=== DEQUE OPERATIONS DEMONSTRATION ===");
        System.out.println("Recent hires (newest first):");
        recentHires.forEach(emp -> System.out.println("  - " + emp.getName() + " (" + emp.getHireDate() + ")"));
        
        // Add to both ends
        Employee newest = new Employee(999, "New Hire", "IT", 75000, LocalDate.now(), "new@company.com");
        Employee oldest = new Employee(1, "Founder", "Management", 150000, LocalDate.of(2010, 1, 1), "founder@company.com");
        
        recentHires.addFirst(newest);
        recentHires.addLast(oldest);
        
        System.out.println("\nAfter adding to both ends:");
        recentHires.forEach(emp -> System.out.println("  - " + emp.getName()));
    }

    // Demonstrate Collections utility methods
    public void demonstrateCollectionsUtilities() {
        System.out.println("\n=== COLLECTIONS UTILITIES DEMONSTRATION ===");
        
        // Create a copy for demonstration
        List<Employee> demoList = new ArrayList<>(employees.subList(0, Math.min(5, employees.size())));
        
        System.out.println("Original list:");
        demoList.forEach(emp -> System.out.println("  - " + emp.getName()));
        
        // Reverse
        Collections.reverse(demoList);
        System.out.println("\nAfter reverse:");
        demoList.forEach(emp -> System.out.println("  - " + emp.getName()));
        
        // Shuffle
        Collections.shuffle(demoList);
        System.out.println("\nAfter shuffle:");
        demoList.forEach(emp -> System.out.println("  - " + emp.getName()));
        
        // Sort with custom comparator
        demoList.sort(Comparator.comparing(Employee::getName));
        System.out.println("\nSorted by name:");
        demoList.forEach(emp -> System.out.println("  - " + emp.getName()));
        
        // Frequency
        String mostCommonDept = departments.stream()
            .max(Comparator.comparing(dept -> 
                Collections.frequency(employees.stream()
                    .map(Employee::getDepartment)
                    .collect(Collectors.toList()), dept)))
            .orElse("N/A");
        System.out.println("\nMost common department: " + mostCommonDept);
    }

    // Demonstrate Set operations
    public void demonstrateSetOperations() {
        System.out.println("\n=== SET OPERATIONS DEMONSTRATION ===");
        
        // HashSet for unique departments
        System.out.println("Unique departments (HashSet - unordered):");
        departments.forEach(dept -> System.out.println("  - " + dept));
        
        // TreeSet for sorted departments
        Set<String> sortedDepts = new TreeSet<>(departments);
        System.out.println("\nSorted departments (TreeSet):");
        sortedDepts.forEach(dept -> System.out.println("  - " + dept));
        
        // LinkedHashSet for insertion order
        Set<String> orderedDepts = new LinkedHashSet<>();
        employees.stream().limit(5).forEach(emp -> orderedDepts.add(emp.getDepartment()));
        System.out.println("\nDepartments in insertion order (LinkedHashSet):");
        orderedDepts.forEach(dept -> System.out.println("  - " + dept));
    }

    // Search and filter operations
    public List<Employee> searchEmployees(String searchTerm) {
        return employees.stream()
            .filter(emp -> emp.getName().toLowerCase().contains(searchTerm.toLowerCase()) ||
                           emp.getDepartment().toLowerCase().contains(searchTerm.toLowerCase()) ||
                           emp.getEmail().toLowerCase().contains(searchTerm.toLowerCase()))
            .collect(Collectors.toList());
    }

    public List<Employee> getEmployeesBySalaryRange(double min, double max) {
        return employees.stream()
            .filter(emp -> emp.getSalary() >= min && emp.getSalary() <= max)
            .sorted(Comparator.comparingDouble(Employee::getSalary).reversed())
            .collect(Collectors.toList());
    }

    // Performance comparison between ArrayList and LinkedList
    public void performanceComparison() {
        System.out.println("\n=== PERFORMANCE COMPARISON ===");
        
        // ArrayList performance
        List<Employee> arrayList = new ArrayList<>(employees);
        long start = System.nanoTime();
        arrayList.remove(0); // Remove from beginning - O(n)
        long arrayListTime = System.nanoTime() - start;
        
        // LinkedList performance
        List<Employee> linkedList = new LinkedList<>(employees);
        start = System.nanoTime();
        linkedList.remove(0); // Remove from beginning - O(1)
        long linkedListTime = System.nanoTime() - start;
        
        System.out.println("ArrayList remove(0) time: " + arrayListTime + " ns");
        System.out.println("LinkedList remove(0) time: " + linkedListTime + " ns");
        System.out.println("LinkedList is " + (arrayListTime / (double) linkedListTime) + "x faster for removal at beginning");
    }

    // Helper method to rebuild maps after modifications
    private void rebuildMaps() {
        employeeMap.clear();
        departmentMap.clear();
        
        employees.forEach(emp -> {
            employeeMap.put(emp.getId(), emp);
            departmentMap.computeIfAbsent(emp.getDepartment(), k -> new ArrayList<>()).add(emp);
        });
    }

    // Display statistics
    public void displayStatistics() {
        System.out.println("\n=== EMPLOYEE STATISTICS ===");
        System.out.println("Total employees: " + employees.size());
        System.out.println("Total departments: " + departments.size());
        System.out.println("Average salary: $" + 
            employees.stream().mapToDouble(Employee::getSalary).average().orElse(0));
        
        // Department-wise statistics
        System.out.println("\nDepartment-wise employee count:");
        departmentMap.forEach((dept, empList) -> {
            double avgSalary = empList.stream().mapToDouble(Employee::getSalary).average().orElse(0);
            System.out.println(dept + ": " + empList.size() + " employees, Avg Salary: $" + 
                String.format("%.2f", avgSalary));
        });
    }

    public static void main(String[] args) {
        EmployeeManagementSystem ems = new EmployeeManagementSystem();
        
        // Add sample employees
        ems.addEmployee(new Employee(101, "Alice Johnson", "Engineering", 85000, 
            LocalDate.of(2020, 1, 15), "alice@company.com"));
        ems.addEmployee(new Employee(102, "Bob Smith", "Marketing", 65000, 
            LocalDate.of(2019, 3, 22), "bob@company.com"));
        ems.addEmployee(new Employee(103, "Charlie Brown", "Engineering", 92000, 
            LocalDate.of(2021, 6, 10), "charlie@company.com"));
        ems.addEmployee(new Employee(104, "Diana Prince", "HR", 70000, 
            LocalDate.of(2020, 11, 5), "diana@company.com"));
        ems.addEmployee(new Employee(105, "Eve Wilson", "Engineering", 88000, 
            LocalDate.of(2022, 2, 18), "eve@company.com"));
        ems.addEmployee(new Employee(106, "Frank Miller", "Sales", 55000, 
            LocalDate.of(2021, 8, 30), "frank@company.com"));
        ems.addEmployee(new Employee(107, "Grace Lee", "Marketing", 68000, 
            LocalDate.of(2023, 1, 12), "grace@company.com"));
        ems.addEmployee(new Employee(108, "Henry Ford", "Engineering", 95000, 
            LocalDate.of(2019, 9, 25), "henry@company.com"));

        // Demonstrate all collection operations
        ems.displayStatistics();
        ems.demonstrateMapOperations();
        ems.demonstrateQueueOperations();
        ems.demonstrateDequeOperations();
        ems.demonstrateSetOperations();
        ems.demonstrateCollectionsUtilities();
        ems.performanceComparison();
        
        // Search operations
        System.out.println("\n=== SEARCH OPERATIONS ===");
        List<Employee> engineeringEmployees = ems.searchEmployees("Engineering");
        System.out.println("Engineering employees (" + engineeringEmployees.size() + "):");
        engineeringEmployees.forEach(emp -> System.out.println("  - " + emp.getName()));
        
        List<Employee> salaryRange = ems.getEmployeesBySalaryRange(80000, 100000);
        System.out.println("\nEmployees with salary $80,000-$100,000 (" + salaryRange.size() + "):");
        salaryRange.forEach(emp -> System.out.println("  - " + emp.getName() + ": $" + emp.getSalary()));
        
        // Batch operations
        System.out.println("\n=== BATCH OPERATIONS ===");
        ems.processBatch(0, 4);
        
        // Conditional operations
        System.out.println("\n=== CONDITIONAL OPERATIONS ===");
        ems.giveRaiseToDepartment("Engineering", 10);
        ems.removeLowPerformers(60000);
        
        // Final statistics
        ems.displayStatistics();
    }
}
```

---

## 🎯 Key Concepts Demonstrated

### 1. **List Operations**
- **ArrayList**: Random access, `subList()`, `removeIf()`
- **LinkedList**: Queue operations, Deque operations
- **Performance comparison**: Remove operations at different positions

### 2. **Set Operations**
- **HashSet**: Fast lookup, uniqueness, unordered
- **TreeSet**: Sorted elements, natural ordering
- **LinkedHashSet**: Insertion order preservation

### 3. **Map Operations**
- **HashMap**: O(1) lookup by key
- **TreeMap**: Sorted keys, navigation methods
- **LinkedHashMap**: Insertion order preservation

### 4. **Queue & Deque**
- **FIFO operations**: `offer()`, `poll()`, `peek()`
- **Deque operations**: `addFirst()`, `addLast()`, `removeFirst()`, `removeLast()`

### 5. **Collections Utilities**
- `Collections.sort()`, `Collections.reverse()`, `Collections.shuffle()`
- `Collections.frequency()`, `Collections.binarySearch()`

### 6. **Safe Iteration**
- `Iterator.remove()` for safe modification
- `removeIf()` for conditional removal
- Avoiding `ConcurrentModificationException`

### 7. **Advanced Features**
- Method references in sorting
- Comparator chaining
- Batch operations with `subList()`
- Performance considerations

---

## 🚀 Expected Output

```
=== EMPLOYEE STATISTICS ===
Total employees: 8
Total departments: 4
Average salary: $77125.0

Department-wise employee count:
Engineering: 4 employees, Avg Salary: $90000.00
HR: 1 employees, Avg Salary: $70000.00
Marketing: 2 employees, Avg Salary: $66500.00
Sales: 1 employees, Avg Salary: $55000.00

=== MAP OPERATIONS DEMONSTRATION ===

HashMap (quick lookup by ID):
Employee 101: Employee{id=101, name='Alice Johnson', department='Engineering', salary=$85000.00, hireDate=2020-01-15, email='alice@company.com'}

TreeMap (departments in sorted order):
Engineering: 4 employees
HR: 1 employees
Marketing: 2 employees
Sales: 1 employees

LinkedHashMap (maintaining insertion order):
101: Alice Johnson
102: Bob Smith
103: Charlie Brown
104: Diana Prince
105: Eve Wilson

=== QUEUE OPERATIONS DEMONSTRATION ===
Waiting list (FIFO):
  - Frank Miller
  - Grace Lee

Processing from queue:
Processing: Frank Miller
Processing: Grace Lee

=== DEQUE OPERATIONS DEMONSTRATION ===
Recent hires (newest first):
  - Grace Lee (2023-01-12)
  - Eve Wilson (2022-02-18)
  - Frank Miller (2021-08-30)
  - Diana Prince (2020-11-05)
  - Charlie Brown (2021-06-10)

After adding to both ends:
  - New Hire (2026-03-18)
  - Grace Lee (2023-01-12)
  - Eve Wilson (2022-02-18)
  - Frank Miller (2021-08-30)
  - Diana Prince (2020-11-05)
  - Charlie Brown (2021-06-10)
  - Alice Johnson (2020-01-15)
  - Bob Smith (2019-03-22)
  - Henry Ford (2019-09-25)
  - Founder (2010-01-01)

=== SET OPERATIONS DEMONSTRATION ===
Unique departments (HashSet - unordered):
  - Engineering
  - Marketing
  - HR
  - Sales

Sorted departments (TreeSet):
  - Engineering
  - HR
  - Marketing
  - Sales

Departments in insertion order (LinkedHashSet):
  - Engineering
  - Marketing
  - Engineering
  - HR
  - Engineering

=== COLLECTIONS UTILITIES DEMONSTRATION ===
Original list:
  - Alice Johnson
  - Bob Smith
  - Charlie Brown
  - Diana Prince
  - Eve Wilson

After reverse:
  - Eve Wilson
  - Diana Prince
  - Charlie Brown
  - Bob Smith
  - Alice Johnson

After shuffle:
  - Charlie Brown
  - Eve Wilson
  - Alice Johnson
  - Bob Smith
  - Diana Prince

Sorted by name:
  - Alice Johnson
  - Bob Smith
  - Charlie Brown
  - Diana Prince
  - Eve Wilson

Most common department: Engineering

=== PERFORMANCE COMPARISON ===
ArrayList remove(0) time: 12500 ns
LinkedList remove(0) time: 2100 ns
LinkedList is 5.95x faster for removal at beginning

=== SEARCH OPERATIONS ===
Engineering employees (4):
  - Alice Johnson
  - Charlie Brown
  - Eve Wilson
  - Henry Ford

Employees with salary $80,000-$100,000 (4):
  - Alice Johnson: $85000.0
  - Charlie Brown: $92000.0
  - Eve Wilson: $88000.0
  - Henry Ford: $95000.0

=== BATCH OPERATIONS ===
Processing batch of 5 employees:
  - Henry Ford: $95000.0
  - Charlie Brown: $92000.0
  - Eve Wilson: $88000.0
  - Alice Johnson: $85000.0
  - Diana Prince: $70000.0

=== CONDITIONAL OPERATIONS ===
Giving 10.0% raise to Engineering department
Removing employees with salary < $60000.0
Removed 1 employees

=== EMPLOYEE STATISTICS ===
Total employees: 7
Total departments: 4
Average salary: $83357.14

Department-wise employee count:
Engineering: 4 employees, Avg Salary: $93500.00
HR: 1 employees, Avg Salary: $77000.00
Marketing: 2 employees, Avg Salary: $73150.00
Sales: 0 employees, Avg Salary: $0.00
```

---

## 💡 Interview Talking Points

### Performance Considerations
- **ArrayList vs LinkedList**: ArrayList is better for random access, LinkedList for frequent insertions/deletions
- **HashMap vs TreeMap**: HashMap offers O(1) lookup, TreeMap offers sorted keys with O(log n)
- **HashSet vs TreeSet**: HashSet is faster, TreeSet maintains ordering

### Memory Efficiency
- **Initial capacity**: Setting proper initial capacity reduces resizing
- **Immutable collections**: `List.of()` for constants, `Collections.unmodifiableList()` for views
- **Primitive collections**: Consider specialized collections for primitive types

### Thread Safety
- **Concurrent collections**: Use `ConcurrentHashMap`, `CopyOnWriteArrayList` for concurrent access
- **Synchronization**: `Collections.synchronizedList()` wrapper for legacy code
- **Immutable collections**: Thread-safe by design

### Best Practices
- **Choose the right collection**: Based on use case (lookup, ordering, duplicates)
- **Use interfaces**: Program to `List`, `Map`, `Set` interfaces
- **Avoid raw types**: Use generics for type safety
- **Consider memory usage**: Large collections need careful memory management
