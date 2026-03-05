# Microfrontend Fundamentals & Core Concepts

These questions focus on the basic definitions, pros/cons, and foundational knowledge expected in interviews for service-based IT companies or consultancy roles.

## 1. What is a Microfrontend? Explain it in simple terms.

**Answer:**
A Microfrontend architecture is a design approach where a monolithic frontend application is broken down into smaller, individual, and semi-independent "micro-apps."

Simply put, instead of having one massive codebase for your entire website (like an e-commerce site where the cart, product catalog, and user profile are all built together), you split them up. One team builds only the "Cart" micro-app, and another builds the "Product Catalog" micro-app.

These separate micro-apps are then stitched together into a single "Host" or "Shell" application, so to the end-user, it looks and feels like one seamless website.

## 2. What are the main advantages of using a Microfrontend architecture?

**Answer:**
The primary benefits are related to organizational scaling and independent deployments:

1.  **Independent Deployments:** The most significant advantage. The "Checkout" team can deploy an update to their checkout flow 10 times a day without coordinating with the "User Profile" team, and without having to redeploy the entire application.
2.  **Team Autonomy:** Teams can operate independently, choosing their own release schedules and internal architectures (as long as they adhere to the overall system contracts).
3.  **Technology Agnosticism (Incremental Upgrades):** It allows teams to upgrade frameworks incrementally. For example, if a company wants to migrate from an old Angular.js app to React, they can rewrite feature by feature (Microfrontend by Microfrontend) without a massive "big bang" rewrite. (Note: While possible, mixing frameworks is generally discouraged for performance reasons unless absolutely necessary for migration).
4.  **Fault Isolation:** If a bug crashes the "Recommendations" Microfrontend, it doesn't necessarily crash the entire application; the rest of the page (like the checkout button) might still work.

## 3. What are the major disadvantages or challenges of using Microfrontends?

**Answer:**
Microfrontends solve organizational problems but introduce significant technical complexity.

1.  **Operational Complexity:** Setting up the CI/CD pipelines, routing, and deployment strategies for 10 small applications is much harder than setting it up for one monolith.
2.  **Performance Overheads (Bundle Size):** If not managed carefully, a user might end up downloading the same libraries (like React or Lodash) multiple times, causing slow load times.
3.  **Inconsistent UI/UX:** Because different teams build different parts of the UI, they can easily look slightly different (different button styles, fonts) unless a very strict, shared Design System is enforced.
4.  **Complex State Management:** Sharing data (like user session or cart items) between two isolated applications running in the browser is challenging and requires careful architectural planning (avoiding tight coupling).
5.  **Difficult Local Development:** For a developer to run the "whole" application locally to test a feature, they might need to spin up 5 different servers simultaneously.

## 4. How does a Microfrontend architecture differ from Microservices?

**Answer:**
They share the exact same philosophy (breaking up a monolith into smaller, independent pieces), but they apply to entirely different layers of the application stack.

*   **Microservices** apply to the **Backend (Server)**. They split backend APIs, databases, and business logic into smaller services (e.g., an Authentication Service, an Order Service) that communicate via network protocols (HTTP/REST, gRPC, Event Queues).
*   **Microfrontends** apply to the **Frontend (Browser)**. They split the User Interface (HTML, CSS, JavaScript) into smaller applications (e.g., a Header App, a Content App) that are composed together in the user's browser or at the edge server.

## 5. When should you NOT use Microfrontends?

**Answer:**
Microfrontends should be avoided in several scenarios:

*   **Small Teams or Startups:** If there are only 3-5 frontend developers working on a project, the overhead of setting up and maintaining MFE infrastructure is not worth it. A well-structured monolithic application (or a mono-repo) is far more efficient.
*   **Highly Coupled Applications:** If every part of the application relies heavily on the state of every other part (e.g., a highly interactive data visualization dashboard where clicking a chart updates 10 other panels instantly), the isolation of MFEs will make state synchronization a nightmare.
*   **When there's no CI/CD maturity:** If a company struggles to automate deployments for a monolith, introducing 10 independent pipelines for MFEs will result in chaos.
