# Practical Coding Questions: Spring Boot Core, Web, and Data JPA

*These problems are designed to test your actual coding ability during a technical interview. Try to write out the code or pseudocode before looking at the suggested architectures.*

---

## 1. REST API Setup: Library Management System
**Problem Statement:**
Design and implement a basic REST API for a "Library Management System".
1. Create a `Book` entity with fields: `id`, `isbn` (must be exactly 13 chars), `title` (cannot be blank), `author`, and `availableCopies` (must be >= 0).
2. Create a `BookDto` to accept incoming requests. Integrate **Spring Validation** so that if a user sends a Book with `availableCopies = -1`, the system rejects the request before hitting the controller logic.
3. Write a `@RestController` to handle basic CRUD operations:
   - `POST /books` (Add a new book)
   - `GET /books/{id}` (Retrieve a book)

**Expected Focus Areas:**
- `@RestController`, `@RequestMapping`.
- `DTO` patterns vs `Entity`.
- **Validation Annotations:** `@NotBlank`, `@Size(min=13, max=13)`, `@Min(0)`.
- Use `@Valid` on the `@RequestBody`.

---

## 2. Global Exception Handling
**Problem Statement:**
Following up on the Library System above, when a requested Book ID is not found in the database, your service throws a custom `ResourceNotFoundException`.
1. Write the `ResourceNotFoundException` class.
2. Implement a global exception handler using `@RestControllerAdvice` that catches this exception.
3. The API should return a standardized JSON error response with a `404 Not Found` HTTP status, looking like this:
```json
{
  "timestamp": "2024-05-12T10:00:00",
  "status": 404,
  "error": "Not Found",
  "message": "Book with ID 5 not found."
}
```

**Expected Focus Areas:**
- `@RestControllerAdvice` and `@ExceptionHandler`.
- Constructing a custom `ErrorResponse` class (using Lombok for brevity).
- Returning `ResponseEntity<ErrorResponse>`.

---

## 3. Spring Data JPA: Relational Mappings & Custom Queries
**Problem Statement:**
Extend the Library System by introducing a `LibraryMember` entity.
1. Establish a **One-to-Many** relationship between `LibraryMember` and `Book` (assuming a book is physically borrowed by one member at a time). Which entity owns the relationship, and where does the `@JoinColumn` go?
2. Create a `BookRepository` interface extending `JpaRepository`.
3. Write a derived query method to find all books by a specific `author` where `availableCopies` is greater than 0.
4. Using JPQL (`@Query`), write a custom query to find all members who currently have a book checked out.

**Expected Focus Areas:**
- `@OneToMany(mappedBy = "member")` on the `LibraryMember` side.
- `@ManyToOne` and `@JoinColumn(name = "member_id")` on the `Book` side (Owning side).
- Derived Method: `List<Book> findByAuthorAndAvailableCopiesGreaterThan(String author, int copies);`
- JPQL: `@Query("SELECT DISTINCT m FROM LibraryMember m JOIN m.borrowedBooks b")`

---

## 4. Transaction Management & Propagation (`REQUIRES_NEW`)
**Problem Statement:**
You are building an E-Commerce application. When a user clicks "Checkout", two things happen:
1. The `InventoryService` deducts the stock.
2. The `AuditLogService` saves a log entry to the database: "User attempted checkout for item X."

**The Catch:** If the `InventoryService` fails (e.g., throwing an `Out_Of_Stock_Exception`), the main transaction rolls back so the order isn't placed. **However**, the audit log MUST still be saved to the database regardless of the main transaction's success or failure, so the business team can track failed checkout attempts.

Write the pseudocode for the `OrderService` and `AuditLogService`, demonstrating the correct use of `@Transactional` propagation to achieve this.

**Expected Focus Areas:**
- `OrderService.checkout()` should be annotated with `@Transactional(propagation = Propagation.REQUIRED)`.
- Inside `checkout()`, a try-catch block is used to call the `AuditLogService`.
- `AuditLogService.logAttempt()` MUST be annotated with `@Transactional(propagation = Propagation.REQUIRES_NEW)`. This suspends the outer transaction, commits the log in its own transaction, and resumes the outer transaction.

---

## 5. Spring Boot Actuator & Custom Health Indicator
**Problem Statement:**
Your microservice depends on an external third-party API (e.g., a Payment Gateway like Stripe) to function correctly.
1. Add Spring Boot Actuator dependencies.
2. Write a custom `HealthIndicator` class that periodically pings the third-party API's health endpoint.
3. If the external API is up, your microservice's `/actuator/health` endpoint should show `UP`. If the external API is down or unreachable, your service's health endpoint should automatically show `DOWN`.

**Expected Focus Areas:**
- Implementing the `HealthIndicator` interface.
- Overriding the `health()` method.
- Returning `Health.up().build()` or `Health.down().withDetail("Stripe API", "Unreachable").build()`.

---

## 6. Consuming External REST APIs (RestTemplate / WebClient)
**Problem Statement:**
Your application needs to fetch the current weather from a public weather API: `https://api.weatherapi.com/v1/current.json?q={city}`.
1. Define a `WeatherResponse` DTO mapped to the expected JSON structure.
2. Write a Spring `@Service` method that takes a `String city` as input.
3. Use `RestTemplate` (or `WebClient`) to perform a GET request to the external URL.
4. Return the temperature parsed from the DTO. Handle the scenario where the city is not found (404 Not Found from the external API).

**Expected Focus Areas:**
- Injecting a `RestTemplate` bean.
- Using `restTemplate.getForObject(url, WeatherResponse.class, city)`.
- Catching `HttpClientErrorException.NotFound` or manually checking the HTTP response status.
