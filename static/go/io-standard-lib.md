# Go `io` Package — Comprehensive Quick Reference

> Complete reference for the `io` package: all sentinel errors, constants, interfaces, functions, and types with their methods. Covers Go 1.21+.

---

## Table of Contents

1. [Sentinel Errors & Variables](#1-sentinel-errors--variables)
2. [Seeker Constants](#2-seeker-constants)
3. [Core Interfaces](#3-core-interfaces)
   - [Single-Method Primitives](#single-method-primitives)
   - [Byte & Rune Interfaces](#byte--rune-interfaces)
   - [Transfer Interfaces](#transfer-interfaces)
   - [Composite Interfaces](#composite-interfaces)
4. [Functions](#4-functions)
   - [Copy Functions](#copy-functions)
   - [Read Functions](#read-functions)
   - [Write Functions](#write-functions)
   - [Constructor Functions](#constructor-functions)
5. [Types & Their Methods](#5-types--their-methods)
   - [`SectionReader`](#sectionreader)
   - [`LimitedReader`](#limitedreader)
   - [`PipeReader`](#pipereader)
   - [`PipeWriter`](#pipewriter)
   - [`OffsetWriter`](#offsetwriter)

---

## 1. Sentinel Errors & Variables

These package-level variables are used as canonical error and sink values throughout the standard library.

---

### `io.EOF`

Returned by `Read` when no more input is available. Callers should treat `EOF` as a graceful signal — not a failure — when the total number of bytes received is non-zero.

```go
var EOF = errors.New("EOF")
```

```go
buf := make([]byte, 512)
for {
    n, err := r.Read(buf)
    process(buf[:n])
    if err == io.EOF {
        break // Clean end of stream
    }
    if err != nil {
        return err // Genuine error
    }
}
```

---

### `io.ErrUnexpectedEOF`

Returned when `EOF` is encountered mid-read while filling a fixed-size block or data structure. Signals truncation.

```go
var ErrUnexpectedEOF = errors.New("unexpected EOF")
```

---

### `io.ErrShortWrite`

Returned when a `Write` call accepted fewer bytes than requested but reported no error. Indicates a misbehaving writer.

```go
var ErrShortWrite = errors.New("short write")
```

---

### `io.ErrShortBuffer`

Returned when a read requires a buffer longer than the one supplied.

```go
var ErrShortBuffer = errors.New("short buffer")
```

---

### `io.ErrNoProgress`

Returned by some clients of an `io.Reader` when many calls to `Read` return `n == 0` and `err == nil`, which is a sign of a broken reader implementation.

```go
var ErrNoProgress = errors.New("multiple Read calls return no data or error")
```

---

### `io.ErrClosedPipe`

Returned by read or write operations on a closed pipe.

```go
var ErrClosedPipe = errors.New("io: read/write on closed pipe")
```

---

### `io.Discard`

An `io.Writer` that discards all bytes written to it without error. Useful for draining a reader when its content is irrelevant.

```go
var Discard io.Writer = devNull(0)
```

```go
// Drain a response body to enable connection reuse
_, err := io.Copy(io.Discard, resp.Body)
```

---

## 2. Seeker Constants

Used as the `whence` argument to `Seek` to specify the reference point for the offset.

| Constant | Value | Description |
|---|---|---|
| `io.SeekStart` | `0` | Seek relative to the beginning of the file |
| `io.SeekCurrent` | `1` | Seek relative to the current position |
| `io.SeekEnd` | `2` | Seek relative to the end of the file |

```go
// Rewind to the beginning of a file
_, err := f.Seek(0, io.SeekStart)

// Jump to 10 bytes before the end
_, err := f.Seek(-10, io.SeekEnd)

// Advance the cursor forward 5 bytes from current position
_, err := f.Seek(5, io.SeekCurrent)
```

---

## 3. Core Interfaces

### Single-Method Primitives

These are the foundational building blocks. All composite interfaces and concrete types in the package are built from combinations of these.

---

#### `io.Reader`

The fundamental streaming input interface. `Read` populates `p` with up to `len(p)` bytes and returns the number of bytes read. It returns `io.EOF` when the stream is exhausted.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

> **Contract:** Callers must always process the `n > 0` bytes returned before inspecting the error. A `Read` call may return both data and `io.EOF` simultaneously.

---

#### `io.Writer`

The fundamental streaming output interface. `Write` writes `len(p)` bytes from `p` and returns the number of bytes written. It must return a non-nil error if `n < len(p)`.

```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```

---

#### `io.Closer`

Implemented by any resource that requires explicit release (files, network connections). `Close` is typically deferred immediately after a successful open.

```go
type Closer interface {
    Close() error
}
```

---

#### `io.Seeker`

Allows random-access repositioning within a stream. `offset` is interpreted relative to `whence` — one of `io.SeekStart`, `io.SeekCurrent`, or `io.SeekEnd`.

```go
type Seeker interface {
    Seek(offset int64, whence int) (int64, error)
}
```

---

### Byte & Rune Interfaces

---

#### `io.ByteReader`

Implemented by readers that can read a single byte at a time without buffering overhead. Used by `encoding/binary` and parsers that process byte-by-byte.

```go
type ByteReader interface {
    ReadByte() (byte, error)
}
```

---

#### `io.ByteScanner`

Extends `ByteReader` with the ability to "unread" the last byte, pushing it back onto the stream for re-reading.

```go
type ByteScanner interface {
    ByteReader
    UnreadByte() error
}
```

---

#### `io.ByteWriter`

Implemented by writers that can write a single byte efficiently.

```go
type ByteWriter interface {
    WriteByte(c byte) error
}
```

---

#### `io.RuneReader`

Implemented by readers that can decode and return a single UTF-8 rune and its byte width. Used by `strings.Reader` and `bytes.Reader`.

```go
type RuneReader interface {
    ReadRune() (r rune, size int, err error)
}
```

---

#### `io.RuneScanner`

Extends `RuneReader` with the ability to push the last decoded rune back onto the stream.

```go
type RuneScanner interface {
    RuneReader
    UnreadRune() error
}
```

---

#### `io.StringWriter`

Implemented by writers that can accept a string directly without first converting it to `[]byte`. Avoids an allocation.

```go
type StringWriter interface {
    WriteString(s string) (n int, err error)
}
```

---

### Transfer Interfaces

---

#### `io.ReaderAt`

Reads from an exact offset in the underlying source without affecting or depending on the source's current seek position. Multiple goroutines may call `ReadAt` on the same source concurrently.

```go
type ReaderAt interface {
    ReadAt(p []byte, off int64) (n int, err error)
}
```

---

#### `io.WriterAt`

Writes data to an exact offset in the underlying sink without affecting the current seek position.

```go
type WriterAt interface {
    WriteAt(p []byte, off int64) (n int, err error)
}
```

---

#### `io.ReaderFrom`

Implemented by writers that know how to read directly from a `Reader`. When the destination implements `ReaderFrom`, `io.Copy` will call `ReadFrom` instead of allocating an intermediate buffer — enabling zero-copy transfers (e.g., `sendfile(2)`).

```go
type ReaderFrom interface {
    ReadFrom(r Reader) (n int64, err error)
}
```

---

#### `io.WriterTo`

Implemented by readers that know how to write directly to a `Writer`. When the source implements `WriterTo`, `io.Copy` will call `WriteTo` instead of looping with an intermediate buffer.

```go
type WriterTo interface {
    WriteTo(w Writer) (n int64, err error)
}
```

---

### Composite Interfaces

These are pre-declared interface combinations. Use them in function signatures to express the exact capability required.

| Interface | Embeds |
|---|---|
| `io.ReadWriter` | `Reader`, `Writer` |
| `io.ReadCloser` | `Reader`, `Closer` |
| `io.WriteCloser` | `Writer`, `Closer` |
| `io.ReadWriteCloser` | `Reader`, `Writer`, `Closer` |
| `io.ReadSeeker` | `Reader`, `Seeker` |
| `io.WriteSeeker` | `Writer`, `Seeker` |
| `io.ReadWriteSeeker` | `Reader`, `Writer`, `Seeker` |

```go
// Accepting ReadSeeker allows both streaming reads and rewinding
func parseReplayable(src io.ReadSeeker) error {
    // First pass: scan
    if err := scan(src); err != nil {
        return err
    }
    // Rewind
    if _, err := src.Seek(0, io.SeekStart); err != nil {
        return err
    }
    // Second pass: parse
    return parse(src)
}
```

---

## 4. Functions

### Copy Functions

---

#### `io.Copy`

Copies from `src` to `dst` until `src` returns `EOF` or an error. Uses a 32 KiB internal buffer unless `src` implements `WriterTo` or `dst` implements `ReaderFrom`, in which case the transfer is delegated to avoid the allocation.

```go
func Copy(dst Writer, src Reader) (written int64, err error)
```

```go
// Stream a file to an HTTP response body
written, err := io.Copy(w, file)
```

---

#### `io.CopyN`

Copies exactly `n` bytes from `src` to `dst`. Returns `ErrUnexpectedEOF` if `src` is exhausted before `n` bytes are read.

```go
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
```

```go
// Copy just the first 4 bytes (e.g., a magic number)
_, err := io.CopyN(buf, f, 4)
```

---

#### `io.CopyBuffer`

Identical to `io.Copy` but uses the caller-supplied `buf` as the intermediate transfer buffer, giving control over allocation. If `buf` is `nil`, an internal buffer is allocated. Ignored if `src` or `dst` implement `WriterTo`/`ReaderFrom`.

```go
func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)
```

```go
buf := make([]byte, 64*1024) // 64 KiB buffer
_, err := io.CopyBuffer(dst, src, buf)
```

---

### Read Functions

---

#### `io.ReadAll`

Reads from `r` until `EOF` and returns the accumulated data. Returns an error if `Read` fails before `EOF`. Suitable for small to medium payloads; avoid on unbounded streams.

```go
func ReadAll(r Reader) ([]byte, error)
```

```go
body, err := io.ReadAll(resp.Body)
```

---

#### `io.ReadFull`

Reads exactly `len(buf)` bytes from `r` into `buf`. Returns `ErrUnexpectedEOF` if `EOF` is hit before the buffer is filled, and `EOF` only if no bytes were read at all.

```go
func ReadFull(r Reader, buf []byte) (n int, err error)
```

```go
header := make([]byte, 8)
if _, err := io.ReadFull(conn, header); err != nil {
    return fmt.Errorf("reading header: %w", err)
}
```

---

#### `io.ReadAtLeast`

Reads from `r` into `buf` until at least `min` bytes have been read. Returns `ErrUnexpectedEOF` if `EOF` is reached before `min` bytes, and `ErrShortBuffer` if `min > len(buf)`.

```go
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)
```

```go
buf := make([]byte, 512)
n, err := io.ReadAtLeast(r, buf, 64) // require at least 64 bytes
```

---

### Write Functions

---

#### `io.WriteString`

Writes the string `s` to `w`. If `w` implements `io.StringWriter`, it calls `WriteString` directly to avoid converting the string to `[]byte`. Falls back to `w.Write([]byte(s))` otherwise.

```go
func WriteString(w Writer, s string) (n int, err error)
```

```go
n, err := io.WriteString(w, "HTTP/1.1 200 OK\r\n")
```

---

### Constructor Functions

---

#### `io.LimitReader`

Returns a `Reader` that reads from `r` but stops with `EOF` after `n` bytes. Useful for enforcing maximum read sizes on untrusted input.

```go
func LimitReader(r Reader, n int64) Reader
```

```go
// Reject request bodies larger than 1 MiB
limited := io.LimitReader(req.Body, 1<<20)
data, err := io.ReadAll(limited)
```

---

#### `io.MultiReader`

Returns a `Reader` that is the logical concatenation of the provided readers. Reads sequentially through each reader until all are exhausted.

```go
func MultiReader(readers ...Reader) Reader
```

```go
// Prepend a header to a stream without copying
header := strings.NewReader("MAGIC\n")
combined := io.MultiReader(header, body)
```

---

#### `io.MultiWriter`

Returns a `Writer` that fans out each `Write` call to all provided writers simultaneously. If any writer returns an error, `MultiWriter` stops and returns that error.

```go
func MultiWriter(writers ...Writer) Writer
```

```go
// Write to both a file and stdout at the same time
mw := io.MultiWriter(os.Stdout, logFile)
fmt.Fprintf(mw, "Initializing server on :%d\n", port)
```

---

#### `io.TeeReader`

Returns a `Reader` that reads from `r` and simultaneously writes every byte read into `w`. There is no internal buffering — the write must complete before the read returns.

```go
func TeeReader(r Reader, w Writer) Reader
```

```go
// Inspect request body while still forwarding it
var buf bytes.Buffer
tee := io.TeeReader(req.Body, &buf)
io.Copy(upstream, tee)
// buf now contains a copy of everything that was read
```

---

#### `io.NopCloser`

Wraps an `io.Reader` in an `io.ReadCloser` whose `Close` method is a no-op. Used to satisfy interfaces that require a `ReadCloser` when you only have a `Reader` without a meaningful close operation.

```go
func NopCloser(r Reader) ReadCloser
```

> As of Go 1.16, if `r` also implements `WriterTo`, the returned `ReadCloser` will also implement `WriterTo`.

```go
// Wrap an in-memory buffer for use where ReadCloser is required
rc := io.NopCloser(strings.NewReader("synthetic body"))
```

---

#### `io.Pipe`

Creates a synchronous, in-memory pipe connecting a `PipeWriter` to a `PipeReader`. Writes on the writer block until one or more reads have consumed all data or the reader is closed. There is no internal buffering.

```go
func Pipe() (*PipeReader, *PipeWriter)
```

```go
pr, pw := io.Pipe()

go func() {
    defer pw.Close()
    json.NewEncoder(pw).Encode(payload)
}()

resp, err := http.Post(url, "application/json", pr)
```

---

#### `io.NewSectionReader`

Returns a `SectionReader` that reads from `r` starting at offset `off` and stops with `EOF` after `n` bytes.

```go
func NewSectionReader(r ReaderAt, off int64, n int64) *SectionReader
```

```go
// Read bytes 512–1023 of a file without seeking
section := io.NewSectionReader(f, 512, 512)
data, err := io.ReadAll(section)
```

---

#### `io.NewOffsetWriter`

Returns an `OffsetWriter` that writes to `w` starting at offset `off`, adjusting each `WriteAt` call's offset transparently.

```go
func NewOffsetWriter(w WriterAt, off int64) *OffsetWriter
```

---

## 5. Types & Their Methods

### `SectionReader`

Presents a sub-region of an `io.ReaderAt` as a self-contained `Reader`, `ReaderAt`, and `Seeker`. Seeking is bounded within the section; positions beyond `n` bytes return `EOF`.

```go
type SectionReader struct { /* unexported */ }
```

#### `(*SectionReader).Read`

Reads from the section into `p`, respecting the section boundary. Returns `io.EOF` when the section is exhausted.

```go
func (s *SectionReader) Read(p []byte) (n int, err error)
```

---

#### `(*SectionReader).ReadAt`

Reads `len(p)` bytes from the section starting at the given offset within the section (not the underlying source). Concurrent calls are safe.

```go
func (s *SectionReader) ReadAt(p []byte, off int64) (n int, err error)
```

---

#### `(*SectionReader).Seek`

Sets the read position within the section. `whence` values (`SeekStart`, `SeekCurrent`, `SeekEnd`) are relative to the section, not the underlying source. Seeking before the start returns an error.

```go
func (s *SectionReader) Seek(offset int64, whence int) (int64, error)
```

---

#### `(*SectionReader).Size`

Returns the total size of the section in bytes.

```go
func (s *SectionReader) Size() int64
```

```go
section := io.NewSectionReader(f, 1024, 4096)
fmt.Println(section.Size()) // 4096
```

---

#### `(*SectionReader).Outer`

Returns the underlying `ReaderAt`, the starting offset, and the total byte limit that the `SectionReader` was created with. Added in Go 1.22.

```go
func (s *SectionReader) Outer() (r ReaderAt, off int64, n int64)
```

---

### `LimitedReader`

A concrete struct (returned by `LimitReader`) that reads from `R` but limits the total bytes to `N`. Once `N` reaches zero, subsequent reads return `EOF`. Inspecting `N` during or after a read reveals how many bytes remain.

```go
type LimitedReader struct {
    R Reader // underlying reader
    N int64  // max bytes remaining
}
```

#### `(*LimitedReader).Read`

Reads up to `min(len(p), N)` bytes from the underlying reader. Decrements `N` by the number of bytes read.

```go
func (l *LimitedReader) Read(p []byte) (n int, err error)
```

```go
lr := &io.LimitedReader{R: r, N: 256}
io.ReadAll(lr)
fmt.Println(lr.N) // 0 if exactly 256 bytes were available
```

---

### `PipeReader`

The read end of a pipe created by `io.Pipe`. Reads block until the writer writes or closes the pipe.

```go
type PipeReader struct { /* unexported */ }
```

#### `(*PipeReader).Read`

Reads data written by the paired `PipeWriter`. Blocks until data is available, the writer is closed, or the writer is closed with an error.

```go
func (r *PipeReader) Read(data []byte) (n int, err error)
```

---

#### `(*PipeReader).Close`

Closes the reader with `ErrClosedPipe`. Subsequent writes on the paired `PipeWriter` will fail with `ErrClosedPipe`.

```go
func (r *PipeReader) Close() error
```

---

#### `(*PipeReader).CloseWithError`

Closes the reader with a custom error. Subsequent writes on the paired `PipeWriter` will fail with that error. If `err` is `nil`, `ErrClosedPipe` is used.

```go
func (r *PipeReader) CloseWithError(err error) error
```

```go
// Abort the pipe with a domain-specific error
pr.CloseWithError(fmt.Errorf("consumer cancelled: %w", ctx.Err()))
```

---

### `PipeWriter`

The write end of a pipe created by `io.Pipe`. Writes block until the reader has consumed the data.

```go
type PipeWriter struct { /* unexported */ }
```

#### `(*PipeWriter).Write`

Writes data to the pipe. Blocks until the reader has consumed all written bytes or the reader is closed. Returns `ErrClosedPipe` if the reader is already closed.

```go
func (w *PipeWriter) Write(data []byte) (n int, err error)
```

---

#### `(*PipeWriter).Close`

Closes the writer with `EOF`. Subsequent reads on the paired `PipeReader` will return `EOF` once buffered data is consumed.

```go
func (w *PipeWriter) Close() error
```

---

#### `(*PipeWriter).CloseWithError`

Closes the writer with a custom error. Subsequent reads on the paired `PipeReader` will return that error. If `err` is `nil`, `EOF` is used.

```go
func (w *PipeWriter) CloseWithError(err error) error
```

```go
pr, pw := io.Pipe()

go func() {
    _, err := fetchData(pw)
    pw.CloseWithError(err) // Propagate upstream errors to the reader
}()
```

---

### `OffsetWriter`

Wraps an `io.WriterAt` and shifts all write operations by a fixed base offset. Added in Go 1.20. Useful for writing into a sub-region of a file or buffer without manually tracking offsets at the call site.

```go
type OffsetWriter struct { /* unexported */ }
```

#### `(*OffsetWriter).Write`

Writes `p` at the current offset (starting from the base offset), then advances the internal cursor.

```go
func (o *OffsetWriter) Write(p []byte) (n int, err error)
```

---

#### `(*OffsetWriter).WriteAt`

Writes `p` at `base + off` in the underlying writer, without advancing the internal cursor.

```go
func (o *OffsetWriter) WriteAt(p []byte, off int64) (n int, err error)
```

---

#### `(*OffsetWriter).Seek`

Moves the internal cursor relative to the base offset. `whence` values follow the standard `io.Seek*` constants.

```go
func (o *OffsetWriter) Seek(offset int64, whence int) (int64, error)
```

```go
// Write into byte 4096–8191 of a file, using zero-based offsets in code
ow := io.NewOffsetWriter(f, 4096)
ow.Write(chunk) // actually writes at file offset 4096
```

---

## Quick Index

| Symbol | Kind | Summary |
|---|---|---|
| `io.EOF` | var | End-of-stream sentinel |
| `io.ErrUnexpectedEOF` | var | EOF encountered mid-block |
| `io.ErrShortWrite` | var | Write accepted fewer bytes than requested |
| `io.ErrShortBuffer` | var | Buffer too small for read |
| `io.ErrNoProgress` | var | Reader returning `0, nil` repeatedly |
| `io.ErrClosedPipe` | var | R/W on a closed pipe |
| `io.Discard` | var | Sink that discards all writes |
| `io.SeekStart` | const | Seek from beginning |
| `io.SeekCurrent` | const | Seek from current position |
| `io.SeekEnd` | const | Seek from end |
| `io.Reader` | interface | `Read([]byte) (int, error)` |
| `io.Writer` | interface | `Write([]byte) (int, error)` |
| `io.Closer` | interface | `Close() error` |
| `io.Seeker` | interface | `Seek(int64, int) (int64, error)` |
| `io.ByteReader` | interface | `ReadByte() (byte, error)` |
| `io.ByteScanner` | interface | `ByteReader` + `UnreadByte` |
| `io.ByteWriter` | interface | `WriteByte(byte) error` |
| `io.RuneReader` | interface | `ReadRune() (rune, int, error)` |
| `io.RuneScanner` | interface | `RuneReader` + `UnreadRune` |
| `io.StringWriter` | interface | `WriteString(string) (int, error)` |
| `io.ReaderAt` | interface | `ReadAt([]byte, int64) (int, error)` |
| `io.WriterAt` | interface | `WriteAt([]byte, int64) (int, error)` |
| `io.ReaderFrom` | interface | `ReadFrom(Reader) (int64, error)` |
| `io.WriterTo` | interface | `WriteTo(Writer) (int64, error)` |
| `io.ReadWriter` | interface | `Reader` + `Writer` |
| `io.ReadCloser` | interface | `Reader` + `Closer` |
| `io.WriteCloser` | interface | `Writer` + `Closer` |
| `io.ReadWriteCloser` | interface | `Reader` + `Writer` + `Closer` |
| `io.ReadSeeker` | interface | `Reader` + `Seeker` |
| `io.WriteSeeker` | interface | `Writer` + `Seeker` |
| `io.ReadWriteSeeker` | interface | `Reader` + `Writer` + `Seeker` |
| `io.Copy` | func | Stream copy with zero-copy fast path |
| `io.CopyN` | func | Copy exactly N bytes |
| `io.CopyBuffer` | func | Copy with caller-supplied buffer |
| `io.ReadAll` | func | Read entire stream into `[]byte` |
| `io.ReadFull` | func | Read exactly `len(buf)` bytes |
| `io.ReadAtLeast` | func | Read at least `min` bytes |
| `io.WriteString` | func | Write string, avoiding `[]byte` allocation |
| `io.LimitReader` | func | Cap a reader at N bytes |
| `io.MultiReader` | func | Concatenate readers sequentially |
| `io.MultiWriter` | func | Fan out writes to multiple writers |
| `io.TeeReader` | func | Read-and-tap to a secondary writer |
| `io.NopCloser` | func | Wrap `Reader` as `ReadCloser` |
| `io.Pipe` | func | Synchronous in-memory pipe |
| `io.NewSectionReader` | func | Sub-region view of a `ReaderAt` |
| `io.NewOffsetWriter` | func | Offset-shifted `WriterAt` wrapper |
| `(*SectionReader).Read` | method | Read within section bounds |
| `(*SectionReader).ReadAt` | method | Read at section-relative offset |
| `(*SectionReader).Seek` | method | Seek within section |
| `(*SectionReader).Size` | method | Section byte count |
| `(*SectionReader).Outer` | method | Recover underlying source and bounds |
| `LimitedReader.R` | field | Underlying reader |
| `LimitedReader.N` | field | Bytes remaining before EOF |
| `(*LimitedReader).Read` | method | Read, decrementing N |
| `(*PipeReader).Read` | method | Blocking read from pipe |
| `(*PipeReader).Close` | method | Close reader with `ErrClosedPipe` |
| `(*PipeReader).CloseWithError` | method | Close reader with custom error |
| `(*PipeWriter).Write` | method | Blocking write to pipe |
| `(*PipeWriter).Close` | method | Signal EOF to reader |
| `(*PipeWriter).CloseWithError` | method | Signal custom error to reader |
| `(*OffsetWriter).Write` | method | Sequential write from base offset |
| `(*OffsetWriter).WriteAt` | method | Write at `base + off` |
| `(*OffsetWriter).Seek` | method | Move cursor within offset writer |
