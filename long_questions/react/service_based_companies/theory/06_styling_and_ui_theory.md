# 🗣️ Theory — Styling & UI in React
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What are the different approaches to styling in React and when do you choose each?"

> *"There are four main styling approaches. Plain CSS files — simple, no tooling overhead, but global scope means class name conflicts in large codebases. CSS Modules — scoped CSS where class names are auto-namespaced per file, preventing collisions. You import the CSS file as an object and use object keys as class names. No JavaScript in your styles, keeps CSS separate. CSS-in-JS — Styled Components, Emotion — lets you write CSS directly in JavaScript as template literals, with full access to props and theme for dynamic styles. Powerful but adds JavaScript bundle overhead and can be harder for CSS-focused developers. Utility-first CSS — Tailwind — you compose utility classes directly in JSX. Extremely fast to write once you know the classes, consistent design system, but JSX can look verbose. My recommendation: CSS Modules for projects without a design system. Tailwind for rapid development. Styled Components or Emotion when you need heavy dynamic styling based on JavaScript state."*

---

## Q: "How do CSS Modules prevent class name collisions?"

> *"CSS Modules work by transforming class names at build time. When you write .button in a CSS module file, the build tool — Vite or webpack — renames it to something like Button_button__kVQ2 — a combination of the component name, class name, and a content hash. This guarantees uniqueness across the entire codebase. In your JavaScript, you import the CSS file as an object — styles — and use styles.button as the className. The transformation is handled automatically by the bundler, so your CSS file looks completely normal. The only rule is the file must be named something.module.css for the bundler to know it's a CSS Module. This approach gives you the familiarity of plain CSS — no JavaScript syntax for styles — with the scoping guarantees of CSS-in-JS. It's also perfectly efficient: all the styles compile to a static stylesheet, there's no JavaScript runtime cost."*

---

## Q: "What is the difference between Styled Components and inline styles?"

> *"Inline styles in React are JavaScript objects passed to the style prop. They're computed at render time and applied directly to the DOM element's style attribute. The limitations: you can't use pseudo-classes like hover, focus, active. No media queries. No CSS animations. They mix presentation logic with component logic. Styled Components are actual CSS — you write CSS syntax in a tagged template literal, it gets extracted to a real stylesheet, and a unique class is applied to your component. You get everything CSS offers: pseudo-classes, media queries, keyframe animations, child selectors. You also get dynamic styles through props — the styled component has access to its props in the template literal. The practical rule: use inline styles only for truly dynamic values that must be computed in JavaScript — like a progress bar width from a percentage variable. For everything else, use CSS Modules or Styled Components for the full power of CSS."*

---

## Q: "How do you implement dark mode in a React app?"

> *"Dark mode has two concerns: detecting the user's preference and toggling between themes. For detection, the prefers-color-scheme media query tells you whether the OS is in dark mode — you can read this with window.matchMedia in JavaScript or apply it directly in CSS. For the implementation: I use CSS variables as the theming mechanism. Define color variables on :root for light mode, and override them on [data-theme='dark'] or the .dark class. The variables cascade to all elements automatically. In React, I store the current theme in state — initialized from localStorage for persistence across sessions, with a fallback to the OS preference. An effect syncs the state to a data attribute on the document element and saves to localStorage. A ThemeContext exposes the current theme and a toggle function. Any component that needs theme-aware styles just uses the CSS variables — they automatically reflect the current theme without any component re-renders."*

---

## Q: "What is Tailwind CSS and how do you handle dynamic classes with it?"

> *"Tailwind is a utility-first CSS framework — instead of writing custom CSS, you compose pre-built utility classes directly in your JSX. Classes like p-4, text-lg, bg-blue-500, flex, border-radius correspond directly to single CSS declarations. It's extremely fast for development once you learn the class names, and because Tailwind purges unused classes in production, the stylesheet is tiny. The dynamic class challenge: Tailwind's purging works by scanning your source files for complete class names. If you construct class names by concatenation — 'bg-' + color — the purger doesn't see the complete class string and removes it. The solution: always use complete class names, and use a mapping object for variants. For conditional classes, use the clsx or classnames library to cleanly compose class names without string concatenation. The Tailwind docs explicitly say: never dynamically construct class strings — always ensure complete class names exist somewhere in your source."*
