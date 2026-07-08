# Go: Control Flow

Go deliberately keeps control flow minimal: one loop keyword (`for`), no
ternary operator, no exceptions, and `switch`/`select` statements that are
more flexible than most C-family languages. This sheet covers all of it.

## `if` Statements

```go
x := 15

if x > 10 {
    fmt.Println("big")
} else if x > 5 {
    fmt.Println("medium")
} else {
    fmt.Println("small")
}
```

Braces are mandatory in Go — there is no single-statement-without-braces
form, which eliminates an entire class of dangling-else bugs.

## `if` with an Initialization Statement

```go
if err := doSomething(); err != nil {
    fmt.Println("error:", err)
}

if val, ok := myMap["key"]; ok {
    fmt.Println("found:", val)
}
```

The variable declared in the `if` statement's init clause (`err`, `val`,
`ok` above) is scoped only to the `if`/`else` block — it doesn't leak into
the surrounding function. This pattern is idiomatic Go, especially paired
with error checks and map/type lookups.

## No Ternary Operator

Go has no `condition ? a : b`. The idiomatic replacements are a full `if`,
or occasionally a small helper function.

```go
var label string
if x > 10 {
    label = "big"
} else {
    label = "small"
}

// Or, for simple cases, a tiny generic helper (Go 1.18+):
func ternary[T any](cond bool, a, b T) T {
    if cond {
        return a
    }
    return b
}
label = ternary(x > 10, "big", "small")
```

## `for`: Go's Only Loop Keyword

Go has no `while`, `do-while`, or `until` — `for` covers every loop shape.

```go
// Traditional three-clause for loop
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// "while" loop: condition only
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}

// Infinite loop
for {
    fmt.Println("forever")
    break // needed to actually stop
}

// for-range over a slice
nums := []int{10, 20, 30}
for index, value := range nums {
    fmt.Println(index, value)
}

// for-range, index only
for index := range nums {
    fmt.Println(index)
}

// for-range, value only (discard index with _)
for _, value := range nums {
    fmt.Println(value)
}

// for-range over a map
m := map[string]int{"a": 1, "b": 2}
for key, value := range m {
    fmt.Println(key, value)
}

// for-range over a string (iterates runes, not bytes)
for i, r := range "héllo" {
    fmt.Printf("%d: %c\n", i, r)
}

// for-range over a channel (loops until channel is closed)
ch := make(chan int)
for value := range ch {
    fmt.Println(value)
}

// for-range over an integer (Go 1.22+): loop N times
for i := range 5 {
    fmt.Println(i) // 0,1,2,3,4
}
```

## `break` and `continue`

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break // exit the loop entirely
    }
    if i%2 == 0 {
        continue // skip to the next iteration
    }
    fmt.Println(i)
}
```

## Labeled `break` and `continue` (Escaping Nested Loops)

```go
outer:
for i := 0; i < 5; i++ {
    for j := 0; j < 5; j++ {
        if j == 3 {
            continue outer // continue the OUTER loop, not the inner one
        }
        if i == 3 {
            break outer // break out of BOTH loops
        }
        fmt.Println(i, j)
    }
}
```

Labels are the only way in Go to break or continue an outer loop from
inside a nested one — there's no equivalent of a labeled `goto`-free
multi-level break in languages that lack this feature.

## `switch`: No Fallthrough by Default

Unlike C, Java, or JavaScript, Go's `switch` cases do **not** fall through
to the next case automatically — each case implicitly breaks.

```go
switch x {
case 1:
    fmt.Println("one")
case 2, 3:
    fmt.Println("two or three") // comma-separated values in one case
case 4, 5, 6:
    fmt.Println("four, five, or six")
default:
    fmt.Println("other")
}
```

## Explicit `fallthrough`

```go
switch x {
case 1:
    fmt.Println("one")
    fallthrough // explicitly fall into the next case
case 2:
    fmt.Println("also runs when x == 1")
default:
    fmt.Println("other")
}
```

`fallthrough` must be the last statement in a case, and it unconditionally
runs the next case's body regardless of whether that case's own condition
would have matched.

## `switch` with No Expression (Replaces Long `if`/`else if` Chains)

```go
switch {
case x < 0:
    fmt.Println("negative")
case x == 0:
    fmt.Println("zero")
case x < 10:
    fmt.Println("single digit")
default:
    fmt.Println("large")
}
```

A bare `switch {}` with boolean cases is idiomatic Go for what other
languages would write as a long `if`/`else if`/`else` chain.

## `switch` with an Initialization Statement

```go
switch x := computeValue(); {
case x > 10:
    fmt.Println("big")
default:
    fmt.Println("small")
}

switch day := time.Now().Weekday(); day {
case time.Saturday, time.Sunday:
    fmt.Println("weekend")
default:
    fmt.Println("weekday")
}
```

## Type Switch (Runtime Type Dispatch on `interface{}`/`any`)

```go
func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("int: %d\n", v)
    case string:
        fmt.Printf("string: %s\n", v)
    case bool:
        fmt.Printf("bool: %v\n", v)
    case []int:
        fmt.Printf("slice of int, len %d\n", len(v))
    case nil:
        fmt.Println("nil value")
    default:
        fmt.Printf("unknown type: %T\n", v)
    }
}
```

A type switch is the idiomatic way to branch on the concrete type stored
in an `interface{}`/`any` value — each case gives `v` the matched concrete
type inside that branch.

## `goto`

```go
i := 0
loop:
if i < 5 {
    fmt.Println(i)
    i++
    goto loop
}
```

`goto` exists but is rarely used; it cannot jump into a block or over a
variable declaration that's still in scope at the target. It occasionally
appears in generated code or for breaking out of deeply nested structures
where labeled break doesn't quite fit.

## `defer`: Deferred Execution, Not Really "Control Flow" but Related

```go
func readFile() {
    f, err := os.Open("file.txt")
    if err != nil {
        return
    }
    defer f.Close() // runs when readFile returns, regardless of how it returns

    // ... use f ...
}
```

`defer` statements run in LIFO order when the enclosing function returns —
useful for guaranteed cleanup that "control flow" (early returns, panics)
shouldn't be able to skip.

```go
func multiDefer() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
    // prints: 3, 2, 1
}
```

## `panic` and `recover`: Go's Exception-Like Mechanism

```go
func safeDivide(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered from panic: %v", r)
        }
    }()

    result = a / b // panics on division by zero
    return
}

func main() {
    result, err := safeDivide(10, 0)
    if err != nil {
        fmt.Println("error:", err)
    } else {
        fmt.Println("result:", result)
    }
}
```

`panic` unwinds the stack, running deferred functions along the way, until
something calls `recover()` inside a deferred function — or the program
crashes if nothing does. This is reserved for truly exceptional,
unrecoverable-by-normal-means situations; ordinary error handling in Go
uses returned `error` values, not panics.

## Error Handling as Control Flow (The Idiomatic Go Pattern)

```go
func process() error {
    result, err := step1()
    if err != nil {
        return fmt.Errorf("step1 failed: %w", err)
    }

    result2, err := step2(result)
    if err != nil {
        return fmt.Errorf("step2 failed: %w", err)
    }

    fmt.Println(result2)
    return nil
}
```

The repeated `if err != nil { return ... }` pattern is Go's primary
"control flow for failure" mechanism — it stands in for exceptions and is
considered idiomatic rather than boilerplate to be avoided.

## `select`: Control Flow Over Channels

```go
ch1 := make(chan string)
ch2 := make(chan string)

select {
case msg1 := <-ch1:
    fmt.Println("received from ch1:", msg1)
case msg2 := <-ch2:
    fmt.Println("received from ch2:", msg2)
case <-time.After(1 * time.Second):
    fmt.Println("timeout")
default:
    fmt.Println("no channel ready, non-blocking")
}
```

`select` waits on multiple channel operations simultaneously and proceeds
with whichever is ready first; if multiple are ready, one is chosen at
random. A `default` case makes the whole `select` non-blocking. This is
Go's control-flow primitive for concurrency, with no direct equivalent in
non-CSP languages.

## `select` in a Loop (Common Concurrency Pattern)

```go
for {
    select {
    case msg := <-messages:
        fmt.Println("got message:", msg)
    case <-done:
        fmt.Println("done signal received, exiting")
        return
    }
}
```

## Early Return Pattern (Guard Clauses)

```go
func validate(age int) error {
    if age < 0 {
        return errors.New("age cannot be negative")
    }
    if age > 150 {
        return errors.New("age unrealistically high")
    }
    // main logic, unindented, after guard clauses
    return nil
}
```

Go strongly favors early returns over nested `if`/`else` — guard clauses
that handle error/edge cases first keep the main logic path flat and
readable.

## Quick Reference

| Construct | Purpose |
|---|---|
| `if` / `else if` / `else` | Conditional branching; braces always required |
| `for` | The only loop keyword; covers while, do-while-ish, infinite, range |
| `range` | Iterate over slices, arrays, maps, strings, channels, or integers |
| `break` / `continue` | Exit or skip a loop iteration; can be labeled |
| `switch` | Multi-way branch; no fallthrough by default |
| `switch { }` | Bare switch with boolean cases, replaces long if/else chains |
| `switch v := x.(type)` | Type switch, dispatch on concrete interface type |
| `goto` | Rarely used unconditional jump within a function |
| `defer` | Schedule a call to run at function return, LIFO order |
| `panic` / `recover` | Exceptional-case unwind and catch, not for routine errors |
| `select` | Wait on multiple channel operations at once |

## Tips

- Reach for the bare `switch {}` form instead of long `if`/`else if`
  chains — it's idiomatic and often more readable.
- Remember `switch` doesn't fall through by default; add `fallthrough`
  explicitly on the rare occasion you actually want it.
- Use labeled `break`/`continue` instead of a boolean "found" flag when
  escaping nested loops — it keeps intent clearer.
- Prefer returning `error` values over `panic`/`recover` for anything a
  caller could reasonably need to handle; reserve panics for programmer
  bugs and truly unrecoverable states.
- `select` with a `default` case is the standard way to make a channel
  operation non-blocking.
