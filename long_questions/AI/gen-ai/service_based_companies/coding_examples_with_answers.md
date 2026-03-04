# Gen-AI Coding Examples with Answers — Service-Based Companies

These are **hands-on implementation questions** that test whether you can actually build GenAI applications using Python, LangChain, OpenAI API, and related tools. Service-based clients expect engineers to deliver working prototypes rapidly.

---

## 1. Basic LLM API Usage

**Q1: Write a Python function that calls the OpenAI Chat API and returns a response.**

```python
from openai import OpenAI

client = OpenAI(api_key="YOUR_API_KEY")

def ask_llm(user_message: str, system_prompt: str = "You are a helpful assistant.") -> str:
    response = client.chat.completions.create(
        model="gpt-4o",
        messages=[
            {"role": "system", "content": system_prompt},
            {"role": "user", "content": user_message}
        ],
        temperature=0.7,
        max_tokens=1024
    )
    return response.choices[0].message.content

# Usage
answer = ask_llm("What is the capital of India?")
print(answer)
```

**Q2: How do you implement streaming LLM responses in Python?**

```python
from openai import OpenAI

client = OpenAI(api_key="YOUR_API_KEY")

def stream_llm(user_message: str):
    stream = client.chat.completions.create(
        model="gpt-4o",
        messages=[{"role": "user", "content": user_message}],
        stream=True
    )
    for chunk in stream:
        delta = chunk.choices[0].delta
        if delta.content:
            print(delta.content, end="", flush=True)

stream_llm("Write a haiku about the monsoon.")
```

---

## 2. RAG Pipeline Implementation

**Q3: Build a minimal RAG pipeline from scratch using LangChain.**

```python
from langchain_openai import ChatOpenAI, OpenAIEmbeddings
from langchain_chroma import Chroma
from langchain.text_splitter import RecursiveCharacterTextSplitter
from langchain_core.prompts import ChatPromptTemplate
from langchain_core.runnables import RunnablePassthrough
from langchain_core.output_parsers import StrOutputParser
from langchain_community.document_loaders import TextLoader

# 1. Load and split documents
loader = TextLoader("company_policy.txt")
docs = loader.load()

splitter = RecursiveCharacterTextSplitter(chunk_size=500, chunk_overlap=50)
chunks = splitter.split_documents(docs)

# 2. Create vector store
embeddings = OpenAIEmbeddings(model="text-embedding-3-small")
vectorstore = Chroma.from_documents(chunks, embedding=embeddings)

# 3. Create retriever
retriever = vectorstore.as_retriever(search_kwargs={"k": 4})

# 4. Prompt template
prompt = ChatPromptTemplate.from_template("""
Answer the question based ONLY on the following context.
If the context doesn't contain the answer, say "I don't know based on available information."

Context:
{context}

Question: {question}
""")

# 5. Build chain
llm = ChatOpenAI(model="gpt-4o-mini", temperature=0)

def format_docs(docs):
    return "\n\n".join(doc.page_content for doc in docs)

rag_chain = (
    {"context": retriever | format_docs, "question": RunnablePassthrough()}
    | prompt
    | llm
    | StrOutputParser()
)

# 6. Query
answer = rag_chain.invoke("What is the leave policy for new employees?")
print(answer)
```

---

## 3. Memory & Conversation History

**Q4: How do you add conversation memory to a LangChain chatbot?**

```python
from langchain_openai import ChatOpenAI
from langchain_core.prompts import ChatPromptTemplate, MessagesPlaceholder
from langchain_core.chat_history import InMemoryChatMessageHistory
from langchain_core.runnables.history import RunnableWithMessageHistory

llm = ChatOpenAI(model="gpt-4o-mini")

prompt = ChatPromptTemplate.from_messages([
    ("system", "You are a helpful assistant. Answer concisely."),
    MessagesPlaceholder(variable_name="chat_history"),
    ("human", "{question}")
])

chain = prompt | llm

# Store for multiple sessions
session_store = {}

def get_session_history(session_id: str) -> InMemoryChatMessageHistory:
    if session_id not in session_store:
        session_store[session_id] = InMemoryChatMessageHistory()
    return session_store[session_id]

chain_with_history = RunnableWithMessageHistory(
    chain,
    get_session_history,
    input_messages_key="question",
    history_messages_key="chat_history"
)

# Turn 1
response1 = chain_with_history.invoke(
    {"question": "My name is Dhruv."},
    config={"configurable": {"session_id": "user_123"}}
)
print(response1.content)

# Turn 2 — model should remember the name
response2 = chain_with_history.invoke(
    {"question": "What is my name?"},
    config={"configurable": {"session_id": "user_123"}}
)
print(response2.content)  # "Your name is Dhruv."
```

---

## 4. Structured Output / Function Calling

**Q5: How do you force an LLM to return structured JSON output using Pydantic?**

```python
from openai import OpenAI
from pydantic import BaseModel
from typing import List

client = OpenAI()

class JobInfo(BaseModel):
    company: str
    role: str
    location: str
    required_skills: List[str]
    experience_years: int

def extract_job_info(job_description: str) -> JobInfo:
    completion = client.beta.chat.completions.parse(
        model="gpt-4o-2024-08-06",
        messages=[
            {"role": "system", "content": "Extract job information from the description."},
            {"role": "user", "content": job_description}
        ],
        response_format=JobInfo
    )
    return completion.choices[0].message.parsed

# Usage
jd = """
TCS is hiring a Senior Python Developer in Bangalore.
5+ years of experience required. Must know Python, FastAPI, Docker, and Kubernetes.
"""
job = extract_job_info(jd)
print(job.model_dump())
# {'company': 'TCS', 'role': 'Senior Python Developer', 'location': 'Bangalore',
#  'required_skills': ['Python', 'FastAPI', 'Docker', 'Kubernetes'], 'experience_years': 5}
```

---

## 5. Document Summarization

**Q6: Write a Python function to summarize a long document using map-reduce chain.**

```python
from langchain_openai import ChatOpenAI
from langchain.chains.combine_documents import create_stuff_documents_chain
from langchain.chains import MapReduceDocumentsChain, ReduceDocumentsChain, LLMChain
from langchain_core.prompts import PromptTemplate
from langchain.text_splitter import RecursiveCharacterTextSplitter

llm = ChatOpenAI(model="gpt-4o-mini", temperature=0)

# Map: summarize each chunk
map_template = "Summarize the following text in 3 bullet points:\n{docs}"
map_prompt = PromptTemplate.from_template(map_template)
map_chain = LLMChain(llm=llm, prompt=map_prompt)

# Reduce: combine all summaries
reduce_template = """
The following are summaries of different sections of a document.
Combine them into one final coherent summary (max 200 words):

{docs}
"""
reduce_prompt = PromptTemplate.from_template(reduce_template)
reduce_chain = LLMChain(llm=llm, prompt=reduce_prompt)
combine_chain = create_stuff_documents_chain(llm, reduce_prompt)

# Build MapReduce chain
map_reduce_chain = MapReduceDocumentsChain(
    llm_chain=map_chain,
    reduce_documents_chain=ReduceDocumentsChain(combine_documents_chain=combine_chain),
    document_variable_name="docs"
)

def summarize_document(text: str) -> str:
    splitter = RecursiveCharacterTextSplitter(chunk_size=1000, chunk_overlap=100)
    docs = splitter.create_documents([text])
    return map_reduce_chain.run(docs)
```

---

## 6. Embedding & Similarity Search

**Q7: How do you compute cosine similarity between two text embeddings?**

```python
import numpy as np
from openai import OpenAI

client = OpenAI()

def get_embedding(text: str, model="text-embedding-3-small") -> list[float]:
    response = client.embeddings.create(input=text, model=model)
    return response.data[0].embedding

def cosine_similarity(vec1: list, vec2: list) -> float:
    a = np.array(vec1)
    b = np.array(vec2)
    return np.dot(a, b) / (np.linalg.norm(a) * np.linalg.norm(b))

# Example
q1 = "How to apply for leave?"
q2 = "What is the process to request time off?"
q3 = "How to compile a Java program?"

e1 = get_embedding(q1)
e2 = get_embedding(q2)
e3 = get_embedding(q3)

print(cosine_similarity(e1, e2))  # ~0.93 (semantically similar)
print(cosine_similarity(e1, e3))  # ~0.67 (semantically different)
```

---

## 7. Simple GenAI Agent with Tools

**Q8: Build a simple tool-calling agent using the OpenAI function calling feature.**

```python
import json
from openai import OpenAI

client = OpenAI()

# Define tools (as JSON schema)
tools = [
    {
        "type": "function",
        "function": {
            "name": "get_weather",
            "description": "Get current weather for a city",
            "parameters": {
                "type": "object",
                "properties": {
                    "city": {"type": "string", "description": "The city name"},
                    "unit": {"type": "string", "enum": ["celsius", "fahrenheit"]}
                },
                "required": ["city"]
            }
        }
    }
]

# Mock tool execution
def get_weather(city: str, unit: str = "celsius") -> str:
    return f"The weather in {city} is 28°{unit[0].upper()} and sunny."

def run_agent(user_query: str) -> str:
    messages = [{"role": "user", "content": user_query}]
    
    response = client.chat.completions.create(
        model="gpt-4o",
        messages=messages,
        tools=tools,
        tool_choice="auto"
    )
    
    message = response.choices[0].message
    
    # If LLM called a tool
    if message.tool_calls:
        tool_call = message.tool_calls[0]
        args = json.loads(tool_call.function.arguments)
        
        # Execute the tool
        result = get_weather(**args)
        
        # Feed result back to LLM
        messages.append(message)
        messages.append({
            "role": "tool",
            "tool_call_id": tool_call.id,
            "content": result
        })
        
        final_response = client.chat.completions.create(
            model="gpt-4o",
            messages=messages
        )
        return final_response.choices[0].message.content
    
    return message.content

print(run_agent("What's the weather like in Mumbai?"))
```
