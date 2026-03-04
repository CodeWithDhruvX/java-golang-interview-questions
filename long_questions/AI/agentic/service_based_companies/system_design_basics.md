# System Design Basics for Agentic AI Workflows (Service-Based Companies)

System design for agentic AI in service-based company interviews is more practical than theoretical. Focus on understanding how to structure agent workflows for real client use cases, avoiding over-engineering.

---

## 1. Design a Customer Support Agent for an E-Commerce Client

**Question:** A client wants to replace their call center with an AI agent that can handle order status, returns, FAQ, and escalation. How would you design this?

**Answer:**

### Architecture Diagram

```
Customer (Web/WhatsApp/App)
         │
         ▼
┌─────────────────────┐
│   API Gateway        │  (FastAPI / Node.js)
└──────────┬──────────┘
           ▼
┌─────────────────────┐
│   Intent Router      │  LLM classifies user query type
│   (GPT-3.5-turbo)   │  → ORDER / FAQ / RETURN / ESCALATE
└──────────┬──────────┘
           │
   ┌───────┼───────┬────────┐
   ▼       ▼       ▼        ▼
ORDER    FAQ     RETURN  ESCALATE
Agent   Agent   Agent    Queue
  │       │       │
  └───────┴───────┘
          │
     Order DB / 
     Policy Docs /
     Returns System
```

### Agent Design (Step-by-Step)

**Step 1: Intent Classification**
```python
INTENT_PROMPT = """Classify this customer query into one of:
ORDER_STATUS, RETURN_REQUEST, FAQ, COMPLAINT, ESCALATE

Query: {query}
Output only the category."""
```

**Step 2: Order Status Agent**
```python
@tool
def get_order_status(order_id: str) -> str:
    """Fetches order status from the order management system."""
    # Calls internal REST API
    response = requests.get(f"https://orders-api/orders/{order_id}")
    return response.json()

@tool
def get_estimated_delivery(order_id: str) -> str:
    """Gets the estimated delivery date for an order."""
    ...
```

**Step 3: FAQ Agent (RAG-based)**
```
- Load FAQ PDFs into ChromaDB vector store
- Agent retrieves top 3 relevant FAQ chunks
- Answers grounded in actual policy (reduces hallucinations)
```

**Step 4: Escalation (HITL)**
```
If the agent:
- Cannot resolve after 3 attempts
- Detects angry or threatening language
- Gets a complaint about wrong amount/fraud

→ Transfer to human agent queue in CRM (Freshdesk/Zendesk)
→ Share full conversation context so agent doesn't start over
```

### Key Points for Interview

| Concern | Solution |
|---|---|
| Wrong order info | Only call verified APIs — never hallucinate order data |
| Customer privacy | Never log PII (name, phone) in LLM prompts |
| Downtime | Fallback: "I'm unable to process this right now. Please call [number]." |
| Cost | Cache FAQ answers (Redis, 1 hour TTL) |
| Language | Translate to English → Agent → Translate back |

---

## 2. Design an Internal HR Chatbot for a 10,000-Employee Company

**Question:** An IT services company wants to deploy an HR chatbot that can answer policy questions, manage leave requests, and show payslips. Design the architecture.

**Answer:**

### Architecture

```
Employee (Teams / Slack / Intranet)
         │
         ▼
┌─────────────────────────────────┐
│         HR Agent                │
│         (LangChain AgentExecutor)│
└─────────────────────────────────┘
         │
   ┌─────┼──────┬─────────────┐
   ▼     ▼      ▼             ▼
Policy  Leave  Payslip    Escalate to
FAQ     Tool   Tool       HR Executive
(RAG)    │      │
         ▼      ▼
      HRMS API  Payroll DB
      (SAP/     (Encrypted)
       Rippling)
```

### Tools the Agent Uses

```python
@tool
def get_leave_balance(employee_id: str) -> str:
    """Check remaining leave balance for an employee."""
    return hrms_api.get_leaves(employee_id)

@tool
def apply_for_leave(employee_id: str, from_date: str, to_date: str, reason: str) -> str:
    """Apply for leave on behalf of the employee. Requires approval."""
    application = hrms_api.apply_leave(employee_id, from_date, to_date, reason)
    return f"Leave applied. Application ID: {application.id}. Pending manager approval."

@tool
def search_hr_policy(query: str) -> str:
    """Search HR policies and answer HR-related questions."""
    # This is a RAG call into the HR policy vector store
    results = policy_vectorstore.similarity_search(query, k=3)
    return "\n".join([r.page_content for r in results])

@tool
def get_payslip(employee_id: str, month: str) -> str:
    """Retrieve payslip for a given month."""
    # Always verify employee_id matches the logged-in user
    if not verify_ownership(employee_id, current_user):
        return "Access denied. You can only view your own payslips."
    return payroll_db.get_payslip(employee_id, month)
```

### Security Considerations

```
1. Authentication: Agent only runs after SSO login (OAuth 2.0)
2. Authorization: Each tool verifies the user only accesses their own data
3. No storage: Conversation not stored (GDPR compliance)
4. Audit log: Every tool call is logged with employee_id + timestamp
5. Sensitive data: Salary info is masked in logs: "₹**,***"
```

### Conversation Flow Example

```
Employee: "How many leaves do I have left?"
Agent:    [calls get_leave_balance(emp_id="EMP1234")]
          "You have 8 casual leaves and 12 earned leaves remaining."

Employee: "Apply for 2 days sick leave from March 10-11"
Agent:    [calls apply_for_leave(...)]
          "Done! Left application #LA-4421 submitted. Your manager will 
           receive an approval request. Expected response: 24 hours."
```

---

## 3. Design an Email Processing Agent

**Question:** A client receives 500 emails daily. Design an agent that reads emails, classifies them, drafts replies, and routes urgent ones to the right team.

**Answer:**

### Workflow

```
Gmail / Outlook (IMAP)
       │
       ▼
┌──────────────────────┐
│ Email Fetcher        │  Runs every 15 minutes (cron job)
└──────────┬───────────┘
           ▼
┌──────────────────────┐
│ Classifier Agent     │  GPT-4o-mini (fast, cheap)
│ Categories:          │
│  - Support Request   │
│  - Sales Lead        │
│  - Invoice           │
│  - Spam              │
│  - Urgent            │
└──────────┬───────────┘
           ▼
┌──────────────────────┐   ┌──────────────────┐
│ Reply Drafting Agent │   │ Human Approval   │
│ (for non-urgent)     │──▶│ Queue (Notion/   │
│                      │   │  CRM)            │
└──────────────────────┘   └──────────────────┘
           │
           ▼ (Human approves)
  Send via Gmail API
```

### Agent Code Structure

```python
from langchain_openai import ChatOpenAI
from langchain_core.tools import tool
import imaplib
import email

llm = ChatOpenAI(model="gpt-4o-mini", temperature=0)
draft_llm = ChatOpenAI(model="gpt-4o", temperature=0.4)  # Better quality for drafts

@tool
def classify_email(subject: str, body: str) -> str:
    """Classifies an email into a category."""
    response = llm.invoke(f"""Classify this email:
    Subject: {subject}
    Body: {body[:500]}
    
    Categories: SUPPORT, SALES_LEAD, INVOICE, SPAM, URGENT
    Respond with only the category.""")
    return response.content

@tool
def draft_reply(original_email: str, category: str, company_context: str) -> str:
    """Drafts a professional reply to an email."""
    response = draft_llm.invoke(f"""You are a professional email writer for {company_context}.
    
    Write a concise, professional reply to this {category} email:
    {original_email}
    
    Keep it under 150 words. Be helpful and specific.""")
    return response.content

def process_emails():
    emails = fetch_unread_emails()  # IMAP fetch
    
    for email_data in emails:
        category = classify_email(email_data["subject"], email_data["body"])
        
        if category == "SPAM":
            archive_email(email_data["id"])
            continue
        
        if category == "URGENT":
            # Immediate Slack/Teams notification
            alert_team(email_data, channel="#urgent-emails")
        
        # Draft reply for human review (never send automatically)
        draft = draft_reply(
            original_email=f"Subject: {email_data['subject']}\n{email_data['body']}",
            category=category,
            company_context="Accenture India Support Team"
        )
        save_draft_for_review(email_data["id"], draft, category)
```

**❗ Important Design Decision:** Never send emails automatically. Always involve a human in the loop for outbound communication. The agent drafts, humans approve.

---

## 4. Design a Document Summarization Agent Pipeline

**Question:** A legal firm wants an agent that processes 100-page legal documents and produces a 1-page summary with key clauses. Design this.

**Answer:**

### Challenge: 100-page document doesn't fit in LLM context (~128K tokens = ~100 pages, but costly)

### Solution: Map-Reduce Summarization

```
100-page PDF
     │
     ▼
[Text Extraction]  → PyMuPDF / pdfplumber
     │
     ▼
[Chunking]  → 20 chunks of 5 pages each
     │
     │ (Parallel)
  ┌──┼──┐... 
  ▼  ▼  ▼
[MAP Phase]  Each chunk → GPT-3.5 summary (20 calls, parallel, fast)
  │  │  │
  └──┼──┘
     │
     ▼
[REDUCE Phase]  20 summaries → GPT-4o → Final 1-page summary
     │
     ▼
[Key Clause Extraction]  GPT-4o extracts: party names, dates, penalties
     │
     ▼
Final Summary + Clause Table (Word/PDF output)
```

### Code Skeleton

```python
from langchain.chains.summarize import load_summarize_chain
from langchain.text_splitter import RecursiveCharacterTextSplitter
from langchain_community.document_loaders import PyMuPDFLoader
from langchain_openai import ChatOpenAI
from concurrent.futures import ThreadPoolExecutor

cheap_llm = ChatOpenAI(model="gpt-3.5-turbo", temperature=0)   # For map phase
powerful_llm = ChatOpenAI(model="gpt-4o", temperature=0)         # For reduce phase

def summarize_legal_document(pdf_path: str) -> dict:
    # Load and chunk
    loader = PyMuPDFLoader(pdf_path)
    docs = loader.load()
    
    splitter = RecursiveCharacterTextSplitter(chunk_size=8000, chunk_overlap=500)
    chunks = splitter.split_documents(docs)
    
    # Map phase: summarize each chunk (run in parallel for speed)
    def summarize_chunk(chunk):
        response = cheap_llm.invoke(
            f"Summarize the key points in this legal document section:\n{chunk.page_content}"
        )
        return response.content
    
    with ThreadPoolExecutor(max_workers=5) as executor:
        chunk_summaries = list(executor.map(summarize_chunk, chunks))
    
    # Reduce phase: combine all summaries
    combined = "\n\n---\n\n".join(chunk_summaries)
    final_summary = powerful_llm.invoke(
        f"Create a clear 1-page summary from these section summaries:\n{combined}"
    )
    
    # Extract key clauses
    clauses = powerful_llm.invoke(
        f"From this document summary, extract: parties involved, dates, "
        f"payment terms, penalties, termination conditions.\n{final_summary.content}"
    )
    
    return {
        "summary": final_summary.content,
        "key_clauses": clauses.content
    }
```

### Cost & Time Estimate

| Phase | Model | ~Cost (100-page doc) | ~Time |
|---|---|---|---|
| Map (20 chunks) | GPT-3.5-turbo | $0.04 | 30s (parallel) |
| Reduce (1 call) | GPT-4o | $0.15 | 10s |
| Clause extraction | GPT-4o | $0.10 | 8s |
| **Total** | | **$0.29** | **~50s** |

---

## 5. How Do You Deploy a LangChain Agent as a REST API?

**Question:** You've built an agent in LangChain. How do you expose it as an API so the frontend can call it?

**Answer:**

```python
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from langchain_openai import ChatOpenAI
from langchain.agents import AgentExecutor, create_openai_tools_agent
from langchain_core.prompts import ChatPromptTemplate, MessagesPlaceholder
from langchain_core.tools import tool
import uuid

app = FastAPI(title="Support Agent API")

# ---- Agent Setup ----
llm = ChatOpenAI(model="gpt-3.5-turbo", temperature=0)

@tool
def answer_faq(question: str) -> str:
    """Answers frequently asked product questions."""
    faqs = {
        "refund policy": "We offer a 30-day money-back guarantee.",
        "shipping time": "Standard shipping takes 5-7 business days.",
    }
    for key, answer in faqs.items():
        if key in question.lower():
            return answer
    return "I don't have a specific answer in our knowledge base."

tools = [answer_faq]
prompt = ChatPromptTemplate.from_messages([
    ("system", "You are a helpful customer support agent."),
    ("human", "{input}"),
    MessagesPlaceholder(variable_name="agent_scratchpad"),
])

agent = create_openai_tools_agent(llm, tools, prompt)
agent_executor = AgentExecutor(agent=agent, tools=tools, max_iterations=5)

# ---- FastAPI Endpoints ----
class QueryRequest(BaseModel):
    user_id: str
    message: str
    session_id: str = None  # Optional: for multi-turn conversation

class QueryResponse(BaseModel):
    session_id: str
    response: str
    status: str

@app.post("/chat", response_model=QueryResponse)
async def chat(request: QueryRequest):
    """Process a user message and return agent response."""
    session_id = request.session_id or str(uuid.uuid4())
    
    try:
        result = agent_executor.invoke({"input": request.message})
        return QueryResponse(
            session_id=session_id,
            response=result["output"],
            status="success"
        )
    except Exception as e:
        raise HTTPException(status_code=500, detail=str(e))

@app.get("/health")
async def health_check():
    return {"status": "healthy"}

# Run with: uvicorn main:app --host 0.0.0.0 --port 8000
```

### Docker Deployment

```dockerfile
FROM python:3.11-slim
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY . .
EXPOSE 8000
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000", "--workers", "4"]
```

**Key Concepts:**
- FastAPI is the go-to framework for deploying LangChain agents (async support)
- `--workers 4` allows 4 concurrent requests
- Health endpoint is essential for K8s liveness probes
- Always validate and sanitize `user_id` from requests
