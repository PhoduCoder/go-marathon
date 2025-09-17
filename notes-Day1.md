To create a module
```
Local only: go mod init helloapp

Public repo: go mod init github.com/<user>/<repo>

go mod init github.com/PhoduCoder/go-practice/go-marathon
```

Each folder is a package in Go and all the code within that folder/package
can be used by any other files within that package irrespective of whether it is 
exposed(capitalized) or not 
```
internal/hello → package hello

cmd/hello → binary entrypoint (with main.go)

Conventions:

internal/ → private code, only importable within this module.

cmd/ → one folder per binary (cmd/hello → hello binary).
```

To test your code, go to the path where go.mod is present 
and then run the below
```
go test ./internal/hello -v
```

White box testing vs black box testing 

Tests live in *_test.go

Function signature:

func TestXxx(t *testing.T) { ... }


Run tests:
```
go test ./internal/hello -v
go test ./... -v   # all tests
```

White-box vs Black-box

White-box (package hello)
Same package as code, can test unexported funcs.
No need to import the package.

Black-box (package hello_test)
Separate package, must import the package.
Tests only public API (exported names).
Runs from the same root of module
However you test function would need to evolve 
Your test function should be in a different package usually by appending the test at end 

In Go:

Any file that contains tests must end with _test.go.

That’s not just a convention — it’s how go test finds test files.

Example: hello_test.go, hello_whitebox_test.go, hello_blackbox_test.go.

Inside those files, you can use either:

package hello → white-box (same package, can test private stuff).

package hello_test → black-box (separate package, must import the real one).

#Running Programs 
```go run ./cmd/hello```

#Build binary:

```go build -o hello ./cmd/hello
./hello
```

Makefile=======

test:
	go test ./... -v

run:
	go run ./cmd/hello
=================

make test 
make run

======
Run whitebox tests 
```
go test ./internal/hello -v
```

=== RUN   TestHello
--- PASS: TestHello (0.00s)
=== RUN   TestHello_BlackBox
--- PASS: TestHello_BlackBox (0.00s)
PASS
ok      github.com/PhoduCoder/go-practice/go-marathon/internal/hello    0.299s

Run all tests 

(base) MacBookPro:go-practice alieninvader$ go test ./... -v
?       github.com/PhoduCoder/go-practice/go-marathon/cmd/hello [no test files]
=== RUN   TestHello
--- PASS: TestHello (0.00s)
=== RUN   TestHello_BlackBox
--- PASS: TestHello_BlackBox (0.00s)
PASS
ok      github.com/PhoduCoder/go-practice/go-marathon/internal/hello    (cached)
(base) MacBookPro:go-practice alieninvader$ 
(base) MacBookPro:go-practice alieninvader$ 