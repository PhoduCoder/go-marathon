## TCP server using Net package

Simple three steps
```
Create a listener
Accept a connection
Read/Write to a connection
Close the connection 
Close the listener
```

#  Create a listener
```
li, err := net.Listen("tcp", ":8080")
defer li.Close()
```

# Accept a connection 
```
conn, err := li.Accept()
defer conn.Close()
```

# Read or Write to a connection
Connection is a type that implements write and read methods which takes a slice of bytes
Since it implements read and write methods, it implements Reader and Writer Interface

So every function which takes a reader or a writer, one can use conn
