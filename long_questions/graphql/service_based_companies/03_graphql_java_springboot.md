# 🟣 GraphQL with Spring Boot & Java — Interview Questions (Service-Based Companies)

This document covers GraphQL integration with Java/Spring Boot, commonly tested in service-based company interviews for Java developers (2–5 years experience).

---

### Q1: How do you set up GraphQL with Spring Boot?

**Answer:**
Spring Boot provides first-class GraphQL support via `spring-boot-starter-graphql` (Spring for GraphQL), available since Spring Boot 2.7+.

**1. Dependencies (Maven):**
```xml
<dependencies>
    <!-- Spring for GraphQL -->
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-graphql</artifactId>
    </dependency>
    <!-- Underlying transport — HTTP + WebSocket -->
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-web</artifactId>
    </dependency>
    <!-- Optional: WebSocket for subscriptions -->
    <dependency>
        <groupId>org.springframework.boot</groupId>
        <artifactId>spring-boot-starter-websocket</artifactId>
    </dependency>
</dependencies>
```

**2. Schema file** (`src/main/resources/graphql/schema.graphqls`):
```graphql
type Query {
  user(id: ID!): User
  users: [User!]!
}

type Mutation {
  createUser(name: String!, email: String!): User!
}

type User {
  id: ID!
  name: String!
  email: String!
  posts: [Post]
}

type Post {
  id: ID!
  title: String!
}
```

**3. Controller (Resolver):**
```java
@Controller
public class UserController {

    @Autowired
    private UserService userService;

    @QueryMapping                          // maps to Query.user
    public User user(@Argument String id) {
        return userService.findById(id);
    }

    @QueryMapping                          // maps to Query.users
    public List<User> users() {
        return userService.findAll();
    }

    @MutationMapping                       // maps to Mutation.createUser
    public User createUser(@Argument String name, @Argument String email) {
        return userService.create(name, email);
    }

    @SchemaMapping(typeName = "User", field = "posts")  // resolves User.posts
    public List<Post> posts(User user) {
        return userService.getPostsByUser(user.getId());
    }
}
```

**4. Application properties:**
```yaml
spring:
  graphql:
    graphiql:
      enabled: true      # Enables /graphiql browser IDE
    path: /graphql       # Endpoint path
    schema:
      locations: classpath:graphql/**/   # Where to look for schema files
      file-extensions: .graphqls, .gqls
```

---

### Q2: Explain `@QueryMapping`, `@MutationMapping`, `@SchemaMapping`, and `@SubscriptionMapping` in Spring for GraphQL.

**Answer:**
These are Spring for GraphQL annotations that map controller methods to GraphQL schema operations.

```java
@Controller
public class ProductController {

    // @QueryMapping — convention-based (method name must match field name)
    @QueryMapping                       // maps to Query.products
    public List<Product> products() {
        return productService.findAll();
    }

    @QueryMapping                       // maps to Query.product
    public Product product(@Argument String id) {
        return productService.findById(id);
    }

    // @MutationMapping — maps to Mutation fields
    @MutationMapping
    public Product createProduct(@Argument CreateProductInput input) {
        return productService.create(input);
    }

    @MutationMapping
    public Boolean deleteProduct(@Argument String id) {
        return productService.delete(id);
    }

    // @SchemaMapping — explicit field mapping
    @SchemaMapping(typeName = "Product", field = "category")
    public Category category(Product product) {
        // parent = Product (first argument)
        return categoryService.findById(product.getCategoryId());
    }

    // Shorthand for @SchemaMapping(typeName = "Product")
    @SchemaMapping
    public Category category(Product product) {
        return categoryService.findById(product.getCategoryId());
    }

    // @SubscriptionMapping — returns a reactive publisher
    @SubscriptionMapping
    public Flux<Product> productCreated() {
        return productService.getProductCreatedStream();
    }
}
```

**Input type mapping:**
```java
// GraphQL input maps to Java record or class
public record CreateProductInput(String name, BigDecimal price, String categoryId) {}

// Spring automatically maps { name, price, categoryId } from GraphQL argument
@MutationMapping
public Product createProduct(@Argument CreateProductInput input) { ... }
```

---

### Q3: How do you solve the N+1 problem in Spring for GraphQL using `@BatchMapping`?

**Answer:**
Spring for GraphQL provides the `@BatchMapping` annotation to solve N+1 by automatically batching resolver calls.

**Without batching (N+1 problem):**
```java
// Called ONCE PER AUTHOR — causes N+1!
@SchemaMapping
public Author author(Book book) {
    return authorService.findById(book.getAuthorId());
}
```

**With `@BatchMapping` (batch all at once):**
```java
@Controller
public class BookController {

    @BatchMapping                              // Spring batches automatically!
    public Map<Book, Author> author(List<Book> books) {
        // books = ALL book objects from the current query
        List<String> authorIds = books.stream()
            .map(Book::getAuthorId)
            .collect(Collectors.toList());

        Map<String, Author> authorMap = authorService
            .findAllByIds(authorIds)
            .stream()
            .collect(Collectors.toMap(Author::getId, Function.identity()));

        // Return Map<Book, Author> — Spring distributes results
        return books.stream()
            .collect(Collectors.toMap(
                Function.identity(),
                book -> authorMap.get(book.getAuthorId())
            ));
    }

    // Alternative — return List (preserving order)
    @BatchMapping
    public List<Author> author(List<Book> books) {
        List<String> ids = books.stream().map(Book::getAuthorId).toList();
        Map<String, Author> map = authorService.findAllByIds(ids)
            .stream().collect(Collectors.toMap(Author::getId, Function.identity()));
        return ids.stream().map(map::get).toList();   // maintain order!
    }
}
```

**Result:** 2 DB queries regardless of list size — Spring for GraphQL uses Project Reactor internally to batch the calls.

---

### Q4: How do you handle authentication and authorization in Spring for GraphQL?

**Answer:**
Spring for GraphQL integrates naturally with **Spring Security**.

**1. Method-Level Security (recommended for GraphQL):**
```java
@Configuration
@EnableMethodSecurity
public class SecurityConfig {

    @Bean
    public SecurityFilterChain securityFilterChain(HttpSecurity http) throws Exception {
        return http
            .csrf(csrf -> csrf.disable())
            .authorizeHttpRequests(auth -> auth
                .requestMatchers("/graphql").permitAll()  // GraphQL endpoint open
                .anyRequest().authenticated()
            )
            .httpBasic(Customizer.withDefaults())
            .build();
    }
}
```

```java
@Controller
public class OrderController {

    @QueryMapping
    @PreAuthorize("isAuthenticated()")           // Must be logged in
    public List<Order> myOrders(Authentication auth) {
        return orderService.findByUser(auth.getName());
    }

    @MutationMapping
    @PreAuthorize("hasRole('ADMIN')")            // Must have ADMIN role
    public Boolean cancelOrder(@Argument String orderId) {
        return orderService.cancel(orderId);
    }

    // Access SecurityContext in resolver
    @QueryMapping
    public UserProfile me() {
        Authentication auth = SecurityContextHolder.getContext().getAuthentication();
        if (auth == null || !auth.isAuthenticated()) {
            throw new AccessDeniedException("Not authenticated");
        }
        return userService.findByUsername(auth.getName());
    }
}
```

**2. Custom Error Handling for security exceptions:**
```java
@Component
public class GraphQLSecurityExceptionHandler implements DataFetcherExceptionResolver {

    @Override
    public Mono<List<GraphQLError>> resolveException(Throwable ex, DataFetchingEnvironment env) {
        if (ex instanceof AccessDeniedException) {
            return Mono.just(List.of(
                GraphqlErrorBuilder.newError(env)
                    .message("Access denied: " + ex.getMessage())
                    .errorType(ErrorType.FORBIDDEN)
                    .build()
            ));
        }
        return Mono.empty();   // Let other handlers process
    }
}
```

---

### Q5: How do you write integration tests for Spring for GraphQL?

**Answer:**
Spring for GraphQL provides `GraphQlTester` for testing without starting a full HTTP server.

**Test setup:**
```java
@SpringBootTest
@AutoConfigureHttpGraphQlTester
class UserControllerTest {

    @Autowired
    private HttpGraphQlTester graphQlTester;

    @Test
    void shouldFetchUser() {
        graphQlTester
            .document("""
                query GetUser($id: ID!) {
                    user(id: $id) {
                        id
                        name
                        email
                    }
                }
            """)
            .variable("id", "1")
            .execute()
            .path("user.name").entity(String.class).isEqualTo("John Doe")
            .path("user.email").entity(String.class).isEqualTo("john@example.com");
    }

    @Test
    void shouldCreateUser() {
        graphQlTester
            .document("""
                mutation CreateUser($name: String!, $email: String!) {
                    createUser(name: $name, email: $email) {
                        id
                        name
                    }
                }
            """)
            .variable("name", "Jane")
            .variable("email", "jane@test.com")
            .execute()
            .path("createUser.name").entity(String.class).isEqualTo("Jane");
    }

    @Test
    void shouldReturnErrorForMissingUser() {
        graphQlTester
            .document("query { user(id: \"999\") { name } }")
            .execute()
            .errors()
            .expect(error -> error.getMessage().contains("not found"));
    }
}
```

**Controller-slice test (no full context):**
```java
@GraphQlTest(UserController.class)
class UserControllerSliceTest {

    @Autowired
    private GraphQlTester graphQlTester;

    @MockBean
    private UserService userService;

    @Test
    void shouldQueryUsers() {
        given(userService.findAll())
            .willReturn(List.of(new User("1", "Alice", "alice@test.com")));

        graphQlTester
            .document("{ users { name } }")
            .execute()
            .path("users[0].name").entity(String.class).isEqualTo("Alice");
    }
}
```

---

*Prepared for technical screening rounds at service-based companies (TCS, Infosys, Wipro, Capgemini, HCL, Cognizant — Java/Spring Boot roles).*
