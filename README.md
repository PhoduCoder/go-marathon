# go-marathon

```
Weeks 1–4: Go Fundamentals (exact tasks)
Week 1

Day 1 – Basics

Write func Hello(name string) string.

Expected: Hello("Gaurav") == "Hello, Gaurav".

CLI: print result from main().

Day 2 – Control flow

func FizzBuzz(n int) []string

“Fizz” (3), “Buzz” (5), “FizzBuzz” (15), else number.

Test 1..20 exact slice.

Day 3 – Functions + errors

func Divide(a, b int) (int, error)

b==0 → ErrDivideByZero.

Tests: normal, divide by zero uses errors.Is.

Day 4 – Arrays/Slices

func Reverse(xs []int) []int

func Max(xs []int) (int, error); empty → error.

Tests: even/odd length, negative numbers.

Day 5 – Maps

func WordCount(s string) map[string]int

Split on spaces/punct (use unicode), lowercase.

Test: "Go go, go!" → {"go":3}.

Day 6 – Structs & methods

type User struct { Name string; Age int }

Methods: IsAdult() bool, Birthday().

Tests: boundary Age==18.

Day 7 – Mini-project

Contact book:

type Contact struct{ Name, Phone string }

type Book struct{ byName map[string]Contact }

Add(c Contact) error (dup name -> error), Get(name string) (Contact,bool), Delete(name string) bool, List() []Contact (sorted by name).

Tests for all ops.

Week 2

Day 8 – Pointers

func Swap(a, b *int)

Test with two ints.

Day 9 – Interfaces

type Shape interface{ Area() float64 }

type Circle{R float64}, type Rect{W,H float64}

func TotalArea(shapes ...Shape) float64

Tests with 2 circles + 1 rect.

Day 10 – Error wrapping

var ErrBadJSON = errors.New("bad json")

func LoadUserJSON(r io.Reader) (User, error); wrap parse errors with %w, empty → ErrBadJSON.

Tests: errors.Is(err, ErrBadJSON).

Day 11 – Packages

Split shape and shape_test.

Add Perimeter() to Circle/Rect and tests.

Day 12 – Table-driven tests

For FizzBuzz and Divide, convert to table tests with subtests.

Day 13 – Benchmarks

Add BenchmarkWordCount_Short/Medium/Long.

Ensure go test -bench . runs.

Day 14 – Review

Create Makefile with test, bench, lint (if using golangci-lint).

All tests green.

Week 3

Day 15 – JSON marshal/unmarshal

type Config struct{ Addr string; TimeoutMS int }

func (c Config) ToJSON() []byte

func FromJSON(b []byte) (Config, error)

Tests: round-trip.

Day 16 – File I/O

func Save(path string, v any) error

func Load[T any](path string, out *T) error

Tests with temp dir.

Day 17 – Goroutines

func SpawnPrint(n int, out chan<- int) (send 1..n)

Test collects all numbers in any order.

Day 18 – Channels (producer/consumer)

func MapInts(in <-chan int, fn func(int) int) <-chan int

Test: square numbers.

Day 19 – WaitGroup

func ParallelMap(xs []int, workers int, fn func(int) int) []int

Test: workers 1 and 4 yield same set (order doesn’t matter).

Day 20 – Context cancel

func FetchAll(ctx context.Context, urls []string) ([]string, error)

Use httptest.Server; cancel after first error.

Tests include timeout.

Day 21 – Project

Concurrent URL fetcher CLI:

fetch --concurrency 5 url1 url2 ... → print status codes.

Tests using cobra’s command execution.

Week 4

Day 22 – Flags

Build wc clone:

wc --lines --words file.txt

Functions: CountLines(r io.Reader) int, CountWords(r io.Reader) int.

Tests with strings.Reader.

Day 23 – Sentinel errors

var ErrNotFound = errors.New("not found")

type Store interface{ Get(k string) (string,error); Put(k,v string) }

type MemStore map[string]string

Tests: Get on missing → ErrNotFound.

Day 24 – Sorting

func TopNByValue(m map[string]int, n int) []string (stable by key on ties)

Tests for ties and n>len.

Day 25 – More testing

Table tests for Store, inject fake clock where relevant.

Day 26 – Simple logger

type Logger struct{ mu sync.Mutex; w io.Writer }

Methods: Info(msg string, kv ...any), Error(err error, msg string, kv ...any)

Tests assert output lines.

Day 27 – To-Do CLI

Commands: add <text>, list, done <id>

Persist to JSON using Day 16 funcs.

Tests run command handlers directly.

Day 28 – Refactor

Move shared helpers to internal/.

All tests pass; go vet clean.

Weeks 5–8: From Basics to Mini-Schedulers
Week 5

Day 29 – DI via interfaces

type KV interface{ Get(string)(string,error); Set(string,string) }

FileKV (JSON file), MemKV.

func NewCache(backing KV) *Cache with read-through/write-through.

Tests use MemKV fake; verifies no file I/O.

Day 30 – Mocks

Create type FakeKV struct{ Data map[string]string; GetHook func(key string) }

Test that GetHook is called; assert interactions.

Day 31 – Binpacking core

Types:

type Pod struct{ Name string; CPU, Mem int } // millicores, MiB
type Node struct{ Name string; CPUCap, MemCap int }


func Pack(pods []Pod, nodes []Node) (placement map[string][]string, unsched []string)

Use first-fit-decreasing by CPU.

Tests: exact fit, overcommit, empty.

Day 32 – Underutilized

func Underutilized(placement map[string][]Pod, nodes map[string]Node, threshold float64) []string

A node is underutilized if usedCPU/CPUCAP < threshold and has ≥1 pod.

Tests: thresholds 0.3/0.5/0.8.

Day 33 – Dry-run consolidation

func PlanDeletes(pods []Pod, nodes []Node, threshold float64) (candidates []string)

For each underutilized node, try removing it and repack remaining pods; if all place → candidate.

Tests with 3 nodes, multiple cases.

Day 34 – “Critical” pods constraint

Treat pods with label critical=true as immovable:

Change Pod to Labels map[string]string.

Update PlanDeletes to reject nodes hosting immovable pods.

Tests: candidate excluded if any critical=true on it.

Day 35 – CLI

consoli plan --pods pods.yaml --nodes nodes.yaml --threshold 0.5

Output JSON:

{"deleteCandidates":["node-b","node-d"]}


E2E test with golden files.

Day 36 – Logging + metrics

Expose /metrics via promhttp with:

consoli_plans_total

consoli_nodes_underutilized

Test: scrape the handler and assert metrics exist.

Week 6

Day 37 – Worker pool

type Pool[T any,R any] struct{ ... }

Methods: Submit(T), Results() <-chan R, Close()

Test: 100 tasks, 5 workers.

Day 38 – Backpressure

Pool has bounded queue; submitting >cap blocks.

Test uses goroutine + select with timeout.

Day 39 – Contextful pool

func NewPool[T,R](ctx context.Context, workers, cap int, fn func(context.Context,T)(R,error))

Cancel ctx → pool drains and stops.

Tests cancel mid-flight.

Day 40 – Integrate pool with packing

Parallelize PlanDeletes by testing each node in the pool.

Tests assert same results as serial.

Day 41 – Determinism

Make outputs sorted and deterministic.

Add property test: run 50 times; same output.

Day 42 – Docs

Write a README with Flow Memo:

ListNodes -> Pack -> Underutilized -> SimulateRemove -> Plan.

Keep to 1 page.

Weeks 7–8: First Controller Runtime (local, no cluster needed)

Day 43 – Fake Reconcile loop

Without Kubernetes: write func Reconcile(state *State) error that:

Ensures state.DeployReplicas[name] == desired[name].

Tests: idempotence—running twice doesn’t change outcome.

Day 44 – Finalizers

Add Finalizers map[string]bool to State.

Delete(name) sets a tombstone; Reconcile cleans related config if finalizer present, then removes finalizer.

Tests for cleanup.

Day 45 – Time-based policy

powerSchedule: weekendOff: true

Reconcile scales deployments to 0 on Sat/Sun (inject a Clock interface).

Tests with fake clock across boundary.

Day 46 – Rate limiting

Add RequeueAfter semantics:

func Reconcile(...) (requeueAfter time.Duration, err error)

Tests that weekend policy requeues for “time to Monday 00:00”.

Day 47 – Metrics

Increment reconcile_total, expose handler; test scrape.

Day 48 – E2E

Compose: read desired state JSON → run Reconcile loop every 1s for 5s → log actions.

Test with small scenario.

Weeks 9–12: Apply to OSS (Levels 100–400 condensed)

Week 9 – Level 100

Exact grep tasks (run and paste notes):

rg -n 'func main\\(' $(go env GOPATH)/pkg/mod -g '!**/vendor/**'

On kubebuilder sample: locate Reconciler.Reconcile and list 3 inputs, 3 effects in a text file.

Write a “flow memo” for that sample (10 bullets max).

Week 10 – Level 200

Generate go-callvis graph for the sample; export SVG.

Set 2 breakpoints with dlv in Reconcile; step once; write down the variable values you observe.

Week 11 – Level 300

Replace your Pack with a simpler best-fit; keep tests. Ensure both pass same fixtures.

Week 12 – Level 400

Hook your PlanDeletes into the fake reconciling loop:

If candidates non-empty, print "Delete:" lines.

Add test that verifies printed plan for a given input.

Test Names & Commands

Use table tests: TestDivide, TestWordCount, TestPack_ExactFit, TestPlanDeletes_CriticalPods, etc.

Run: go test ./... -race

Bench (where present): go test -bench . ./...

Data Files (examples to create in testdata/)

pods.yaml

- name: p1
  cpu: 200
  mem: 256
- name: p2
  cpu: 500
  mem: 256
- name: sys
  cpu: 50
  mem: 64
  labels: {critical: "true"}


nodes.yaml

- name: node-a
  cpucap: 1000
  memcap: 1024
- name: node-b
  cpucap: 600
  memcap: 512

```