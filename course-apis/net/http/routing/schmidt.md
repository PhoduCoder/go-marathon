## Julien Schmidt Router

This is a more performant http router compared to http.NewServeMux()
It is also more flexible with parameters handing

It uses a radial tree [ which holds paths in tree like structure]
/car /cart /catalogue/ cat paths are stored as

ca -> r -> t 
   -> t -> alogue 



Handler in this router also has the params in the function signature
```
func handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
```

```
func handler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    name := ps.ByName("name")
    fmt.Fprintf(w, "hello, %s!\n", name)
}
```

httpRouter.Params is basically 
```
type Params []Param

type Param struct {
    Key   string
    Value string
}
```

ByName() just loops through that slice and returns the matching value.