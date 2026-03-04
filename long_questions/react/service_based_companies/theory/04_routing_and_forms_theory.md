# 🗣️ Theory — Routing & Forms
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "How does React Router v6 differ from v5?"

> *"React Router v6 has several important changes from v5. First, the Route component now uses element prop instead of component or render — you pass JSX directly: element=. Second, Routes replaces Switch as the wrapper for multiple Route components — and Routes does automatic best-match selection rather than first-match, so you no longer need to add exact to every route. Third, nested routes are dramatically improved — you can define nested Route elements inside parent routes, and the Outlet component renders the active child route in the parent layout. This eliminates the pattern of manually figuring out nested routing. Fourth, the navigation hook changed from useHistory to useNavigate — and the API is cleaner. Fifth, there's built-in support for relative links. Most importantly, the architecture encourages co-locating route definitions with their component trees rather than one big route config at the top level."*

---

## Q: "How do you implement protected routes in React?"

> *"Protected routes redirect unauthenticated users to the login page. The pattern in React Router v6: create a ProtectedRoute component that checks auth state. If not authenticated, render Navigate to redirect to login. If authenticated, render Outlet to render the matched child route. Then in your router definition, wrap groups of routes that require auth with the ProtectedRoute component as a layout route — the parent Route has no path but has element pointing to ProtectedRoute, and the actual routes are its children. This is cleaner than duplicating an auth check in every protected component. A nice pattern is to save the intended destination using location state when redirecting to login, so after the user logs in you can navigate them directly to where they were trying to go. You read this from location.state?.from in the login component and navigate there on successful login."*

---

## Q: "What is the difference between controlled and uncontrolled forms?"

> *"In a controlled form, React state drives every input field. Each field has a value prop and an onChange handler. Every keystroke updates state, which triggers a re-render, which updates the input. This gives you total control — you can validate on change, transform the input, sync fields. In an uncontrolled form, the DOM owns the input value — you just use defaultValue for the initial value and attach a ref to read the final value when the form submits. React Hook Form is interesting because it's built around uncontrolled inputs by default — it uses refs to read values on submit — which means no re-renders on every keystroke. This makes it much more performant for large forms. Then when validation fails, it can register errors and trigger targeted re-renders for only the invalid fields. You get the performance of uncontrolled with the power of controlled, with much less boilerplate than either."*

---

## Q: "How do you handle form validation properly?"

> *"There are three levels of validation to handle. First — native browser validation — using required, type, min, max attributes. These are zero-effort but limited and browser-styled. Second — manual controlled validation — running a validate function on submit, setting an errors state object, and rendering error messages next to fields. This is sufficient for simple forms. Third — library validation with React Hook Form plus Zod or Yup for schema-based validation. You define a schema that describes the shape and constraints of your form data, and the library automatically validates against it and exposes errors per field. This is the industry standard for serious forms. The important UX patterns: show errors only after a field has been touched — not before — to avoid overwhelming a new user. Show inline errors adjacent to the field, not in a list at the top. Disable the submit button while the form is submitting to prevent double submission."*

---

## Q: "What is the purpose of useNavigate vs Link vs NavLink?"

> *"These serve different use cases. Link is the base navigation component — it renders an anchor tag that navigates to a route without a page reload. Use it anywhere you need a navigable element. NavLink is Link with active state awareness — it knows if its target route is currently active and lets you apply className or style conditionally. It's perfect for navigation menus where you want to highlight the current page. useNavigate is for programmatic navigation — inside event handlers and async code, not in JSX. You'd use it to redirect after a form submission, after a successful login, or when some condition is met. The replace option for navigate is important — navigate('/home', { replace: true }) replaces the current history entry instead of pushing a new one, so the user can't click Back to return to where they were."*

---

## Q: "How do you persist URL state with useSearchParams?"

> *"useSearchParams lets you read and write the URL query string — the part after the question mark — as a React-stateful value. This is valuable because URL state is shareable and bookmarkable — unlike useState. If a user filters a product list and shares the URL, the recipient sees the same filters. The hook returns the params and a setter. Reading a param: searchParams.get('page') — returns the string value or null. Writing: setSearchParams({ category: 'electronics', sort: 'price' }) — replaces the entire query string. For updating just one param without losing others, spread the existing params: setSearchParams(prev => { prev.set('page', '2'); return prev; }). Common patterns: pagination, filters, tabs, search queries — anything where 'what the user is looking at' should be in the URL. The tradeoff is that URL state is always strings — you need to parse numbers and booleans explicitly."*
