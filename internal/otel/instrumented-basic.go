package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

// The initTracer returns a shutdown function so that
// one can flush spans before exit
func initTracer() func(context.Context) error {

	//Create an exporter to send the spans to , in this case stdout
	exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatal(err)
	}

	//Creates the TracerProvider , sdktrace.WithBatcher buffers & exports traces in batches
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
	)

	//The tracer provider is set as the global default
	otel.SetTracerProvider(tp)

	//Main can call this and flush spans
	return tp.Shutdown
}

func main() {

	// Ensure that the spans are exported
	// before the process exits
	shutdown := initTracer()
	defer shutdown(context.Background())

	//Starting an empty context
	// Contexts are used to share contextual information between go functions and/or coroutines
	ctx := context.Background()

	// Creates a tracer object
	//demo tells which component or library created the span
	tr := otel.Tracer("demo")

	fmt.Println("start main")

	//Starts a span named "main".
	//It returns a new context (ctx) that now contains this span.
	// defer span.End() records the end time automatically.
	ctx, span := tr.Start(ctx, "main")
	defer span.End()

	doWork(ctx, tr)

	fmt.Println("end main")
}

func doWork(ctx context.Context, tr trace.Tracer) {
	fmt.Println("start doWork")

	//Creates a child span under "main" (because ctx already contains the parent span).
	ctx, span := tr.Start(ctx, "doWork")
	defer span.End()

	for i := 0; i < 3; i++ {
		//_, child := tr.Start(ctx, fmt.Sprintf("step-%d", i))
		fmt.Println("step", i)
		time.Sleep(50 * time.Millisecond)
		//child.End()
	}

	fmt.Println("end doWork")
}
