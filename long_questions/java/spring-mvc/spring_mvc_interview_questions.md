# Spring MVC Interview Questions

## Core Architecture & Workflow

### 1. What is the Front Controller Design Pattern in Spring MVC?
Consider the **DispatcherServlet** as the "Front Controller". It receives *all* incoming HTTP requests and delegates them to the appropriate controllers. This centralizes logic like routing, security, and view selection.

### 2. Explain the Request Flow in Spring MVC.
1.  **Request received:** `DispatcherServlet` receives the HTTP request.
2.  **Handler Mapping:** `DispatcherServlet` asks `HandlerMapping` to find the correct controller method.
3.  **Controller Execution:** `DispatcherServlet` passes the request to the `HandlerAdapter`, which executes the controller method.
4.  **Model & View:** The controller processes the request and returns a `ModelAndView` object (data + view name).
5.  **View Resolution:** `DispatcherServlet` asks `ViewResolver` to find the actual view file (e.g., JSP, Thymeleaf template).
6.  **Rendering:** The view is rendered with the model data and the response is sent back to the user.

### 3. What is the `DispatcherServlet`?
It is the heart of Spring MVC. It acts as the bridge between the Servlet container (like Tomcat) and the Spring application. It manages the entire request processing lifecycle.

## Controllers & Annotations

### 4. What is the difference between `@Controller` and `@RestController`?
- **`@Controller`:** Used for traditional MVC. It typically returns a view name (String). You need `@ResponseBody` on mapped methods to return data directly (JSON/XML).
- **`@RestController`:** A convenience annotation that combines `@Controller` and `@ResponseBody`. It implies that every method returns domain objects instead of a view, writing directly to the HTTP response body (ideal for REST APIs).

### 5. What is `@RequestMapping`?
It maps HTTP requests to handler methods of MVC and REST controllers. It can be used at the class level and method level.
- `method`: GET, POST, etc.
- `value`/`path`: URL pattern.
- `produces`/`consumes`: Media types (JSON, XML).

### 6. Explain `@PathVariable` vs `@RequestParam`.
- **`@PathVariable`:** Extracts values from the URI path itself.
    - Example: `/users/5` -> `@GetMapping("/users/{id}") public User getUser(@PathVariable String id)`
- **`@RequestParam`:** Extracts values from query parameters.
    - Example: `/users?id=5` -> `@GetMapping("/users") public User getUser(@RequestParam String id)`

### 7. What is `@ModelAttribute`?
It binds a method parameter or method return value to a named model attribute, exposed to a web view. It can be used to prepare data before a controller method is invoked.

## View Resolution & Handling

### 8. What is a `ViewResolver`?
It maps logical view names returned by a controller to actual view resources.
- Example: Returns `"home"` -> Resolver finds `/WEB-INF/views/home.jsp`.

### 9. How do you handle exceptions in Spring MVC?
- **@ExceptionHandler:** Defined within a controller to handle exceptions thrown by methods in *that* controller.
- **@ControllerAdvice:** Global exception handling. Classes annotated with this can handle exceptions across *all* controllers.

### 10. What are Interceptors in Spring MVC?
`HandlerInterceptor` allows you to intercept requests:
- `preHandle()`: Before the controller is executed.
- `postHandle()`: After the controller executes but before the view is rendered.
- `afterCompletion()`: After the complete request has finished.
Useful for logging, authentication checks, or adding common model attributes.

## Advanced MVC Topics

### 11. Filters vs. Interceptors?
- **Filters:** Part of the Servlet standard (Java EE). They run *before* the request reaches the `DispatcherServlet`. Good for low-level tasks (compression, encoding, security).
- **Interceptors:** Part of the Spring MVC framework. They run *inside* the Spring context, giving access to the handler (controller) and model. Good for app-specific logic (authorization checks, logging execution time).

### 12. How do you validate input in Spring MVC?
Use the Java Bean Validation API (JSR-380/Hibernate Validator).
1.  Add constraints to DTO: `@NotNull`, `@Size`, `@Email`.
2.  Annotate controller argument with `@Valid` or `@Validated`.
3.  Handle `MethodArgumentNotValidException` or check `BindingResult`.

### 13. How do you handle file uploads?
Use `MultipartFile` interface in the controller method.
```java
@PostMapping("/upload")
public String handleFileUpload(@RequestParam("file") MultipartFile file) {
    // ...
}
```
You may need to configure `spring.servlet.multipart.max-file-size` properties.

### 14. What is Content Negotiation in Spring MVC?
It's the process of determining the best representation for a given resource when multiple representations are available.
- **Strategy:** Accept Header, Request Parameter (e.g., `?format=json`), or File Extension usage.
- Spring uses `ContentNegotiatingViewResolver` to select the correct view (JSON, XML, HTML) based on the client's request.
