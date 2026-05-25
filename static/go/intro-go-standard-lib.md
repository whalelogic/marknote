# Go Standard Library Reference (Curated)

This document provides the most commonly used functions, methods, and types from selected Go packages. Each entry includes the function signature and an example.

---

## 📦 strings

### `strings.Contains`
```go
func Contains(s, substr string) bool
```
Checks if a substring exists within a string.
```go
fmt.Println(strings.Contains("golang", "go")) // true
```

### `strings.Split`
```go
func Split(s, sep string) []string
```
Splits a string by a separator.
```go
parts := strings.Split("a,b,c", ",")
fmt.Println(parts) // [a b c]
```

### `strings.Join`
```go
func Join(elems []string, sep string) string
```
Joins a slice of strings with a separator.
```go
s := strings.Join([]string{"a", "b", "c"}, ",")
fmt.Println(s) // "a,b,c"
```

### `strings.TrimSpace`
```go
func TrimSpace(s string) string
```
Removes leading and trailing whitespace.
```go
fmt.Println(strings.TrimSpace("   hello  ")) // "hello"
```

### `strings.ReplaceAll`
```go
func ReplaceAll(s, old, new string) string
```
Replaces all occurrences of `old` with `new`.
```go
fmt.Println(strings.ReplaceAll("foo bar foo", "foo", "baz")) // "baz bar baz"
```

---

## 📦 io/ioutil (Legacy)
*Note: Most functionality has been moved to `os` and `io` in modern Go.*

### `ioutil.ReadFile`
```go
func ReadFile(filename string) ([]byte, error)
```
Reads the entire file into memory.
```go
data, err := ioutil.ReadFile("example.txt")
if err != nil {
	panic(err)
}
fmt.Println(string(data))
```

### `ioutil.WriteFile`
```go
func WriteFile(filename string, data []byte, perm os.FileMode) error
```
Writes data to a file.
```go
ioutil.WriteFile("out.txt", []byte("hello"), 0644)
```

---

## 📦 bytes

### `bytes.Buffer`
```go
type Buffer struct
func (b *Buffer) Write(p []byte) (n int, err error)
func (b *Buffer) WriteString(s string) (n int, err error)
func (b *Buffer) String() string
```
Efficiently builds or manipulates byte slices and strings.
```go
var buf bytes.Buffer
buf.WriteString("Hello, ")
buf.WriteString("World!")
fmt.Println(buf.String()) // "Hello, World!"
```

---

## 📦 context

### `context.Background`
```go
func Background() Context
```
Returns an empty, non-nil Context.
```go
ctx := context.Background()
```

### `context.WithCancel`
```go
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
```
Returns a copy of parent with a new Done channel.
```go
ctx, cancel := context.WithCancel(context.Background())
go func() {
	time.Sleep(2 * time.Second)
	cancel()
}()
<-ctx.Done() // context cancelled
```

### `context.WithTimeout`
```go
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
```
Returns a context that is canceled when the timeout expires.
```go
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
<-ctx.Done()
fmt.Println("timeout:", ctx.Err())
```

---

## 📦 encoding/json

### `json.Marshal`
```go
func Marshal(v any) ([]byte, error)
```
Converts a Go value to JSON.
```go
type User struct { Name string; Age int }
user := User{"Alice", 30}
b, _ := json.Marshal(user)
fmt.Println(string(b)) // {"Name":"Alice","Age":30}
```

### `json.Unmarshal`
```go
func Unmarshal(data []byte, v any) error
```
Parses JSON into a Go value.
```go
var user User
json.Unmarshal([]byte(`{"Name":"Bob","Age":25}`), &user)
fmt.Println(user.Name) // Bob
```

### `json.NewDecoder`
```go
func NewDecoder(r io.Reader) *Decoder
```
Creates a decoder to read JSON from a stream.
```go
r := strings.NewReader(`{"Name":"Eve"}`)
var user User
json.NewDecoder(r).Decode(&user)
fmt.Println(user.Name) // Eve
```

---

## 📦 net/http

### `http.Get`
```go
func Get(url string) (resp *Response, err error)
```
Performs an HTTP GET request.
```go
resp, _ := http.Get("https://example.com")
defer resp.Body.Close()
body, _ := io.ReadAll(resp.Body)
fmt.Println(string(body))
```

### `http.Post`
```go
func Post(url, contentType string, body io.Reader) (resp *Response, err error)
```
Performs an HTTP POST request.
```go
resp, _ := http.Post("https://example.com", "application/json", strings.NewReader(`{"k":"v"}`))
defer resp.Body.Close()
```

### `http.ListenAndServe`
```go
func ListenAndServe(addr string, handler Handler) error
```
Starts an HTTP server.
```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
})
http.ListenAndServe(":8080", nil)
```

### `http.ResponseWriter`
Interface used to send responses.
```go
http.HandleFunc("/json", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"ok"}`))
})
```

---

## 📦 gRPC

### `grpc.NewServer`
```go
func NewServer(opts ...ServerOption) *Server
```
Creates a new gRPC server.
```go
s := grpc.NewServer()
```

### `grpc.Dial`
```go
func Dial(target string, opts ...DialOption) (*ClientConn, error)
```
Connects to a gRPC server.
```go
conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
if err != nil { panic(err) }
defer conn.Close()
```

---

## ✅ Summary

- **strings**: Text manipulation.
- **os / io**: Modern file and I/O operations.
- **bytes.Buffer**: Efficient byte buffer.
- **context**: Control cancellation and timeouts.
- **encoding/json**: JSON processing.
- **net/http**: Robust HTTP clients and servers.
- **grpc**: High-performance RPC.
