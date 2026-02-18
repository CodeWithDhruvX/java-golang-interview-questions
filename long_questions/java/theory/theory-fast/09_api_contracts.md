# API Versioning, Documentation & Contracts Interview Questions (111-117)

## API Versioning & Strategy

### 111. What are different API versioning strategies?
"There are three main ways.

1.  **URL Versioning**: Put the version in the path (`/v1/users`). It’s explicit and easy to see in logs, but arguably breaks REST principles (a resource shouldn't change just because the version changes).
2.  **Header Versioning**: Send a custom header (`X-API-Version: 1`). It keeps URLs clean but is harder to test in a browser.
3.  **Content Negotiation (Accept Header)**: Clients ask for `application/vnd.mycompany.v1+json`. This is the most 'RESTful' but hardest to implement.

I prefer URL versioning because the developer experience (DX) is simpler."

### 112. Pros and cons of URL vs header-based versioning?
"**URL Versioning** pros: Easy to cache, easy to bookmark, easy to debug. Cons: Can lead to URL bloat.

**Header Versioning** pros: Keeps your URLs clean (`/users` is always `/users`). Your API surface area looks smaller. Cons: It’s invisible. You can't just copy-paste a URL to a colleague; you have to say 'hit this URL with these headers'. And caching proxies (CDNs) need extra configuration (Vary header) to cache different versions properly."

### 113. What is OpenAPI / Swagger used for?
"It’s the standard for describing REST APIs.

It allows us to write a YAML or JSON file that defines every endpoint: input parameters, output schemas, and error codes.

We use it to generate **interactive documentation** (Swagger UI) so frontend devs can try out the API. And we use it to **auto-generate client SDKs** (in TypeScript, Python, etc.) so we don't have to write fetch code manually. It ensures the documentation and code stay in sync."

### 114. What is backward compatibility?
"It means newer versions of your service can still support older clients.

If I add a new field `phoneNumber` to the user profile, that’s backward compatible—old clients will just ignore it.
If I *rename* `email` to `emailAddress`, that breaks compatibility—old clients expecting `email` will crash.

Maintaining backward compatibility is critical in microservices so we can deploy services independently without needing a 'big bang' release where everyone updates at once."

### 115. What is a contract-first approach?
"Contract-first means we define the API specification (the OpenAPI YAML) *before* we write any code.

We agree on the contract with the Frontend and Mobile teams. This allows them to start building their UI using mock data (generated from the contract) while we build the backend in parallel.

It prevents the 'integration hell' at the end of a sprint where the frontend realizes the backend built the wrong thing."

### 116. How do you deprecate an API safely?
"It’s a slow process.

1.  **Mark it**: Add `@Deprecated` in code and headers (`Sunset` header).
2.  **Monitor**: Use logs to see who is still calling it.
3.  **Communication**: Notify consumers (email, Slack) with a deadline.
4.  **Brownouts**: Start failing requests intentionally for short periods (1 minute) to get their attention.
5.  **Shutdown**: Finally remove the code.

You never just delete an endpoint overnight."

### 117. How do you handle breaking changes in APIs?
"If I absolutely must make a breaking change (like changing a data type), I introduce a **new version** (`/v2/users`).

I keep `/v1/users` running, possibly as a proxy that adapts the old request format to the new internal logic, or just maintaining the old code for a grace period.

This allows clients to migrate at their own pace. Forced upgrades are a reliable way to make enemies."
