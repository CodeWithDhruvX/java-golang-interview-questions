# 🗣️ Theory — State Management in React
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "When should you use Redux vs Context vs Zustand?"

> *"The short answer is: Context for simple, stable data — like theme or current user — accessed by many components. Redux for large applications with complex state logic, many developers, or when you need strong conventions and DevTools time-travel debugging. Zustand for everything in between — it's Redux without the boilerplate. The key insight is that Context is NOT a state manager — it's a data transport mechanism. It doesn't have selectors, derived state, or efficient subscriptions. Every consumer re-renders on any context change. If you have state that changes frequently or complex interdependencies, you want a real state manager. Zustand fixes Context's performance issues with a subscription model. Redux is worth it when you have a large team that benefits from enforced patterns and when you want RTK Query for server state as well."*

---

## Q: "What is Redux Toolkit and how does it fix the old Redux pain points?"

> *"Old Redux was famously boilerplate-heavy — you needed action type constants, action creators, reducer switch statements, and connect HOCs, all spread across multiple files. Redux Toolkit — RTK — is the official modern Redux, and it eliminates almost all of that. The centerpiece is createSlice — you give it a name, initialState, and a reducers object, and RTK generates the action creators and action types for you. It also ships with Immer built in, which means you can write 'mutating' state updates that are actually immutable under the hood. There's createAsyncThunk for async operations — it auto-dispatches pending, fulfilled, and rejected actions. And RTK Query for data fetching — similar to React Query but tightly integrated with the Redux store. The mental model shift is: you still think in reducers and actions, but you write a fraction of the code."*

---

## Q: "What is derived state and why is redundant state a problem?"

> *"Derived state is data that can be computed from existing state — it shouldn't be stored separately. If you have a list of products and a filter value in state, the filtered list is derived — you compute it from those two values rather than storing it as a third state variable. The problem with redundant state is synchronization — you now have two sources of truth that need to stay in sync. When the product list changes, you have to remember to also update the filtered list. Miss that update and you have a bug. With derived state, you compute it during render — or with useMemo for expensive computations — and it's always correct because it's always freshly computed from the authoritative sources. Redux's createSelector and React's useMemo are the tools for memoizing derived state so you don't recompute it on every render."*

---

## Q: "Explain server state vs client state."

> *"Client state is UI state that lives only in the browser — the current modal being open, which tab is active, form input values. This belongs in useState or a client state manager like Zustand. Server state is data that actually lives on the server and is just cached on the client — user profiles, product lists, orders. This is fundamentally different: it can become stale, it needs to be synchronized with the server, it requires loading and error tracking, and you might have the same data fetched by multiple components. The problem is that most people reach for useState plus useEffect to manage server state, and then they end up building a poor version of a data fetching library with no caching, no deduplication, and manual loading state. React Query — TanStack Query — exists specifically for server state and handles all of this automatically."*

---

## Q: "How do Zustand subscriptions prevent unnecessary re-renders?"

> *"In Context, all consumers re-render when any part of the context value changes — there's no way to subscribe to just a part of it. Zustand uses a different model: each useStore call takes a selector function — a function that picks a slice of state. Zustand's internal subscription tracks that selector. When state changes, Zustand runs the selector on the new state and the old state and compares the results. If the results are equal — reference equality by default — the component doesn't re-render. So if you have a store with user, cart, and theme, and only the cart updates, only the components that subscribed to cart's slice re-render. Components subscribed to user and theme stay unchanged. This is the key performance advantage over Context, and it's why Zustand is preferred for stores with multiple pieces of frequently updating state."*

---

## Q: "What is optimistic updating and how do you implement it?"

> *"Optimistic updating means updating the UI immediately when a user action happens, without waiting for the server to confirm it. You assume the operation will succeed, show the result right away, and revert if the server responds with an error. It makes apps feel instant. The implementation with React Query: on the mutation's onMutate callback, you cancel outgoing queries to prevent them overwriting your optimistic update, then manually set the cache to the expected result. You save the old data in the context for rollback. If onError is triggered, you restore the old data from context and trigger a refetch. If onSuccess, you invalidate the query so it refreshes from the server to get any server-computed fields. This pattern is common for todo lists, 'like' buttons, cart operations — anything where the success path is highly predictable and the latency would otherwise hurt UX."*

---

## Q: "When would you choose Jotai or Recoil over Redux?"

> *"Both Jotai and Recoil are atomic state libraries — state is split into minimal atoms. Components subscribe to precisely the atoms they need, so only they re-render when that atom changes. This is more granular than Redux where a selector subscribes to a slice of the store. I'd choose atomic state when: you have a complex web of inter-related state where different components need different subsets, like a spreadsheet or a rich text editor. When you need async derived state — Jotai's async atoms are elegant. Or when you want fine-grained subscriptions without the Redux mental model overhead. Between Jotai and Recoil, I'd generally pick Jotai today — it's simpler, has no key strings, is more actively maintained, and has a smaller bundle. Recoil is Meta-backed but its development velocity has slowed. The main reason to choose Redux over atomic: team size, established conventions, and RTK Query."*
