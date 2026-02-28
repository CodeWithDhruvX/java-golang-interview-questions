# Spring MVC and REST APIs - Interview Questions and Answers

## 1. Explain the Spring MVC Architecture and the request flow.
**Answer:**
Spring MVC (Model-View-Controller) is a framework built on the Servlet API that provides a robust architecture for developing web applications.

**Request Flow (The Front Controller Pattern):**
1. **DispatcherServlet:** Every HTTP request destined for a Spring MVC application is first intercepted by the `DispatcherServlet` (the Front Controller).
2. **HandlerMapping:** The `DispatcherServlet` consults one or more `HandlerMapping`s to decide which Controller (handler method) should process the request. It maps URLs to specific controller methods (often configured via `@RequestMapping` or `@GetMapping`).
3. **Controller Execution:** The `DispatcherServlet` routes the request to the chosen `Controller`. The Controller executes the business logic (usually by calling service layer methods) and processes the data.
4. **Model and View (Traditional Web):** The Controller returns a `ModelAndView` object to the `DispatcherServlet`. The `Model` contains the application data, and the `View` is the logical string name of the view template (e.g., "home").
5. **ViewResolver (Traditional Web):** The `DispatcherServlet` consults a `ViewResolver` to map the logical view name to a physical view template file (like JSP, Thymeleaf, FreeMarker).
6. **Rendering:** The physical view is rendered using the Model data, and the HTML response is returned to the client.
7. **REST APIs (`@RestController`):** If the controller is a REST controller, steps 4, 5, and 6 are bypassed. Instead, the Java object returned by the handler method is directly serialized into JSON/XML (using libraries like Jackson) by `HttpMessageConverters` and written straight to the HTTP response body.

## 2. What is the difference between `@Controller` and `@RestController` annotations?
**Answer:**
- **`@Controller`:** This annotation marks a class as a traditional Spring MVC controller. Handler methods in a `@Controller` typically return a `String` (logical view name) which is resolved by a `ViewResolver` to render an HTML page. If you want a specific method in a `@Controller` to return data (like JSON) directly to the response body instead of returning a view, you must annotate that specific method with `@ResponseBody`.
- **`@RestController`:** This is a convenience annotation introduced in Spring 4.0. It is a meta-annotation that combines `@Controller` and `@ResponseBody`. When a class is annotated with `@RestController`, every handler method automatically applies `@ResponseBody`. This means the returned object is always serialized into JSON/XML bypassing the view resolution mechanism. It is the standard for building REST APIs.

## 3. How do you handle exceptions globally in a Spring Boot REST API?
**Answer:**
Global exception handling ensures that API clients receive consistent and standardized error responses regardless of where an exception occurs within the application. Spring provides `@ControllerAdvice` and `@ExceptionHandler` for this purpose.

**Mechanism:**
1. **`@ControllerAdvice` / `@RestControllerAdvice`:** You create a specialized class and annotate it with `@RestControllerAdvice` (which implies `@ResponseBody` on all methods). This class acts as a global interceptor for exceptions thrown by any `@RequestMapping` controller.
2. **`@ExceptionHandler`:** Inside this class, you define methods annotated with `@ExceptionHandler(SpecificExceptionClass.class)`. Spring MVC will invoke the matching method whenever that specific exception (or its subclasses) is thrown anywhere in the application.

**Standardizing the Response:**
Instead of returning a plain string or a raw error trace, you should create a custom `ErrorResponse` or `ApiError` class containing fields like `timestamp`, `status`, `error`, `message`, and `path`. The `@ExceptionHandler` method should construct this object and return it wrapped in a `ResponseEntity` to set the appropriate HTTP status code.

```java
@RestControllerAdvice
public class GlobalExceptionHandler {

    @ExceptionHandler(ResourceNotFoundException.class)
    public ResponseEntity<ErrorResponse> handleResourceNotFound(ResourceNotFoundException ex, WebRequest request) {
        ErrorResponse error = new ErrorResponse(LocalDateTime.now(), HttpStatus.NOT_FOUND.value(), ex.getMessage());
        return new ResponseEntity<>(error, HttpStatus.NOT_FOUND);
    }
}
```

## 4. How do you perform data validation in Spring Boot?
**Answer:**
Spring Boot integrates seamlessly with the **Jakarta Bean Validation API (formerly Java Bean Validation / Hibernate Validator)**.

**Steps representing the validation process:**
1. **Add Dependency:** Ensure `spring-boot-starter-validation` is in your `pom.xml` / `build.gradle`.
2. **Annotate DTOs:** Add validation constraint annotations to the fields of your Request DTO (Data Transfer Object). Common annotations include:
    - `@NotNull`, `@NotEmpty`, `@NotBlank`
    - `@Size(min=, max=)`, `@Min`, `@Max`
    - `@Email`, `@Pattern(regexp=...)`
3. **Trigger Validation with `@Valid`:** In your Rest Controller, prepend the `@RequestBody` parameter with the `@Valid` (or Spring's `@Validated`) annotation. This tells Spring to enforce the validation rules before entering the method body.
4. **Handling Validation Errors:** If validation fails, Spring throws a `MethodArgumentNotValidException`. By default, this returns a 400 Bad Request. You should intercept this exception in your `@RestControllerAdvice` global exception handler to parse the `BindingResult` and return a user-friendly JSON map of field names and their specific error messages.

## 5. What is the DTO (Data Transfer Object) Pattern, and why is it used in REST APIs?
**Answer:**
A **Data Transfer Object (DTO)** is an object used to encapsulate data and send it from one subsystem of an application to another, specifically between the Client and the Controller in a REST API.

**Why Use DTOs instead of Database Entities?**
1. **Separation of Concerns:** Entities are mapped to database tables and represent the persistence model. DTOs represent the API contract (presentation model). Mixing them couples the API to the database schema.
2. **Security:** Exposing Entities directly can lead to unintended data leaks (e.g., accidentally serializing a user's password hash or internal IDs). DTOs allow you to strictly control which fields are exposed to the client.
3. **Over-posting/Mass Assignment Protection:** When accepting data, clients shouldn't be able to update fields they aren't authorized to touch (like `isAdmin`, `createdAt`). Using a specific Request DTO containing only allowed fields prevents this.
4. **Tailoring Data:** You often need to aggregate data from multiple Entities to serve a single API response, or flatten complex structures. DTOs allow structuring data exactly as the client UI needs it.
5. **Validation:** Validation constraints belonging to the API endpoint (e.g., "password must be 8 chars for registration") are placed on the DTO, not on the underlying database Entity.

## 6. How do you map between Entity and DTO objects? (ModelMapper & MapStruct)
**Answer:**
Writing repetitive getter/setter boilerplate code to copy data between Entities and DTOs is tedious and error-prone. Mapping libraries automate this.

**1. ModelMapper:**
- An older, reflection-based mapping library.
- It attempts to intelligently figure out the mapping between source and destination fields based on matching names.
- It is easy to set up (just inject the `ModelMapper` bean and call `mapper.map(entity, Dto.class)`), but it uses reflection at runtime, which can be slower for high-throughput applications compared to compile-time generation.

**2. MapStruct (Recommended by many):**
- A newer, annotation processor-based library.
- Instead of reflection at runtime, it generates the actual mapping code (plain getters and setters) at compile time.
- It is extremely fast (no runtime overhead), type-safe (errors are caught during compilation), and very easy to debug because the generated implementation classes are readable Java code.
- You define an interface annotated with `@Mapper(componentModel = "spring")`, declare mapping methods, and MapStruct implements them.

## 7. What is Lombok and how does it reduce boilerplate code? List some common annotations.
**Answer:**
**Project Lombok** is a Java library that automatically plugs into your editor and build tools to generate boilerplate code at compile time, reducing clutter and improving readability.

It works via annotation processing. You add annotations to your classes, and Lombok modifies the Abstract Syntax Tree (AST) before generating the `.class` files to inject getters, setters, constructors, etc.

**Common Annotations:**
- `@Getter` / `@Setter`: Generates getter and setter methods for all non-static fields.
- `@NoArgsConstructor`: Generates a default constructor with no arguments.
- `@AllArgsConstructor`: Generates a constructor taking an argument for every field.
- `@RequiredArgsConstructor`: Generates a constructor for fields that mandate initialization (e.g., `final` fields or fields marked `@NonNull`), excellent for Constructor Dependency Injection.
- `@Data`: A shortcut annotation that bundles `@ToString`, `@EqualsAndHashCode`, `@Getter` on all fields, `@Setter` on all non-final fields, and `@RequiredArgsConstructor`. Often used on DTOs (but be cautious using it on Hibernate Entities due to `hashCode` and `equals` complications).
- `@Builder`: Implements the Builder design pattern, allowing fluent object creation (e.g., `User.builder().name("John").age(30).build()`).
- `@Slf4j`: Instantiates a SLF4J logger instance for the class automatically.

## 8. Explain how to integrate and consume Third-Party APIs using Spring's `RestTemplate` (or `WebClient`).
**Answer:**
Spring provides classes to act as HTTP clients to consume external APIs.

**1. `RestTemplate` (Traditional / Synchronous):**
- Introduced early in Spring, it provides a convenient, synchronous, blocking API to make HTTP requests (GET, POST, PUT, DELETE) to external services.
- It handles the serialization of Java objects to JSON request bodies and deserialization of JSON responses back into Java objects.
- **Usage:** You typically define it as a `@Bean` and autowire it into a service. You use methods like `getForObject()`, `postForEntity()`, or the generic `exchange()` method to execute requests.
- *Note:* While still widely used, Spring has placed `RestTemplate` in maintenance mode conceptually, favoring `WebClient` for general-purpose use.

**2. `WebClient` (Modern / Reactive / Asynchronous):**
- Introduced in Spring 5 WebFlux, it is a non-blocking, reactive client.
- It can be used both synchronously (by blocking the result stream) and asynchronously (using reactive types like `Mono` and `Flux`).
- Based on Project Reactor, it offers a fluent API builder pattern and handles high concurrency much more efficiently than `RestTemplate` because it doesn't block threads while waiting for network responses.

## 9. What is Swagger / OpenAPI, and how do you integrate it into a Spring Boot application?
**Answer:**
**OpenAPI Specification (OAS)** is a standard, language-agnostic interface description for REST APIs. **Swagger** consists of tools (like Web UI and Code Generators) built around the OpenAPI Specification.

Integrating OpenAPI provides automated, interactive API documentation that allows developers to see available endpoints, required request bodies, and expected response formats directly from the browser without looking at the source code.

**Integration (using `springdoc-openapi` for Spring Boot 3 / OpenAPI 3):**
1. **Dependency:** Add the `springdoc-openapi-starter-webmvc-ui` dependency to your project.
2. **Auto-Configuration:** Spring Boot auto-configures the application. Upon starting, the framework scans your `@RestController`s, `@RequestMapping`s, and DTOs to generate the OpenAPI JSON schema on the fly.
3. **UI Access:** You can access the visual, interactive Swagger UI page at `http://localhost:8080/swagger-ui.html`.
4. **Customization:** You can customize the documentation using annotations like `@Tag` (on controllers to group them), `@Operation` (on methods to define summary and description), and `@Schema` (on DTO fields to provide examples).

## 10. Explain the significance of standardizing / reformatting the Response Object in REST APIs.
**Answer:**
When building REST APIs, returning raw domain objects, primitives, or different structural formats for different endpoints creates confusion for frontend developers and external clients. Standardizing the response object means enforcing a consistent JSON structure for *every* API response (both successful data and errors).

**Common Standard Wrapper Structure:**
```json
{
  "success": true,
  "message": "Operation completed successfully",
  "data": { ... payload ... },
  "error": null
}
```

**Benefits:**
1. **Predictability:** Clients always know where to look for data (the `data` field) and whether the request succeeded without strictly relying solely on HTTP status codes (though status codes should still be correct).
2. **Uniform Error Handling:** Frontend clients can write a single, central interceptor to intercept all incoming API responses, check the `success` flag or `error` object, and display global toast notifications or error dialogs.
3. **Metadata Support:** A wrapper easily accommodates adding metadata later, such as pagination details (`page`, `size`, `totalElements`) or correlation IDs for tracing, without breaking the existing contract.
4. **Implementation:** In Spring, this can be achieved manually by having controllers return a `ApiResponse<T>` generic class payload, or automatically using `ResponseBodyAdvice` to intercept and wrap every outgoing response globally before it is written to the HTTP stream.
