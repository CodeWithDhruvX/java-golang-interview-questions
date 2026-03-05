# Microfrontend Frameworks & Setup

These questions verify practical knowledge of how to actually build and configure a Microfrontend system, focusing on popular tools like Single-SPA and Webpack Module Federation.

## 1. Which popular tools or frameworks are used to implement a Microfrontend architecture?

**Answer:**
There are several ways to implement MFEs, ranging from simple to complex:

1.  **Webpack 5 Module Federation (WMF):** The most popular modern approach. It's a feature built directly into Webpack 5 that allows different Webpack builds (applications) to dynamically load code from each other at runtime.
2.  **Single-SPA:** A mature, framework-agnostic routing framework specifically designed for MFEs. It handles the lifecycle (mounting, unmounting) of different micro-apps (React, Angular, Vue) within a single host page.
3.  **Iframes:** The oldest and simplest method. A host page simply embeds independent applications using `<iframe>` tags. While providing perfect isolation (CSS and JS), it creates poor user experiences (scrolling issues, cross-frame communication is hard).
4.  **Web Components:** Modern browser native technologies (Custom Elements, Shadow DOM). Each MFE is exported as a custom HTML element (e.g., `<my-cart-app></my-cart-app>`).
5.  **Server-Side Includes (SSI) / Edge Side Includes (ESI):** Composing the page on the server or CDN level before it reaches the browser. Excellent for SEO but harder for highly interactive JS applications.

## 2. Explain the basic concept of "Webpack Module Federation." How is it different from just using an NPM package?

**Answer:**
Webpack Module Federation acts like a runtime NPM package manager directly in the browser.

*   **NPM Packages (Build-Time):** If team A builds an NPM package (`@company/button`) and team B wants to use it, team B must `npm install` it. When team B builds their application, the button code is physically bundled into team B's final JavaScript file. If team A updates the button, team B **must** `npm update`, rebuild, and redeploy their entire application to see the change.
*   **Module Federation (Runtime):** Team A exposes the `Button` via Webpack. Team B's Webpack config points to Team A's deployed URL. The `Button` code is **never** bundled into Team B's code. Instead, when a user visits Team B's website, the browser dynamically fetches the `Button` code directly from Team A's server. If Team A deploys an update, Team B gets it instantly without rebuilding.

## 3. In Single-SPA, what is an "Application Shell," and what are its primary responsibilities?

**Answer:**
The Application Shell (or Host Application) is the central orchestration layer. It is the very first piece of code that the browser downloads and executes.

**Primary Responsibilities:**
1.  **Global Routing:** It listens to browser URL changes and determines which child Microfrontend should be active. For example, if the URL is `/orders`, it knows to load the "Orders App".
2.  **Loading and Lifecycle Management:** It fetches the JavaScript bundles for the Remote Microfrontends, mounts them into the DOM when needed, and crucially, *unmounts* them (destroys them to free memory) when the user navigates away.
3.  **Global State/Auth (Optional but common):** It often handles user authentication (checking tokens) and might expose a global state (like a user context or translation library) to the child applications.
4.  **Shared Layout:** It usually renders the parts of the page that never change, like the main Navigation Header or Footer.

## 4. How do you configure a React application to act as a "Remote" in Webpack Module Federation?

**Answer:**
You need to modify the `webpack.config.js` of the React application by adding the `ModuleFederationPlugin`.

**Basic Configuration Steps:**
1.  **Name:** Give the remote application a unique name (e.g., `app_remote_cart`).
2.  **Filename:** Specify the output filename for the remote entry file, usually `remoteEntry.js`. This is the file the Host will fetch.
3.  **Exposes:** Define the components or modules you want to share with the Host. Use key-value pairs where the key is the alias (e.g., `./CartWidget`) and the value is the path to the local file (e.g., `./src/components/CartWidget`).
4.  **Shared:** Define which libraries (like React or ReactDOM) should be shared so the browser doesn't download them twice.

**Code Example:**
```javascript
const ModuleFederationPlugin = require('webpack/lib/container/ModuleFederationPlugin');

module.exports = {
  // ... other webpack config ...
  plugins: [
    new ModuleFederationPlugin({
      name: 'remoteCart', // The name the Host will use
      filename: 'remoteEntry.js',
      exposes: {
        './CartWidget': './src/components/CartWidget', // What we are sharing
      },
      shared: {
        react: { singleton: true, requiredVersion: '^18.0.0' },
        'react-dom': { singleton: true, requiredVersion: '^18.0.0' }
      }
    }),
  ],
};
```
