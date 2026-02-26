Amazon = Client

Google = Authorization Server

You = Resource Owner

Amazon backend = Resource Server (for Amazon's own data)

==========
Step 1 — “Link to Service ABC?”

You’re on Amazon login page.

You click: “Sign in with Google”

In diagram:

User → App XYZ (Amazon)
===========

Step 2 — Authorization Request

Amazon redirects your browser to:

https://accounts.google.com/o/oauth2/v2/auth

That is Google’s:

➡ Authorization Endpoint

This request includes:

response_type=code
client_id=amazon-client-id
redirect_uri=https://amazon.com/callback
scope=openid email profile
state=xyz123
nonce=abc456

This is OIDC not just oAuth2.0 because of scope=openid

===============

Step 3 & 4 — Google Shows Login Screen

Google displays:

Login form

“Amazon wants access to your email and profile”

This is the green “Authorization Page” box in your diagram.

You enter credentials.
================

Step 5 — You Approve

You click “Allow”.

Now Google knows:

You authenticated

You consented
================

Step 6 — Authorization Code Issued

Google redirects browser back to:

https://amazon.com/callback?code=AUTH_CODE_123

This is the authorization code (short-lived).

At this point:

❗ Amazon still has NO tokens.
It only has a temporary code.
===============
Step 7 — Token Endpoint

Now Amazon’s backend server makes a secure call to Google:

POST https://oauth2.googleapis.com/token

It sends:

authorization code

client_id

client_secret

redirect_uri

This is the Token Endpoint in the diagram
===============
Step 8 — Google Returns Tokens

Google responds with:

{
  "access_token": "ACCESS123",
  "id_token": "IDTOKEN456",
  "expires_in": 3600,
  "token_type": "Bearer"
}

Now here is where OAuth and OIDC split.
==========

OAUTH VS OIDC DIFFERENCE

Access Token (OAuth)

Purpose:
👉 Allows Amazon to call Google APIs.

Example:
Amazon might call:

https://www.googleapis.com/oauth2/v3/userinfo

With:

Authorization: Bearer ACCESS123

Google verifies token and returns your profile info.

Access Token = “permission to access API”
=========

ID Token (OIDC Token)

This is new.

The ID Token is a JWT that proves:

You authenticated

Google authenticated you

It was meant for Amazon

Example ID token payload:


```
{
  "iss": "https://accounts.google.com",
  "sub": "google-user-123",
  "aud": "amazon-client-id",
  "exp": 1716239022,
  "iat": 1716235422,
  "email": "you@gmail.com"
}
```
Amazon verifies:

Signature (using Google’s public keys)

aud matches Amazon

iss is Google

exp valid

This token is NEVER sent to Google APIs.

It is for Amazon only.

================
What Actually Logs You Into Amazon?

Not the access token.

Not the authorization code.

👉 The ID token.

Amazon uses:

sub as your unique Google ID

email to match or create account

Then Amazon creates its OWN session cookie.

After that:
Google tokens are no longer used.

================
Important: Introspection Endpoint

Used only when:

Access tokens are opaque

Resource server cannot validate locally

================
Amazon + Google Example:

Does Google use introspection?

Usually no.

Because:

Google access tokens are JWT or verifiable

Amazon validates ID token using Google’s public keys
================
What Actually Happens After Login?

After Google authentication:

Amazon:

Verifies ID token

Creates its own user session

Sets Amazon session cookie

From that moment:
You are logged into Amazon — not Google.
==================
aud is REQUIRED in OIDC ID tokens

aud is not defined by core OAuth 2.0

But modern OAuth access tokens (especially JWT ones) almost always include aud
===============
OAuth 2.0 does not define token structure at all.

Access tokens can be:

Opaque string

JWT

Anything

If it's opaque → no visible aud.

If it's JWT → most implementations include aud, but it's convention, not core spec requirement.

Modern best practice:

Always include aud in JWT access tokens.

Why?
To prevent token replay across APIs.
============



=====

In Client Credentials Flow, you use OAuth 2.0, not OpenID Connect.

Let’s break this down very clearly.

🔷 Why Client Credentials Flow Is OAuth2 Only
Core reason:

There is no end-user involved.

And:

OIDC = Identity layer for authenticating a user
Client Credentials = Machine-to-machine authorization

So OIDC does not apply.

🔹 What Is Client Credentials Flow?

Defined in OAuth 2.0

Used when:

A backend service talks to another backend service

A daemon job calls an API

A microservice calls another microservice

No human is present

Example:

Order Service → Payment Service

Analytics Job → Reporting API

🔷 Flow Overview

Client (service) sends:

POST /token
grant_type=client_credentials
client_id=order-service
client_secret=supersecret

Authorization Server returns:

{
  "access_token": "abc123",
  "token_type": "Bearer",
  "expires_in": 3600
}

Service calls API:

Authorization: Bearer abc123

That’s it.

No:

Redirect

Login page

ID token

User session

🔷 Why OIDC Is Not Used Here

OpenID Connect exists to answer:

Who is the authenticated user?

Client Credentials Flow has no user.

There is no:

sub representing a person

login screen

user consent screen

The “subject” in client credentials flow is the client itself.

Example JWT access token might look like:

{
  "iss": "https://auth.internal",
  "sub": "order-service",
  "aud": "payment-service",
  "exp": 1716239022,
  "scope": "charge:customer"
}

Notice:

sub = service identity

Not a human user

No ID token exists

🔷 When Do We Use OIDC?

Use OIDC when:

A user logs in

You need authentication

You need identity claims (email, name, etc.)

Example flows:

Authorization Code Flow

Authorization Code + PKCE

Hybrid Flow

🔷 Clean Comparison
Feature	Client Credentials	OIDC Auth Code
User involved	❌ No	✅ Yes
Login screen	❌	✅
ID token	❌	✅
Access token	✅	✅
Purpose	Service auth	User auth
🔷 Important Subtlety

You can use OIDC server infrastructure to issue client credentials tokens.

But the flow itself:

Is purely OAuth 2.0

Does not involve OIDC identity layer

Think of OIDC as:

OAuth + user identity

Remove the user → you're back to plain OAuth.