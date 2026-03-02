# 📚 Go Practical Basics — Reading Progress Tracker

> **Format of each question:** "Predict the output / Spot the bug / Does it compile?" style.
> **How to use this file:** Mark `[ ]` → `[x]` as you complete each section. Update the file status row at the top too.

---

## 📁 Files Overview (Recommended Reading Order)

| # | File | Topics | Questions | Status |
|---|------|---------|-----------|--------|
| 1 | [go_basics_fundamentals_snippets.md](go_basics_fundamentals_snippets.md) | Variables · Types · Control Flow · Functions · Closures · Pointers · Slices · Maps · Structs · Interfaces · Error Handling · Defer | **100 Q** | ⬜ Not Started |
| 2 | [go_basics_indepth_snippets.md](go_basics_indepth_snippets.md) | Deep-dive gotchas on all the above — scope, type system, closures, goroutines, channels | **100 Q** | ⬜ Not Started |
| 3 | [go_methods_and_io_snippets.md](go_methods_and_io_snippets.md) | Value vs Pointer Receivers · Method Sets · io.Reader/Writer · Functional Options · Channel patterns | **~65 Q** | ⬜ Not Started |
| 4 | [go_modern_go_snippets.md](go_modern_go_snippets.md) | new/make · Slice tricks · Map gotchas · Struct Embedding · Interface nil trap · Generics · errors.Is/As · select · sync | **~80 Q** | ⬜ Not Started |
| 5 | [go_context_testing_reflect_snippets.md](go_context_testing_reflect_snippets.md) | context.Context · Goroutine leaks · strings.Builder · Table-driven tests · Benchmarks · reflect | **~72 Q** | ⬜ Not Started |
| 6 | [go_service_company_coverage.md](go_service_company_coverage.md) | fmt · strings/strconv · math · os/File I/O · JSON · time · net/http · sort · log · testing | **~75 Q** | ⬜ Not Started |
| 7 | [go_stdlib_patterns_snippets.md](go_stdlib_patterns_snippets.md) | flag · regexp · runtime · path/filepath · unicode/utf8 · embed · net/url · os/exec · math/rand | **~90 Q** | ⬜ Not Started |

> 💡 **Total: ~582 questions across 7 files**

---

## ✅ Section-Level Progress

### File 1 — `go_basics_fundamentals_snippets.md`
*🎯 Priority: HIGH — Start here for any interview*

- [ ] **Section 1:** Variables, Constants & Types (Q1–Q15)
- [ ] **Section 2:** Control Flow (Q16–Q28)
- [ ] **Section 3:** Functions, Closures & Defer (Q29–Q44)
- [ ] **Section 4:** Pointers (Q45–Q52)
- [ ] **Section 5:** Strings & Runes (Q53–Q62)
- [ ] **Section 6:** Arrays, Slices & Maps (Q63–Q78)
- [ ] **Section 7:** Structs & Interfaces (Q79–Q91)
- [ ] **Section 8:** Error Handling (Q92–Q96)
- [ ] **Section 9:** Goroutines Basics & Misc (Q97–Q100)

---

### File 2 — `go_basics_indepth_snippets.md`
*🎯 Priority: HIGH — Product companies love these tricky edge cases*

- [ ] **Section 1:** Variables, Scope & Type System Deep Dives (Q1–Q18)
- [ ] **Section 2:** Control Flow Deep Dives (Q19–Q30)
- [ ] **Section 3:** Functions & Closures Deep Dives (Q31–Q45)
- [ ] **Section 4:** Pointers Deep Dives (Q46–Q56)
- [ ] **Section 5:** Strings & Runes Deep Dives (Q57–Q67)
- [ ] **Section 6:** Slices & Maps Deep Dives (Q68–Q83)
- [ ] **Section 7:** Structs & Interfaces Deep Dives (Q84–Q95)
- [ ] **Section 8:** Error Handling Deep Dives (Q96–Q100)

---

### File 3 — `go_methods_and_io_snippets.md`
*🎯 Priority: HIGH — Interface satisfaction bugs come up constantly*

- [ ] **Section 1:** Value Receivers vs Pointer Receivers (Q1–Q15)
- [ ] **Section 2:** io.Reader / io.Writer Patterns (Q16–Q38)
- [ ] **Section 3:** Functional Options Pattern (Q39–Q50)
- [ ] **Section 4:** Channel Composition Patterns (Q51–Q65)

---

### File 4 — `go_modern_go_snippets.md`
*🎯 Priority: HIGH — Modern Go patterns, Generics, sync primitives*

- [ ] **Section 1:** `new` vs `make`, Slice & Map Gotchas (Q1–Q16)
- [ ] **Section 2:** Struct Embedding & Promoted Methods (Q17–Q25)
- [ ] **Section 3:** The Interface nil Trap & errors.Is / errors.As (Q26–Q38)
- [ ] **Section 4:** Generics (Q39–Q52)
- [ ] **Section 5:** `select`, Channel Directions & Patterns (Q53–Q66)
- [ ] **Section 6:** sync.Mutex, sync.RWMutex, sync.Once, sync.Map (Q67–Q80)

---

### File 5 — `go_context_testing_reflect_snippets.md`
*🎯 Priority: MEDIUM — Essential for production Go / senior roles*

- [ ] **Section 1:** context.Context (Q1–Q22)
- [ ] **Section 2:** Goroutine Leaks (Q23–Q32)
- [ ] **Section 3:** strings.Builder & bytes.Buffer (Q33–Q44)
- [ ] **Section 4:** Table-Driven Tests & testing.T (Q45–Q60)
- [ ] **Section 5:** reflect Basics (Q61–Q72)

---

### File 6 — `go_service_company_coverage.md`
*🎯 Priority: MEDIUM — Service company stdlib breadth questions*

- [ ] **Section 1:** fmt Package Deep Dives (Q1–Q12)
- [ ] **Section 2:** strings & strconv Package (Q13–Q24)
- [ ] **Section 3:** math Package (Q25–Q29)
- [ ] **Section 4:** os Package & File I/O (Q30–Q38)
- [ ] **Section 5:** JSON Encoding / Decoding (Q39–Q46)
- [ ] **Section 6:** time Package (Q47–Q54)
- [ ] **Section 7:** sort Package (Q55–Q59)
- [ ] **Section 8:** log Package (Q60–Q63)
- [ ] **Section 9:** Basic net/http (Q64–Q70)
- [ ] **Section 10:** Basic Testing Patterns (Q71–Q75)

---

### File 7 — `go_stdlib_patterns_snippets.md`
*🎯 Priority: LOW-MEDIUM — Good for completeness and CLI tool roles*

- [ ] **Section 1:** flag Package (Q1–Q12)
- [ ] **Section 2:** regexp Basics (Q13–Q24)
- [ ] **Section 3:** runtime Package (Q25–Q34)
- [ ] **Section 4:** path/filepath Package (Q35–Q44)
- [ ] **Section 5:** unicode/utf8 Package (Q45–Q54)
- [ ] **Section 6:** embed Package (Q55–Q60)
- [ ] **Section 7:** net/url & os/exec (Q61–Q72)

---

## 🗺️ Topic → File Quick Reference

| Topic | File |
|-------|------|
| Variables, Types, Constants | File 1 → File 2 |
| Control Flow (for/switch/goto) | File 1 → File 2 |
| Functions, Closures, Defer | File 1 → File 2 |
| Pointers | File 1 → File 2 |
| Strings & Runes | File 1 → File 2, File 6 |
| Arrays, Slices | File 1 → File 2, File 4 |
| Maps | File 1 → File 2, File 4 |
| Structs & Embedding | File 1 → File 2, File 4 |
| Interfaces & Method Sets | File 3 → File 1, File 4 |
| Value vs Pointer Receiver | File 3 → File 2 |
| io.Reader / io.Writer | File 3 → File 6 |
| Goroutines (basic) | File 1 → File 2 |
| Goroutine Leaks | File 5 |
| context.Context | File 5 |
| select & Channels | File 4 → File 2 |
| sync (Mutex/Once/Map) | File 4 |
| Generics | File 4 |
| errors.Is / errors.As | File 4 → File 2 |
| Interface nil trap | File 4 → File 2 |
| fmt package | File 6 → File 1 |
| strings / strconv | File 6 → File 3 |
| JSON | File 6 |
| time | File 6 |
| net/http basics | File 6 |
| sort | File 6 → File 7 |
| testing / benchmarks | File 5 |
| reflect | File 5 |
| flag | File 7 |
| regexp | File 7 |
| runtime | File 7 |
| path/filepath | File 7 |
| unicode/utf8 | File 7 |
| embed | File 7 |

---

## 📌 Study Strategy

| Goal | Priority Files |
|------|----------------|
| **Interview tomorrow (any)** | File 1 §1–5, File 4 §1–3 |
| **Service company prep** | Files 1 + 6 |
| **Product company prep** | Files 1 + 2 + 3 + 4 |
| **Senior / backend role** | Files 2 + 3 + 4 + 5 |
| **Complete coverage** | All 7 files in order |
