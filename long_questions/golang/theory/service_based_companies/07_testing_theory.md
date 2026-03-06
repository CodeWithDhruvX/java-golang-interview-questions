# 🗣️ Theory — Testing in Go
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "How do you write unit tests in Go?"

> *"Go has a built-in testing framework in the `testing` package — no third-party tools needed. Test files end with `_test.go` and live alongside the source files. Test functions start with `Test` and accept `*testing.T`. Inside, you call methods like `t.Errorf` to log failures and continue, or `t.Fatalf` to stop the test immediately on failure. You run them with `go test ./...`. The beauty of Go's approach is that tests are just regular Go code — there's no magic test framework to learn."*

---

## Q: "What is table-driven testing and why is it the idiomatic way?"

> *"Table-driven testing is the preferred Go pattern for writing tests with multiple inputs and expected outputs. You define a slice of test cases — each case has fields for input, expected output, and usually a name. Then you loop over the slice, calling `t.Run(tt.name, func(t *testing.T) { ... })` for each case. Using `t.Run` creates subtests that appear as separate entries in output and can be run individually. The advantage is: adding new test scenarios is just adding a new entry to the table — you don't write a whole new test function. It also keeps test data separated from test logic."*

---

## Q: "What is `testify` and why do people use it?"

> *"Testify is the most popular third-party testing library in Go. It provides `assert` and `require` packages that give you readable assertion functions — instead of `if got != want { t.Errorf(...) }`, you write `assert.Equal(t, want, got)`. The difference between `assert` and `require`: assert logs the failure and the test continues, require stops the test immediately on failure. Testify also has a `mock` package for generating mocks from interfaces, and a `suite` package if you need setup/teardown across a group of tests. Most professional Go projects use testify."*

---

## Q: "How do you test HTTP handlers? What is `httptest`?"

> *"The `net/http/httptest` package lets you test HTTP handlers without starting a real server. You create a fake request with `http.NewRequest()`, a response recorder with `httptest.NewRecorder()`, and then call your handler function directly with those. After the call, you inspect the recorder's status code and body. This is fast, isolated, and doesn't require network access. For integration tests, you can use `httptest.NewServer()` which starts a real server on a random port — you get back a URL and can make real HTTP requests against it, then shut it down after the test."*

---

## Q: "How do you mock interfaces for testing?"

> *"The standard pattern is to define your dependencies as interfaces and inject them. Your service takes a `UserRepository` interface, not a concrete `*PostgresRepo`. In tests, you pass a mock implementation. For simple cases, you just write a hand-crafted struct that implements the interface. For more complex cases where you want to set expectations — like 'this method should be called exactly twice' — you use `testify/mock`. You define a struct that embeds `mock.Mock`, implement each method using `m.Called(args)`, then configure expectations in the test setup with `mockRepo.On('GetByID', 1).Return(user, nil)`."*

---

## Q: "How do you measure test coverage and what is a good target?"

> *"You run `go test -coverprofile=coverage.out ./...` which produces a coverage profile. Then `go tool cover -html=coverage.out` opens a browser showing which lines are covered — green for covered, red for not. For a specific percentage, use `go test -cover ./...`. As for targets — 100% is not the goal. Chasing 100% leads to tests that exist just to cover lines, not to verify behavior. I'd say aim for 70–80% on core business logic. Some packages have a lot of boilerplate or infrastructure code that's hard to test in isolation and it's fine to leave those at lower coverage."*
