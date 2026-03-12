// Online Java Compiler
// Use this editor to write, compile and run your Java code online

// salary descending order and software enginner designation
// id,name,age,designation,salary
import java.util.*;

class Employee{
    public int id,age,salary;
    public String name,designation;
    
    public Employee(int id,int age,int salary,String name,String designation){
        this.id=id;
        this.age=id;
        this.salary=salary;
        this.name=name;
        this.designation=designation;
    }
    
}


class Main {
    public static void main(String[] args) {
        ArrayList<Employee> listOfEmployee=new ArrayList<Employee>(List.of(
            new Employee(1,12,24000,"aman","Software Enginner"),
            new Employee(1,12,32000,"deep","HR"),
            new Employee(1,12,35000,"aman","")
            ));
            
        ArrayList<Employee> countedEmployees=listOfEmployee.stream()
        .filter(ele->ele.designation.equalsTo("Software Enginner"));
        
        for (Employee employee:countedEmployees){
            System.out.println("Employee with software enginner"+employee.name);
        }
    }
}




sudo code:

step 1: i'll collect all the employee records 
step 2: filter the records to desination that is matching with "Software Enginner" position
Step 3: sorting the salary with descending order.
Step 4: then i'll collect(gropuing) the records and displayed it.