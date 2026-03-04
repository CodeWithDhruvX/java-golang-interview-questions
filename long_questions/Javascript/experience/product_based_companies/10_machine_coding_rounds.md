# ⚡ 10 — Machine Coding Round Questions
> **Most Asked in Product-Based Companies** | 🔴 Difficulty: Hard

---

## 🔑 What is the Machine Coding Round?
Build a **mini feature/component** in 60–90 minutes using plain JS or React.
Evaluated on: **correctness, code quality, component design, edge cases, extensibility.**

---

## ❓ Top Machine Coding Problems

### Q1. Build an Autocomplete / Search Suggestion Component

```javascript
// Requirements:
// - Input shows dropdown suggestions as user types
// - Debounce API calls (300ms)
// - Keyboard navigation (↑↓ to navigate, Enter to select, Esc to close)
// - Highlight matching text in suggestions
// - Handle loading and error states

class Autocomplete {
    #el;
    #input;
    #dropdown;
    #items = [];
    #activeIndex = -1;
    #debounceTimer;
    #fetchSuggestions;

    constructor(container, fetchSuggestions) {
        this.#fetchSuggestions = fetchSuggestions;
        this.#el = container;
        this.#render();
        this.#attachEvents();
    }

    #render() {
        this.#el.innerHTML = `
            <div class="autocomplete" role="combobox" aria-expanded="false">
                <input type="text" id="search" aria-autocomplete="list"
                       aria-controls="suggestions" autocomplete="off"
                       placeholder="Search..."/>
                <ul id="suggestions" role="listbox" hidden></ul>
            </div>
        `;
        this.#input    = this.#el.querySelector('input');
        this.#dropdown = this.#el.querySelector('ul');
    }

    #attachEvents() {
        this.#input.addEventListener('input', () => this.#onInput());
        this.#input.addEventListener('keydown', (e) => this.#onKeydown(e));
        document.addEventListener('click', (e) => {
            if (!this.#el.contains(e.target)) this.#close();
        });
    }

    #onInput() {
        const query = this.#input.value.trim();
        clearTimeout(this.#debounceTimer);
        if (!query) { this.#close(); return; }

        this.#debounceTimer = setTimeout(async () => {
            this.#setLoading(true);
            try {
                this.#items = await this.#fetchSuggestions(query);
                this.#renderItems(query);
            } catch (err) {
                this.#renderError();
            } finally {
                this.#setLoading(false);
            }
        }, 300);
    }

    #renderItems(query) {
        if (!this.#items.length) {
            this.#dropdown.innerHTML = `<li class="no-results">No results found</li>`;
        } else {
            this.#dropdown.innerHTML = this.#items.map((item, i) => `
                <li role="option" data-index="${i}" aria-selected="false">
                    ${this.#highlight(item, query)}
                </li>
            `).join('');

            this.#dropdown.querySelectorAll('li[data-index]').forEach(li => {
                li.addEventListener('mousedown', (e) => {
                    e.preventDefault(); // prevent input blur before click
                    this.#select(+li.dataset.index);
                });
            });
        }
        this.#open();
    }

    // Bold the matching portion
    #highlight(text, query) {
        const regex = new RegExp(`(${query.replace(/[.*+?^${}()|[\]\\]/g, '\\$&')})`, 'gi');
        return text.replace(regex, '<mark>$1</mark>');
    }

    #onKeydown(e) {
        const options = this.#dropdown.querySelectorAll('li[data-index]');
        if (!options.length) return;

        switch (e.key) {
            case 'ArrowDown':
                e.preventDefault();
                this.#activeIndex = (this.#activeIndex + 1) % options.length;
                this.#updateActive(options);
                break;
            case 'ArrowUp':
                e.preventDefault();
                this.#activeIndex = (this.#activeIndex - 1 + options.length) % options.length;
                this.#updateActive(options);
                break;
            case 'Enter':
                if (this.#activeIndex >= 0) this.#select(this.#activeIndex);
                break;
            case 'Escape':
                this.#close();
                break;
        }
    }

    #updateActive(options) {
        options.forEach((li, i) => {
            li.setAttribute('aria-selected', i === this.#activeIndex);
            if (i === this.#activeIndex) li.scrollIntoView({ block: 'nearest' });
        });
    }

    #select(index) {
        this.#input.value = this.#items[index];
        this.#el.dispatchEvent(new CustomEvent('select', { detail: this.#items[index] }));
        this.#close();
    }

    #open() {
        this.#dropdown.hidden = false;
        this.#el.querySelector('.autocomplete').setAttribute('aria-expanded', 'true');
        this.#activeIndex = -1;
    }

    #close() {
        this.#dropdown.hidden = true;
        this.#el.querySelector('.autocomplete').setAttribute('aria-expanded', 'false');
        this.#activeIndex = -1;
    }

    #setLoading(loading) {
        this.#input.setAttribute('aria-busy', loading);
    }

    #renderError() {
        this.#dropdown.innerHTML = `<li class="error">Failed to load suggestions</li>`;
        this.#open();
    }
}

// Usage
const ac = new Autocomplete(document.getElementById('app'), async (query) => {
    const res = await fetch(`/api/search?q=${encodeURIComponent(query)}`);
    return res.json(); // returns string[]
});

ac.addEventListener('select', (e) => console.log('Selected:', e.detail));
```

---

### Q2. Build an Infinite Scroll Feed

```javascript
// Requirements:
// - Load items when user scrolls near the bottom
// - Show skeleton loading cards
// - Handle network error with retry
// - No duplicate fetches (debounce + lock)
// - Support "all items loaded" state

class InfiniteScroll {
    #container;
    #page = 1;
    #loading = false;
    #allLoaded = false;
    #observer;
    #fetchPage;

    constructor(container, fetchPage) {
        this.#container = container;
        this.#fetchPage = fetchPage;
        this.#setupObserver();
        this.#loadMore(); // load initial page
    }

    #setupObserver() {
        // Sentinel element at the bottom of the list
        const sentinel = document.createElement('div');
        sentinel.id = 'scroll-sentinel';
        this.#container.appendChild(sentinel);

        this.#observer = new IntersectionObserver(
            ([entry]) => { if (entry.isIntersecting) this.#loadMore(); },
            { rootMargin: '200px' } // start loading 200px before hitting bottom
        );
        this.#observer.observe(sentinel);
    }

    async #loadMore() {
        if (this.#loading || this.#allLoaded) return;
        this.#loading = true;

        const skeletons = this.#showSkeletons(3);

        try {
            const { items, hasMore } = await this.#fetchPage(this.#page);
            skeletons.forEach(s => s.remove());
            items.forEach(item => this.#renderItem(item));

            if (!hasMore) {
                this.#allLoaded = true;
                this.#observer.disconnect();
                this.#showEndMessage();
            }
            this.#page++;
        } catch (err) {
            skeletons.forEach(s => s.remove());
            this.#showError(() => {
                this.#loading = false;
                this.#loadMore(); // retry on button click
            });
        } finally {
            this.#loading = false;
        }
    }

    #renderItem({ id, title, image, description }) {
        const sentinel = document.getElementById('scroll-sentinel');
        const card = document.createElement('article');
        card.className = 'feed-card';
        card.innerHTML = `
            <img src="${image}" alt="${title}" loading="lazy"/>
            <h3>${title}</h3>
            <p>${description}</p>
        `;
        this.#container.insertBefore(card, sentinel);
    }

    #showSkeletons(count) {
        return Array.from({ length: count }, () => {
            const skeleton = document.createElement('div');
            skeleton.className = 'skeleton-card';
            skeleton.innerHTML = `
                <div class="skeleton skeleton-image"></div>
                <div class="skeleton skeleton-text"></div>
                <div class="skeleton skeleton-text short"></div>
            `;
            this.#container.insertBefore(skeleton, document.getElementById('scroll-sentinel'));
            return skeleton;
        });
    }

    #showError(onRetry) {
        const error = document.createElement('div');
        error.className = 'error-state';
        error.innerHTML = `<p>Failed to load</p><button>Try Again</button>`;
        error.querySelector('button').addEventListener('click', () => {
            error.remove();
            onRetry();
        });
        this.#container.appendChild(error);
    }

    #showEndMessage() {
        const msg = document.createElement('p');
        msg.className = 'end-message';
        msg.textContent = "You've reached the end!";
        this.#container.appendChild(msg);
    }
}

// Usage
new InfiniteScroll(document.getElementById('feed'), async (page) => {
    const res = await fetch(`/api/posts?page=${page}&limit=10`);
    return res.json(); // { items: [...], hasMore: boolean }
});
```

---

### Q3. Build a Debounced Multi-Select Filter

```javascript
// Requirements:
// - Filter a list of products by multiple categories, price range, rating
// - Debounce filter application
// - URL sync (update query params as filters change)
// - Show active filter count badge

class ProductFilter {
    #products;
    #filters = { categories: new Set(), minPrice: 0, maxPrice: Infinity, minRating: 0 };
    #onFilter;
    #debounceTimer;

    constructor(products, onFilter) {
        this.#products = products;
        this.#onFilter = onFilter;
        this.#restoreFromURL();
    }

    setCategory(category, active) {
        active ? this.#filters.categories.add(category)
               : this.#filters.categories.delete(category);
        this.#apply();
    }

    setPriceRange(min, max) {
        this.#filters.minPrice = min;
        this.#filters.maxPrice = max;
        this.#apply();
    }

    setMinRating(rating) {
        this.#filters.minRating = rating;
        this.#apply();
    }

    #apply() {
        clearTimeout(this.#debounceTimer);
        this.#debounceTimer = setTimeout(() => {
            const { categories, minPrice, maxPrice, minRating } = this.#filters;

            const filtered = this.#products.filter(p => {
                if (categories.size > 0 && !categories.has(p.category)) return false;
                if (p.price < minPrice || p.price > maxPrice) return false;
                if (p.rating < minRating) return false;
                return true;
            });

            this.#syncURL();
            this.#onFilter(filtered, this.#getActiveCount());
        }, 200);
    }

    #getActiveCount() {
        let count = this.#filters.categories.size;
        if (this.#filters.minPrice > 0)        count++;
        if (this.#filters.maxPrice < Infinity)  count++;
        if (this.#filters.minRating > 0)        count++;
        return count;
    }

    #syncURL() {
        const params = new URLSearchParams();
        if (this.#filters.categories.size)  params.set('cats', [...this.#filters.categories].join(','));
        if (this.#filters.minPrice)          params.set('minP', this.#filters.minPrice);
        if (this.#filters.maxPrice < Infinity) params.set('maxP', this.#filters.maxPrice);
        if (this.#filters.minRating)         params.set('rating', this.#filters.minRating);
        // Update URL without page reload
        window.history.replaceState(null, '', `?${params.toString()}`);
    }

    #restoreFromURL() {
        const params = new URLSearchParams(window.location.search);
        if (params.get('cats')) params.get('cats').split(',').forEach(c => this.#filters.categories.add(c));
        if (params.get('minP')) this.#filters.minPrice = +params.get('minP');
        if (params.get('maxP')) this.#filters.maxPrice = +params.get('maxP');
        if (params.get('rating')) this.#filters.minRating = +params.get('rating');
    }

    reset() {
        this.#filters = { categories: new Set(), minPrice: 0, maxPrice: Infinity, minRating: 0 };
        this.#apply();
    }
}
```

---

### Q4. Build a Star Rating Component (Vanilla JS)

```javascript
// Requirements:
// - Interactive star rating (1–5)
// - Hover preview before selection
// - Half-star support
// - Read-only mode
// - Accessible (keyboard + screen reader)

class StarRating {
    #container;
    #value = 0;
    #hovered = 0;
    #max;
    #readOnly;
    #onChange;

    constructor(container, { max = 5, value = 0, readOnly = false, onChange } = {}) {
        this.#container = container;
        this.#max       = max;
        this.#value     = value;
        this.#readOnly  = readOnly;
        this.#onChange  = onChange;
        this.#render();
    }

    #render() {
        this.#container.innerHTML = `
            <div class="star-rating" role="group" aria-label="Rating">
                ${Array.from({ length: this.#max }, (_, i) => `
                    <button
                        class="star"
                        data-value="${i + 1}"
                        aria-label="${i + 1} star${i + 1 > 1 ? 's' : ''}"
                        aria-pressed="${i + 1 <= this.#value}"
                        ${this.#readOnly ? 'disabled' : ''}
                    >★</button>
                `).join('')}
                <span class="rating-value" aria-live="polite">${this.#value || 'Not rated'}</span>
            </div>
        `;

        if (!this.#readOnly) this.#attachEvents();
        this.#updateDisplay(this.#value);
    }

    #attachEvents() {
        const stars = this.#container.querySelectorAll('.star');

        stars.forEach(star => {
            star.addEventListener('click', () => {
                const val = +star.dataset.value;
                this.#value = val;
                this.#onChange?.(val);
                this.#updateDisplay(val);
                this.#updateAria(val);
            });

            star.addEventListener('mouseenter', () => {
                this.#hovered = +star.dataset.value;
                this.#updateDisplay(this.#hovered);
            });
        });

        this.#container.querySelector('.star-rating').addEventListener('mouseleave', () => {
            this.#hovered = 0;
            this.#updateDisplay(this.#value);
        });
    }

    #updateDisplay(activeValue) {
        this.#container.querySelectorAll('.star').forEach((star, i) => {
            const filled = i + 1 <= activeValue;
            star.classList.toggle('filled', filled);
            star.style.color = filled ? '#f59e0b' : '#d1d5db';
        });
        const label = this.#container.querySelector('.rating-value');
        if (label) label.textContent = activeValue ? `${activeValue}/${this.#max}` : 'Not rated';
    }

    #updateAria(val) {
        this.#container.querySelectorAll('.star').forEach((star, i) => {
            star.setAttribute('aria-pressed', i + 1 <= val);
        });
    }

    getValue() { return this.#value; }
    setValue(val) { this.#value = val; this.#updateDisplay(val); }
}

// Usage
const rating = new StarRating(document.getElementById('rating-widget'), {
    max: 5,
    value: 3,
    onChange: (val) => console.log('Rating selected:', val)
});
```

---

### Q5. Build a Drag-and-Drop Kanban Board

```javascript
// Requirements:
// - Columns: Todo, In Progress, Done
// - Drag cards between columns
// - Add new cards, delete cards
// - Persist state to localStorage

class KanbanBoard {
    #state;
    #el;
    #draggedCard = null;

    constructor(container) {
        this.#el = container;
        this.#state = this.#load() || {
            columns: {
                todo:       { title: "📋 Todo",        cards: [] },
                inProgress: { title: "⚡ In Progress",  cards: [] },
                done:       { title: "✅ Done",          cards: [] }
            }
        };
        this.#render();
    }

    #load() {
        try {
            return JSON.parse(localStorage.getItem('kanban'));
        } catch { return null; }
    }

    #save() {
        localStorage.setItem('kanban', JSON.stringify(this.#state));
    }

    #render() {
        this.#el.innerHTML = `
            <div class="kanban-board">
                ${Object.entries(this.#state.columns).map(([id, col]) => `
                    <div class="column" data-column="${id}">
                        <h2>${col.title} <span class="count">${col.cards.length}</span></h2>
                        <div class="cards-container" data-column="${id}">
                            ${col.cards.map(card => this.#cardHTML(card, id)).join('')}
                        </div>
                        <button class="add-card" data-column="${id}">+ Add Card</button>
                    </div>
                `).join('')}
            </div>
        `;
        this.#attachDragEvents();
        this.#attachButtonEvents();
    }

    #cardHTML(card, columnId) {
        return `
            <div class="card" draggable="true" data-id="${card.id}" data-column="${columnId}">
                <p>${card.text}</p>
                <button class="delete-card" data-id="${card.id}" data-column="${columnId}">×</button>
            </div>
        `;
    }

    #attachDragEvents() {
        // Drag start
        this.#el.addEventListener('dragstart', (e) => {
            const card = e.target.closest('.card');
            if (!card) return;
            this.#draggedCard = { id: card.dataset.id, fromColumn: card.dataset.column };
            card.classList.add('dragging');
            e.dataTransfer.effectAllowed = 'move';
        });

        // Drag over column drop zones
        this.#el.addEventListener('dragover', (e) => {
            e.preventDefault();
            e.dataTransfer.dropEffect = 'move';
            const zone = e.target.closest('.cards-container');
            document.querySelectorAll('.cards-container').forEach(z => z.classList.remove('drag-over'));
            zone?.classList.add('drag-over');
        });

        // Drop on column
        this.#el.addEventListener('drop', (e) => {
            e.preventDefault();
            const zone = e.target.closest('.cards-container');
            document.querySelectorAll('.cards-container').forEach(z => z.classList.remove('drag-over'));

            if (!zone || !this.#draggedCard) return;
            const toColumn   = zone.dataset.column;
            const { id, fromColumn } = this.#draggedCard;

            if (fromColumn === toColumn) return;

            const fromCards = this.#state.columns[fromColumn].cards;
            const cardIndex = fromCards.findIndex(c => c.id === id);
            const [card]    = fromCards.splice(cardIndex, 1);
            this.#state.columns[toColumn].cards.push(card);

            this.#save();
            this.#render();
            this.#draggedCard = null;
        });

        // Drag end cleanup
        this.#el.addEventListener('dragend', () => {
            document.querySelectorAll('.card').forEach(c => c.classList.remove('dragging'));
            document.querySelectorAll('.cards-container').forEach(z => z.classList.remove('drag-over'));
        });
    }

    #attachButtonEvents() {
        this.#el.addEventListener('click', (e) => {
            // Add card
            if (e.target.matches('.add-card')) {
                const text = prompt('Card title:')?.trim();
                if (!text) return;
                const col = e.target.dataset.column;
                this.#state.columns[col].cards.push({ id: `card-${Date.now()}`, text });
                this.#save();
                this.#render();
            }

            // Delete card
            if (e.target.matches('.delete-card')) {
                const { id, column } = e.target.dataset;
                const cards = this.#state.columns[column].cards;
                const idx = cards.findIndex(c => c.id === id);
                cards.splice(idx, 1);
                this.#save();
                this.#render();
            }
        });
    }
}

new KanbanBoard(document.getElementById('app'));
```

---

### Q6. Build a Custom Modal Manager

```javascript
// Requirements:
// - Open/close modals with animation
// - Stack multiple modals (nested)
// - Keyboard Escape to close topmost
// - Focus trap inside modal
// - Scroll lock on body

class ModalManager {
    #stack = [];
    #scrollY = 0;

    open(content, { title, onClose, closeOnOverlay = true } = {}) {
        const id = `modal-${Date.now()}`;

        const overlay = document.createElement('div');
        overlay.className = 'modal-overlay';
        overlay.setAttribute('role', 'dialog');
        overlay.setAttribute('aria-modal', 'true');
        overlay.setAttribute('aria-labelledby', `${id}-title`);

        overlay.innerHTML = `
            <div class="modal" id="${id}">
                <div class="modal-header">
                    <h2 id="${id}-title">${title || ''}</h2>
                    <button class="close-btn" aria-label="Close modal">×</button>
                </div>
                <div class="modal-body">${typeof content === 'string' ? content : ''}</div>
            </div>
        `;

        if (typeof content !== 'string') {
            overlay.querySelector('.modal-body').appendChild(content);
        }

        document.body.appendChild(overlay);
        this.#lockScroll();

        // Animate in
        requestAnimationFrame(() => overlay.classList.add('open'));

        // Focus first focusable element
        const focusable = overlay.querySelectorAll('button, input, [tabindex]:not([tabindex="-1"])');
        focusable[0]?.focus();

        const close = () => this.#close(overlay, onClose);

        overlay.querySelector('.close-btn').addEventListener('click', close);
        if (closeOnOverlay) {
            overlay.addEventListener('click', (e) => {
                if (e.target === overlay) close();
            });
        }

        this.#stack.push({ overlay, close, focusable });
        return close; // return close function for programmatic closing
    }

    #close(overlay, onClose) {
        overlay.classList.remove('open');
        overlay.addEventListener('transitionend', () => {
            overlay.remove();
            this.#stack = this.#stack.filter(m => m.overlay !== overlay);
            if (!this.#stack.length) this.#unlockScroll();
            // Restore focus to previous modal if stacked
            const prev = this.#stack[this.#stack.length - 1];
            prev?.focusable[0]?.focus();
            onClose?.();
        }, { once: true });
    }

    #lockScroll() {
        if (this.#stack.length === 0) {
            this.#scrollY = window.scrollY;
            document.body.style.cssText = `overflow: hidden; position: fixed; top: -${this.#scrollY}px; width: 100%;`;
        }
    }

    #unlockScroll() {
        document.body.style.cssText = '';
        window.scrollTo(0, this.#scrollY);
    }

    closeTop() {
        this.#stack[this.#stack.length - 1]?.close();
    }
}

// Global Escape key handler
const modal = new ModalManager();
document.addEventListener('keydown', (e) => {
    if (e.key === 'Escape') modal.closeTop();
});

// Usage
modal.open('<p>Hello from modal!</p>', {
    title: 'Welcome',
    onClose: () => console.log('Modal closed')
});
```

---

## 📋 Machine Coding Evaluation Checklist

Before submitting, verify:

| Criterion | Checklist |
|---|---|
| **Correctness** | All requirements met, edge cases handled |
| **Code Quality** | Meaningful names, no magic numbers, clean logic |
| **Modularity** | Separated concerns, easy to extend |
| **Accessibility** | ARIA roles, keyboard support, screen reader labels |
| **Performance** | Debounced events, no unnecessary re-renders, event delegation |
| **Error Handling** | Network errors, empty states, invalid input |
| **No Memory Leaks** | Removed event listeners, cleared timers |
| **Persistence** | localStorage/URL sync where needed |
