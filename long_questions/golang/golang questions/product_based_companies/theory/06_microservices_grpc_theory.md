# 🗣️ Theory — Microservices & gRPC in Go
> **Conversational / Spoken Format** | For quick verbal recall in interviews

---

## Q: "What is gRPC and how is it different from REST?"

> *"gRPC is a high-performance RPC framework from Google that uses HTTP/2 as its transport and Protocol Buffers as its data format. Compared to REST: gRPC is significantly faster because Protobuf is binary — much more compact than JSON — and HTTP/2 supports multiplexing multiple requests on one connection. gRPC also has a strict schema defined in `.proto` files and generates client/server code automatically. REST is more flexible and human-readable, works natively in browsers, and has a bigger ecosystem of tooling. In practice, teams use gRPC for internal service-to-service communication where performance matters, and REST for external-facing public APIs where accessibility matters."*

---

## Q: "What is protobuf and why is it used with gRPC?"

> *"Protocol Buffers — protobuf — is Google's binary serialization format. You define your data structures and service contracts in a `.proto` file, then run `protoc` with the Go plugin to generate Go code — structs for your messages, interface definitions for your services, and client stub implementations. The generated code does all the serialization and deserialization. Protobuf is 3 to 10 times smaller and 5 to 10 times faster to serialize and deserialize than JSON. The trade-off is it's binary — you can't just `curl` a gRPC endpoint and read the response. Tools like `grpcurl` exist for that, or you use gRPC reflection."*

---

## Q: "What are gRPC interceptors? How are they like HTTP middleware?"

> *"Interceptors in gRPC are the equivalent of middleware in HTTP frameworks. A unary interceptor wraps a single RPC call — you get the request, can modify the context, call the handler, then process the response or error. A streaming interceptor wraps streaming calls. Common uses: authentication — extract and validate JWT from metadata; logging — log method name, duration, and error; distributed tracing — extract trace context from metadata and inject into span; monitoring — record request counts and latencies. You register interceptors when creating the server: `grpc.NewServer(grpc.ChainUnaryInterceptor(auth, logging, tracing))`."*

---

## Q: "What are the different types of gRPC calls?"

> *"gRPC supports four communication patterns. First, unary — one request, one response, like a regular function call. Second, server streaming — client sends one request, server sends a stream of responses — useful for things like streaming a large dataset or pushing real-time updates. Third, client streaming — client sends a stream of requests, server responds once — good for file uploads. Fourth, bidirectional streaming — both sides send streams simultaneously — useful for chat applications or real-time collaboration. The pattern is defined in the `.proto` file with the `stream` keyword. Each pattern generates different interface methods in the Go code."*

---

## Q: "What is Kafka and how do you use it in a Go microservice architecture?"

> *"Kafka is a distributed event streaming platform used as a high-throughput, durable message queue between services. In a microservice context: one service produces events to a Kafka topic — like 'order.created' — and one or more consumer services subscribe to that topic. Kafka decouples producers from consumers, handles traffic spikes by buffering, and provides replay capability since messages are stored durably. In Go, the most popular Kafka client is `github.com/IBM/sarama`. You create a producer, publish events as byte-serialized JSON or Protobuf messages, and on the consumer side, use a consumer group that distributes partitions across multiple instances automatically."*

---

## Q: "How do you handle service-to-service authentication in gRPC?"

> *"The standard approach is JWT or mTLS. With JWT: the client includes a token in gRPC metadata — similar to an HTTP Authorization header — and a server-side interceptor extracts and validates it on every call. With mTLS — mutual TLS — both client and server present certificates. The server verifies the client's certificate against a trusted CA, and vice versa. mTLS is stronger because it's at the transport layer and doesn't require passing tokens. In Kubernetes environments, service meshes like Istio or Linkerd can handle mTLS transparently — every pod gets a certificate managed by the mesh — so your application code doesn't need to do anything certificate-related."*
