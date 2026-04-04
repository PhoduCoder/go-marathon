# What not to pass in contexts

Don’t misuse context

Context is NOT for passing business data.

Use it only for:

request-scoped metadata (trace IDs, auth info)
cancellation / deadlines

❌ Avoid:

large objects
core application data
things required for function logic

Rule of thumb
Small metadata → OK
Anything important for logic → pass explicitly as function params

=======

Pitfalls two in contexts 

Avoid using built-in types (like string) as keys
This can cause collisions across packages.

To understand this better, say you passed like this 

```
ctx := context.WithValue(ctx, "user", "gaurav")
```

The "user" is not scoped to our package.

Say package one 
ctx = context.WithValue(ctx, "user", "gaurav")

Package two 
ctx = context.WithValue(ctx, "user", 12345) // user ID as int

Retrieval

val := ctx.Value("user")

Depends on who wrote last 
"gaurav"
or "12345"

=====
How to avoid this?

```
type contextKey string //declaring contextKey as a separate type 

const metaKey contextKey = "meta" // creating a variable metaKey of type contextKey

//Passing that now instead metaKey
ctx := context.WithValue(ctx, metaKey, &RequestMeta{
    UserID: "123",
})
```

=====

complex structs are fine to be passed in contexts values 

You can pass:

structs
pointers to structs (preferred for large objects)
maps, slices, etc.

👉 Prefer pointer:

```
ctx := context.WithValue(ctx, metaKey, &meta)
```

=====

👉 Strings do NOT get namespaced by package.
👉 Custom types DO.

======