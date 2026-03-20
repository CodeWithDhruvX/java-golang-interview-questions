# 🔴 DevOps, CI/CD & Build Tools for Angular

> 🚀 **Product-Based companies** (Swiggy, Zepto, startups): Full CI/CD pipeline, Docker, Nx
> 🏭 **Service-Based companies** (TCS, Wipro): Build basics, deployment, ng build configurations
>
> 🎯 **Experience Level:** 🟡 **Mid-Level (2–4 yrs)** | 🔴 **Senior (5+ yrs)**

---

### 1. What is the Angular build process? 🟢 | 🏭

"Angular's build process transforms TypeScript + HTML templates + SCSS into optimized browser assets:

```
TypeScript + HTML Templates + SCSS
         ↓
   Angular Compiler (ngtsc/Ivy)      ← Compiles templates to TypeScript
         ↓
   TypeScript Compiler (tsc)         ← Type checking
         ↓
   Bundler (Webpack / esbuild)       ← Bundle + tree shake + minify
         ↓
   dist/app/
     ├── main.js                     ← App code
     ├── polyfills.js                ← Browser compatibility
     ├── runtime.js                  ← Webpack runtime
     ├── styles.css                  ← Compiled styles
     └── chunk-xxx.js               ← Lazy-loaded chunks
```

```bash
ng build                         # Development build (unoptimized)
ng build --configuration=production  # Production build (AOT + minify + tree shake)
```"

#### In Depth
Angular 16+ uses **esbuild** as the default bundler (replacing Webpack for most scenarios) via `@angular-devkit/build-angular:application`. esbuild is **10–100x faster** than Webpack because it's written in Go and uses native OS threads. Build times for large apps dropped from 60–120 seconds to 5–15 seconds. However, some Webpack-specific customizations (module federation) still require Webpack via the `browser` builder.

---

### 2. What is `angular.json` and what can you configure? 🟡 | 🏭

"`angular.json` is the Angular workspace configuration file that controls **how ng CLI commands build and serve** the project.

Key sections:
```json
{
  \"projects\": {
    \"my-app\": {
      \"architect\": {
        \"build\": {
          \"builder\": \"@angular-devkit/build-angular:application\",
          \"options\": {
            \"outputPath\": \"dist/my-app\",
            \"index\": \"src/index.html\",
            \"browser\": \"src/main.ts\",
            \"styles\": [\"src/styles.scss\"],
            \"assets\": [\"src/favicon.ico\", \"src/assets\"],
            \"budgets\": [
              { \"type\": \"initial\", \"maximumWarning\": \"500kb\", \"maximumError\": \"1mb\" },
              { \"type\": \"anyComponentStyle\", \"maximumWarning\": \"4kb\" }
            ]
          },
          \"configurations\": {
            \"production\": {
              \"optimization\": true,
              \"outputHashing\": \"all\",   // Cache-busting file names
              \"sourceMap\": false,
              \"namedChunks\": false
            },
            \"staging\": {
              \"fileReplacements\": [
                { \"replace\": \"src/environments/environment.ts\",
                  \"with\": \"src/environments/environment.staging.ts\" }
              ]
            }
          }
        }
      }
    }
  }
}
```"

#### In Depth
The **`budgets`** configuration is critical for production discipline. It causes `ng build` to fail if any bundle exceeds the `maximumError` threshold. This acts as a **CI gate** preventing accidental bundle bloat from sneaking into production. I set `initial` budget at 500KB warning / 1MB error, and per-component style budget at 4KB — this keeps the app lean and forces conscious decisions about which libraries to include.

---

### 3. How to set up a CI/CD pipeline for Angular? 🟡 | 🏭🚀

"A typical Angular CI/CD pipeline for **GitHub Actions**:

```yaml
# .github/workflows/angular-ci.yml
name: Angular CI/CD

on:
  push:
    branches: [main, develop]
  pull_request:
    branches: [main]

jobs:
  build-and-test:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-node@v4
        with:
          node-version: '20'
          cache: 'npm'

      - name: Install dependencies
        run: npm ci

      - name: Lint
        run: npm run lint

      - name: Unit tests
        run: npm run test -- --no-watch --no-progress --code-coverage

      - name: Build production
        run: npm run build -- --configuration=production

      - name: Upload coverage to Codecov
        uses: codecov/codecov-action@v3
        with:
          files: ./coverage/lcov.info

  deploy:
    needs: build-and-test
    if: github.ref == 'refs/heads/main'
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v4
      - name: Deploy to Firebase Hosting
        run: |
          npm ci
          npm run build -- --configuration=production
          npx firebase deploy --only hosting
        env:
          FIREBASE_TOKEN: ${{ secrets.FIREBASE_TOKEN }}
```"

#### In Depth
For large Angular projects, I optimize CI time with:

1. **Nx affected** — Only test/build projects changed by a PR: `nx affected --target=test`
2. **Caching** — Cache `node_modules` and Nx incremental build cache between runs
3. **Parallelization** — Run tests for multiple apps in parallel: `nx run-many --target=test --parallel=3`
4. **Docker-based builds** — Deterministic builds regardless of CI environment

---

### 4. How to Dockerize an Angular app? 🟡 | 🏭🚀

"A two-stage Docker build keeps the final image small:

```dockerfile
# Stage 1: Build
FROM node:20-alpine AS build
WORKDIR /app
COPY package*.json ./
RUN npm ci --only=production
COPY . .
RUN npm run build -- --configuration=production

# Stage 2: Serve with nginx
FROM nginx:alpine
COPY --from=build /app/dist/my-app/browser /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
```

```nginx
# nginx.conf — critical for Angular Router (HTML5 History API)
server {
  listen 80;
  root /usr/share/nginx/html;
  index index.html;

  # Serve Angular for all non-file routes
  location / {
    try_files $uri $uri/ /index.html;
  }

  # Cache static assets aggressively
  location ~* \.(js|css|png|jpg|ico)$ {
    expires 1y;
    add_header Cache-Control "public, immutable";
  }
}
```

```bash
docker build -t my-angular-app:latest .
docker run -p 8080:80 my-angular-app:latest
```"

#### In Depth
The `try_files $uri $uri/ /index.html` nginx config is **mandatory** for Angular apps using the HTML5 History API (`PathLocationStrategy`). Without it, refreshing at `/products/123` returns a 404 because nginx looks for a file literally at that path. Angular's client-side routing only works after `index.html` is loaded. The `/index.html` fallback ensures Angular bootstraps from any route, then takes over routing client-side.

---

### 5. What is Nx and how does it help Angular? 🔴 | 🚀

"**Nx** is a monorepo tool that extends Angular CLI for **large-scale development**. It manages multiple apps and libraries in a single repository.

```
myco/                              # Nx workspace root
├── apps/
│   ├── customer-portal/           # Angular app
│   ├── admin-panel/               # Another Angular app
│   └── api/                       # NestJS backend
├── libs/
│   ├── products/
│   │   ├── data-access/           # API + NgRx store
│   │   ├── ui/                    # Product components
│   │   └── util/                  # Types, utils
│   ├── shared/
│   │   ├── ui/                    # Shared UI components (design system)
│   │   └── util/                  # Shared utilities, types
│   └── auth/
│       └── data-access/           # Auth service
```

Key Nx benefits:
- `nx affected:test` — Only test what changed (drastically faster CI)
- `nx graph` — Visualize dependency graph
- `nx generate` — Custom generators for your team's conventions
- `@nx/enforce-module-boundaries` — Lint rule preventing architectural violations"

#### In Depth
Nx's **computation caching** is its killer feature. When you run `nx test shared-ui` and the code hasn't changed since last run, Nx returns the **cached result instantly** — no re-running tests. This cache can be shared across the team via **Nx Cloud** — if a colleague already tested the same commit, your CI pull the result from the cache rather than running the tests again. In large codebases, this reduces CI times from 30–60 minutes to 2–5 minutes for typical PRs.

---

### 6. How to optimize Angular production build? 🟡 | 🏭🚀

"Production optimization checklist:

```bash
# Standard production build (covers most optimizations automatically)
ng build --configuration=production
```

**What's included automatically in production:**
| Optimization | Effect |
|---|---|
| AOT compilation | No compiler in browser |
| Tree shaking | Remove unused code |
| Minification (Terser) | Smaller JS files |
| CSS minification | Smaller style files |
| Output hashing | Cache-busting filenames |
| Build optimizer | Angular-specific tree shaking |

**Additional optimizations I apply manually:**
```typescript
// 1. Enable Gzip/Brotli at server level (nginx/firebase)
// 2. Enable differential loading (ES5/ES2015 targets) — automatic in Angular 13+
// 3. Set strict budgets in angular.json
// 4. Analyze bundle: npx webpack-bundle-analyzer dist/stats.json
// 5. Lazy load heavy libraries (chart.js, PDF generator)
// 6. Use CDN for assets and fonts
// 7. Service Worker for caching: ng add @angular/pwa
```"

#### In Depth
**Brotli compression** (`br`) is the most impactful server-side optimization: Angular's production bundles compress 15–20% smaller with Brotli vs Gzip. Almost all modern browsers support Brotli. Serving a 300KB `main.js` Brotli-compressed over the wire as 80KB dramatically improves FCP and TTI on mobile networks. I configure nginx to serve pre-compressed files: `gzip_static on; brotli_static on;` — Angular CLI can be configured to output `.gz` and `.br` files during build via esbuild options.

---

### 7. What is environment configuration in Angular? 🟢 | 🏭

"Angular supports **environment-specific configurations** via `environment.ts` files:

```typescript
// src/environments/environment.ts (development)
export const environment = {
  production: false,
  apiBaseUrl: 'http://localhost:3000/api',
  wsUrl: 'ws://localhost:3001',
  featureFlags: {
    darkMode: true,
    betaFeatures: true,
  },
  logLevel: 'debug'
};

// src/environments/environment.prod.ts (production)
export const environment = {
  production: true,
  apiBaseUrl: 'https://api.myapp.com/v2',
  wsUrl: 'wss://ws.myapp.com',
  featureFlags: {
    darkMode: true,
    betaFeatures: false,
  },
  logLevel: 'error'
};
```

Angular replaces `environment.ts` with environment-specific file at build time:
```json
// angular.json
\"fileReplacements\": [
  {
    \"replace\": \"src/environments/environment.ts\",
    \"with\": \"src/environments/environment.prod.ts\"
  }
]
```"

#### In Depth
For **sensitive values** (API keys, secrets), I never put them directly in `environment.ts` files (which are committed to Git). Instead, I use `environment.ts` for build-time references like API URLs and use **runtime environment injection** for secrets: the server injects `window.ENV_VARS = { apiKey: '...' }` from environment variables (Docker/Kubernetes secrets) into the HTML before serving it. This way, secrets live in the deployment platform, not the codebase.

---

### 15. What is the specific command used to update Angular version and dependencies? 🟢 | 🏭

"The specific command used to update Angular version and dependencies is:

```bash
ng update @angular/core @angular/cli
```

This command:
- Updates `@angular/core` and `@angular/cli` to their latest compatible versions
- Automatically updates related dependencies (like `@angular/common`, `@angular/forms`, etc.)
- Runs automated migrations for breaking changes
- Updates `package.json` and `angular.json` configurations

For updating specific packages:
```bash
# Update individual packages
ng update @angular/material
ng update @ngrx/store

# Update all packages to latest compatible versions
ng update --all
```

Before updating, always:
1. Commit your code to version control
2. Check the Angular update guide for breaking changes
3. Run `ng update --dry-run` to preview changes
4. Test thoroughly after update"

#### In Depth
The `ng update` command uses **Angular DevKit** schematics to automate migrations. When updating major versions, it can:
- Refactor deprecated APIs (like `HttpModule` → `HttpClientModule`)
- Update template syntax (like `[hidden]` → `*ngIf`)
- Modify `angular.json` configuration changes
- Add missing imports and remove deprecated ones

For **complex updates** (like Angular 14→16 with standalone components), I run updates incrementally:
```bash
ng update @angular/core@14 @angular/cli@14  # First to latest 14.x
ng update @angular/core@15 @angular/cli@15  # Then to 15.x
ng update @angular/core@16 @angular/cli@16  # Finally to 16.x
```

This approach makes debugging easier and reduces risk of breaking changes."
