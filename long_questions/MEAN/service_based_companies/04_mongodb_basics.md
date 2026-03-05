# MongoDB Basics (Service-Based Companies)

MongoDB is the chosen database for both MEAN and MERN stacks. You are expected to know NoSQL basics, document structures, basic CRUD operations, and how to interface with MongoDB using an ODM like Mongoose.

## Core MongoDB Concepts

### 1. What is MongoDB? How does it differ from a traditional Relational Database (RDBMS)?
*   **MongoDB**: A NoSQL, document-oriented database. Data is stored in collections as flexible, JSON-like documents (BSON). It scales horizontally well and allows for a flexible, schema-less (or dynamic schema) design.
*   **RDBMS (SQL)**: Stores data in rigid tables with rows and columns. Schemas must be predefined. Standardizes relations using Foreign Keys and relies heavily on complex JOINs and ACID transactions (though modern MongoDB supports multi-document transactions too).

### 2. What are Collections and Documents?
*   **Document**: The basic unit of data in MongoDB, consisting of key-value pairs (similar to a JSON object). E.g., `{ "name": "John", "age": 30 }`. Corresponds to a "Row" in SQL.
*   **Collection**: A grouping of MongoDB documents. Corresponds to a "Table" in SQL. Documents within a single collection do not enforce a strict schema.

### 3. What is the `_id` field?
Every document in MongoDB requires an `_id` field that acts as a primary key. It must be unique within a collection. If you don't provide one upon insertion, MongoDB automatically generates a 12-byte `ObjectId` that contains a timestamp, machine identifier, process id, and an incrementing counter.

## CRUD Operations

### 4. Provide basic examples of CRUD (Create, Read, Update, Delete) operations in the MongoDB shell.
*   **Create**: `db.users.insertOne({ name: "Alice", age: 25 })`
*   **Read**:
    *   Find all: `db.users.find()`
    *   Find with condition: `db.users.find({ age: { $gt: 20 } })`
*   **Update**: `db.users.updateOne({ name: "Alice" }, { $set: { age: 26 } })`
*   **Delete**: `db.users.deleteOne({ name: "Alice" })`

### 5. Explain some basic query operators: `$eq`, `$gt`, `$lt`, `$in`, `$or`, `$and`.
*   `$eq`: Matches values that are equal to a specified value.
*   `$gt` / `$gte`: Matches values that are greater than / greater than or equal to.
*   `$lt` / `$lte`: Matches values that are less than / less than or equal to.
*   `$in`: Matches any of the values specified in an array.
*   `$or`: Joins query clauses with a logical OR.
*   `$and`: Joins query clauses with a logical AND.

### 6. What is the Aggregation Framework? (Basic understanding)
The aggregation framework is a pipeline-based method for processing data records and returning computed results. Operations are grouped into stages.
Common stages:
*   `$match`: Filters documents (like `find`).
*   `$group`: Groups documents by a specified key and performs accumulations (e.g., `$sum`, `$avg`).
*   `$sort`: Sorts documents.
*   `$project`: Reshapes documents (includes/excludes fields).

## Mongoose (Node.js ODM)

### 7. What is Mongoose and why use it?
Mongoose is an Object Data Modeling (ODM) library for MongoDB and Node.js.
While raw MongoDB is schema-less, Mongoose allows you to define rigid schemas at the application layer. This provides:
*   Data validation (before saving to the DB).
*   Enforced structure.
*   Middleware/Hooks (e.g., hashing a password before saving).
*   Easier syntax for writing complex queries.

### 8. What is the difference between a Schema and a Model in Mongoose?
*   **Schema**: A blueprint that defines the structure of your document, default values, and validation rules.
*   **Model**: A compiled version of the Schema. The Model provides the interface (the static methods like `.find()`, `.create()`) to interact with the underlying MongoDB collection.

### 9. How do you implement references (relationships) in Mongoose?
Since MongoDB is NoSQL, "JOINs" are handled via the `$lookup` aggregation or Mongoose's `populate()` method.
You define references in the schema using the `ObjectId` type and a `ref` property.
```javascript
const postSchema = new mongoose.Schema({
    title: String,
    authorId: { type: mongoose.Schema.Types.ObjectId, ref: 'User' }
});

// Later, to retrieve the post WITH the user data:
Post.find().populate('authorId').exec();
```
