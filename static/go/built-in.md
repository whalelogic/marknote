# Go Standard Library & Built-in Quick Reference Guide

> A curated, production-grounded cheat sheet of the most frequently used built-in functions, types, and standard library primitives in Go — organized for fast, syntax-anchored lookup.

---

## Table of Contents

1. [Built-in Functions & Primitives](#1-built-in-functions--primitives)
2. [I/O & File System (`io`, `os`, `io/fs`)](#2-io--file-system-io-os-iofs)
3. [String Manipulation & Formatting (`strings`, `fmt`, `strconv`)](#3-string-manipulation--formatting-strings-fmt-strconv)
4. [Data Serialization (`encoding/json`)](#4-data-serialization-encodingjson)
5. [Time & Durations (`time`)](#5-time--durations-time)
6. [Networking & HTTP (`net/http`)](#6-networking--http-nethttp)
7. [Concurrency & Synchronization (`sync`, `context`)](#7-concurrency--synchronization-sync-context)
8. [Common Idiomatic Patterns](#8-common-idiomatic-patterns)

---

## 1. Built-in Functions & Primitives

Go provides a set of predeclared functions and types that do not require an import statement. These are essential for managing dynamic data structures and memory allocation.

---

### `make`

Initializes and allocates slices, maps, or channels. Unlike `new`, `make` returns an initialized (not zeroed) value of type `T`.

```go
// Slice: func make([]T, len, cap) []T
slice := make([]string, 0, 10)

// Map: func make(map[K]V, hint) map[K]V
hashmap := make(map[string]int, 100)

// Channel: func make(chan T, buffer) chan T
ch := make(chan bool, 1)
```

---

### `new`

Allocates zeroed storage for a new item of type `T` and returns its address (`*T`).

```go
// func new(T) *T
ptr := new(int) // *ptr is 0
```

---

### `append`

Appends elements to the end of a slice. If the underlying capacity is insufficient, a new array is allocated.

```go
// func append(slice []T, elems ...T) []T
slice := []int{1, 2}
slice = append(slice, 3, 4)

// Appending a slice to another slice using the variadic ... unpacker
anotherSlice := []int{5, 6}
slice = append(slice, anotherSlice...)
```

---

### `copy`

Copies elements from a source slice to a destination slice. Returns the number of elements copied — the minimum of `len(dst)` and `len(src)`.

```go
// func copy(dst, src []T) int
src := []int{1, 2, 3}
dst := make([]int, len(src))
bytesCopied := copy(dst, src)
```

---

### `delete`

Removes the element with the specified key from a map. If the key is not present, `delete` is a no-op.

```go
// func delete(m map[K]V, key K)
m := map[string]int{"a": 1, "b": 2}
delete(m, "a")
```

---

### `len` & `cap`

- **`len`** — Returns the length of a slice, map, string, array, or channel buffer.
- **`cap`** — Returns the capacity of a slice, array, or channel buffer.

```go
// func len(v Type) int
// func cap(v Type) int
slice := make([]int, 3, 5)
l := len(slice) // 3
c := cap(slice) // 5
```

---

### `panic` & `recover`

Used to manage abrupt control flow. `panic` stops ordinary execution; `recover` allows a program to intercept a panicking goroutine inside a deferred function.

```go
// func panic(v any)
// func recover() any
defer func() {
    if r := recover(); r != nil {
        fmt.Printf("Recovered from panic: %v\n", r)
    }
}()
panic("unrecoverable failure state reached")
```

---

### `close`

Closes a channel. Must be executed only by the sender, never the receiver. Sending to or closing an already closed channel induces a panic.

```go
// func close(ch chan<- T)
ch := make(chan int)
close(ch)
```

---

## 2. I/O & File System (`io`, `os`, `io/fs`)

These packages handle persistent storage interactions, OS abstractions, and data streaming interfaces.

---

### Types & Interfaces

#### `io.Reader`

The foundational interface representing an input stream.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

#### `io.Writer`

The foundational interface representing an output stream.

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

#### `fs.FileInfo`

Describes a file system object returned by stat operations.

```go
type FileInfo interface {
    Name() string       // base name of the file
    Size() int64        // length in bytes for regular files
    Mode() FileMode     // file mode bits
    ModTime() time.Time // modification time
    IsDir() bool        // abbreviation for Mode().IsDir()
    Sys() any           // underlying data source (can return nil)
}
```

#### `os.File`

An encapsulation representing an open file descriptor. Implements `io.Reader`, `io.Writer`, `io.Closer`, and `io.Seeker`.

```go
type File struct { /* ... */ }
```

---

### Functions & Methods

#### `os.Open`

Opens a named file for reading (`O_RDONLY`).

```go
// func Open(name string) (*File, error)
file, err := os.Open("config.json")
if err != nil {
    return err
}
defer file.Close()
```

#### `os.OpenFile`

Generalized open call with custom flag configuration (e.g., `O_CREATE`, `O_WRONLY`, `O_APPEND`) and permissions.

```go
// func OpenFile(name string, flag int, perm FileMode) (*File, error)
file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
```

#### `os.Create`

Creates or truncates the named file. Maps to `O_RDWR|O_CREATE|O_TRUNC`.

```go
// func Create(name string) (*File, error)
file, err := os.Create("output.txt")
```

#### `os.ReadFile`

Reads an entire file into memory, handling open and close automatically.

```go
// func ReadFile(name string) ([]byte, error)
data, err := os.ReadFile("data.raw")
```

#### `os.WriteFile`

Writes a byte slice to a named file, creating it if necessary or truncating if it exists.

```go
// func WriteFile(name string, data []byte, perm FileMode) error
err := os.WriteFile("output.txt", []byte("payload"), 0644)
```

#### `os.Stat`

Returns a `fs.FileInfo` describing the named file. Commonly used to check for file existence.

```go
// func Stat(name string) (FileInfo, error)
info, err := os.Stat("path/to/target")
if os.IsNotExist(err) {
    // File does not exist
}
```

#### `io.ReadAll`

Reads from an `io.Reader` until EOF or error, returning the accumulated data.

```go
// func ReadAll(r Reader) ([]byte, error)
bytes, err := io.ReadAll(response.Body)
```

#### `io.Copy`

Copies from an `io.Reader` to an `io.Writer` until EOF or error. Avoids allocating a large intermediate slice.

```go
// func Copy(dst Writer, src Reader) (written int64, err error)
written, err := io.Copy(outputFile, inputFile)
```

---

## 3. String Manipulation & Formatting (`strings`, `fmt`, `strconv`)

Tools for formatting, text parsing, conversions, and optimized string construction.

---

### `fmt.Printf` & `fmt.Sprintf`

Formatted standard output vs. raw formatted string construction.

| Verb | Description |
|------|-------------|
| `%v` | Default format representation |
| `%+v` | Struct with field names included |
| `%T` | Go syntax type representation |

```go
// func Printf(format string, a ...any) (n int, err error)
// func Sprintf(format string, a ...any) string
fmt.Printf("User: %+v\n", user)
msg := fmt.Sprintf("Error code: %d", 404)
```

---

### `fmt.Errorf`

Formats a string and returns it as an `error`. Using `%w` creates a wrapped error compatible with `errors.Is` and `errors.As`.

```go
// func Errorf(format string, a ...any) error
err := fmt.Errorf("failed validation step: %w", errTarget)
```

---

### `strings.Contains`

Reports whether a substring is within a target string.

```go
// func Contains(s, substr string) bool
found := strings.Contains("containerized", "tain")
```

---

### `strings.Split`

Slices a string into substrings separated by a delimiter, returning a slice.

```go
// func Split(s, sep string) []string
elements := strings.Split("a,b,c,d", ",")
```

---

### `strings.Join`

Concatenates string slice elements into a single string using a specified separator.

```go
// func Join(elems []string, sep string) string
result := strings.Join([]string{"foo", "bar"}, "-") // "foo-bar"
```

---

### `strings.HasPrefix` & `strings.HasSuffix`

Tests whether a string begins with a prefix or ends with a suffix.

```go
// func HasPrefix(s, prefix string) bool
// func HasSuffix(s, suffix string) bool
isJSON := strings.HasSuffix("payload.json", ".json")
```

---

### `strings.ToLower` & `strings.ToUpper`

Returns a copy of the string mapped to its lowercase or uppercase equivalent.

```go
// func ToLower(s string) string
lower := strings.ToLower("GoLang")
```

---

### `strings.TrimSpace`

Removes leading and trailing whitespace as defined by Unicode.

```go
// func TrimSpace(s string) string
clean := strings.TrimSpace(" \t content \n ")
```

---

### `strconv.Atoi` & `strconv.Itoa`

- **`Atoi`** — ASCII to Integer. Equivalent to `ParseInt(s, 10, 0)`.
- **`Itoa`** — Integer to ASCII. Equivalent to `FormatInt(int64(i), 10)`.

```go
// func Atoi(s string) (int, error)
// func Itoa(i int) string
val, err := strconv.Atoi("42")
str := strconv.Itoa(100)
```

---

### `strconv.ParseBool` / `strconv.ParseFloat` / `strconv.ParseInt`

Converts string representations to their typed equivalents.

```go
// func ParseInt(s string, base int, bitSize int) (i int64, err error)
i64, err := strconv.ParseInt("1234", 10, 64)
```

---

## 4. Data Serialization (`encoding/json`)

Handles unmarshaling raw payloads into runtime types and serializing native types into JSON.

---

### Types

#### `json.RawMessage`

A raw encoded JSON value. Delays JSON decoding and allows unparsed JSON to pass through types.

```go
type RawMessage []byte
```

---

### Functions & Methods

#### `json.Marshal`

Returns the JSON encoding of an interface value.

```go
// func Marshal(v any) ([]byte, error)
bytes, err := json.Marshal(userProfile)
```

#### `json.Unmarshal`

Parses JSON-encoded data and stores the result in the value pointed to by `v`. `v` must be a pointer.

```go
// func Unmarshal(data []byte, v any) error
var cfg Config
err := json.Unmarshal(rawBytes, &cfg)
```

#### `json.NewEncoder` & `json.NewDecoder`

Designed for streaming over open network connections or file references directly.

```go
// func NewEncoder(w io.Writer) *Encoder
// func NewDecoder(r io.Reader) *Decoder
err := json.NewDecoder(request.Body).Decode(&payload)
```

---

## 5. Time & Durations (`time`)

Primitives for precision clocks, execution measurement, and time formatting.

---

### Types

#### `time.Time`

Represents an instant in time with nanosecond precision.

```go
type Time struct { /* ... */ }
```

#### `time.Duration`

Represents elapsed time between two instants as an `int64` nanosecond count.

```go
type Duration int64
```

---

### Functions & Methods

#### `time.Now`

Returns the current local time.

```go
// func Now() Time
current := time.Now()
```

#### `time.Since`

Returns the time elapsed since `t`. Shorthand for `time.Now().Sub(t)`.

```go
// func Since(t Time) Duration
start := time.Now()
// execute task...
duration := time.Since(start)
```

#### `time.Parse`

Parses a formatted string and returns the time value it represents. The layout uses the reference time `Mon Jan 2 15:04:05 MST 2006`.

```go
// func Parse(layout, value string) (Time, error)
t, err := time.Parse("2006-01-02", "2026-05-25")
```

#### `time.Time.Format`

Returns a textual representation of the time value formatted according to the layout.

```go
// func (t Time) Format(layout string) string
str := time.Now().Format("2006/01/02 15:04:05")
```

#### `time.Sleep`

Pauses the current goroutine for at least the duration `d`.

```go
// func Sleep(d Duration)
time.Sleep(500 * time.Millisecond)
```

#### `time.After` & `time.Tick`

- **`After`** — Waits for the duration to elapse, then sends the current time on the returned channel.
- **`Tick`** — Convenience wrapper for `NewTicker`; delivers time ticks at regular intervals. Use cautiously to avoid leaks.

```go
// func After(d Duration) <-chan Time
select {
case res := <-ch:
    // process result
case <-time.After(1 * time.Second):
    // handle timeout
}
```

---

## 6. Networking & HTTP (`net/http`)

Implements HTTP client and server abstractions over low-level network primitives.

---

### Types & Interfaces

#### `http.ResponseWriter`

The interface used by an HTTP handler to construct an HTTP response.

```go
type ResponseWriter interface {
    Header() Header
    Write([]byte) (int, error)
    WriteHeader(statusCode int)
}
```

#### `http.Request`

Represents an incoming server request or an outbound client request.

```go
type Request struct {
    Method        string
    URL           *url.URL
    Header        Header
    Body          io.ReadCloser
    ContentLength int64
    // ...
}
```

#### `http.Client`

An HTTP client wrapper for executing requests. Safe for concurrent use by multiple goroutines.

```go
type Client struct {
    Transport     RoundTripper
    CheckRedirect func(req *Request, via []*Request) error
    Jar           CookieJar
    Timeout       time.Duration
}
```

---

### Functions & Methods

#### `http.HandleFunc`

Registers a handler function for the given pattern in the `DefaultServeMux`.

```go
// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("OK"))
})
```

#### `http.ListenAndServe`

Listens on the TCP network address, then calls `Serve` with the handler to handle incoming requests.

```go
// func ListenAndServe(addr string, handler Handler) error
log.Fatal(http.ListenAndServe(":8080", nil))
```

#### `http.Get` / `http.Post` / `http.Do`

- **`Get` / `Post`** — Convenience wrappers for the default HTTP client.
- **`Do`** — Sends an HTTP request and returns a response; allows full control via a custom `http.Request`.

```go
// func (c *Client) Do(req *Request) (*Response, error)
resp, err := http.DefaultClient.Do(req)
if err != nil {
    return err
}
defer resp.Body.Close()
```

---

## 7. Concurrency & Synchronization (`sync`, `context`)

Primitives for coordinating concurrent workflows, protecting shared state, and managing execution lifecycles.

---

### Types & Interfaces

#### `sync.WaitGroup`

Waits for a collection of goroutines to finish.

```go
type WaitGroup struct { /* ... */ }

// Essential methods:
// func (wg *WaitGroup) Add(delta int)
// func (wg *WaitGroup) Done()
// func (wg *WaitGroup) Wait()
```

#### `sync.Mutex` & `sync.RWMutex`

- **`Mutex`** — A mutual exclusion lock to protect critical sections.
- **`RWMutex`** — A reader/writer lock. Multiple readers can hold it simultaneously; writers require exclusive access.

```go
type Mutex struct { /* ... */ }
type RWMutex struct { /* ... */ }

// Essential methods:
// func (m *Mutex) Lock()
// func (m *Mutex) Unlock()
// func (rw *RWMutex) RLock()
// func (rw *RWMutex) RUnlock()
```

#### `sync.Once`

Guarantees that a function is executed exactly once, regardless of how many goroutines invoke it.

```go
type Once struct { /* ... */ }

// Essential method:
// func (o *Once) Do(f func())
```

#### `context.Context`

Carries deadlines, cancellation signals, and request-scoped values across API boundaries and between goroutines.

```go
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key any) any
}
```

---

### Functions & Methods

#### `context.Background` & `context.TODO`

- **`Background`** — Returns a non-nil, empty `Context`. Used by `main`, initialization, and tests as the root context.
- **`TODO`** — Returns a non-nil, empty `Context` for use when the appropriate context is unclear or not yet available.

```go
// func Background() Context
ctx := context.Background()
```

#### `context.WithCancel` / `context.WithTimeout`

Returns a copy of the parent context with a new `Done` channel, closed when the cancellation function is called or the timeout expires.

```go
// func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel() // Always call cancel to release context resources
```

---

## 8. Common Idiomatic Patterns

---

### Standard Error Assignment & Validation Flow

Go enforces explicit error handling. Check `err != nil` immediately after any operation that may fail.

```go
// Explicit assignment within if scope initialization block
if err := operation(); err != nil {
    return fmt.Errorf("operation failed: %w", err)
}
```

---

### Multi-Value Map Lookup ("comma ok" idiom)

Fetching a missing key returns the zero value of the map's value type. Use two-value assignment to distinguish between a missing key and a stored zero value.

```go
// value, ok := map[key]
val, ok := cache["session_id"]
if !ok {
    return ErrSessionNotFound
}
// val is safe to use
```

---

### Type Assertions ("comma ok" idiom)

To extract a concrete value from an interface, use a type assertion. The two-value variant prevents runtime panics on failure.

```go
// value, ok := interfaceValue.(ConcreteType)
var raw any = "message text"

str, ok := raw.(string)
if !ok {
    return ErrInvalidType
}
```

---

### Non-Blocking Channel Operations

Include a `default` case in a `select` statement to perform non-blocking sends or receives.

```go
select {
case msg := <-ch:
    fmt.Println("Received message:", msg)
default:
    fmt.Println("No message queue items ready; proceeding immediately.")
}
```

---

### Worker Pool Goroutine Synchronization Pattern

Uses `sync.WaitGroup` to coordinate safe concurrent execution and teardown of multiple worker goroutines.

```go
func WorkerPool(jobs <-chan int, results chan<- int, workerCount int) {
    var wg sync.WaitGroup

    for w := 1; w <= workerCount; w++ {
        wg.Add(1)
        go func(workerID int) {
            defer wg.Done()
            for job := range jobs {
                results <- job * 2
            }
        }(w)
    }

    // Wait for all workers to finish before closing the results channel
    go func() {
        wg.Wait()
        close(results)
    }()
}
```

---

*Reference covers Go standard library as of Go 1.21+.*
