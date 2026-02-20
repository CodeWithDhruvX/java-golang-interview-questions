# 🟢 **Freshers: HTML & CSS Basics**

### 1. What is the difference between HTML and HTML5?
"HTML5 is the latest version of HTML, and it brought massive improvements for modern web development. The biggest difference I notice is the introduction of **semantic tags** like `<article>`, `<section>`, `<header>`, and `<footer>`, which replaced the endless sea of `<div>` tags we used to rely on. 

It also introduced native multimedia support with `<audio>` and `<video>` tags, eliminating the need for bulky third-party plugins like Flash. Plus, it gave us APIs for **local storage**, which is incredibly useful for saving user preferences without constantly hitting the backend."

#### Indepth
HTML5 also introduced the `<canvas>` element for dynamic 2D and 3D graphics rendering via JavaScript. Furthermore, HTML5 has stricter parsing rules, which means modern browsers can better handle malformed syntax without breaking the page layout compared to HTML4.

---

### 2. What are semantic HTML elements and why do they matter?
"Semantic elements are tags that clearly describe their meaning to both the browser and the developer. Examples include `<article>`, `<aside>`, `<nav>`, and `<main>`. 

I use them because they provide immense value for **Accessibility (a11y)**, allowing screen readers to interpret the page structure correctly for visually impaired users. They also boost **SEO** (Search Engine Optimization) because search engine crawlers can easily distinguish between main content and supplementary content. Plus, they just make the code so much cleaner to read."

#### Indepth
While a `<div>` implies a generic container without intrinsic meaning, semantic tags convey intent. For example, screen readers will provide navigation commands specifically for `<nav>` sections, allowing users to jump directly to site navigation instead of tabbing through every element on the page.

---

### 3. What is the difference between `id` and `class`?
"Both are attributes used to target elements for styling or JavaScript manipulation, but their scoping is different. 

An `id` must be **unique** within a single HTML document. I use it for one-of-a-kind elements, like a main navigation bar or a specific target for a URL anchor. A `class`, on the other hand, can be applied to **multiple** elements. I use classes for reusable styles, like `.btn-primary` or `.card`.

Also, in CSS, an `id` selector (`#`) has higher specificity than a `class` selector (`.`)."

#### Indepth
Because IDs must be strictly unique, relying heavily on them for CSS styling is considered an anti-pattern in modern architectures (like BEM) because it causes specificity wars. It is best practice to use classes for styling and reserve IDs for JavaScript hooks, form labels (`for` attribute), or fragment identifiers in URLs.

---

### 4. What are inline, block, and inline-block elements?
"These define how elements behave in the document flow.

A **Block-level** element (like `<div>` or `<p>`) always starts on a new line and takes up the full width available to it. 
An **Inline** element (like `<span>` or `<a>`) only takes up as much width as its content and does not start on a new line. I can't set explicit width or height on them.
An **Inline-block** element is a hybrid. It sits inline with text, but unlike a pure inline element, it respects the `width` and `height` properties I assign to it."

#### Indepth
The default display behavior can be altered using CSS (`display: block | inline | inline-block`). Note that block styling also applies vertical margins correctly, whereas applying vertical margins to true inline elements will not affect surrounding elements in the normal document flow.

---

### 5. What is the Box Model?
"The CSS Box Model is the foundation of web layout. Every element on a web page is essentially a rectangular box consisting of four layers.

From the inside out:
1. **Content**: The actual text, image, or child element.
2. **Padding**: Transparent space *inside* the border, around the content.
3. **Border**: The line surrounding the padding and content.
4. **Margin**: Transparent space *outside* the border, pushing other elements away.

Understanding this is crucial because, by default, padding and borders increase the total rendered size of an element beyond its defined `width` and `height`."

#### Indepth
To avoid calculating complex widths manually, modern CSS almost always uses `box-sizing: border-box`. This tells the browser to include padding and border within the explicitly declared `width` and `height`, preventing elements from unexpectedly overflowing their containers.

---

### 6. What is the difference between `margin` and `padding`?
"`Margin` is the space *outside* an element's border. I use it to create distance between different elements on the page, like separating two cards.

`Padding` is the space *inside* an element's border, between the border and the content itself. I use it to give the content some breathing room, like making a button larger and easier to click without changing the text size."

#### Indepth
One crucial behavior regarding margins is **Margin Collapse**. When vertical margins of two adjoining block elements touch, they combine into a single margin equal to the larger of the two, rather than summing together. Padding never collapses. Furthermore, backgrounds apply to the padding area, but not the margin area.

---

### 7. What is the difference between `display: none` and `visibility: hidden`?
"Both hide an element from view, but they impact the document layout differently.

When I use `display: none`, the element is completely removed from the document flow. It’s as if the element doesn't exist, and surrounding elements will collapse into the space it previously occupied.

When I use `visibility: hidden`, the element is invisible, but it still takes up its exact original space in the layout, leaving an empty 'hole' on the page."

#### Indepth
Neither state allows for user interaction (e.g., clicks). However, `opacity: 0` is another alternative where the element is invisible but still takes up space *and* can still receive pointer events (clicks). Additionally, elements with `display: none` are completely ignored by screen readers, whereas `visibility: hidden` hides content from both sighted users and screen readers.

---

### 8. Explain the different positioning types (`static`, `relative`, `absolute`, `fixed`, `sticky`).
"These properties determine how an element is placed on the page.

- **Static** is the default. The element sits in the normal document flow. `top`/`left`/`z-index` properties have no effect.
- **Relative** positions the element relative to its *normal* position. Moving it via `top` or `left` leaves a blank space where it originally was.
- **Absolute** removes the element from normal flow and positions it relative to its closest *positioned* ancestor. 
- **Fixed** is similar to absolute, but it’s positioned relative to the browser *viewport*. It stays glued to the screen even as I scroll.
- **Sticky** behaves like `relative` until a certain scroll threshold is met, at which point it acts like `fixed`."

#### Indepth
For `absolute` positioning, if no ancestor has a position assigned (i.e., they are all `static`), the element will be positioned relative to the initial containing block (usually the `<html>` element). `z-index` only works on elements that have a positioning value other than `static`.

---

### 9. What is the difference between `em` and `rem`?
"Both are relative units used for scalable typography and layout. 

`em` is relative to the font-size of its **direct parent**. This can lead to compounding issues; if I nest `.child` inside `.parent` and both use `2em`, the child is 4 times the baseline size.

`rem` stands for 'root em'. It is always relative to the root element (`<html>`), which is typically 16px by default. I strongly prefer `rem` for general layout and font sizing because it provides consistent sizing across the entire app without complex nesting math."

#### Indepth
While `rem` is preferred for consistent global sizing, `em` is incredibly powerful for component-level scaling. For example, if you build a button component and set its padding to `0.5em 1em`, the padding will automatically scale perfectly if you only change the `font-size` of the button, making it easy to create small, medium, and large button variants.

---

### 10. What is the purpose of the `alt` attribute in images?
"The `alt` (alternative text) attribute provides a textual description of an image. 

I include it for three main reasons: First, **Accessibility**; screen readers read the `alt` text aloud so visually impaired users understand the image's context. Second, **Failsafe**; if the image fails to load due to a bad connection, the browser displays this text instead. Finally, **SEO**; search engines use it to index the image properly."

#### Indepth
If an image is purely decorative (like a background flourish loaded via an `<img>` tag) and conveys no contextual information, you should still include the attribute but leave it empty (`alt=""`). This explicitly tells screen readers to skip the image altogether, rather than them reading out the jarring raw filename.
