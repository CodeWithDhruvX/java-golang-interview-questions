# Client Integrations – WhatsApp, Teams & Slack Agents (Service-Based Companies)

One of the most common deliverables at service-based companies is connecting an AI agent to a **messaging platform** a client already uses. This is heavily asked in Accenture, Infosys, and Wipro interviews because it's what they actually build for clients daily.

---

## 1. WhatsApp Agent Integration (Most Requested by Indian Clients)

**Q: A client wants to deploy their customer support AI agent on WhatsApp. How would you integrate it?**

**Answer:**

WhatsApp Business API (via **Meta/360Dialog/Twilio**) is the standard channel for Indian enterprise clients (banks, e-commerce, telecom). The agent logic is the same as any LangChain agent — only the I/O layer changes.

### Architecture

```
Customer (WhatsApp)
        │ sends message
        ▼
Meta WhatsApp Business API
        │ webhook POST
        ▼
┌────────────────────┐
│  Your Webhook       │  FastAPI / Flask endpoint
│  Server             │  (Hosted on Azure/AWS)
└────────┬───────────┘
         │ processes message
         ▼
┌────────────────────┐
│   LangChain Agent  │  (with tools: order lookup, FAQ RAG, etc.)
└────────┬───────────┘
         │ response text
         ▼
Meta WhatsApp API  → sends reply to customer
```

### Full Implementation

```python
# pip install fastapi uvicorn langchain-openai twilio

from fastapi import FastAPI, Request, HTTPException
from langchain_openai import ChatOpenAI
from langchain.agents import AgentExecutor, create_openai_tools_agent
from langchain_core.prompts import ChatPromptTemplate, MessagesPlaceholder
from langchain_core.tools import tool
from langchain.memory import ConversationSummaryBufferMemory
import httpx
import hashlib
import hmac
import json

app = FastAPI(title="WhatsApp Agent Webhook")
llm = ChatOpenAI(model="gpt-4o", temperature=0)

# --- Memory store (use Redis in production, dict for demo) ---
sessions: dict[str, ConversationSummaryBufferMemory] = {}

def get_or_create_memory(phone_number: str) -> ConversationSummaryBufferMemory:
    """Each WhatsApp number gets its own memory session."""
    if phone_number not in sessions:
        sessions[phone_number] = ConversationSummaryBufferMemory(
            llm=llm,
            max_token_limit=2000,    # Summarize when history gets long
            return_messages=True,
            memory_key="chat_history"
        )
    return sessions[phone_number]

# --- Tools ---
@tool
def get_order_status(order_id: str) -> str:
    """Get order status by order ID (e.g., 'ORD-12345')."""
    # Call your actual order API
    return json.dumps({"order_id": order_id, "status": "Shipped", "eta": "2026-03-10"})

@tool
def get_store_info(query: str) -> str:
    """Answer store-related FAQs: hours, location, policies."""
    faqs = {
        "hours": "We are open Monday–Saturday, 9 AM to 9 PM IST.",
        "return": "Returns accepted within 30 days with original packaging.",
        "location": "Nearest store: Phoenix Mall, Bangalore. Use maps.google.com/stylekart"
    }
    for key, answer in faqs.items():
        if key in query.lower():
            return answer
    return "Sorry, I don't have specific information on that. Please call 1800-XXX-XXXX."

tools = [get_order_status, get_store_info]
prompt = ChatPromptTemplate.from_messages([
    ("system", """You are a WhatsApp customer support assistant for StyleKart.
    Be concise — WhatsApp users prefer short, clear messages.
    Use emojis sparingly for friendliness. 
    Never send messages longer than 3 short paragraphs.
    If you can't help, escalate: 'Let me connect you to our team. One moment...'"""),
    MessagesPlaceholder(variable_name="chat_history"),
    ("human", "{input}"),
    MessagesPlaceholder(variable_name="agent_scratchpad"),
])
agent = create_openai_tools_agent(llm, tools, prompt)

# --- WhatsApp Webhook Verification ---
WHATSAPP_VERIFY_TOKEN = "your-custom-verify-token"
WHATSAPP_APP_SECRET = "your-meta-app-secret"

@app.get("/webhook")
async def verify_webhook(request: Request):
    """Meta sends a GET to verify your webhook during setup."""
    params = dict(request.query_params)
    if params.get("hub.verify_token") == WHATSAPP_VERIFY_TOKEN:
        return int(params.get("hub.challenge", 0))
    raise HTTPException(status_code=403, detail="Invalid verify token")

# --- Message Handler ---
@app.post("/webhook")
async def handle_whatsapp_message(request: Request):
    """Receives incoming WhatsApp messages from Meta."""
    body = await request.body()

    # 1. Verify signature (security: ensures request is from Meta)
    signature = request.headers.get("X-Hub-Signature-256", "")
    expected_sig = "sha256=" + hmac.new(
        WHATSAPP_APP_SECRET.encode(), body, hashlib.sha256
    ).hexdigest()
    if not hmac.compare_digest(signature, expected_sig):
        raise HTTPException(status_code=403, detail="Invalid signature")

    data = json.loads(body)

    # 2. Extract message from Meta's webhook payload
    try:
        entry = data["entry"][0]["changes"][0]["value"]
        messages_list = entry.get("messages", [])
        if not messages_list:
            return {"status": "no_message"}  # Delivery receipts etc.

        message = messages_list[0]
        sender_phone = message["from"]           # Phone number like "919876543210"
        message_text = message["text"]["body"]   # The actual message text
        message_id = message["id"]

    except (KeyError, IndexError):
        return {"status": "ignored"}

    # 3. Run the agent with per-user memory
    memory = get_or_create_memory(sender_phone)
    executor = AgentExecutor(
        agent=agent, tools=tools, memory=memory,
        max_iterations=4, handle_parsing_errors=True
    )

    try:
        result = executor.invoke({"input": message_text})
        response_text = result["output"]
    except Exception as e:
        response_text = "I'm having trouble right now. Please try again or call 1800-XXX-XXXX."

    # 4. Send reply via WhatsApp API
    await send_whatsapp_message(sender_phone, response_text)
    return {"status": "processed"}

async def send_whatsapp_message(to_phone: str, message: str):
    """Sends a reply back via Meta WhatsApp Cloud API."""
    PHONE_NUMBER_ID = "your-whatsapp-phone-number-id"
    ACCESS_TOKEN = "your-meta-permanent-access-token"

    async with httpx.AsyncClient() as client:
        await client.post(
            f"https://graph.facebook.com/v17.0/{PHONE_NUMBER_ID}/messages",
            headers={"Authorization": f"Bearer {ACCESS_TOKEN}"},
            json={
                "messaging_product": "whatsapp",
                "to": to_phone,
                "type": "text",
                "text": {"body": message}
            }
        )
```

---

## 2. Microsoft Teams Bot (Common in IT Services/Enterprise)

**Q: A large enterprise client uses Microsoft Teams. How do you deploy your agent as a Teams bot?**

**Answer:**

```python
# pip install botframework-integration-aiohttp botbuilder-core

from botbuilder.core import ActivityHandler, TurnContext, MessageFactory
from botbuilder.schema import Activity
from langchain_openai import ChatOpenAI
from langchain.agents import AgentExecutor, create_openai_tools_agent
from langchain.memory import ConversationBufferWindowMemory
from langchain_core.tools import tool
from langchain_core.prompts import ChatPromptTemplate, MessagesPlaceholder
import json

llm = ChatOpenAI(model="gpt-4o", temperature=0)

# --- Tools tailored for enterprise context ---
@tool
def search_it_documentation(query: str) -> str:
    """Search internal IT help desk documentation and policies."""
    # In production: query your SharePoint/Confluence knowledge base
    return f"IT Policy result for '{query}': Submit tickets via ServiceNow at https://helpdesk.company.com"

@tool
def check_system_status(system_name: str) -> str:
    """Check if a company system is up or down."""
    statuses = {
        "sap": "✅ Operational",
        "email": "✅ Operational",
        "vpn": "⚠️ Degraded performance in Chennai region",
        "teams": "✅ Operational"
    }
    return statuses.get(system_name.lower(), f"Status for {system_name}: Unknown")

tools = [search_it_documentation, check_system_status]
user_memories: dict[str, ConversationBufferWindowMemory] = {}

class ITSupportBot(ActivityHandler):
    """Teams bot class — extends ActivityHandler from Bot Framework SDK."""

    async def on_message_activity(self, turn_context: TurnContext):
        user_id = turn_context.activity.from_property.id
        user_message = turn_context.activity.text

        # Per-user memory
        if user_id not in user_memories:
            user_memories[user_id] = ConversationBufferWindowMemory(
                k=5,  # Keep last 5 exchanges
                return_messages=True,
                memory_key="chat_history"
            )
        
        memory = user_memories[user_id]
        prompt = ChatPromptTemplate.from_messages([
            ("system", "You are an IT support assistant on Microsoft Teams. Be professional and concise."),
            MessagesPlaceholder(variable_name="chat_history"),
            ("human", "{input}"),
            MessagesPlaceholder(variable_name="agent_scratchpad"),
        ])
        
        agent = create_openai_tools_agent(llm, tools, prompt)
        executor = AgentExecutor(agent=agent, tools=tools, memory=memory, max_iterations=4)
        
        # Show "typing" indicator while agent thinks
        await turn_context.send_activity(Activity(type="typing"))
        
        try:
            result = executor.invoke({"input": user_message})
            response = result["output"]
        except Exception:
            response = "I'm having trouble right now. Please contact IT support directly."
        
        await turn_context.send_activity(MessageFactory.text(response))

    async def on_members_added_activity(self, members_added, turn_context: TurnContext):
        """Welcome message when bot is added to a chat."""
        for member in members_added:
            if member.id != turn_context.activity.recipient.id:
                await turn_context.send_activity(
                    "👋 Hi! I'm the IT Support Assistant. Ask me about:\n"
                    "• System status\n• IT policies\n• How to raise tickets"
                )
```

### Registering the Teams Bot (Azure Bot Service)

```bash
# 1. Register your bot in Azure Bot Service
az bot create \
  --resource-group myResourceGroup \
  --name "it-support-bot" \
  --kind webapp \
  --location eastus

# 2. Configure Teams channel
az bot msteams create \
  --resource-group myResourceGroup \
  --name "it-support-bot"

# 3. Deploy your FastAPI + Bot Framework app to Azure App Service
# 4. Upload Teams app manifest (app.zip) to Teams Admin Centre
```

---

## 3. Slack Agent Integration

**Q: How do you build an agent bot that lives inside a customer's Slack workspace?**

**Answer:**

```python
# pip install slack-bolt langchain-openai

from slack_bolt import App
from slack_bolt.adapter.fastapi import SlackRequestHandler
from fastapi import FastAPI, Request
from langchain_openai import ChatOpenAI
from langchain.agents import AgentExecutor, create_openai_tools_agent
from langchain_core.tools import tool
from langchain_core.prompts import ChatPromptTemplate, MessagesPlaceholder
import json

# Initialize Slack app (get tokens from api.slack.com)
slack_app = App(
    token="xoxb-your-bot-token",        # Bot User OAuth Token
    signing_secret="your-signing-secret" # Used to verify Slack requests
)

llm = ChatOpenAI(model="gpt-4o-mini", temperature=0)

@tool
def get_deploy_status(service_name: str) -> str:
    """Check the deployment status of a service (DevOps use case)."""
    statuses = {
        "api-gateway": "✅ v2.3.1 deployed 2h ago",
        "user-service": "⚠️ Rollback in progress",
        "payment-service": "🔴 Deploy failed - check Jenkins"
    }
    return statuses.get(service_name, f"No deploy info found for {service_name}")

@tool
def create_jira_ticket(summary: str, priority: str = "Medium") -> str:
    """Create a Jira ticket for engineering issues."""
    # In production: call Jira REST API
    ticket_id = "ENG-" + str(hash(summary))[-4:]
    return f"Created Jira ticket {ticket_id}: '{summary}' (Priority: {priority})"

tools = [get_deploy_status, create_jira_ticket]

def get_agent_response(user_message: str, channel_history: list) -> str:
    """Run the LangChain agent and return the response."""
    prompt = ChatPromptTemplate.from_messages([
        ("system", """You are a DevOps assistant bot in Slack.
        Be concise. Use Slack markdown for formatting (*bold*, `code`, ```code blocks```).
        Always ask before creating Jira tickets — confirm with the user first."""),
        ("human", "{input}"),
        MessagesPlaceholder(variable_name="agent_scratchpad"),
    ])
    agent = create_openai_tools_agent(llm, tools, prompt)
    executor = AgentExecutor(agent=agent, tools=tools, max_iterations=4)
    result = executor.invoke({"input": user_message})
    return result["output"]

# --- Slack Event Handlers ---

# Handle @mentions of the bot in any channel
@slack_app.event("app_mention")
def handle_mention(event, say, client):
    """Called when someone @mentions the bot."""
    # Remove the bot mention from the text
    user_message = event["text"].split(">", 1)[-1].strip()
    user_id = event["user"]
    channel = event["channel"]
    thread_ts = event.get("thread_ts", event["ts"])  # Reply in thread

    # Show typing indicator
    client.reactions_add(channel=channel, timestamp=event["ts"], name="thinking_face")

    response = get_agent_response(user_message, [])
    
    # Remove thinking emoji and reply in thread
    client.reactions_remove(channel=channel, timestamp=event["ts"], name="thinking_face")
    say(text=response, thread_ts=thread_ts)

# Handle direct messages to the bot
@slack_app.event("message")
def handle_direct_message(event, say):
    """Called for DMs — no @mention needed."""
    if event.get("channel_type") == "im":  # im = direct message
        user_message = event.get("text", "")
        response = get_agent_response(user_message, [])
        say(response)

# --- FastAPI wrapper for Slack webhook ---
api = FastAPI()
handler = SlackRequestHandler(slack_app)

@api.post("/slack/events")
async def endpoint(req: Request):
    return await handler.handle(req)
```

---

## 4. Comparison for Client Conversations

| Platform | When to Recommend | Setup Complexity | Client Org Type |
|---|---|---|---|
| **WhatsApp** | Customer-facing B2C, India market | Medium (Meta approval needed) | Retail, BFSI, Telecom |
| **Microsoft Teams** | Internal enterprise use, M365 clients | Medium-Low (Azure Bot Service) | Large enterprises, BFSI |
| **Slack** | Tech startups, DevOps tooling | Low (well-documented API) | IT services, startups |
| **Web Chat Widget** | Embed on website (Freshchat, Intercom) | Lowest | Any |

**Interview Talking Point:**
> *"For client integrations, I always check two things first: where does the end user already spend their time, and what's the client's data residency requirement. For Indian retail clients, WhatsApp is non-negotiable — 90% of their customers are on WhatsApp. For a large BFSI client's internal helpdesk, Microsoft Teams is the obvious choice since they're already on M365. The agent logic is identical in both cases — only the webhook layer changes."*
