# Angular Cheatsheet Q&A

**User:** Dhruv Shah (dhruvshahlinkedin@gmail.com)  
**Created:** 2/11/2026 16:08:57  
**Updated:** 2/11/2026 16:14:30  
**Exported:** 2/11/2026 16:14:48  
**Link:** [https://chatgpt.com/c/698c5c37-addc-83a7-8686-590d11903d42](https://chatgpt.com/c/698c5c37-addc-83a7-8686-590d11903d42)  

## Prompt:
angular and javascript cheatsheet that most service based comapnies asked

2/2

## Response:
Perfect 👍 Dhruv — since you’re targeting **service-based companies (Infosys, TCS, Wipro, Accenture, Capgemini, Cognizant, etc.)**, they usually focus on:

*   Strong fundamentals
*   Real project understanding
*   Debugging knowledge
*   Practical scenarios
*   Not very deep framework internals like product companies

So here’s a **high-impact Angular + JavaScript Interview Cheatsheet** tailored exactly for service-based companies.

* * *

🔥 JAVASCRIPT CHEATSHEET (Service Company Focus)
================================================

1️⃣ Core Concepts (Must Be Very Strong)
---------------------------------------

### ✅ var vs let vs const

*   `var` → function scoped, hoisted, can redeclare
*   `let` → block scoped, no redeclare
*   `const` → block scoped, cannot reassign

* * *

### ✅ Hoisting

*   Variables declared with `var` are hoisted (initialized as undefined)
*   Functions are fully hoisted

* * *

### ✅ == vs ===

*   `==` → type coercion
*   `===` → strict comparison (recommended)

* * *

### ✅ Data Types

*   Primitive → string, number, boolean, null, undefined, symbol, bigint
*   Non-Primitive → object, array, function

* * *

### ✅ Closures (Very Frequently Asked)

Function remembers its outer scope even after outer function is executed.

```js
function outer() {
  let count = 0;
  return function inner() {
    count++;
    return count;
  }
}
```

Used in:

*   Data hiding
*   Private variables
*   Event handlers

* * *

### ✅ Event Loop (Important)

JS is:

*   Single-threaded
*   Non-blocking
*   Uses call stack + callback queue + event loop

Understand:

*   setTimeout
*   Promises
*   Async/await
*   Microtask vs Macrotask

* * *

### ✅ Promises

States:

*   Pending
*   Fulfilled
*   Rejected

Methods:

*   .then()
*   .catch()
*   .finally()

* * *

### ✅ Async / Await

Cleaner syntax for Promises.

```js
async function getData() {
  const data = await fetch(url);
}
```

* * *

### ✅ Arrow Functions

*   No own `this`
*   Short syntax

* * *

### ✅ this Keyword

Depends on:

*   Global context
*   Object method
*   Arrow function
*   Constructor

* * *

### ✅ Array Important Methods

*   map()
*   filter()
*   reduce()
*   forEach()
*   find()
*   some()
*   every()

* * *

### ✅ Spread & Rest Operator

```js
const arr2 = [...arr1];
function sum(...nums) {}
```

* * *

### ✅ Debouncing & Throttling (Scenario Based)

*   Debounce → API call after user stops typing
*   Throttle → Limit execution frequency

* * *

### ✅ Local Storage vs Session Storage

*   Local → persists
*   Session → tab-based

* * *

* * *

🚀 ANGULAR CHEATSHEET (Service Company Focus)
=============================================

* * *

1️⃣ Angular Architecture
------------------------

*   Component-based
*   Uses TypeScript
*   Follows MVC pattern

* * *

2️⃣ Angular Lifecycle Hooks (Very Important)
--------------------------------------------

Order:

1.  ngOnChanges
2.  ngOnInit
3.  ngDoCheck
4.  ngAfterContentInit
5.  ngAfterViewInit
6.  ngOnDestroy

Most asked:

*   ngOnInit
*   ngOnDestroy
*   ngOnChanges

* * *

3️⃣ Data Binding
----------------

4 Types:

*   Interpolation → `{{}}`
*   Property Binding → `[value]`
*   Event Binding → `(click)`
*   Two-way Binding → `[(ngModel)]`

* * *

4️⃣ Directives
--------------

Types:

*   Structural → \*ngIf, \*ngFor
*   Attribute → ngClass, ngStyle
*   Custom directives

* * *

5️⃣ Services & Dependency Injection (Very Important)
----------------------------------------------------

*   Services → reusable logic
*   DI → inject services into components
*   @Injectable()

* * *

6️⃣ Observables (Very Frequently Asked)
---------------------------------------

*   Used for async data
*   HttpClient returns Observable
*   Need to subscribe()

Important operators:

*   map
*   filter
*   switchMap
*   mergeMap
*   tap
*   catchError

* * *

7️⃣ Angular Routing
-------------------

*   RouterModule
*   Lazy loading
*   Route Guards

Types of guards:

*   CanActivate
*   CanDeactivate
*   Resolve

* * *

8️⃣ Forms
---------

### Template-driven forms

*   Simple forms
*   Uses ngModel

### Reactive forms (More important)

*   FormControl
*   FormGroup
*   Validators

* * *

9️⃣ Pipes
---------

*   Built-in → date, uppercase, currency
*   Custom pipe

* * *

🔟 HTTP Client
--------------

*   GET
*   POST
*   PUT
*   DELETE
*   Interceptors

* * *

1️⃣1️⃣ Change Detection
-----------------------

*   Default
*   OnPush

Service companies love asking:  
👉 When should you use OnPush?

* * *

1️⃣2️⃣ Angular vs React (Comparison Question)
---------------------------------------------

They often ask this casually.

* * *

1️⃣3️⃣ Performance Optimization Questions
-----------------------------------------

*   Lazy loading
*   TrackBy in ngFor
*   Unsubscribe in ngOnDestroy
*   OnPush strategy

* * *

🔥 MOST ASKED SCENARIO QUESTIONS
================================

### ✅ How do you handle memory leaks?

→ Unsubscribe from Observables

* * *

### ✅ How do you share data between components?

*   Parent to child → @Input
*   Child to parent → @Output
*   Unrelated → Shared service

* * *

### ✅ How do you optimize API calls?

*   Debounce
*   switchMap
*   Caching

* * *

### ✅ What is AOT vs JIT?

*   AOT → compile at build time
*   JIT → compile in browser

* * *

🎯 What Service-Based Companies Focus On
========================================

They check:

*   Fundamentals clarity
*   Real project explanation
*   Problem-solving
*   Debugging approach
*   Communication skills

They do NOT usually go extremely deep into:

*   Internal Angular rendering engine
*   Zone.js internals
*   Advanced RxJS internals

* * *

💎 Bonus: How to Answer Smartly
===============================

Instead of giving definition, say:

"In my current project, we used reactive forms because we needed dynamic validation..."

Service companies LOVE project-based answers.

* * *

If you want, I can now give you:

*   🔥 100+ Angular interview rapid-fire questions
*   🔥 100+ JavaScript rapid-fire questions
*   🔥 Real project-based scenario questions
*   🔥 2–3 years experience mock interview
*   🔥 4–6 years experience mock interview

Tell me your experience level and I’ll customize it for you 🚀

## Prompt:
is this enough to crack the interviews?

## Response:
Short answer: **No — this alone is not enough.**  
But… it is a **strong foundation**.

Let me explain clearly, especially for **service-based companies**.

* * *

🎯 What Actually Cracks Service-Based Interviews
================================================

Most candidates fail NOT because they don’t know concepts…

They fail because:

*   ❌ They can’t explain clearly
*   ❌ They can’t relate answers to real projects
*   ❌ They panic in scenario questions
*   ❌ They don’t know debugging approach
*   ❌ Their fundamentals are weak

The cheatsheet I gave you = **Theory Coverage (40%)**

But cracking interview needs:

| Area | Weight |
| --- | --- |
| JavaScript fundamentals | 25% |
| Angular practical knowledge | 25% |
| Real project explanation | 25% |
| Scenario & debugging ability | 15% |
| Communication confidence | 10% |

* * *

🔥 What You Still Need
======================

1️⃣ You must prepare Project-Based Answers
------------------------------------------

For example:

Instead of:

> “Observable is used for async operations.”

Say:

> “In my current project, we used Observables with switchMap to cancel previous API calls in search functionality to avoid unnecessary backend load.”

This level is what impresses interviewers.

* * *

2️⃣ You must prepare Debugging Questions
----------------------------------------

They may ask:

*   What if API is not hitting?
*   Why is ngOnInit not triggering?
*   Why is component not updating?
*   Why memory leak happens?
*   Why is change detection not working?

If you can solve these → You’re ahead of 70% candidates.

* * *

3️⃣ You must know Basic DSA (Light Level)
-----------------------------------------

Service companies may ask:

*   Reverse string
*   Remove duplicates from array
*   Count characters
*   Fibonacci
*   Sorting

Nothing too advanced, but you must be comfortable.

* * *

4️⃣ You must prepare HR + Behavioral
------------------------------------

Very important in service companies:

*   Why do you want to switch?
*   Why should we hire you?
*   Explain your project architecture.
*   What was your biggest challenge?

* * *

💡 Reality Check
================

If you:

*   Just read cheatsheet → ❌ Not enough
*   Practice explaining answers aloud → ✔ Good
*   Prepare 30–40 real scenario questions → ✔ Very good
*   Do 2–3 mock interviews → ✔ Strong chances

* * *

🎯 What I Recommend For You
===========================

Since you're serious about growth and interviews:

### Step 1

Master:

*   Closures
*   Event loop
*   Observables
*   Lifecycle hooks
*   Reactive forms
*   Routing

### Step 2

Prepare:

*   25 scenario-based Angular questions
*   25 scenario-based JavaScript questions

### Step 3

Prepare:

*   One solid project explanation (end-to-end)

* * *

⚡ Honest Answer
===============

For:

*   2–3 years experience → With practice, YES you can crack it.
*   4–6 years → You need deeper architecture + performance + system-level thinking.

* * *

Now tell me:

👉 How many years of experience do you have in Angular/Frontend?

I’ll give you a **custom crack-the-interview roadmap** based on your level.

## Prompt:
please give me that question bank as a cheatsheet

## Response:
Perfect Dhruv 👍  
Here is your **Angular + JavaScript Service-Based Company Interview Question Bank Cheatsheet**.

This is focused on:

*   Infosys
*   TCS
*   Wipro
*   Accenture
*   Capgemini
*   Cognizant
*   Mphasis
*   LTIMindtree
*   Tech Mahindra

These are the type of questions they actually ask.

* * *

🔥 JAVASCRIPT QUESTION BANK (Must Prepare)
==========================================

🧠 Core Fundamentals
--------------------

1.  What is the difference between var, let, and const?
2.  What is hoisting?
3.  What is temporal dead zone?
4.  What is closure? Give a real use case.
5.  What is event loop?
6.  What is call stack?
7.  What are microtasks and macrotasks?
8.  Difference between == and ===?
9.  What is type coercion?
10.  What are primitive and non-primitive types?
11.  What is undefined vs null?
12.  What is NaN?
13.  What is truthy and falsy?
14.  What is a higher-order function?
15.  What is callback hell?

* * *

⚡ Async JavaScript
------------------

16.  What is a Promise?
17.  States of Promise?
18.  What is async/await?
19.  Difference between Promise and async/await?
20.  What happens if you don't use await?
21.  How does setTimeout work internally?
22.  What is fetch API?
23.  How do you handle API errors?
24.  What is Promise.all?
25.  What is Promise.race?

* * *

🔥 this & Functions
-------------------

26.  What is this in global scope?
27.  What is this inside object method?
28.  What is this in arrow function?
29.  What is bind, call, apply?
30.  What is function currying?
31.  What are IIFE?
32.  What is prototype?
33.  What is prototypal inheritance?

* * *

📦 Arrays & Objects
-------------------

34.  Difference between map and forEach?
35.  What does reduce do?
36.  How do you remove duplicates from array?
37.  How do you deep clone an object?
38.  Difference between shallow copy and deep copy?
39.  What is destructuring?
40.  What is spread operator?
41.  What is rest parameter?
42.  How to merge two arrays?
43.  How to check if object is empty?
44.  How to iterate object properties?

* * *

🚀 Scenario-Based JS
--------------------

45.  How do you debounce an API call?
46.  How do you throttle a scroll event?
47.  How do you prevent memory leaks in JS?
48.  What happens if two async calls return in different order?
49.  How do you optimize performance in large list rendering?
50.  How do you debug JavaScript errors?

* * *

🚀 ANGULAR QUESTION BANK
========================

* * *

🏗 Angular Basics
-----------------

1.  What is Angular?
2.  What is SPA?
3.  What is TypeScript?
4.  What is Angular CLI?
5.  Explain Angular architecture.
6.  What is a component?
7.  What is a module?
8.  What is metadata?
9.  What is decorator?
10.  What is bootstrap array?

* * *

🔁 Lifecycle Hooks
------------------

11.  Explain Angular lifecycle hooks.
12.  Difference between ngOnInit and constructor?
13.  When does ngOnChanges trigger?
14.  What is ngOnDestroy used for?
15.  Real use case of ngAfterViewInit?

* * *

🔗 Data Binding
---------------

16.  Types of data binding?
17.  Difference between property and attribute binding?
18.  What is two-way binding?
19.  How does \[(ngModel)\] work internally?

* * *

🧩 Directives
-------------

20.  Difference between structural and attribute directives?
21.  What is \*ngIf?
22.  What is \*ngFor?
23.  What is trackBy and why is it used?
24.  How do you create custom directive?

* * *

💉 Services & DI
----------------

25.  What is dependency injection?
26.  What is @Injectable?
27.  What is providedIn: 'root'?
28.  How do you share data between components?
29.  Singleton service in Angular?

* * *

🌊 Observables & RxJS (Very Important)
--------------------------------------

30.  What is Observable?
31.  Difference between Promise and Observable?
32.  What is subscription?
33.  What happens if you don't unsubscribe?
34.  What is switchMap?
35.  Difference between mergeMap and switchMap?
36.  What is pipe?
37.  What is Subject?
38.  What is BehaviorSubject?
39.  What is ReplaySubject?

* * *

🌐 HTTP & API
-------------

40.  How do you call API in Angular?
41.  What is HttpClient?
42.  How do you handle errors?
43.  What is interceptor?
44.  How do you add token in header?
45.  How do you cancel previous API calls?

* * *

🛣 Routing
----------

46.  What is Angular routing?
47.  What is lazy loading?
48.  What is route guard?
49.  Types of guards?
50.  What is router outlet?

* * *

📝 Forms
--------

51.  Difference between template-driven and reactive forms?
52.  What is FormGroup?
53.  What is FormControl?
54.  How do you add custom validator?
55.  How do you handle dynamic forms?

* * *

⚡ Performance
-------------

56.  What is change detection?
57.  Default vs OnPush?
58.  What is zone.js?
59.  How do you improve Angular performance?
60.  How do you prevent memory leaks?

* * *

🔥 HIGH-PROBABILITY SCENARIO QUESTIONS
======================================

These are gold.

1.  How do you handle large data table?
2.  How do you implement search with debounce?
3.  How do you manage state in Angular?
4.  How do you handle authentication?
5.  How do you protect routes?
6.  How do you implement role-based access?
7.  What happens if API fails?
8.  How do you handle loader globally?
9.  How do you share data between unrelated components?
10.  How do you optimize repeated API calls?

* * *

💎 Must Prepare One Project Explanation
=======================================

They will definitely ask:

*   Explain your current project architecture.
*   What was your role?
*   How many modules?
*   How do you handle API errors?
*   How do you deploy?
*   How do you optimize performance?

If you answer this confidently → 50% interview is done.

* * *

🎯 Final Advice
===============

If you:

*   Prepare these 100+ questions
*   Practice explaining answers aloud
*   Prepare 1 strong project explanation
*   Do 2 mock interviews

You can confidently crack most service-based Angular interviews.

* * *

Now tell me:

How many years of Angular experience do you have?

I’ll give you:

*   🔥 2–3 year level focused questions OR
*   🔥 4–6 year advanced architecture level questions

Customized only for your level.

## Prompt:
is this enough to crack the interviews

## Response:
Dhruv, I’ll give you a very honest answer.

👉 **This question bank is necessary.  
But it is NOT sufficient by itself.**

Let me explain clearly.

* * *

🎯 What Actually Cracks Service-Based Interviews
================================================

Most candidates:

*   ✔ Know definitions
*   ❌ Fail in scenario questions
*   ❌ Cannot explain their project clearly
*   ❌ Panic when interviewer cross-questions

Interviewers don’t check if you memorized 100 questions.

They check:

1.  Can you think?
2.  Can you explain clearly?
3.  Do you understand real-world usage?
4.  Can you debug?

* * *

🧠 The Real Formula to Crack It
===============================

Think like this:

### 1️⃣ Theory (40%)

That cheatsheet covers this.

But theory alone won't select you.

* * *

### 2️⃣ Scenario-Based Thinking (30%)

Example:

❓ “User is typing in search box and API is getting called too many times. What will you do?”

Correct approach:

*   Use debounceTime
*   Use switchMap
*   Cancel previous requests
*   Handle errors

If you answer like this → Strong candidate.

If you say:

> “We can use debounce”

Too basic.

* * *

### 3️⃣ Project Explanation (20%)

They will definitely ask:

*   Explain your current project architecture.
*   How do components communicate?
*   How do you handle authentication?
*   How do you optimize performance?

If you fumble here → interview is gone.

If you explain clearly → you win.

* * *

### 4️⃣ Communication & Confidence (10%)

Service companies value:

*   Clear communication
*   Structured answers
*   Calm behavior
*   Logical thinking

* * *

🚨 Brutal Truth
===============

If you:

*   Just read the cheatsheet → ❌ Not enough
*   Practice answering aloud → ✔ Good
*   Prepare scenario answers → ✔ Strong
*   Prepare project architecture → ✔ Very strong
*   Do mock interview → ✔ High chance

* * *

💡 Let Me Test You
==================

Answer this:

You have a component that calls an API in ngOnInit.  
But user navigates away quickly.  
What problem can happen?  
How will you fix it?

If you can answer this properly → you are interview ready.

* * *

Now tell me:

How many years of Angular experience do you have?

Because preparation strategy for:

*   2–3 years
*   4–6 years

Is completely different.

## Prompt:
Is this enough  to crack the interviews?

## Response:
Dhruv, I’ll answer you very directly.

👉 **If you properly prepare that cheatsheet + practice explaining + prepare your project well… then YES, it is enough to crack most service-based company interviews.**

But…

👉 **If you only read it once and memorize definitions… then NO.**

* * *

🎯 What “Enough” Actually Means
===============================

For service-based companies (Infosys, TCS, Wipro, Accenture, Capgemini, Mphasis etc.), they usually check:

### ✅ Strong fundamentals

*   JS basics (closures, async, array methods)
*   Angular basics (lifecycle, DI, routing, forms)

### ✅ Real project understanding

*   How you used Observables
*   How you handled API errors
*   How you optimized performance
*   How components communicate

### ✅ Scenario thinking

*   Debounce search
*   Cancel API calls
*   Memory leaks
*   Route protection
*   Authentication flow

If you can answer these confidently → you are ahead of 70% candidates.

* * *

🔥 The Real Difference
======================

Average Candidate:

> “Observable is used for async.”

Strong Candidate:

> “In our project, we used switchMap in search functionality to cancel previous API calls and avoid race conditions.”

That practical explanation is what cracks interviews.

* * *

📊 Honest Breakdown
===================

If you:

*   ✔ Understand all JS fundamentals deeply
*   ✔ Understand Angular concepts practically
*   ✔ Prepare 1 strong project explanation
*   ✔ Practice answering aloud
*   ✔ Solve 10–15 basic coding questions

👉 Your chances are HIGH.

* * *

🚨 Where Most People Fail
=========================

*   They don’t know why they used something.
*   They cannot explain architecture.
*   They give one-line answers.
*   They panic in cross-questioning.

* * *

💡 Final Answer
===============

For:

*   2–3 years experience → Yes, this preparation is enough.
*   4–6 years → You need architecture + performance + advanced RxJS knowledge also.

* * *

If you want, I can now give you:

*   🔥 A 7-day crash plan
*   🔥 A 14-day serious preparation plan
*   🔥 A mock interview simulation
*   🔥 Real project explanation template

Tell me:  
How many years of experience do you have?



---
Powered by [ChatGPT Exporter](https://www.chatgptexporter.com)