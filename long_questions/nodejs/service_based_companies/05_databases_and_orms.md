# 🗄️ 05 — Databases and ORMs (Mongoose/Sequelize)
> **Most Asked in Service-Based Companies** | 🗄️ Difficulty: Medium

---

## 🔑 Must-Know Topics
- Connecting Node.js to a database (MongoDB/MySQL)
- SQL vs NoSQL considerations
- Mongoose (Schemas, Models, Queries)
- Object-Relational Mapping (ORM) vs Object Document Mapper (ODM)
- Database Connection Pooling
- Preventing SQL Injection

---

## ❓ Frequently Asked Questions

### Q1. What is the difference between an ODM and an ORM in Node.js?

**Answer:**
Both serve to act as an abstraction layer between your application code and the database, translating objects in code to data rows/documents in the database.

- **ORM (Object-Relational Mapping):** Used for **SQL/Relational Databases** (e.g., MySQL, PostgreSQL). It maps JavaScript objects to tables and rows.
  - *Popular Node.js ORMs:* Sequelize, TypeORM, Prisma.
- **ODM (Object Document Mapper):** Used for **NoSQL/Document Databases** (e.g., MongoDB). It maps JavaScript objects to JSON-like documents.
  - *Popular Node.js ODM:* Mongoose.

---

### Q2. How do you connect a Node.js application to MongoDB using Mongoose?

**Answer:**
Mongoose is the standard ODM for MongoDB and Node.js.

```javascript
const mongoose = require('mongoose');

// Connect to MongoDB
mongoose.connect('mongodb://localhost:27017/myDatabase', {
    useNewUrlParser: true,
    useUnifiedTopology: true
})
.then(() => console.log('Successfully connected to MongoDB!'))
.catch((error) => console.error('Connection failed:', error));

// Define a Schema
const userSchema = new mongoose.Schema({
    name: { type: String, required: true },
    email: { type: String, required: true, unique: true },
    age: Number
});

// Compile into a Model
const User = mongoose.model('User', userSchema);

// Example usage
async function createUser() {
    const newUser = new User({ name: 'Alice', email: 'alice@example.com', age: 25 });
    await newUser.save();
}
```

---

### Q3. Explain the difference between `Schema` and `Model` in Mongoose.

**Answer:**

- **Schema:** A blueprint that defines the structure of the document, default values, validators, etc. It does NOT interact directly with the database.
- **Model:** A compiled version of the Schema. It acts as a constructor function that creates and reads documents from the underlying MongoDB database. It provides the interface to interact with the database (e.g., `.save()`, `.find()`, `.updateOne()`).

---

### Q4. How do you handle relationships in MongoDB (Mongoose)? What is `populate()`?

**Answer:**
Although MongoDB is NoSQL, you often need to reference documents in other collections. We achieve this using `refs` and the `populate()` method.

```javascript
// Author Schema
const authorSchema = new mongoose.Schema({
    name: String,
    bio: String
});
const Author = mongoose.model('Author', authorSchema);

// Book Schema (References Author)
const bookSchema = new mongoose.Schema({
    title: String,
    authorId: { type: mongoose.Schema.Types.ObjectId, ref: 'Author' } // Reference
});
const Book = mongoose.model('Book', bookSchema);

// Querying with Populate
async function getBookWithAuthor() {
    // Instead of just getting the authorId string, populate() replaces it with the actual Author document.
    const book = await Book.findOne({ title: 'Node.js Guide' }).populate('authorId');
    console.log(book.authorId.name);
}
```

---

### Q5. How does Database Connection Pooling work, and why is it important?

**Answer:**
Establishing a new database connection involves network latency and authentication, making it a "heavy" operation. 

**Connection Pooling** maintains a cache ("pool") of pre-established, open connections to the database. When your Node.js app needs to query the database, it requests an open connection from the pool.
- Once the query finishes, the connection is returned to the pool to be reused, rather than closed.
- **Why it's important:** It dramatically improves the performance and scalability of an application, reducing latency for incoming requests.

*(Libraries like Mongoose and `pg` handle connection pooling by default).*

---

### Q6. How do you prevent SQL Injection in Node.js?

**Answer:**
SQL Injection occurs when user input is concatenated directly into a database query, allowing attackers to execute malicious SQL statements.

**How to prevent it:**
1. **Use Parameterized Queries / Prepared Statements:** Instead of combining strings, use parameters (`?` or `$1`). The database driver ensures the user input is treated strictly as data, not executable code.
   ```javascript
   // VULNERABLE
   db.query(`SELECT * FROM users WHERE email = '${req.body.email}'`);

   // SECURE (Parameterized)
   db.query('SELECT * FROM users WHERE email = ?', [req.body.email]);
   ```
2. **Use an ORM/Query Builder:** Libraries like Sequelize or Knex.js automatically escape user input.
   ```javascript
   // SECURE (Sequelize)
   User.findAll({ where: { email: req.body.email } });
   ```
3. **Input Validation & Sanitization:** Ensure the input is exactly what you expect (e.g., using libraries like Joi or express-validator).
