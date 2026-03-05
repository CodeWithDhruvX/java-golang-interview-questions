# Behavioral Story Bank Template

The best FAANG candidates don't memorize 50 answers to 50 questions. 
They prepare **4 to 6 incredibly strong "Master Stories"** using the STAR method. Because a *single* complex story can be used to answer many different questions (e.g., the same story about a production outage can answer a question about "failure", "working under pressure", "deep diving into a bug", or "dealing with an angry customer").

## How to use this template:
1. Brainstorm 5 different impactful scenarios from your career. Make sure to include:
   - Your biggest technical achievement.
   - Your biggest failure / mistake.
   - A time you had a tough conflict with a coworker/manager.
   - A time you led a project or mentored someone.
   - A time you dealt with extreme ambiguity or changing requirements.
2. Fill out the grid below. **Memorize the metrics and data**.

---

### Story 1: The Major Technical Win / Deep Dive
*(Use for: "Most complex project", "Proudest moment", "Diving deep", "Delivering results")*

- **S (Situation):** [Context, the stakes. E.g., The core API was crashing during Black Friday.]
- **T (Task):** [What did YOU have to do? E.g., Identify the root cause and patch it within 24 hours.]
- **A (Action):** [Steps YOU took. E.g., 1. Setup DataDog tracing. 2. Identified N+1 query loop. 3. Wrote a batched SQL query to bypass the ORM.]
- **R (Result & Metrics):** [The concrete impact. E.g., Dropped latency from 3s to 120ms. Handled 30k Requests/Second. Saved $2M in lost revenue.]

### Story 2: The Brutal Failure and the Pivot
*(Use for: "A time you failed", "Missed a deadline", "Made a mistake", "Learned a lesson")*

- **S (Situation):** [Context. E.g., Migrating the user DB from Postgres to Mongo.]
- **T (Task):** [What were you trying to do? E.g., Run a zero-downtime migration over the weekend.]
- **A (Action):** [What went wrong, and how did you FIX it? E.g., The schema translation caused data corruption. Instead of hiding it, I immediately alerted leadership, rolled back the DB using our snapshot, and spent the entire night writing a script to clean the corrupted user profiles.]
- **R (Result & Metrics):** [The learning. E.g., 500 users experienced 10 mins of downtime. But I took ownership, created an automated testing suite for migrations, and the next attempt ran flawlessly.]

### Story 3: The Interpersonal Conflict
*(Use for: "Disagreed with someone", "Difficult teammate", "Persuading someone", "Earn Trust")*

- **S (Situation):** [Context. E.g., A senior dev wanted to use microservices, I wanted a monolith for our small MVP.]
- **T (Task):** [The goal. E.g., We had to finalize the architecture by Friday to start coding.]
- **A (Action):** [How did you handle the people aspect? E.g., Instead of arguing, I scheduled a 1-on-1. I brought objective data showing that our team size (3 people) couldn't maintain the overhead of 10 microservices. I listened fully to their concerns about scalability. We compromised on a modular monolith.]
- **R (Result & Metrics):** [The outcome. E.g., The MVP shipped 1 month early. The relationship with the senior dev improved because I handled the dispute with data, not emotion.]

### Story 4: Leading through Ambiguity
*(Use for: "Vague requirements", "Stepped up as a leader", "Invent & Simplify", "Moving Fast")*

- **S (Situation):** [Context. E.g., PM left the company mid-project, leaving behind a 2-sentence Jira ticket.]
- **T (Task):** [The goal. E.g., Build a GDPR compliance deletion tool.]
- **A (Action):** [How did you create order from chaos? E.g., I stepped up as the unofficial PM. I scheduled meetings with Legal to define exactly what data needed deleting. I drafted a 2-page tech spec, got buy-in from the engineering lead, and broke it down into 5 Jira tickets for the team.]
- **R (Result & Metrics):** [The outcome. E.g., Shipped compliance tool 2 weeks before the EU deadline. Prevented potential massive fines. Received a spot-bonus for stepping outside my strict coding duties.]

---

### 🔥 Pro-Tip: The "So What?" Test
After writing down your Result, ask yourself: *"So what?"*
If your result is "I finished the feature," the interviewer will think "So what? That's your job."
If your result is "I finished the feature which increased user retention by 5% and brought in $50k ARR," you pass.
