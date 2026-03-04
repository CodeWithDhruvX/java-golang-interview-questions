# 🧪 07 — Testing & Debugging
> **Most Asked in Service-Based Companies** | 🧪 Difficulty: Medium

---

## 🔑 Must-Know Topics
- Unit Testing vs Integration Testing
- Popular Testing Frameworks (Jest, Mocha)
- Mocking and Stubbing
- Testing Asynchronous Code
- Debugging Node.js applications

---

## ❓ Frequently Asked Questions

### Q1. What is the difference between Unit Testing and Integration Testing?

**Answer:**

- **Unit Testing:** Tests individual, isolated pieces of code (like a single function or class method). The goal is to ensure that a specific component works as intended in isolation. Dependencies (like databases or external APIs) are usually mocked or stubbed.
- **Integration Testing:** Tests how multiple units or components work together. For instance, testing an Express route that connects to a database, creates a user, and returns a JSON response. It ensures the whole system or a subsystem functions correctly.

---

### Q2. Write a simple Unit Test using Jest.

**Answer:**
Jest is a popular testing framework by Facebook. Let's write a simple math function and test it.

*math.js*
```javascript
const add = (a, b) => a + b;
module.exports = { add };
```

*math.test.js*
```javascript
const { add } = require('./math');

// Describe block groups related tests
describe('Math Functions', () => {

    // Test or 'it' defines a single test case
    test('should add two positive numbers correctly', () => {
        const result = add(2, 3);
        // expect() is an assertion
        expect(result).toBe(5); 
    });
});
```
To run it, configure `"test": "jest"` in `package.json` and run `npm test`.

---

### Q3. How do you test asynchronous code in Jest?

**Answer:**
When testing asynchronous code, you must ensure the test runner waits for the asynchronous operation to complete before asserting the result.

**Using Promises:**
Return the Promise from the test.
```javascript
test('fetches data successfully', () => {
    return fetchData().then(data => {
        expect(data).toBe('peanut butter');
    });
});
```

**Using Async/Await (Preferred):**
Mark the test function as `async`.
```javascript
test('fetches data successfully', async () => {
    const data = await fetchData();
    expect(data).toBe('peanut butter');
});

test('handles errors', async () => {
    expect.assertions(1); // Ensure the catch block is hit
    try {
        await fetchError();
    } catch (e) {
        expect(e).toMatch('error');
    }
});
```

---

### Q4. What is Mocking and why is it used? Can you give an example?

**Answer:**
Mocking means replacing a real object, function, or module with a fake one (a "mock") during testing. 

**Why use it?**
- To isolate the unit being tested from external dependencies (databases, third-party APIs).
- To make tests faster and more reliable (no network latency).
- To simulate error scenarios (like a database connection failure).

**Example in Jest:**
Suppose we want to test a function that fetches user data from an API, but we don't want to make an actual HTTP request during the test.
```javascript
// Function being tested
const axios = require('axios');
async function getUser(id) {
    const response = await axios.get(`https://api.example.com/users/${id}`);
    return response.data;
}

// Test file
jest.mock('axios'); // Mock the entire axios library

test('should fetch a user', async () => {
    const testData = { name: 'Alice' };
    
    // Define what the mocked axios.get should resolve to
    axios.get.mockResolvedValue({ data: testData }); 
    
    const user = await getUser(1);
    
    expect(user).toEqual(testData);
    expect(axios.get).toHaveBeenCalledWith('https://api.example.com/users/1');
});
```

---

### Q5. How do you debug a Node.js application?

**Answer:**
There are several ways to debug in Node.js:

1. **`console.log()`:** The simplest and most common method for quick debugging, though not suitable for complex data flows.
2. **Built-in Node Inspector:** You can run Node.js with the `--inspect` flag:
   ```bash
   node --inspect app.js
   # Or --inspect-brk to break on the first line
   ```
   You can then open `chrome://inspect` in Google Chrome and use the Chrome DevTools to set breakpoints, step through code, and inspect variables.
3. **IDE Debuggers:** Modern IDEs like Visual Studio Code have powerful built-in debuggers for Node.js. You can configure `launch.json` to attach to running processes or start the app in debug mode.
