# 🧪 06 — Testing React Applications
> **Most Asked in Product-Based Companies** | 🟡 Difficulty: Medium–Hard

---

## 🔑 Must-Know Topics
- Testing philosophy — what to test
- React Testing Library (RTL) — queries and interactions
- Jest setup and configuration
- Mocking modules, API calls, timers
- Testing custom hooks
- Cypress for E2E testing
- Accessibility testing

---

## ❓ Most Asked Questions

### Q1. What is the React Testing Library philosophy?

```jsx
// RTL guiding principle:
// "The more your tests resemble the way your software is used, the more confidence they give you."
// — Kent C. Dodds

// ❌ Enzyme (implementation details testing)
const wrapper = shallow(<Counter />);
wrapper.instance().handleClick();         // testing internals
expect(wrapper.state('count')).toBe(1);   // testing state implementation

// ✅ RTL (user-centric testing)
render(<Counter />);
userEvent.click(screen.getByRole('button', { name: 'Increment' }));
expect(screen.getByText('Count: 1')).toBeInTheDocument();

// RTL query priority (use in this order):
// 1. getByRole         — best for accessibility (button, heading, link)
// 2. getByLabelText    — for form inputs (finds by label association)
// 3. getByPlaceholderText — for input with placeholder
// 4. getByText         — visible text content
// 5. getByDisplayValue — current value of form control
// 6. getByAltText      — image alt text
// 7. getByTitle        — title attribute
// 8. getByTestId       — last resort — data-testid attribute
```

---

### Q2. How do you test async components in RTL?

```jsx
// Testing component with async data fetching
import { render, screen, waitFor } from '@testing-library/react';
import userEvent from '@testing-library/user-event';
import { rest } from 'msw';
import { setupServer } from 'msw/node';
import UserList from './UserList';

// MSW (Mock Service Worker) — intercept requests at network level
const server = setupServer(
  rest.get('/api/users', (req, res, ctx) => {
    return res(ctx.json([
      { id: 1, name: 'Alice', email: 'alice@test.com' },
      { id: 2, name: 'Bob',   email: 'bob@test.com' },
    ]));
  })
);

beforeAll(() => server.listen());
afterEach(() => server.resetHandlers());
afterAll(() => server.close());

describe('UserList', () => {
  it('shows loading state then users', async () => {
    render(<UserList />);

    // Assert loading state
    expect(screen.getByText('Loading...')).toBeInTheDocument();

    // Wait for async render
    const alice = await screen.findByText('Alice');
    expect(alice).toBeInTheDocument();
    expect(screen.getByText('Bob')).toBeInTheDocument();
    expect(screen.queryByText('Loading...')).not.toBeInTheDocument();
  });

  it('shows error state on API failure', async () => {
    server.use(
      rest.get('/api/users', (req, res, ctx) =>
        res(ctx.status(500), ctx.json({ message: 'Server Error' }))
      )
    );

    render(<UserList />);
    expect(await screen.findByText(/error/i)).toBeInTheDocument();
  });
});
```

---

### Q3. How do you test user interactions?

```jsx
import { render, screen } from '@testing-library/react';
import userEvent from '@testing-library/user-event';

describe('LoginForm', () => {
  // Setup user-event (v14 API)
  const user = userEvent.setup();

  it('submits valid form data', async () => {
    const onSubmit = jest.fn();
    render(<LoginForm onSubmit={onSubmit} />);

    // Find elements by accessible role/label
    await user.type(
      screen.getByLabelText('Email'),
      'test@example.com'
    );
    await user.type(
      screen.getByLabelText('Password'),
      'password123'
    );
    await user.click(screen.getByRole('button', { name: 'Login' }));

    expect(onSubmit).toHaveBeenCalledWith({
      email: 'test@example.com',
      password: 'password123'
    });
  });

  it('shows validation error for empty email', async () => {
    render(<LoginForm onSubmit={jest.fn()} />);

    await user.click(screen.getByRole('button', { name: 'Login' }));

    expect(screen.getByText('Email is required')).toBeInTheDocument();
    expect(screen.getByRole('button', { name: 'Login' })).not.toBeDisabled();
  });

  it('disables submit button while loading', async () => {
    const slowSubmit = () => new Promise(resolve => setTimeout(resolve, 1000));
    render(<LoginForm onSubmit={slowSubmit} />);

    await user.type(screen.getByLabelText('Email'), 'a@b.com');
    await user.type(screen.getByLabelText('Password'), 'pass123');
    await user.click(screen.getByRole('button', { name: 'Login' }));

    expect(screen.getByRole('button', { name: 'Logging in...' })).toBeDisabled();
  });
});
```

---

### Q4. How do you test custom hooks?

```jsx
import { renderHook, act } from '@testing-library/react';
import useCounter from './useCounter';

describe('useCounter', () => {
  it('initializes with default value', () => {
    const { result } = renderHook(() => useCounter());
    expect(result.current.count).toBe(0);
  });

  it('initializes with provided value', () => {
    const { result } = renderHook(() => useCounter(10));
    expect(result.current.count).toBe(10);
  });

  it('increments count', () => {
    const { result } = renderHook(() => useCounter());

    act(() => {
      result.current.increment();
    });

    expect(result.current.count).toBe(1);
  });

  it('resets count', () => {
    const { result } = renderHook(() => useCounter(5));

    act(() => {
      result.current.increment();
      result.current.reset();
    });

    expect(result.current.count).toBe(5); // back to initial
  });
});

// Testing hook with context
it('useCart hook adds items', () => {
  const wrapper = ({ children }) => (
    <CartProvider>{children}</CartProvider>
  );

  const { result } = renderHook(() => useCart(), { wrapper });

  act(() => {
    result.current.addItem({ id: 1, name: 'Hat', price: 29 });
  });

  expect(result.current.items).toHaveLength(1);
  expect(result.current.total).toBe(29);
});
```

---

### Q5. How do you mock modules and APIs in Jest?

```jsx
// Mock an entire module
jest.mock('../api/userService');
import * as userService from '../api/userService';

// Type the mock
const mockFetchUser = userService.fetchUser as jest.MockedFunction<typeof userService.fetchUser>;

it('renders user data', async () => {
  mockFetchUser.mockResolvedValue({ id: 1, name: 'Alice', age: 30 });

  render(<UserProfile userId={1} />);

  expect(await screen.findByText('Alice')).toBeInTheDocument();
  expect(mockFetchUser).toHaveBeenCalledWith(1);
});

// Mock implementation
mockFetchUser.mockImplementation((id) => {
  if (id === 1) return Promise.resolve({ name: 'Alice' });
  return Promise.reject(new Error('User not found'));
});

// Reset mocks
afterEach(() => jest.clearAllMocks());

// Mock localStorage
const localStorageMock = {
  getItem:  jest.fn(),
  setItem:  jest.fn(),
  removeItem: jest.fn(),
  clear:    jest.fn(),
};
Object.defineProperty(window, 'localStorage', { value: localStorageMock });

// Spy on functions
const consoleSpy = jest.spyOn(console, 'error').mockImplementation(() => {});
// ... test ...
consoleSpy.mockRestore();
```

---

### Q6. What is Cypress and how do you write E2E tests?

```javascript
// cypress/e2e/login.cy.js
describe('Login Flow', () => {
  beforeEach(() => {
    // Intercept API (no real server needed)
    cy.intercept('POST', '/api/login', {
      statusCode: 200,
      body: { token: 'fake-jwt', user: { id: 1, name: 'Alice' } }
    }).as('loginRequest');

    cy.visit('/login');
  });

  it('successfully logs in with valid credentials', () => {
    cy.get('[data-testid="email-input"]').type('alice@example.com');
    cy.get('[data-testid="password-input"]').type('password123');
    cy.get('[data-testid="login-button"]').click();

    cy.wait('@loginRequest');
    cy.url().should('include', '/dashboard');
    cy.contains('Welcome, Alice').should('be.visible');
  });

  it('shows error for invalid credentials', () => {
    cy.intercept('POST', '/api/login', { statusCode: 401 });

    cy.get('[data-testid="email-input"]').type('wrong@example.com');
    cy.get('[data-testid="password-input"]').type('wrongpass');
    cy.get('[data-testid="login-button"]').click();

    cy.contains('Invalid credentials').should('be.visible');
    cy.url().should('include', '/login');
  });
});
```
