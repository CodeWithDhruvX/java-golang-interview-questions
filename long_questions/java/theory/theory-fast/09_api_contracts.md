# API Versioning, Documentation & Contracts Interview Questions (111-117)

## API Versioning & Strategy

### 111. What are different API versioning strategies?
"There are three main ways.

1.  **URL Versioning**: Put the version in the path (`/v1/users`). It’s explicit and easy to see in logs, but arguably breaks REST principles (a resource shouldn't change just because the version changes).
2.  **Header Versioning**: Send a custom header (`X-API-Version: 1`). It keeps URLs clean but is harder to test in a browser.
3.  **Content Negotiation (Accept Header)**: Clients ask for `application/vnd.mycompany.v1+json`. This is the most 'RESTful' but hardest to implement.

I prefer URL versioning because of developer experience (DX) is simpler."

**Spoken Format:**
"API versioning is like deciding how to label your products in a store.

**URL Versioning** is like putting version number on the product box itself - `/v2/products`. It's obvious and easy to see, but it breaks REST principles because the same URL now returns different data.

**Header Versioning** is like having a small tag on your product that only the scanner can read - `X-API-Version: 2`. The product stays the same but you can evolve it.

The tradeoffs are:

URL versioning is easier for developers and caching, but can lead to messy URLs and confusion about what API does what.

Header versioning is cleaner URLs and better for REST, but invisible to users and harder to debug.

I prefer URL versioning for public APIs where developer experience matters, and header versioning for internal APIs where consistency is key."

### 112. Pros and cons of URL vs header-based versioning?
"**URL Versioning** pros: Easy to cache, easy to bookmark, easy to debug. Cons: Can lead to URL bloat.

**Header Versioning** pros: Keeps your URLs clean (`/users` is always `/users`). Your API surface area looks smaller. Cons: It’s invisible. You can't just copy-paste a URL to a colleague; you have to say 'hit this URL with these headers'. And caching proxies (CDNs) need extra configuration (Vary header) to cache different versions properly."

**Spoken Format:**
"API versioning strategies are like choosing how to organize your store.

**URL Versioning** is like having separate aisles for each product version. It's easy to find what you need, but the store can get cluttered.

**Header Versioning** is like having a single aisle with different products on the same shelf. It looks cleaner, but you need to know what you're looking for.

The key is to balance ease of use with maintainability. URL versioning is great for public APIs where simplicity matters, while header versioning is better for internal APIs where consistency is key."

### 113. What is OpenAPI / Swagger used for?
"It’s the standard for describing REST APIs.

It allows us to write a YAML or JSON file that defines every endpoint: input parameters, output schemas, and error codes.

We use it to generate **interactive documentation** (Swagger UI) so frontend devs can try out the API. And we use it to **auto-generate client SDKs** (in TypeScript, Python, etc.) so we don't have to write fetch code manually. It ensures documentation and code stay in sync."

**Spoken Format:**
"OpenAPI and Swagger are like having a blueprint and construction crew for your API building.

**OpenAPI** is the architectural blueprint - it defines exactly what your API looks like: every endpoint, what parameters it accepts, what it returns, and what errors it can produce.

**Swagger** is like the construction crew that takes your blueprint and automatically builds:
- Interactive documentation so developers can try your API without reading a manual
- Auto-generated client SDKs in multiple languages so teams don't have to write API calls manually
- Code examples and testing tools

The beauty is that when you update your OpenAPI specification, Swagger automatically updates the documentation and SDKs. It's like having a smart blueprint that keeps everyone synchronized.

This prevents the classic problem where documentation gets out of date with the actual API, or where frontend and backend teams have different understandings of how the API works. Everything stays in sync!"

### 114. What is backward compatibility?
"It means newer versions of your service can still support older clients.

If I add a new field `phoneNumber` to the user profile, that’s backward compatible—old clients will just ignore it.
If I *rename* `email` to `emailAddress`, that breaks compatibility—old clients expecting `email` will crash.

Maintaining backward compatibility is critical in microservices so we can deploy services independently without needing a 'big bang' release where everyone updates at once."

**Spoken Format:**
"Backward compatibility is like promising customers that your new product upgrades won't break their existing setup.

Imagine you sell coffee machines to businesses. When you release version 2, you promise that version 1 machines will still work with your coffee pods.

**Adding fields** is like adding a new button to your machine - old machines ignore it, so they keep working.

**Renaming fields** is like changing the coffee pod size - old machines won't recognize the new pods and will break.

**Removing features** is like removing the espresso shot option - customers who relied on it will be upset.

The key is that breaking changes create enemies. I only make them when absolutely necessary, and I give plenty of notice.

In microservices, this is even more critical because you might have 50 different services. You can't coordinate a 'big bang' release across all of them. Backward compatibility lets you upgrade services independently without breaking the whole system!"

### 115. What is a contract-first approach?
"Contract-first means we define the API specification (the OpenAPI YAML) *before* we write any code.

We agree on the contract with the Frontend and Mobile teams. This allows them to start building their UI using mock data (generated from the contract) while we build the backend in parallel.

It prevents 'integration hell' at the end of a sprint where frontend realizes backend built the wrong thing."

**Spoken Format:**
"Contract-first development is like having architects and construction crews working from the same blueprints.

Instead of backend team building something and then telling frontend what they built, everyone agrees on the blueprint first.

The OpenAPI contract defines exactly what the API will do - every endpoint, data structure, error format.

Frontend team can start building their UI immediately using mock data that follows the contract.

Backend team builds the real implementation knowing exactly what frontend expects.

This prevents the classic 'integration hell' where frontend builds their entire application, then at the end discovers that backend built something different.

Both teams work in parallel from day one, using the same blueprint. It's like having architects and construction crews working from the same plans instead of waiting for each other!"

### 116. How do you deprecate an API safely?
"It’s a slow process.

1.  **Mark it**: Add `@Deprecated` in code and headers (`Sunset` header).
2.  **Monitor**: Use logs to see who is still calling it.
3.  **Communication**: Notify consumers (email, Slack) with a deadline.
4.  **Brownouts**: Start failing requests intentionally for short periods (1 minute) to get their attention.
5.  **Shutdown**: Finally remove the code.

You never just delete an endpoint overnight."

**Spoken Format:**
"API deprecation is like retiring a product model - you need to do it carefully to avoid upsetting customers.

The process is like a careful phase-out plan:

**1. Mark it Deprecated** - This is like putting 'Discontinued' label on the product. You're telling customers this model won't be available forever.

**2. Monitor Usage** - Check how many people are still using the old endpoint. This is like tracking sales of the discontinued model.

**3. Communicate** - Send emails, add headers, post notices. It's like sending letters to all customers about the upcoming changes.

**4. Brownouts** - Start failing requests intentionally for short periods to get developers' attention. This is like making the product occasionally malfunction to encourage migration.

**5. Graceful Shutdown** - Finally, return proper error codes and eventually turn off the endpoint. This is like closing the production line for that model.

The key is to never surprise your users. Give them plenty of time and support to migrate. Breaking changes without notice is how you lose customer trust!""

### 117. How do you handle breaking changes in APIs?
"If I absolutely must make a breaking change (like changing a data type), I introduce a **new version** (`/v2/users`).

I keep `/v1/users` running, possibly as a proxy that adapts the old request format to the new internal logic, or just maintaining the old code for a grace period.

This allows clients to migrate at their own pace. Forced upgrades are a reliable way to make enemies."

**Spoken Format:**
"API evolution is like planning city growth - you need to add new roads without destroying existing ones.

When I absolutely must make a breaking change, I introduce a **new version** - like building a new highway system that runs alongside the old one.

The old version (`/v1/users`) keeps working exactly as before, serving existing clients.

The new version (`/v2/users`) can have improved features and better design.

This approach lets different types of traffic coexist:
- Old clients can continue using the old highway
- New clients get the benefits of the improved system
- No forced migration - everyone upgrades when they're ready

It's like having multiple generations of roads in your city. The old roads still work, but new roads have better traffic flow. Everyone can travel at their own pace.

Forced upgrades are like suddenly closing all old roads and telling everyone they must use the new highway immediately - that creates chaos and enemies!

Gradual evolution respects your users' investment and timing.""
