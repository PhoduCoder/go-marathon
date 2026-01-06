## Design patterns notes on concurrency 

It is important to know which coroutines own the channel 
and which utilize them

Unidirectional channels help us with this problem
channel owners have a write-access view into the channel (chan or chan<-),
and channel utilizers only have a read-only view into the channel (<-chan)

The coroutine that owns the channel [Channel owners]
a) Should create the channel
b) Perform write to the channel or pass ownership to another coroutine 
c) Close the channel 

All the above are encapsulated and exposed via a reader channel

## chan-design1.go has the implementation

