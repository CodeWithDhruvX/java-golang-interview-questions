# Practical Coding Questions: Spring Security, JWT, and Testing

*These problems are designed to test your actual coding ability during a technical interview. Try to write out the code or pseudocode before looking at the suggested architectures.*

---

## 1. Stateless Authentication Flow (JWT)
**Problem Statement:**
Design a complete flow for User Registration, Login, and accessing a protected resource using JWT (JSON Web Tokens).
1. **Registration:** Create an endpoint `POST /auth/register` taking `{username, password}`. Detail how the password is saved.
2. **Login:** Create an endpoint `POST /auth/login` taking `{username, password}`. If the credentials match, the server generates and returns a signed JWT containing the user's role.
3. **Protected API:** Create an endpoint `GET /api/dashboard`. This endpoint must intercept the incoming `Authorization: Bearer <token>` header, validate the token signature, and extract the username to populate the `SecurityContext`.

**Expected Focus Areas:**
- **Password Encoding:** Using `BCryptPasswordEncoder` in the Registration service.
- **AuthenticationManager:** Using Spring's `DaoAuthenticationProvider` and a custom `UserDetailsService` to verify login credentials.
- **JwtUtil:** Writing a utility class using `io.jsonwebtoken` to `.signWith(SecretKey)` and `.parseClaimsJws()`.
- **JwtAuthenticationFilter:** A custom `OncePerRequestFilter` that intercepts requests, validates the token, and calls `SecurityContextHolder.getContext().setAuthentication(...)`.

---

## 2. Role-Based Access Control (RBAC)
**Problem Statement:**
You have three user roles in your application: `ROLE_USER`, `ROLE_MANAGER`, and `ROLE_ADMIN`.
1. Configure your `SecurityFilterChain` so that any endpoint starting with `/api/admin/` is strictly accessible only by `ROLE_ADMIN`.
2. Configure your `SecurityFilterChain` so that the `/api/manager/reports` endpoint is accessible by *both* `ROLE_MANAGER` and `ROLE_ADMIN`.
3. Use Method-Level Security (`@PreAuthorize`) on a specific Service method: `public void deleteUser(Long userId)`. This method should only be executable if the currently authenticated user's ID matches the `userId` parameter, OR if the user is an `ADMIN`.

**Expected Focus Areas:**
- `.requestMatchers("/api/admin/**").hasRole("ADMIN")` vs `.hasAnyRole("ADMIN", "MANAGER")`.
- `@EnableMethodSecurity` configuration.
- SpEL (Spring Expression Language): `@PreAuthorize("hasRole('ADMIN') or #userId == authentication.principal.id")`.

---

## 3. Custom Request Filters (Rate Limiting Simulation)
**Problem Statement:**
Your API is being hammered by a specific malicious IP address.
1. Write a custom Spring Security Filter (`BlockIpFilter`) extending `OncePerRequestFilter`.
2. The filter must check the incoming request's IP (`request.getRemoteAddr()`).
3. If the IP matches a specific bad IP (e.g., "192.168.1.100"), the filter should immediately reject the request returning a `403 Forbidden` status code and a custom JSON message, *without* letting the request proceed down the filter chain to the controller.

**Expected Focus Areas:**
- `response.setStatus(HttpServletResponse.SC_FORBIDDEN)`.
- `response.getWriter().write("{ \"error\": \"IP Blocked\" }")`.
- Ensuring `filterChain.doFilter(...)` is **NOT** called if the IP is blocked.
- Registering the filter in `SecurityFilterChain`: `http.addFilterBefore(...)`.

---

## 4. Unit Testing with Mockito (`UserService` logic)
**Problem Statement:**
You have a `UserService` with a `registerUser(UserDto dto)` method. The method:
  a) Checks if the username exists using `UserRepository.existsByUsername()`. If it does, it throws `UserAlreadyExistsException`.
  b) Encodes the pure text password using `PasswordEncoder.encode()`.
  c) Saves the mapped entity using `UserRepository.save()`.

Write a JUnit 5 test class for this service method.
1. Test the **Success Scenario:** Ensure `save()` is called and the password was encoded.
2. Test the **Failure Scenario:** Ensure the `UserAlreadyExistsException` is thrown and `save()` is NEVER called.

**Expected Focus Areas:**
- `@ExtendWith(MockitoExtension.class)`, `@Mock`, `@InjectMocks`.
- `when(userRepository.existsByUsername("admin")).thenReturn(true)`.
- `assertThrows(UserAlreadyExistsException.class, () -> userService.registerUser(dto))`.
- `verify(userRepository, never()).save(any(User.class))`.
- `verify(passwordEncoder, times(1)).encode("rawPassword")`.

---

## 5. Integration Testing a REST API (`MockMvc`)
**Problem Statement:**
Write a Spring Boot Integration test for a completely secured endpoint: `GET /api/user/profile`.
1. The test class should spin up the web context using `@SpringBootTest` and `@AutoConfigureMockMvc`.
2. Write a test case that makes a `mockMvc.perform(get("/api/user/profile"))` request WITHOUT any authentication. Assert that the response status is `401 Unauthorized`.
3. Write a second test case that simulates a logged-in user with the username "john.doe" and role "USER". Assert that the request succeeds (`200 OK`) and the JSON returned contains `{"username": "john.doe"}`.

**Expected Focus Areas:**
- `@SpringBootTest`, `@AutoConfigureMockMvc`.
- `MockMvc` injection.
- Testing unauthorized: `mockMvc.perform(...).andExpect(status().isUnauthorized())`.
- Simulating security identity (Spring Security Test dependency): `@WithMockUser(username="john.doe", roles={"USER"})`.
- JSON assertions: `.andExpect(jsonPath("$.username").value("john.doe"))`.
