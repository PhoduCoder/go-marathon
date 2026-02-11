Even after all the methods available there is the HandlerFunc

Why is this needed?

We know that for any type to be passed as handler, it has to implement the method with below signature.

ServeHTTP(ResponseWriter, *request)

===
But usually one writes the function like this 

func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello"))
}

Now this can't be substitued by a handler since it doesn't implement the ServeHTTP method 

So you define a type whose underlying type is a function 
type HandlerFunc func(ResponseWriter, *Request)
HandlerFunc is a type whose underlying type is a function.

Now you use the above type and implement the ServeHTTP method for it 

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}

==========
so instead of creating a empty struct say and then creating a ServeHTTP method on that struct, you define the method on the function itself.

=======
When you define 

type HandlerFuncNew func(ResponseWriter, *Request)
Any function with this signature[func(ResponseWriter, *Request)] becomes a HandlerFuncNew type 

Eg. 
NAMED FUNCTION 
func hello(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("hello"))
}

var h HandlerFuncNew = hello
OR 
h := HandlerFuncNew(hello)

ANONYMOUS FUNCTION 
h := HandlerFuncNew(func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("anonymous"))
})

======
Now once you have defined your new type ( which behind the scene represent a function), you will 
then define the ServeHTTP method as below 

```
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
```

======
So these three steps involved

a) define a function which has (w ResponseWriter, r *Request) signature

b) Define a new type which is a function defined with (w ResponseWriter, r *Request) signature
c) then implement the serveHTTP method passing the function type as the receiver argument.

d) After that use the function defined in a) in http.Handle function as a handler 


======

```
package main
import "net/http" 

func hello(w http.ResponseWriter, r *http.Request) {
     w.Write([]byte("hi")) 
    }
type HandlerFuncNew func(http.ResponseWriter, *http.Request) 

func (f HandlerFuncNew) ServeHTTP(w http.ResponseWriter, r *http.Request) {
     f(w, r) } 

var h HandlerFuncNew = hello 

func main() {
     // default mux 
     http.Handle("/hello", h ) 
     http.ListenAndServe(":8080", nil) }
```


OR THIS WAY 

```
package main

import "net/http"

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

type HandlerFuncNew func(http.ResponseWriter, *http.Request)

func (f HandlerFuncNew) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func main() {
	// default mux
	http.Handle("/hello", HandlerFuncNew(hello)) // ✅ Handle expects a Handler

	http.ListenAndServe(":8080", nil)
}

```