# Behavioral & Leadership - Interview Answers

> 🎯 **Focus:** Use the **STAR Method** (Situation, Task, Action, Result). Be honest, but strategic.

### 1. Tell me about a time you disagreed with a coworker. (Conflict Resolution)
"**Situation**: We were designing a new API. My peer wanted to use SOAP because he was familiar with it, but I advocated for REST/JSON for better mobile compatibility.
**Action**: Instead of arguing opinions, I built a quick prototype of both. I showed that the JSON payload was 40% smaller and easier for our frontend team to consume. I acknowledged that SOAP had better strict contract enforcement but proposed we use OpenAPI (Swagger) with REST to get the same benefit.
**Result**: He agreed with the data. We went with REST, and the mobile team shipped the feature 2 weeks early because integration was so easy."

---

### 2. Describe a time you failed. (Handling Failure)
"**Situation**: Early in my career, I deployed a hotfix directly to production on a Friday evening to fix a critical bug.
**Action**: I missed a database migration script. The app crashed for 20 minutes.
**Result**: I immediately rolled back. But I didn't hide it. I wrote a post-mortem.
**Learning**: I implemented a strict CI/CD pipeline policy: 'No manual deployments' and 'No deployments on Fridays'. It was a painful lesson, but it forced us to automate our safety checks."

---

### 3. What is your biggest weakness?
"I sometimes get too deep into optimization before it's needed (Premature Optimization).
I might spend 3 hours trying to make a query 5ms faster when it runs only once a day.
**How I manage it**: I now write 'Make it work, Make it right, Make it fast' on a sticky note. I force myself to ship the working solution first, and only optimize if monitoring tools show it's actually a bottleneck."

---

### 4. Tell me about a challenging technical problem you solved.
"**Situation**: Our order processing system was getting slower every day.
**Investigation**: I used a profiler (VisualVM) and found that we were loading the entire 'Order History' for a user just to check their current status. It was a massive N+1 problem.
**Action**: I refactored the JPA query to use a 'Fetch Join' and implemented a Redis cache for active orders.
**Result**: API latency dropped from 2 seconds to 50ms, and database load decreased by 60%."

---

### 5. How do you handle tight deadlines?
"**Communication is key**.
If I see a deadline is at risk, I flag it early—not on the day it's due.
I usually propose a **Scope Cut**. 'We can't deliver all 5 features by Friday, but we can deliver the core 3 features perfectly, and push the other 2 to next sprint.'
Managers usually appreciate the honesty and options rather than a surprise delay."

---

### 6. Why do you want to leave your current company?
"I've learned a ton there, especially about [mention skill, e.g., Microservices].
However, I feel I've reached a plateau. My current team is in maintenance mode, and I'm eager to work on [mention challenge, e.g., high-scale distributed systems] which your company is known for. I want to be in an environment where I'm challenged to grow technically."

---

### 7. Describe a time you showed leadership (even if not a manager).
"**Situation**: Our daily standups were taking 45 minutes and draining energy.
**Action**: I suggested we switch to a 'Walking the Board' format—focusing on tickets that are blocked, rather than everyone giving a status update.
**Result**: Standups dropped to 15 minutes. The team was happier, and we started identifying blockers much faster."
