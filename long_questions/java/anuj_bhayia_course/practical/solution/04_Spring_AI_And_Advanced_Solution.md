# Solution: Spring AI, AOP, Scheduling, and CI/CD

## 1. Spring AI: Retrieval-Augmented Generation (RAG) Pipeline

*(Note: To run these AI examples for free locally, replace the Spring AI OpenAI starter dependency with the `spring-ai-ollama-spring-boot-starter` dependency in your `pom.xml` and install [Ollama](https://ollama.com/) with a model like `llama3` or `mistral`)*

**Solution:**

```java
import org.springframework.ai.chat.client.ChatClient;
import org.springframework.ai.chat.prompt.Prompt;
import org.springframework.ai.chat.prompt.PromptTemplate;
import org.springframework.ai.document.Document;
import org.springframework.ai.vectorstore.SearchRequest;
import org.springframework.ai.vectorstore.VectorStore;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

@Service
public class HrChatService {

    private final VectorStore vectorStore;
    private final ChatClient chatClient;

    public HrChatService(VectorStore vectorStore, ChatClient.Builder chatClientBuilder) {
        this.vectorStore = vectorStore;
        this.chatClient = chatClientBuilder.build();
    }

    public String askQuestion(String query) {
        // 1. Retrieval (The 'R' in RAG) - Find top 3 relevant chunks
        List<Document> similarDocuments = vectorStore.similaritySearch(
                SearchRequest.query(query).withTopK(3)
        );

        // Convert the matched documents into a single contextual string
        String contextString = similarDocuments.stream()
                .map(Document::getContent)
                .collect(Collectors.joining("\n\n"));

        // 2. Augmentation (The 'A' in RAG) - Inject context into Prompt
        String template = "You are a helpful HR Assistant. Answer the user's question using ONLY the following context.\nContext:\n{context}\n\nQuestion: {question}";
        PromptTemplate promptTemplate = new PromptTemplate(template);
        
        // Render the template with the dynamic values
        Prompt prompt = promptTemplate.create(Map.of(
                "context", contextString,
                "question", query
        ));

        // 3. Generation (The 'G' in RAG) - Ask the LLM
        return chatClient.prompt(prompt)
                .call()
                .content();
    }
}
```

---

## 2. Spring AI: Integrating Function Calling (Tools)

**Solution:**

```java
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.context.annotation.Description;
import java.util.function.Function;

// 1. The Request/Response Classes
public record FlightRequest(String flightNumber, String date) {}
public record FlightStatus(String status, String gate, String estimatedArrival) {}

@Configuration
public class FlightToolsConfig {

    // 2. The Tool Definition
    @Bean
    @Description("Get real-time flight status (gate, arrival time, delays) for a given flight number and date.")
    public Function<FlightRequest, FlightStatus> flightStatusFunction() {
        return request -> {
            // Live API call simulation
            System.out.println("LLM triggered tool to lookup flight: " + request.flightNumber());
            return new FlightStatus("On Time", "B12", "14:30 PM");
        };
    }
}

// 3. Using the Tool in the ChatClient
@Service
public class TravelAgentService {
    
    private final ChatClient chatClient;

    public TravelAgentService(ChatClient.Builder builder) {
        this.chatClient = builder
                .defaultSystem("You are a helpful travel assistant.")
                // Grants the model permission to use exactly this bean if it needs to
                .defaultFunctions("flightStatusFunction") 
                .build();
    }

    public String ask(String userInput) {
        return chatClient.prompt()
                .user(userInput)
                .call()
                .content();
    }
}
```

---

## 3. Aspect-Oriented Programming (AOP): Custom Logging

**Solution:**

```java
// 1. The Custom Annotation
import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

@Target(ElementType.METHOD)
@Retention(RetentionPolicy.RUNTIME)
public @interface LogExecutionTime {
}

// 2 & 3. The Aspect
import org.aspectj.lang.ProceedingJoinPoint;
import org.aspectj.lang.annotation.Around;
import org.aspectj.lang.annotation.Aspect;
import org.springframework.stereotype.Component;

@Aspect
@Component
public class LoggingAspect {

    // Advice indicating WHICH methods to intercept (Pointcut expression)
    @Around("@annotation(com.example.LogExecutionTime)")
    public Object logExecutionTime(ProceedingJoinPoint joinPoint) throws Throwable {
        
        long start = System.currentTimeMillis();
        
        // 4. Continue execution of the actual method
        Object proceed = joinPoint.proceed();
        
        long executionTime = System.currentTimeMillis() - start;
        
        String methodName = joinPoint.getSignature().getName();
        System.out.println("AOP Log: " + methodName + " executed in " + executionTime + "ms");
        
        return proceed; // Crucial: Return the actual result to the caller!
    }
}
```

---

## 4. Task Scheduling & Java Executor Framework

**Solution:**

```java
@Configuration
@EnableScheduling
@EnableAsync
// 3. Configure the custom ThreadPool
public class AsyncConfig {
    
    @Bean(name = "jobExecutor")
    public Executor taskExecutor() {
        ThreadPoolTaskExecutor executor = new ThreadPoolTaskExecutor();
        executor.setCorePoolSize(10);
        executor.setMaxPoolSize(20);
        executor.setQueueCapacity(5000);
        executor.setThreadNamePrefix("JobThread-");
        executor.initialize();
        return executor;
    }
}

@Service
public class SubscriptionService {

    private final UserRepository userRepository;
    private final EmailService emailService; // Assuming this has the inner @Async method

    public SubscriptionService(UserRepository userRepository, EmailService emailService) {
        this.userRepository = userRepository;
        this.emailService = emailService;
    }

    // 1. Schedule at 1:00 AM every day
    // "Second Minute Hour DayOfMonth Month DayOfWeek"
    @Scheduled(cron = "0 0 1 * * ?")
    public void checkExpiredSubscriptionsJob() {
        List<User> expiredUsers = userRepository.findExpiredSubscriptions();
        
        // 2. Submit batches asynchronously without blocking the single scheduler thread
        for (User user : expiredUsers) {
            emailService.sendExpirationWarningAsync(user);
        }
    }
}

@Service
public class EmailService {
    
    // 3. Runs completely independently in the custom configured thread pool
    @Async("jobExecutor")
    public void sendExpirationWarningAsync(User user) {
        // Send actual email here. Takes time...
        System.out.println("Processing email for " + user.getId() + " on thread: " + Thread.currentThread().getName());
    }
}
```

---

## 5. CI/CD: AWS CodePipeline Buildspec

**Solution:**

```yaml
# buildspec.yml (located in root of the project repository)
version: 0.2

phases:
  install:
    runtime-versions:
      # 2. Specify the exact language/version needed
      java: corretto17
  
  pre_build:
    commands:
      - echo "- Starting the pre-build phase"
      - mvn clean # Ensure workspace is empty of old artifacts
  
  build:
    commands:
      - echo "- Starting the build phase"
      # 3. Build the JAR, but skip integration/unit tests for speed if they ran earlier in the pipeline
      - mvn package -DskipTests
  
  post_build:
    commands:
      - echo "- Build completed successfully"

artifacts:
  files:
    # 4. Grab the successfully compiled JAR file to pass to CodeDeploy (or Docker phase)
    - 'target/*.jar'
  discard-paths: yes # Flattens the file structure so CodeDeploy finds it easily
```
