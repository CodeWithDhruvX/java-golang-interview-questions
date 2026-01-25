# OAuth 2.0 & OIDC (OpenID Connect)

## 1. The Problem (Why OAuth?)
Before OAuth, if you wanted Yelp to access your Google Contacts, you had to give Yelp your Google password. This is insecure.
**OAuth 2.0** allows you to give Yelp a "Valet Key" (Access Token) that allows it to only read contacts, without sharing your password.

## 2. Roles
1.  **Resource Owner**: You (The user).
2.  **Client**: The application (Yelp, Spotify).
3.  **Authorization Server**: The server that issues tokens (Google, Facebook).
4.  **Resource Server**: The server ensuring the data (Gmail API).

## 3. The Flow (Authorization Code Grant)
This is the standard flow for Server-side apps.

1.  **User clicks "Login with Google"**.
2.  **Redirect**: Client redirects Browser to Google's Auth Server.
    *   `GET /authorize?response_type=code&client_id=XYZ&scope=email`
3.  **Consent**: User logs into Google and clicks "Allow".
4.  **Callback**: Google redirects Browser back to Client with a temporary **Auth Code**.
    *   `GET /callback?code=AUTH_CODE_123`
5.  **Exchange**: Client (backend) sends Auth Code + Client Secret to Google.
    *   `POST /token`
6.  **Token**: Google validates code and returns **Access Token** (and Refresh Token).
7.  **Access**: Client uses Access Token to call Resource Server.

## 4. OIDC (OpenID Connect)
OAuth 2.0 is for **Authorization** (Accessing resources). It doesn't tell the Client *who* the user is.
**OIDC** adds **Authentication** layer on top of OAuth 2.0.

*   **ID Token**: A JWT (JSON Web Token) returned alongside the Access Token.
    *   Contains user info: `{ "sub": "123", "name": "John Doe", "email": "john@gmail.com" }`.
    *   Signed by the Identity Provider (Google).

## 5. Token Types
1.  **Access Token**: Short-lived (e.g., 1 hour). Used to access APIs. Bearer token.
2.  **Refresh Token**: Long-lived (e.g., 30 days). Used to get a new Access Token when the old one expires, without forcing the user to log in again.
3.  **ID Token**: Used by the Client to identify the user (UI customization).

## 6. Interview Questions
1.  **Difference between 401 and 403?**
    *   *Ans*: 401 = Missing/Invalid Token (Authentication). 403 = Valid Token, but not allowed scope (Authorization).
2.  **Why Authorization Code flow and not Implicit flow?**
    *   *Ans*: Implicit flow returns tokens in the URL hash, which is visible to browser history and JavaScript (XSS risk). Code flow keeps tokens on the backend.
3.  **How to revoke access?**
    *   *Ans*: Cannot easily revoke Access Tokens (stateless JWTs) until they expire. This is why they are short-lived. You revoke the Refresh Token database entry.
