# 📘 02 — DOM Manipulation & Browser Events
> **Most Asked in Service-Based Companies** | 🟢 Difficulty: Easy–Medium

---

## 🔑 Must-Know Topics
- DOM selection and manipulation
- Event listeners, bubbling, delegation
- Form validation
- `fetch` API basics
- `localStorage` and `sessionStorage`

---

## ❓ Most Asked Questions

### Q1. How do you select and manipulate DOM elements?

```javascript
// Selection methods
const byId     = document.getElementById('btn');            // fastest for IDs
const firstDiv = document.querySelector('.card');          // CSS selector, first match
const allCards = document.querySelectorAll('.card');       // NodeList of all matches
const byClass  = document.getElementsByClassName('card');  // Live HTMLCollection
const byTag    = document.getElementsByTagName('div');     // Live HTMLCollection

// NodeList vs HTMLCollection:
// NodeList: querySelectorAll returns static snapshot
// HTMLCollection: getElementsBy* returns live (updates with DOM changes)

// Iterating NodeList
allCards.forEach(card => console.log(card.textContent));
// Or convert to array
Array.from(allCards).map(card => card.classList);
[...allCards].filter(card => card.dataset.active === "true");

// Creating and adding elements
const newItem = document.createElement('li');
newItem.textContent = 'New Item';
newItem.className = 'list-item active';
newItem.setAttribute('data-id', '123');

const list = document.getElementById('myList');
list.appendChild(newItem);          // add to end
list.prepend(newItem);              // add to start
list.insertBefore(newItem, list.firstChild); // insert before specific child

// Modern: insertAdjacentHTML (fast, no full re-parse)
list.insertAdjacentHTML('beforeend', '<li class="item">New Item</li>');

// Removing elements
newItem.remove();           // modern API
list.removeChild(newItem);  // older API

// Modifying attributes and styles
const btn = document.querySelector('#submit');
btn.disabled = true;
btn.style.backgroundColor = '#667eea';
btn.classList.add('active');
btn.classList.remove('inactive');
btn.classList.toggle('selected');
btn.classList.contains('selected'); // true/false
```

---

### Q2. Implement event delegation for a dynamic list.

```javascript
// Event delegation: attach ONE listener to parent, handle events from children
// Efficient — works for dynamically added elements

// ❌ Wrong: listeners on each item (breaks for dynamically added items)
document.querySelectorAll('.todo-item').forEach(item => {
    item.addEventListener('click', handleClick); // won't work for new items
});

// ✅ Correct: delegate to parent
const todoList = document.getElementById('todo-list');

todoList.addEventListener('click', (event) => {
    // Check which element was clicked
    const item = event.target.closest('.todo-item');
    if (!item) return;

    if (event.target.matches('.delete-btn')) {
        item.remove();
    } else if (event.target.matches('.checkbox')) {
        item.classList.toggle('completed');
    } else if (event.target.matches('.edit-btn')) {
        startEditMode(item);
    }
});

// Function to add new items (delegation handles clicks automatically)
function addTodoItem(text) {
    todoList.insertAdjacentHTML('beforeend', `
        <li class="todo-item" data-id="${Date.now()}">
            <input type="checkbox" class="checkbox">
            <span class="text">${text}</span>
            <button class="edit-btn">Edit</button>
            <button class="delete-btn">Delete</button>
        </li>
    `);
}

// Form with delegates
const form = document.querySelector('#app');
form.addEventListener('input', (e) => {
    if (e.target.matches('input[type="text"]')) validateText(e.target);
    if (e.target.matches('input[type="email"]')) validateEmail(e.target);
});
```

---

### Q3. How do you handle form validation with JavaScript?

```javascript
// Form validation approaches: HTML5 built-in + Custom JS

// HTML5 built-in attributes:
// required, minlength, maxlength, pattern, type="email", min, max

// Custom validation
const form = document.getElementById('signupForm');

const validators = {
    name: (value) => {
        if (!value.trim()) return "Name is required";
        if (value.trim().length < 2) return "Name must be at least 2 characters";
        return null; // null = valid
    },
    email: (value) => {
        if (!value) return "Email is required";
        if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value)) return "Enter a valid email";
        return null;
    },
    password: (value) => {
        if (!value) return "Password is required";
        if (value.length < 8) return "Password must be at least 8 characters";
        if (!/(?=.*[a-z])(?=.*[A-Z])(?=.*\d)/.test(value))
            return "Must contain uppercase, lowercase, and number";
        return null;
    }
};

function showError(fieldId, message) {
    const field = document.getElementById(fieldId);
    const errorEl = document.getElementById(fieldId + 'Error');
    field.classList.add('error');
    if (errorEl) errorEl.textContent = message;
}

function clearError(fieldId) {
    const field = document.getElementById(fieldId);
    const errorEl = document.getElementById(fieldId + 'Error');
    field.classList.remove('error');
    if (errorEl) errorEl.textContent = '';
}

form.addEventListener('submit', (e) => {
    e.preventDefault(); // prevent page reload
    let isValid = true;

    const formData = new FormData(form);
    for (const [field, validator] of Object.entries(validators)) {
        const value = formData.get(field) || '';
        const error = validator(value);
        if (error) {
            showError(field, error);
            isValid = false;
        } else {
            clearError(field);
        }
    }

    if (isValid) submitForm(Object.fromEntries(formData));
});

// Real-time validation on blur
form.querySelectorAll('input').forEach(input => {
    input.addEventListener('blur', () => {
        const validator = validators[input.name];
        const error = validator?.(input.value);
        error ? showError(input.name, error) : clearError(input.name);
    });
});
```

---

### Q4. Explain the `fetch` API and how to handle errors.

```javascript
// fetch: modern API for making HTTP requests (returns Promise)

// Basic GET request
async function getUsers() {
    try {
        const response = await fetch('https://jsonplaceholder.typicode.com/users');
        
        // fetch only rejects on network errors, NOT HTTP errors!
        if (!response.ok) {
            throw new Error(`HTTP Error: ${response.status} ${response.statusText}`);
        }
        
        const users = await response.json();
        return users;
    } catch (error) {
        console.error('Fetch failed:', error.message);
        throw error; // re-throw so caller knows it failed
    }
}

// POST request with body
async function createUser(userData) {
    const response = await fetch('/api/users', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${getToken()}`
        },
        body: JSON.stringify(userData)
    });

    if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || 'Failed to create user');
    }

    return response.json();
}

// Practical: API wrapper
const api = {
    async request(url, options = {}) {
        const response = await fetch(url, {
            headers: { 'Content-Type': 'application/json' },
            ...options
        });

        if (!response.ok) {
            throw new Error(`API Error ${response.status}`);
        }

        return response.json();
    },

    get(url)          { return this.request(url); },
    post(url, data)   { return this.request(url, { method: 'POST', body: JSON.stringify(data) }); },
    put(url, data)    { return this.request(url, { method: 'PUT', body: JSON.stringify(data) }); },
    delete(url)       { return this.request(url, { method: 'DELETE' }); }
};

// Usage
const users = await api.get('/api/users');
const newUser = await api.post('/api/users', { name: "Alice", email: "alice@example.com" });
```

---

### Q5. How does browser storage work (localStorage vs sessionStorage vs cookies)?

```javascript
// localStorage: persists across sessions until manually cleared (~5MB)
// sessionStorage: cleared when browser/tab is closed (~5MB)
// Cookies: sent to server in every request, configurable expiry (~4KB)

// localStorage operations
localStorage.setItem('theme', 'dark');
const theme = localStorage.getItem('theme'); // 'dark'
localStorage.removeItem('theme');
localStorage.clear(); // removes ALL items

// ✅ Always handle JSON for objects
const user = { id: 1, name: "Alice", role: "admin" };
localStorage.setItem('currentUser', JSON.stringify(user));
const storedUser = JSON.parse(localStorage.getItem('currentUser'));

// Helper utility
const storage = {
    set(key, value) {
        try {
            localStorage.setItem(key, JSON.stringify(value));
        } catch (e) {
            console.error('Storage set failed:', e);
        }
    },
    get(key, defaultValue = null) {
        try {
            const item = localStorage.getItem(key);
            return item ? JSON.parse(item) : defaultValue;
        } catch {
            return defaultValue;
        }
    },
    remove(key) { localStorage.removeItem(key); }
};

// sessionStorage: same API, but tab-scoped
sessionStorage.setItem('formProgress', JSON.stringify({ step: 2 }));

// Cookie operations (manual — no native API)
function setCookie(name, value, days = 7) {
    const expires = new Date(Date.now() + days * 864e5).toUTCString();
    document.cookie = `${name}=${value}; expires=${expires}; path=/; SameSite=Lax`;
}

function getCookie(name) {
    return document.cookie
        .split('; ')
        .find(row => row.startsWith(name + '='))
        ?.split('=')[1];
}

// Comparison: use localStorage for user preferences; sessionStorage for multi-step forms
// cookies for server-readable data (auth tokens with HttpOnly)
```

---

### Q6. How does `debounce` work? Implement it.

```javascript
// Debounce: delay function execution until N ms of idleness
// Use case: search input, window resize, auto-save

function debounce(fn, delay) {
    let timeoutId;
    return function(...args) {
        clearTimeout(timeoutId);                     // reset every time
        timeoutId = setTimeout(() => {
            fn.apply(this, args);                    // call after delay
        }, delay);
    };
}

// Usage: search
const searchInput = document.getElementById('search');

const handleSearch = debounce(async (query) => {
    if (!query.trim()) return;
    const results = await fetchSearchResults(query);
    displayResults(results);
}, 300); // wait 300ms after user stops typing

searchInput.addEventListener('input', (e) => handleSearch(e.target.value));

// Throttle: execute at most once per N ms (for scroll, resize)
function throttle(fn, interval) {
    let lastTime = 0;
    return function(...args) {
        const now = Date.now();
        if (now - lastTime >= interval) {
            lastTime = now;
            fn.apply(this, args);
        }
    };
}

const handleScroll = throttle(() => {
    const scrollPercent = (window.scrollY / document.body.scrollHeight) * 100;
    document.getElementById('progress').style.width = scrollPercent + '%';
}, 100);

window.addEventListener('scroll', handleScroll);

// Key difference:
// Debounce: waits until user STOPS (search, resize, auto-save)
// Throttle: limits frequency (scroll, API calls, game loop)
```

---

### Q7. Explain event propagation with a practical example.

```javascript
// Events propagate in 3 phases:
// Capture: root → parent → target (few listener use this)
// Target: at the actual element
// Bubble: target → parent → root (default, most common)

// Practical example: Closeable Modal
document.addEventListener('DOMContentLoaded', () => {
    const modal = document.getElementById('modal');
    const overlay = document.getElementById('overlay');

    // Close when clicking overlay (outside modal)
    overlay.addEventListener('click', (e) => {
        // Only close if we clicked the overlay itself, not the modal inside it
        if (e.target === overlay) {
            closeModal();
        }
    });

    // Prevent clicks inside modal from bubbling to overlay
    modal.addEventListener('click', (e) => {
        e.stopPropagation();
    });

    // Close on Escape key
    document.addEventListener('keydown', (e) => {
        if (e.key === 'Escape') closeModal();
    });

    function closeModal() {
        overlay.classList.add('hidden');
    }
});

// Event prevent default
const links = document.querySelectorAll('a[href^="#"]');
links.forEach(link => {
    link.addEventListener('click', (e) => {
        e.preventDefault(); // prevent jump to anchor
        const target = document.querySelector(link.getAttribute('href'));
        target?.scrollIntoView({ behavior: 'smooth' });
    });
});

const form = document.querySelector('form');
form.addEventListener('submit', (e) => {
    e.preventDefault(); // prevent page reload
    validateAndSubmit(e.target);
});
```

---

### Q8. What is `setTimeout`, `setInterval`, and how do they behave?

```javascript
// setTimeout: execute function ONCE after delay
const timerId = setTimeout(() => {
    console.log("Executed after 1 second");
}, 1000);

clearTimeout(timerId); // cancel before it fires

// setInterval: execute function REPEATEDLY every N ms
const intervalId = setInterval(() => {
    console.log("Runs every second");
}, 1000);

clearInterval(intervalId); // stop the interval

// ⚠️ setTimeout with 0ms delay: not ZERO — runs after current call stack clears
console.log("1");
setTimeout(() => console.log("2"), 0);
console.log("3");
// Output: 1, 3, 2 — setTimeout goes to task queue!

// Real-world: polling API
function pollForStatus(taskId, intervalMs = 2000) {
    const poll = setInterval(async () => {
        const status = await checkStatus(taskId);
        if (status === 'completed' || status === 'failed') {
            clearInterval(poll); // stop polling
            handleCompletion(status);
        }
    }, intervalMs);

    // Timeout: stop polling after 30s even if not done
    setTimeout(() => {
        clearInterval(poll);
        handleTimeout();
    }, 30000);
}

// ✅ Better: recursive setTimeout (prevents overlap on slow operations)
function pollSafely(taskId) {
    async function check() {
        const status = await checkStatus(taskId);
        if (status !== 'completed') {
            setTimeout(check, 2000); // next call AFTER previous completes
        }
    }
    check();
}
```
