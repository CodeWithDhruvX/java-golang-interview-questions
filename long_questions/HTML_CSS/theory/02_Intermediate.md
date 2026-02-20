# 🟡 **Intermediate: Advanced HTML & CSS Layouts**

### 1. What is the difference between `defer` and `async` in script loading?
"Both attributes allow the browser to continue parsing the HTML document while the script downloads in the background, rather than blocking rendering entirely.

When I use `async`, the script executes exactly when it finishes downloading, regardless of where the HTML parser is. This means scripts might execute out of order. I use this for independent third-party scripts like Google Analytics.

When I use `defer`, the script downloads in the background but *waits* to execute until the entire HTML document is fully parsed. I use this for scripts that rely on the DOM being fully loaded, like React or Vue initialization scripts, and they are guaranteed to execute in the order they appear in the document."

#### Indepth
Prior to these attributes, the best practice was to put `<script>` tags immediately before the closing `</body>` tag to avoid render blocking. Today, placing `<script defer src="...">` in the `<head>` is the modern standard because the browser can start fetching the script immediately while still rendering the page.

---

### 2. What is the difference between `localStorage`, `sessionStorage`, and cookies?
"These are three different ways to store client-side data.

**Cookies** are tiny (max 4KB) and are automatically sent to the server with *every* HTTP request. I use them for session tokens and authentication states.
**sessionStorage** isolates data per tab. Once the specific browser tab is closed, the data is completely erased. I use this for temporary state like step-by-step form wizards.
**localStorage** persists data permanently across sessions and tabs until explicitly cleared by the user or JavaScript. It holds much more data (usually 5MB+). I use this for user preferences like dark mode."

#### Indepth
Both Web Storage APIs (`localStorage` and `sessionStorage`) are synchronous and accessed on the main thread, meaning large reads/writes can block the UI. If you need to store significant amounts of complex data locally in a non-blocking way, IndexedDB is the correct, albeit more complex, API to use.

---

### 3. How do you calculate CSS Specificity?
"CSS specificity determines which style rule wins when multiple rules target the same element. It uses a four-tier point system.

From highest priority to lowest:
1. **Inline styles** (`style="..."`) 
2. **IDs** (`#header`)
3. **Classes, Attributes, and Pseudo-classes** (`.btn`, `[type='text']`, `:hover`)
4. **Elements and Pseudo-elements** (`div`, `::before`)

I calculate this mentally as a comma-separated value: (inline, id, class, element). So a rule like `#nav .link a:hover` has 1 ID, 2 Classes (because `:hover` counts as a class), and 1 Element. The higher leftmost number always wins."

#### Indepth
If two rules have the exact same specificity score, the one declared *latest* in the CSS file takes precedence. The `!important` declaration overrides normal specificity rules entirely, meaning even an inline style will lose to a tag selector marked `!important`. It should be used very sparingly as it makes debugging difficult.

---

### 4. What is Flexbox, and when do you use it?
"Flexbox is a one-dimensional layout system that excels at aligning items and distributing space either in a single row or a single column.

Before Flexbox, horizontal layouts required tedious CSS floats or table hacks. With Flexbox, assigning `display: flex` on a container immediately aligns its children logically. 

I use it constantly for small-to-medium layout components: navigation bars, aligning icons next to text, vertical centering (`align-items: center; justify-content: center`), and making rows of content fluid with `flex-grow`."

#### Indepth
The core power of Flexbox lies in its dynamic alignment capabilities: `justify-content` aligns items along the main axis (row or column), while `align-items` aligns them along the cross axis. The `flex` shorthand property `flex: 1 0 auto` sets `flex-grow`, `flex-shrink`, and `flex-basis` simultaneously.

---

### 5. What is CSS Grid, and how does it differ from Flexbox?
"CSS Grid is a two-dimensional layout system, meaning it can handle both rows and columns simultaneously.

While Flexbox is content-out (the items dictate the layout width), Grid is layout-in (you define the rigid grid tracks, and items snap into them). 

I use Grid for overall page architectures and complex, rigid layouts, like a classic header-sidebar-content-footer dashboard (`grid-template-areas` is amazing for this) or responsive photo galleries. I often use *both* together: Grid for the macro layout, and Flexbox for the micro layout within individual grid cells."

#### Indepth
A very powerful Grid pattern for responsive design without media queries is `grid-template-columns: repeat(auto-fit, minmax(250px, 1fr))`. This tells the browser to automatically wrap elements into as many columns as will fit on the screen, heavily simplifying component logic.

---

### 6. What is Stacking Context in CSS?
"Stacking Context is a 3D conceptualization of HTML elements along the z-axis (facing the user). It determines the painting order of elements when they overlap.

Most developers think `z-index` simply determines what goes on top, but it only works within the same Stacking Context. If I have a modal with a `z-index` of 9999 inside a container that creates a new stacking context with a `z-index` of 1, that modal will never appear above an adjacent container with a `z-index` of 2.

I create new stacking contexts deliberately to isolate components (like a complex dropdown menu) so their internal z-indexes don't bleed out and affect the rest of the page."

#### Indepth
While `position: absolute`/`relative` with `z-index` is the most common way to create a stacking context, several other properties trigger it implicitly. For instance, `opacity` less than 1, `transform`, `filter`, or `will-change` will immediately trap child z-indexes within that element, which is a frequent source of "why isn't my z-index working?" bugs.

---

### 7. What is the BEM methodology?
"BEM stands for **Block, Element, Modifier**. It is a CSS naming convention that makes large codebases predictable and easy to scale.

1. **Block**: A standalone entity that is meaningful on its own (e.g., `.card`)
2. **Element**: A part of a block that has no standalone meaning and is semantically tied to its block (e.g., `.card__title` – separated by double underscores)
3. **Modifier**: A flag on a block or element to change its appearance or state (e.g., `.card--dark` or `.card__title--large` – separated by double hyphens).

I use it because it completely prevents CSS specificity wars. Because every class is unique and flat, I never have to write heavily nested selectors like `#main .container .card h2`."

#### Indepth
BEM pairs exceptionally well with CSS preprocessors like SASS. You can use the parent selector `&` to nest elements cleanly: `.card { &__title { color: red; } }`. It results in highly modular, reusable components that translate perfectly into modern frameworks like React or Vue.

---

### 8. How do you implement Responsive Design?
"Responsive design ensures a website looks good on all devices, from mobile phones to 4K desktop monitors. 

My approach involves three key principles:
1. **Fluid Grids**: I use relative units like `%`, `vw`, or `fr` (in Grid) instead of hardcoded pixels for layouts so containers expand dynamically.
2. **Flexible Media**: I set `max-width: 100%` on images and videos so they scale down inside smaller containers rather than overflowing them.
3. **Media Queries**: I use `@media (min-width: 768px)` to apply different layouts based on the screen width, such as shifting a Flexbox column into a row on desktop viewports."

#### Indepth
The modern preferred approach is **Mobile-First Design**. You write the default CSS for mobile screens first without any media queries, and then use `min-width` media queries to progressively enhance the layout for tablets and desktops. This results in cleaner, faster-loading CSS for mobile users who frequently have slower network connections.

---

### 9. How do you center a div horizontally and vertically?
"This is the classic CSS interview question, and modern CSS has made it trivial.

My preferred way is using **Flexbox**:
1. On the parent container, I apply `display: flex`.
2. `justify-content: center` centers it horizontally (along the main axis).
3. `align-items: center` centers it vertically (along the cross axis).

If I need to use **CSS Grid**, it’s even shorter:
1. Apply `display: grid` to the parent.
2. Use `place-items: center;` which combines horizontal and vertical centering perfectly."

#### Indepth
Before Grid and Flexbox, the fallback method involved absolute positioning. You would position the child: `position: absolute; top: 50%; left: 50%;`. Because this aligned the *top-left corner* of the child to the center of the parent, you then had to pull it back by exactly half its width and height using: `transform: translate(-50%, -50%);`.

---

### 10. Explain HTTP form submissions: GET vs POST.
"When building HTML forms (`<form method="...">`), choosing the correct HTTP verb is critical.

I use **GET** when I want to retrieve data without changing server state, like a search form or filtering a list. The form data is appended directly to the URL string (e.g., `?search=shoes`). This makes the query bookmarkable and shareable, but it is highly insecure for sensitive data.

I use **POST** when the form submission will change server state—like creating a user account, uploading a file, or processing payment. The data is sent securely in the HTTP request body instead of the URL."

#### Indepth
While HTML5 `<form>` tags only support GET and POST natively, modern web apps use JavaScript (Fetch/Axios) to intercept form submissions and send PUT, PATCH, or DELETE requests according to RESTful API principles. Additionally, when uploading files, the form must explicitly define `enctype="multipart/form-data"` to handle binary data correctly.
