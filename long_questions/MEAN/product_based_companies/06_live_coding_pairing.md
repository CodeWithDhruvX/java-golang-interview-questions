# Live Coding & Pairing Scenarios (Product-Based Companies)

In product-based company interviews, you will often face 45-60 minute pairing sessions or live coding exercises. These aren't just algorithmic questions (Leetcode); they are practical, real-world feature implementations evaluating how you write clean, modular, and performant code.

## Scenario 1: Developing a Custom React Hook / Angular Directive

### The prompt:
"Write a custom React Hook (or Angular Directive) called `useDebounce`/`DebounceDirective` that delays invoking a function until after a specified wait time has elapsed since the last time the debounced function was invoked."

### Why they ask it:
Tests your understanding of closures, timers, React component lifecycles (`useEffect` cleanup), and controlling function execution frequency (a core optimization skill).

### React Implementation Focus:
A strong candidate will immediately recognize that `useEffect` needs a cleanup function (`clearTimeout`) to prevent memory leaks and cancel pending state updates if the component unmounts or the value changes rapidly.

```javascript
import { useState, useEffect } from 'react';

// Live Coding Example
export function useDebounce(value, delay) {
    const [debouncedValue, setDebouncedValue] = useState(value);

    useEffect(() => {
        // Set up the timeout
        const handler = setTimeout(() => {
            setDebouncedValue(value);
        }, delay);

        // Crucial: The cleanup function. Clears the timeout if 'value' or 'delay'
        // changes before the timeout completes, thus resetting the "clock".
        return () => {
            clearTimeout(handler);
        };
    }, [value, delay]); // Only re-call if value or delay changes

    return debouncedValue;
}
```

## Scenario 2: Building an API Rate Limiter Middleware in Node.js

### The prompt:
"Implement an Express middleware that limits users to making a maximum of 100 requests per hour per IP address. If they exceed this, return a `429 Too Many Requests` status."

### Why they ask it:
Tests your knowledge of Express middleware patterns, handling request objects (getting the IP), state management, and recognizing scale limitations.

### Implementation Focus:
*   **Level 1 (Basic, usually acceptable for first draft)**: Storing IP addresses and counts in a Javascript `Map` object in MEMORY.
*   **Level 2 (The Follow-up)**: The interviewer asks, "What happens if we scale to 10 Node.js servers behind a load balancer?" The candidate must identify that an in-memory `Map` won't work across multiple servers.
*   **Level 3 (Senior solution)**: Live coding the integration of a Redis client to store the IP rates, utilizing Redis `INCR` and `EXPIRE` commands for an atomic, distributed rate limiter.

```javascript
// Level 1 In-Memory implementation (Live Coding)
const rateLimitMap = new Map();

const rateLimiter = (req, res, next) => {
    const ip = req.ip;
    const windowMs = 60 * 60 * 1000; // 1 hour
    const limit = 100;

    if (!rateLimitMap.has(ip)) {
        rateLimitMap.set(ip, { count: 1, timer: setTimeout(() => rateLimitMap.delete(ip), windowMs) });
        return next();
    }

    const userData = rateLimitMap.get(ip);
    if (userData.count >= limit) {
        return res.status(429).json({ error: 'Too many requests' });
    }

    userData.count++;
    next();
};
```

## Scenario 3: Aggregation Pipeline Debugging/Construction

### The prompt:
You are given a MongoDB dataset containing a messy `Orders` collection. You are asked to write an aggregation pipeline that finds the top 3 selling product categories by total revenue within the last 30 days.

### Why they ask it:
Evaluates whether you know how to process data efficiently on the database side rather than pulling massive datasets into Node.js memory.

### Implementation Focus:
Using `$match` first (to filter the 30-day window early, utilizing indexes), then `$unwind` (if order items are in arrays), `$group` (to calculate sums), `$sort`, and `$limit`.

```javascript
// Live coding expectation:
db.orders.aggregate([
    // 1. Match ONLY orders from the last 30 days immediately (Fast filtering)
    { $match: { createdAt: { $gte: new Date(new Date().setDate(new Date().getDate()-30)) } } },
    // 2. Unwind the items array so we can group by category
    { $unwind: "$items" },
    // 3. Group by category and sum the price * quantity
    {
        $group: {
            _id: "$items.category",
            totalRevenue: { $sum: { $multiply: ["$items.price", "$items.quantity"] } }
        }
    },
    // 4. Sort descending by revenue
    { $sort: { totalRevenue: -1 } },
    // 5. Limit to top 3
    { $limit: 3 }
])
```

## Key Tips for Live Coding in MEAN/MERN:
1.  **Talk Out Loud**: Silence is deadly. Explain your thought process. If you are stuck, communicate exactly *why* you are stuck.
2.  **Make it Work, Then Make it Fast**: Do not over-optimize early. Write the naive solution (e.g., in-memory map), ensure it passes the basic tests, and *then* discuss how you would refactor it for scale (e.g., Redis).
3.  **Handle Errors Proactively**: Before the interviewer asks, add `try/catch` blocks in your Async/Await functions and `if (!data) return res.status(404)` checks. It shows senior-level foresight.
