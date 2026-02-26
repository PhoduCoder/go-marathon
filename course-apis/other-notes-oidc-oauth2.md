
1. OAuth 2.0 vs OIDC
OAuth 2.0

Authorization framework

Grants API access

Defines flows

Does NOT define token structure

OIDC

Built on OAuth 2.0

Adds identity layer

Introduces ID Token (JWT)

2. Token Types
ID Token

Purpose: Authentication
Audience: Client app
Format: JWT (required)

Required claims:

iss

sub

aud

exp

iat

Access Token

Purpose: API access
Audience: Resource server
Format: Opaque or JWT

Common claims (if JWT):

iss

sub

aud

exp

scope

3. Critical Claims Explained
iss

Token issuer
Must match expected authorization server

sub

Unique user identifier
Must be:

Stable

Never reassigned

Unique per issuer

Identity key is:

(iss, sub)
aud

Intended recipient(s)

Prevents:

Token replay across apps

Token substitution attacks

Validation rule:

my_identifier ∈ token.aud
exp

Expiration time
Prevents long-term replay

scope

Permission set for APIs

Example:

read:orders write:orders
4. Authorization Code Flow (High-Level)

Client redirects user to Authorization Endpoint

User authenticates

Authorization server issues authorization code

Client exchanges code at Token Endpoint

Authorization server returns:

ID token

Access token

5. Federation Model Example (Amazon + Google)
External Identity

Google issues ID token:

iss = accounts.google.com
sub = 11384723984723984723
aud = amazon-client-id

Amazon:

Validates token

Maps (iss, sub) to internal user

Creates Amazon session

Internal Authorization

Amazon services use:

Amazon session

Amazon internal tokens

Google tokens are not reused.

6. Authorization Server Endpoints

Minimum OAuth:

Authorization Endpoint

Token Endpoint

Optional:

Introspection Endpoint (for opaque tokens)

OIDC adds:

JWKS Endpoint

UserInfo Endpoint

Discovery Endpoint

7. Security Principles

Always validate:

Signature

iss

aud

exp

Never:

Trust email as unique identifier

Skip audience validation

Accept tokens from unknown issuer

8. Mental Model

OAuth = “Can this app access that API?”

OIDC = “Who is this user?”

aud = “Who is allowed to accept this token?”

sub = “Who is this token about?”
============