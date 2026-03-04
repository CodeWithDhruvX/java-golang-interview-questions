# Cloud AI Agent Platforms – Interview Q&A (Service-Based Companies)

Service-based companies (Accenture, Infosys, TCS, Wipro, Capgemini) deploy agents on **client-approved cloud stacks** — most often Azure, AWS, or GCP. Interviewers will ask: "How would you build this on Azure?" rather than "Design the algorithm." This file covers those platform-specific questions.

---

## 1. Azure AI Foundry — Building Agents on Microsoft's Platform

**Q: A client's entire stack is on Azure. How would you build an Agentic AI solution using Azure services?**

**Answer:**

**Azure AI Foundry** (formerly Azure AI Studio) is Microsoft's unified platform for building, deploying, and monitoring AI agents in enterprise Azure environments.

### Core Azure Services for Agentic AI

| Service | Role in Agent Stack | Equivalent |
|---|---|---|
| **Azure OpenAI Service** | LLM backend (GPT-4o, GPT-4o-mini) | OpenAI API |
| **Azure AI Search** | Vector store + hybrid search for RAG | Pinecone / Weaviate |
| **Azure AI Foundry Agents** | Managed agent runtime with built-in tools | LangGraph / AutoGen |
| **Azure Functions** | Custom tool execution (serverless) | FastAPI tool endpoints |
| **Azure Cosmos DB** | Agent state & conversation history | Redis + Postgres |
| **Azure Service Bus** | Message queue for event-driven agents | Kafka |
| **Azure Key Vault** | Secure storage for API keys/secrets | AWS Secrets Manager |
| **Azure Monitor / App Insights** | Observability and alerting | Prometheus + Grafana |

### Building an Agent with Azure AI Foundry SDK

```python
# pip install azure-ai-projects azure-ai-inference azure-identity

from azure.ai.projects import AIProjectClient
from azure.ai.projects.models import (
    AgentsApiToolChoiceOptionMode,
    FunctionTool,
    ToolSet,
)
from azure.identity import DefaultAzureCredential
import json

# Connect to your Azure AI Foundry project
project_client = AIProjectClient.from_connection_string(
    credential=DefaultAzureCredential(),
    conn_str="eastus.api.azureml.ms;your-subscription-id;your-rg;your-project"
)

# Define custom tools (same JSON schema as OpenAI function calling)
def get_order_status(order_id: str) -> str:
    """Fetches order status from internal ERP system."""
    # In production: call your ERP API (SAP, Oracle, etc.)
    return json.dumps({"order_id": order_id, "status": "Shipped", "eta": "2026-03-10"})

def calculate_refund(order_id: str, reason: str) -> str:
    """Calculates eligible refund amount based on policy."""
    return json.dumps({"order_id": order_id, "refund_amount": 2500, "currency": "INR"})

# Register functions as tools
functions = FunctionTool(functions=[get_order_status, calculate_refund])
toolset = ToolSet()
toolset.add(functions)

# Create the agent on Azure
agent = project_client.agents.create_agent(
    model="gpt-4o",                         # Deployed in your Azure OpenAI resource
    name="customer-support-agent",
    instructions="""You are a customer support agent for StyleKart. 
    Use the provided tools to help customers with order status and refunds.
    Always verify order details before processing any refund.
    For refunds above ₹5,000, escalate to a human agent.""",
    toolset=toolset,
    headers={"x-ms-enable-preview": "true"}
)
print(f"Created agent: {agent.id}")

# Create a thread (conversation session)
thread = project_client.agents.create_thread()

# Add user message
project_client.agents.create_message(
    thread_id=thread.id,
    role="user",
    content="I want to return my order ORD-88231. The product was damaged."
)

# Run the agent
run = project_client.agents.create_and_process_run(
    thread_id=thread.id,
    assistant_id=agent.id
)

# Get the response
messages = project_client.agents.list_messages(thread_id=thread.id)
for message in messages.data:
    if message.role == "assistant":
        print(f"Agent: {message.content[0].text.value}")
```

### Azure AI Search for RAG (Enterprise Document Search)

```python
from azure.search.documents import SearchClient
from azure.search.documents.models import VectorizedQuery
from azure.core.credentials import AzureKeyCredential
from openai import AzureOpenAI

# Azure OpenAI for embeddings
aoai_client = AzureOpenAI(
    azure_endpoint="https://your-resource.openai.azure.com/",
    api_key="your-azure-openai-key",
    api_version="2024-02-01"
)

# Azure AI Search client
search_client = SearchClient(
    endpoint="https://your-search-service.search.windows.net",
    index_name="company-documents",
    credential=AzureKeyCredential("your-search-key")
)

def rag_search(query: str, top_k: int = 5) -> list[dict]:
    """Hybrid search: combines semantic vector search + keyword search (BM25)."""
    
    # Generate query embedding using Azure OpenAI
    embedding_response = aoai_client.embeddings.create(
        model="text-embedding-ada-002",    # Your Azure OpenAI embedding deployment
        input=query
    )
    query_vector = embedding_response.data[0].embedding
    
    # Hybrid search query (vector + keyword)
    vector_query = VectorizedQuery(
        vector=query_vector,
        k_nearest_neighbors=top_k,
        fields="content_vector"            # The vector field in your search index
    )
    
    results = search_client.search(
        search_text=query,                 # BM25 keyword search
        vector_queries=[vector_query],     # Semantic vector search
        select=["title", "content", "source_url", "last_updated"],
        top=top_k
    )
    
    return [
        {
            "title": r["title"],
            "content": r["content"],
            "source": r["source_url"],
            "score": r["@search.score"]
        }
        for r in results
    ]

# Usage in an agent tool
def answer_from_company_docs(question: str) -> str:
    docs = rag_search(question)
    context = "\n\n".join([f"[{d['title']}]: {d['content']}" for d in docs])
    
    response = aoai_client.chat.completions.create(
        model="gpt-4o",
        messages=[
            {"role": "system", "content": "Answer based on the company documents. Cite sources."},
            {"role": "user", "content": f"Context:\n{context}\n\nQuestion: {question}"}
        ]
    )
    return response.choices[0].message.content
```

---

## 2. AWS Bedrock Agents — Building on Amazon's Platform

**Q: How do you build an AI agent using AWS Bedrock for a client on AWS?**

**Answer:**

**Amazon Bedrock Agents** is AWS's fully managed service for building agents with access to 30+ foundation models (Claude, Llama, Titan, Mistral).

### Key AWS Services for Agents

| Service | Role |
|---|---|
| **Amazon Bedrock** | LLM access + managed agent runtime |
| **Amazon Kendra / OpenSearch** | Enterprise RAG knowledge base |
| **AWS Lambda** | Custom tool execution |
| **Amazon DynamoDB** | Agent state and conversation history |
| **Amazon S3** | Document storage for RAG |
| **AWS Secrets Manager** | API keys for tools |
| **Amazon CloudWatch** | Monitoring and alerting |

### Building a Bedrock Agent

```python
import boto3
import json

# Initialize Bedrock clients
bedrock_client = boto3.client('bedrock-agent-runtime', region_name='us-east-1')

# Invoke an existing Bedrock Agent
def invoke_bedrock_agent(agent_id: str, agent_alias_id: str, user_message: str, session_id: str) -> str:
    """
    Calls a Bedrock Agent with a user message.
    session_id maintains conversation history automatically.
    """
    response = bedrock_client.invoke_agent(
        agentId=agent_id,
        agentAliasId=agent_alias_id,
        sessionId=session_id,
        inputText=user_message,
        enableTrace=True           # Capture reasoning steps for observability
    )

    # Bedrock returns a streaming response
    full_response = ""
    trace_events = []

    for event in response['completion']:
        if 'chunk' in event:
            chunk_data = event['chunk']
            if 'bytes' in chunk_data:
                full_response += chunk_data['bytes'].decode('utf-8')
        
        # Capture trace for debugging (shows agent's thought process)
        if 'trace' in event:
            trace_events.append(event['trace'])

    return full_response

# Create an Action Group (equivalent to defining tools in LangChain)
# Action Groups connect your Lambda functions to the Bedrock Agent
def create_action_group_schema() -> dict:
    """
    OpenAPI schema that defines what tools the agent can use.
    Bedrock uses this to understand when and how to call your Lambda.
    """
    return {
        "openapi": "3.0.0",
        "info": {"title": "Customer Support API", "version": "1.0.0"},
        "paths": {
            "/get-order-status": {
                "get": {
                    "summary": "Get order status by order ID",
                    "operationId": "getOrderStatus",
                    "parameters": [
                        {
                            "name": "orderId",
                            "in": "query",
                            "required": True,
                            "schema": {"type": "string"},
                            "description": "The order ID, e.g. 'ORD-12345'"
                        }
                    ],
                    "responses": {
                        "200": {
                            "description": "Order status details",
                            "content": {
                                "application/json": {
                                    "schema": {
                                        "type": "object",
                                        "properties": {
                                            "status": {"type": "string"},
                                            "estimatedDelivery": {"type": "string"}
                                        }
                                    }
                                }
                            }
                        }
                    }
                }
            }
        }
    }

# Lambda function that handles tool calls from Bedrock Agent
def lambda_handler(event, context):
    """
    This Lambda is called by Bedrock when the agent wants to use a tool.
    event contains: actionGroup, apiPath, httpMethod, parameters
    """
    action_group = event['actionGroup']
    api_path = event['apiPath']
    
    if api_path == '/get-order-status':
        order_id = next(
            p['value'] for p in event['parameters'] if p['name'] == 'orderId'
        )
        # Call your actual order system here
        result = {"status": "Shipped", "estimatedDelivery": "2026-03-10"}
        
    elif api_path == '/process-refund':
        # ... handle refund tool
        result = {"refundId": "REF-789", "amount": 2500, "status": "Initiated"}
    
    else:
        result = {"error": f"Unknown API path: {api_path}"}
    
    # Response format required by Bedrock
    return {
        "messageVersion": "1.0",
        "response": {
            "actionGroup": action_group,
            "apiPath": api_path,
            "httpMethod": event['httpMethod'],
            "httpStatusCode": 200,
            "responseBody": {
                "application/json": {
                    "body": json.dumps(result)
                }
            }
        }
    }
```

---

## 3. Google Vertex AI Agents — Building on GCP

**Q: How do you build an AI agent on Google Cloud for a client using GCP?**

**Answer:**

```python
# pip install google-cloud-aiplatform vertexai

import vertexai
from vertexai.preview.generative_models import (
    GenerativeModel,
    Tool,
    FunctionDeclaration,
)

# Initialize Vertex AI
vertexai.init(project="your-gcp-project", location="us-central1")

# Define tools using Vertex AI's function declaration format
get_product_info = FunctionDeclaration(
    name="get_product_info",
    description="Get detailed information about a product by its product code.",
    parameters={
        "type": "object",
        "properties": {
            "product_code": {
                "type": "string",
                "description": "Product code like 'ELEC-001'"
            }
        },
        "required": ["product_code"]
    }
)

check_stock = FunctionDeclaration(
    name="check_stock",
    description="Check current stock availability for a product at a specific warehouse.",
    parameters={
        "type": "object",
        "properties": {
            "product_code": {"type": "string"},
            "warehouse_city": {"type": "string", "enum": ["mumbai", "delhi", "bangalore"]}
        },
        "required": ["product_code"]
    }
)

# Create tool set
retail_tools = Tool(function_declarations=[get_product_info, check_stock])

# Initialize Gemini model with tools
model = GenerativeModel(
    "gemini-1.5-pro",
    tools=[retail_tools],
    system_instruction="You are a retail assistant. Use tools to fetch accurate product and stock information."
)

def run_vertex_agent(user_query: str) -> str:
    chat = model.start_chat()
    
    # Tool implementations
    def execute_tool(tool_name: str, tool_args: dict) -> str:
        if tool_name == "get_product_info":
            # Call your product API
            return f"Product: {tool_args['product_code']} | Price: ₹45,000 | Category: Electronics"
        elif tool_name == "check_stock":
            return f"Stock in {tool_args.get('warehouse_city', 'all')}: 15 units available"
        return "Tool not found"
    
    # Agentic loop
    response = chat.send_message(user_query)
    
    while True:
        # Check if model wants to call a function
        if not response.candidates[0].function_calls:
            break  # No tool call = final answer
        
        function_calls = response.candidates[0].function_calls
        function_responses = []
        
        for fc in function_calls:
            result = execute_tool(fc.name, dict(fc.args))
            function_responses.append({
                "name": fc.name,
                "response": {"result": result}
            })
        
        # Send tool results back to model
        response = chat.send_message(function_responses)
    
    return response.text

print(run_vertex_agent("Is the phone ELEC-002 available in Bangalore?"))
```

---

## 4. Platform Comparison for Client Recommendations

**Q: A new client asks which cloud AI platform to use for their agent. How do you recommend?**

| Factor | Recommend Azure | Recommend AWS | Recommend GCP |
|---|---|---|---|
| **Existing cloud** | Client is on Azure/M365 | Client is on AWS | Client is Google Workspace/GCP |
| **LLM preference** | Best GPT-4o access (direct MS partnership) | Best Claude access (Anthropic on Bedrock) | Best Gemini access (Google's own) |
| **Enterprise compliance** | Azure Government, GDPR | AWS GovCloud | Google Workspace integration |
| **Indian clients** | Azure India Central DC | AWS Mumbai/Hyderabad | GCP Mumbai DC |
| **Existing skillset** | Team knows .NET/Azure DevOps | Team knows AWS CDK/Lambda | Team knows GCP/Kubernetes |

**Interview Talking Point:**
> *"My recommendation is always cloud-agnostic first — I'd start by understanding where the client's data already lives and what their compliance requirements are. Moving data to a new cloud just for the AI layer is rarely worth the cost and risk. If they're on Azure, we use Azure OpenAI + AI Search. If AWS, we use Bedrock. All three platforms have reached parity for standard agentic use cases."*

---

## 5. Security & Compliance for Enterprise Agents

Common compliance requirements from Indian IT clients (BFSI, healthcare, government):

```
BFSI (Banks, Insurance):
  ✅ Data must stay in India: Use Azure India Central / AWS Mumbai
  ✅ No data sent to US endpoints
  ✅ All LLM calls through a private endpoint (not public internet)
  ✅ Conversation logs encrypted at rest + in transit
  ✅ PII must be masked before reaching LLM (use Azure Presidio)

Healthcare (HIPAA/PDPA):
  ✅ PHI (patient health info) must NEVER enter LLM prompt
  ✅ De-identify before any AI processing
  ✅ Audit logs for every tool call (who accessed what, when)

Government:
  ✅ Use only approved models (not all open models cleared)
  ✅ Sovereign cloud deployments required
  ✅ Air-gapped options: deploy open-source LLM on-prem (Llama via vLLM)
```

```python
# Azure Presidio — PII scrubbing before LLM
from presidio_analyzer import AnalyzerEngine
from presidio_anonymizer import AnonymizerEngine

analyzer = AnalyzerEngine()
anonymizer = AnonymizerEngine()

def scrub_pii(text: str) -> str:
    """Remove PII before sending to LLM."""
    results = analyzer.analyze(
        text=text,
        entities=["PHONE_NUMBER", "EMAIL_ADDRESS", "CREDIT_CARD", "IN_AADHAAR", "PERSON"],
        language="en"
    )
    anonymized = anonymizer.anonymize(text=text, analyzer_results=results)
    return anonymized.text

# Usage before any LLM call
user_message = "My name is Rahul Sharma, Aadhaar 1234-5678-9012. My order is stuck."
safe_message = scrub_pii(user_message)
# → "My name is <PERSON>, Aadhaar <IN_AADHAAR>. My order is stuck."

response = llm.invoke(safe_message)  # LLM never sees the actual PII
```
