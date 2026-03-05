# Web Accessibility (a11y) (Product-Based Companies)

## 1. What is Web Accessibility (a11y), and why is it crucial for product companies?
**Answer:**
Web Accessibility (often abbreviated as a11y, where 11 represents the number of letters between 'a' and 'y') refers to the practice of designing and developing websites, tools, and technologies so that people with disabilities can use them.
*   **Crucial for product companies because:**
    *   **Inclusivity & Market Reach:** It ensures all users, including those with visual, auditory, motor, or cognitive disabilities, can access the product, expanding the audience base.
    *   **Legal Compliance:** Many countries have strict laws (like the ADA in the US or the EAA in Europe) requiring digital accessibility, with hefty fines for non-compliance.
    *   **SEO & UX:** Many accessibility best practices (like semantic HTML, proper contrast, logical structure) inherently improve SEO and the overall user experience for *everyone*.

## 2. What is ARIA, and when should you use it?
**Answer:**
ARIA (Accessible Rich Internet Applications) is a set of attributes defined by the W3C that supplement HTML. They provide additional semantics to elements, telling assistive technologies (like screen readers) what an element is, what its state is, and what it does.
*   **Golden Rule of ARIA:** *No ARIA is better than bad ARIA.*
*   **When to use it:** You should *always* prefer native HTML elements (`<button>`, `<nav>`, `<dialog>`) because they have built-in accessibility. Use ARIA *only* when you are building complex, custom UI components (like custom dropdowns, tabs, or modals) where native HTML lacks the necessary semantics to convey the component's role, state (e.g., `aria-expanded`, `aria-checked`), or properties (e.g., `aria-controls`, `aria-describedby`).

## 3. Explain the difference between `aria-hidden="true"`, `hidden` attribute, and `.sr-only` (visually hidden class).
**Answer:**
*   **`hidden` (HTML5 attribute):** Tells the browser not to render the element *at all*. It is equivalent to `display: none` in CSS. The element is removed from both the visual layout and the accessibility tree (screen readers won't see it).
*   **`aria-hidden="true"`:** Tells assistive technologies (screen readers) to ignore the element (it removes it from the accessibility tree). However, it remains visually present on the screen. Used for decorative elements or duplicate content that might confuse screen reader users.
*   **`.sr-only` (Screen Reader Only Class):** A CSS technique used to hide content *visually*, but keep it available in the DOM for screen readers to announce.
    ```css
    .sr-only {
      position: absolute;
      width: 1px;
      height: 1px;
      padding: 0;
      margin: -1px;
      overflow: hidden;
      clip: rect(0, 0, 0, 0);
      white-space: nowrap;
      border: 0;
    }
    ```
    Useful for providing audible context where visual context is implied (e.g., a search button that only has a magnifying glass icon).

## 4. How do you handle Focus Management in a Single Page Application (SPA)?
**Answer:**
In a traditional multi-page website, moving to a new page resets the focus to the top of the new document. In SPAs (like React, Angular, Vue), routing changes the DOM without a full page reload, leading to a "loss of focus" where the screen reader doesn't realize the page content has changed.
*   **Solution:** You must manually manage focus.
    1.  When a route changes, programmatically move the focus to a logical element on the new 'page' (often a main heading like `<h1>`, which should have `tabindex="-1"` so it can receive programmatic focus without being in the natural tab flow).
    2.  For custom modals/dialogs, you must trap focus *inside* the modal while it's open, and return focus to the trigger button when it closes.

## 5. What is the `tabindex` attribute, and what are its standard values?
**Answer:**
The `tabindex` attribute indicates whether an element can receive keyboard focus and in what order.
*   **`tabindex="0"`:** Adds the element to the natural tab order (e.g., making a custom `<div>` focusable like a button).
*   **`tabindex="-1"`:** Removes the element from the natural tab order (it cannot be reached via the Tab key), *but* it allows the element to receive focus programmatically via JavaScript (e.g., `element.focus()`). Essential for custom widgets and SPA routing.
*   **`tabindex="> 0"` (Positive numbers):** *Anti-pattern.* It explicitly sets a tab order, overriding the logical source-code order. This almost always creates a thoroughly confusing, jumping experience for users. Do not use this; fix the DOM structure instead.

## 6. What are the WCAG guidelines for Color Contrast?
**Answer:**
WCAG (Web Content Accessibility Guidelines) mandates specific contrast ratios between foreground text and its background color to ensure readability for users with low vision or color blindness.
*   **AA Standard (Minimum requirement for most companies):**
    *   **Normal text:** Ratio of at least **4.5:1**.
    *   **Large text** (18pt or 14pt bold): Ratio of at least **3:1**.
    *   **UI Components & Graphical Objects** (e.g., input borders, icons): Ratio of at least **3:1**.
*   **AAA Standard (Enhanced requirement):** Requires 7:1 for normal text and 4.5:1 for large text.

## 7. How do you make HTML Forms accessible?
**Answer:**
Forms are frequent points of failure. Best practices include:
1.  **Labels:** Every input `<input>`, `<select>`, `<textarea>` *must* have an associated `<label>`. The `id` of the input must match the `for` attribute of the label.
2.  **Fieldsets:** Group related controls (like radio buttons or checkboxes) within a `<fieldset>` and use a `<legend>` to describe the group.
3.  **Error Handling:**
    *   Clearly associate error messages with the specific input field using `aria-describedby` or `aria-errormessage`.
    *   Use `aria-invalid="true"` on the input when it fails validation.
    *   Programmatically move focus to the first invalid field upon form submission failure.
4.  **Avoid Placeholder as Label:** Placeholders disappear when the user starts typing, causing strain on short-term memory. They also often fail color contrast tests.

## 8. What is a "Skip to Main Content" link, and why is it needed?
**Answer:**
It's a visually hidden anchor link placed at the very top of the HTML document (the first focusable element). When a keyboard user hits the Tab key upon loading the page, this link becomes visible and receives focus.
**Why:** Many sites have complex navigation menus at the top. Sighted users can visually skip over them to read the main content. Without a "skip link," a keyboard-only or screen reader user must manually tab through every single navigation link *on every single page* just to reach the main article. The skip link targets an ID on the `<main>` tag, allowing them to bypass the boilerplate.

## 9. How do you construct an accessible button containing only an icon?
**Answer:**
If a button only has an SVG or an icon font glyph, a screen reader will either announce nothing or read out a random file name.
*   **Solution 1 (Visually Hidden Text):**
    ```html
    <button>
      <svg aria-hidden="true">...</svg>
      <span class="sr-only">Close Modal</span>
    </button>
    ```
    (The `aria-hidden` prevents the screen reader from trying to read the SVG code itself).
*   **Solution 2 (aria-label):**
    ```html
    <button aria-label="Close Modal">
      <svg aria-hidden="true">...</svg>
    </button>
    ```

## 10. Explain the difference between `role="presentation"` and `role="none"`.
**Answer:**
They are exactly the same. They both remove the semantic meaning of an element from the accessibility tree, treating it purely as a visual, decorative element.
For example, if you use a `<table>` purely for layout purposes (which is a bad practice nowadays, but common in HTML email), you would apply `role="presentation"` to tell screen readers not to announce "Table with x rows and y columns."
`role="none"` was introduced in ARIA 1.1 because `presentation` was sometimes misunderstood to mean "hide this element," when it only means "strip the semantics."
