# 🗣️ Theory — Go Internals & Runtime
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "How does Go's garbage collector work?"

> *"Go uses a concurrent, tri-color mark-and-sweep garbage collector. The tri-color refers to how objects are classified during a collection cycle. White objects are candidates for collection — not yet visited. Gray objects have been discovered but their references haven't been fully scanned. Black objects are fully scanned and are reachable. The GC starts by marking all roots — globals and goroutine stacks — as gray. Then it scans gray objects, turning their references gray and marking the parent black. When no gray objects remain, everything still white is unreachable and gets swept — their memory reclaimed. The clever part is that most of this happens concurrently with your program running, using write barriers to track pointer changes that happen during the collection."*

---

## Q: "What is escape analysis in Go?"

> *"Escape analysis is what the Go compiler does to decide whether a variable lives on the stack or the heap. The stack is faster — allocation is just moving a pointer — and stack memory is automatically freed when the function returns. The heap is slower and must be garbage collected. The compiler allocates to the stack if it can prove the variable doesn't outlive the function. It moves to the heap — 'escapes' — when: you return a pointer to a local variable, you store a variable in an interface (because the runtime doesn't know the concrete type's size at compile time), or you capture a variable in a goroutine closure. You can see the compiler's decisions with `go build -gcflags='-m'`."*

---

## Q: "What are the stop-the-world pauses in Go's GC? How has Go improved them?"

> *"Stop-the-world, or STW, means the GC pauses all goroutines so it can safely do work without data races. Early versions of Go had long STW pauses — hundreds of milliseconds — which was a dealbreaker for latency-sensitive apps. Over the years, the Go team moved almost everything to concurrent phases. Today, STW pauses are typically in the microsecond range — well under a millisecond. There are still two brief STW phases: one at the start of a GC cycle to enable write barriers, and one at the end to finalize. The concurrent mark phase — the bulk of the work — runs while goroutines continue executing. This makes Go viable for real-time applications like game servers and trading systems."*

---

## Q: "How are interfaces represented in memory? Explain the nil interface gotcha."

> *"An interface value in memory is two machine words: an itab pointer and a data pointer. The itab — or interface table — contains the type's metadata and a table of function pointers for the methods the type implements. The data pointer points to the actual value. A nil interface has BOTH words set to zero. Here's the gotcha: if you assign a typed nil pointer to an interface — like `var d *Dog = nil; var a Animal = d` — the interface is NOT nil. It has the type `*Dog` in the itab field, it just has a nil data pointer. So `a == nil` is false. This subtlety causes real bugs, particularly when returning typed nil errors from functions."*

---

## Q: "What are write barriers in Go's GC?"

> *"A write barrier is a piece of code inserted by the compiler at pointer write locations. Its purpose is to inform the garbage collector about pointer changes that happen while the concurrent mark phase is running. Without write barriers, the GC might not notice that a newly allocated object has been pointed to, and incorrectly free it — a use-after-free bug. Go's hybrid write barrier — introduced in Go 1.14 — marks both the old and new referent of a pointer write as grey when a pointer is modified. This ensures the GC can't miss any live objects even when marking concurrently. The trade-off is a small overhead on every pointer write — typically around 10-20% during a GC cycle."*

---

## Q: "How do you tune the GC in production Go applications?"

> *"The main knob is `GOGC`, which defaults to 100 — meaning the GC triggers when the heap doubles from its size after the previous collection. Lowering GOGC — say to 50 — makes the GC more aggressive, reducing memory usage at the cost of more CPU. Raising it — to 200 — means less frequent GC, using more memory but consuming less CPU. In Go 1.19, they added `GOMEMLIMIT` — a soft ceiling on total memory. When memory approaches the limit, the GC becomes more aggressive automatically. This is often the better way to tune rather than manually adjusting GOGC, because it's adaptive to actual memory pressure."*

---

## Q: "What is cooperative vs preemptive scheduling in Go?"

> *"Early Go used cooperative scheduling — goroutines would only yield at certain safe points, like function calls or channel operations. If a goroutine ran a tight loop with no function calls, it could starve other goroutines on that OS thread. Go 1.14 introduced asynchronous preemption — the scheduler can now interrupt any goroutine at any point, even in a tight loop, by sending a signal to the OS thread. This is done safely by the runtime inserting check points. The result is much fairer scheduling — no goroutine can starve others indefinitely. This also improves STW latency because the GC doesn't have to wait for goroutines to reach safe points."*
