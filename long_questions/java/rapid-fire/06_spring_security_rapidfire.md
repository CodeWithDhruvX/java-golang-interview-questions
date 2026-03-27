# 🔐 Spring Security & REST API Security (Rapid-Fire)

> 🔑 **Master Keyword:** **"AAJCP"** → Authentication, Authorization, JWT, CSRF/CORS, Passwords

---

## 🛡️ Section 1: Core Security Concepts

### Q1: Authentication vs Authorization?
🔑 **Keyword: "WHO-WHAT"** → Auth**N**=WHO you are, Auth**Z**=WHAT you can do

| Concept | Question | Example |
|---|---|---|
| **Authentication** | "Who are you?" | Login with username + password |
| **Authorization** | "What are you allowed to do?" | Admin vs User access |

---

### Q2: Spring Security Filter Chain Architecture?
🔑 **Keyword: "DFC"** → DelegatingFilterProxy→FilterChainProxy→SecurityFilterChain

```
HTTP Request
    → DelegatingFilterProxy (Servlet-level bridge to Spring)
        → FilterChainProxy (selects correct chain)
            → SecurityFilterChain (list of security filters)
                → UsernamePasswordAuthenticationFilter
                → JwtAuthenticationFilter (custom)
                → ExceptionTranslationFilter
                → FilterSecurityInterceptor
                    → Controller (finally!)
```

---

### Q3: `SecurityContext` and `SecurityContextHolder`?
🔑 **Keyword: "CTH"** → Context=Auth-storage, Holder=ThreadLocal-accessor

```java
// Get current authenticated user from anywhere
Authentication auth = SecurityContextHolder.getContext().getAuthentication();
String username = ((UserDetails) auth.getPrincipal()).getUsername();
List<GrantedAuthority> roles = (List) auth.getAuthorities();
```

- **SecurityContext:** Holds current `Authentication` object
- **SecurityContextHolder:** ThreadLocal-backed access to SecurityContext

---

### Q4: `UserDetailsService` interface?
🔑 **Keyword: "LUN"** → Load-User-byUsername bridge to your user store

```java
@Service
public class MyUserDetailsService implements UserDetailsService {
    @Autowired
    private UserRepository repo;

    @Override
    public UserDetails loadUserByUsername(String username) throws UsernameNotFoundException {
        User user = repo.findByUsername(username)
            .orElseThrow(() -> new UsernameNotFoundException("Not found: " + username));
        return org.springframework.security.core.userdetails.User
            .withUsername(user.getUsername())
            .password(user.getPassword())
            .roles(user.getRole())
            .build();
    }
}
```

---

### Q5: Spring Security Configuration (Modern — Spring Security 5.7+)?
🔑 **Keyword: "SFB"** → SecurityFilterChain as Bean (no more WebSecurityConfigurerAdapter)

```java
@Configuration
@EnableWebSecurity
public class SecurityConfig {

    @Bean
    public SecurityFilterChain filterChain(HttpSecurity http) throws Exception {
        http
            .csrf(csrf -> csrf.disable())           // disable for REST APIs
            .authorizeHttpRequests(auth -> auth
                .requestMatchers("/api/public/**").permitAll()
                .requestMatchers("/api/admin/**").hasRole("ADMIN")
                .anyRequest().authenticated()
            )
            .sessionManagement(s -> s.sessionCreationPolicy(SessionCreationPolicy.STATELESS))
            .addFilterBefore(jwtFilter, UsernamePasswordAuthenticationFilter.class);
        return http.build();
    }

    @Bean
    public PasswordEncoder passwordEncoder() {
        return new BCryptPasswordEncoder();
    }
}
```

---

## 🔑 Section 2: JWT Authentication

### Q6: What is JWT? Structure?
🔑 **Keyword: "HPS"** → Header.Payload.Signature (3 parts, dot-separated)

```
eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1c2VyMSJ9.abc123signature
     └── Header ──┘  └──── Payload ────┘ └── Signature ──┘
```

- **Header:** Algorithm (`HS256`, `RS256`) + type (`JWT`)
- **Payload:** Claims (sub=username, exp=expiry, roles)
- **Signature:** HMAC(header + payload + secret) — verifies integrity

---

### Q7: JWT Flow in Spring Security?

🔑 **Keyword: "LVFS"** → Login→Validate→Filter→SecurityContext→Success

```
1. POST /auth/login {username, password}
2. Server validates credentials → generates JWT
3. Client stores JWT (localStorage / cookie)
4. Client sends: Authorization: Bearer <jwt-token>
5. JwtFilter extracts token → validates signature → sets SecurityContext
6. API request processed (stateless — no session!)
```

```java
// Custom JWT Filter
@Component
public class JwtAuthFilter extends OncePerRequestFilter {
    @Override
    protected void doFilterInternal(HttpServletRequest req, HttpServletResponse res, FilterChain chain) {
        String header = req.getHeader("Authorization");
        if (header != null && header.startsWith("Bearer ")) {
            String token = header.substring(7);
            String username = jwtUtil.extractUsername(token);
            if (username != null && SecurityContextHolder.getContext().getAuthentication() == null) {
                UserDetails userDetails = userDetailsService.loadUserByUsername(username);
                if (jwtUtil.validateToken(token, userDetails)) {
                    UsernamePasswordAuthenticationToken auth =
                        new UsernamePasswordAuthenticationToken(userDetails, null, userDetails.getAuthorities());
                    SecurityContextHolder.getContext().setAuthentication(auth);
                }
            }
        }
        chain.doFilter(req, res);
    }
}
```

---

### Q8: JWT vs Session-based Auth?
🔑 **Keyword: "JSS"** → JWT=Stateless, Session=Server-state

| Feature | JWT | Session |
|---|---|---|
| Server state | ❌ None (stateless) | ✅ Stores session |
| Scalability | ✅ Easy (no session sync) | ❌ Needs sticky sessions / Redis |
| Token revocation | ❌ Hard (valid until expiry) | ✅ Easy (delete session) |
| Best for | Microservices, REST APIs | Traditional web apps |

---

## 🌐 Section 3: CORS, CSRF & Security Attacks

### Q9: CSRF — What is it and when to disable?
🔑 **Keyword: "CSRF-C-disable"** → CSRF=Cookie-based attack, safe to disable for stateless JWT APIs

- **CSRF Attack:** Evil website tricks logged-in user's browser to make state-changing requests
- **Protection:** CSRF token in form/request header — server validates it
- **Disable for REST APIs?** YES — because stateless JWT APIs don't use session cookies → CSRF not applicable

```java
.csrf(csrf -> csrf.disable()) // safe for REST + JWT
```

---

### Q10: CORS — What and How to configure?
🔑 **Keyword: "CORS-Allow"** → CORS=browser-prevents-cross-domain, configure=whitelist-origins

```java
// Global CORS config
@Bean
public CorsConfigurationSource corsConfigurationSource() {
    CorsConfiguration config = new CorsConfiguration();
    config.setAllowedOrigins(List.of("https://myfrontend.com"));
    config.setAllowedMethods(List.of("GET", "POST", "PUT", "DELETE"));
    config.setAllowedHeaders(List.of("Authorization", "Content-Type"));
    UrlBasedCorsConfigurationSource source = new UrlBasedCorsConfigurationSource();
    source.registerCorsConfiguration("/**", config);
    return source;
}
```

---

### Q11: CORS vs CSRF?
🔑 **Keyword: "CORS-allow, CSRF-block"** → CORS=allow-legitimate-cross-origin, CSRF=block-malicious-requests

| | CORS | CSRF |
|---|---|---|
| What | Browser security restricts cross-domain requests | Attack: tricks browser to make unwanted requests |
| Protects | Client (data leakage) | Server (unauthorized actions) |
| Configure to | **Allow** trusted origins | **Block** unauthorized requests |

---

## 🔒 Section 4: Method Security

### Q12: Method-Level Security Annotations?
🔑 **Keyword: "SPA"** → Secured/PreAuthorize/Annotations

```java
@EnableMethodSecurity  // enable in configuration

// @PreAuthorize — most powerful (SpEL support)
@PreAuthorize("hasRole('ADMIN')")
public void deleteUser(Long id) { }

@PreAuthorize("hasRole('USER') and #id == authentication.principal.id")
public User getProfile(Long id) { }  // user can only get their own profile

// @PostAuthorize — check after method runs
@PostAuthorize("returnObject.owner == authentication.name")
public Document getDocument(Long id) { }

// @Secured — simple role check (legacy)
@Secured("ROLE_ADMIN")
public void adminTask() { }
```

---

### Q13: Password Encoding — BCrypt?
🔑 **Keyword: "BSW"** → BCrypt=Salted+Work-factor (one-way hash)

```java
@Bean
public PasswordEncoder passwordEncoder() {
    return new BCryptPasswordEncoder(); // work factor = 10 by default
}

// Encode on registration
String encoded = passwordEncoder.encode("myPassword123");
// Verify on login
boolean matches = passwordEncoder.matches("myPassword123", encodedFromDB);
```

**Never store plain-text passwords!** BCrypt has built-in salt → same password produces different hashes.

---

### Q14: OAuth2 Flow?
🔑 **Keyword: "UCAST"** → User→Client→AuthServer→Token→ResourceServer

```
1. User tries to access resource via Client app
2. Client redirects to Authorization Server (Google/GitHub)
3. User authenticates + grants permission
4. Authorization Server issues Access Token (+ optionally Refresh Token)
5. Client uses Access Token to call Resource Server API
```

Spring Boot OAuth2 login:
```java
.oauth2Login(oauth2 -> oauth2
    .authorizationEndpoint(auth -> auth.baseUri("/oauth2/authorize"))
    .redirectionEndpoint(redirect -> redirect.baseUri("/oauth2/callback/*"))
)
```

---

### Q15: `AuthenticationProvider` — Custom Authentication?
🔑 **Keyword: "APCA"** → AuthenticationProvider=Custom-Auth-logic

```java
@Component
public class CustomAuthProvider implements AuthenticationProvider {
    @Override
    public Authentication authenticate(Authentication auth) throws AuthenticationException {
        String username = auth.getName();
        String password = auth.getCredentials().toString();
        // custom validation (OTP, biometric, etc.)
        if (isValid(username, password)) {
            return new UsernamePasswordAuthenticationToken(username, password, getAuthorities());
        }
        throw new BadCredentialsException("Invalid credentials");
    }

    @Override
    public boolean supports(Class<?> auth) {
        return UsernamePasswordAuthenticationToken.class.isAssignableFrom(auth);
    }
}
```

---

### Q16: Role Hierarchy?
🔑 **Keyword: "RHIA"** → Role-Hierarchy=Inherit-below-roles-Automatically

```java
@Bean
public RoleHierarchy roleHierarchy() {
    RoleHierarchyImpl hierarchy = new RoleHierarchyImpl();
    hierarchy.setHierarchy("ROLE_ADMIN > ROLE_MODERATOR > ROLE_USER > ROLE_GUEST");
    return hierarchy;
}
// ADMIN automatically passes ROLE_USER checks — no need to assign both
```

---

### Q17: Adding Custom Filter to Security Chain?
🔑 **Keyword: "ABB"** → AddFilterBefore/After/At-certain-filter

```java
// JWT filter registered BEFORE the standard auth filter
http.addFilterBefore(jwtAuthFilter, UsernamePasswordAuthenticationFilter.class);

// Logging filter AFTER authentication
http.addFilterAfter(loggingFilter, UsernamePasswordAuthenticationFilter.class);
```

---

*End of File — Spring Security & REST API Security*
