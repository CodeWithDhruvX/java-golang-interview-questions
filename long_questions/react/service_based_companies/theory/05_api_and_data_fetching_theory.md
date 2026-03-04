# 🗣️ Theory — API & Data Fetching
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "Why is React Query better than useEffect + useState for data fetching?"

> *"The naive approach — fetch in useEffect, store in useState — looks simple but hides a lot of problems you have to solve yourself. No caching — every time a component mounts, even if you just had that data, it fetches again. No deduplication — if two components fetch the same data simultaneously, you make two network requests. No background refetching — data can go stale without the user knowing. Manual loading and error states in every component. Manual cache invalidation after mutations. React Query solves all of this. You define a queryKey and a queryFn, and React Query handles caching, deduplication, stale-time management, background refetching on window focus, retry on failure, and loading and error states. Components that share the same queryKey share the same cache — zero duplicate requests. When data changes, you call invalidateQueries and React Query refetches. The code is dramatically simpler and more correct than the DIY approach."*

---

## Q: "How do you handle errors from API calls properly in React?"

> *"Error handling should be layered. At the component level: catch the error from useQuery or the fetch call and render contextually appropriate UI — for a search results component, showing 'Search failed, try again' with a retry button makes sense. For unexpected errors in a component tree during render — not async errors — use an Error Boundary. At the API client level: use Axios interceptors or a fetch wrapper to handle auth errors globally — a 401 should redirect to login regardless of which component triggered it. Server errors — 500s — should log to your error tracking service like Sentry. Network errors — no connection — should be distinguished from server errors and shown with an appropriate offline message. React Query handles retries automatically — by default it retries failed requests up to 3 times with exponential backoff. You configure retry behavior per query to avoid retrying user errors like 400s and 404s that won't change on retry."*

---

## Q: "What is the difference between GET and POST requests and when do you use each?"

> *"GET requests are for reading data — they have no body, the parameters go in the URL as query strings, they're idempotent — calling them multiple times produces the same result — and they can be cached by browsers and CDNs. You'd use GET for fetching product lists, user profiles, search results. POST requests are for submitting data — they have a body, don't appear in browser history, and are not idempotent — calling a POST twice might create two records. For REST APIs: POST to create a resource, PUT or PATCH to update, DELETE to remove. In React, useEffect or React Query handles GET fetches. For POST — form submissions, mutations — you use the Axios post method or fetch with method: 'POST' and a JSON body. With React Query, mutations use useMutation. Setting Content-Type: application/json tells the server to parse the body as JSON — without this, your data arrives as a plain string."*

---

## Q: "What is CORS and whose responsibility is it to fix?"

> *"CORS — Cross-Origin Resource Sharing — is a browser security mechanism. By default, JavaScript running on one origin — say localhost:3000 — cannot make requests to a different origin — like localhost:5000 or api.example.com — because the browser blocks them. The server opt-in to allow specific origins by including Access-Control-Allow-Origin headers in its response. This is entirely a server-side configuration — the browser enforces the restriction, but the fix is on the backend. In development, the easiest solution is the Vite or CRA proxy — you configure the dev server to forward requests to your API, so the requests appear to come from the same origin. In production, the API server must include the appropriate CORS headers for the frontend domain. A common misconception: CORS doesn't protect your API from server-to-server requests — only from browsers. It's not a security measure for your API, it's a mechanism that lets your API opt in to cross-origin browser requests."*

---

## Q: "How do you use environment variables and why can't you put secrets in them?"

> *"Environment variables are a way to inject configuration that differs between environments — development versus staging versus production. In React, you prefix them with REACT_APP_ for CRA or VITE_ for Vite projects. You read them in code at build time — process.env.REACT_APP_API_URL or import.meta.env.VITE_API_URL. The critical security point: these values are embedded in your JavaScript bundle at build time and are visible to anyone who downloads your app. The browser is a public client — there are no secrets in the browser. So you should only put public configuration in client-side environment variables: API URLs, your Stripe publishable key, analytics IDs — things that are designed to be public. Private keys, database passwords, signing secrets — those belong only on the server, in server-side environment variables that never reach the browser. If you need to use a secret in a frontend operation, create a backend endpoint that uses the secret and call that endpoint."*
