# 🗣️ Theory — Stored Procedures, Functions & Triggers
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What is a stored procedure and why would you use one?"

> *"A stored procedure is a named, precompiled block of SQL and procedural logic stored in the database itself. It can accept IN parameters for input, OUT parameters to return values, and INOUT for both. You call it with CALL procedure_name(args). The classic arguments for stored procedures: they reduce network overhead — instead of multiple round-trips from the application to the database, the logic runs inside the database in one call. They're precompiled so repeated execution is faster. They enforce consistent business logic at the database level. But I'd be honest in an interview — modern practice has moved away from storing business logic in stored procedures because they're hard to version-control, hard to test, and tightly couple application logic to the database. I use stored procedures mostly for complex multi-step data operations like ETL, batch processing, or administrative tasks — not for core business logic."*

---

## Q: "What is the difference between a stored procedure and a stored function?"

> *"The key difference is how they return values and how you use them. A stored procedure can produce output via OUT parameters and can return zero or more result sets — but you CALL it as a standalone statement. A stored function always returns exactly one scalar value via a RETURN statement, and you embed it inside SQL expressions — you can use it in a SELECT list, a WHERE clause, anywherean expression is valid. Functions must be deterministic or explicitly declared non-deterministic, and they shouldn't have side effects — they shouldn't modify data, though technically they can. The practical guideline: use a function when you need a reusable computation that fits into a SQL expression — like calculating a bonus, formatting a string, or converting a unit. Use a procedure when you need multi-step logic, multiple result sets, transaction management, or error handling with SIGNAL."*

---

## Q: "How do triggers work? When would you use them?"

> *"A trigger is automatically executed code attached to a table that fires in response to INSERT, UPDATE, or DELETE operations. Each trigger fires either BEFORE or after the data change. BEFORE triggers can modify the NEW values before they're written — useful for validation, normalization, or setting derived values. AFTER triggers fire after the data is written — useful for audit logging, propagating changes to other tables, or maintaining denormalized data. Best use cases: audit trails — whenever salary changes, log old and new values to an audit table. Cascade computed columns — update a summary table whenever detail rows change. Input normalization — automatically trim whitespace or title-case a name before insert. The cautions: triggers add invisible overhead to DML operations — a developer running a bulk INSERT might not realize 100,000 trigger firings are happening. They can also cause cascading effects that are hard to debug. I use them selectively and always document them clearly."*

---

## Q: "What is an event in MySQL and how is it different from a trigger?"

> *"An event is MySQL's built-in job scheduler — it runs SQL code on a time-based schedule, like a cron job inside the database. You create an event with a schedule — AT a specific time for a one-shot event, or EVERY interval for recurring events. Examples: DELETE old log rows every night at 2 AM, run ANALYZE TABLE weekly, send a report query to a log table monthly. An event is time-triggered. A trigger is operation-triggered — it fires in response to a specific DML action on a specific table. Events need the event_scheduler to be running — SET GLOBAL event_scheduler = ON or configure it in my.cnf. I use events for maintenance tasks that need to happen automatically — cleanup jobs, rolling up statistics, expiring old sessions. The downside is that events are database-internal — they're less visible to the operations team than a cron job or a scheduled task in a job scheduler like Kubernetes CronJob, so I prefer application-level scheduling for anything critical."*

---

## Q: "What is a cursor in a stored procedure? When should you avoid it?"

> *"A cursor lets you iterate over a result set row by row inside a stored procedure. You declare the cursor with a SELECT statement, open it, fetch rows one at a time in a loop, do your logic per row, and close it when done. The critical thing to emphasize in an interview: cursors are the last resort. SQL is set-based — the engine is optimized to process entire sets of rows at once. Row-by-row cursor iteration throws away that optimization and basically writes a loop that processes one row at a time. For 1 million rows, that's 1 million iterations. Almost anything you'd use a cursor for can be rewritten as a set-based operation — UPDATE with a JOIN, INSERT INTO ... SELECT, or a window function. Cursors are acceptable when the per-row logic is genuinely complex and cannot be expressed in set-based SQL — for example, calling external procedures per row or doing complex conditional branching that's too hard to express declaratively."*

---

## Q: "What is a view? What are updatable vs non-updatable views?"

> *"A view is a named, saved SELECT query — it acts as a virtual table. You query it as if it were a real table, but underneath, the database re-runs the stored SELECT each time. Views are useful for: simplifying complex queries that involve many JOINs — define the view once, query it simply everywhere. Security — create a view that exposes only specific columns like name and department, hiding sensitive columns like salary or SSN, and grant users access to the view rather than the underlying table. Encapsulating business logic like active_customers being defined as customers where status='active'. An updatable view is one where MySQL can trace each column in the view back to exactly one base table column — then INSERT, UPDATE, and DELETE on the view propagate to the underlying table. Non-updatable views contain aggregates, DISTINCT, GROUP BY, UNION, or subqueries — MySQL can't determine which base rows to modify, so they're read-only. Regular views don't improve performance — the query still runs fresh every time. For performance gains, you'd manually materialize a view into a real table and refresh it periodically."*

---

## Q: "Why does MySQL require the DELIMITER command for stored procedures?"

> *"MySQL's command-line client uses semicolons as the statement terminator — when it sees a semicolon, it sends whatever it has collected to the server as a complete statement. Inside a stored procedure or trigger, you have multiple statements separated by semicolons. Without changing the delimiter, the client would send the first statement inside the procedure body as a complete statement before you finish defining the procedure — it would be a syntax error. So you change the delimiter to something else — typically two dollar signs or double forward slashes — and now the client waits until it sees that new delimiter before sending the statement to the server. You wrap your entire CREATE PROCEDURE ... END with the changed delimiter, and the client sends the whole thing as one unit. Then you change the delimiter back to semicolon. This is strictly a client-side concern — it only applies when using the mysql command-line client or script files. In application code using a driver like Go's database/sql, you just pass the procedure body as a string and there's no delimiter issue."*

