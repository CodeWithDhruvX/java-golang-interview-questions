// Online Java Compiler
// Use this editor to write, compile and run your Java code online

// salary descending order and software enginner designation
// id,name,age,designation,salary
import java.util.*;
import java.util.stream.Collectors;

class Employee{
    public int id,age,salary;
    public String name,designation;
    
    public Employee(int id,int age,int salary,String name,String designation){
        this.id=id;
        this.age=age;
        this.salary=salary;
        this.name=name;
        this.designation=designation;
    }
    
    @Override
    public String toString(){
        return "ID: " + id + ", Name: " + name + ", Age: " + age + ", Designation: " + designation + ", Salary: " + salary;
    }
}


class Main {
    public static void main(String[] args) {
        ArrayList<Employee> listOfEmployee=new ArrayList<Employee>(List.of(
            new Employee(1,25,24000,"aman","Software Engineer"),
            new Employee(2,30,32000,"deep","HR"),
            new Employee(3,28,35000,"rahul","Software Engineer"),
            new Employee(4,32,45000,"priya","Software Engineer"),
            new Employee(5,27,28000,"neha","Developer")
            ));
            
        // Step 1: Collect all employee records (already done)
        // Step 2: Filter records with "Software Engineer" designation
        // Step 3: Sort by salary in descending order
        // Step 4: Collect and display the results
        
        List<Employee> softwareEngineers = listOfEmployee.stream()
            .filter(emp -> "Software Engineer".equals(emp.designation))
            .sorted((e1, e2) -> Integer.compare(e2.salary, e1.salary)) // Sort by salary descending
            .collect(Collectors.toList());
        
        System.out.println("Software Engineers sorted by salary (descending order):");
        System.out.println("=====================================================");
        
        for (Employee employee : softwareEngineers){
            System.out.println(employee);
        }
        
        // Alternative approach using grouping and counting
        System.out.println("\nAdditional Analysis:");
        System.out.println("====================");
        
        Map<String, List<Employee>> employeesByDesignation = listOfEmployee.stream()
            .collect(Collectors.groupingBy(emp -> emp.designation != null ? emp.designation : "Unknown"));
        
        System.out.println("Employees grouped by designation:");
        employeesByDesignation.forEach((designation, employees) -> {
            System.out.println("\n" + designation + " (" + employees.size() + " employees):");
            employees.stream()
                .sorted((e1, e2) -> Integer.compare(e2.salary, e1.salary))
                .forEach(emp -> System.out.println("  " + emp));
        });
    }
}


sudo code:

step 1: i'll collect all the employee records 
step 2: filter the records to desination that is matching with "Software Enginner" position
Step 3: sorting the salary with descending order.
Step 4: then i'll collect(gropuing) the records and displayed it.