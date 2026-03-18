# HTML & CSS Basics (Service-Based Companies)

## 1. What is semantic HTML, and why is it important?
**Answer:**
Semantic HTML introduces meaning to the web page rather than just presentation. Tags like `<header>`, `<footer>`, `<article>`, `<section>`, `<nav>`, and `<main>` clearly describe their physical meaning to both the browser and the developer.
*   **Importance:**
    *   **Accessibility:** Screen readers and assistive technologies use semantic tags to provide a better browsing experience for users with disabilities.
    *   **SEO (Search Engine Optimization):** Search engines give more weight to semantic tags than non-semantic tags (like `<div>` or `<span>`), helping properly index the content.
    *   **Maintainability:** Easier for developers to read and understand the HTML structure.

## 2. Explain the difference between `display: none` and `visibility: hidden`.
**Answer:**
*   `display: none`: Removes the element completely from the document flow. The element takes up no space, and the browser repaints the layout as if the element never existed.
*   `visibility: hidden`: Hides the element visually, but it still occupies its allocated space in the DOM. It affects the layout of the surrounding elements exactly as it would if it were visible.

## 3. What are CSS selectors, and what is its specificity?
**Answer:**
CSS Selectors are rules used to target specific HTML elements for styling. Common types include:
*   **Universal selector:** `*`
*   **Element selector:** `div`, `p`, `h1`
*   **Class selector:** `.classname`
*   **ID selector:** `#idname`
*   **Attribute selector:** `[type="text"]`
*   **Pseudo-classes & Pseudo-elements:** `:hover`, `::before`

**Specificity** determines which CSS rule is applied by the browsers. It's calculated based on a weight hierarchy (Inline > ID > Class/Attribute/Pseudo-class > Element/Pseudo-element).
If two selectors apply to the same element, the one with higher specificity wins. The `!important` rule overrides all normal specificity rules.

## 4. What is the difference between inline, inline-block, and block elements?
**Answer:**
*   **Block elements (`display: block`):** Start on a new line and take up the entire width available. (e.g., `<div>`, `<p>`, `<h1>`). You can set both `width` and `height`.
*   **Inline elements (`display: inline`):** Do not start on a new line; they only take up as much width as necessary. (e.g., `<span>`, `<a>`, `<strong>`). You *cannot* set `width` or `height` (top and bottom margins are also ignored, but left and right apply).
*   **Inline-block elements (`display: inline-block`):** Similar to inline elements, they do not start on a new line. However, unlike inline elements, you *can* set their `width` and `height`.

## 5. How do you center an element horizontally and vertically using CSS?
**Answer:**
*   **Using Flexbox (Modern, Preferred):**
    ```css
    .parent {
        display: flex;
        justify-content: center; /* horizontal */
        align-items: center; /* vertical */
        height: 100vh; /* Needed for vertical centering */
    }
    ```
*   **Using Grid:**
    ```css
    .parent {
        display: grid;
        place-items: center; /* centers both horizontally and vertically */
        height: 100vh;
    }
    ```
*   **Using Absolute Positioning and Transforms:**
    ```css
    .parent {
        position: relative;
    }
    .child {
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
    }
    ```

## 6. What is the CSS Box Model? Briefly describe its components.
**Answer:**
The CSS Box Model is essentially a box that wraps around every HTML element. It determines how elements are sized and spaced in relation to one another.
From inside out, its components are:
1.  **Content:** The actual content of the box, where text or images appear.
2.  **Padding:** Clears an area around the content. The padding is transparent and sits inside the element's background.
3.  **Border:** A border that goes around the padding and content.
4.  **Margin:** Clears an area outside the border. The margin is transparent and separates the element from other elements.

## 7. What is the difference between relative, absolute, fixed, and sticky positioning?
**Answer:**
*   **`static` (Default):** The element is positioned according to the normal document flow. `top`, `right`, `bottom`, `left`, and `z-index` properties have no effect.
*   **`relative`:** Positioned relative to its *normal static position*. Other elements are not adjusted to fit into any gap left by the element.
*   **`absolute`:** Positioned relative to its *closest positioned ancestor* (an ancestor with a position other than static). If no positioned ancestor is found, it uses the document body. It is removed from the normal document flow.
*   **`fixed`:** Positioned relative to the *viewport* (browser window). It stays in the same place even if the page is scrolled. Emoved from the document flow.
*   **`sticky`:** Toggles between relative and fixed depending on the scroll position. It's treated as relative until it crosses a specified threshold, at which point it acts as fixed.

## 8. What are Pseudo-classes and Pseudo-elements? Provide examples.
**Answer:**
*   **Pseudo-classes:** Used to define a special *state* of an element. Single colon (`:`).
    *   *Examples:* `:hover`, `:active`, `:focus`, `:first-child`, `:nth-child()`.
*   **Pseudo-elements:** Used to style a specific *part* of an element. Double colon (`::`).
    *   *Examples:* `::before` (inserts content before the element content), `::after`, `::first-line`, `::first-letter`, `::selection`.

## 9. How do you implement a CSS reset or normalize.css, and why?
**Answer:**
Browsers have default, built-in CSS stylesheets (user-agent styles) that apply varying margins, paddings, and styles to HTML elements.
*   **CSS Reset:** completely strips away all default styling (e.g., sets margins/padding to 0 on everything), forcing the developer to declare all styles explicitly.
*   **Normalize.css:** A softer approach that preserves useful browser defaults while fixing bugs and ensuring consistency across different browsers. It creates a level playing field rather than erasing everything.
*   **Why:** To ensure cross-browser consistency before applying custom styles.

## 10. Can you explain HTML `<canvas>` vs `<svg>`?
**Answer:**
*   **Canvas (`<canvas>`):** Raster-based (pixel-based). Drawn via JavaScript. Best for complex animations, games, or manipulating thousands of objects. When scaled up, it becomes pixelated. Does not support native DOM events for elements drawn inside it.
*   **SVG (`<svg>`):** Vector-based. Drawn via XML format. Best for logos, icons, and illustrations that need to scale infinitely without losing quality. Each element within the SVG is available in the DOM, so you can easily attach event handlers and style them with CSS.

## 11. What are the key features of HTML5?
**Answer:**
HTML5 introduced numerous features that revolutionized web development:

### **Semantic Elements**
- `<article>`, `<section>`, `<header>`, `<footer>`, `<nav>`, `<main>`, `<aside>` - Provide meaning to content structure
- `<figure>`, `<figcaption>` - Images with captions
- `<time>`, `<mark>`, `<details>`, `<summary>` - Enhanced content semantics

### **Multimedia Support**
- `<audio>` - Native audio playback without plugins
- `<video>` - Native video playback without Flash
- `<canvas>` - 2D/3D graphics rendering via JavaScript
- `<svg>` - Scalable vector graphics

### **Form Enhancements**
- **New Input Types:** `email`, `tel`, `url`, `search`, `date`, `time`, `datetime-local`, `month`, `week`, `number`, `range`, `color`
- **New Attributes:** `placeholder`, `required`, `pattern`, `min`, `max`, `step`, `autofocus`, `autocomplete`
- **New Elements:** `<datalist>`, `<progress>`, `<meter>`

### **Storage APIs**
- **localStorage** - Persistent client-side storage (no expiration)
- **sessionStorage** - Session-based storage (clears on tab close)
- **IndexedDB** - Client-side database for structured data

### **Graphics & Media**
- **Canvas API** - Dynamic graphics and animations
- **WebGL** - 3D graphics rendering
- **Drag & Drop API** - Native drag and drop functionality

### **Communication APIs**
- **WebSockets** - Real-time bidirectional communication
- **Server-Sent Events** - Server-to-client updates
- **Web Workers** - Background JavaScript processing
- **Geolocation API** - Location services
- **WebRTC** - Real-time audio/video communication

### **Performance & Optimization**
- **Script Loading:** `defer` and `async` attributes for better performance
- **Resource Hints:** `preload`, `prefetch`, `preconnect`
- **Service Workers** - Offline functionality and advanced caching

### **Device Integration**
- **Device Orientation API** - Detect device rotation
- **Vibration API** - Haptic feedback
- **Camera/Microphone Access** - Media capture capabilities

### **Accessibility Improvements**
- **ARIA attributes** integration
- **Enhanced semantic markup** for screen readers
- **Improved `tabindex`** and focus management

### **Other Features**
- **ContentEditable** - Inline editing capabilities
- **Data attributes** (`data-*`) - Custom data storage
- **Cross-Origin Resource Sharing (CORS)** support
- **Better error handling and parsing**
