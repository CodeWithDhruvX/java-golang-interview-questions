# 🗣️ Theory — Hooks & State
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "Explain useState and when you'd use the functional update form."

> *"useState is the hook for local component state — you give it an initial value and it returns a tuple of the current value and a setter function. The functional update form — passing a function to the setter instead of a value — is important for two scenarios. First, when the new state depends on the previous state, especially if multiple updates might be batched. If you call setCount(count + 1) three times in a row, all three reads capture the same stale count. But if you write setCount(prev => prev + 1) three times, each function receives the most recent value from the queue. Second, for expensive initial state — you can pass a function to useState as the initializer, and it's only called once during mount. For example, parsing localStorage. Without the function form, JSON.parse runs on every render — wasteful. With the function form, it runs exactly once."*

---

## Q: "Walk me through the useEffect dependency array — what are the rules?"

> *"The dependency array tells React when to re-run the effect. No array means run after every render. Empty array means run once after mount — like componentDidMount. Array with values means run after the first render and again whenever any dependency changes. The rule is: every reactive value referenced inside the effect must be in the dependency array. Reactive values are props, state, and anything derived from them. If you omit a dependency, the effect captures a stale closure — a snapshot from when it last ran, which may be outdated. The ESLint rule exhaustive-deps enforces this. Common escape hatches: if the effect only needs to run with the initial value and you genuinely don't want it to re-run when something changes, consider whether the dependency belongs in a ref instead. useRef values are deliberately excluded because they don't trigger re-renders and don't form stale closures."*

---

## Q: "What are the two uses of useRef?"

> *"useRef has two distinct uses. First — DOM access. Attaching a ref to a JSX element gives you a direct reference to the underlying DOM node in ref.current. You'd use this for focusing an input programmatically, measuring dimensions with getBoundingClientRect, or integrating with third-party libraries that need a DOM node. Second — mutable instance variable. useRef creates a container that persists across renders. Crucially, changing ref.current doesn't trigger a re-render. This is useful for storing the previous value of a prop or state, holding a timer ID so you can cancel it, storing a flag that tracks if the component is mounted -- for preventing state updates after unmount. The key mental model is: useState for values that affect rendering; useRef for values that affect behavior but not rendering."*

---

## Q: "What is the difference between useEffect and a regular event handler?"

> *"This is a fundamental React question. Event handlers are the right place for user-initiated actions — clicks, form submissions, key presses. Effects are for synchronization with external systems — when some state or prop changes, you need the outside world to stay in sync. If you're thinking 'when the user clicks this button, fetch some data' — that might feel like an effect, but it's actually an event handler. You can call fetch directly in the onClick handler. Effects should answer the question: when THIS VALUE changes — through any means, not just user action — what external thing needs to be synchronized? A WebSocket subscription that should connect when userId changes, a document title that should reflect the current page name, an analytics event that fires when a user lands on a route. The common mistake: putting logic in useEffect just to 'do it after render' when it should be in an event handler."*

---

## Q: "When would you use useReducer over useState?"

> *"useReducer shines over useState in a few specific situations. When you have complex state transitions where the next state depends on the previous in a non-trivial way — especially when multiple fields change together based on the same action. A form with validation is a classic example: submitting sets loading to true, clears errors, and might reset certain fields — that's multiple state updates from one action. With multiple useState calls, you'd call setLoading, setErrors, and setFormData separately. With useReducer, it's one dispatch({ type: 'SUBMIT' }). The second scenario: when the next state depends on both current state and the action in a way that needs to be testable. Reducers are pure functions — you can unit test them without rendering. The rule of thumb: start with useState. When you find yourself writing complex setState logic with multiple dependent updates, or when you want to test state logic in isolation, switch to useReducer."*

---

## Q: "Explain the Rules of Hooks and why they exist."

> *"The Rules of Hooks are two: only call hooks at the top level — not inside conditions, loops, or nested functions — and only call hooks from React functions — functional components or custom hooks. The reason is how React tracks hook state. React maintains a linked list of hook states per component, ordered by the sequence in which hooks are called. On every render, React walks that list in order to match hook calls to their stored state. If you call a hook conditionally, the order of hooks can differ between renders — React reads from the wrong position in the list and state gets scrambled. This is a fundamental architectural constraint of how React implemented hooks without needing to identify them by name. The ESLint plugin eslint-plugin-react-hooks enforces these rules automatically — you should always have this lint rule enabled and never disable it."*
