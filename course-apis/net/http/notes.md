
# First is the handler interface

type Handler interface {
    serveHTTP(ResponseWriter, *Request)
}


func ListenAndServe(addr string, handler Handler) error 

func ListenAndServeTLS(addr, certfile, keyfile string, handler Handler) error 


# ServerMUX notes 
When you want to do routing , you have many choices 

a) Get the URL and do a switch case, while possible, THIS IS NOT ELEGANT 

b) Use a serverMux, create a new ServerMux and then call the serverMux.Handle(/path, handler)
in the http.ListenAndServe(":8080", mux), we pass the mux 

c) However there exist a default server mux and we can use that too
That saves us from defining a mux and we can just use the http.Handle function

d) Now one more way is to go from handle to handleFunc 

