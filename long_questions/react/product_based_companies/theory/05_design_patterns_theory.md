# 🗣️ Theory — Design Patterns in React
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What is the Compound Components pattern and when would you use it?"

> *"Compound Components is a pattern where a parent component and several child components work together sharing implicit state — similar to how the native select and option HTML elements work together. The parent owns the state and exposes it via context. Each sub-component reads from that context to know how to render itself. The beauty is the API it produces: the consumer gets a declarative, readable DSL-like interface — Tabs > Tabs.List > Tabs.Tab — instead of passing twenty different configuration props to a single monolithic component. The alternative — 'prop soup' — would be one component with activeTab, tabList, onTabChange, renderPanel, and so on. Compound Components shine when building component libraries: Accordion, Tabs, Select, Dropdown, Modal with multiple sections. It separates the behavior control — in the parent — from the rendering decisions — in the consumer."*

---

## Q: "How are custom hooks the evolution of higher-order components?"

> *"Both HOCs and custom hooks solve the same problem: reusing stateful logic across multiple components. HOCs do it by wrapping a component and injecting props — you get a new component back. Custom hooks do it by extracting the logic into a reusable function that any component can call. The problems HOCs solved but poorly: HOC wrappers create extra layers in the DevTools component tree. Multiple HOCs can collide on prop names. They're hard to type generically in TypeScript. Custom hooks fix all of these. A hook like useAuth() or useWindowSize() adds no extra components to the tree — the logic lives inside the component that calls it. There's no prop name collision because each hook has its own scope. They're straightforward to type. The only scenario where an HOC is still cleaner than a hook is when you need to wrap a component you don't own — a third-party component — because you can't add hook calls to it."*

---

## Q: "What is the Provider pattern and how do you structure it?"

> *"The Provider pattern is about creating a component that wraps a subtree and makes functionality — state, actions, configuration — available to any descendant without explicit prop passing. You implement it with React Context: create a context, create a Provider component that holds the state and exposes it via the context value, and create a custom hook — useMyContext — that components use to access it. The custom hook serves two purposes: it's a convenient accessor, and it throws a helpful error if called outside the Provider — catching developer mistakes early. For complex contexts, you often split the context into a stable part — state that doesn't change often — and a dynamic part — frequently updating values — to prevent over-rendering. In large apps, you'll have nested providers for different concerns: AuthProvider at the top, then ThemeProvider, then CartProvider — each providing only what its subtree needs."*

---

## Q: "What is the Container/Presentational pattern and is it still relevant?"

> *"The Container/Presentational pattern separates components by responsibility. Containers do logic — they fetch data, manage state, handle business rules — and they have no direct styling. Presentational or 'dumb' components just render whatever props they receive — they're pure UI with no awareness of where the data comes from. This separation has real benefits: presentational components are easy to test with varying prop combinations, easy to document in Storybook, and trivially reusable. Containers can change their data source without touching the UI. Is it still relevant? Yes and no. Hooks somewhat blurred the line — now you can extract logic into a custom hook inside a component without making a separate container component. But the underlying principle — separate logic from presentation — is very much still good design. The pattern is now often implemented as components using custom hooks rather than splitting into two separate component files."*

---

## Q: "What is a headless component and why is it powerful?"

> *"A headless component provides behavior and accessibility without any visual styling. It encapsulates the hard parts — keyboard navigation, ARIA attributes, focus management, state — and gives the consumer complete control over how it looks. Radix UI is the best example: their Dialog component handles focus trapping, ESC to close, ARIA roles, scroll locking, and animation primitives — but applies zero CSS. You bring your own styles through className. This is powerful because it solves the hardest problem in UI libraries — the styling escape hatch problem. Traditional component libraries like Bootstrap bake in design decisions that are hard to override. Headless components have no opinions on appearance. They're especially relevant in enterprise apps where every company has a specific design system that conflicts with third-party styles. You get the heavily tested, accessible behavior for free while keeping full design flexibility."*

---

## Q: "When do you use render props versus a custom hook?"

> *"Both patterns share logic with other components — the question is when each is more appropriate. Render props — passing a function as a prop that the component calls to render something using its internal state — were the solution before hooks. The pattern has a specific remaining use case: when you need to co-locate rendered output from a shared logic component with its consumer's markup. For example, a DragAndDrop container that needs to coordinate exactly which element is being dragged and render a preview. Custom hooks are better for almost all other cases — pure logic extraction without UI. The hook calls in any component that needs the logic; no extra component in the tree; no callback prop ceremony. If you look at a render prop and none of the render prop's arguments affect what you render — you're just using it for side effects or computed values — that's a sign it should be a custom hook."*
