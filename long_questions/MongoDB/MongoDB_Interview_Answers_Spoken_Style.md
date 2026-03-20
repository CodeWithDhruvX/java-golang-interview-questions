# MongoDB Interview Questions & Answers (Spoken Style Format)

## Service-Based Companies (Fresher to Mid-Level)

### 1. What is MongoDB and how does it differ from traditional SQL databases?

**How to Explain in Interview (Spoken style format):**

"MongoDB is a NoSQL, document-oriented database that stores data in flexible, JSON-like documents called BSON. Unlike traditional SQL databases that use rigid tables with rows and columns, MongoDB uses collections and documents.

The key differences are:
- MongoDB has a flexible schema - documents in the same collection can have different structures
- It uses JSON/BSON format instead of tables
- No complex JOINs - we embed related data or use references
- Horizontal scaling is easier with sharding
- It's more suitable for unstructured data and rapid development

For example, in a user collection, one user document might have an 'age' field as a number, while another might have it as a string - MongoDB allows this flexibility, whereas SQL would require strict data types."

### 2. What are collections and documents in MongoDB?

**How to Explain in Interview (Spoken style format):**

"In MongoDB, a **document** is the basic unit of data, similar to a row in SQL. It's a JSON-like object with key-value pairs. For example, `{ name: 'John', age: 30, email: 'john@example.com' }` is a document.

A **collection** is a group of documents, similar to a table in SQL. But unlike SQL tables, collections don't enforce a strict schema - different documents in the same collection can have different fields.

So if I have a users collection, one document might have fields for name, email, and age, while another document might have name, email, and phone number. MongoDB handles this flexibility naturally."

### 3. What is the `_id` field in MongoDB?

**How to Explain in Interview (Spoken style format):**

"Every document in MongoDB must have an `_id` field, which acts as the primary key. It must be unique within a collection.

If I don't provide an `_id` when inserting a document, MongoDB automatically generates a 12-byte ObjectId. This ObjectId is actually composed of:
- A 4-byte timestamp
- A 5-byte random value (machine identifier)
- A 3-byte incrementing counter

This structure ensures uniqueness across distributed systems and also gives us the creation time of the document. I can also provide my own `_id` if needed, like using a UUID or a business-specific identifier."

### 4. How do you perform basic CRUD operations in MongoDB?

**How to Explain in Interview (Spoken style format):**

"CRUD operations in MongoDB are straightforward:

**Create** - I use `insertOne()` or `insertMany()`:
```javascript
db.users.insertOne({ name: 'Alice', age: 25 })
```

**Read** - I use `find()` or `findOne()`:
```javascript
// Find all users
db.users.find()

// Find users older than 20
db.users.find({ age: { $gt: 20 } })

// Find one specific user
db.users.findOne({ email: 'alice@example.com' })
```

**Update** - I use `updateOne()` or `updateMany()` with operators like `$set`:
```javascript
db.users.updateOne({ name: 'Alice' }, { $set: { age: 26 } })
```

**Delete** - I use `deleteOne()` or `deleteMany()`:
```javascript
db.users.deleteOne({ name: 'Alice' })
```

The key is using query operators like `$gt`, `$in`, `$or` for complex conditions and update operators like `$set`, `$push`, `$pull` for modifications."

### 5. What is Mongoose and why do we use it?

**How to Explain in Interview (Spoken style format):**

"Mongoose is an Object Data Modeling (ODM) library for MongoDB and Node.js. While MongoDB is schema-less, Mongoose allows us to define schemas at the application level.

I use Mongoose because it provides:
- **Schema validation** - I can define required fields, data types, and validation rules
- **Middleware/hooks** - I can run code before or after operations, like hashing passwords before saving
- **Easier queries** - It provides a cleaner syntax for complex operations
- **Type safety** - Better IntelliSense and error catching in development

For example, I can define a user schema:
```javascript
const userSchema = new mongoose.Schema({
  name: { type: String, required: true },
  email: { type: String, unique: true },
  age: { type: Number, min: 18 }
});
```

This ensures data integrity before it even reaches the database."

### 6. What is the difference between Schema and Model in Mongoose?

**How to Explain in Interview (Spoken style format):**

"In Mongoose, a **Schema** is the blueprint that defines the structure of documents - it specifies fields, data types, validation rules, and defaults.

A **Model** is the compiled version of the Schema that provides the interface to interact with the database collection.

Think of it like this: Schema is the architectural drawing, Model is the actual building.

I first create a schema:
```javascript
const userSchema = new mongoose.Schema({
  name: String,
  email: String
});
```

Then I compile it into a model:
```javascript
const User = mongoose.model('User', userSchema);
```

Now I can use the User model to perform operations like `User.find()`, `User.create()`, etc. The model has all the methods I need to work with the users collection."

### 7. How do you handle relationships in MongoDB?

**How to Explain in Interview (Spoken style format):**

"MongoDB doesn't have foreign keys like SQL, but I can handle relationships in two ways:

**Embedding** - I nest related documents inside a parent document. This is good for one-to-many relationships where the related data is always needed together.

For example, embedding comments in a blog post:
```javascript
{
  title: "My Post",
  comments: [
    { text: "Great post!", author: "John" },
    { text: "Thanks for sharing", author: "Jane" }
  ]
}
```

**Referencing** - I store the `_id` of related documents and use Mongoose's `populate()` to fetch them when needed.

For example, posts referencing users:
```javascript
const postSchema = new mongoose.Schema({
  title: String,
  author: { type: mongoose.Schema.Types.ObjectId, ref: 'User' }
});

// Later fetch post with author details
Post.find().populate('author').exec();
```

I choose embedding for read-heavy scenarios and referencing when data is accessed independently."

---

## Product-Based Companies (Mid to Senior Level)

### 8. Explain Sharding in MongoDB and how to choose a Shard Key

**How to Explain in Interview (Spoken style format):**

"Sharding is MongoDB's method for horizontal scaling - it distributes data across multiple servers (shards) to handle large datasets and high throughput.

The most critical decision is choosing the **shard key** - this determines how data is distributed. A good shard key should have:

**High Cardinality** - Many unique values to distribute data evenly. If I use a field with only a few values like 'status' (active/inactive), data will cluster on few shards.

**Non-Monotonically Increasing** - If I use timestamps or sequential IDs, all new writes go to the same shard (hot spot). I might use a hashed shard key to distribute these evenly.

**Query Isolation** - Ideally, my queries should include the shard key so MongoDB can target specific shards instead of querying all shards (scatter-gather).

For example, in an e-commerce system, I might use `userId` as shard key for user data because it has high cardinality and queries usually filter by user. For orders, I might use a compound key like `{ userId: 1, orderDate: 1 }` or a hashed key on `orderId`."

### 9. What is a Replica Set and how does failover work?

**How to Explain in Interview (Spoken style format):**

"A Replica Set is a group of MongoDB servers that maintain the same data copy for high availability and redundancy.

It typically consists of:
- **Primary node** - Handles all write operations and replicates changes to secondaries
- **Secondary nodes** - Replicate the primary's operations and can serve read queries
- **Arbiter** (optional) - Votes in elections but doesn't store data

During failover:
1. Secondaries send heartbeats to monitor the primary
2. If the primary is unreachable for more than 10 seconds, an election is triggered
3. Eligible secondaries vote based on priority and who has the most up-to-date data
4. The node with majority votes becomes the new primary
5. Applications automatically reconnect to the new primary

This ensures the database remains available even if the primary server fails. The election process usually takes 10-20 seconds, during which the cluster is unavailable for writes but can still serve reads from secondaries."

### 10. Explain different types of indexes in MongoDB and when to use them

**How to Explain in Interview (Spoken style format):**

"Indexes in MongoDB are special data structures that improve query performance. The main types are:

**Single Field Index** - Index on one field. Good for queries that filter or sort on a single field.
```javascript
db.users.createIndex({ email: 1 })
```

**Compound Index** - Index on multiple fields. The order is crucial - it follows the ESR rule: Equality fields first, then Sort fields, then Range fields.
```javascript
db.users.createIndex({ status: 1, createdAt: -1, age: 1 })
```

**Multikey Index** - Index on array fields. Automatically created when indexing an array field.

**Text Index** - For full-text search on string content.
```javascript
db.articles.createIndex({ content: "text" })
```

**Geospatial Index** - For location-based queries using `$near` or `$geoWithin`.

**Partial/Sparse Indexes** - Only index documents that meet certain criteria, saving memory.
```javascript
db.users.createIndex({ email: 1 }, { partialFilterExpression: { email: { $exists: true } } })
```

I use compound indexes for common query patterns, text indexes for search functionality, and partial indexes when I only need to index a subset of documents."

### 11. Does MongoDB support ACID transactions? When should you use them?

**How to Explain in Interview (Spoken style format):**

"Yes, MongoDB supports ACID transactions since version 4.0 for replica sets and 4.2 for sharded clusters. However, I use them sparingly because they come with performance costs.

Transactions in MongoDB work across multiple documents and collections, ensuring all operations succeed or fail together. I implement them using sessions:

```javascript
const session = await mongoose.startSession();
session.startTransaction();

try {
  await User.updateOne({ _id: userId }, { $inc: { balance: -100 } }, { session });
  await Account.updateOne({ _id: accountId }, { $inc: { balance: 100 } }, { session });
  await session.commitTransaction();
} catch (error) {
  await session.abortTransaction();
} finally {
  session.endSession();
}
```

I only use transactions when absolutely necessary, like:
- Financial operations (money transfers)
- Multi-document consistency requirements
- Complex business logic that can't be modeled with embedded documents

For most cases, I prefer embedding related data in a single document or using eventual consistency patterns, as transactions limit horizontal scalability and add locking overhead."

### 12. What is the Aggregation Framework and how do you optimize it?

**How to Explain in Interview (Spoken style format):**

"The Aggregation Framework is MongoDB's powerful pipeline for data processing and transformation. It processes documents through multiple stages to produce computed results.

Common stages include:
- `$match` - Filter documents (like WHERE clause)
- `$group` - Group documents and perform calculations (like GROUP BY)
- `$sort` - Sort documents
- `$project` - Reshape documents (select/exclude fields)
- `$lookup` - Join collections (like LEFT JOIN)
- `$unwind` - Deconstruct array fields

For example, to get monthly sales:
```javascript
db.orders.aggregate([
  { $match: { status: "completed" } },
  { $group: { _id: { $month: "$createdAt" }, total: { $sum: "$amount" } } },
  { $sort: { _id: 1 } }
])
```

To optimize aggregations:
1. **Place `$match` early** - Reduce documents processed in later stages
2. **Use indexes** - Ensure `$match` stages use indexed fields
3. **Limit with `$limit`** - Reduce dataset size early
4. **Avoid `$lookup` when possible** - Consider embedding instead
5. **Use `$project`** - Remove unnecessary fields early to reduce memory usage

I also use `explain()` to analyze aggregation performance and identify bottlenecks."

### 13. How do you implement pagination in MongoDB?

**How to Explain in Interview (Spoken style format):**

"I implement pagination using `skip()` and `limit()`, but I also consider performance implications:

**Basic pagination:**
```javascript
// Page 1 (10 items per page)
db.users.find().limit(10)

// Page 2
db.users.find().skip(10).limit(10)

// Page 3
db.users.find().skip(20).limit(10)
```

However, `skip()` becomes slow for large offsets because MongoDB still has to scan through all the skipped documents.

For better performance, I use **range-based pagination**:
```javascript
// First page
db.users.find().sort({ _id: 1 }).limit(10)

// Next page using last seen _id
db.users.find({ _id: { $gt: lastId } }).sort({ _id: 1 }).limit(10)
```

This is much faster because it uses the index to start directly from the last seen document.

I also implement **cursor-based pagination** for real-time applications:
```javascript
db.users.find({ _id: { $gt: cursor } }).sort({ _id: 1 }).limit(10)
```

And I always include total count for UI:
```javascript
const total = await users.countDocuments(query);
const users = await users.find(query).skip(offset).limit(limit);
```

This gives the frontend both the data and pagination metadata."

### 14. What is the difference between `$lookup` and embedding?

**How to Explain in Interview (Spoken style format):**

"`$lookup` and embedding are two different approaches to handle relationships in MongoDB.

**Embedding** means storing related data within the same document:
```javascript
{
  _id: 1,
  title: "Blog Post",
  author: { name: "John", email: "john@example.com" },
  comments: [
    { text: "Great post!", author: "Jane" }
  ]
}
```

**`$lookup`** is like a SQL JOIN - it fetches related documents from other collections:
```javascript
db.posts.aggregate([
  { $lookup: {
    from: "users",
    localField: "authorId",
    foreignField: "_id",
    as: "author"
  }}
])
```

I choose embedding when:
- Data is always accessed together
- The embedded data is relatively small
- Read performance is critical
- Data doesn't change independently

I choose `$lookup` when:
- Data is accessed independently
- Related data is large or changes frequently
- I need to avoid data duplication
- Different parts of data have different access patterns

For example, I embed user profile data in posts for fast display, but use `$lookup` for separate user management operations. The key is considering the read-write patterns and data consistency requirements."

### 15. How do you monitor and optimize MongoDB performance?

**How to Explain in Interview (Spoken style format):**

"I monitor and optimize MongoDB performance through several approaches:

**Monitoring Tools:**
- **MongoDB Atlas** or **Cloud Manager** for cloud deployments
- **MongoDB Compass** for visual query analysis
- **`db.stats()`** and **`collection.stats()`** for database metrics
- **`explain()`** to analyze query execution plans

**Key Metrics I Track:**
- Query execution time and slow queries
- Index usage and efficiency
- Memory usage and page faults
- Connection count and queue depth
- Disk I/O and storage usage

**Optimization Techniques:**
1. **Index Optimization** - Create appropriate indexes based on query patterns
2. **Query Optimization** - Use `explain()` to identify full collection scans
3. **Connection Pooling** - Reuse connections instead of creating new ones
4. **Read Preference** - Route reads to secondary nodes when appropriate
5. **Write Concern** - Balance between durability and performance
6. **Schema Design** - Use embedding vs referencing appropriately

**Example Optimization Process:**
```javascript
// First, analyze slow query
db.users.find({ age: { $gt: 25 } }).sort({ name: 1 }).explain("executionStats")

// Create compound index based on analysis
db.users.createIndex({ age: 1, name: 1 })

// Verify improvement
db.users.find({ age: { $gt: 25 } }).sort({ name: 1 }).explain("executionStats")
```

I also set up alerts for metrics like slow query count, connection usage, and disk space to proactively identify performance issues."

---

## Advanced Topics (Senior/Lead Level)

### 16. How do you design a MongoDB schema for a high-traffic e-commerce application?

**How to Explain in Interview (Spoken style format):**

"For a high-traffic e-commerce application, I'd design the schema considering read/write patterns, scalability, and performance:

**Products Collection:**
```javascript
{
  _id: ObjectId,
  name: String,
  price: Number,
  category: String,
  inventory: Number,
  variants: [{ // Embedded for fast access
    size: String,
    color: String,
    stock: Number
  }],
  reviews: [{ // Recent reviews embedded
    rating: Number,
    comment: String,
    userId: ObjectId,
    date: Date
  }],
  // Indexes: category, price, inventory
}
```

**Orders Collection (Sharded by userId):**
```javascript
{
  _id: ObjectId,
  userId: ObjectId,
  items: [{
    productId: ObjectId,
    quantity: Number,
    price: Number
  }],
  total: Number,
  status: String,
  createdAt: Date,
  // Shard key: userId for even distribution
}
```

**Users Collection:**
```javascript
{
  _id: ObjectId,
  email: String,
  profile: {
    name: String,
    addresses: [{ // Embedded addresses
      type: String, // shipping/billing
      street: String,
      city: String
    }]
  },
  // Separate collections for orders, cart to avoid document size limits
}
```

**Design Decisions:**
- **Embed frequently accessed data** like product variants and recent reviews
- **Reference large/infrequent data** like full order history
- **Shard orders by userId** for even distribution and query isolation
- **Use compound indexes** on category+price for product searches
- **Implement TTL indexes** for session data and temporary carts

This design balances read performance with write scalability while keeping document sizes manageable."

### 17. How do you handle data consistency in a distributed MongoDB deployment?

**How to Explain in Interview (Spoken style format):**

"In distributed MongoDB deployments, I handle data consistency through several strategies:

**Write Concern Levels:**
- **w:1** - Acknowledged by primary (fastest, potential data loss)
- **w:majority** - Acknowledged by majority of replicas (default for transactions)
- **w:all** - Acknowledged by all replicas (slowest, strongest consistency)

I choose based on criticality - financial operations use `w:majority`, while analytics might use `w:1`.

**Read Concern Levels:**
- **local** - Read from primary without waiting for replication
- **majority** - Read data acknowledged by majority (strong consistency)
- **linearizable** - Strongest consistency, reads latest committed data

**Causal Consistency:**
```javascript
const session = await mongoose.startSession();
session.startTransaction();

// All operations in this session are causally consistent
await User.updateOne({ _id: userId }, { $inc: { balance: -100 } }, { session });
await Order.create({ userId, amount: 100 }, { session });

await session.commitTransaction();
```

**Eventual Consistency Patterns:**
- Use **Change Streams** to sync data across services
- Implement **Saga Pattern** for distributed transactions
- Use **Compensating Transactions** for rollback scenarios

**Example: Order Processing**
```javascript
// 1. Create order with w:majority
const order = await Order.create(orderData, { w: 'majority' });

// 2. Update inventory (can be eventual)
await Inventory.updateOne({ productId }, { $inc: { stock: -quantity } });

// 3. If inventory fails, compensate
if (inventoryError) {
  await Order.updateOne({ _id: order._id }, { status: 'failed' });
  // Refund payment
}
```

I also implement **idempotent operations** and **retry logic** to handle network partitions and temporary failures."

### 18. How do you implement a multi-tenant architecture in MongoDB?

**How to Explain in Interview (Spoken style format):**

"For multi-tenant architecture in MongoDB, I consider three main approaches:

**1. Database-per-Tenant:**
Each tenant gets their own database. This provides complete isolation but can be expensive and hard to manage at scale.

**2. Collection-per-Tenant:**
Each tenant gets separate collections within one database. Better resource utilization but still complex to manage.

**3. Shared Collection with Tenant ID (Most Common):**
All tenants share collections with a `tenantId` field:

```javascript
// User schema with tenant isolation
const userSchema = new mongoose.Schema({
  tenantId: { type: String, required: true, index: true },
  name: String,
  email: String,
  // Other fields...
});

// Compound index for tenant-specific queries
userSchema.index({ tenantId: 1, email: 1 }, { unique: true });
```

**Implementation Strategies:**

**Middleware for Automatic Tenant Isolation:**
```javascript
// Global middleware to filter by tenant
app.use((req, res, next) => {
  req.tenantId = getTenantFromRequest(req);
  next();
});

// Mongoose middleware to auto-filter queries
userSchema.pre(/^find/, function() {
  this.where({ tenantId: this.getOptions().tenantId });
});
```

**Connection Pooling:**
```javascript
// Single connection pool with tenant-aware queries
mongoose.connect(process.env.MONGO_URI);

// All queries automatically include tenant filter
const users = await User.find().setOptions({ tenantId: 'tenant123' });
```

**Security Considerations:**
- Always validate tenant ownership
- Use row-level security patterns
- Implement proper authentication/authorization
- Audit logs for cross-tenant access attempts

**Scaling Considerations:**
- Shard by `{ tenantId: 1, _id: 1 }` for even distribution
- Use separate databases for large enterprise tenants
- Implement resource quotas per tenant

This approach provides good isolation while maximizing resource utilization and simplifying maintenance."

---

## Common Coding Questions

### 19. Write a MongoDB query to find duplicate documents

**How to Explain in Interview (Spoken style format):**

"To find duplicate documents, I use the aggregation framework with `$group` to count occurrences and `$match` to filter duplicates:

```javascript
// Find duplicate emails in users collection
db.users.aggregate([
  {
    $group: {
      _id: "$email",  // Group by email field
      count: { $sum: 1 },  // Count occurrences
      docs: { $push: "$_id" }  // Collect document IDs
    }
  },
  {
    $match: {
      count: { $gt: 1 }  // Only groups with more than 1 document
    }
  }
])
```

If I need to actually remove duplicates, I can:
```javascript
// Keep the first occurrence, remove others
db.users.aggregate([
  { $sort: { _id: 1 } },  // Sort to ensure consistent first document
  {
    $group: {
      _id: "$email",
      docs: { $push: "$_id" },
      count: { $sum: 1 }
    }
  },
  { $match: { count: { $gt: 1 } } }
]).forEach(function(doc) {
  // Remove all except the first one
  doc.docs.shift();  // Remove first element (keep it)
  db.users.deleteMany({ _id: { $in: doc.docs } });
});
```

For complex duplicate detection based on multiple fields:
```javascript
db.users.aggregate([
  {
    $group: {
      _id: { name: "$name", email: "$email" },  // Multiple fields
      count: { $sum: 1 },
      duplicates: { $push: "$_id" }
    }
  },
  { $match: { count: { $gt: 1 } } }
])
```

This approach is flexible and can handle any duplicate detection scenario."

### 20. How do you implement a full-text search in MongoDB?

**How to Explain in Interview (Spoken style format):**

"I implement full-text search using MongoDB's text indexes:

**Creating a Text Index:**
```javascript
// Single field text index
db.articles.createIndex({ title: "text" })

// Multiple fields text index with weights
db.articles.createIndex({
  title: "text",
  content: "text",
  tags: "text"
}, {
  weights: {
    title: 10,    // Higher weight = more important
    content: 5,
    tags: 8
  },
  name: "article_text_index"
})
```

**Performing Text Search:**
```javascript
// Basic text search
db.articles.find({ $text: { $search: "mongodb tutorial" } })

// Text search with sorting by relevance score
db.articles.find(
  { $text: { $search: "mongodb tutorial" } },
  { score: { $meta: "textScore" } }
).sort({ score: { $meta: "textScore" } })

// Excluding words from search
db.articles.find({ $text: { $search: "mongodb -tutorial" } })

// Exact phrase search
db.articles.find({ $text: { $search: "\"mongodb tutorial\"" } })
```

**Advanced Search Features:**
```javascript
// Search with additional filters
db.articles.find({
  $text: { $search: "mongodb" },
  category: "technology",
  published: true
})

// Case-insensitive search with diacritic sensitivity
db.articles.find({
  $text: { 
    $search: "café",
    $caseSensitive: false,
    $diacriticSensitive: false
  }
})
```

**Performance Considerations:**
- Text indexes can be large, so I only index relevant fields
- I use compound indexes for common search + filter combinations
- For large datasets, I consider Elasticsearch for advanced search features
- I implement result caching for frequently searched terms

This provides good search capabilities for most applications without needing external search engines."

### 21. What MongoDB operations update documents and create them if they don't exist?

**How to Explain in Interview (Spoken style format):**

"MongoDB provides several operations that can update existing documents or create new ones if they don't exist - this is called 'upsert' behavior.

The main operations that support upsert are:

**`updateOne()` with upsert:**
```javascript
// Updates if found, creates if not found
await User.updateOne(
  { email: 'new@example.com' },
  { $set: { name: 'New User', createdAt: new Date() } },
  { upsert: true }
)
```

**`updateMany()` with upsert:**
```javascript
// Updates multiple matching documents or creates one if none match
await Product.updateMany(
  { category: 'electronics' },
  { $set: { discount: 10 } },
  { upsert: true }
)
```

**`findOneAndUpdate()` with upsert:**
```javascript
// Finds, updates, and returns the document (creates if not found)
const user = await User.findOneAndUpdate(
  { email: 'new@example.com' },
  { $set: { name: 'New User', createdAt: new Date() } },
  { upsert: true, new: true }
)
```

**`replaceOne()` with upsert:**
```javascript
// Replaces entire document or creates new one
await User.replaceOne(
  { email: 'user@example.com' },
  { name: 'Updated User', email: 'user@example.com', age: 30 },
  { upsert: true }
)
```

**Key Points:**
- **upsert: true** enables the 'create if not found' behavior
- **new: true** in `findOneAndUpdate()` returns the created/updated document
- For `updateOne()` and `updateMany()`, use `getUpsertedId()` to get the ID of created documents
- Upsert is useful for idempotent operations, counters, and configuration data

**Real-world Example:**
```javascript
// User login counter - increments if exists, creates with count=1 if new
const result = await User.updateOne(
  { email: 'user@example.com' },
  { 
    $setOnInsert: { name: 'New User', createdAt: new Date() },
    $inc: { loginCount: 1 },
    $set: { lastLogin: new Date() }
  },
  { upsert: true }
)

if (result.upsertedId) {
  console.log('New user created with ID:', result.upsertedId)
} else {
  console.log('Existing user updated')
}
```

I use upsert operations frequently for user profiles, configuration settings, and any data that should be initialized with default values on first access."

---

## Quick Reference Summary

### Most Asked Questions by Company Type:

**Service-Based Companies:**
1. MongoDB vs SQL differences
2. CRUD operations
3. Mongoose basics
4. Schema vs Model
5. Basic relationships

**Product-Based Companies:**
1. Sharding and replica sets
2. Indexing strategies
3. Aggregation framework
4. Transactions and consistency
5. Performance optimization

**Senior/Lead Level:**
1. Schema design for scale
2. Multi-tenant architecture
3. Distributed consistency
4. Advanced aggregation
5. Monitoring and optimization

### Key Concepts to Master:
- Document vs Collection structure
- Embedding vs Referencing
- Index types and optimization
- Aggregation pipeline stages
- Replica set architecture
- Sharding strategies
- Transaction usage patterns
- Performance monitoring tools
- Upsert operations and use cases

---

# Complete Mongoose Keywords Explained (Interview Style)

## 1️⃣ Core Setup Keywords

### mongoose.connect()
**How to Explain in Interview:**
"`mongoose.connect()` is the first thing I use to establish a connection to my MongoDB database. It takes the connection URI and optional configuration options."

```javascript
// Basic connection
mongoose.connect('mongodb://localhost:27017/myapp')

// With options for production
mongoose.connect(process.env.MONGO_URI, {
  useNewUrlParser: true,
  useUnifiedTopology: true,
  maxPoolSize: 10,  // Connection pool size
  serverSelectionTimeoutMS: 5000,
  socketTimeoutMS: 45000,
})
```

**Key Points:**
- Always handle connection errors with `.catch()` or try-catch
- Use environment variables for connection strings
- Configure connection pooling for production
- Handle connection events like 'connected', 'error', 'disconnected'

### mongoose.disconnect()
**How to Explain in Interview:**
"`mongoose.disconnect()` gracefully closes all database connections. I use this in application shutdown handlers or testing cleanup."

```javascript
// Graceful shutdown
async function shutdown() {
  try {
    await mongoose.disconnect()
    console.log('Database disconnected successfully')
    process.exit(0)
  } catch (error) {
    console.error('Error disconnecting:', error)
    process.exit(1)
  }
}

// Handle SIGTERM (Docker, Kubernetes)
process.on('SIGTERM', shutdown)
```

### mongoose.Schema
**How to Explain in Interview:**
"`mongoose.Schema` is the blueprint for my documents. It defines the structure, data types, validation rules, and behavior of my data."

```javascript
const userSchema = new mongoose.Schema({
  name: {
    type: String,
    required: [true, 'Name is required'],
    trim: true,
    maxlength: 50
  },
  email: {
    type: String,
    required: true,
    unique: true,
    lowercase: true
  },
  age: {
    type: Number,
    min: [18, 'Must be at least 18'],
    max: 120
  }
}, {
  timestamps: true,  // Adds createdAt, updatedAt
  toJSON: { virtuals: true },
  toObject: { virtuals: true }
})
```

### mongoose.model()
**How to Explain in Interview:**
"`mongoose.model()` compiles my schema into a model, which gives me the interface to interact with the database collection. The model provides all the CRUD methods."

```javascript
// Compile schema into model
const User = mongoose.model('User', userSchema)

// Now I can use User methods
const user = new User({ name: 'John', email: 'john@example.com' })
await user.save()

// Or use static methods
const users = await User.find({ age: { $gte: 18 } })
```

**Key Points:**
- First argument is the model name (singular, capitalized)
- Second argument is the schema
- Mongoose automatically creates a collection with the plural, lowercase name
- Models are cached, so calling `mongoose.model()` multiple times returns the same model

### mongoose.Types.ObjectId
**How to Explain in Interview:**
"`mongoose.Types.ObjectId` represents MongoDB's unique identifier. I use it for creating references between documents and for querying by ID."

```javascript
// Creating a new ObjectId
const newId = new mongoose.Types.ObjectId()

// Validating if a string is a valid ObjectId
if (mongoose.Types.ObjectId.isValid(userId)) {
  const user = await User.findById(userId)
}

// Using in schema for references
const postSchema = new mongoose.Schema({
  author: {
    type: mongoose.Types.ObjectId,
    ref: 'User',
    required: true
  }
})
```

---

## 2️⃣ Schema Data Types

### String
**How to Explain in Interview:**
"String is for text data. I can add validation like required, minlength, maxlength, enum for predefined values, and transformations like lowercase, trim."

```javascript
name: {
  type: String,
  required: true,
  trim: true,
  minlength: 2,
  maxlength: 50,
  enum: ['John', 'Jane', 'Bob']  // Only these values allowed
}
```

### Number
**How to Explain in Interview:**
"Number handles numeric data. I can set min/max constraints and use it for calculations, sorting, and range queries."

```javascript
age: {
  type: Number,
  min: 0,
  max: 150,
  default: 0
},
price: {
  type: Number,
  min: [0, 'Price cannot be negative'],
  get: v => v.toFixed(2)  // Format when retrieving
}
```

### Boolean
**How to Explain in Interview:**
"Boolean is for true/false values. Perfect for flags like isActive, isAdmin, hasPaid. I can set default values."

```javascript
isActive: {
  type: Boolean,
  default: true
},
isAdmin: {
  type: Boolean,
  default: false
}
```

### Date
**How to Explain in Interview:**
"Date handles timestamps. I use it for createdAt, updatedAt, birthDate, expirationDate. Mongoose can auto-manage timestamps."

```javascript
birthDate: {
  type: Date,
  required: true,
  validate: {
    validator: function(value) {
      return value < new Date()  // Must be in the past
    },
    message: 'Birth date must be in the past'
  }
},
expiresAt: {
  type: Date,
  default: Date.now,
  expires: 3600  // Auto-delete after 1 hour
}
```

### Buffer
**How to Explain in Interview:**
"Buffer stores binary data like images, files, or encrypted data. It's useful when I need to store small files directly in the database."

```javascript
profilePicture: {
  type: Buffer,
  required: false
},
encryptedData: {
  type: Buffer,
  get: function(data) {
    // Decrypt when retrieving
    return decrypt(data)
  },
  set: function(data) {
    // Encrypt when saving
    return encrypt(data)
  }
}
```

### ObjectId
**How to Explain in Interview:**
"ObjectId creates references to other documents. This is how I implement relationships in MongoDB."

```javascript
author: {
  type: mongoose.Schema.Types.ObjectId,
  ref: 'User',
  required: true
},
comments: [{
  type: mongoose.Schema.Types.ObjectId,
  ref: 'Comment'
}]
```

### Array
**How to Explain in Interview:**
"Array stores lists of data. I can have arrays of strings, numbers, objects, or even references to other documents."

```javascript
tags: [String],  // Array of strings
scores: [Number],  // Array of numbers
addresses: [{  // Array of objects
  street: String,
  city: String,
  isPrimary: {
    type: Boolean,
    default: false
  }
}],
friends: [{  // Array of references
  type: mongoose.Schema.Types.ObjectId,
  ref: 'User'
}]
```

### Map
**How to Explain in Interview:**
"Map is for dynamic key-value pairs. Perfect for when I don't know the exact keys beforehand."

```javascript
metadata: {
  type: Map,
  of: String  // Values will be strings
},
preferences: {
  type: Map,
  of: {  // Nested schema for values
    theme: { type: String, enum: ['light', 'dark'] },
    notifications: Boolean
  }
}
```

### Mixed
**How to Explain in Interview:**
"Mixed accepts any data type. It's flexible but should be used sparingly since it loses type safety and validation benefits."

```javascript
extraData: {
  type: mongoose.Schema.Types.Mixed,
  default: {}
},
config: {
  type: mongoose.Schema.Types.Mixed,
  validate: {
    validator: function(value) {
      // Custom validation for mixed type
      return typeof value === 'object'
    },
    message: 'Config must be an object'
  }
}
```

### Decimal128
**How to Explain in Interview:**
"Decimal128 provides high-precision decimal numbers. Perfect for financial calculations where floating-point precision issues matter."

```javascript
price: {
  type: mongoose.Schema.Types.Decimal128,
  required: true,
  min: 0
},
balance: {
  type: mongoose.Schema.Types.Decimal128,
  default: 0.00,
  get: function(value) {
    return parseFloat(value.toString())
  }
}
```

### UUID
**How to Explain in Interview:**
"UUID stores universally unique identifiers. I use this when I need globally unique IDs or when integrating with external systems."

```javascript
identifier: {
  type: mongoose.Schema.Types.UUID,
  default: () => mongoose.Types.UUID.randomUUID(),
  unique: true,
  required: true
},
externalId: {
  type: mongoose.Schema.Types.UUID,
  required: false  // Optional external system ID
}
```

---

## 3️⃣ Schema Field Options

### required
**How to Explain in Interview:**
"`required` ensures a field must be present. I can use it as a boolean or provide a custom error message."

```javascript
email: {
  type: String,
  required: true  // Simple required
},
name: {
  type: String,
  required: [true, 'Name is mandatory']  // Custom message
},
age: {
  type: Number,
  required: function() {
    return this.isActive  // Conditional requirement
  }
}
```

### default
**How to Explain in Interview:**
"`default` provides automatic values when a field isn't specified. I can use static values, functions, or even reference other fields."

```javascript
status: {
  type: String,
  enum: ['active', 'inactive'],
  default: 'active'  // Static default
},
createdAt: {
  type: Date,
  default: Date.now  // Function default
},
code: {
  type: String,
  default: function() {
    return Math.random().toString(36).substr(2, 9)  // Generated default
  }
},
fullName: {
  type: String,
  default: function() {
    return `${this.firstName} ${this.lastName}`  // Reference other fields
  }
}
```

### unique
**How to Explain in Interview:**
"`unique` creates a unique index to prevent duplicate values. I use this for emails, usernames, and other identifiers."

```javascript
email: {
  type: String,
  unique: true,
  required: true
},
username: {
  type: String,
  unique: true,
  sparse: true  // Allow multiple nulls
}
```

### lowercase/uppercase
**How to Explain in Interview:**
"`lowercase` and `uppercase` automatically transform string values. Great for normalizing emails, usernames, and codes."

```javascript
email: {
  type: String,
  lowercase: true,
  trim: true
},
code: {
  type: String,
  uppercase: true
}
```

### trim
**How to Explain in Interview:**
"`trim` removes whitespace from the beginning and end of strings. Essential for user input to prevent issues with spaces."

```javascript
name: {
  type: String,
  required: true,
  trim: true
},
searchTerm: {
  type: String,
  trim: true,
  lowercase: true
}
```

### minlength/maxlength
**How to Explain in Interview:**
"`minlength` and `maxlength` validate string length. I use these for passwords, usernames, and content fields."

```javascript
password: {
  type: String,
  required: true,
  minlength: [8, 'Password must be at least 8 characters'],
  maxlength: 128
},
username: {
  type: String,
  minlength: 3,
  maxlength: 30,
  unique: true
}
```

### enum
**How to Explain in Interview:**
"`enum` restricts values to a predefined set. Perfect for status fields, categories, and any field with limited options."

```javascript
status: {
  type: String,
  enum: ['pending', 'approved', 'rejected'],
  default: 'pending'
},
priority: {
  type: String,
  enum: {
    values: ['low', 'medium', 'high'],
    message: 'Priority must be low, medium, or high'
  }
}
```

### match
**How to Explain in Interview:**
"`match` uses regular expressions for validation. I use it for email formats, phone numbers, and custom patterns."

```javascript
email: {
  type: String,
  required: true,
  match: [/^\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/, 'Invalid email format']
},
phone: {
  type: String,
  match: [/^\+?[1-9]\d{1,14}$/, 'Invalid phone number format']
},
zipCode: {
  type: String,
  match: [/^\d{5}(-\d{4})?$/, 'Invalid ZIP code format']
}
```

### select
**How to Explain in Interview:**
"`select` controls whether fields are included by default in queries. I use this to exclude sensitive data like passwords."

```javascript
password: {
  type: String,
  required: true,
  select: false  // Excluded by default
},
secretKey: {
  type: String,
  select: false
},
publicInfo: {
  type: String,
  select: true  // Included by default (default behavior)
}

// To include excluded fields:
const user = await User.findById(id).select('+password')
```

### immutable
**How to Explain in Interview:**
"`immutable` prevents fields from being changed after creation. Perfect for timestamps, original values, and audit fields."

```javascript
createdAt: {
  type: Date,
  default: Date.now,
  immutable: true
},
originalPrice: {
  type: Number,
  required: true,
  immutable: true
},
version: {
  type: Number,
  default: 1,
  immutable: true
}
```

### alias
**How to Explain in Interview:**
"`alias` provides an alternative field name. Useful for API compatibility or when the database field name differs from the API field name."

```javascript
firstName: {
  type: String,
  required: true,
  alias: 'first_name'  // Can also access as first_name
},
lastName: {
  type: String,
  alias: 'last_name'
}

// Both work:
user.firstName  // or user.first_name
user.firstName = 'John'  // or user.first_name = 'John'
```

### index
**How to Explain in Interview:**
"`index` creates database indexes for better query performance. I use this on frequently queried fields."

```javascript
email: {
  type: String,
  required: true,
  unique: true,
  index: true  // Creates unique index automatically
},
name: {
  type: String,
  index: true  // Simple index
},
createdAt: {
  type: Date,
  index: -1  // Descending index
}
```

### sparse
**How to Explain in Interview:**
"`sparse` creates indexes that only include documents with the indexed field. Perfect for optional unique fields."

```javascript
email: {
  type: String,
  unique: true,
  sparse: true  // Allows multiple nulls
},
referralCode: {
  type: String,
  unique: true,
  sparse: true  // Only index if field exists
}
```

---

## 4️⃣ Schema Options

### timestamps
**How to Explain in Interview:**
"`timestamps` automatically adds createdAt and updatedAt fields. It's a convenient way to track when documents are created and modified."

```javascript
const userSchema = new mongoose.Schema({
  name: String,
  email: String
}, {
  timestamps: true  // Adds createdAt, updatedAt
})

// Custom timestamp field names
const postSchema = new mongoose.Schema({
  title: String,
  content: String
}, {
  timestamps: {
    createdAt: 'created_at',
    updatedAt: 'updated_at'
  }
})
```

### versionKey
**How to Explain in Interview:**
"`versionKey` controls the document versioning field. By default, Mongoose adds `__v` field for version control."

```javascript
const schema = new mongoose.Schema({
  name: String
}, {
  versionKey: '__v'  // Default behavior
})

// Disable versioning
const schemaNoVersion = new mongoose.Schema({
  name: String
}, {
  versionKey: false
})

// Custom version key
const schemaCustomVersion = new mongoose.Schema({
  name: String
}, {
  versionKey: 'version'
})
```

### collection
**How to Explain in Interview:**
"`collection` specifies a custom collection name. By default, Mongoose uses the pluralized, lowercase model name."

```javascript
const User = mongoose.model('User', userSchema, 'users')  // Explicit
const User = mongoose.model('User', userSchema, { collection: 'app_users' })  // Custom name
```

### strict
**How to Explain in Interview:**
"`strict` controls whether fields not in the schema are allowed. I usually keep it true for data integrity."

```javascript
const strictSchema = new mongoose.Schema({
  name: String
}, {
  strict: true  // Default - rejects unknown fields
})

const flexibleSchema = new mongoose.Schema({
  name: String
}, {
  strict: false  // Allows unknown fields
})

// Throw error for unknown fields
const throwSchema = new mongoose.Schema({
  name: String
}, {
  strict: 'throw'
})
```

### minimize
**How to Explain in Interview:**
"`minimize` removes empty objects from documents. Helps keep documents clean and save space."

```javascript
const schema = new mongoose.Schema({
  name: String,
  metadata: {
    type: Map,
    of: String
  }
}, {
  minimize: true  // Default - removes empty objects
})

// Document { name: 'John', metadata: {} } becomes { name: 'John' }
```

### toJSON/toObject
**How to Explain in Interview:**
"`toJSON` and `toObject` customize how documents are converted to JSON or plain objects. I use these to control output format."

```javascript
const schema = new mongoose.Schema({
  name: String,
  password: String,
  secret: String
}, {
  toJSON: {
    transform: function(doc, ret) {
      delete ret.password
      delete ret.secret
      delete ret.__v
      return ret
    }
  },
  toObject: {
    virtuals: true,
    transform: function(doc, ret) {
      ret.id = ret._id
      delete ret._id
      return ret
    }
  }
})
```

---

## 5️⃣ Document Methods

### save()
**How to Explain in Interview:**
"`save()` persists a document to the database. I use it for both creating new documents and updating existing ones."

```javascript
// Creating a new document
const user = new User({ name: 'John', email: 'john@example.com' })
await user.save()

// Updating an existing document
const user = await User.findById(id)
user.name = 'Jane'
await user.save()

// With validation options
await user.save({ validateBeforeSave: false })

// Handling errors
try {
  await user.save()
} catch (error) {
  if (error.name === 'ValidationError') {
    // Handle validation errors
  }
}
```

### remove()
**How to Explain in Interview:**
"`remove()` deletes a document from the database. I use this when I have the document instance already loaded."

```javascript
const user = await User.findById(id)
await user.remove()

// Alternative with callback (older syntax)
user.remove(function(err) {
  if (err) return handleError(err)
  console.log('Document removed')
})
```

### validate()
**How to Explain in Interview:**
"`validate()` runs schema validation without saving. I use this for pre-save validation or form validation."

```javascript
const user = new User({ name: '', email: 'invalid' })

try {
  await user.validate()
  console.log('Document is valid')
} catch (error) {
  console.log('Validation errors:', error.errors)
}

// Validate specific fields
await user.validateSync(['name', 'email'])
```

### isModified()
**How to Explain in Interview:**
"`isModified()` checks if a field has been changed. I use this in middleware and conditional logic."

```javascript
const user = await User.findById(id)
user.name = 'New Name'

if (user.isModified('name')) {
  console.log('Name was modified')
}

if (user.isModified(['name', 'email'])) {
  console.log('Name or email was modified')
}

if (user.isModified()) {
  console.log('Document was modified')
}

// Get modified paths
const modifiedPaths = user.modifiedPaths()
```

### toObject()
**How to Explain in Interview:**
"`toObject()` converts a Mongoose document to a plain JavaScript object. I use this when I need to manipulate the data without Mongoose methods."

```javascript
const user = await User.findById(id)
const plainObject = user.toObject()

// With options
const userObject = user.toObject({
  virtuals: true,  // Include virtual fields
  getters: true,   // Apply getters
  transform: function(doc, ret) {
    delete ret.__v
    return ret
  }
})
```

### toJSON()
**How to Explain in Interview:**
"`toJSON()` converts a document to JSON format. This is automatically called when sending responses in Express."

```javascript
const user = await User.findById(id)
const json = user.toJSON()

// Express automatically calls toJSON()
res.json(user)  // Uses toJSON() internally

// Custom JSON output
const customJson = user.toJSON({
  transform: function(doc, ret) {
    ret.id = ret._id
    delete ret._id
    delete ret.__v
    delete ret.password
    return ret
  }
})
```

---

## 6️⃣ Model Query Methods

### create()
**How to Explain in Interview:**
"`create()` is a convenient method that combines creating a new document and saving it in one operation."

```javascript
// Create single document
const user = await User.create({
  name: 'John Doe',
  email: 'john@example.com',
  age: 30
})

// Create multiple documents
const users = await User.create([
  { name: 'Alice', email: 'alice@example.com' },
  { name: 'Bob', email: 'bob@example.com' }
])

// With options
const user = await User.create(
  { name: 'John', email: 'john@example.com' },
  { validateBeforeSave: false }
)
```

### find()
**How to Explain in Interview:**
"`find()` retrieves multiple documents that match the query criteria. It returns a query that I can chain with other methods."

```javascript
// Find all documents
const users = await User.find()

// Find with criteria
const activeUsers = await User.find({ isActive: true })

// With operators
const adults = await User.find({ age: { $gte: 18 } })

// Chain with other methods
const users = await User.find({ status: 'active' })
  .sort({ name: 1 })
  .limit(10)
  .select('name email')
```

### findOne()
**How to Explain in Interview:**
"`findOne()` retrieves the first document that matches the query criteria. I use this when I expect only one result."

```javascript
// Find one by criteria
const user = await User.findOne({ email: 'john@example.com' })

// With sorting
const oldestUser = await User.findOne().sort({ age: -1 })

// Chain with select
const user = await User.findOne({ email: email })
  .select('name email age')
  .populate('friends')
```

### findById()
**How to Explain in Interview:**
"`findById()` retrieves a document by its _id. This is the most efficient way to get a specific document."

```javascript
// Find by ID
const user = await User.findById('507f1f77bcf86cd799439011')

// With field selection
const user = await User.findById(id, 'name email')

// With options
const user = await User.findById(id, null, { lean: true })

// Handle not found
const user = await User.findById(id)
if (!user) {
  throw new Error('User not found')
}
```

### updateOne()
**How to Explain in Interview:**
"`updateOne()` updates a single document that matches the query. I use update operators like $set, $push, $pull."

```javascript
// Basic update
const result = await User.updateOne(
  { _id: userId },
  { $set: { name: 'New Name' } }
)

// Multiple fields
await User.updateOne(
  { email: 'john@example.com' },
  {
    $set: { name: 'John Updated', age: 31 },
    $inc: { loginCount: 1 },
    $push: { tags: 'premium' }
  }
)

// With upsert (create if not found)
await User.updateOne(
  { email: 'new@example.com' },
  { $set: { name: 'New User' } },
  { upsert: true }
)
```

### updateMany()
**How to Explain in Interview:**
"`updateMany()` updates all documents that match the query. I use this for bulk operations."

```javascript
// Update all active users
const result = await User.updateMany(
  { isActive: true },
  { $set: { lastLogin: new Date() } }
)

// Bulk update with conditions
await User.updateMany(
  { age: { $lt: 18 } },
  { $set: { status: 'minor' } }
)

// Result contains matchedCount and modifiedCount
console.log(`Matched: ${result.matchedCount}, Modified: ${result.modifiedCount}`)
```

### deleteOne()
**How to Explain in Interview:**
"`deleteOne()` deletes the first document that matches the query. I use this for removing specific documents."

```javascript
// Delete by criteria
await User.deleteOne({ email: 'john@example.com' })

// Delete by ID
await User.deleteOne({ _id: userId })

// Returns deletion result
const result = await User.deleteOne({ name: 'John' })
console.log(`Deleted ${result.deletedCount} document`)
```

### deleteMany()
**How to Explain in Interview:**
"`deleteMany()` deletes all documents that match the query. I use this for bulk deletions."

```javascript
// Delete all inactive users
const result = await User.deleteMany({ isActive: false })

// Delete by date range
await User.deleteMany({
  createdAt: {
    $lt: new Date('2020-01-01')
  }
})

// Delete all (with caution)
await User.deleteMany({})
```

### countDocuments()
**How to Explain in Interview:**
"`countDocuments()` counts documents matching the query. It's accurate but slower as it scans the collection."

```javascript
// Count all documents
const totalUsers = await User.countDocuments()

// Count with criteria
const activeUsers = await User.countDocuments({ isActive: true })

// Count with complex query
const adultUsers = await User.countDocuments({
  age: { $gte: 18 },
  status: 'active'
})
```

### estimatedDocumentCount()
**How to Explain in Interview:**
"`estimatedDocumentCount()` provides an estimated count using collection metadata. It's faster but less accurate."

```javascript
// Fast estimate of total documents
const estimatedTotal = await User.estimatedDocumentCount()

// Good for pagination totals
const total = await User.estimatedDocumentCount()
const users = await User.find().limit(10)
```

---

## 7️⃣ Advanced Query Methods

### findOneAndUpdate()
**How to Explain in Interview:**
"`findOneAndUpdate()` finds a document, updates it, and returns the updated document. It's atomic and efficient."

```javascript
// Update and return the new document
const updatedUser = await User.findOneAndUpdate(
  { email: 'john@example.com' },
  { $set: { name: 'John Updated', lastLogin: new Date() } },
  { new: true }  // Return the updated document
)

// Return the original document
const originalUser = await User.findOneAndUpdate(
  { _id: userId },
  { $inc: { loginCount: 1 } },
  { new: false }  // Default behavior
)

// With upsert
const user = await User.findOneAndUpdate(
  { email: 'new@example.com' },
  { $set: { name: 'New User', createdAt: new Date() } },
  { upsert: true, new: true }
)
```

### findOneAndDelete()
**How to Explain in Interview:**
"`findOneAndDelete()` finds a document, deletes it, and returns the deleted document."

```javascript
// Delete and return the deleted document
const deletedUser = await User.findOneAndDelete({
  email: 'john@example.com'
})

// With sorting (deletes the first match after sorting)
const deletedUser = await User.findOneAndDelete(
  { status: 'inactive' },
  { sort: { createdAt: 1 } }
)

// With projection
const deletedUser = await User.findOneAndDelete(
  { _id: userId },
  { projection: { name: 1, email: 1 } }
)
```

### findByIdAndUpdate()
**How to Explain in Interview:**
"`findByIdAndUpdate()` is a shortcut for finding by ID and updating. It's commonly used for API endpoints."

```javascript
// Update user by ID
const updatedUser = await User.findByIdAndUpdate(
  userId,
  { $set: { name: 'Updated Name' } },
  { new: true, runValidators: true }
)

// Complex update
const user = await User.findByIdAndUpdate(
  userId,
  {
    $set: { profile: { bio: 'Updated bio' } },
    $push: { tags: 'updated' },
    $unset: { temporaryField: 1 }
  },
  { new: true }
)
```

### findByIdAndDelete()
**How to Explain in Interview:**
"`findByIdAndDelete()` is a shortcut for finding by ID and deleting."

```javascript
// Delete user by ID
const deletedUser = await User.findByIdAndDelete(userId)

// Handle not found
const user = await User.findByIdAndDelete(userId)
if (!user) {
  throw new Error('User not found')
}
```

### lean()
**How to Explain in Interview:**
"`lean()` returns plain JavaScript objects instead of Mongoose documents. It's much faster for read operations."

```javascript
// Fast queries for read-only operations
const users = await User.find({ isActive: true }).lean()

// Good for API responses
const users = await User.find()
  .select('name email')
  .lean()
  .limit(100)

// No Mongoose methods available
users[0].save // undefined
users[0].name // Works - it's a plain object
```

### limit()
**How to Explain in Interview:**
"`limit()` restricts the number of documents returned. Essential for pagination and performance."

```javascript
// Get first 10 users
const users = await User.find().limit(10)

// Pagination
const page = 2
const limit = 10
const skip = (page - 1) * limit

const users = await User.find()
  .skip(skip)
  .limit(limit)
```

### skip()
**How to Explain in Interview:**
"`skip()` skips a specified number of documents. Used with limit() for pagination."

```javascript
// Skip first 5 documents
const users = await User.find().skip(5)

// Pagination
const getUsers = async (page = 1, limit = 10) => {
  const skip = (page - 1) * limit
  return await User.find()
    .skip(skip)
    .limit(limit)
    .sort({ createdAt: -1 })
}
```

### sort()
**How to Explain in Interview:**
"`sort()` orders the results. 1 for ascending, -1 for descending."

```javascript
// Sort by name ascending
const users = await User.find().sort({ name: 1 })

// Sort by age descending, then name ascending
const users = await User.find().sort({ age: -1, name: 1 })

// Sort by date
const recentUsers = await User.find().sort({ createdAt: -1 })

// Text search relevance sorting
const articles = await Article.find(
  { $text: { $search: 'mongodb' } },
  { score: { $meta: 'textScore' } }
).sort({ score: { $meta: 'textScore' } })
```

### select()
**How to Explain in Interview:**
"`select()` specifies which fields to include or exclude in the results. Good for performance and security."

```javascript
// Include only specific fields
const users = await User.find().select('name email')

// Exclude specific fields
const users = await User.find().select('-password -secret')

// Mix include and exclude
const users = await User.find()
  .select('name email -password')

// Include virtual fields
const users = await User.find()
  .select('name email')
  .populate('friends')
```

---

## 8️⃣ Reference & Population

### ref
**How to Explain in Interview:**
"`ref` creates references between documents. It stores the ObjectId of another document and allows population."

```javascript
// Reference another model
const postSchema = new mongoose.Schema({
  title: String,
  author: {
    type: mongoose.Schema.Types.ObjectId,
    ref: 'User',  // Reference to User model
    required: true
  },
  comments: [{
    type: mongoose.Schema.Types.ObjectId,
    ref: 'Comment'
  }]
})

// Reference with conditions
const categorySchema = new mongoose.Schema({
  name: String,
  posts: [{
    type: mongoose.Schema.Types.ObjectId,
    ref: 'Post'
  }]
})
```

### populate()
**How to Explain in Interview:**
"`populate()` replaces references with actual documents. It's like JOIN in SQL but more flexible."

```javascript
// Basic population
const posts = await Post.find().populate('author')

// Populate multiple fields
const posts = await Post.find()
  .populate('author')
  .populate('comments')

// Populate with field selection
const posts = await Post.find()
  .populate('author', 'name email avatar')
  .populate('comments', 'text createdAt')

// Nested population
const posts = await Post.find()
  .populate({
    path: 'comments',
    populate: {
      path: 'author',
      select: 'name'
    }
  })

// Population with conditions
const activePosts = await Post.find()
  .populate({
    path: 'author',
    match: { isActive: true },
    select: 'name email'
  })

// Limit populated documents
const posts = await Post.find()
  .populate({
    path: 'comments',
    options: { limit: 5, sort: { createdAt: -1 } }
  })
```

---

## 9️⃣ Middleware Hooks

### pre() / post()
**How to Explain in Interview:**
"`pre()` and `post()` hooks allow me to run code before or after specific operations. Great for validation, logging, and data transformation."

```javascript
// Pre-save hook
userSchema.pre('save', function(next) {
  if (this.isModified('password')) {
    this.password = bcrypt.hashSync(this.password, 10)
  }
  next()
})

// Post-save hook
userSchema.post('save', function(doc) {
  console.log('User saved:', doc._id)
  // Send welcome email
})

// Pre-remove hook
userSchema.pre('remove', function(next) {
  // Clean up related data
  Post.deleteMany({ author: this._id }).exec()
  next()
})
```

### Hook Types
**How to Explain in Interview:**
"Different hooks run at different times. I use save for document operations, validate for validation, and find for query operations."

```javascript
// Save hook - runs before save
schema.pre('save', function(next) {
  this.updatedAt = new Date()
  next()
})

// Validate hook - runs during validation
schema.pre('validate', function(next) {
  if (this.email && !this.email.includes('@')) {
    next(new Error('Invalid email'))
  } else {
    next()
  }
})

// Find hook - modifies queries
schema.pre(/^find/, function(next) {
  this.where({ isActive: true })  // Only find active documents
  next()
})

// Update hook
schema.pre('updateOne', function(next) {
  this.set({ updatedAt: new Date() })
  next()
})

// Delete hook
schema.pre('deleteOne', function(next) {
  console.log('Deleting document')
  next()
})
```

---

## 🔟 Advanced Schema Features

### schema.methods
**How to Explain in Interview:**
"`schema.methods` adds instance methods to documents. These are available on document instances."

```javascript
// Add instance method
userSchema.methods.getFullName = function() {
  return `${this.firstName} ${this.lastName}`
}

// Async instance method
userSchema.methods.updateProfile = async function(profileData) {
  Object.assign(this.profile, profileData)
  return await this.save()
}

// Usage
const user = await User.findById(id)
const fullName = user.getFullName()
await user.updateProfile({ bio: 'New bio' })
```

### schema.statics
**How to Explain in Interview:**
"`schema.statics` adds static methods to the model. These are available on the model itself."

```javascript
// Add static method
userSchema.statics.findByEmail = function(email) {
  return this.findOne({ email: email })
}

// Static method with aggregation
userSchema.statics.getStats = function() {
  return this.aggregate([
    { $group: { _id: null, total: { $sum: 1 } } }
  ])
}

// Usage
const user = await User.findByEmail('john@example.com')
const stats = await User.getStats()
```

### schema.virtual()
**How to Explain in Interview:**
"`schema.virtual()` creates virtual fields that don't exist in the database but are computed on the fly."

```javascript
// Virtual field
userSchema.virtual('fullName').get(function() {
  return `${this.firstName} ${this.lastName}`
})

// Virtual with setter
userSchema.virtual('fullName').set(function(value) {
  const parts = value.split(' ')
  this.firstName = parts[0]
  this.lastName = parts[1]
})

// Virtual for population
userSchema.virtual('posts', {
  ref: 'Post',
  localField: '_id',
  foreignField: 'author'
})

// Usage
const user = await User.findById(id).populate('posts')
console.log(user.fullName)  // Virtual field
```

### schema.query
**How to Explain in Interview:**
"`schema.query` adds custom query helpers that can be chained with other query methods."

```javascript
// Custom query helper
userSchema.query.byName = function(name) {
  return this.find({ name: new RegExp(name, 'i') })
}

// Chainable query helpers
userSchema.query.active = function() {
  return this.where({ isActive: true })
}

userSchema.query.adults = function() {
  return this.where({ age: { $gte: 18 } })
}

// Usage
const users = await User.find()
  .byName('john')
  .active()
  .adults()
  .limit(10)
```

### plugin()
**How to Explain in Interview:**
"`plugin()` adds reusable functionality to schemas. Great for common features like timestamps, pagination, and soft deletes."

```javascript
// Create a plugin
const timestampPlugin = function(schema, options) {
  schema.add({ createdAt: Date, updatedAt: Date })
  
  schema.pre('save', function(next) {
    if (!this.createdAt) this.createdAt = new Date()
    this.updatedAt = new Date()
    next()
  })
}

// Use the plugin
userSchema.plugin(timestampPlugin)
postSchema.plugin(timestampPlugin)

// Popular plugins
const mongoosePaginate = require('mongoose-paginate')
schema.plugin(mongoosePaginate)
```

### discriminator()
**How to Explain in Interview:**
"`discriminator()` creates schema inheritance. Perfect for when I have a base model with different types."

```javascript
// Base schema
const personSchema = new mongoose.Schema({
  name: String,
  email: String
}, { discriminatorKey: 'type' })

const Person = mongoose.model('Person', personSchema)

// Employee discriminator
const Employee = Person.discriminator('Employee',
  new mongoose.Schema({
    employeeId: String,
    department: String
  })
)

// Customer discriminator
const Customer = Person.discriminator('Customer',
  new mongoose.Schema({
    customerId: String,
    orders: [String]
  })
)

// Usage
const employee = await Employee.findOne({ name: 'John' })
const customer = await Customer.findOne({ name: 'Jane' })
const allPeople = await Person.find()  // Returns all types
```

---

## Summary

This covers all 60+ Mongoose keywords with practical examples and interview-style explanations. The key is understanding not just what each keyword does, but when and why to use it in real applications.

**Most Important for Interviews:**
1. Schema and Model fundamentals
2. CRUD operations (create, find, update, delete)
3. Population and references
4. Middleware hooks
5. Validation and field options
6. Query methods and chaining
7. Performance considerations (lean, indexes, pagination)

**Remember to explain:**
- What the keyword does
- When to use it
- Why it's useful
- Practical examples
- Performance implications
- Common pitfalls to avoid
