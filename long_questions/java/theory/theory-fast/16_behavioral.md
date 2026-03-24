# Behavioral & Real-World Interview Questions (161-170)

## Engineering Culture & Soft Skills

### 161. Describe a production issue you handled end-to-end.
"Once, our payment processing latency spiked 10x during a sale event.

I immediately got on the incident call and checked the metrics. The DB CPU was fine, but the App Threads were maxed out. I took a thread dump and found hundreds of threads blocked waiting for a 3rd party fraud check API.

I realized their API was slow, causing backpressure on us. I hot-patched a change to reduce the timeout from 5s to 1s and enabled a circuit breaker.

Latency recovered immediately. Later, I refactored the fraud check to be asynchronous so it wouldn't block the main checkout flow."

**Spoken Format:**
"Handling production issues is like being a detective who solves crimes in progress.

When a critical issue happens during a major sale, you need to work systematically:

**1. Immediate Response** - Like securing the crime scene first to prevent further damage.

**2. Investigation** - Like gathering evidence from multiple sources:
- Metrics dashboard shows you where the problem is (CPU spikes, thread issues)
- Thread dump reveals who the suspects are (blocked threads)
- Logs tell you what actually happened (error messages, timeouts)

**3. Root Cause Analysis** - Like connecting all the evidence to understand the real story.

**4. Solution Implementation** - Like implementing the fix based on your findings.

**5. Prevention** - Like improving security to prevent similar crimes.

The key is to be methodical - gather evidence, analyze patterns, implement solutions, and prevent recurrences. Don't just fix symptoms, solve the underlying crime!"

### 162. How do you prioritize bug fixes vs new features?
"It’s a balancing act.

If it's a **P0 critical bug** (data loss, security, widespread outage), I drop everything and fix it. If it's a **P2/P3 bug** (minor UI glitch), I weigh it against the value of the new feature. I usually reserve 20% of my sprint capacity for 'tech debt & bugs'. If we ignore bugs for too long, we lose customer trust. If we ignore features, we lose market share. I try to group small bugs into a 'Bug Bash' day to clear them out efficiently."

**Spoken Format:**
"Prioritizing work is like being a smart project manager who balances urgent fixes with long-term progress.

**P0 Critical Issues** are like the building being on fire - you drop everything and fight the fire immediately. Customer safety and company survival come first.

**P2/P3 Minor Issues** are like having a leaky faucet - it's annoying but won't destroy the building. You can fix it during regular maintenance.

**The Strategy**:
- Allocate capacity for both types of issues
- For critical issues: Use emergency resources immediately
- For minor issues: Schedule regular maintenance time
- Track both types to ensure neither gets neglected
- Communicate priorities clearly to stakeholders

The key insight: Ignoring small issues leads to big problems later. Addressing both types ensures your system stays reliable and your customers stay happy!"

### 163. How do you handle technical debt?
"I treat tech debt like financial debt. A mortgage is fine; credit card debt is bad. Taking on debt (quick & dirty code) to meet a critical deadline is okay, *if* we have a plan to pay it back."

**Spoken Format:**
"Technical debt is like borrowing against your future productivity.

**Good Debt** (mortgage) - Taking a loan to buy a house that will generate value. It's a strategic investment with a clear repayment plan.

**Bad Debt** (credit card) - Using a credit card for daily expenses with high interest. It provides temporary relief but creates long-term problems.

**The Strategy**:
- Track debt with interest (how much it's costing you in maintenance)
- Have a repayment plan (refactoring roadmap)
- Don't take on new debt without understanding the cost
- Sometimes it's okay to take on strategic debt for critical deadlines
- But make sure the 'interest payments' (refactoring effort) don't exceed the value generated

The key is to treat technical debt as a business decision, not just a technical problem. Every quick fix should include a plan to pay back the borrowed time!"

I document debt with `// TODO` comments and Jira tickets immediately. Then, during sprint planning, I advocate for picking up at least one 'debt payoff' task. If we don't, the interest (development slowness) eventually compounds and kills our velocity."

### 164. How do you disagree with a design decision?
"I focus on **data and trade-offs**, not opinions.

Instead of saying 'I don't like this', I say 'If we choose Option A, we gain speed but lose consistency. Have we considered the risk of data corruption?'

If the team (or lead) still decides on Option A, I commit to it fully. Usage of 'disagree and commit' is important. I don't sabotage the project by saying 'I told you so' later. We win or lose as a team."

### 165. How do you estimate backend work?
"I break the task down until each piece is no smaller than half a day.

I look at complexity, not just lines of code.
-   Is there a DB migration? (buffer +1 day)
-   Is there a 3rd party integration? (buffer +2 days for bad docs)
-   Do I need to touch legacy code? (buffer +50%)

I always give a range ('3 to 5 days') rather than a specific date, and I communicate early if I see I'm going to miss it."

### 166. How do you do code reviews effectively?
"I look for logic errors and security issues first, not just formatting (linters handle that).

I ask questions rather than giving orders. 'Have you considered handling the null case here?' vs 'Fix this null check.'

I also praise good code. 'This refactor is really clean, nice job.'

And I try to turn reviews around quickly (within 4 hours). Blocking a teammate is the worst thing I can do for team velocity."

### 167. How do you handle on-call incidents?
"Calmly.

1.  **Acknowledge**: 'I'm looking into it.'
2.  **Mitigate**: Stop the bleeding. Roll back the deploy, or flip a feature flag. Fixing the root cause comes *later*.
3.  **Communicate**: Update the status page or stakeholders every 30 mins.
4.  **Post-Mortem**: After it's over, write a blameless RCA (Root Cause Analysis). How did this happen? How do we prevent it automatically next time?"

### 168. What trade-offs have you made for delivery speed?
"In my last startup, we needed to launch a referral program in 3 days.

Ideally, I would have built a scalable, event-driven service with a dedicated graph database.

Instead, I accepted the trade-off: I built a synchronous, monolithic implementation using our existing Postgres DB. I knew it wouldn't scale to 1M users, but it got us to market to *validate* if users even wanted referrals.

We later refactored it when traffic grew. Speed was the feature."

### 169. How do you ensure code quality under deadlines?
"I cut **scope**, not **quality**.

If we have a hard deadline, I’d rather deliver 3 solid features than 5 buggy ones. I’ll talk to the PM: 'We can't do the fancy animated dashboard, but we can do a simple table that works perfectly.'

I also rely heavily on automated tests. They are my safety net. If I skip writing tests to 'save time', I usually end up spending twice as much time debugging later."

### 170. What would you improve in your last project?
"We built a monolithic service that became a 'distributed monolith'. We split the code into services but they shared the same database.

This meant if one service locked a table, others halted. It was the worst of both worlds.

If I could do it again, I would have been stricter about Database-per-Service boundaries from day one, even if it meant more work syncing data initially. It would have saved us months of pain later."
