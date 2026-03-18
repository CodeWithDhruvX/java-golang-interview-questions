# Web/API Development: Practical Programs

**Goal**: Master modern web development with Spring Boot REST APIs, JSON processing, and HTTP client operations.

## Prerequisites

Add these dependencies to your `pom.xml` (Maven):

```xml
<dependencies>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-web</artifactId>
        <version>3.2.0</version>
    </dependency>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-data-jpa</artifactId>
        <version>3.2.0</version>
    </dependency>
    <dependency>
        <groupId>com.h2database</groupId>
        <artifactId>h2</artifactId>
        <scope>runtime</scope>
    </dependency>
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-validation</artifactId>
        <version>3.2.0</version>
    </dependency>
    <dependency>
        <groupId>com.fasterxml.jackson.core</groupId>
        <artifactId>jackson-databind</artifactId>
    </dependency>
    <dependency>
        <groupId>org.apache.httpcomponents</groupId>
        <artifactId>httpclient</artifactId>
        <version>4.5.14</version>
    </dependency>
</dependencies>
```

## 1. Spring Boot REST API

### User Management System

```java
// Model/Entity class
@Entity
@Table(name = "users")
public class User {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    
    @Column(nullable = false, unique = true)
    @NotBlank(message = "Username is required")
    private String username;
    
    @Column(nullable = false)
    @Email(message = "Email should be valid")
    private String email;
    
    @Column(nullable = false)
    @NotBlank(message = "Name is required")
    private String name;
    
    @Column
    private Integer age;
    
    @Column
    private String department;
    
    @Enumerated(EnumType.STRING)
    private UserStatus status;
    
    @CreationTimestamp
    private LocalDateTime createdAt;
    
    @UpdateTimestamp
    private LocalDateTime updatedAt;
    
    // Constructors
    public User() {}
    
    public User(String username, String email, String name, Integer age, String department) {
        this.username = username;
        this.email = email;
        this.name = name;
        this.age = age;
        this.department = department;
        this.status = UserStatus.ACTIVE;
    }
    
    // Getters and Setters
    public Long getId() { return id; }
    public void setId(Long id) { this.id = id; }
    
    public String getUsername() { return username; }
    public void setUsername(String username) { this.username = username; }
    
    public String getEmail() { return email; }
    public void setEmail(String email) { this.email = email; }
    
    public String getName() { return name; }
    public void setName(String name) { this.name = name; }
    
    public Integer getAge() { return age; }
    public void setAge(Integer age) { this.age = age; }
    
    public String getDepartment() { return department; }
    public void setDepartment(String department) { this.department = department; }
    
    public UserStatus getStatus() { return status; }
    public void setStatus(UserStatus status) { this.status = status; }
    
    public LocalDateTime getCreatedAt() { return createdAt; }
    public void setCreatedAt(LocalDateTime createdAt) { this.createdAt = createdAt; }
    
    public LocalDateTime getUpdatedAt() { return updatedAt; }
    public void setUpdatedAt(LocalDateTime updatedAt) { this.updatedAt = updatedAt; }
}

enum UserStatus {
    ACTIVE, INACTIVE, SUSPENDED
}

// DTO (Data Transfer Object)
public class UserDTO {
    private Long id;
    private String username;
    private String email;
    private String name;
    private Integer age;
    private String department;
    private String status;
    private String createdAt;
    
    // Constructors
    public UserDTO() {}
    
    public UserDTO(User user) {
        this.id = user.getId();
        this.username = user.getUsername();
        this.email = user.getEmail();
        this.name = user.getName();
        this.age = user.getAge();
        this.department = user.getDepartment();
        this.status = user.getStatus().toString();
        this.createdAt = user.getCreatedAt() != null ? 
            user.getCreatedAt().toString() : null;
    }
    
    // Getters and Setters
    public Long getId() { return id; }
    public void setId(Long id) { this.id = id; }
    
    public String getUsername() { return username; }
    public void setUsername(String username) { this.username = username; }
    
    public String getEmail() { return email; }
    public void setEmail(String email) { this.email = email; }
    
    public String getName() { return name; }
    public void setName(String name) { this.name = name; }
    
    public Integer getAge() { return age; }
    public void setAge(Integer age) { this.age = age; }
    
    public String getDepartment() { return department; }
    public void setDepartment(String department) { this.department = department; }
    
    public String getStatus() { return status; }
    public void setStatus(String status) { this.status = status; }
    
    public String getCreatedAt() { return createdAt; }
    public void setCreatedAt(String createdAt) { this.createdAt = createdAt; }
}

// Repository interface
@Repository
public interface UserRepository extends JpaRepository<User, Long> {
    Optional<User> findByUsername(String username);
    Optional<User> findByEmail(String email);
    List<User> findByDepartment(String department);
    List<User> findByStatus(UserStatus status);
    List<User> findByNameContainingIgnoreCase(String name);
    
    @Query("SELECT u FROM User u WHERE u.age BETWEEN :minAge AND :maxAge")
    List<User> findByAgeBetween(@Param("minAge") Integer minAge, @Param("maxAge") Integer maxAge);
}

// Service layer
@Service
@Transactional
public class UserService {
    
    @Autowired
    private UserRepository userRepository;
    
    public List<UserDTO> getAllUsers() {
        return userRepository.findAll().stream()
                .map(UserDTO::new)
                .collect(Collectors.toList());
    }
    
    public Optional<UserDTO> getUserById(Long id) {
        return userRepository.findById(id)
                .map(UserDTO::new);
    }
    
    public Optional<UserDTO> getUserByUsername(String username) {
        return userRepository.findByUsername(username)
                .map(UserDTO::new);
    }
    
    public UserDTO createUser(User user) {
        // Validate unique username and email
        if (userRepository.existsByUsername(user.getUsername())) {
            throw new RuntimeException("Username already exists");
        }
        if (userRepository.existsByEmail(user.getEmail())) {
            throw new RuntimeException("Email already exists");
        }
        
        User savedUser = userRepository.save(user);
        return new UserDTO(savedUser);
    }
    
    public Optional<UserDTO> updateUser(Long id, User userDetails) {
        return userRepository.findById(id)
                .map(user -> {
                    user.setName(userDetails.getName());
                    user.setEmail(userDetails.getEmail());
                    user.setAge(userDetails.getAge());
                    user.setDepartment(userDetails.getDepartment());
                    user.setStatus(userDetails.getStatus());
                    return userRepository.save(user);
                })
                .map(UserDTO::new);
    }
    
    public boolean deleteUser(Long id) {
        if (userRepository.existsById(id)) {
            userRepository.deleteById(id);
            return true;
        }
        return false;
    }
    
    public List<UserDTO> getUsersByDepartment(String department) {
        return userRepository.findByDepartment(department).stream()
                .map(UserDTO::new)
                .collect(Collectors.toList());
    }
    
    public List<UserDTO> getActiveUsers() {
        return userRepository.findByStatus(UserStatus.ACTIVE).stream()
                .map(UserDTO::new)
                .collect(Collectors.toList());
    }
    
    public List<UserDTO> searchUsers(String keyword) {
        return userRepository.findByNameContainingIgnoreCase(keyword).stream()
                .map(UserDTO::new)
                .collect(Collectors.toList());
    }
    
    public List<UserDTO> getUsersByAgeRange(Integer minAge, Integer maxAge) {
        return userRepository.findByAgeBetween(minAge, maxAge).stream()
                .map(UserDTO::new)
                .collect(Collectors.toList());
    }
}

// REST Controller
@RestController
@RequestMapping("/api/users")
@Validated
public class UserController {
    
    @Autowired
    private UserService userService;
    
    // GET all users
    @GetMapping
    public ResponseEntity<List<UserDTO>> getAllUsers() {
        List<UserDTO> users = userService.getAllUsers();
        return ResponseEntity.ok(users);
    }
    
    // GET user by ID
    @GetMapping("/{id}")
    public ResponseEntity<UserDTO> getUserById(@PathVariable Long id) {
        return userService.getUserById(id)
                .map(ResponseEntity::ok)
                .orElse(ResponseEntity.notFound().build());
    }
    
    // GET user by username
    @GetMapping("/username/{username}")
    public ResponseEntity<UserDTO> getUserByUsername(@PathVariable String username) {
        return userService.getUserByUsername(username)
                .map(ResponseEntity::ok)
                .orElse(ResponseEntity.notFound().build());
    }
    
    // POST create user
    @PostMapping
    public ResponseEntity<?> createUser(@Valid @RequestBody User user) {
        try {
            UserDTO createdUser = userService.createUser(user);
            return ResponseEntity.created(URI.create("/api/users/" + createdUser.getId()))
                    .body(createdUser);
        } catch (RuntimeException e) {
            return ResponseEntity.badRequest().body(e.getMessage());
        }
    }
    
    // PUT update user
    @PutMapping("/{id}")
    public ResponseEntity<?> updateUser(@PathVariable Long id, 
                                        @Valid @RequestBody User userDetails) {
        return userService.updateUser(id, userDetails)
                .map(ResponseEntity::ok)
                .orElse(ResponseEntity.notFound().build());
    }
    
    // DELETE user
    @DeleteMapping("/{id}")
    public ResponseEntity<Void> deleteUser(@PathVariable Long id) {
        if (userService.deleteUser(id)) {
            return ResponseEntity.noContent().build();
        }
        return ResponseEntity.notFound().build();
    }
    
    // GET users by department
    @GetMapping("/department/{department}")
    public ResponseEntity<List<UserDTO>> getUsersByDepartment(@PathVariable String department) {
        List<UserDTO> users = userService.getUsersByDepartment(department);
        return ResponseEntity.ok(users);
    }
    
    // GET active users
    @GetMapping("/active")
    public ResponseEntity<List<UserDTO>> getActiveUsers() {
        List<UserDTO> users = userService.getActiveUsers();
        return ResponseEntity.ok(users);
    }
    
    // GET search users
    @GetMapping("/search")
    public ResponseEntity<List<UserDTO>> searchUsers(@RequestParam String keyword) {
        List<UserDTO> users = userService.searchUsers(keyword);
        return ResponseEntity.ok(users);
    }
    
    // GET users by age range
    @GetMapping("/age-range")
    public ResponseEntity<List<UserDTO>> getUsersByAgeRange(@RequestParam Integer minAge,
                                                            @RequestParam Integer maxAge) {
        List<UserDTO> users = userService.getUsersByAgeRange(minAge, maxAge);
        return ResponseEntity.ok(users);
    }
    
    // Custom response with pagination
    @GetMapping("/paginated")
    public ResponseEntity<Map<String, Object>> getPaginatedUsers(
            @RequestParam(defaultValue = "0") int page,
            @RequestParam(defaultValue = "10") int size) {
        
        // This would require implementing pagination in the repository
        Map<String, Object> response = new HashMap<>();
        response.put("users", userService.getAllUsers());
        response.put("currentPage", page);
        response.put("pageSize", size);
        response.put("totalElements", userService.getAllUsers().size());
        
        return ResponseEntity.ok(response);
    }
}

// Global Exception Handler
@ControllerAdvice
public class GlobalExceptionHandler {
    
    @ExceptionHandler(MethodArgumentNotValidException.class)
    public ResponseEntity<Map<String, String>> handleValidationExceptions(
            MethodArgumentNotValidException ex) {
        
        Map<String, String> errors = new HashMap<>();
        ex.getBindingResult().getFieldErrors().forEach(error -> 
            errors.put(error.getField(), error.getDefaultMessage()));
        
        return ResponseEntity.badRequest().body(errors);
    }
    
    @ExceptionHandler(RuntimeException.class)
    public ResponseEntity<String> handleRuntimeException(RuntimeException ex) {
        return ResponseEntity.badRequest().body(ex.getMessage());
    }
    
    @ExceptionHandler(Exception.class)
    public ResponseEntity<String> handleGeneralException(Exception ex) {
        return ResponseEntity.status(HttpStatus.INTERNAL_SERVER_ERROR)
                .body("An unexpected error occurred");
    }
}

// Main Application class
@SpringBootApplication
public class UserManagementApplication {
    public static void main(String[] args) {
        SpringApplication.run(UserManagementApplication.class, args);
    }
    
    @Bean
    public CommandLineRunner loadData(UserService userService) {
        return args -> {
            // Load sample data
            User user1 = new User("john_doe", "john@example.com", "John Doe", 30, "IT");
            User user2 = new User("jane_smith", "jane@example.com", "Jane Smith", 28, "HR");
            User user3 = new User("bob_wilson", "bob@example.com", "Bob Wilson", 35, "Finance");
            User user4 = new User("alice_brown", "alice@example.com", "Alice Brown", 32, "IT");
            
            userService.createUser(user1);
            userService.createUser(user2);
            userService.createUser(user3);
            userService.createUser(user4);
            
            System.out.println("Sample data loaded successfully!");
        };
    }
}
```

## 2. JSON Processing with Jackson

### Advanced JSON Operations

```java
import com.fasterxml.jackson.annotation.*;
import com.fasterxml.jackson.core.*;
import com.fasterxml.jackson.databind.*;
import com.fasterxml.jackson.datatype.jsr310.JavaTimeModule;
import java.io.*;
import java.time.LocalDateTime;
import java.util.*;

// Complex POJO with Jackson annotations
@JsonInclude(JsonInclude.Include.NON_NULL)
@JsonPropertyOrder({"id", "name", "email", "age", "address", "skills", "projects", "metadata"})
public class Employee {
    
    @JsonProperty("employee_id")
    private Long id;
    
    @JsonProperty("full_name")
    private String name;
    
    @JsonProperty("email_address")
    private String email;
    
    @JsonProperty("age")
    private Integer age;
    
    @JsonProperty("address")
    private Address address;
    
    @JsonProperty("skills")
    private List<String> skills;
    
    @JsonProperty("projects")
    private List<Project> projects;
    
    @JsonIgnore
    private String password; // Won't be included in JSON
    
    @JsonProperty("created_at")
    @JsonFormat(pattern = "yyyy-MM-dd HH:mm:ss")
    private LocalDateTime createdAt;
    
    @JsonProperty("metadata")
    private Map<String, Object> metadata;
    
    // Constructors
    public Employee() {}
    
    public Employee(Long id, String name, String email, Integer age) {
        this.id = id;
        this.name = name;
        this.email = email;
        this.age = age;
        this.skills = new ArrayList<>();
        this.projects = new ArrayList<>();
        this.metadata = new HashMap<>();
        this.createdAt = LocalDateTime.now();
    }
    
    // Getters and Setters with annotations
    @JsonProperty("employee_id")
    public Long getId() { return id; }
    
    @JsonProperty("employee_id")
    public void setId(Long id) { this.id = id; }
    
    @JsonProperty("full_name")
    public String getName() { return name; }
    
    @JsonProperty("full_name")
    public void setName(String name) { this.name = name; }
    
    @JsonProperty("email_address")
    public String getEmail() { return email; }
    
    @JsonProperty("email_address")
    public void setEmail(String email) { this.email = email; }
    
    public Integer getAge() { return age; }
    public void setAge(Integer age) { this.age = age; }
    
    public Address getAddress() { return address; }
    public void setAddress(Address address) { this.address = address; }
    
    public List<String> getSkills() { return skills; }
    public void setSkills(List<String> skills) { this.skills = skills; }
    
    public List<Project> getProjects() { return projects; }
    public void setProjects(List<Project> projects) { this.projects = projects; }
    
    @JsonIgnore
    public String getPassword() { return password; }
    
    @JsonIgnore
    public void setPassword(String password) { this.password = password; }
    
    @JsonProperty("created_at")
    public LocalDateTime getCreatedAt() { return createdAt; }
    
    @JsonProperty("created_at")
    public void setCreatedAt(LocalDateTime createdAt) { this.createdAt = createdAt; }
    
    @JsonProperty("metadata")
    public Map<String, Object> getMetadata() { return metadata; }
    
    @JsonProperty("metadata")
    public void setMetadata(Map<String, Object> metadata) { this.metadata = metadata; }
    
    // Custom serialization method
    @JsonGetter("display_name")
    public String getDisplayName() {
        return name + " (" + email + ")";
    }
    
    // Add skill method
    public void addSkill(String skill) {
        if (skills == null) {
            skills = new ArrayList<>();
        }
        skills.add(skill);
    }
    
    // Add project method
    public void addProject(Project project) {
        if (projects == null) {
            projects = new ArrayList<>();
        }
        projects.add(project);
    }
}

class Address {
    @JsonProperty("street")
    private String street;
    
    @JsonProperty("city")
    private String city;
    
    @JsonProperty("state")
    private String state;
    
    @JsonProperty("zip_code")
    private String zipCode;
    
    @JsonProperty("country")
    private String country;
    
    // Constructors
    public Address() {}
    
    public Address(String street, String city, String state, String zipCode, String country) {
        this.street = street;
        this.city = city;
        this.state = state;
        this.zipCode = zipCode;
        this.country = country;
    }
    
    // Getters and Setters
    public String getStreet() { return street; }
    public void setStreet(String street) { this.street = street; }
    
    public String getCity() { return city; }
    public void setCity(String city) { this.city = city; }
    
    public String getState() { return state; }
    public void setState(String state) { this.state = state; }
    
    @JsonProperty("zip_code")
    public String getZipCode() { return zipCode; }
    
    @JsonProperty("zip_code")
    public void setZipCode(String zipCode) { this.zipCode = zipCode; }
    
    public String getCountry() { return country; }
    public void setCountry(String country) { this.country = country; }
}

class Project {
    @JsonProperty("project_id")
    private String projectId;
    
    @JsonProperty("project_name")
    private String projectName;
    
    @JsonProperty("start_date")
    @JsonFormat(pattern = "yyyy-MM-dd")
    private String startDate;
    
    @JsonProperty("end_date")
    @JsonFormat(pattern = "yyyy-MM-dd")
    private String endDate;
    
    @JsonProperty("status")
    private String status;
    
    @JsonProperty("technologies")
    private List<String> technologies;
    
    // Constructors
    public Project() {}
    
    public Project(String projectId, String projectName, String startDate, String endDate, String status) {
        this.projectId = projectId;
        this.projectName = projectName;
        this.startDate = startDate;
        this.endDate = endDate;
        this.status = status;
        this.technologies = new ArrayList<>();
    }
    
    // Getters and Setters
    @JsonProperty("project_id")
    public String getProjectId() { return projectId; }
    
    @JsonProperty("project_id")
    public void setProjectId(String projectId) { this.projectId = projectId; }
    
    @JsonProperty("project_name")
    public String getProjectName() { return projectName; }
    
    @JsonProperty("project_name")
    public void setProjectName(String projectName) { this.projectName = projectName; }
    
    @JsonProperty("start_date")
    public String getStartDate() { return startDate; }
    
    @JsonProperty("start_date")
    public void setStartDate(String startDate) { this.startDate = startDate; }
    
    @JsonProperty("end_date")
    public String getEndDate() { return endDate; }
    
    @JsonProperty("end_date")
    public void setEndDate(String endDate) { this.endDate = endDate; }
    
    public String getStatus() { return status; }
    public void setStatus(String status) { this.status = status; }
    
    public List<String> getTechnologies() { return technologies; }
    public void setTechnologies(List<String> technologies) { this.technologies = technologies; }
    
    public void addTechnology(String technology) {
        if (technologies == null) {
            technologies = new ArrayList<>();
        }
        technologies.add(technology);
    }
}

public class JsonProcessingDemo {
    
    private static ObjectMapper objectMapper;
    
    static {
        objectMapper = new ObjectMapper();
        objectMapper.registerModule(new JavaTimeModule());
        objectMapper.configure(SerializationFeature.INDENT_OUTPUT, true);
        objectMapper.configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false);
    }
    
    public static void main(String[] args) {
        try {
            // 1. Basic serialization
            System.out.println("=== Basic JSON Serialization ===");
            Employee employee = createSampleEmployee();
            String json = objectMapper.writeValueAsString(employee);
            System.out.println(json);
            
            // 2. Basic deserialization
            System.out.println("\n=== Basic JSON Deserialization ===");
            Employee deserializedEmployee = objectMapper.readValue(json, Employee.class);
            System.out.println("Deserialized: " + deserializedEmployee.getName());
            
            // 3. Working with JSON arrays
            System.out.println("\n=== JSON Array Processing ===");
            List<Employee> employees = Arrays.asList(
                createSampleEmployee(),
                createSampleEmployee2(),
                createSampleEmployee3()
            );
            
            String employeesJson = objectMapper.writeValueAsString(employees);
            System.out.println("Employees JSON:");
            System.out.println(employeesJson);
            
            // 4. Parse JSON array
            List<Employee> parsedEmployees = objectMapper.readValue(employeesJson, 
                new TypeReference<List<Employee>>() {});
            System.out.println("Parsed " + parsedEmployees.size() + " employees");
            
            // 5. Working with JSON as Tree
            System.out.println("\n=== JSON Tree Processing ===");
            JsonNode jsonNode = objectMapper.readTree(json);
            System.out.println("Employee ID: " + jsonNode.get("employee_id").asText());
            System.out.println("Employee Name: " + jsonNode.get("full_name").asText());
            
            // 6. Extract nested data
            if (jsonNode.has("address")) {
                JsonNode addressNode = jsonNode.get("address");
                System.out.println("City: " + addressNode.get("city").asText());
            }
            
            // 7. Modify JSON as Tree
            System.out.println("\n=== JSON Tree Modification ===");
            ObjectNode modifiedNode = (ObjectNode) jsonNode;
            modifiedNode.put("bonus_eligible", true);
            modifiedNode.put("department", "Engineering");
            
            String modifiedJson = objectMapper.writeValueAsString(modifiedNode);
            System.out.println("Modified JSON:");
            System.out.println(modifiedJson);
            
            // 8. JSON from/to File
            System.out.println("\n=== File Operations ===");
            writeJsonToFile(employee, "employee.json");
            Employee fileEmployee = readJsonFromFile("employee.json");
            System.out.println("Read from file: " + fileEmployee.getName());
            
            // 9. JSON to Map
            System.out.println("\n=== JSON to Map ===");
            Map<String, Object> employeeMap = objectMapper.readValue(json, 
                new TypeReference<Map<String, Object>>() {});
            System.out.println("Converted to Map:");
            employeeMap.forEach((key, value) -> System.out.println(key + ": " + value));
            
            // 10. Custom serialization
            System.out.println("\n=== Custom Serialization ===");
            String customJson = serializeWithCustomFormat(employee);
            System.out.println("Custom format:");
            System.out.println(customJson);
            
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    
    private static Employee createSampleEmployee() {
        Employee employee = new Employee(1L, "John Doe", "john.doe@company.com", 30);
        
        Address address = new Address("123 Main St", "New York", "NY", "10001", "USA");
        employee.setAddress(address);
        
        employee.addSkill("Java");
        employee.addSkill("Spring Boot");
        employee.addSkill("React");
        employee.addSkill("Docker");
        
        Project project1 = new Project("P001", "E-commerce Platform", "2023-01-01", "2023-12-31", "Active");
        project1.addTechnology("Spring Boot");
        project1.addTechnology("React");
        project1.addTechnology("PostgreSQL");
        employee.addProject(project1);
        
        Project project2 = new Project("P002", "Mobile App", "2023-06-01", "2024-03-31", "In Progress");
        project2.addTechnology("React Native");
        project2.addTechnology("Node.js");
        employee.addProject(project2);
        
        employee.getMetadata().put("department", "Engineering");
        employee.getMetadata().put("salary", 95000);
        employee.getMetadata().put("remote_worker", true);
        
        return employee;
    }
    
    private static Employee createSampleEmployee2() {
        Employee employee = new Employee(2L, "Jane Smith", "jane.smith@company.com", 28);
        
        Address address = new Address("456 Oak Ave", "San Francisco", "CA", "94102", "USA");
        employee.setAddress(address);
        
        employee.addSkill("Python");
        employee.addSkill("Django");
        employee.addSkill("Vue.js");
        employee.addSkill("AWS");
        
        Project project = new Project("P003", "Data Analytics", "2023-03-01", "2023-11-30", "Completed");
        project.addTechnology("Python");
        project.addTechnology("TensorFlow");
        project.addTechnology("AWS");
        employee.addProject(project);
        
        employee.getMetadata().put("department", "Data Science");
        employee.getMetadata().put("salary", 105000);
        employee.getMetadata().put("remote_worker", true);
        
        return employee;
    }
    
    private static Employee createSampleEmployee3() {
        Employee employee = new Employee(3L, "Bob Wilson", "bob.wilson@company.com", 35);
        
        Address address = new Address("789 Pine Rd", "Austin", "TX", "78701", "USA");
        employee.setAddress(address);
        
        employee.addSkill("C#");
        employee.addSkill(".NET");
        employee.addSkill("Angular");
        employee.addSkill("Azure");
        
        Project project = new Project("P004", "Enterprise System", "2023-02-01", "2024-01-31", "In Progress");
        project.addTechnology(".NET Core");
        project.addTechnology("Angular");
        project.addTechnology("Azure");
        employee.addProject(project);
        
        employee.getMetadata().put("department", "Enterprise Solutions");
        employee.getMetadata().put("salary", 110000);
        employee.getMetadata().put("remote_worker", false);
        
        return employee;
    }
    
    private static void writeJsonToFile(Employee employee, String filename) throws IOException {
        objectMapper.writeValue(new File(filename), employee);
        System.out.println("JSON written to " + filename);
    }
    
    private static Employee readJsonFromFile(String filename) throws IOException {
        return objectMapper.readValue(new File(filename), Employee.class);
    }
    
    private static String serializeWithCustomFormat(Employee employee) throws JsonProcessingException {
        // Create a custom view
        ObjectNode customNode = objectMapper.createObjectNode();
        customNode.put("id", employee.getId());
        customNode.put("name", employee.getName());
        customNode.put("email", employee.getEmail());
        
        // Add skills as comma-separated string
        if (employee.getSkills() != null) {
            customNode.put("skills", String.join(", ", employee.getSkills()));
        }
        
        // Add project count
        if (employee.getProjects() != null) {
            customNode.put("project_count", employee.getProjects().size());
        }
        
        return objectMapper.writeValueAsString(customNode);
    }
}
```

## 3. HTTP Client Operations

### REST API Client with Apache HttpClient

```java
import org.apache.http.*;
import org.apache.http.client.*;
import org.apache.http.client.methods.*;
import org.apache.http.entity.*;
import org.apache.http.impl.client.*;
import org.apache.http.util.*;
import com.fasterxml.jackson.databind.ObjectMapper;
import java.io.*;
import java.util.*;

public class ApiClientDemo {
    
    private static final String BASE_URL = "https://jsonplaceholder.typicode.com";
    private static final ObjectMapper objectMapper = new ObjectMapper();
    private static final CloseableHttpClient httpClient = HttpClients.createDefault();
    
    public static void main(String[] args) {
        try {
            // 1. GET request
            System.out.println("=== GET Request ===");
            getUsers();
            getUserById(1);
            
            // 2. POST request
            System.out.println("\n=== POST Request ===");
            createNewUser();
            
            // 3. PUT request
            System.out.println("\n=== PUT Request ===");
            updateUser(1);
            
            // 4. DELETE request
            System.out.println("\n=== DELETE Request ===");
            deleteUser(1);
            
            // 5. Custom API client
            System.out.println("\n=== Custom API Client ===");
            CustomApiClient apiClient = new CustomApiClient();
            apiClient.performAllOperations();
            
        } catch (Exception e) {
            e.printStackTrace();
        } finally {
            try {
                httpClient.close();
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }
    
    // GET all users
    private static void getUsers() throws IOException {
        HttpGet request = new HttpGet(BASE_URL + "/users");
        
        try (CloseableHttpResponse response = httpClient.execute(request)) {
            System.out.println("Response Code: " + response.getStatusLine().getStatusCode());
            
            String responseBody = EntityUtils.toString(response.getEntity());
            System.out.println("Response Body: " + responseBody.substring(0, Math.min(200, responseBody.length())) + "...");
            
            // Parse JSON array
            List<Map<String, Object>> users = objectMapper.readValue(responseBody, 
                new com.fasterxml.jackson.core.type.TypeReference<List<Map<String, Object>>>() {});
            System.out.println("Total users: " + users.size());
        }
    }
    
    // GET user by ID
    private static void getUserById(int id) throws IOException {
        HttpGet request = new HttpGet(BASE_URL + "/users/" + id);
        
        try (CloseableHttpResponse response = httpClient.execute(request)) {
            System.out.println("Response Code: " + response.getStatusLine().getStatusCode());
            
            String responseBody = EntityUtils.toString(response.getEntity());
            System.out.println("User: " + responseBody);
        }
    }
    
    // POST create new user
    private static void createNewUser() throws IOException {
        HttpPost request = new HttpPost(BASE_URL + "/users");
        
        // Create user data
        Map<String, Object> userData = new HashMap<>();
        userData.put("name", "John Doe");
        userData.put("username", "johndoe");
        userData.put("email", "john.doe@example.com");
        userData.put("address", Map.of(
            "street", "123 Main St",
            "city", "New York",
            "zipcode", "10001"
        ));
        userData.put("phone", "555-1234");
        userData.put("website", "johndoe.com");
        userData.put("company", Map.of(
            "name", "Tech Corp",
            "catchPhrase", "Innovation at its best"
        ));
        
        String json = objectMapper.writeValueAsString(userData);
        request.setEntity(new StringEntity(json, ContentType.APPLICATION_JSON));
        
        try (CloseableHttpResponse response = httpClient.execute(request)) {
            System.out.println("Response Code: " + response.getStatusLine().getStatusCode());
            
            String responseBody = EntityUtils.toString(response.getEntity());
            System.out.println("Created User: " + responseBody);
        }
    }
    
    // PUT update user
    private static void updateUser(int id) throws IOException {
        HttpPut request = new HttpPut(BASE_URL + "/users/" + id);
        
        Map<String, Object> updateData = new HashMap<>();
        updateData.put("name", "John Updated");
        updateData.put("email", "john.updated@example.com");
        updateData.put("phone", "555-5678");
        
        String json = objectMapper.writeValueAsString(updateData);
        request.setEntity(new StringEntity(json, ContentType.APPLICATION_JSON));
        
        try (CloseableHttpResponse response = httpClient.execute(request)) {
            System.out.println("Response Code: " + response.getStatusLine().getStatusCode());
            
            String responseBody = EntityUtils.toString(response.getEntity());
            System.out.println("Updated User: " + responseBody);
        }
    }
    
    // DELETE user
    private static void deleteUser(int id) throws IOException {
        HttpDelete request = new HttpDelete(BASE_URL + "/users/" + id);
        
        try (CloseableHttpResponse response = httpClient.execute(request)) {
            System.out.println("Response Code: " + response.getStatusLine().getStatusCode());
            
            String responseBody = EntityUtils.toString(response.getEntity());
            System.out.println("Delete Response: " + responseBody);
        }
    }
}

// Custom API Client wrapper
class CustomApiClient {
    private static final String BASE_URL = "https://jsonplaceholder.typicode.com";
    private final CloseableHttpClient httpClient;
    private final ObjectMapper objectMapper;
    
    public CustomApiClient() {
        this.httpClient = HttpClients.custom()
                .setMaxConnTotal(100)
                .setMaxConnPerRoute(20)
                .build();
        this.objectMapper = new ObjectMapper();
    }
    
    public void performAllOperations() {
        try {
            // Get all posts
            List<Post> posts = getAllPosts();
            System.out.println("Retrieved " + posts.size() + " posts");
            
            // Get specific post
            Post post = getPostById(1);
            System.out.println("Post title: " + post.getTitle());
            
            // Get comments for post
            List<Comment> comments = getCommentsForPost(1);
            System.out.println("Post has " + comments.size() + " comments");
            
            // Create new post
            Post newPost = createPost("My New Post", "This is the content of my new post.");
            System.out.println("Created post with ID: " + newPost.getId());
            
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    
    public List<Post> getAllPosts() throws IOException {
        HttpGet request = new HttpGet(BASE_URL + "/posts");
        
        try (CloseableHttpResponse response = httpClient.execute(request)) {
            String responseBody = EntityUtils.toString(response.getEntity());
            return objectMapper.readValue(responseBody, 
                new com.fasterxml.jackson.core.type.TypeReference<List<Post>>() {});
        }
    }
    
    public Post getPostById(int id) throws IOException {
        HttpGet request = new HttpGet(BASE_URL + "/posts/" + id);
        
        try (CloseableHttpResponse response = httpClient.execute(request)) {
            String responseBody = EntityUtils.toString(response.getEntity());
            return objectMapper.readValue(responseBody, Post.class);
        }
    }
    
    public List<Comment> getCommentsForPost(int postId) throws IOException {
        HttpGet request = new HttpGet(BASE_URL + "/posts/" + postId + "/comments");
        
        try (CloseableHttpResponse response = httpClient.execute(request)) {
            String responseBody = EntityUtils.toString(response.getEntity());
            return objectMapper.readValue(responseBody, 
                new com.fasterxml.jackson.core.type.TypeReference<List<Comment>>() {});
        }
    }
    
    public Post createPost(String title, String body) throws IOException {
        HttpPost request = new HttpPost(BASE_URL + "/posts");
        
        Map<String, Object> postData = new HashMap<>();
        postData.put("title", title);
        postData.put("body", body);
        postData.put("userId", 1);
        
        String json = objectMapper.writeValueAsString(postData);
        request.setEntity(new StringEntity(json, ContentType.APPLICATION_JSON));
        
        try (CloseableHttpResponse response = httpClient.execute(request)) {
            String responseBody = EntityUtils.toString(response.getEntity());
            return objectMapper.readValue(responseBody, Post.class);
        }
    }
    
    public void close() throws IOException {
        httpClient.close();
    }
}

// POJO classes for JSON parsing
class Post {
    private int id;
    private int userId;
    private String title;
    private String body;
    
    // Getters and Setters
    public int getId() { return id; }
    public void setId(int id) { this.id = id; }
    
    public int getUserId() { return userId; }
    public void setUserId(int userId) { this.userId = userId; }
    
    public String getTitle() { return title; }
    public void setTitle(String title) { this.title = title; }
    
    public String getBody() { return body; }
    public void setBody(String body) { this.body = body; }
}

class Comment {
    private int id;
    private int postId;
    private String name;
    private String email;
    private String body;
    
    // Getters and Setters
    public int getId() { return id; }
    public void setId(int id) { this.id = id; }
    
    public int getPostId() { return postId; }
    public void setPostId(int postId) { this.postId = postId; }
    
    public String getName() { return name; }
    public void setName(String name) { this.name = name; }
    
    public String getEmail() { return email; }
    public void setEmail(String email) { this.email = email; }
    
    public String getBody() { return body; }
    public void setBody(String body) { this.body = body; }
}
```

## Practice Exercises

1. **REST API**: Build a complete CRUD API for a book management system
2. **JSON Processing**: Create a configuration management system with JSON files
3. **HTTP Client**: Build a weather data aggregator that calls multiple APIs
4. **Advanced API**: Implement file upload/download endpoints with proper error handling

## Interview Questions

1. What's the difference between @RestController and @Controller?
2. How do you handle exceptions in Spring Boot REST APIs?
3. What are the benefits of using DTOs instead of returning entities directly?
4. How do you implement pagination in REST APIs?
5. What's the difference between GET and POST methods beyond idempotency?
