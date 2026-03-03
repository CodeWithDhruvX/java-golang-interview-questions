# 🗣️ Theory — Testing React Applications
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What is React Testing Library's philosophy and how does it differ from Enzyme?"

> *"React Testing Library's philosophy is summarized in its guiding principle: test your components the way users use them, not the way they're implemented. Enzyme was the previous generation of React testing — it gave you shallow rendering, access to component state and lifecycle methods, and the ability to call individual functions directly. The problem is you end up testing implementation details. If you refactor a component — same behavior, different internal structure — your Enzyme tests break even though nothing for the user changed. RTL doesn't give you access to component internals. You render the component, interact with it the way a user would — by clicking, typing — and assert on what's visible or accessible in the DOM. Refactoring the implementation doesn't break your tests. This produces tests that give you real confidence because they actually verify user-facing behavior."*

---

## Q: "How does MSW — Mock Service Worker — compare to manually mocking fetch?"

> *"MSW is a service worker that intercepts actual network requests at the browser or Node level — your components make real fetch calls, the service worker intercepts them before they leave the browser, and returns mock responses. Manually mocking fetch with jest.spyOn or global.fetch = jest.fn() works but has drawbacks: you're patching the global, which can bleed between tests; you have to mock at the right level; and it couples your tests to the exact fetch call implementation. MSW lives at the network boundary — your component can use fetch, axios, or anything else, and the MSW handler intercepts regardless. Tests look exactly like production code. The same MSW handlers can be used in development — so you use the same mocks for local dev and tests. If you switch the HTTP layer from fetch to axios, your tests don't change because MSW doesn't care which library made the call."*

---

## Q: "What is the difference between unit, integration, and e2e tests? How do you balance them?"

> *"Unit tests test a single function or component in isolation — dependencies are mocked. Integration tests test how multiple units work together — like a form component with real children, real hooks, but mocked API. End-to-end tests run the full application in a real browser with a real server. The testing trophy — Kent C. Dodds' model for React — suggests most tests should be integration tests, with a smaller number of unit tests for utility functions and a smaller number of e2e tests for critical paths. This differs from the classic testing pyramid which puts most tests as unit tests. The reason: in React, a highly unit-tested component with mocked children tells you little about whether the complete feature works. Integration tests with RTL — rendering real components with real context but mocked network — give much more confidence per test. E2E tests are expensive but irreplaceable for login, checkout, and signup flows."*

---

## Q: "How do you test components that use React Query?"

> *"Components that use React Query need the QueryClientProvider in their test setup. The standard approach is creating a fresh QueryClient per test to avoid state leaking between tests. You wrap the render call in a utility that provides the QueryClient wrapper. For the actual data, you have two options: use MSW to mock the network — which is the most realistic approach since your component uses the real React Query hooks against real mocked endpoints. Or use the queryClient.setQueryData to pre-populate the cache — faster and simpler for straightforward success cases. For testing loading and error states, MSW is essential — you configure the server handler to respond slowly or with an error. Using waitFor or findBy* queries in RTL handles the async nature correctly — RTL will retry the assertion until it passes or timeout occurs."*

---

## Q: "What should you NOT test in React components?"

> *"The things I actively avoid testing: styling and CSS — whether something has the right color or margin is better validated with visual regression tools like Chromatic or Percy, since CSS doesn't belong in JavaScript assertions. Implementation details — which function was called internally, what state variable holds what value, how many times a child component rendered. If the visible outcome is correct, the implementation is correct. Third-party library behavior — I don't test that React State updates correctly when I call setState, or that React Router navigates when I click a Link. Those libraries have their own tests. Over-mocked unit tests — when you mock everything except the component itself, you're testing that the component calls its mocked dependencies correctly, not that the feature works. I also don't test things that can't break — a component that just renders static HTML with no logic doesn't need testing."*

---

## Q: "How do you test that a component handles errors from an API correctly?"

> *"I use MSW to mock the API server returning error responses. First, the test renders the component. Then it waits for the error state to appear in the DOM — using findByText or findByRole with a regex matching an error message. The key assertions are: the error message is visible, any loading indicator is gone, and optionally a retry button is present. For testing the retry behavior, I'd click the retry and then use server.use to override the handler to return success data on the next call. The important detail is using findBy* queries — with the 'find' prefix — which return Promises that RTL retries until the element appears or a timeout occurs. You need the async version because API responses are asynchronous. Using getBy* for async content will fail immediately before the data arrives."*
