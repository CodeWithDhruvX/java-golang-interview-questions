# State Management & Communication in Microfrontends

One of the strict rules of Microfrontends is loose coupling. These questions explore how to share data without creating a monolith out of your frontend applications.

## 1. How do you share state between two microfrontends (e.g., a "User Profile" Remote and a "Shopping Cart" Remote)? 

**Answer:**
The golden rule of Microfrontends is: **Avoid sharing state if possible.** If you must share state, keep it as minimal as possible (like simple primitives: IDs or tokens).

If communication is necessary, here are the patterns, from best to worst:

1.  **Custom Events (Pub/Sub) - *Recommended*:**
    *   *Implementation:* Use the native browser API (`window.dispatchEvent(new CustomEvent('ITEM_ADDED', { detail: { id: 1 }}))`).
    *   *Pros:* Extremely decoupled. The Cart MFE emits an event; it doesn't care who listens. The Header MFE listens and updates the cart counter. Tech agnostic (React can talk to Angular).
    *   *Cons:* Can become difficult to track "who fired what" in very large applications without strict documentation conventions.
2.  **URL / Query Parameters - *Recommended for specific cases*:**
    *   *Implementation:* Passing simple identifiers in the URL (e.g., `/checkout?productId=123`).
    *   *Pros:* Trivially easy, bookmarkable, perfectly decoupled.
    *   *Cons:* Only works for very small, non-sensitive data strings.
3.  **Global Store (Zustand, Redux, Context exposed via Shell) - *Use with Extreme Caution*:**
    *   *Implementation:* The Host app initializes a Redux store and passes it down as props or exposes it via a shared context/window object.
    *   *Pros:* Easy to use if you are already familiar with the state library.
    *   *Cons:* Creates tight coupling. All MFEs must now agree on the state library and its structure. If the shape of the global store changes, multiple independent teams must coordinate updates, destroying the benefits of MFEs.
4.  **Backend for Frontend (BFF) / Shared Database:**
    *   *Implementation:* State isn't shared on the client at all. The Cart MFE posts an update to the backend. The Header MFE polls or uses a WebSocket to the backend to get the latest cart count.
    *   *Pros:* The ultimate source of truth remains on the server.
    *   *Cons:* Increased network traffic and latency.

## 2. Suppose an authentication token is required by all Microfrontends. How do you manage authentication and token sharing?

**Answer:**
Authentication is a cross-cutting concern that should be handled centrally by the Host/App Shell to avoid duplicating logic in every MFE.

**Recommended Flow:**
1.  **The Shell Handles Auth:** The App Shell is responsible for checking if the user is authenticated on initial load. If not, it redirects to the login page (or handles OAuth/OIDC flows).
2.  **Token Storage:** The Shell receives the JWT (JSON Web Token) and stores it securely.
    *   *Best practice:* Store tokens in `HttpOnly` securely flagged cookies. This prevents XSS attacks and requires zero Javascript state sharing—the browser automatically attaches the cookie to API requests made by any MFE.
    *   *Alternative:* If APIs require an `Authorization: Bearer <token>` header, the Shell stores the token in memory (or sessionStorage).
3.  **Sharing the Token (If `HttpOnly` cookies aren't used):**
    *   If the token is in memory, the Shell can expose a secure, read-only getter function to the Remotes (e.g., injected via props during mounting, or mapped via a singleton module in Webpack Federation).
    *   Avoid storing sensitive tokens in `localStorage`, as they are highly vulnerable to XSS.

**Handling Expiration:**
The Shell should manage token refresh logic (refresh tokens). Remotes should be entirely agnostic to the refresh process; they simply attempt to make API calls using the currently provided credentials.

## 3. How do you ensure that a Custom Event schema doesn't break when one team updates their MFE?

**Answer:**
When relying on Custom Events (event bus) for cross-MFE communication, establishing a contract is vital to prevent breaking changes.

1.  **Event Schema Definitions:**
    *   Create a shared, strongly-typed internal library (e.g., an NPM package containing TypeScript interfaces or Zod schemas).
    *   Teams define their event payloads: `export interface CartAddedEvent { productId: string; quantity: number; }`
    *   Both the emitting team and the listening team consume this type, ensuring compile-time safety.
2.  **Versioning Events:**
    *   If a breaking change is needed (e.g., changing `productId` to a numeric type), version the event string rather than breaking the existing one.
    *   Old event: `APP/CART/ITEM_ADDED_V1`
    *   New event: `APP/CART/ITEM_ADDED_V2`
    *   The emitting MFE should fire both events temporarily until the listening team has migrated to V2.
3.  **Event Registry:**
    *   Document all cross-boundary events in a central registry (like a Wiki or Backstage portal). Every event must list its owner, payload schema, and intended consumers.
