# Practical UI Implementation & Edge Cases (Product-Based Companies)

## 1. How do you build a custom-styled Checkbox or Radio Button using CSS?
**Answer:**
Native `<input type="checkbox">` and `<input type="radio">` are notoriously difficult to style uniformly across browsers. The standard practice involves hiding the native input and styling a pseudo-element.
1.  **Structure:** Wrap the input and the visible label text inside a `<label>`.
    ```html
    <label class="custom-checkbox">
      <input type="checkbox" class="sr-only"> <!-- Visually hide, but keep focusable -->
      <span class="checkmark"></span>
      Accept Terms
    </label>
    ```
2.  **Hide Native Input:** Visually hide the input using `.sr-only` (screen-reader only class) or `opacity: 0; position: absolute;` so it remains focusable and accessible. Never use `display: none` or `visibility: hidden`, as this breaks accessibility.
3.  **Style the Custom Mark:** Style the `span.checkmark` to look like an unchecked box.
4.  **Handle the Checked State:** Use the `:checked` pseudo-class and the adjacent sibling combinator (`+`) or general sibling combinator (`~`) to style the checkmark when the hidden input is checked.
    ```css
    input:checked + .checkmark {
      background-color: blue;
    }
    input:checked + .checkmark::after {
       content: '✓'; /* Or use CSS pseudo-elements to draw a check */
       color: white;
    }
    ```
5.  **Handle Focus state:** Be sure to apply focus styles when the hidden input is navigated to via keyboard: `input:focus + .checkmark { outline: 2px solid blue; }`.

## 2. Explain CSS Logical Properties and how they help with i18n (Internationalization).
**Answer:**
Traditionally, we used physical properties like `margin-left`, `padding-right`, `border-top`, etc.
However, when translating a site for languages that read Right-to-Left (RTL), like Arabic or Hebrew, `margin-left` stays on the left, breaking the layout. You would have to write complex overrides using `[dir="rtl"]`.
**CSS Logical Properties** map CSS properties to the flow of text (writing mode) rather than the physical screen.
*   **Inline Axis:** The direction text flows (Left-to-Right in English, Right-to-Left in Arabic).
    *   Left becomes `inline-start`
    *   Right becomes `inline-end`
*   **Block Axis:** The direction lines stack (Top-to-Bottom in English).
    *   Top becomes `block-start`
    *   Bottom becomes `block-end`
*   **Example:** Instead of `margin-left: 10px;`, use `margin-inline-start: 10px;`.
Now, if you switch the document direction to RTL (`<html dir="rtl">`), the margin automatically switches to the right side without any additional CSS.

## 3. How do you create a Dropdown Menu solely with CSS (No JavaScript)?
**Answer:**
This is commonly done using the `:hover` pseudo-class (for mouse users) and `:focus-within` (for keyboard accessibility).
1.  **Structure:** A parent container holding the trigger (e.g., a button or link) and the dropdown content (usually a `<ul>`).
2.  **Default State:** Hide the dropdown content using `display: none` or `opacity: 0; visibility: hidden;` (preferred for transitions).
3.  **Active State:** Use a CSS combinator to show the dropdown content when the parent container is hovered or focused globally.
    ```css
    .dropdown .dropdown-content {
      visibility: hidden;
      opacity: 0;
      position: absolute;
    }

    /* Target the content when the parent is hovered or has focus inside it */
    .dropdown:hover .dropdown-content,
    .dropdown:focus-within .dropdown-content {
      visibility: visible;
      opacity: 1;
    }
    ```

## 4. What are CSS Feature Queries (`@supports`), and when should you use them?
**Answer:**
`@supports` is a CSS at-rule that lets you perform feature detection directly inside your stylesheet, bypassing the need for JavaScript libraries like Modernizr.
It allows you to test whether a browser supports a specific CSS property-value pair before applying a block of styles.
**Use Case (Progressive Enhancement):**
You want to use modern CSS Grid for layout, but you need to support much older browsers that only understand Flexbox or Floats.
```css
/* 1. Base Fallback styles (for older browsers) */
.layout {
  display: flex;
}

/* 2. Enhanced styles (for modern browsers) */
@supports (display: grid) {
  .layout {
    display: grid;
    grid-template-columns: 1fr 1fr;
    /* Grid overrides flexbox */
  }
}
```

## 5. How does native HTML5 Form Validation work?
**Answer:**
HTML5 introduced built-in form validation, reducing the reliance on JavaScript for basic checks.
*   **Validation Attributes:**
    *   `required`: Forces the field to be filled out.
    *   `min` / `max`: Sets limits for `type="number"` or dates.
    *   `pattern`: Uses a regular expression the input must match (e.g., `pattern="[0-9]{3}"` for a 3-digit number).
    *   `minlength` / `maxlength`: Sets character limits.
    *   `type="email"`, `type="url"`: Native pattern matching for these formats.
*   **CSS Pseudo-classes:** You can style elements based on their valid state.
    *   `:valid`: Styles the input when it meets all constraints.
    *   `:invalid`: Styles the input when it fails constraints.
    *   `:user-invalid`: (Modern) Only applies the invalid style *after* the user has interacted with the field, preventing a sea of red errors when the form first loads.
*   **Form Submission:** The browser prevents the form from submitting (`onsubmit` is not fired) if any of these constraints fail, displaying a native tooltip with an error message.

## 6. How do you create an accessible, responsive Tooltip using pure CSS?
**Answer:**
You can create tooltips using `data-*` attributes and CSS pseudo-elements.
1.  **HTML:** Add a custom data attribute to the element you want the tooltip on.
    ```html
    <button class="tooltip" data-tooltip="This is extra info">Hover Me</button>
    ```
2.  **CSS Setup:** Set the parent to `position: relative`. Use `::after` to create the tooltip bubble based on the data attribute's content.
    ```css
    .tooltip {
      position: relative;
    }
    .tooltip::after {
      content: attr(data-tooltip); /* Pull the text from the attribute */
      position: absolute;
      bottom: 120%; /* Position above the button */
      left: 50%;
      transform: translateX(-50%);
      opacity: 0;
      pointer-events: none; /* Prevents the tooltip from interfering with clicks */
      transition: opacity 0.2s;
    }
    ```
3.  **Hover/Focus State:** Reveal the tooltip.
    ```css
    .tooltip:hover::after,
    .tooltip:focus::after {
      opacity: 1;
    }
    ```

## 7. Explain the `aspect-ratio` property in modern CSS.
**Answer:**
Historically, developers used the "Padding Hack" (setting `padding-top` to a percentage, like `56.25%` for 16:9, on a zero-height container) to maintain a specific aspect ratio for embedded videos or images (preventing layout shifts as they load).
The modern `aspect-ratio` property solves this natively and elegantly.
```css
.card-image {
  width: 100%;
  aspect-ratio: 16 / 9; /* Width is established by layout, height is calculated automatically */
  object-fit: cover; /* Prevents the image source from stretching */
}
```
This guarantees the box will always hold its shape as it scales responsively.
