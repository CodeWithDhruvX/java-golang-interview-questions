# 📦 JavaScript Arrays, Objects & Scenarios (Detailed Answers)

## 34. Difference between map and forEach?
| Feature | `map` | `forEach` |
| :--- | :--- | :--- |
| **Return Value** | Returns a new array. | Returns `undefined`. |
| **Chaining?** | Yes (`.map().filter()...`). | No. |
| **Immutability** | Does not modify original array. | Does not modify original array (unless elements are objects). |

**Use `map` when you want to transform elements.**
**Use `forEach` when you want to perform side effects (e.g., logging, saving to DB).**

## 35. What does reduce do?
Executes a reducer function on each element of the array, resulting in a **single output value**.
Typical use cases: Summing numbers, flattening arrays, grouping objects.

```javascript
const numbers = [1, 2, 3, 4];
const sum = numbers.reduce((accumulator, currentValue) => accumulator + currentValue, 0);
console.log(sum); // 10
```

## 36. How do you remove duplicates from array?
1.  **Set**: `[...new Set(array)]` (Cleanest, O(N)).
2.  **Filter**: `array.filter((item, index) => array.indexOf(item) === index)` (O(N^2)).

## 37. How do you deep clone an object?
1.  **JSON**: `JSON.parse(JSON.stringify(obj))`
    *   *Pros*: Simple.
    *   *Cons*: Fails with Dates, Functions, undefined, Circular references.
2.  **structuredClone**: `structuredClone(obj)` (Modern, natively supported).
3.  **Recursion**: Write a custom deepClone function.
4.  **Lodash**: `_.cloneDeep(obj)`.

## 38. Difference between shallow copy and deep copy?
*   **Shallow Copy**: Creates a new object, but inserts references into it. If you modify a nested object in the copy, the original is also modified. (`Object.assign`, spread operator `...`).
*   **Deep Copy**: Creates a new object and recursively copies all nested objects. The copy and original share nothing.

## 39. What is destructuring?
Expression that makes it possible to unpack values from arrays, or properties from objects, into distinct variables.

```javascript
const user = { id: 1, name: "Alice" };
const { name } = user; // "Alice"

const colors = ["red", "green"];
const [firstColor] = colors; // "red"
```

## 40. What is spread operator?
(`...`) Allows an iterable (array/string) to be expanded in places where zero or more arguments or elements are expected.

```javascript
const arr1 = [1, 2];
const arr2 = [...arr1, 3, 4]; // [1, 2, 3, 4] (Copy + Add)
```

## 41. What is rest parameter?
(`...args`) Collects all remaining arguments into an array. Must be the last parameter used in a function definition.

```javascript
function sum(...numbers) {
  return numbers.reduce((a, b) => a + b, 0);
}
```

## 42. How to merge two arrays?
*   `const merged = [...arr1, ...arr2];` (Modern)
*   `const merged = arr1.concat(arr2);` (Functions)

## 43. How to check if object is empty?
*   `Object.keys(obj).length === 0` (Most reliable).
*   `JSON.stringify(obj) === "{}"` (Slower).

## 44. How to iterate object properties?
*   `for (let key in obj)`: Iterates over enumerable properties (including inherited ones).
*   `Object.keys(obj)`: Array of own enumerable keys.
*   `Object.entries(obj)`: Array of `[key, value]` pairs.

---

## 45. How do you debounce an API call? (Scenario)
**Problem**: User types fast, triggering 100 API calls.
**Solution**: Wait for user to stop typing for `n` ms before calling API.

```javascript
function debounce(func, delay) {
  let timeoutId;
  return function(...args) {
    if (timeoutId) clearTimeout(timeoutId);
    timeoutId = setTimeout(() => {
        func.apply(this, args);
    }, delay);
  };
}

const search = debounce(() => console.log("API Call"), 300);
// Used on input 'keyup' event
```

## 46. How do you throttle a scroll event? (Scenario)
**Problem**: Scroll event fires too often (performance issue).
**Solution**: Ensure function runs at most once every `n` ms.

```javascript
function throttle(func, limit) {
  let inThrottle;
  return function(...args) {
    if (!inThrottle) {
      func.apply(this, args);
      inThrottle = true;
      setTimeout(() => inThrottle = false, limit);
    }
  }
}
```

## 47. How do you prevent memory leaks in JS?
1.  **Clear Timers**: Always `clearTimeout` / `clearInterval` when component unmounts.
2.  **Remove Listeners**: `element.removeEventListener()` when done.
3.  **Closures**: Be careful with closures holding references to large objects/DOM elements.
4.  **Detached DOM**: Nullify references to DOM nodes removed from the document.

## 48. What happens if two async calls return in different order? (Race Condition)
If first request takes 3s and second takes 1s, the first might overwrite the second's result.
**Fix**:
1.  **Cancel Previous**: AbortController (fetch) or `switchMap` (RxJS).
2.  **Versioning**: Ignore results from older requests.

## 49. How do you optimize performance in large list rendering?
**Virtualization (Windowing)**.
Render only the items currently visible in the viewport. As user scrolls, recycle DOM nodes.
Libraries: `react-window`, `cdk-virtual-scroll-viewport` (Angular).

## 50. How do you debug JavaScript errors?
1.  **Console**: `console.log`, `console.table`, `console.dir`.
2.  **Debugger**: Add `debugger;` line in code to pause execution in DevTools.
3.  **Network Tab**: Check for failed API requests (400/500 errors).
4.  **Breakpoints**: Click on line number in Sources tab to pause and inspect variables.
