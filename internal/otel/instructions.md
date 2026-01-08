## Creating Traces Steps

```
To create traces you need:

a) A tracer provider (the SDK “engine”).

b) An exporter (where spans go: terminal, Tempo, etc).

c) Spans around the work you want to measure.

d) Context propagation inside the process (context.Context) so child spans attach to parents.
```


# Download the go binaries 
go get go.opentelemetry.io/otel go.opentelemetry.io/otel/exporters/stdout/stdouttrace go.opentelemetry.io/otel/sdk/trace


#OTel packages:
```
otel: global entry point (get tracer, set tracer provider)

stdouttrace: exporter that prints spans to terminal 

sdktrace: the SDK implementation of tracing (tracer provider, batcher, etc)
```
