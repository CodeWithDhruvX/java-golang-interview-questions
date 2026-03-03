# 🗣️ Theory — Testing Basics in React
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "Why do we write tests for React components?"

> *"Tests give you confidence that your code works correctly and catches regressions when you change it. Without tests, every change requires manually clicking through the app to verify nothing broke — which is slow, incomplete, and boring. With good tests, you can refactor a component with confidence. In a professional setting, tests are particularly important because multiple developers work on the same codebase — tests document expected behavior and catch each other's mistakes. The sweet spot for service-based company interviews is demonstrating that you know the basics: setting up RTL, rendering a component, querying elements by role or text, simulating user events, and making assertions. Companies at this level aren't expecting advanced testing concepts like MSW or custom render utilities — they want to see you write a practical, readable test that verifies real behavior."*

---

## Q: "What is React Testing Library and what are its core query methods?"

> *"React Testing Library is the standard tool for testing React components from the user's perspective. You render a component with its render function, and then query the output using methods available on the screen object. The most important queries: getByRole — the preferred method — finds elements by their ARIA role and optional accessible name. This is what a screen reader would see: getByRole('button', { name: 'Submit' }) finds a button with the text Submit. getByLabelText finds form inputs by their associated label text. getByText finds elements by their text content. getByPlaceholderText finds inputs by placeholder. The query naming also tells you about error handling: getBy throws if not found — use when the element must exist. queryBy returns null if not found — use when testing that something is absent. findBy returns a Promise — use for async content that appears after some time. Always prefer higher-priority queries like getByRole over lower-priority ones like getByTestId."*

---

## Q: "How do you simulate user interactions in RTL?"

> *"React Testing Library's companion library user-event simulates real browser interaction more accurately than the old fireEvent. With user-event v14, you first call userEvent.setup() to get a configured user object, then use its methods: user.click() for clicking any element, user.type() for typing text into an input character by character — just like a real user would — user.clear() to clear a field, user.selectOptions() for dropdowns. The important detail about user.type is that it fires all the keyboard events for each character — keydown, keypress, keyup, and input — which triggers onChange handlers correctly for controlled inputs. The old fireEvent.change directly fires the change event without simulating keystrokes — which could produce different behavior than real typing. All user-event methods return Promises — you need to await them in your tests. userEvent.setup() is called outside the test body, and the configured user is shared across assertions within a describe block."*

---

## Q: "What is snapshot testing and when should you use it?"

> *"Snapshot testing captures the rendered output of a component as a text file — the snapshot — and on subsequent test runs, compares the output against the saved snapshot. If the output changed, the test fails. You update the snapshot with a flag when the change is intentional. Snapshot tests are fast to write — one expect(container).toMatchSnapshot() — and they catch accidental visual regressions. The limitations: snapshots of complex components are very large files that are hard to review in code review. When something changes, developers often just update the snapshot without checking if the change was correct — they become a source of false confidence. They test structure, not behavior. My opinion: snapshots are appropriate for small utility components that render markup based on props and have no state or interactions — like an Avatar component. For anything with behavior — forms, lists, interactive components — write explicit interaction tests instead. Interaction tests are harder to write but provide real confidence."*

---

## Q: "What does 'testing implementation details' mean and why is it bad?"

> *"Testing implementation details means writing tests that verify how a component is built internally — which state variable holds a value, which function is called, how many times a child component rendered — rather than what the user experiences. These tests fail when you refactor the implementation even if the behavior is identical. For example: testing that clicking a button calls handleSubmit is testing implementation detail. Testing that clicking the button triggers a network request and shows a success message — that's testing behavior. The problem is fragility: tests should tell you when something is broken from the user's perspective. If a test fails because you renamed an internal function, that test was never providing value — it was just creating maintenance burden. RTL is designed to discourage implementation detail testing — it doesn't give you access to component state or internal functions. It forces you to test through the DOM, the same way the user experiences the component."*

---

## Q: "How do you test a component that shows different UI based on a prop?"

> *"This is the simplest form of component testing and a great starting point for service company interviews. You render the component once with each variant of the prop and assert that the expected UI is shown. For a StatusBadge component that shows different colors and text based on a status prop: render with status='active' and assert the text 'Active' is visible; render again with status='inactive' and assert 'Inactive' is visible. Use separate it blocks for each variant so each test is independent and failures are isolated. A good pattern is multiple renders in a describe block with a clear name. Assert both what should be present and what should be absent — getByText to confirm what's there, queryByText to confirm what's not. Clean up between renders happens automatically — RTL clears the DOM after each test by default, so you can render multiple times in the same test file without interference."*
