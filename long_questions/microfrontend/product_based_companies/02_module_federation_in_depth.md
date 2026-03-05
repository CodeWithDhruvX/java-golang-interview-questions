# Webpack Module Federation In-Depth

Webpack 5's Module Federation (WMF) is currently the industry standard for implementing runtime Microfrontends. Product-based companies will test your deep, practical knowledge of this technology.

## 1. Explain exactly how Webpack Module Federation resolves shared dependencies at runtime. What happens if there's a version mismatch?

**Answer:**
Module Federation handles shared dependencies through a mechanism that creates a shared scope during initialization.

**How it works:**
1.  **Initialization:** When the Host application boots up, Webpack initializes a "shared scope". It registers the host's versions of shared libraries (e.g., React 18.2.0) into this scope.
2.  **Remote Loading:** When a Remote MFE is loaded, Webpack injects the currently established shared scope into the Remote's initialization process.
3.  **Resolution Logic:** The Remote checks the shared scope:
    *   If the required library exists and satisfies the Semantic Versioning (SemVer) requirement specified in the Remote's `webpack.config.js` (e.g., `^18.0.0`), the Remote *does not* download its own bundled version. It uses the Host's version.
    *   If the library doesn't exist, or the version in the scope is incompatible (e.g., Host has 17.0.0, Remote needs `^18.0.0`), the Remote will download its own bundled "fallback" version.

**Handling Version Mismatches:**
If a version mismatch occurs, you might end up downloading React twice, causing bundle bloat. Worse, if a library like React (which expects a single instance) is loaded twice, it will throw fatal errors (e.g., Invalid Hook Call Warning).

**Mitigation (The `singleton` flag):**
To prevent this, you enforce the `singleton: true` option in the shared configuration for critical libraries (React, ReactDOM, Vue).
*   If `singleton: true` is set, and a version mismatch occurs, Webpack will **force** the use of the highest available version in the shared scope, and will throw a warning in the console. E.g., if Host provides React 17 and Remote requires React 18 (both set as singletons), it will try to use React 18 (if the remote downloads it), potentially breaking the Host if the Host isn't compatible with 18.
*   To strictly enforce versions and fail fast, you can use `strictVersion: true` alongside `singleton: true`. This throws a runtime error instead of a warning if versions don't match exactly.

## 2. In Module Federation, what is the difference between an `exposes` configuration and a `remotes` configuration? Can an application be both a Host and a Remote?

**Answer:**
Yes, an application can absolutely be both, and this is a common pattern known as "Bi-directional" or "Omni-directional" routing.

*   `exposes`: Defines exactly which local modules/components this Webpack build will make available to *other* applications. It packages these exposed modules into a separate file path (usually `remoteEntry.js`).
    ```javascript
    // Remote App Webpack Config
    exposes: {
      './Button': './src/components/Button',
      './CheckoutFlow': './src/pages/Checkout',
    }
    ```
*   `remotes`: Defines the external URLs or paths where this Webpack build should look for modules exposed by *other* applications. It maps a local alias to the remote `remoteEntry.js` file.
    ```javascript
    // Host App Webpack Config
    remotes: {
      CheckoutApp: 'checkout@http://localhost:3002/remoteEntry.js',
    }
    ```

**Bi-directional Example:**
A `Dashboard_App` might act as a Host to consume a `Widget` from the `Analytics_App` (`remotes: { Analytics: ... }`).
Simultaneously, the `Dashboard_App` might expose its `UserProfile` component (`exposes: { './UserProfile': ... }`) so that the `Analytics_App` can consume it. Therefore, both are Hosts and Remotes simultaneously.

## 3. The `remoteEntry.js` file URL is typically hardcoded in the `webpack.config.js`. How do you handle dynamic URLs for different environments (Dev, Staging, Prod) in a CI/CD pipeline?

**Answer:**
Hardcoding URLs like `http://localhost:3001/remoteEntry.js` in the Webpack config is only good for local development. For production, you cannot hardcode the URLs because they change per environment.

There are three primary ways to handle dynamic remote URLs:

1.  **Environment Variables at Build Time (Simplest, but least flexible):**
    You pass `process.env.REMOTE_URL` into Webpack during the build step.
    *   *Drawback:* This tightly couples the build artifact to a specific environment. A Docker image built for Staging cannot be promoted directly to Production because the URL is baked into the JS bundle. You have to rebuild.
2.  **Promise-based Dynamic Remotes at Runtime (Recommended):**
    Webpack allows you to resolve the remote URL at runtime by returning a Promise in the `remotes` configuration.
    ```javascript
    remotes: {
      app2: `promise new Promise(resolve => {
        // Fetch URL from a global config object injected at runtime via index.html
        const remoteUrl = window._env_.REMOTE_APP2_URL;
        const script = document.createElement('script');
        script.src = remoteUrl;
        script.onload = () => {
          const proxy = { get: (request) => window.app2.get(request), init: (arg) => { window.app2.init(arg) } };
          resolve(proxy);
        }
        document.head.appendChild(script);
      })`
    }
    ```
    *   *How it works:* You inject a small configuration script into the Host's `index.html` (e.g., via a Kubernetes ConfigMap or a server-rendered `<script>` tag). Webpack uses this global variable to fetch the remote at runtime. This allows building once and deploying anywhere.
3.  **Module Federation Dashboard / Dynamic Remote Plugins:**
    Using specialized plugins (like `@module-federation/node` or Medusa) that act as registries to fetch the latest remote URLs dynamically from a centralized server.

## 4. You deploy an update to a Remote MFE, but the Host application is still serving the old version. What went wrong, and how do you fix it?

**Answer:**
This is almost always a caching issue involving the `remoteEntry.js` file.

**The Problem:**
1.  The Host application requests `app2@https://cdn.com/app2/remoteEntry.js`.
2.  The browser (or a CDN like Cloudflare/Cloudfront) aggressively caches `remoteEntry.js` because it has a static filename.
3.  Even though you deployed a new version of App2, the Host continues to load the cached `remoteEntry.js`, which points to the old, hashed JavaScript chunks.

**The Fix:**
You must implement a strict Cache-Control policy specifically for the `remoteEntry.js` file.

1.  **Cache-Control Header:** Configure your CDN or S3 bucket to serve `remoteEntry.js` with `Cache-Control: no-cache, no-store, must-revalidate`. This forces the browser to check the server for a new version on every request.
2.  **Chunk Hashing:** The actual application chunks (e.g., `src_components_Button_js.[hash].js`) *should* be heavily cached long-term (`Cache-Control: public, max-age=31536000`), because their filenames contain unique hashes that change on every build.
3.  **App Shell Check:** Because `remoteEntry.js` is tiny (usually a few KBs), forcing the browser to fetch it fresh does not impact performance significantly. Only if the hash inside `remoteEntry.js` points to a new chunk will the browser download the new heavy JS files.
