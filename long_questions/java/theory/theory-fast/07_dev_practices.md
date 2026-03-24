# Git, Linux & Debugging Interview Questions (91-100)

## Git & Development Practices

### 91. Difference between `merge` and `rebase`?
"They both integrate changes from one branch to another, but the history looks different.

**Merge** creates a new 'merge commit'. It preserves the history exactly as it happened. If you look at the graph, you see branches diverging and coming back together. It’s 'true' history but can get messy.

**Rebase** rewrites history. It takes your commits from the feature branch and replays them on top of the master branch. The result is a perfectly linear history, which is much cleaner to read.

In my team, we rebase local branches to keep them clean, but always merge (squash merge) when bringing feature branches into master."

**Spoken Format:**
"Think of `merge` and `rebase` as two different ways to combine stories in a book.

**Merge** is like taking two chapters and creating a new chapter that says 'Chapter A and Chapter B were combined'. It preserves the exact history - you can see exactly how the stories came together, even if they diverged.

**Rebase** is like taking Chapter B and rewriting it to come after Chapter A. The history looks like Chapter A happened first, then Chapter B was added on top. It creates a clean, linear story.

The tradeoff: Merge preserves the messy truth of how things actually happened, while Rebase creates a clean but fictional timeline.

In practice, I use rebase for my local work to keep my history clean, but when bringing features into the main branch, I use merge to preserve the actual development history.

It's like keeping a personal diary clean (rebase) but maintaining an official timeline for the team (merge)."

### 92. What is a pull request?
"A Pull Request (PR) is a request to merge your code changes into the main codebase.

It’s the core of our code review process. I push my branch, open a PR, and my teammates review it. They check for bugs, style issues, and logic gaps. The CI pipeline also runs tests on the PR automatically.

Once everything is green and approved, it gets merged. It's a quality gate."

**Spoken Format:**
"Pull Requests are like sending your chapter to an editor for review before it gets published in the book.

When I finish working on a feature, I create a Pull Request - it's like saying 'Hey editors, please review my chapter before it goes into the main book'.

The process works like this:
1. I push my changes to a separate branch
2. I open a PR asking to merge it into the main branch
3. My teammates review it - they check for bugs, style issues, and logic gaps
4. The CI pipeline automatically runs tests to make sure nothing breaks
5. Once everyone approves and tests pass, the PR gets merged

It's like having a quality control system where multiple people have to approve changes before they become official. This prevents bad code from getting into the main codebase.

PRs are the backbone of modern development - they ensure code quality and knowledge sharing across the team!"

### 93. How do you resolve merge conflicts?
"First, I don't panic. A conflict just means Git doesn't know which version of a line to pick.

I usually run `git rebase main` on my feature branch. Git stops at the conflict. I open the file in my IDE (IntelliJ), which shows a 3-way view: 'Yours', 'Theirs', and 'Result'.

I pick the changes I want (or combine them), save the file, `git add`, and `git rebase --continue`. It's part of the job."

**Spoken Format:**
"Merge conflicts are like two people trying to edit the same document at the same time.

When Git finds a conflict, it's like saying 'I don't know which version is correct - here are both options'.

The solution is to use the 3-way view Git provides:
- **Yours**: Your version of the changes
- **Theirs**: The version from the other branch
- **Result**: The final merged version

I look at both versions, decide what I want to keep (or combine parts of both), edit the file to remove the conflict markers, then continue with `git rebase --continue`.

The key is not to panic - conflicts are normal in team development. They're just Git's way of asking for human intervention when it can't automatically decide.

It's like being a mediator who helps two people agree on the final version of their shared document!"

### 94. What is CI/CD?
"CI (Continuous Integration) is the practice of merging code frequently. Every time I push, a build server (Jenkins/GitHub Actions) compiles the code and runs unit tests. If it fails, I fix it immediately.

CD (Continuous Delivery/Deployment) takes it further. Once the build passes, it automatically deploys the app to a staging environment (Delivery) or even straight to production (Deployment).

It replaces the old 'release night' panic with a boring, automated process."

**Spoken Format:**
"CI/CD is like having an automated assembly line for your software.

**Continuous Integration (CI)** is like having quality control at every step of the assembly line. Every time a developer adds code (pushes), the system automatically:
- Compiles the code to make sure it builds
- Runs unit tests to ensure it works
- Checks code quality and style
If anything fails, the assembly line stops immediately for that developer to fix.

**Continuous Deployment/Delivery (CD)** is like having the assembly line automatically package and ship the finished products. Once code passes all quality checks, it gets deployed to staging or production automatically.

This eliminates the old 'release night' panic where everyone stays late, manually deploying code and hoping nothing breaks.

CI/CD makes software development predictable, fast, and much less stressful!"

### 95. Common Linux commands you use?
"I live in the terminal.

-   `grep`: Searching logs (`grep "ERROR" app.log`).
-   `tail -f`: Watching logs in real-time.
-   `ps -ef | grep java`: Checking if my process is running.
-   `netstat` or `lsof`: Checking ports.
-   `curl`: Testing APIs.
-   `chmod`: Changing permissions.
-   `top` / `htop`: Checking CPU/Memory usage."

**Spoken Format:**
"Linux commands are like having a Swiss Army knife for system administration - each tool solves specific problems.

For searching logs, I use `grep` - it's like having a super-fast search function that can find any pattern in massive log files. Instead of scrolling through thousands of lines, I just say 'find all ERROR messages'.

For watching logs in real-time, I use `tail -f` - it's like having a live news feed that shows me what's happening right now.

For checking processes, `ps -ef | grep java` is like having a task manager that shows me exactly what Java processes are running.

For network diagnostics, `netstat` shows me which ports are open - like checking which doors are unlocked in my building.

For performance monitoring, `top` and `htop` are like having a dashboard that shows me which applications are using the most resources.

These commands help me understand what's happening in the system without needing fancy GUI tools!"

### 96. How do you check if a port is open?
"On a remote server, I use `telnet host port` or `nc -zv host port` (Netcat).

If I'm checking strictly on the local machine to see what's listening, I use `netstat -tulpn` or `lsof -i :8080`.

This is usually the first step when a service fails to start—checking if the port is already taken."

**Spoken Format:**
"Checking if a port is open is like checking if a door is unlocked before trying to enter.

For remote servers, I use `telnet` or `nc` (netcat) - it's like trying to knock on the door to see if anyone answers.

For local machines, I use `netstat -tulpn` or `lsof -i :8080` - it's like checking which doors in my building are currently open.

The process is:
1. Service fails to start with 'Port already in use' error
2. I check if something else is actually using that port
3. If yes, I either stop the other service or use a different port
4. If no, I investigate why my service can't bind to that port

It's the first step in debugging - like trying to understand why you can't park in your usual spot before calling a tow truck.

This simple check often reveals the obvious problem - another service is already running where you expected to start yours!"

### 97. Difference between HTTP and HTTPS?
"HTTPS is simply HTTP over a secure encrypted connection (SSL/TLS).

With HTTP, data is sent in plaintext. Anyone on the network (Wi-Fi) can sniff my password.

With HTTPS, the data is encrypted using a public/private key exchange during the handshake. Even if someone intercepts the packet, it looks like garbage. Modern web browsers and APIs enforce HTTPS everywhere."

**Spoken Format:**
"HTTP vs HTTPS is like the difference between sending a postcard and sending a sealed letter.

**HTTP** is like sending a postcard - anyone who handles it can read the message. If someone intercepts it, they can see everything you wrote. It's fast but completely insecure.

**HTTPS** is like sending a sealed letter with a tamper-proof envelope. The message is encrypted, and only the intended recipient can open it. Even if someone intercepts it, they just see gibberish.

The magic happens during the handshake - your browser and the server exchange secret keys to establish a secure connection. After that, all communication is encrypted.

Modern browsers and APIs automatically enforce HTTPS - they refuse to send postcards (HTTP) and only accept sealed letters (HTTPS).

It's like the postal service deciding to only handle secure, sealed envelopes for important communications!"

### 98. What is TCP vs UDP?
"**TCP** is connection-oriented. It guarantees delivery and order. If a packet is lost, it retransmits. It’s reliable but has overhead. We use it for HTTP, Databases, FTP—where data integrity is critical.

**UDP** is connectionless. It fires packets and forgets. If one is lost, it’s gone. It’s unreliable but extremely fast. We use it for video streaming, gaming, or VoIP—where a slight glitch is better than buffering."

**Spoken Format:**
"TCP vs UDP is like the difference between registered mail and instant messaging.

**TCP** is like registered mail - it guarantees delivery and order. If a packet gets lost, the system automatically resends it. You know exactly when your message was delivered. It's reliable but has overhead.

**UDP** is like instant messaging - you fire off a message and hope it arrives. If it gets lost, it's just gone. There's no guarantee of delivery or order.

The tradeoff is speed vs. reliability:
- TCP: Slower but reliable - perfect for websites, file transfers, databases
- UDP: Faster but unreliable - perfect for video streaming, online gaming, voice calls

For video calls, you'd rather have a frozen frame than wait for a delayed one. For gaming, you'd rather have your character move immediately than wait for the server to confirm.

Choose based on whether you need guaranteed delivery or maximum speed!"

### 99. How do you debug production issues?
"It’s a detective process.

1.  **Reproduce**: Can I reproduce it locally or in staging? If yes, great.
2.  **Logs**: If not, I go to the logs (Splunk/ELK). I search for the Request ID or Error stack trace.
3.  **Metrics**: I check Grafana. Was CPU high? Was the DB slow?
4.  **Recent Changes**: Did we deploy anything recently? 90% of the time, it’s a bad config change.

If it's a hard crash, I analyze the heap dump or thread dump."

**Spoken Format:**
"Debugging production issues is like being a detective at a crime scene - you need to follow the evidence systematically.

The process is:

1. **Reproduce** - Can I make the problem happen again in a safe environment? If yes, I can experiment with solutions safely.

2. **Logs** - If not, I go to the crime scene (production logs). I use tools like Splunk or ELK to search for clues - error messages, stack traces, request IDs.

3. **Metrics** - I check the security cameras (Grafana) - was CPU high? Was memory usage unusual? Did database response time spike?

4. **Recent Changes** - What changed right before the crime? 90% of the time, it's a recent deployment or configuration change.

5. **Deep Analysis** - If it's a crash, I analyze the evidence (heap dumps, thread dumps) to understand exactly what happened.

The key is to be systematic - don't just randomly change things. Follow the evidence from logs to metrics to code changes. That's how you solve production mysteries!"

### 100. How do you read and analyze stack traces?
"I start from the **top** to see the Exception type (`NullPointerException`) and the message.

Then I scan **down** looking for *my* package names (`com.mycompany...`). I ignore the hundreds of lines of Spring/Hibernate framework noise.

Once I find the first line of my code in the trace, that's usually where the gun is smoking. I go to that line in the IDE and work backwards to see how the state got there."

**Spoken Format:**
"Reading stack traces is like being a medical examiner - you need to find the cause of death by working backwards from the symptoms.

The process is:

1. **Start at the top** - The first thing I look for is the exception type and message. This is like finding the cause of death on the death certificate.

2. **Scan down for my code** - I look for package names that match my application (`com.mycompany...`). I ignore hundreds of lines of framework code that aren't relevant.

3. **Find the smoking gun** - The first line of my own code that appears in the trace is usually where the problem originated. It's like finding the wound that caused the death.

4. **Work backwards** - Once I find the problem line, I trace backwards through the method calls to understand how the program got to that state.

Stack traces are the program's final words - they tell you exactly what happened before the crash. Your job is to be a good detective and listen to what they're saying!"
