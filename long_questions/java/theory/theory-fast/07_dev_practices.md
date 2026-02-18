# Git, Linux & Debugging Interview Questions (91-100)

## Git & Development Practices

### 91. Difference between `merge` and `rebase`?
"They both integrate changes from one branch to another, but the history looks different.

**Merge** creates a new 'merge commit'. It preserves the history exactly as it happened. If you look at the graph, you see branches diverging and coming back together. It’s 'true' history but can get messy.

**Rebase** rewrites history. It takes your commits from the feature branch and replays them on top of the master branch. The result is a perfectly linear history, which is much cleaner to read.

In my team, we rebase local branches to keep them clean, but always merge (squash merge) when bringing feature branches into master."

### 92. What is a pull request?
"A Pull Request (PR) is a request to merge your code changes into the main codebase.

It’s the core of our code review process. I push my branch, open a PR, and my teammates review it. They check for bugs, style issues, and logic gaps. The CI pipeline also runs tests on the PR automatically.

Once everything is green and approved, it gets merged. It’s a quality gate."

### 93. How do you resolve merge conflicts?
"First, I don't panic. A conflict just means Git doesn't know which version of a line to pick.

I usually run `git rebase main` on my feature branch. Git stops at the conflict. I open the file in my IDE (IntelliJ), which shows a 3-way view: 'Yours', 'Theirs', and 'Result'.

I pick the changes I want (or combine them), save the file, `git add`, and `git rebase --continue`. It’s part of the job."

### 94. What is CI/CD?
"CI (Continuous Integration) is the practice of merging code frequently. Every time I push, a build server (Jenkins/GitHub Actions) compiles the code and runs unit tests. If it fails, I fix it immediately.

CD (Continuous Delivery/Deployment) takes it further. Once the build passes, it automatically deploys the app to a staging environment (Delivery) or even straight to production (Deployment).

It replaces the old 'release night' panic with a boring, automated process."

## Linux, Networking & Debugging

### 95. Common Linux commands you use?
"I live in the terminal.

-   `grep`: Searching logs (`grep "ERROR" app.log`).
-   `tail -f`: Watching logs in real-time.
-   `ps -ef | grep java`: Checking if my process is running.
-   `netstat` or `lsof`: Checking ports.
-   `curl`: Testing APIs.
-   `chmod`: Changing permissions.
-   `top` / `htop`: Checking CPU/Memory usage."

### 96. How do you check if a port is open?
"On a remote server, I use `telnet host port` or `nc -zv host port` (Netcat).

If I'm checking strictly on the local machine to see what's listening, I use `netstat -tulpn` or `lsof -i :8080`.

This is usually the first step when a service fails to start—checking if the port is already taken."

### 97. Difference between HTTP and HTTPS?
"HTTPS is simply HTTP over a secure encrypted connection (SSL/TLS).

With HTTP, data is sent in plaintext. Anyone on the network (Wi-Fi) can sniff my password.

With HTTPS, the data is encrypted using a public/private key exchange during the handshake. Even if someone intercepts the packet, it looks like garbage. Modern web browsers and APIs enforce HTTPS everywhere."

### 98. What is TCP vs UDP?
"**TCP** is connection-oriented. It guarantees delivery and order. If a packet is lost, it retransmits. It’s reliable but has overhead. We use it for HTTP, Databases, FTP—where data integrity is critical.

**UDP** is connectionless. It fires packets and forgets. If one is lost, it’s gone. It’s unreliable but extremely fast. We use it for video streaming, gaming, or VoIP—where a slight glitch is better than buffering."

### 99. How do you debug production issues?
"It’s a detective process.

1.  **Reproduce**: Can I reproduce it locally or in staging? If yes, great.
2.  **Logs**: If not, I go to the logs (Splunk/ELK). I search for the Request ID or Error stack trace.
3.  **Metrics**: I check Grafana. Was CPU high? Was the DB slow?
4.  **Recent Changes**: Did we deploy anything recently? 90% of the time, it’s a bad config change.

If it's a hard crash, I analyze the heap dump or thread dump."

### 100. How do you read and analyze stack traces?
"I start from the **top** to see the Exception type (`NullPointerException`) and the message.

Then I scan **down** looking for *my* package names (`com.mycompany...`). I ignore the hundreds of lines of Spring/Hibernate framework noise.

Once I find the first line of my code in the trace, that’s usually where the gun is smoking. I go to that line in the IDE and work backwards to see how the state got there."
