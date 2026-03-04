# 🗣️ Theory — Deployment & DevOps for React
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What happens when you run `npm run build` for a React app?"

> *"Running the build command kicks off a production compilation process. For a Vite app, esbuild processes all your TypeScript and JSX, resolving imports and transforming syntax. Then Rollup bundles everything — splitting the code into chunks based on your dynamic imports. Each chunk gets a content hash in its filename, like main.a3f5b2.js — this enables long-term caching because the filename changes only when the content changes. CSS is also extracted and minified. The output goes into a dist or build folder and is entirely static — just HTML, CSS, JavaScript, and assets. Nothing requires Node.js to serve — you can put these files on any static hosting service. The key characteristics of a production build versus development: all console logging is stripped, minification removes whitespace and shortens variable names, dead code is tree-shaken away, and React's development warnings are eliminated. The result is typically 60-80% smaller than development code."*

---

## Q: "What is the SPA routing problem with static hosting and how do you fix it?"

> *"A React SPA has one entry point — index.html. React Router handles routing client-side. The problem: if you navigate to /user/profile in the browser, the React app handles it and the URL looks right. But if you refresh the page or share that URL, the static server tries to find a file at /user/profile — which doesn't exist. It returns a 404. The fix is to configure your hosting to serve index.html for all routes — let the server always return index.html and let React Router handle the routing. On Netlify, you add a redirects file: /* to /index.html with a 200 status. On Nginx, it's the try_files directive — try the file, try the folder, fall back to index.html. On Vercel it's done automatically for framework projects. On AWS S3 with CloudFront, you configure a custom error response for 403 and 404 to return index.html with a 200 status. Every static host has this configuration option, but the default behavior is to 404 on unknown paths."*

---

## Q: "What is a CI/CD pipeline and what does it typically do for a React app?"

> *"CI/CD stands for Continuous Integration and Continuous Deployment. The CI part — Continuous Integration — means every push to a branch automatically runs a set of checks: install dependencies, run linting, run tests, check TypeScript types, and build the project. If any step fails, the developer is notified immediately. This prevents broken code from reaching main. The CD part — Continuous Deployment — means when a push lands on the main branch and CI passes, the deployment happens automatically without anyone pressing a button. For a React app, that means the build output gets pushed to your hosting — Vercel, Netlify, AWS S3. Tools like GitHub Actions make this declarative — you write a YAML file describing the pipeline steps, and GitHub runs them on their servers. The practical effect: developers can merge confidently knowing tests pass, and users see changes live within minutes of merging. The culture shift it enables: small, frequent deployments rather than big risky releases."*

---

## Q: "What are the differences between Vercel, Netlify, and AWS S3 for hosting React?"

> *"Vercel is optimized for Next.js — it was built by the Next.js creators — but hosts any static site or serverless function. It has the best developer experience: connect a GitHub repo and every push previews automatically, production deploys on merge to main. Free tier is generous. Edge network is fast globally. For a pure Vite or CRA app, Vercel is my first choice for simplicity. Netlify is similar — great DX, automatic preview deployments, generous free tier. It had form handling and identity features built in which were unique, but Vercel caught up. Both are ideal for small to medium projects and indie developers. AWS S3 with CloudFront is what you'd use in an enterprise setting where you need more control, compliance requirements, or integration with other AWS services. It's significantly more complex to set up — you need to configure the bucket, set up CloudFront distribution, configure the SPA routing workaround, set up HTTPS — but infinitely flexible and very cheap at scale for static files."*

---

## Q: "How do you use Docker to containerize a React app and why would you?"

> *"Containerizing a React app with Docker makes sense in a few scenarios: when your team wants consistent build environments, when you're deploying to Kubernetes or other container orchestration, or when the React app is just one service among many in a docker-compose setup. A React app is static files, so the Dockerfile typically uses a multi-stage build. Stage one uses a Node image to install dependencies and run the build — npm ci then npm run build. Stage two uses a lean Nginx image and just copies the build output from stage one. The result is a tiny image — maybe 20MB — that serves the static files with Nginx. Nginx also handles the SPA routing problem with a try_files configuration. You'd expose port 80 inside the container. The multi-stage approach is important: if you used a single Node-based image, you'd ship all your dev dependencies and build tooling into the final image unnecessarily. The two-stage approach keeps the production image minimal."*

---

## Q: "What environment variables can you safely include in a React app?"

> *"The main rule: anything in a React client-side environment variable is visible to everyone in the browser — treat it as public. The browser downloads your JavaScript bundle and anyone can read it. So you should only include values that are designed to be public: API base URLs, your Stripe or PayPal publishable key — not the secret key, analytics IDs like your Google Analytics tracking ID, feature flag configuration, application version numbers. What you should never include: private API keys, database connection strings, signing secrets, OAuth client secrets, any credential that gives server-level access. For values that need to be secret, create a backend endpoint. The frontend calls your backend API; your backend uses the secret to call a third party. The frontend never touches the secret. In practice: Vite variables are prefixed VITE_ and accessed via import.meta.env. CRA variables are prefixed REACT_APP_ and accessed via process.env. Anything without the prefix is excluded from the bundle entirely."*
