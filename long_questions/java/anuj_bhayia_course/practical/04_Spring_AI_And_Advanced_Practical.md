# Practical Coding Questions: Spring AI, AOP, Scheduling, and CI/CD

*These problems are designed to test your actual coding ability during a technical interview. Try to write out the code or pseudocode before looking at the suggested architectures.*

---

## 1. Spring AI: Retrieval-Augmented Generation (RAG) Pipeline
**Problem Statement:**
You are building an internal HR chatbot that answers employee questions based strictly on the company's uploaded PDF handbooks.

1. Inject a `VectorStore` and a `ChatClient` into your `HrChatService`.
2. The user passes a string `query` (e.g., "What is the maternity leave policy?"). Write the method that performs a similarity search on the `VectorStore` using the query, retrieving the top 3 most relevant document chunks.
3. Write a `PromptTemplate` that includes two placeholders: `{context}` and `{question}`.
4. Pass the retrieved chunks into the `{context}` variable, the user's query into `{question}`, and execute the prompt against the `ChatClient`, returning a generated `String`.

**Expected Focus Areas:**
- Using `vectorStore.similaritySearch(SearchRequest.query(query).withTopK(3))`.
- Extracting text from `List<Document>` into a string.
- Using `new PromptTemplate("Answer using {context}. Question: {question}")`.
- Passing a `Map.of("context", contextString, "question", query)` to the template.
- Utilizing `chatClient.prompt(prompt).call().content()`.

---

## 2. Spring AI: Integrating Function Calling (Tools)
**Problem Statement:**
You want your LLM to be able to look up live flight statuses.
1. Write a standard Java method `public FlightStatus getLiveFlightStatus(FlightRequest request)` that simulates querying a live API.
2. Annotate this method as a Spring `@Bean` and annotate it with `@Description` explaining exactly what it does so the LLM understands it. Wrap it in a `Function<FlightRequest, FlightStatus>` interface.
3. Modify the `ChatClient` prompt options so the LLM has permission to invoke this specific function tool if it decides it needs flight data to answer the user's question.

**Expected Focus Areas:**
- `@Description("Get real-time flight status for a given flight number and date")`.
- Generating the bean: `@Bean public Function<FlightRequest, FlightStatus> flightStatusFunction() { ... }`.
- Using `PromptOptionsBuilder`.
- Providing the function name to the client: `.withFunction("flightStatusFunction")`.

---

## 3. Aspect-Oriented Programming (AOP): Custom Logging
**Problem Statement:**
You realize that scattering `System.currentTimeMillis()` across hundreds of service methods to track performance is messy and violates the Single Responsibility Principle.

1. Create a custom Java Annotation called `@LogExecutionTime`.
2. Write a Spring Aspect class.
3. Use the `@Around` advice to intercept any method execution that is annotated with `@LogExecutionTime`.
4. The Aspect should capture the start time, proceed with the original method execution using `ProceedingJoinPoint`, capture the end time, calculate the duration, and log it to the console along with the original method's name.

**Expected Focus Areas:**
- Target method definition: `@Target(ElementType.METHOD) @Retention(RetentionPolicy.RUNTIME) public @interface LogExecutionTime {}`.
- Defining the Aspect: `@Aspect @Component class LoggingAspect`.
- The Advice: `@Around("@annotation(com.example.LogExecutionTime)") public Object logExecutionTime(ProceedingJoinPoint joinPoint) throws Throwable { ... }`.
- Ensuring the Aspect returns `joinPoint.proceed()`.

---

## 4. Task Scheduling & Java Executor Framework
**Problem Statement:**
Your application needs to run a background job that checks for expired subscriptions.
1. Write a Service method annotated with Spring's `@Scheduled` annotation to run exactly at **1:00 AM every day** using a Cron expression.
2. The job pulls 10,000 expired user IDs from the database. Since processing them sequentially takes too long, configure a custom `ThreadPoolTaskExecutor` (Max pool size: 20).
3. Process the 10,000 users concurrently by submitting individual tasks (or batches) to the configured custom ThreadPool using Spring's `@Async`.

**Expected Focus Areas:**
- Scheduling Configuration: `@EnableScheduling` class annotation.
- Method Annotation: `@Scheduled(cron = "0 0 1 * * ?")`.
- Async Configuration: `@EnableAsync` with a custom `@Bean` returning `ThreadPoolTaskExecutor`.
- Annotating the inner processing method with `@Async("myCustomExecutor")` so the main scheduler thread doesn't block.

---

## 5. CI/CD: AWS CodePipeline Buildspec
**Problem Statement:**
Your team uses AWS CodeBuild and CodePipeline. You need to provide the instructions for AWS to pull your source code from GitHub, test it, compile it, and spit out the JAR for deployment.

1. Write a standard `buildspec.yml` file.
2. Under the `install` phase, specify the Java 17 runtime.
3. Under the `build` phase, run the Maven command to cleanly build the package without running the integration tests again (assume they run in a separate PR phase).
4. Under the `artifacts` phase, correctly identify the compiled `.jar` file located in the `target/` directory so CodeBuild knows to pass it to the CodeDeploy phase.

**Expected Focus Areas:**
- Accurate YAML structure: `version: 0.2`, `phases:`, `install:`, `build:`, `artifacts:`.
- Specifying runtimes: `runtime-versions: java: corretto17`.
- The maven build command: `mvn clean package -DskipTests`.
- Artifacts mapping: `files: - 'target/*.jar'`, `discard-paths: yes`.
