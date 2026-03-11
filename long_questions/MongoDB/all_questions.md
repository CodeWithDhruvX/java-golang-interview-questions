Here is a **50+ important keyword cheat sheet** for **Mongoose**. This covers **most keywords developers commonly use in real projects**.

---

# 🧠 Mongoose 50+ Keywords Cheat Sheet

## 1️⃣ Core Setup

Basic keywords used when starting.

| Keyword                 | Purpose            |
| ----------------------- | ------------------ |
| mongoose.connect()      | Connect to MongoDB |
| mongoose.disconnect()   | Disconnect DB      |
| mongoose.Schema         | Create schema      |
| mongoose.model()        | Create model       |
| mongoose.Types.ObjectId | ObjectId type      |

---

# 2️⃣ Schema Data Types

Used to define field types.

| Keyword    | Type                  |
| ---------- | --------------------- |
| String     | Text                  |
| Number     | Numeric value         |
| Boolean    | true/false            |
| Date       | Date values           |
| Buffer     | Binary data           |
| ObjectId   | Reference ID          |
| Array      | List                  |
| Map        | Key-value object      |
| Mixed      | Any type              |
| Decimal128 | High precision number |
| UUID       | Unique identifier     |

Example

```js
const userSchema = new mongoose.Schema({
  name: String,
  age: Number,
  isActive: Boolean
});
```

---

# 3️⃣ Schema Field Options

Used to control behavior of fields.

| Keyword   | Purpose               |
| --------- | --------------------- |
| required  | Field must exist      |
| default   | Default value         |
| unique    | Unique index          |
| lowercase | Convert to lowercase  |
| uppercase | Convert to uppercase  |
| trim      | Remove spaces         |
| minlength | Minimum length        |
| maxlength | Maximum length        |
| enum      | Allowed values        |
| match     | Regex validation      |
| select    | Include/exclude field |
| immutable | Prevent updates       |
| alias     | Alternate field name  |
| index     | Create index          |
| sparse    | Sparse index          |

Example

```js
email:{
 type:String,
 required:true,
 unique:true,
 lowercase:true
}
```

---

# 4️⃣ Schema Options

| Keyword    | Purpose                       |
| ---------- | ----------------------------- |
| timestamps | Adds createdAt & updatedAt    |
| versionKey | Controls __v field            |
| collection | Custom collection name        |
| strict     | Allow/disallow unknown fields |
| minimize   | Remove empty objects          |
| toJSON     | Customize JSON output         |
| toObject   | Customize object output       |

Example

```js
new mongoose.Schema({}, { timestamps:true })
```

---

# 5️⃣ Document Methods

| Method       | Purpose            |
| ------------ | ------------------ |
| save()       | Save document      |
| remove()     | Delete document    |
| validate()   | Validate data      |
| isModified() | Check field change |
| toObject()   | Convert to object  |
| toJSON()     | Convert to JSON    |

---

# 6️⃣ Model Query Methods

| Method                   | Purpose            |
| ------------------------ | ------------------ |
| create()                 | Insert document    |
| find()                   | Find multiple docs |
| findOne()                | Find single doc    |
| findById()               | Find by ID         |
| updateOne()              | Update one         |
| updateMany()             | Update many        |
| deleteOne()              | Delete one         |
| deleteMany()             | Delete many        |
| countDocuments()         | Count docs         |
| estimatedDocumentCount() | Fast count         |

---

# 7️⃣ Advanced Query Methods

| Method              | Purpose                |
| ------------------- | ---------------------- |
| findOneAndUpdate()  | Update + return        |
| findOneAndDelete()  | Delete + return        |
| findByIdAndUpdate() | Update by ID           |
| findByIdAndDelete() | Delete by ID           |
| lean()              | Return plain JS object |
| limit()             | Limit results          |
| skip()              | Skip results           |
| sort()              | Sort results           |
| select()            | Select fields          |

Example

```js
User.find().sort({age:-1}).limit(5)
```

---

# 8️⃣ Reference & Population

| Keyword    | Purpose                  |
| ---------- | ------------------------ |
| ref        | Reference another model  |
| populate() | Load referenced document |

Example

```js
userId:{
 type:mongoose.Schema.Types.ObjectId,
 ref:"User"
}
```

---

# 9️⃣ Middleware Hooks

| Hook      | Purpose           |
| --------- | ----------------- |
| pre()     | Run before event  |
| post()    | Run after event   |
| save      | Before/after save |
| validate  | Validation hook   |
| find      | Query hook        |
| updateOne | Update hook       |
| deleteOne | Delete hook       |

Example

```js
schema.pre("save",function(next){
 console.log("Before saving");
 next();
});
```

---

# 🔟 Advanced Schema Features

| Keyword          | Purpose              |
| ---------------- | -------------------- |
| schema.methods   | Instance methods     |
| schema.statics   | Static methods       |
| schema.virtual() | Virtual fields       |
| schema.query     | Custom query helpers |
| plugin()         | Add plugins          |
| discriminator()  | Schema inheritance   |

---

# ⚡ Most Important 25 (Must Know)

If you're preparing for **interviews or real projects**, focus on these:

```
Schema
model
connect
String
Number
Boolean
Date
ObjectId
required
default
unique
enum
trim
timestamps
save
create
find
findOne
findById
updateOne
deleteOne
populate
ref
pre
post
```

---

✅ **Total Keywords Covered:**
**60+ Mongoose keywords**

---

💡 If you want, I can also show you:

* **Complete Mongoose CRUD project example** (very useful for interviews)
* **Top 30 Mongoose interview questions**
* **Mongoose schema design best practices used in production**.
