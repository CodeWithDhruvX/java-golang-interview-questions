# HTTP, HTTPS, and DNS

## 1. HTTP vs HTTPS

### HTTP (HyperText Transfer Protocol)
*   Application layer protocol for transmitting hypermedia documents.
*   **Stateless**: Server doesn't retain information about past requests.
*   **Text-based**: Headers and Body are sent as plain text.

### HTTPS (HTTP Secure)
*   Extension of HTTP that uses encryption for secure communication.
*   **Encryption**: Uses **TLS (Transport Layer Security)** / SSL.
*   **Port**: HTTP uses 80, HTTPS uses 443.

## 2. SSL/TLS Handshake
How HTTPS establishes security:
1.  **Client Hello**: Client sends supported cypher suites and random number.
2.  **Server Hello**: Server picks a cypher, sends its **Digital Certificate** (Public Key) and random number.
3.  **Verification**: Client verifies the certificate with a Certificate Authority (CA).
4.  **Key Exchange**: Client encrypts a "Pre-Master Secret" using Server's Public Key. Only Server can decrypt it with its Private Key.
5.  **Session Key**: Both sides generate a Symmetric Session Key.
6.  **Secure**: All further communication is encrypted using this Session Key (Symmetric Encryption is faster).

## 3. methods & Status Codes

### Methods
*   **GET**: Retrieve resources. Idempotent.
*   **POST**: Create resources. Not Idempotent.
*   **PUT**: Update/Replace resource. Idempotent.
*   **PATCH**: Partial update.
*   **DELETE**: Delete resource.

### Status Codes
*   **2xx (Success)**: 200 OK, 201 Created.
*   **3xx (Redirection)**: 301 Moved Permanently, 304 Not Modified (Cache).
*   **4xx (Client Error)**: 400 Bad Request, 401 Unauthorized, 403 Forbidden, 404 Not Found.
*   **5xx (Server Error)**: 500 Internal Server Error, 502 Bad Gateway.

## 4. DNS (Domain Name System)
Translates domain names (google.com) to IP addresses (142.250.1.1).

### The Lookup Flow
1.  **Browser Cache**: Checks if recently visited.
2.  **OS Cache**: Checks hosts file or OS cache.
3.  **Recursive Resolver**: ISP's DNS Server (e.g., 8.8.8.8) is queried.
4.  **Root Server (.)**: Resolver asks Root "Where is .com?"
5.  **TLD Server (.com)**: Root replies with TLD server IP. Resolver asks TLD "Where is google.com?"
6.  **Authoritative Name Server**: TLD replies with Google's Name Server. Resolver asks "What is IP of google.com?"
7.  **Final Answer**: Authoritative server returns the IP.

## 5. Interview Questions
1.  **What happens when you type google.com in browser?**
    *   *Ans*: DNS Lookup -> TCP Handshake -> TLS Handshake -> HTTP GET Request -> Server Processing -> HTTP Response -> Browser Rendering.
2.  **What is a Sticky Session?**
    *   *Ans*: A Load Balancer feature where requests from the same user (IP or Cookie) are always routed to the same backend server. Useful if local state is stored on server (though generally bad practice in stateless architecture).
3.  **Difference between 401 and 403?**
    *   *Ans*:
        *   **401 Unauthorized**: "Who are you?" (Authentication failed/missing).
        *   **403 Forbidden**: "I know who you are, but strict NO." (Authorization failed, e.g., Admin only page accessed by user).
