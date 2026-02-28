# Solution: Spring Security, JWT, and Testing

## 1. Stateless Authentication Flow (JWT)

**Solution:**

```java
// 1. JWT Utility Component
@Component
public class JwtUtil {
    private final SecretKey key = Jwts.SIG.HS256.key().build(); // generate strong key
    private final long expiration = 86400000; // 1 day

    public String generateToken(String username, String role) {
        return Jwts.builder()
                .subject(username)
                .claim("role", role)
                .issuedAt(new Date())
                .expiration(new Date(System.currentTimeMillis() + expiration))
                .signWith(key)
                .compact();
    }

    public boolean validateToken(String token) {
        try {
            Jwts.parser().verifyWith(key).build().parseSignedClaims(token);
            return true;
        } catch (JwtException e) {
            return false;
        }
    }
    
    public String extractUsername(String token) {
        return Jwts.parser().verifyWith(key).build()
                .parseSignedClaims(token).getPayload().getSubject();
    }
}

// 2. Authentication Filter
@Component
public class JwtAuthenticationFilter extends OncePerRequestFilter {
    @Autowired private JwtUtil jwtUtil;
    @Autowired private UserDetailsService userDetailsService;

    @Override
    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain chain) 
            throws ServletException, IOException {
        
        String authHeader = request.getHeader("Authorization");
        if (authHeader != null && authHeader.startsWith("Bearer ")) {
            String token = authHeader.substring(7);
            
            if (jwtUtil.validateToken(token) && SecurityContextHolder.getContext().getAuthentication() == null) {
                String username = jwtUtil.extractUsername(token);
                UserDetails userDetails = userDetailsService.loadUserByUsername(username);
                
                UsernamePasswordAuthenticationToken auth = new UsernamePasswordAuthenticationToken(
                        userDetails, null, userDetails.getAuthorities()
                );
                SecurityContextHolder.getContext().setAuthentication(auth);
            }
        }
        chain.doFilter(request, response); // Vital to continue the chain
    }
}

// 3. Controller
@RestController
@RequestMapping("/auth")
public class AuthController {
    
    @Autowired private AuthenticationManager authenticationManager;
    @Autowired private JwtUtil jwtUtil;
    @Autowired private PasswordEncoder passwordEncoder;
    @Autowired private UserRepository userRepository; // Direct access for brevity

    @PostMapping("/register")
    public ResponseEntity<String> register(@RequestBody UserDto dto) {
        User user = new User();
        user.setUsername(dto.getUsername());
        // CRITICAL: Encode the password before saving
        user.setPassword(passwordEncoder.encode(dto.getPassword()));
        user.setRole("ROLE_USER");
        userRepository.save(user);
        return ResponseEntity.ok("User registered successfully");
    }

    @PostMapping("/login")
    public ResponseEntity<String> login(@RequestBody UserDto dto) {
        // Automatically uses UserDetailsService and PasswordEncoder to verify
        Authentication auth = authenticationManager.authenticate(
                new UsernamePasswordAuthenticationToken(dto.getUsername(), dto.getPassword())
        );
        
        // If it passes, generate the token
        String jwt = jwtUtil.generateToken(auth.getName(), auth.getAuthorities().iterator().next().getAuthority());
        return ResponseEntity.ok(jwt);
    }
}
```

---

## 2. Role-Based Access Control (RBAC)

**Solution:**

```java
// 1 & 2. Security Filter Chain Configuration
@Configuration
@EnableWebSecurity
@EnableMethodSecurity // CRITICAL for step 3 to work!
public class SecurityConfig {

    @Bean
    public SecurityFilterChain filterChain(HttpSecurity http, JwtAuthenticationFilter jwtFilter) throws Exception {
        return http
            .csrf(csrf -> csrf.disable()) // Disable for stateless REST APIs
            .sessionManagement(sess -> sess.sessionCreationPolicy(SessionCreationPolicy.STATELESS))
            .authorizeHttpRequests(auth -> auth
                .requestMatchers("/auth/**").permitAll() // Public endpoints
                .requestMatchers("/api/admin/**").hasRole("ADMIN") // Only Admin
                .requestMatchers("/api/manager/reports").hasAnyRole("ADMIN", "MANAGER") // Admin OR Manager
                .anyRequest().authenticated()
            )
            .addFilterBefore(jwtFilter, UsernamePasswordAuthenticationFilter.class)
            .build();
    }
    
    @Bean
    public PasswordEncoder passwordEncoder() {
        return new BCryptPasswordEncoder();
    }
}

// 3. Method-Level Security
@Service
public class UserService {
    
    // SpEL: Allows the operation if the user is an ADMIN, or if the user's ID matches the parameter.
    @PreAuthorize("hasRole('ADMIN') or #userId == authentication.principal.id")
    public void deleteUser(Long userId) {
        // ... deletion logic
    }
}
```

---

## 3. Custom Request Filters (Rate Limiting Simulation)

**Solution:**

```java
@Component
public class BlockIpFilter extends OncePerRequestFilter {

    private final String BLOCKED_IP = "192.168.1.100";

    @Override
    protected void doFilterInternal(HttpServletRequest request, HttpServletResponse response, FilterChain filterChain)
            throws ServletException, IOException {
        
        String clientIp = request.getRemoteAddr();
        
        if (BLOCKED_IP.equals(clientIp)) {
            // Intercept and reject immediately
            response.setStatus(HttpServletResponse.SC_FORBIDDEN);
            response.setContentType("application/json");
            response.getWriter().write("{ \"error\": \"IP Blocked due to malicious activity\" }");
            
            // CRITICAL: Do NOT call filterChain.doFilter() here!
            return; 
        }

        // Allow normal requests to proceed
        filterChain.doFilter(request, response);
    }
}
```

---

## 4. Unit Testing with Mockito (`UserService` logic)

**Solution:**

```java
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.extension.ExtendWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.junit.jupiter.MockitoExtension;
import org.springframework.security.crypto.password.PasswordEncoder;
import static org.mockito.Mockito.*;
import static org.junit.jupiter.api.Assertions.*;

@ExtendWith(MockitoExtension.class) // Instructs JUnit to process the Mockito annotations
public class UserServiceTest {

    @Mock
    private UserRepository userRepository;

    @Mock
    private PasswordEncoder passwordEncoder;

    @InjectMocks
    private UserService userService; // Injects the above mocks into this real instance

    @Test
    void registerUser_Success() {
        // Arrange
        UserDto dto = new UserDto("admin", "rawPassword");
        
        when(userRepository.existsByUsername("admin")).thenReturn(false);
        when(passwordEncoder.encode("rawPassword")).thenReturn("encodedPassword");
        
        // Act
        userService.registerUser(dto);
        
        // Assert
        verify(passwordEncoder, times(1)).encode("rawPassword");
        verify(userRepository, times(1)).save(any(User.class)); // Ensures save was called
    }

    @Test
    void registerUser_ThrowsExceptionWhenUserExists() {
        // Arrange
        UserDto dto = new UserDto("admin", "rawPassword");
        
        when(userRepository.existsByUsername("admin")).thenReturn(true);
        
        // Act & Assert
        assertThrows(UserAlreadyExistsException.class, () -> {
            userService.registerUser(dto);
        });
        
        // Verify that execution stopped before encoding and saving
        verify(passwordEncoder, never()).encode(anyString());
        verify(userRepository, never()).save(any(User.class));
    }
}
```

---

## 5. Integration Testing a REST API (`MockMvc`)

**Solution:**

```java
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.autoconfigure.web.servlet.AutoConfigureMockMvc;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.security.test.context.support.WithMockUser;
import org.springframework.test.web.servlet.MockMvc;

import static org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get;
import static org.springframework.test.web.servlet.result.MockMvcResultMatchers.*;

@SpringBootTest 
@AutoConfigureMockMvc // Bootstraps the full web context and MockMvc
public class UserControllerIntegrationTest {

    @Autowired
    private MockMvc mockMvc;

    @Test
    void getProfile_WithoutAuth_Returns401() throws Exception {
        mockMvc.perform(get("/api/user/profile"))
               .andExpect(status().isUnauthorized());
    }

    @Test
    @WithMockUser(username = "john.doe", roles = {"USER"}) // Simulates standard Spring Security context
    void getProfile_WithAuth_ReturnsProfile() throws Exception {
        mockMvc.perform(get("/api/user/profile"))
               .andExpect(status().isOk())
               .andExpect(jsonPath("$.username").value("john.doe"));
    }
}
```
