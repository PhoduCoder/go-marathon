Example scenario:

Order Service → wants to call → Payment Service

They both trust → Authorization Server

So we have:

Role	                Example
Client	                Order Service
Authorization Server	Auth Server
Resource Server	        Payment Service
============

Step 1 — Client Authenticates to Authorization Server

Order Service sends:

```
POST /token
Content-Type: application/x-www-form-urlencoded

grant_type=client_credentials
client_id=order-service
client_secret=super-secret
scope=charge:customer
```

What’s happening:

Order Service proves its identity

No user involved

This is service authentication

==========
Step 2 — Authorization Server Validates Client

Auth server checks:

Is client_id valid?

Is client_secret correct?

Is it allowed requested scope?

Is it allowed to access Payment Service?

If valid → proceed.

==========
Step 3 — Authorization Server Issues Access Token

Example JWT access token:

{
  "iss": "https://auth.internal",
  "sub": "order-service",
  "aud": "payment-service",
  "scope": "charge:customer",
  "exp": 1716239022,
  "iat": 1716235422
}

Important differences from OIDC:

No ID token

No nonce

No user identity

sub = service identityStep 3 — Authorization Server Issues Access Token

```
Example JWT access token:

{
  "iss": "https://auth.internal",
  "sub": "order-service",
  "aud": "payment-service",
  "scope": "charge:customer",
  "exp": 1716239022,
  "iat": 1716235422
}
```

Important differences from OIDC:

No ID token

No nonce

No user identity

sub = service identity

=======
Step 4 — Client Calls Resource Server

Order Service now calls:

POST /charge
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9..

=======
Step 5 — Resource Server Validates Token

Payment Service checks:

Signature valid?

iss trusted?

aud == payment-service?

exp not expired?

Scope allows charge:customer?

If yes → execute request.
=======

What Is NOT Present

Client Credentials Flow does NOT include:

❌ Authorization endpoint
❌ Redirects
❌ Login screen
❌ User consent
❌ ID token
❌ aud = client_id

Because:

There is no user to authenticate.
======