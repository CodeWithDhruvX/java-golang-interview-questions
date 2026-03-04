# Agentic AI – Practical Coding Examples with Answers (Service-Based Companies)

These are hands-on coding questions asked at service-based companies (TCS, Infosys, Wipro, Accenture, Capgemini, Cognizant). The focus is on building working prototypes using LangChain, OpenAI, or similar tools — not deep infrastructure.

---

## 1. Build a Simple Chatbot Agent with Memory (LangChain)

**Question:** Create a conversational chatbot using LangChain that remembers the conversation history.

**Answer:**

```python
from langchain_openai import ChatOpenAI
from langchain.memory import ConversationBufferMemory
from langchain.chains import ConversationChain

# Initialize LLM
llm = ChatOpenAI(model="gpt-3.5-turbo", temperature=0.7)

# Memory to store conversation history
memory = ConversationBufferMemory()

# Create a conversation chain
conversation = ConversationChain(
    llm=llm,
    memory=memory,
    verbose=True  # Shows the full prompt being sent
)

# Simulate a conversation
print(conversation.predict(input="Hi! My name is Dhruv."))
# Output: "Hello Dhruv! How can I help you today?"

print(conversation.predict(input="What is my name?"))
# Output: "Your name is Dhruv!" — It remembers!

print(conversation.predict(input="What programming languages do you know?"))
```

**Key Concepts:**
- `ConversationBufferMemory` stores the full chat history
- For very long conversations, use `ConversationSummaryMemory` (summarizes old messages)
- `ConversationChain` automatically injects memory into every prompt

---

## 2. Build a Custom Tool and Run an Agent (LangChain)

**Question:** How do you create a custom tool and use it with a LangChain agent?

**Answer:**

```python
from langchain_openai import ChatOpenAI
from langchain_core.tools import tool
from langchain.agents import AgentExecutor, create_openai_tools_agent
from langchain_core.prompts import ChatPromptTemplate, MessagesPlaceholder

# Step 1: Create a custom tool using the @tool decorator
@tool
def get_employee_info(employee_id: str) -> str:
    """Retrieves information about an employee by their employee ID."""
    # Simulating a database call
    employees = {
        "EMP001": {"name": "Rahul Sharma", "department": "Engineering", "salary": 80000},
        "EMP002": {"name": "Priya Patel", "department": "HR", "salary": 60000},
    }
    employee = employees.get(employee_id)
    if employee:
        return str(employee)
    return f"Employee {employee_id} not found"

@tool
def calculate_tax(salary: float) -> str:
    """Calculates the income tax for a given annual salary in INR."""
    if salary <= 250000:
        tax = 0
    elif salary <= 500000:
        tax = (salary - 250000) * 0.05
    else:
        tax = 12500 + (salary - 500000) * 0.20
    return f"Tax for salary ₹{salary}: ₹{tax:.2f}"

# Step 2: Set up the LLM and tools list
llm = ChatOpenAI(model="gpt-4o", temperature=0)
tools = [get_employee_info, calculate_tax]

# Step 3: Create the prompt
prompt = ChatPromptTemplate.from_messages([
    ("system", "You are an HR assistant. Use the provided tools to answer queries."),
    ("human", "{input}"),
    MessagesPlaceholder(variable_name="agent_scratchpad"),
])

# Step 4: Create and run the agent
agent = create_openai_tools_agent(llm, tools, prompt)
agent_executor = AgentExecutor(agent=agent, tools=tools, verbose=True)

# Run queries
result = agent_executor.invoke({
    "input": "What is the tax for employee EMP001?"
})
print(result["output"])
# Agent will: 1) call get_employee_info("EMP001"), 
#              2) extract salary=80000, 
#              3) call calculate_tax(80000),
#              4) give final answer
```

**Key Concepts:**
- The `@tool` decorator automatically generates the JSON schema from the docstring
- The docstring is crucial — the LLM reads it to decide when to use the tool
- `verbose=True` lets you see the agent's reasoning: Thought → Action → Observation

---

## 3. RAG (Retrieval-Augmented Generation) with Documents

**Question:** Build a simple RAG system that can answer questions from uploaded PDF documents.

**Answer:**

```python
from langchain_community.document_loaders import PyPDFLoader
from langchain.text_splitter import RecursiveCharacterTextSplitter
from langchain_openai import OpenAIEmbeddings, ChatOpenAI
from langchain_community.vectorstores import Chroma
from langchain.chains import RetrievalQA

# Step 1: Load the document
loader = PyPDFLoader("company_policy.pdf")
documents = loader.load()
print(f"Loaded {len(documents)} pages")

# Step 2: Split into chunks
text_splitter = RecursiveCharacterTextSplitter(
    chunk_size=1000,    # Each chunk = 1000 characters
    chunk_overlap=200   # 200 character overlap to avoid splitting context
)
chunks = text_splitter.split_documents(documents)
print(f"Created {len(chunks)} chunks")

# Step 3: Create embeddings and store in vector database
embeddings = OpenAIEmbeddings()
vector_store = Chroma.from_documents(
    documents=chunks,
    embedding=embeddings,
    persist_directory="./chroma_db"  # Save to disk
)

# Step 4: Create a retrieval chain
llm = ChatOpenAI(model="gpt-3.5-turbo", temperature=0)
qa_chain = RetrievalQA.from_chain_type(
    llm=llm,
    chain_type="stuff",    # "stuff" = puts all retrieved docs into one prompt
    retriever=vector_store.as_retriever(search_kwargs={"k": 3}),  # Top 3 results
    return_source_documents=True
)

# Step 5: Ask questions
result = qa_chain.invoke({"query": "What is the company's leave policy?"})
print("Answer:", result["result"])
print("Sources:", [doc.page_content[:100] for doc in result["source_documents"]])
```

**Key Concepts:**
- Chunking strategy matters: too small = context loss, too large = dilutes relevance
- `chunk_overlap` prevents important information being split across chunks
- `k=3` means fetch top 3 most similar chunks
- Always return source documents to verify the answer is grounded in reality

---

## 4. Multi-Agent System with CrewAI (Beginner-Friendly)

**Question:** How do you build a multi-agent system where a researcher and a writer collaborate?

**Answer:**

```python
from crewai import Agent, Task, Crew
from langchain_openai import ChatOpenAI

llm = ChatOpenAI(model="gpt-3.5-turbo", temperature=0.7)

# Step 1: Define agents with their roles
researcher = Agent(
    role="Research Analyst",
    goal="Research and gather accurate information on the given topic",
    backstory="You are an experienced researcher who finds detailed, accurate information.",
    llm=llm,
    verbose=True
)

writer = Agent(
    role="Content Writer",
    goal="Write clear, engaging content based on provided research",
    backstory="You are a professional writer who creates well-structured blog posts.",
    llm=llm,
    verbose=True
)

# Step 2: Define tasks for each agent
research_task = Task(
    description="Research the topic: Benefits of using AI agents in software development. "
                "Find 5 key benefits with real-world examples.",
    agent=researcher,
    expected_output="A list of 5 key benefits with examples"
)

writing_task = Task(
    description="Using the research provided, write a 300-word blog post about "
                "AI agents in software development.",
    agent=writer,
    expected_output="A 300-word blog post",
    context=[research_task]  # Writer gets researcher's output as input
)

# Step 3: Create and run the crew
crew = Crew(
    agents=[researcher, writer],
    tasks=[research_task, writing_task],
    verbose=True
)

result = crew.kickoff()
print("\n=== FINAL BLOG POST ===")
print(result)
```

**Key Concepts:**
- `context=[research_task]` → writing_task automatically receives researcher's output
- Each agent has a clear `role`, `goal`, and `backstory` — these shape behavior
- Tasks execute sequentially by default (researcher first, then writer)
- CrewAI abstracts away the complex orchestration logic

---

## 5. OpenAI Function Calling — Simple Example

**Question:** Demonstrate how to use OpenAI's function calling feature to call a weather API.

**Answer:**

```python
import json
import openai

client = openai.OpenAI()

# Step 1: Define function schema
functions = [
    {
        "name": "get_current_weather",
        "description": "Get the current weather for a city",
        "parameters": {
            "type": "object",
            "properties": {
                "city": {
                    "type": "string",
                    "description": "City name, e.g., Bangalore, Mumbai"
                },
                "unit": {
                    "type": "string",
                    "enum": ["celsius", "fahrenheit"]
                }
            },
            "required": ["city"]
        }
    }
]

# Step 2: Stub function (in production this calls a real API)
def get_current_weather(city: str, unit: str = "celsius") -> str:
    weather_data = {
        "Bangalore": {"temp": 24, "condition": "Cloudy"},
        "Mumbai": {"temp": 32, "condition": "Humid"},
        "Delhi": {"temp": 18, "condition": "Foggy"}
    }
    data = weather_data.get(city, {"temp": 28, "condition": "Unknown"})
    return json.dumps({
        "city": city,
        "temperature": data["temp"],
        "unit": unit,
        "condition": data["condition"]
    })

# Step 3: Call the model
messages = [{"role": "user", "content": "What's the weather in Bangalore?"}]

response = client.chat.completions.create(
    model="gpt-3.5-turbo",
    messages=messages,
    functions=functions,
    function_call="auto"
)

msg = response.choices[0].message

# Step 4: If the model wants to call a function, execute it
if msg.function_call:
    function_name = msg.function_call.name
    arguments = json.loads(msg.function_call.arguments)
    
    print(f"Model wants to call: {function_name}({arguments})")
    
    # Execute the function
    function_result = get_current_weather(**arguments)
    
    # Step 5: Send result back to model
    messages.append(msg)
    messages.append({
        "role": "function",
        "name": function_name,
        "content": function_result
    })
    
    final_response = client.chat.completions.create(
        model="gpt-3.5-turbo",
        messages=messages
    )
    print("Final Answer:", final_response.choices[0].message.content)
```

**Key Concepts:**
- The LLM does NOT execute the function — your code does
- The LLM outputs a structured JSON with the function name and arguments
- You execute the function, pass the result back, and the LLM generates a natural language response
- This pattern is the foundation for all agentic tool use

---

## 6. Handling Agent Failures and Error Recovery

**Question:** How do you add error handling and retry logic to a LangChain agent?

**Answer:**

```python
from langchain_openai import ChatOpenAI
from langchain_core.tools import tool
from langchain.agents import AgentExecutor, create_openai_tools_agent
from langchain_core.prompts import ChatPromptTemplate, MessagesPlaceholder
import time

llm = ChatOpenAI(model="gpt-3.5-turbo", temperature=0)

@tool
def fetch_stock_price(symbol: str) -> str:
    """Fetches the current stock price for a given stock symbol like RELIANCE, TCS."""
    # Simulating occasional API failures
    import random
    if random.random() < 0.3:  # 30% chance of failure
        raise Exception(f"API rate limit exceeded for {symbol}")
    return f"{symbol}: ₹{random.randint(1000, 5000)}"

tools = [fetch_stock_price]

prompt = ChatPromptTemplate.from_messages([
    ("system", "You are a stock market assistant."),
    ("human", "{input}"),
    MessagesPlaceholder(variable_name="agent_scratchpad"),
])

agent = create_openai_tools_agent(llm, tools, prompt)

# Key error handling parameters
agent_executor = AgentExecutor(
    agent=agent,
    tools=tools,
    verbose=True,
    max_iterations=5,           # Stop after 5 steps (prevents infinite loops)
    max_execution_time=30,      # Stop after 30 seconds
    handle_parsing_errors=True, # Don't crash on malformed LLM output
    return_intermediate_steps=True  # Return all steps for debugging
)

# Retry wrapper for transient failures
def run_with_retry(input_text: str, max_retries: int = 3) -> str:
    last_error = None
    for attempt in range(max_retries):
        try:
            result = agent_executor.invoke({"input": input_text})
            return result["output"]
        except Exception as e:
            last_error = e
            wait_time = 2 ** attempt  # Exponential backoff: 1s, 2s, 4s
            print(f"Attempt {attempt + 1} failed: {e}. Retrying in {wait_time}s...")
            time.sleep(wait_time)
    
    return f"Failed after {max_retries} attempts. Last error: {last_error}"

result = run_with_retry("What is the current price of RELIANCE stock?")
print(result)
```

**Key Concepts:**
- `max_iterations=5` — prevents the agent from running indefinitely
- `handle_parsing_errors=True` — catches malformed JSON from the LLM
- Exponential backoff (`2 ** attempt`) is the standard pattern for API retries
- `return_intermediate_steps=True` helps with debugging and auditability

---

## 7. Building a Simple Agent with LangGraph (Beginner Level)

**Question:** Create a basic LangGraph workflow where an agent chooses between answering directly or using a search tool.

**Answer:**

```python
from typing import TypedDict, List
from langgraph.graph import StateGraph, END
from langchain_openai import ChatOpenAI
from langchain_community.tools import DuckDuckGoSearchRun
from langchain_core.messages import HumanMessage, AIMessage, BaseMessage

llm = ChatOpenAI(model="gpt-3.5-turbo", temperature=0)
search_tool = DuckDuckGoSearchRun()

# Define the state (data that flows through the graph)
class State(TypedDict):
    messages: List[BaseMessage]
    needs_search: bool
    search_result: str

# Node 1: LLM decides if it needs to search
def router_node(state: State) -> State:
    user_message = state["messages"][-1].content
    
    # Ask LLM if it needs to search
    decision = llm.invoke([
        HumanMessage(content=f"""Does this question require searching the internet for current information?
        Question: {user_message}
        Reply only YES or NO.""")
    ])
    
    needs_search = "YES" in decision.content.upper()
    return {"needs_search": needs_search}

# Node 2: Search if needed
def search_node(state: State) -> State:
    user_message = state["messages"][-1].content
    result = search_tool.run(user_message)
    return {"search_result": result}

# Node 3: Generate final answer
def answer_node(state: State) -> State:
    user_message = state["messages"][-1].content
    
    if state.get("search_result"):
        # Answer using search results
        response = llm.invoke([
            HumanMessage(content=f"Question: {user_message}\n\nSearch results: {state['search_result']}\n\nAnswer based on the search results.")
        ])
    else:
        # Answer from knowledge
        response = llm.invoke(state["messages"])
    
    new_messages = state["messages"] + [AIMessage(content=response.content)]
    return {"messages": new_messages}

# Routing function
def decide_next_step(state: State) -> str:
    return "search" if state["needs_search"] else "answer"

# Build the graph
graph = StateGraph(State)
graph.add_node("router", router_node)
graph.add_node("search", search_node)
graph.add_node("answer", answer_node)

graph.set_entry_point("router")
graph.add_conditional_edges("router", decide_next_step)
graph.add_edge("search", "answer")
graph.add_edge("answer", END)

app = graph.compile()

# Run it
result = app.invoke({
    "messages": [HumanMessage(content="Who won the 2024 IPL?")],
    "needs_search": False,
    "search_result": ""
})

print(result["messages"][-1].content)
```

**Key Concepts:**
- `StateGraph` is a directed graph where each node transforms the state
- `conditional_edges` implements branching logic (if-else at the graph level)
- Every node receives the full state and returns only what it changed
- LangGraph makes the agent's flow explicit and debuggable
