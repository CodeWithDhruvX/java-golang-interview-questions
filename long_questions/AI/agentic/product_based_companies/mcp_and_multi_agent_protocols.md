# MCP & Multi-Agent Communication Protocols – Interview Q&A (Product-Based Companies)

These topics are increasingly asked at AI-native companies (Anthropic, OpenAI, Cohere, Glean) and forward-thinking product teams at larger companies. MCP is rapidly becoming the industry standard for agent-tool integration.

---

## 1. What is MCP (Model Context Protocol)?

**Q: What is Anthropic's Model Context Protocol and why was it created?**

**Answer:**

**MCP (Model Context Protocol)** is an open standard introduced by Anthropic in November 2024 that provides a unified interface for LLMs to connect to external data sources and tools.

### The Problem MCP Solves

Before MCP, every AI application had to build custom integrations for each tool:

```
Without MCP (M × N problem):
  Agent A ────── Custom code ──────► Google Drive
  Agent A ────── Custom code ──────► GitHub
  Agent A ────── Custom code ──────► Slack
  Agent B ────── Custom code ──────► Google Drive (again!)
  Agent B ────── Custom code ──────► Postgres (again!)

  Every new agent × every new tool = new custom integration
  Result: N agents × M tools = N×M bespoke integrations to maintain
```

```
With MCP (M + N problem):
  Agent A ──► MCP Client ──► MCP Server ──► Google Drive
  Agent B ──► MCP Client ──► MCP Server ──► Google Drive (same server!)
  Agent B ──► MCP Client ──► MCP Server ──► GitHub
  
  Write the MCP Server ONCE per tool, any agent can use it
  Result: N agent clients + M tool servers (not N×M)
```

### MCP Architecture

```
┌─────────────────────────────────────────────────────────┐
│                    MCP HOST                              │
│  (Claude Desktop, Cursor, your custom application)       │
│                                                          │
│  ┌──────────────┐     ┌──────────────┐                  │
│  │  MCP Client  │────►│  MCP Client  │                  │
│  │  (Agent A)   │     │  (Agent B)   │                  │
│  └──────┬───────┘     └──────┬───────┘                  │
└─────────┼────────────────────┼────────────────────────  ┘
          │                    │
    ┌─────┴──────┐      ┌──────┴──────┐
    │ MCP Server │      │ MCP Server  │
    │ (GitHub)   │      │ (Postgres)  │
    └────────────┘      └─────────────┘
```

### MCP Primitives

MCP servers expose three types of capabilities:

| Primitive | Description | Example |
|---|---|---|
| **Tools** | Functions the LLM can call (model-controlled) | `create_github_issue`, `run_sql_query` |
| **Resources** | Data the LLM can read (application-controlled) | File contents, DB records, API responses |
| **Prompts** | Pre-built prompt templates (user-controlled) | "Summarize this PR", "Review this SQL" |

---

## 2. How Does MCP Work Under the Hood?

**Q: Explain the MCP communication protocol. How does a client discover and call tools on a server?**

**Answer:**

MCP uses **JSON-RPC 2.0** over one of three transports:
- `stdio` — standard input/output (for local tools, most common)
- **SSE (Server-Sent Events)** — for remote HTTP servers
- **WebSocket** — for bidirectional real-time communication

### Server Discovery & Tool Calling Flow

```
Step 1: Initialization
  Client ──► {method: "initialize", params: {capabilities: {...}}} ──► Server
  Server ◄── {result: {capabilities: {tools: {}, resources: {}}}}   ──► Client

Step 2: Tool Discovery
  Client ──► {method: "tools/list"} ──► Server
  Server ◄── {result: {tools: [
    {name: "read_file", description: "...", inputSchema: {type: "object", ...}},
    {name: "write_file", description: "...", inputSchema: {...}}
  ]}}

Step 3: LLM decides which tool to call → Client executes:
  Client ──► {method: "tools/call", params: {name: "read_file", arguments: {path: "/etc/config"}}}
  Server ◄── {result: {content: [{type: "text", text: "file contents here..."}]}}
```

### Building a Simple MCP Server (Python SDK)

```python
# pip install mcp

from mcp.server import Server
from mcp.server.models import InitializationOptions
from mcp.server.stdio import stdio_server
import mcp.types as types
import json

# Create the MCP server
app = Server("company-data-server")

# Register tools
@app.list_tools()
async def list_tools() -> list[types.Tool]:
    return [
        types.Tool(
            name="get_employee_info",
            description="Retrieve employee information by employee ID from the HR system.",
            inputSchema={
                "type": "object",
                "properties": {
                    "employee_id": {
                        "type": "string",
                        "description": "The employee ID, e.g. 'EMP-1234'"
                    }
                },
                "required": ["employee_id"]
            }
        ),
        types.Tool(
            name="search_company_docs",
            description="Search internal company documentation and policies.",
            inputSchema={
                "type": "object",
                "properties": {
                    "query": {"type": "string", "description": "Search query"},
                    "max_results": {"type": "integer", "default": 5}
                },
                "required": ["query"]
            }
        )
    ]

# Handle tool calls
@app.call_tool()
async def call_tool(name: str, arguments: dict) -> list[types.TextContent]:
    if name == "get_employee_info":
        employee_id = arguments["employee_id"]
        # In production: call your HR API
        employee = {
            "id": employee_id,
            "name": "Rahul Sharma",
            "department": "Engineering",
            "manager": "Priya Patel"
        }
        return [types.TextContent(type="text", text=json.dumps(employee))]
    
    elif name == "search_company_docs":
        query = arguments["query"]
        results = vector_store.similarity_search(query, k=arguments.get("max_results", 5))
        content = "\n---\n".join([doc.page_content for doc in results])
        return [types.TextContent(type="text", text=content)]
    
    else:
        raise ValueError(f"Unknown tool: {name}")

# Run the server on stdio (for local use with Claude Desktop)
async def main():
    async with stdio_server() as (read_stream, write_stream):
        await app.run(read_stream, write_stream, InitializationOptions(
            server_name="company-data-server",
            server_version="1.0.0",
        ))

if __name__ == "__main__":
    import asyncio
    asyncio.run(main())
```

### Connecting Claude to Your MCP Server

```json
// claude_desktop_config.json
{
  "mcpServers": {
    "company-data": {
      "command": "python",
      "args": ["/path/to/your/mcp_server.py"],
      "env": {
        "DATABASE_URL": "postgresql://...",
        "HR_API_KEY": "your-api-key"
      }
    }
  }
}
```

---

## 3. MCP vs. OpenAI Function Calling vs. LangChain Tools

**Q: How does MCP compare to OpenAI function calling and LangChain tools?**

**Answer:**

| Feature | OpenAI Function Calling | LangChain Tools | MCP |
|---|---|---|---|
| **Scope** | OpenAI models only | LangChain ecosystem | Cross-model, cross-framework |
| **Transport** | HTTPS API | In-process Python | stdio, SSE, WebSocket |
| **Tool Discovery** | Manual (you define in code) | Manual (you register tools) | Automatic (`tools/list`) |
| **Reusability** | Per-application | Per-agent | Cross-application (write once) |
| **Standard** | Proprietary | Proprietary | Open standard (growing ecosystem) |
| **Best for** | OpenAI-only stacks | LangChain-heavy stacks | Multi-model, multi-application setups |
| **Adoption** | Universal | LangChain users | Growing rapidly (Claude, Cursor, Cline, etc.) |

**Interview Talking Point:**
> *"We initially built our tools using LangChain's `@tool` decorator, but as we started integrating with more LLM providers and internal applications, we migrated to MCP. Now the same GitHub integration can be used by our Claude-powered IDE assistant, our LangGraph agent, and our internal chatbot — we just wrote the MCP server once."*

---

## 4. Multi-Agent Communication Patterns

**Q: How do agents communicate with each other in a multi-agent system? What patterns exist?**

**Answer:**

There are three primary communication patterns:

### Pattern 1: Hierarchical (Supervisor → Worker)

```python
from langgraph.graph import StateGraph, END
from langchain_openai import ChatOpenAI
from langchain_core.messages import HumanMessage

llm = ChatOpenAI(model="gpt-4o", temperature=0)

# Supervisor decides which worker to call
SUPERVISOR_SYSTEM = """You are a supervisor that routes tasks to specialized workers.
Workers available:
- 'researcher': Gathers information from databases and the web
- 'analyst': Analyzes data and produces insights  
- 'writer': Produces final reports and summaries
- 'FINISH': When the task is complete

Given the current state of the task, respond with ONLY the next worker name."""

def supervisor_node(state: dict) -> dict:
    response = llm.invoke([
        {"role": "system", "content": SUPERVISOR_SYSTEM},
        {"role": "user", "content": f"Task: {state['task']}\nWork done so far: {state.get('results', [])}"}
    ])
    return {"next_worker": response.content.strip()}

def researcher_node(state: dict) -> dict:
    findings = f"[Researcher] Found: market data for {state['task']}"
    return {"results": state.get("results", []) + [findings]}

def analyst_node(state: dict) -> dict:
    insights = f"[Analyst] Insights: trends identified from researcher findings"
    return {"results": state.get("results", []) + [insights]}

def writer_node(state: dict) -> dict:
    report = f"[Writer] Final Report: comprehensive summary of all findings"
    return {"final_output": report}

def route_to_worker(state: dict) -> str:
    next_worker = state.get("next_worker", "FINISH")
    if next_worker == "FINISH":
        return END
    return next_worker

# Build the graph
graph = StateGraph(dict)
graph.add_node("supervisor", supervisor_node)
graph.add_node("researcher", researcher_node)
graph.add_node("analyst", analyst_node)
graph.add_node("writer", writer_node)

graph.set_entry_point("supervisor")
graph.add_conditional_edges("supervisor", route_to_worker)
graph.add_edge("researcher", "supervisor")  # Workers always report back to supervisor
graph.add_edge("analyst", "supervisor")
graph.add_edge("writer", "supervisor")

app = graph.compile()
```

### Pattern 2: Peer-to-Peer (Agent-to-Agent Messaging)

Used in frameworks like **AutoGen** where agents can directly message each other:

```python
# AutoGen example: Two agents collaborate directly
import autogen

config_list = [{"model": "gpt-4o", "api_key": "your-key"}]

# Agent 1: Coder
coder = autogen.AssistantAgent(
    name="Coder",
    system_message="You are an expert Python programmer. Write clean, tested code.",
    llm_config={"config_list": config_list}
)

# Agent 2: Code Reviewer
reviewer = autogen.AssistantAgent(
    name="CodeReviewer",
    system_message="""You are a strict code reviewer. 
    Review code for: bugs, security issues, performance, readability.
    If the code passes all checks, say APPROVED.""",
    llm_config={"config_list": config_list}
)

# Human proxy to initiate
user_proxy = autogen.UserProxyAgent(
    name="UserProxy",
    human_input_mode="NEVER",
    is_termination_msg=lambda msg: "APPROVED" in msg.get("content", ""),
    code_execution_config={"work_dir": "coding", "use_docker": False}
)

# Start the conversation
user_proxy.initiate_chat(
    coder,
    message="Write a Python function to find the longest palindrome in a string."
)
# Coder writes code → Reviewer reviews → Coder fixes → Reviewer approves
```

### Pattern 3: Event-Driven (Message Bus / Kafka)

For large-scale distributed multi-agent systems:

```
                    ┌─────────────────┐
                    │  Kafka / Redis  │  ← Message Bus
                    │  Streams        │
                    └───────┬─────────┘
                            │
          ┌─────────────────┼─────────────────┐
          │                 │                 │
    ┌─────▼──────┐   ┌──────▼─────┐   ┌──────▼──────┐
    │ Ingestion  │   │ Processing │   │  Output     │
    │ Agent      │   │ Agent      │   │  Agent      │
    │            │   │            │   │             │
    │ Reads      │   │ Analyzes   │   │ Sends       │
    │ emails     │   │ sentiment  │   │ reports     │
    └────────────┘   └────────────┘   └─────────────┘

Each agent publishes events to Kafka topics:
  email_ingested → {email_id, content, timestamp}
  analysis_done  → {email_id, sentiment, category}
  report_ready   → {report_id, recipient, content}
```

```python
# Simplified event-driven agent using Redis Streams
import redis
import json
from langchain_openai import ChatOpenAI

r = redis.Redis(host='localhost', port=6379, decode_responses=True)
llm = ChatOpenAI(model="gpt-4o-mini", temperature=0)

class AnalysisAgent:
    """Consumes ingested emails, publishes analysis results"""
    
    def run(self):
        while True:
            # Read from stream (blocks until new message)
            messages = r.xread({"email_ingested": "$"}, block=5000, count=10)
            
            for stream_name, stream_messages in (messages or []):
                for msg_id, data in stream_messages:
                    email_content = data["content"]
                    
                    # Process with LLM
                    analysis = llm.invoke(
                        f"Classify this email's sentiment and category:\n{email_content}"
                    ).content
                    
                    # Publish result to next stream
                    r.xadd("analysis_done", {
                        "email_id": data["email_id"],
                        "analysis": analysis,
                        "processed_by": "AnalysisAgent-v1"
                    })
```

---

## 5. Agent-to-Agent Trust & Authentication

**Q: How do you ensure that one agent can't impersonate another or access data it shouldn't?**

**Answer:**

```python
import jwt
import time
from functools import wraps

SECRET_KEY = "your-secret-key"

def generate_agent_token(agent_id: str, allowed_tools: list[str], ttl_seconds: int = 300) -> str:
    """Issue a short-lived JWT for an agent with specific tool permissions."""
    payload = {
        "agent_id": agent_id,
        "allowed_tools": allowed_tools,    # Principle of Least Privilege
        "issued_at": time.time(),
        "expires_at": time.time() + ttl_seconds
    }
    return jwt.encode(payload, SECRET_KEY, algorithm="HS256")

def verify_agent_token(token: str, required_tool: str) -> bool:
    """Verify an agent's token before allowing tool access."""
    try:
        payload = jwt.decode(token, SECRET_KEY, algorithms=["HS256"])
        
        # Check expiry
        if time.time() > payload["expires_at"]:
            raise PermissionError("Agent token expired")
        
        # Check tool authorization (Principle of Least Privilege)
        if required_tool not in payload["allowed_tools"]:
            raise PermissionError(f"Agent {payload['agent_id']} not authorized for {required_tool}")
        
        return True
    
    except jwt.InvalidTokenError as e:
        raise PermissionError(f"Invalid agent token: {e}")

def authenticated_tool(tool_name: str):
    """Decorator that enforces agent authentication for tool access."""
    def decorator(func):
        @wraps(func)
        def wrapper(*args, agent_token: str = None, **kwargs):
            if not agent_token:
                raise PermissionError("No agent token provided")
            verify_agent_token(agent_token, tool_name)
            return func(*args, **kwargs)
        return wrapper
    return decorator

# Usage
@authenticated_tool("delete_user")  # Only agents explicitly granted this tool can call it
def delete_user(user_id: str, agent_token: str):
    # This tool requires explicit authorization
    print(f"Deleting user {user_id}")

# Researcher agent: can only read
researcher_token = generate_agent_token(
    agent_id="researcher-001",
    allowed_tools=["search_database", "read_documents"]
)

# Executor agent: can write (restricted)
executor_token = generate_agent_token(
    agent_id="executor-001",
    allowed_tools=["write_record", "send_notification"]
    # Note: "delete_user" NOT in this list
)
```

**Key Interview Point:** In multi-agent systems, **each agent should have its own identity and minimal privilege set**. A compromised worker agent should not be able to escalate to actions above its permission level — this is the Principle of Least Privilege applied to agents.
