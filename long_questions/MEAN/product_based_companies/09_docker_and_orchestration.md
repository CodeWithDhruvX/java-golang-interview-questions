# Docker and Orchestration (Product-Based Companies)

Modern backend development rarely involves copying files straight to a server. Applications are containerized for consistency and orchestrated for scale. You must understand Docker fundamentals and how it applies to Node/MongoDB.

## Docker Fundamentals

### 1. What is Docker, and why is it preferred over Virtual Machines?
*   **Virtual Machines (VMs)**: Emulate an entire hardware system, requiring a full, hefty Guest Operating System (Linux/Windows) on top of the Host OS for *each* VM. Heavy resource overhead, slow startup.
*   **Docker (Containers)**: Virtualize the OS at the application level. Multiple containers share the same Host OS Kernel. They are lightweight, start in milliseconds, port perfectly across environments ("it works on my machine" guarantee), and use very little CPU/RAM overhead.

### 2. Describe an optimized `Dockerfile` for a Node.js API.
An optimized Dockerfile utilizes **Multi-stage builds** and caching, and limits security risks.

```dockerfile
# Stage 1: Build
FROM node:18-alpine AS builder
# Set working directory
WORKDIR /app
# Copy package files FIRST to leverage Docker layer caching
COPY package*.json ./
# Install ALL dependencies (including devDependencies for building/TypeScript)
RUN npm ci
# Copy remaining source code
COPY . .
# Run build script (e.g., compile TypeScript to pure JS)
RUN npm run build

# Stage 2: Production
FROM node:18-alpine
WORKDIR /app
# Only copy package files again
COPY package*.json ./
# Install ONLY production dependencies (keeps image size tiny)
RUN npm ci --only=production
# Copy ONLY the built artifacts from the 'builder' stage
COPY --from=builder /app/dist ./dist
# Switch to non-root user for security
USER node
EXPOSE 3000
# Run the built output
CMD ["node", "dist/index.js"]
```

### 3. What is `.dockerignore` and why is it critical in Node.js?
Like `.gitignore`, it tells Docker which files/folders to ignore when executing a `COPY . .` command.
It is **critical** to ignore `node_modules`. If you copy your local `node_modules` into a Linux container, native C++ bindings (like `bcrypt` or `node-sass` built for your Mac/Windows machine) will instantly crash when executed inside the Linux container. Let Docker install dependencies internally.

## Advanced Containerization concepts

### 4. How does `docker-compose` fit into the MERN/MEAN stack?
`docker-compose` is a tool for defining and running multi-container Docker applications locally using a YAML file.
Instead of manually running `docker run` commands for Node, MongoDB, and Redis, you define them all in `docker-compose.yml`.
*   A single `docker-compose up` will spin up the database, cache, and API, place them on the same internal Docker network (so they can communicate using service names like `mongodb://mongo:27017` instead of IP addresses), and handle port mapping.

### 5. In a distributed, containerized environment (like Kubernetes or AWS ECS), how do you persist MongoDB data?
Containers are ephemeral (temporary). If a MongoDB container crashes and is restarted by Docker/Kubernetes, all data inside it is wiped out.
*   **Volumes**: You must attach external storage Volumes to the container. The data lives on the physical host machine's drive (or a cloud block storage like AWS EBS), and is mounted into the container at runtime. If the container dies, the data persists on the volume and is attached to the replacement container.
*   *(Note: In large enterprise environments, it is often best practice to use managed database services like MongoDB Atlas or AWS DocumentDB rather than hosting DBs in containers yourself, to offload backup, scaling, and patching responsibilities).*
