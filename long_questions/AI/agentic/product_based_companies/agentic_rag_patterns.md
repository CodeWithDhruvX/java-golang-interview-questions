# Agentic RAG Patterns – Interview Q&A (Product-Based Companies)

Standard RAG (Retrieve → Augment → Generate) is a solved problem for most interviewers. What differentiates senior candidates is knowledge of **Agentic RAG** — where the agent has the intelligence to adapt its retrieval strategy based on the query. This is heavily tested at AI-native companies (Cohere, Databricks, Glean, Notion AI) and research teams at Google and Microsoft.

---

## 1. What Is the Difference Between Naive RAG, Advanced RAG, and Agentic RAG?

**Answer:**

| Level | Description | Limitation |
|---|---|---|
| **Naive RAG** | Query → Vector Search → Stuff chunks into prompt → LLM answers | Fixed retrieval, no feedback, no retry |
| **Advanced RAG** | Adds query rewriting, re-ranking, hybrid search | Still linear, can't change strategy mid-flight |
| **Agentic RAG** | LLM *decides* how to retrieve, evaluates result quality, and adapts | Higher cost and latency; best for complex multi-hop queries |

```
Naive RAG:   Query ──► Retrieve ──► Generate

Agentic RAG: Query ──► Plan retrieval strategy
                            │
                       ┌────▼──────────────────────┐
                       │  Is retrieval sufficient? │
                       └────┬─────────────┬────────┘
                       YES  │             │  NO
                            ▼             ▼
                         Generate    Re-query / try
                                     different source
```

---

## 2. What is Adaptive RAG?

**Q: Explain Adaptive RAG. When does it improve over standard RAG?**

**Answer:**

**Adaptive RAG** uses a classifier (usually a small fine-tuned model or the LLM itself) to route each query to the most appropriate retrieval strategy *before* retrieval happens.

### Routing Strategies

```python
from langchain_openai import ChatOpenAI
from langchain_core.messages import HumanMessage

router_llm = ChatOpenAI(model="gpt-4o-mini", temperature=0)

def route_query(question: str) -> str:
    """
    Routes query to:
    - 'no_retrieval'  : factual Q the LLM knows (e.g., "What is Python?")
    - 'single_hop'    : one document/source needed
    - 'multi_hop'     : requires combining info from multiple sources
    - 'web_search'    : needs real-time data (post-cutoff)
    """
    response = router_llm.invoke([
        HumanMessage(content=f"""Classify this question for retrieval strategy:
        Question: {question}
        
        Choose one: no_retrieval / single_hop / multi_hop / web_search
        Respond with ONLY the category.""")
    ])
    return response.content.strip().lower()

# Router dispatch
def adaptive_rag(question: str) -> str:
    strategy = route_query(question)
    
    if strategy == "no_retrieval":
        return llm.invoke(question).content
    
    elif strategy == "single_hop":
        docs = vector_store.similarity_search(question, k=3)
        return generate_answer(question, docs)
    
    elif strategy == "multi_hop":
        return multi_hop_agent(question)   # Decompose → retrieve per sub-question → synthesize
    
    elif strategy == "web_search":
        results = web_search_tool.run(question)
        return generate_answer(question, results)

# Usage
print(adaptive_rag("What is the capital of France?"))          # → no_retrieval
print(adaptive_rag("What is our company's leave policy?"))     # → single_hop
print(adaptive_rag("Who won the 2025 IPL?"))                   # → web_search
print(adaptive_rag("Compare Q3 revenue across all divisions")) # → multi_hop
```

### When to Use Adaptive RAG

- **High-traffic systems** where 40%+ of queries don't need retrieval (saves cost)
- **Mixed-domain agents** where some domains need web search and others need internal docs
- Routing accuracy needs to be evaluated separately as a classifier task

---

## 3. What is Self-RAG?

**Q: Explain the Self-RAG paper's approach. How does an agent decide if retrieval is even necessary?**

**Answer:**

**Self-RAG** (Asai et al., 2023) teaches the LLM to use special **reflection tokens** to self-assess its own outputs:

| Token | Meaning |
|---|---|
| `[Retrieve]` | The model decides it needs to retrieve before answering |
| `[IsREL]` | The retrieved document IS relevant to the question |
| `[IsREL=No]` | The retrieved document is NOT relevant — retrieve again |
| `[IsSUP]` | The generated response IS supported by the retrieved context |
| `[IsSUP=No]` | The response contradicts the context — regenerate |
| `[IsUSE]` | The final response is useful to the user |

### Simplified Implementation (Mimicking Self-RAG behavior)

```python
from langchain_openai import ChatOpenAI
from langchain_core.messages import HumanMessage, SystemMessage

llm = ChatOpenAI(model="gpt-4o", temperature=0)

def self_rag(question: str, vector_store, max_retries: int = 2) -> str:
    
    # Step 1: Does the LLM need retrieval?
    needs_retrieval_check = llm.invoke([
        SystemMessage(content="Respond only YES or NO."),
        HumanMessage(content=f"Do you need external documents to answer this accurately?\nQuestion: {question}")
    ])
    
    if "NO" in needs_retrieval_check.content.upper():
        return llm.invoke([HumanMessage(content=question)]).content
    
    for attempt in range(max_retries):
        # Step 2: Retrieve
        docs = vector_store.similarity_search(question, k=3)
        context = "\n\n".join([d.page_content for d in docs])
        
        # Step 3: Relevance check (IsREL)
        relevance_check = llm.invoke([
            SystemMessage(content="Respond only RELEVANT or NOT_RELEVANT."),
            HumanMessage(content=f"Is this context relevant to the question?\nQuestion: {question}\nContext: {context[:500]}")
        ])
        
        if "NOT_RELEVANT" in relevance_check.content.upper():
            # Re-query with rewritten question
            question = rewrite_query(question)  # Query rewriting
            continue
        
        # Step 4: Generate answer
        answer = llm.invoke([
            SystemMessage(content="Answer using ONLY the provided context. If context is insufficient, say so."),
            HumanMessage(content=f"Context:\n{context}\n\nQuestion: {question}")
        ]).content
        
        # Step 5: Faithfulness check (IsSUP)
        support_check = llm.invoke([
            SystemMessage(content="Respond only SUPPORTED or NOT_SUPPORTED."),
            HumanMessage(content=f"Is this answer fully supported by the context?\nContext: {context[:500]}\nAnswer: {answer}")
        ])
        
        if "SUPPORTED" in support_check.content.upper():
            return answer
        
        # Not supported → retry with stricter prompt
        question = f"{question} (Note: previous answer was not grounded. Be more conservative.)"
    
    return "I could not generate a reliably grounded answer for this question."

def rewrite_query(original_query: str) -> str:
    response = llm.invoke([
        HumanMessage(content=f"Rewrite this search query to be more specific:\n{original_query}")
    ])
    return response.content
```

**Key Interview Point:** Self-RAG is expensive (multiple LLM calls per query) but produces the most faithful answers. Use it for **high-stakes domains** like legal, medical, or financial RAG where hallucinations are unacceptable.

---

## 4. What is Corrective RAG (CRAG)?

**Q: What is CRAG and how does it improve retrieval robustness?**

**Answer:**

**Corrective RAG** (Yan et al., 2024) adds a **retrieval evaluator** that scores the quality of retrieved documents. Based on the score, it takes one of three corrective actions:

```
Retrieved Documents
        │
        ▼
  ┌─────────────┐
  │  Evaluator  │  Scores relevance: HIGH / LOW / AMBIGUOUS
  └──────┬──────┘
         │
    ┌────┴──────────────────┐
    │                       │
 HIGH (>0.8)         LOW (<0.5) or AMBIGUOUS
    │                       │
    ▼                       ▼
 Use docs directly    Web search to
 for generation       supplement/replace
```

### CRAG Implementation with LangGraph

```python
from typing import TypedDict, List
from langgraph.graph import StateGraph, END
from langchain_openai import ChatOpenAI
from langchain_community.tools import DuckDuckGoSearchRun

llm = ChatOpenAI(model="gpt-4o", temperature=0)
web_search = DuckDuckGoSearchRun()

class CRAGState(TypedDict):
    question: str
    documents: List[str]
    relevance_score: float
    web_results: str
    final_answer: str

def retrieve_node(state: CRAGState) -> CRAGState:
    docs = vector_store.similarity_search(state["question"], k=4)
    return {"documents": [d.page_content for d in docs]}

def evaluate_relevance_node(state: CRAGState) -> CRAGState:
    """Score how relevant the retrieved docs are (0.0 to 1.0)"""
    context = "\n".join(state["documents"][:2])  # Sample first 2 docs
    
    response = llm.invoke([
        HumanMessage(content=f"""Rate how relevant these documents are to answering the question.
        Question: {state['question']}
        Documents: {context[:600]}
        
        Respond with a number between 0.0 (completely irrelevant) and 1.0 (perfectly relevant).
        Respond with ONLY the number.""")
    ])
    
    try:
        score = float(response.content.strip())
    except ValueError:
        score = 0.5  # Default to ambiguous if parsing fails
    
    return {"relevance_score": score}

def web_search_node(state: CRAGState) -> CRAGState:
    results = web_search.run(state["question"])
    return {"web_results": results}

def generate_node(state: CRAGState) -> CRAGState:
    # Use docs if highly relevant, web results if not, or combine if ambiguous
    score = state["relevance_score"]
    
    if score >= 0.8:
        context = "\n".join(state["documents"])
    elif score <= 0.4:
        context = state.get("web_results", "No web results available")
    else:  # Ambiguous: combine both
        doc_context = "\n".join(state["documents"])
        web_context = state.get("web_results", "")
        context = f"Internal Documents:\n{doc_context}\n\nWeb Search:\n{web_context}"
    
    answer = llm.invoke([
        HumanMessage(content=f"Answer based on this context:\n{context}\n\nQuestion: {state['question']}")
    ]).content
    
    return {"final_answer": answer}

# Routing based on relevance score
def route_after_evaluation(state: CRAGState) -> str:
    score = state["relevance_score"]
    if score >= 0.8:
        return "generate"       # Docs are good, use directly
    else:
        return "web_search"     # Need to supplement or replace

# Build CRAG graph
graph = StateGraph(CRAGState)
graph.add_node("retrieve", retrieve_node)
graph.add_node("evaluate", evaluate_relevance_node)
graph.add_node("web_search", web_search_node)
graph.add_node("generate", generate_node)

graph.set_entry_point("retrieve")
graph.add_edge("retrieve", "evaluate")
graph.add_conditional_edges("evaluate", route_after_evaluation)
graph.add_edge("web_search", "generate")
graph.add_edge("generate", END)

crag_app = graph.compile()

result = crag_app.invoke({
    "question": "What is our company's policy on remote work?",
    "documents": [], "relevance_score": 0.0,
    "web_results": "", "final_answer": ""
})
print(result["final_answer"])
```

---

## 5. When to Use Each Pattern — Summary

| Pattern | When to Use | Cost | Latency |
|---|---|---|---|
| **Naive RAG** | Simple single-doc Q&A, low stakes | Low | Low |
| **Adaptive RAG** | Mixed query types, high traffic | Medium | Medium |
| **Self-RAG** | Medical/legal/financial — faithfulness critical | High | High |
| **CRAG** | Docs may be outdated or incomplete; need web fallback | Medium-High | Medium |
| **Multi-hop Agent** | Complex research, comparing across multiple sources | High | High |

### Interview Answer to "How would you pick a RAG strategy?"

> *"I start by understanding the failure mode I'm most worried about. If hallucination is unacceptable (medical), I'd go Self-RAG with faithfulness checking. If our internal knowledge base is often stale, I'd go CRAG with a web fallback. For a mixed-query system where 40% of questions don't need retrieval at all, I'd add Adaptive RAG routing to save cost. For simple use cases, Naive RAG with re-ranking is usually sufficient and much cheaper."*

---

## 6. HyDE — Hypothetical Document Embedding

**Q: What is HyDE and when is it better than standard query embedding?**

**Answer:**

**HyDE (Hypothetical Document Embeddings)** — instead of embedding the raw user question, you ask the LLM to generate a *hypothetical answer document*, then embed THAT for retrieval.

```python
from langchain_openai import ChatOpenAI, OpenAIEmbeddings

llm = ChatOpenAI(model="gpt-4o-mini", temperature=0.3)
embeddings = OpenAIEmbeddings()

def hyde_retrieve(question: str, vector_store, k: int = 4):
    # Step 1: Generate a hypothetical answer (with hallucinations — that's OK!)
    hypothetical_doc = llm.invoke(
        f"Write a detailed answer to this question as if you were an expert document:\n{question}"
    ).content
    
    # Step 2: Embed the hypothetical answer (not the question)
    # Hypothetical answer uses domain-specific vocabulary → better matches real docs
    hypo_embedding = embeddings.embed_query(hypothetical_doc)
    
    # Step 3: Retrieve using the hypothetical embedding
    docs = vector_store.similarity_search_by_vector(hypo_embedding, k=k)
    
    # Step 4: Generate real answer from retrieved docs
    context = "\n\n".join([d.page_content for d in docs])
    return llm.invoke(
        f"Answer this question based on the context:\nContext: {context}\nQuestion: {question}"
    ).content

# Why It Works:
# Query: "What is the sick leave policy?"
# HyDE: "Employees are entitled to 12 days of paid sick leave per year..."
# → The hypothetical answer's vocabulary matches what's actually in the HR policy document
# → Better retrieval than embedding the short question directly
```

**When HyDE Wins:** Short, ambiguous queries where the question vocabulary doesn't match the document vocabulary. **When HyDE Hurts:** Very factual queries where the LLM's hallucinated "answer" leads retrieval in the wrong direction.
