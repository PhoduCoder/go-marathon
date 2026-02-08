
# First is the handler interface

type Handler interface {
    serveHTTP(ResponseWriter, *Request)
}


func ListenAndServe(addr string, handler Handler) error 

func ListenAndServeTLS(addr, certfile, keyfile string, handler Handler) error 

