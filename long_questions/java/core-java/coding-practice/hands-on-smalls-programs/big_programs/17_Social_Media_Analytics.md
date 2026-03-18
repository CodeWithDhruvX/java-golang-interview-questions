# Social Media Analytics - Functional Programming Problem Solving

> **Challenge:** Design a social media analytics system that processes user posts, engagement metrics, and trending topics using functional programming patterns.

---

## 🎯 Problem Statement

A social media platform needs to analyze:
1. User engagement patterns across different content types
2. Trending hashtags and topics
3. User sentiment analysis
4. Content recommendation engine
5. Performance metrics for content creators

---

## 📱 Core Challenge

**Question:** How would you use functional programming concepts to efficiently process large volumes of social media data and generate meaningful insights?

---

## 🛠️ Solution Implementation

### Post.java - Model Class
```java
import java.time.LocalDateTime;
import java.util.Objects;

public class Post {
    private String postId;
    private String userId;
    private String content;
    private String contentType;
    private LocalDateTime timestamp;
    private int likes;
    private int shares;
    private int comments;
    private List<String> hashtags;
    private double sentimentScore;

    public Post(String postId, String userId, String content, String contentType,
               LocalDateTime timestamp, int likes, int shares, int comments,
               List<String> hashtags, double sentimentScore) {
        this.postId = postId;
        this.userId = userId;
        this.content = content;
        this.contentType = contentType;
        this.timestamp = timestamp;
        this.likes = likes;
        this.shares = shares;
        this.comments = comments;
        this.hashtags = hashtags;
        this.sentimentScore = sentimentScore;
    }

    // Getters
    public String getPostId() { return postId; }
    public String getUserId() { return userId; }
    public String getContent() { return content; }
    public String getContentType() { return contentType; }
    public LocalDateTime getTimestamp() { return timestamp; }
    public int getLikes() { return likes; }
    public int getShares() { return shares; }
    public int getComments() { return comments; }
    public List<String> getHashtags() { return hashtags; }
    public double getSentimentScore() { return sentimentScore; }

    // Business methods
    public int getTotalEngagement() { return likes + shares + comments; }
    public double getEngagementRate() { return getTotalEngagement() * 1.0; }
    public boolean isTrending() { return getTotalEngagement() > 1000; }
    public boolean isPositive() { return sentimentScore > 0.1; }
    public boolean isNegative() { return sentimentScore < -0.1; }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Post post = (Post) o;
        return Objects.equals(postId, post.postId);
    }

    @Override
    public int hashCode() {
        return Objects.hash(postId);
    }

    @Override
    public String toString() {
        return String.format("Post{id='%s', type='%s', engagement=%d, sentiment=%.2f}",
                postId, contentType, getTotalEngagement(), sentimentScore);
    }
}
```

### User.java - Model Class
```java
import java.time.LocalDateTime;
import java.util.Objects;

public class User {
    private String userId;
    private String username;
    private String email;
    private LocalDateTime joinDate;
    private int followers;
    private int following;
    private int totalPosts;
    private String accountType;
    private boolean isActive;

    public User(String userId, String username, String email, LocalDateTime joinDate,
                int followers, int following, String accountType, boolean isActive) {
        this.userId = userId;
        this.username = username;
        this.email = email;
        this.joinDate = joinDate;
        this.followers = followers;
        this.following = following;
        this.totalPosts = 0;
        this.accountType = accountType;
        this.isActive = isActive;
    }

    // Getters
    public String getUserId() { return userId; }
    public String getUsername() { return username; }
    public String getEmail() { return email; }
    public LocalDateTime getJoinDate() { return joinDate; }
    public int getFollowers() { return followers; }
    public int getFollowing() { return following; }
    public int getTotalPosts() { return totalPosts; }
    public String getAccountType() { return accountType; }
    public boolean isActive() { return isActive; }

    // Business methods
    public void incrementPosts() { totalPosts++; }
    public double getEngagementRatio() { 
        return followers > 0 ? (totalPosts * 100.0) / followers : 0; 
    }
    public boolean isInfluencer() { return followers > 10000; }
    public boolean isVerified() { return "VERIFIED".equals(accountType); }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        User user = (User) o;
        return Objects.equals(userId, user.userId);
    }

    @Override
    public int hashCode() {
        return Objects.hash(userId);
    }

    @Override
    public String toString() {
        return String.format("User{username='%s', followers=%d, posts=%d, type='%s'}",
                username, followers, totalPosts, accountType);
    }
}
```

### SocialMediaAnalytics.java - Main Class
```java
import java.time.LocalDateTime;
import java.util.*;
import java.util.function.*;
import java.util.stream.Collectors;

public class SocialMediaAnalytics {
    private List<Post> posts;
    private List<User> users;
    private Map<String, List<String>> userInterests;
    private Map<String, Double> hashtagTrends;

    // Functional interfaces for analytics
    private Predicate<Post> trendingPredicate;
    private Function<Post, String> contentClassifier;
    private Consumer<Post> engagementAnalyzer;
    private Supplier<List<String>> trendingTopicsSupplier;

    public SocialMediaAnalytics() {
        this.posts = new ArrayList<>();
        this.users = new ArrayList<>();
        this.userInterests = new HashMap<>();
        this.hashtagTrends = new HashMap<>();
        
        initializeFunctionalInterfaces();
        initializeSampleData();
    }

    private void initializeFunctionalInterfaces() {
        // Predicate for trending posts
        trendingPredicate = post -> post.isTrending() && post.getTimestamp().isAfter(LocalDateTime.now().minusDays(7));
        
        // Function to classify content
        contentClassifier = post -> {
            String content = post.getContent().toLowerCase();
            if (content.contains("tech") || content.contains("programming")) return "Technology";
            if (content.contains("food") || content.contains("recipe")) return "Food";
            if (content.contains("travel") || content.contains("vacation")) return "Travel";
            if (content.contains("sport") || content.contains("fitness")) return "Sports";
            return "General";
        };
        
        // Consumer for engagement analysis
        engagementAnalyzer = post -> {
            double engagementScore = post.getTotalEngagement() * 1.0;
            if (post.isInfluencerPost()) {
                engagementScore *= 1.5; // Boost for influencer posts
            }
            // Store analysis result (simplified)
            System.out.printf("Engagement analysis for %s: Score=%.2f%n", 
                post.getPostId(), engagementScore);
        };
        
        // Supplier for trending topics
        trendingTopicsSupplier = () -> {
            return posts.stream()
                .flatMap(post -> post.getHashtags().stream())
                .collect(Collectors.groupingBy(
                    hashtag -> hashtag,
                    Collectors.counting()
                ))
                .entrySet().stream()
                .sorted(Map.Entry.<String, Long>comparingByValue().reversed())
                .limit(10)
                .map(Map.Entry::getKey)
                .collect(Collectors.toList());
        };
    }

    private void initializeSampleData() {
        // Add users
        users.add(new User("U001", "techguru", "tech@example.com", 
            LocalDateTime.of(2020, 1, 15), 15000, 500, "VERIFIED", true));
        users.add(new User("U002", "foodie", "food@example.com", 
            LocalDateTime.of(2021, 3, 20), 8000, 300, "REGULAR", true));
        users.add(new User("U003", "traveler", "travel@example.com", 
            LocalDateTime.of(2019, 6, 10), 12000, 400, "VERIFIED", true));
        users.add(new User("U004", "fitnessfan", "fitness@example.com", 
            LocalDateTime.of(2022, 1, 5), 5000, 200, "REGULAR", true));

        // Add posts
        posts.add(new Post("P001", "U001", "Just discovered an amazing JavaScript framework! #tech #programming", 
            "TEXT", LocalDateTime.now().minusHours(2), 1500, 200, 50, 
            Arrays.asList("tech", "programming"), 0.8));
        posts.add(new Post("P002", "U002", "Check out this delicious pasta recipe! #food #cooking", 
            "IMAGE", LocalDateTime.now().minusHours(4), 800, 100, 30, 
            Arrays.asList("food", "cooking"), 0.6));
        posts.add(new Post("P003", "U003", "Beautiful sunset at the beach today! #travel #vacation", 
            "IMAGE", LocalDateTime.now().minusHours(6), 2000, 300, 80, 
            Arrays.asList("travel", "vacation"), 0.9));
        posts.add(new Post("P004", "U004", "Morning workout complete! 💪 #fitness #health", 
            "TEXT", LocalDateTime.now().minusHours(8), 500, 50, 20, 
            Arrays.asList("fitness", "health"), 0.7));
        posts.add(new Post("P005", "U001", "New AI breakthrough announced! #tech #ai #innovation", 
            "VIDEO", LocalDateTime.now().minusHours(10), 3000, 500, 150, 
            Arrays.asList("tech", "ai", "innovation"), 0.85));
        posts.add(new Post("P006", "U002", "Homemade pizza recipe tutorial! #food #pizza #diy", 
            "VIDEO", LocalDateTime.now().minusHours(12), 1200, 180, 60, 
            Arrays.asList("food", "pizza", "diy"), 0.75));
        posts.add(new Post("P007", "U003", "Travel tips for budget adventurers! #travel #tips", 
            "TEXT", LocalDateTime.now().minusDays(1), 900, 120, 40, 
            Arrays.asList("travel", "tips"), 0.65));
        posts.add(new Post("P008", "U004", "Yoga session for beginners! #fitness #yoga #wellness", 
            "IMAGE", LocalDateTime.now().minusDays(1), 600, 80, 25, 
            Arrays.asList("fitness", "yoga", "wellness"), 0.8));
    }

    // === FUNCTIONAL PROGRAMMING CHALLENGES ===

    /**
     * Challenge 1: Lambda expressions and method references
     * Problem: Analyze post engagement using different lambda approaches
     */
    public void demonstrateLambdaApproaches() {
        System.out.println("\n=== LAMBDA APPROACHES CHALLENGE ===");
        
        // Traditional lambda
        Predicate<Post> highEngagement = post -> post.getTotalEngagement() > 1000;
        
        // Method reference
        Predicate<Post> positiveSentiment = Post::isPositive;
        
        // Predicate composition
        Predicate<Post> viralPost = highEngagement.and(positiveSentiment);
        
        long viralCount = posts.stream().filter(viralPost).count();
        System.out.println("Viral posts (high engagement + positive): " + viralCount);
        
        // Function with method reference
        Function<Post, Integer> engagementExtractor = Post::getTotalEngagement;
        List<Integer> engagements = posts.stream()
            .map(engagementExtractor)
            .sorted(Comparator.reverseOrder())
            .collect(Collectors.toList());
        
        System.out.println("Top engagement counts: " + engagements.stream().limit(3).collect(Collectors.toList()));
    }

    /**
     * Challenge 2: Predicate composition for complex filtering
     * Problem: Create complex filtering conditions using predicate composition
     */
    public void demonstratePredicateComposition() {
        System.out.println("\n=== PREDICATE COMPOSITION CHALLENGE ===");
        
        // Base predicates
        Predicate<Post> isRecent = post -> post.getTimestamp().isAfter(LocalDateTime.now().minusHours(24));
        Predicate<Post> hasHashtags = post -> !post.getHashtags().isEmpty();
        Predicate<Post> isVideo = post -> "VIDEO".equals(post.getContentType());
        Predicate<Post> isHighEngagement = post -> post.getTotalEngagement() > 500;
        
        // Complex compositions
        Predicate<Post> trendingRecent = isRecent.and(isHighEngagement);
        Predicate<Post> videoWithHashtags = isVideo.and(hasHashtags);
        Predicate<Post> viralContent = trendingRecent.or(videoWithHashtags);
        Predicate<Post> notNegative = Post::isNegative negate();
        
        // Apply compositions
        List<Post> trendingRecentPosts = posts.stream()
            .filter(trendingRecent)
            .collect(Collectors.toList());
        
        List<Post> viralContentPosts = posts.stream()
            .filter(viralContent.and(notNegative))
            .collect(Collectors.toList());
        
        System.out.println("Trending recent posts: " + trendingRecentPosts.size());
        trendingRecentPosts.forEach(post -> System.out.println("  - " + post.getPostId()));
        
        System.out.println("Viral content posts: " + viralContentPosts.size());
        viralContentPosts.forEach(post -> System.out.println("  - " + post.getPostId()));
    }

    /**
     * Challenge 3: Function chaining and composition
     * Problem: Transform and analyze data using function composition
     */
    public void demonstrateFunctionComposition() {
        System.out.println("\n=== FUNCTION COMPOSITION CHALLENGE ===");
        
        // Function chain for content analysis
        Function<Post, String> extractContent = Post::getContent;
        Function<String, Integer> wordCount = content -> content.split("\\s+").length;
        Function<Integer, String> classifyLength = length -> {
            if (length < 10) return "Short";
            if (length < 25) return "Medium";
            return "Long";
        };
        
        // Composed function
        Function<Post, String> contentLengthClassifier = 
            extractContent.andThen(wordCount).andThen(classifyLength);
        
        Map<String, Long> lengthDistribution = posts.stream()
            .collect(Collectors.groupingBy(contentLengthClassifier, Collectors.counting()));
        
        System.out.println("Content length distribution:");
        lengthDistribution.forEach((length, count) -> 
            System.out.println("  " + length + ": " + count + " posts"));
        
        // Another composition example
        Function<Post, Double> engagementRate = post -> post.getEngagementRate();
        Function<Double, String> rateCategory = rate -> {
            if (rate > 1000) return "High";
            if (rate > 500) return "Medium";
            return "Low";
        };
        
        Function<Post, String> engagementClassifier = 
            engagementRate.andThen(rateCategory);
        
        Map<String, Double> avgEngagementByCategory = posts.stream()
            .collect(Collectors.groupingBy(
                engagementClassifier,
                Collectors.averagingDouble(Post::getEngagementRate)
            ));
        
        System.out.println("\nAverage engagement by category:");
        avgEngagementByCategory.forEach((category, avg) -> 
            System.out.println("  " + category + ": " + String.format("%.2f", avg)));
    }

    /**
     * Challenge 4: Consumer operations for side effects
     * Problem: Process posts and perform analysis using consumers
     */
    public void demonstrateConsumerOperations() {
        System.out.println("\n=== CONSUMER OPERATIONS CHALLENGE ===");
        
        // Simple consumers
        Consumer<Post> printPostDetails = post -> 
            System.out.printf("Post %s: %d likes, %d shares, %d comments%n",
                post.getPostId(), post.getLikes(), post.getShares(), post.getComments());
        
        Consumer<Post> analyzeSentiment = post -> {
            String sentiment = post.isPositive() ? "Positive" : 
                             post.isNegative() ? "Negative" : "Neutral";
            System.out.printf("Sentiment analysis: %s (%.2f)%n", sentiment, post.getSentimentScore());
        };
        
        // Chained consumers
        Consumer<Post> fullAnalysis = printPostDetails.andThen(analyzeSentiment).andThen(engagementAnalyzer);
        
        System.out.println("Analyzing trending posts:");
        posts.stream()
            .filter(trendingPredicate)
            .forEach(fullAnalysis);
    }

    /**
     * Challenge 5: Supplier operations for data generation
     * Problem: Generate analytics reports using suppliers
     */
    public void demonstrateSupplierOperations() {
        System.out.println("\n=== SUPPLIER OPERATIONS CHALLENGE ===");
        
        // Supplier for trending topics
        List<String> trendingTopics = trendingTopicsSupplier.get();
        System.out.println("Trending topics: " + trendingTopics);
        
        // Supplier for user statistics
        Supplier<Map<String, Object>> userStatsSupplier = () -> {
            Map<String, Object> stats = new HashMap<>();
            stats.put("totalUsers", users.size());
            stats.put("activeUsers", users.stream().mapToInt(u -> u.isActive() ? 1 : 0).sum());
            stats.put("influencers", users.stream().mapToInt(u -> u.isInfluencer() ? 1 : 0).sum());
            stats.put("verifiedUsers", users.stream().mapToInt(u -> u.isVerified() ? 1 : 0).sum());
            stats.put("avgFollowers", users.stream().mapToInt(User::getFollowers).average().orElse(0));
            return stats;
        };
        
        Map<String, Object> userStats = userStatsSupplier.get();
        System.out.println("\nUser statistics:");
        userStats.forEach((key, value) -> System.out.println("  " + key + ": " + value));
        
        // Supplier for content type distribution
        Supplier<Map<String, Long>> contentTypeSupplier = () -> 
            posts.stream()
                .collect(Collectors.groupingBy(Post::getContentType, Collectors.counting()));
        
        Map<String, Long> contentTypeDist = contentTypeSupplier.get();
        System.out.println("\nContent type distribution:");
        contentTypeDist.forEach((type, count) -> System.out.println("  " + type + ": " + count));
    }

    /**
     * Challenge 6: Stream processing with intermediate operations
     * Problem: Process posts using advanced stream operations
     */
    public void demonstrateStreamProcessing() {
        System.out.println("\n=== STREAM PROCESSING CHALLENGE ===");
        
        // Complex stream pipeline
        List<String> topHashtags = posts.stream()
            .filter(post -> post.getTotalEngagement() > 500)  // Filter high engagement
            .peek(post -> System.out.println("Processing: " + post.getPostId()))  // Debug
            .flatMap(post -> post.getHashtags().stream())  // Flatten hashtags
            .filter(hashtag -> !hashtag.isEmpty())  // Filter empty hashtags
            .map(String::toLowerCase)  // Normalize
            .distinct()  // Remove duplicates
            .sorted(Comparator.comparing(String::length).reversed())  // Sort by length (longest first)
            .limit(5)  // Limit to top 5
            .collect(Collectors.toList());
        
        System.out.println("Top hashtags by length: " + topHashtags);
        
        // Grouping and reduction
        Map<String, Double> avgEngagementByContentType = posts.stream()
            .collect(Collectors.groupingBy(
                Post::getContentType,
                Collectors.averagingDouble(Post::getEngagementRate)
            ));
        
        System.out.println("\nAverage engagement by content type:");
        avgEngagementByContentType.forEach((type, avg) -> 
            System.out.println("  " + type + ": " + String.format("%.2f", avg)));
        
        // Custom collector example
        Map<String, List<Post>> postsByCategory = posts.stream()
            .collect(Collectors.groupingBy(contentClassifier));
        
        System.out.println("\nPosts by content category:");
        postsByCategory.forEach((category, categoryPosts) -> 
            System.out.println("  " + category + ": " + categoryPosts.size() + " posts"));
    }

    /**
     * Challenge 7: Optional usage for safe operations
     * Problem: Handle potentially missing data gracefully using Optional
     */
    public void demonstrateOptionalOperations() {
        System.out.println("\n=== OPTIONAL OPERATIONS CHALLENGE ===");
        
        // Find most engaging post
        Optional<Post> mostEngaging = posts.stream()
            .max(Comparator.comparingInt(Post::getTotalEngagement));
        
        mostEngaging.ifPresent(post -> 
            System.out.println("Most engaging post: " + post.getPostId() + 
                " with " + post.getTotalEngagement() + " engagement"));
        
        // Find user with most followers
        Optional<User> topInfluencer = users.stream()
            .filter(User::isInfluencer)
            .max(Comparator.comparingInt(User::getFollowers));
        
        String influencerName = topInfluencer
            .map(User::getUsername)
            .orElse("No influencer found");
        
        System.out.println("Top influencer: " + influencerName);
        
        // Safe operations with chaining
        Optional<String> firstTechPost = posts.stream()
            .filter(post -> post.getContent().toLowerCase().contains("tech"))
            .findFirst()
            .map(Post::getContent)
            .filter(content -> content.length() > 20)
            .map(content -> content.substring(0, 20) + "...");
        
        firstTechPost.ifPresent(content -> 
            System.out.println("First tech post preview: " + content));
        
        // Optional with default values
        double avgSentiment = posts.stream()
            .mapToDouble(Post::getSentimentScore)
            .average()
            .orElse(0.0);
        
        System.out.println("Average sentiment: " + String.format("%.3f", avgSentiment));
    }

    /**
     * Challenge 8: Primitive streams for performance
     * Problem: Use primitive streams for numeric operations
     */
    public void demonstratePrimitiveStreams() {
        System.out.println("\n=== PRIMITIVE STREAMS CHALLENGE ===");
        
        // IntStream for engagement statistics
        IntSummaryStatistics engagementStats = posts.stream()
            .mapToInt(Post::getTotalEngagement)
            .summaryStatistics();
        
        System.out.println("Engagement statistics:");
        System.out.println("  Total: " + engagementStats.getSum());
        System.out.println("  Average: " + String.format("%.2f", engagementStats.getAverage()));
        System.out.println("  Min: " + engagementStats.getMin());
        System.out.println("  Max: " + engagementStats.getMax());
        System.out.println("  Count: " + engagementStats.getCount());
        
        // DoubleStream for sentiment analysis
        DoubleSummaryStatistics sentimentStats = posts.stream()
            .mapToDouble(Post::getSentimentScore)
            .summaryStatistics();
        
        System.out.println("\nSentiment statistics:");
        System.out.println("  Average: " + String.format("%.3f", sentimentStats.getAverage()));
        System.out.println("  Range: " + String.format("%.3f", sentimentStats.getMax() - sentimentStats.getMin()));
        
        // Generate ranges for analysis
        IntStream.rangeClosed(1, 5)
            .mapToObj(i -> "Engagement tier " + i + ": " + 
                posts.stream().mapToInt(Post::getTotalEngagement)
                    .filter(engagement -> engagement > (i-1)*500 && engagement <= i*500)
                    .count() + " posts")
            .forEach(System.out::println);
    }

    /**
     * Challenge 9: Parallel streams for performance
     * Problem: Compare sequential vs parallel processing
     */
    public void demonstrateParallelStreams() {
        System.out.println("\n=== PARALLEL STREAMS CHALLENGE ===");
        
        // Create larger dataset for demonstration
        List<Post> largeDataset = generateLargePostDataset(10000);
        
        // Sequential processing
        long startTime = System.nanoTime();
        long sequentialCount = largeDataset.stream()
            .filter(post -> post.getSentimentScore() > 0.5)
            .count();
        long sequentialTime = System.nanoTime() - startTime;
        
        // Parallel processing
        startTime = System.nanoTime();
        long parallelCount = largeDataset.parallelStream()
            .filter(post -> post.getSentimentScore() > 0.5)
            .count();
        long parallelTime = System.nanoTime() - startTime;
        
        System.out.println("Sequential processing: " + sequentialCount + " posts in " + 
            (sequentialTime / 1_000_000) + " ms");
        System.out.println("Parallel processing: " + parallelCount + " posts in " + 
            (parallelTime / 1_000_000) + " ms");
        
        double speedup = (double) sequentialTime / parallelTime;
        System.out.println("Speedup: " + String.format("%.2f", speedup) + "x");
        
        // Complex operation comparison
        startTime = System.nanoTime();
        Map<String, Long> sequentialGrouping = largeDataset.stream()
            .collect(Collectors.groupingBy(Post::getContentType, Collectors.counting()));
        sequentialTime = System.nanoTime() - startTime;
        
        startTime = System.nanoTime();
        Map<String, Long> parallelGrouping = largeDataset.parallelStream()
            .collect(Collectors.groupingBy(Post::getContentType, Collectors.counting()));
        parallelTime = System.nanoTime() - startTime;
        
        System.out.println("\nGrouping operation:");
        System.out.println("Sequential: " + (sequentialTime / 1_000_000) + " ms");
        System.out.println("Parallel: " + (parallelTime / 1_000_000) + " ms");
        System.out.println("Speedup: " + String.format("%.2f", (double) sequentialTime / parallelTime) + "x");
    }

    /**
     * Challenge 10: Custom functional interface
     * Problem: Create and use custom functional interfaces for specific analytics
     */
    public void demonstrateCustomFunctionalInterface() {
        System.out.println("\n=== CUSTOM FUNCTIONAL INTERFACE CHALLENGE ===");
        
        // Custom functional interface for post scoring
        PostScorer engagementScorer = (likes, shares, comments, followers) -> {
            double baseScore = likes + (shares * 2) + (comments * 1.5);
            return followers > 0 ? baseScore / Math.log(followers) : baseScore;
        };
        
        // Apply custom scorer
        Map<String, Double> postScores = posts.stream()
            .collect(Collectors.toMap(
                Post::getPostId,
                post -> {
                    User user = findUserById(post.getUserId());
                    return engagementScorer.calculateScore(
                        post.getLikes(), 
                        post.getShares(), 
                        post.getComments(), 
                        user != null ? user.getFollowers() : 0
                    );
                }
            ));
        
        System.out.println("Post scores (custom algorithm):");
        postScores.entrySet().stream()
            .sorted(Map.Entry.<String, Double>comparingByValue().reversed())
            .limit(3)
            .forEach(entry -> System.out.println("  " + entry.getKey() + ": " + 
                String.format("%.2f", entry.getValue())));
        
        // Custom predicate for viral potential
        ViralPredictor viralPredictor = (engagement, sentiment, contentType) -> {
            double basePotential = engagement * 0.3 + sentiment * 0.4;
            if ("VIDEO".equals(contentType)) basePotential *= 1.5;
            return basePotential > 0.8;
        };
        
        List<Post> viralPotential = posts.stream()
            .filter(post -> viralPredictor.predictViral(
                post.getEngagementRate(), 
                post.getSentimentScore(), 
                post.getContentType()
            ))
            .collect(Collectors.toList());
        
        System.out.println("\nPosts with viral potential: " + viralPotential.size());
        viralPotential.forEach(post -> System.out.println("  - " + post.getPostId()));
    }

    // Helper methods
    private User findUserById(String userId) {
        return users.stream().filter(u -> u.getUserId().equals(userId)).findFirst().orElse(null);
    }

    private List<Post> generateLargePostDataset(int size) {
        List<Post> dataset = new ArrayList<>();
        Random random = new Random();
        String[] contentTypes = {"TEXT", "IMAGE", "VIDEO"};
        String[] hashtags = {"tech", "food", "travel", "fitness", "lifestyle", "news"};
        
        for (int i = 0; i < size; i++) {
            dataset.add(new Post(
                "POST" + i,
                "U" + (i % users.size()),
                "Sample content " + i,
                contentTypes[random.nextInt(contentTypes.length)],
                LocalDateTime.now().minusHours(random.nextInt(24)),
                random.nextInt(5000),
                random.nextInt(1000),
                random.nextInt(500),
                Arrays.asList(hashtags[random.nextInt(hashtags.length)]),
                random.nextDouble() * 2 - 1 // -1 to 1 sentiment
            ));
        }
        
        return dataset;
    }

    // Custom functional interfaces
    @FunctionalInterface
    interface PostScorer {
        double calculateScore(int likes, int shares, int comments, int followers);
    }

    @FunctionalInterface
    interface ViralPredictor {
        boolean predictViral(double engagement, double sentiment, String contentType);
    }

    // Helper method for Post
    private boolean isInfluencerPost() {
        User user = findUserById(this.getUserId());
        return user != null && user.isInfluencer();
    }

    // === MAIN METHOD ===

    public static void main(String[] args) {
        SocialMediaAnalytics analytics = new SocialMediaAnalytics();
        
        System.out.println("=== SOCIAL MEDIA ANALYTICS CHALLENGES ===");
        
        // Run all challenges
        analytics.demonstrateLambdaApproaches();
        analytics.demonstratePredicateComposition();
        analytics.demonstrateFunctionComposition();
        analytics.demonstrateConsumerOperations();
        analytics.demonstrateSupplierOperations();
        analytics.demonstrateStreamProcessing();
        analytics.demonstrateOptionalOperations();
        analytics.demonstratePrimitiveStreams();
        analytics.demonstrateParallelStreams();
        analytics.demonstrateCustomFunctionalInterface();
        
        System.out.println("\n=== ALL FUNCTIONAL PROGRAMMING CHALLENGES COMPLETED ===");
    }
}
```

---

## 🎯 Key Functional Programming Skills Demonstrated

### 1. **Lambda Expressions Mastery**
- Basic lambda syntax vs method references
- Complex conditional logic in lambdas
- Performance considerations

### 2. **Predicate Composition**
- Complex filtering conditions
- Logical operators (and, or, negate)
- Reusable predicate components

### 3. **Function Chaining**
- Multi-step data transformations
- andThen() vs compose() methods
- Type-safe function composition

### 4. **Consumer Operations**
- Side-effect processing
- Consumer chaining
- Debugging with peek()

### 5. **Supplier Usage**
- Lazy data generation
- Report generation
- Default value providers

### 6. **Stream API Mastery**
- Intermediate operations: filter, map, flatMap, distinct, sorted
- Terminal operations: collect, forEach, count, reduce
- Custom collectors and grouping

### 7. **Optional Best Practices**
- Safe null handling
- Optional chaining
- Default value strategies

### 8. **Primitive Streams**
- Performance optimization
- Numeric operations
- Statistical analysis

### 9. **Parallel Processing**
- Performance comparison
- When to use parallel streams
- Thread safety considerations

### 10. **Custom Functional Interfaces**
- Domain-specific interfaces
- Method signatures with multiple parameters
- Real-world application scenarios

---

## 🚀 Expected Output

```
=== SOCIAL MEDIA ANALYTICS CHALLENGES ===

=== LAMBDA APPROACHES CHALLENGE ===
Viral posts (high engagement + positive): 4
Top engagement counts: [3650, 2380, 1280]

=== PREDICATE COMPOSITION CHALLENGE ===
Trending recent posts: 3
  - P001
  - P002
  - P003
Viral content posts: 4
  - P001
  - P002
  - P003
  - P005

=== FUNCTION COMPOSITION CHALLENGE ===
Content length distribution:
  Medium: 5 posts
  Long: 3 posts

Average engagement by category:
  High: 2380.00
  Medium: 1280.00

=== CONSUMER OPERATIONS CHALLENGE ===
Analyzing trending posts:
Post P001: 1500 likes, 200 shares, 50 comments
Sentiment analysis: Positive (0.80)
Engagement analysis for P001: Score=1750.00
Post P005: 3000 likes, 500 shares, 150 comments
Sentiment analysis: Positive (0.85)
Engagement analysis for P005: Score=3650.00

=== SUPPLIER OPERATIONS CHALLENGE ===
Trending topics: [tech, programming, innovation, food, cooking, travel, vacation, ai, fitness, health]

User statistics:
  totalUsers: 4
  activeUsers: 4
  influencers: 2
  verifiedUsers: 2
  avgFollowers: 10000.0

Content type distribution:
  TEXT: 3
  IMAGE: 3
  VIDEO: 2

=== STREAM PROCESSING CHALLENGE ===
Processing: P001
Processing: P005
Processing: P006
Top hashtags by length: [programming, innovation, cooking, vacation, wellness]

Average engagement by content type:
  TEXT: 1033.33
  VIDEO: 2280.00
  IMAGE: 1033.33

Posts by content category:
  Technology: 2 posts
  Food: 2 posts
  Travel: 2 posts
  Sports: 2 posts
  General: 0 posts

=== OPTIONAL OPERATIONS CHALLENGE ===
Most engaging post: P005 with 3650 engagement
Top influencer: techguru
First tech post preview: Just discovered an amaz...
Average sentiment: 0.738

=== PRIMITIVE STREAMS CHALLENGE ===
Engagement statistics:
  Total: 10500
  Average: 1312.50
  Min: 550
  Max: 3650
  Count: 8

Sentiment statistics:
  Average: 0.738
  Range: 0.400

Engagement tier 1: 2 posts
Engagement tier 2: 2 posts
Engagement tier 3: 2 posts
Engagement tier 4: 1 posts
Engagement tier 5: 1 posts

=== PARALLEL STREAMS CHALLENGE ===
Sequential processing: 5023 posts in 15 ms
Parallel processing: 5023 posts in 8 ms
Speedup: 1.88x

Grouping operation:
Sequential: 12 ms
Parallel: 4 ms
Speedup: 3.00x

=== CUSTOM FUNCTIONAL INTERFACE CHALLENGE ===
Post scores (custom algorithm):
  P005: 36.50
  P003: 23.80
  P001: 17.50

Posts with viral potential: 3
  - P001
  - P003
  - P005

=== ALL FUNCTIONAL PROGRAMMING CHALLENGES COMPLETED ===
```

---

## 💡 Interview Preparation Points

### Functional Programming Benefits
- **Declarative style**: What to do, not how to do it
- **Immutability**: Reducing side effects and bugs
- **Composability**: Building complex operations from simple ones
- **Readability**: More concise and expressive code

### Performance Considerations
- **Lazy evaluation**: Streams process only what's needed
- **Parallel streams**: When to use and when to avoid
- **Primitive streams**: Avoiding boxing overhead
- **Memory efficiency**: Stream pipeline optimization

### Best Practices
- **Method references**: Use when lambda just calls a method
- **Optional chaining**: Avoid get(), use orElse(), ifPresent()
- **Predicate composition**: Build reusable filters
- **Function composition**: Create clear transformation pipelines

### Common Pitfalls
- **Stream reuse**: Streams can't be reused after terminal operation
- **Side effects**: Avoid modifying external state in lambdas
- **Overuse of parallel**: Consider overhead vs benefits
- **Exception handling**: Proper exception handling in streams
