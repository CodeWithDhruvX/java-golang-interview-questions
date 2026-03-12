# Spring MVC Interview Questions

## Core Architecture & Workflow

### 1. What is the Front Controller Design Pattern in Spring MVC?
Consider the **DispatcherServlet** as the "Front Controller". It receives *all* incoming HTTP requests and delegates them to the appropriate controllers. This centralizes logic like routing, security, and view selection.

**Explanation:** The Front Controller pattern consolidates request handling through a single entry point, eliminating duplicate logic across multiple controllers and providing a consistent request processing pipeline.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the Front Controller Design Pattern in Spring MVC?
**Your Response:** The Front Controller pattern in Spring MVC is implemented by the DispatcherServlet. Instead of having multiple servlets handling different URL patterns, all requests come to one central servlet - the DispatcherServlet. This servlet then acts like a traffic controller, routing each request to the appropriate handler. This centralizes common logic like security checks, locale handling, and view resolution, making the application more maintainable and consistent.

### 2. Explain the Request Flow in Spring MVC.
1.  **Request received:** `DispatcherServlet` receives the HTTP request.
2.  **Handler Mapping:** `DispatcherServlet` asks `HandlerMapping` to find the correct controller method.
3.  **Controller Execution:** `DispatcherServlet` passes the request to the `HandlerAdapter`, which executes the controller method.
4.  **Model & View:** The controller processes the request and returns a `ModelAndView` object (data + view name).
5.  **View Resolution:** `DispatcherServlet` asks `ViewResolver` to find the actual view file (e.g., JSP, Thymeleaf template).
6.  **Rendering:** The view is rendered with the model data and the response is sent back to the user.

**Explanation:** This request processing pipeline ensures separation of concerns, with each component handling a specific responsibility, making the framework highly modular and extensible.

### How to Explain in Interview (Spoken style format)
**Interviewer:** Can you explain the request flow in Spring MVC?
**Your Response:** When a request comes into a Spring MVC application, it follows a well-defined pipeline. First, the DispatcherServlet receives the request as the front controller. Then it consults HandlerMapping to determine which controller method should handle the request. The HandlerAdapter executes that method, which returns a ModelAndView object containing both the data and the view name. The ViewResolver finds the actual view template, and finally the view renders the data back to the user. This structured approach makes the framework predictable and easy to extend.

### 3. What is the `DispatcherServlet`?
It is the heart of Spring MVC. It acts as the bridge between the Servlet container (like Tomcat) and the Spring application. It manages the entire request processing lifecycle.

**Explanation:** The DispatcherServlet integrates the web container's servlet API with Spring's application context, enabling Spring's dependency injection and MVC features within a standard servlet environment.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is the DispatcherServlet?
**Your Response:** The DispatcherServlet is the central component of Spring MVC - it's essentially the engine that drives the entire framework. It acts as the bridge between the web server's servlet container and Spring's application context. Every request goes through the DispatcherServlet, which coordinates all the other components like handler mappings, controllers, and view resolvers. It's what makes Spring MVC work by managing the complete request lifecycle from when it hits the server until the response is sent back.

## Controllers & Annotations

### 4. What is the difference between `@Controller` and `@RestController`?
- **`@Controller`:** Used for traditional MVC. It typically returns a view name (String). You need `@ResponseBody` on mapped methods to return data directly (JSON/XML).
- **`@RestController`:** A convenience annotation that combines `@Controller` and `@ResponseBody`. It implies that every method returns domain objects instead of a view, writing directly to the HTTP response body (ideal for REST APIs).

**Explanation:** @RestController is a specialized version of @Controller that eliminates the need to add @ResponseBody to each method, making it ideal for building RESTful APIs.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between @Controller and @RestController?
**Your Response:** The main difference is their intended use case. @Controller is designed for traditional web applications where you return view names like 'home' or 'user/profile'. If I want to return JSON data from a @Controller method, I need to add @ResponseBody. @RestController is a convenience annotation that combines @Controller and @ResponseBody, so every method automatically returns data directly to the response body instead of a view name. I use @RestController for building REST APIs and @Controller for traditional MVC applications with server-side rendering.

### 5. What is `@RequestMapping`?
It maps HTTP requests to handler methods of MVC and REST controllers. It can be used at the class level and method level.
- `method`: GET, POST, etc.
- `value`/`path`: URL pattern.
- `produces`/`consumes`: Media types (JSON, XML).

**Explanation:** @RequestMapping provides a flexible way to map URLs to controller methods, supporting various HTTP methods, path variables, and content negotiation.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is @RequestMapping?
**Your Response:** @RequestMapping is Spring's primary annotation for mapping web requests to controller methods. I can use it at the class level to define a base path for all methods in that controller, and at the method level for specific endpoints. It supports various attributes like method to specify HTTP verbs, path or value for URL patterns, and produces/consumes for content negotiation. There are also specialized shortcuts like @GetMapping, @PostMapping, and @PutMapping that make the code more readable.

### 6. Explain `@PathVariable` vs `@RequestParam`.
- **`@PathVariable`:** Extracts values from the URI path itself.
    - Example: `/users/5` -> `@GetMapping("/users/{id}") public User getUser(@PathVariable String id)`
- **`@RequestParam`:** Extracts values from query parameters.
    - Example: `/users?id=5` -> `@GetMapping("/users") public User getUser(@RequestParam String id)`

**Explanation:** @PathVariable is used for RESTful URLs where data is part of the path structure, while @RequestParam handles traditional query string parameters.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between @PathVariable and @RequestParam?
**Your Response:** They extract data from different parts of the URL. @PathVariable pulls values directly from the URL path itself, like getting the user ID from /users/123. It's great for RESTful APIs where the resource identifier is part of the URL structure. @RequestParam extracts values from the query string, like getting parameters from /users?id=123. I use @PathVariable for required resource identifiers and @RequestParam for optional filters, pagination parameters, or search criteria.

### 7. What is `@ModelAttribute`?
It binds a method parameter or method return value to a named model attribute, exposed to a web view. It can be used to prepare data before a controller method is invoked.

**Explanation:** @ModelAttribute ensures that common data is available to multiple views without repetitive code, promoting DRY principles and consistent view data.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is @ModelAttribute?
**Your Response:** @ModelAttribute is used to add attributes to the model that will be available to the view. I can use it in two ways - either on a method to automatically add its return value to the model before every request handler method executes, or on a method parameter to bind request parameters to a model object. It's really useful for populating common data like dropdown lists or user information that multiple views need, without having to repeat the code in every controller method.

## View Resolution & Handling

### 8. What is a `ViewResolver`?
It maps logical view names returned by a controller to actual view resources.
- Example: Returns `"home"` -> Resolver finds `/WEB-INF/views/home.jsp`.

**Explanation:** View Resolvers provide a clean separation between controller logic and view technology, allowing controllers to remain agnostic about the actual view implementation.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is a ViewResolver?
**Your Response:** A ViewResolver is responsible for converting the logical view names returned by controllers into actual view resources. When my controller returns "home", the ViewResolver figures out that this means it should render /WEB-INF/views/home.jsp or home.html depending on the configuration. This decouples my controllers from the specific view technology - I can switch from JSP to Thymeleaf without changing my controller code, just by reconfiguring the ViewResolver.

### 9. How do you handle exceptions in Spring MVC?
- **@ExceptionHandler:** Defined within a controller to handle exceptions thrown by methods in *that* controller.
- **@ControllerAdvice:** Global exception handling. Classes annotated with this can handle exceptions across *all* controllers.

**Explanation:** Spring's exception handling provides a centralized way to manage errors, ensuring consistent error responses and proper logging across the application.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle exceptions in Spring MVC?
**Your Response:** Spring provides two main approaches for exception handling. I can use @ExceptionHandler within a specific controller to handle exceptions thrown by that controller's methods. For global exception handling across all controllers, I use @ControllerAdvice classes that contain @ExceptionHandler methods. This allows me to handle exceptions consistently throughout the application, return appropriate error responses, and implement proper logging. I can also create custom exception classes and map them to specific HTTP status codes.

### 10. What are Interceptors in Spring MVC?
`HandlerInterceptor` allows you to intercept requests:
- `preHandle()`: Before the controller is executed.
- `postHandle()`: After the controller executes but before the view is rendered.
- `afterCompletion()`: After the complete request has finished.
Useful for logging, authentication checks, or adding common model attributes.

**Explanation:** Interceptors provide cross-cutting concerns processing, allowing common logic to be applied to multiple handlers without duplication.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What are Interceptors in Spring MVC?
**Your Response:** Interceptors allow me to intercept requests at different points in the processing pipeline. I can implement three methods: preHandle runs before the controller method executes, postHandle runs after the controller but before the view renders, and afterCompletion runs after the entire request is complete. They're perfect for cross-cutting concerns like logging request execution time, performing authentication checks, or adding common data to all models. Unlike filters, interceptors have access to the Spring context and can work with the ModelAndView object.

## Advanced MVC Topics

### 11. Filters vs. Interceptors?
- **Filters:** Part of the Servlet standard (Java EE). They run *before* the request reaches the `DispatcherServlet`. Good for low-level tasks (compression, encoding, security).
- **Interceptors:** Part of the Spring MVC framework. They run *inside* the Spring context, giving access to the handler (controller) and model. Good for app-specific logic (authorization checks, logging execution time).

**Explanation:** Filters operate at the servlet container level with broader scope, while interceptors operate within the Spring MVC framework with access to Spring-specific components.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What's the difference between Filters and Interceptors?
**Your Response:** The key difference is where they operate in the request processing chain. Filters are part of the Java servlet specification and run before the request even reaches Spring MVC - they're good for low-level concerns like request/response compression, character encoding, or security authentication. Interceptors are Spring-specific and run within the Spring MVC context, so they have access to the controller and model objects. I use filters for protocol-level concerns and interceptors for application-level concerns like business logic validation or performance monitoring.

### 12. How do you validate input in Spring MVC?
Use the Java Bean Validation API (JSR-380/Hibernate Validator).
1.  Add constraints to DTO: `@NotNull`, `@Size`, `@Email`.
2.  Annotate controller argument with `@Valid` or `@Validated`.
3.  Handle `MethodArgumentNotValidException` or check `BindingResult`.

**Explanation:** Bean Validation provides a declarative way to enforce input rules, with Spring automatically handling validation and error reporting.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you validate input in Spring MVC?
**Your Response:** I use the Java Bean Validation API with Spring's validation support. First, I add validation annotations like @NotNull, @Size, @Email to my DTO fields. Then I annotate the controller method parameter with @Valid or @Validated to trigger validation. Spring automatically validates the input and either throws MethodArgumentNotValidException or populates the BindingResult with error details. I can then handle these errors globally with @ControllerAdvice or check the BindingResult in my controller method to return appropriate error responses to the user.

### 13. How do you handle file uploads?
Use `MultipartFile` interface in the controller method.
```java
@PostMapping("/upload")
public String handleFileUpload(@RequestParam("file") MultipartFile file) {
    // ...
}
```
You may need to configure `spring.servlet.multipart.max-file-size` properties.

**Explanation:** Spring's MultipartFile abstraction handles the complexities of file uploads, including multipart request parsing and file streaming.

### How to Explain in Interview (Spoken style format)
**Interviewer:** How do you handle file uploads?
**Your Response:** I handle file uploads using Spring's MultipartFile interface. I create a controller method that accepts MultipartFile as a parameter, typically with @RequestParam. Spring automatically handles the multipart request parsing and gives me access to the file's contents, name, and metadata. I can then process the file - save it to disk, store it in a database, or upload it to cloud storage. I also configure properties like max-file-size and max-request-size in application.properties to control upload limits and prevent denial of service attacks.

### 14. What is Content Negotiation in Spring MVC?
It's the process of determining the best representation for a given resource when multiple representations are available.
- **Strategy:** Accept Header, Request Parameter (e.g., `?format=json`), or File Extension usage.
- Spring uses `ContentNegotiatingViewResolver` to select the correct view (JSON, XML, HTML) based on the client's request.

**Explanation:** Content negotiation enables the same endpoint to serve different formats based on client preferences, promoting API flexibility and reuse.

### How to Explain in Interview (Spoken style format)
**Interviewer:** What is Content Negotiation in Spring MVC?
**Your Response:** Content negotiation is Spring's mechanism for determining the best response format based on what the client requests. The same endpoint can return JSON, XML, or HTML depending on the client's Accept header, a URL parameter like ?format=json, or even the file extension. Spring uses ContentNegotiatingViewResolver to automatically select the appropriate view resolver based on the requested content type. This allows me to build flexible APIs that can serve multiple client types - browsers get HTML, mobile apps get JSON, and legacy systems get XML, all from the same controller method.
