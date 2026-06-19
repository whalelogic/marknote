# The Pragmatic Rust Syntax & Keyword Reference

## Part 1: Symbols and Operators

| Symbol | Name | Description / Use Case | Example |
| :--- | :--- | :--- | :--- |
| `!` | Macro / Not | 1. Logic NOT. 2. Suffix for macro invocation. 3. "Never" return type. | `println!()`, `!true`, `-> !` |
| `?` | Try Operator | Unwraps `Result` or `Option`. Returns error/None to caller on failure. | `let f = File::open("a.txt")?;` |
| `.` | Member Access | Accesses fields of a struct or methods of a trait/impl. | `person.name`, `vec.len()` |
| `..` | Range (Exclusive) | Creates a range from start to up-to-but-not-including end. | `1..10` (1 to 9) |
| `..=` | Range (Inclusive) | Creates a range including the end value. | `1..=10` (1 to 10) |
| `..` | Functional Update | In struct literals, copies remaining fields from another instance. | `User { email, ..user1 }` |
| `::` | Path Separator | Accesses items in modules, namespaces, or associated functions. | `std::io::stdin()`, `Vec::new()` |
| `::<_>` | Turbofish | Provides explicit type arguments to a generic function/method. | `collect::<Vec<i32>>()` |
| `&` | Reference / And | 1. Bitwise AND. 2. Create an immutable reference (borrow). | `let x = &y;`, `5 & 1` |
| `&mut` | Mutable Borrow | Creates a unique mutable reference to a value. | `let x = &mut y;` |
| `*` | Dereference / Mult | 1. Multiplication. 2. Follow a pointer/reference to its value. | `let val = *reference;` |
| `_` | Wildcard / Ignore | 1. Pattern match catch-all. 2. Ignore variable binding. | `let _ = hide_me();`, `match { _ => ... }` |
| `,` | Separator | Separates arguments, elements in arrays/vecs, or struct fields. | `fn(a: i32, b: i32)` |
| `;` | Terminator | Ends a statement. If omitted at end of function, returns expression. | `let x = 5;` |
| `:` | Type Ascription | Associates a name with a specific type. | `let x: i32 = 5;` |
| `->` | Return Arrow | Denotes the return type of a function or closure. | `fn add() -> i32` |
| `=>` | Match Arm | Separates a match pattern from the resulting expression. | `Some(x) => println!("{x}")` |
| `&'a` | Lifetime | Explicitly labels how long a reference must remain valid. | `&'a str`, `fn foo<'a>(...)` |
| `<T>` | Generics | Declares or uses a generic type parameter `T`. | `struct Box<T> { ... }` |
| `{:?}` | Debug Format | Format specifier for the `Debug` trait (printing internals). | `println!("{:?}", vec);` |
| `{:#?}` | Pretty Debug | Format specifier for multi-line, indented debug printing. | `println!("{:#?}", complex_struct);` |
| `\|...\|` | Closure | Defines parameters for an anonymous function (lambda). | `let c = \|x, y\| x + y;` |
| `[...]` | Array / Slice | Declares array types or accesses indices. | `let a: [i32; 3]`, `vec[0]` |
| `(...)` | Tuple / Group | 1. Group expressions. 2. Declare tuples. 3. Call functions. | `let t = (1, "hi");` |
| `{...}` | Block / Scope | Defines a scope or a struct/enum body. Expressions return last value. | `let x = { 5 };` |
| `@` | Binding Pattern | Binds a matched value to a name while also testing the pattern. | `n @ 1..=9 => ...` |
| `#[...]` | Attribute (Outer) | Annotates the item that follows (derive, cfg, inline, etc.). | `#[derive(Debug, Clone)]` |
| `#![...]` | Attribute (Inner) | Annotates the enclosing item, typically the crate root. | `#![allow(dead_code)]` |
| `<<` / `>>` | Bit Shift | Shifts bits left or right. | `1 << 3` (equals 8) |
| `^` | XOR | Bitwise exclusive OR. | `0b1010 ^ 0b1100` |
| `\|` | Bitwise OR / Alt Pattern | 1. Bitwise OR. 2. Combine patterns in `match`. | `0b1010 \| 0b0101`, `1 \| 2 => ...` |
| `as_ref()` | Borrow as Ref | Converts `Option<T>` → `Option<&T>` without consuming. | `opt.as_ref().map(\|v\| v.len())` |
| `..` | Rest Pattern | Ignores remaining fields/elements in a destructure. | `let (x, ..) = (1, 2, 3);` |

## Part 2: Keywords and Reserved Terms

| Keyword | Definition | Practical Usage |
| :--- | :--- | :--- |
| `as` | Casting | Perform primitive casting or rename imports. `x as u64`, `use std::io as bio;` |
| `async` | Async Function | Marks a function or block as asynchronous; returns a `Future`. `async fn fetch() -> Result<...>` |
| `await` | Await Future | Suspends execution until a `Future` resolves. `let data = fetch().await?;` |
| `break` | Exit Loop | Break out of a loop (can also return a value from `loop`). `break 5;` |
| `const` | Constant | Compile-time constant with a mandatory type. `const MAX: u32 = 100;` |
| `continue` | Skip | Skip the rest of the current loop iteration. |
| `crate` | Root Module | Refers to the root of the current crate. |
| `dyn` | Dynamic Dispatch | Indicates a trait object (vtable lookup at runtime). `&dyn MyTrait` |
| `else` | Fallback | Defines the fallback block for an `if` expression. |
| `enum` | Enumeration | Defines a type that can be one of several variants. |
| `extern` | FFI | Links to external libraries or defines C-calling conventions. |
| `fn` | Function | Declares a function or a function pointer type. |
| `for` | Loop | Iterates over an iterator. `for x in 0..5 { ... }` |
| `if` | Conditional | Branching expression. Returns a value if used as an expression. |
| `if let` | Conditional Bind | Combines `if` and pattern matching. Avoids a full `match` for one variant. `if let Some(v) = opt { ... }` |
| `impl` | Implementation | Implements methods or traits for a struct or enum. |
| `impl Trait` | Return Position | Hides a concrete return type behind a trait bound (opaque type). `fn make() -> impl Iterator<Item=i32>` |
| `in` | Part of `for` | Used in `for` loops to specify the collection/range. |
| `let` | Binding | Binds a value to a variable name. |
| `let else` | Binding / Diverge | Like `if let`, but the `else` branch must diverge (return, break, panic). `let Ok(v) = res else { return; };` |
| `loop` | Infinite Loop | Unconditional loop. Useful for retries or long-running tasks. |
| `match` | Pattern Match | Exhaustive branching based on pattern matching. |
| `mod` | Module | Declares or defines a module. |
| `move` | Capture by Value | Forces a closure to take ownership of captured variables. |
| `mut` | Mutable | Allows a variable or reference to be modified. |
| `pub` | Public | Makes a module, struct, or function visible outside its parent. |
| `pub(crate)` | Crate-Visible | Restricts visibility to within the current crate. |
| `pub(super)` | Parent-Visible | Restricts visibility to the parent module. |
| `ref` | By Reference | Used in patterns to bind by reference rather than moving. |
| `return` | Return | Returns a value from a function early. |
| `Self` | Self Type | The type being implemented (inside an `impl` block). |
| `self` | Self Instance | The instance being operated on (method receiver). |
| `static` | Global | Global variable with a `'static` lifetime. |
| `struct` | Structure | Defines a custom data type with fields. |
| `trait` | Interface | Defines shared behavior that types can implement. |
| `type` | Alias | Creates a new name for an existing type. `type Bytes = Vec<u8>;` |
| `unsafe` | Unsafe | Opts into features the compiler can't verify (raw pointers, etc). |
| `use` | Import | Brings a path into the current scope. |
| `where` | Constraints | Defines complex generic bounds at the end of a signature. |
| `while` | Loop while | Conditional loop. `while x < 5 { ... }` |
| `while let` | Loop / Bind | Loops as long as a pattern matches. `while let Some(v) = iter.next() { ... }` |

## Part 3: Common Standard Types & Enums

| Term | Context | Description |
| :--- | :--- | :--- |
| `Option<T>` | Enum | Represents a value that might be missing (`Some(T)` or `None`). |
| `Result<T, E>` | Enum | Represents success (`Ok(T)`) or failure (`Err(E)`). |
| `Some(v)` | Variant | Part of `Option`. Wraps an existing value. |
| `None` | Variant | Part of `Option`. Represents the absence of value. |
| `Ok(v)` | Variant | Part of `Result`. Represents a successful operation. |
| `Err(e)` | Variant | Part of `Result`. Represents a failed operation. |
| `Vec<T>` | Struct | A growable, heap-allocated array (Vector). |
| `String` | Struct | A heap-allocated, UTF-8 encoded string. |
| `&str` | Type | A string slice (reference to string data). |
| `Box<T>` | Struct | A heap-allocated smart pointer for single ownership. |
| `Rc<T>` | Struct | Reference-counted pointer for shared ownership (single-threaded). |
| `Arc<T>` | Struct | Atomically reference-counted pointer for shared ownership across threads. |
| `Cell<T>` | Struct | Interior mutability for `Copy` types (no borrow checker enforcement). |
| `RefCell<T>` | Struct | Interior mutability with runtime borrow checking (single-threaded). |
| `Mutex<T>` | Struct | Mutual exclusion primitive for thread-safe interior mutability. |
| `RwLock<T>` | Struct | Like `Mutex`, but allows multiple concurrent readers. |
| `HashMap<K, V>` | Struct | Hash-based key-value store. Not ordered. |
| `BTreeMap<K, V>` | Struct | Key-value store sorted by key. Useful when ordering matters. |
| `HashSet<T>` | Struct | Unordered collection of unique values. |
| `BTreeSet<T>` | Struct | Ordered collection of unique values. |
| `VecDeque<T>` | Struct | Double-ended queue backed by a ring buffer. |
| `Range<T>` | Struct | The type produced by `start..end`. Implements `Iterator`. |
| `RangeInclusive<T>` | Struct | The type produced by `start..=end`. Implements `Iterator`. |
| `PathBuf` | Struct | Owned, mutable filesystem path. Analogous to `String`. |
| `Path` | Struct | Borrowed filesystem path slice. Analogous to `&str`. |
| `OsString` | Struct | Owned OS-native string (not guaranteed UTF-8). |
| `Duration` | Struct | A span of time (`std::time::Duration`). |
| `Instant` | Struct | A monotonic clock timestamp, used for elapsed time measurement. |

## Part 4: Common Traits

| Trait | Description | Common Use |
| :--- | :--- | :--- |
| `Debug` | Enables `{:?}` formatting. | `#[derive(Debug)]` on nearly every type. |
| `Display` | Enables `{}` formatting (user-facing). | `impl Display for MyType`. |
| `Clone` | Explicit deep copy via `.clone()`. | `#[derive(Clone)]` |
| `Copy` | Implicit bitwise copy (no move semantics). | Scalars, small value types. |
| `PartialEq` / `Eq` | Equality comparisons (`==`, `!=`). | `#[derive(PartialEq, Eq)]` |
| `PartialOrd` / `Ord` | Ordering comparisons (`<`, `>`, `<=`, `>=`). | Required for sorting or use in `BTreeMap`. |
| `Hash` | Enables use as a `HashMap` or `HashSet` key. | `#[derive(Hash)]` |
| `Default` | Provides a zero-value via `T::default()`. | `#[derive(Default)]` |
| `Iterator` | Core iteration protocol. Requires `fn next() -> Option<Self::Item>`. | `impl Iterator for MyIter`. |
| `From<T>` / `Into<T>` | Infallible type conversion. Implementing `From` gives `Into` for free. | `String::from("hello")`, `x.into()` |
| `TryFrom<T>` / `TryInto<T>` | Fallible type conversion returning `Result`. | `u8::try_from(256u32)` → `Err` |
| `AsRef<T>` | Cheap reference-to-reference conversion. Used to write flexible APIs. | `fn open(p: impl AsRef<Path>)` |
| `Deref` | Overloads `*`. Enables coercions (`Box<T>` → `T`, `String` → `&str`). | Implemented by smart pointers. |
| `Drop` | Custom destructor logic run when a value goes out of scope. | `impl Drop for MyType { fn drop(&mut self) { ... } }` |
| `Send` | Marker: safe to transfer ownership across threads. | Auto-derived; `Arc<T>` is `Send`, `Rc<T>` is not. |
| `Sync` | Marker: safe to share a reference across threads. | Auto-derived; `Mutex<T>` is `Sync`. |
| `Fn` / `FnMut` / `FnOnce` | Closure call traits. `FnOnce` consumes, `FnMut` mutates, `Fn` borrows. | Used as trait bounds: `fn run(f: impl Fn())`. |
| `Error` | Base trait for error types. Requires `Debug + Display`. | `impl Error for MyError {}` |
| `Write` | Byte or formatted output (files, sockets, `Vec<u8>`). | `use std::io::Write; file.write_all(b"data")?;` |
| `Read` | Byte input from a source. | `use std::io::Read; file.read_to_string(&mut s)?;` |
| `BufRead` | Buffered line-by-line reading. | `for line in reader.lines() { ... }` |

## Part 5: Practical Syntax Examples

### 1. The Turbofish and Collect

```rust
// split returns an iterator. collect needs to know the destination type.
let items = "apple,banana,cherry"
    .split(',')
    .collect::<Vec<&str>>();
```

### 2. Match with Pattern Guards and Enums

```rust
match number {
    0 => println!("Zero"),
    n if n < 0 => println!("Negative"),
    1..=10 => println!("Small positive"),
    _ => println!("Something else"),
}
```

### 3. Closure with Move

```rust
let data = vec![1, 2, 3];
// 'move' takes ownership of 'data' so it can be used in the closure/thread
let handle = std::thread::spawn(move || {
    println!("{:?}", data);
});
```

### 4. Struct Update Syntax

```rust
let user1 = User { name: "Alice".to_string(), active: true, score: 10 };
let user2 = User {
    name: "Bob".to_string(),
    ..user1 // active and score are copied from user1
};
```

### 5. Error Handling with `?`

```rust
fn get_id() -> Result<u32, MyError> {
    let s = std::fs::read_to_string("id.txt")?; // returns Err early if read fails
    let id = s.trim().parse::<u32>()?;          // returns Err early if parse fails
    Ok(id)
}
```

### 6. `if let` and `while let`

```rust
// if let — handle one variant without a full match
if let Some(value) = map.get("key") {
    println!("Found: {value}");
}

// while let — drain an iterator or channel
while let Some(item) = stack.pop() {
    process(item);
}
```

### 7. `let`-`else` for Early Returns

```rust
fn parse_port(s: &str) -> u16 {
    // The else branch must diverge (return, panic, break, continue, etc.)
    let Ok(n) = s.parse::<u16>() else {
        panic!("invalid port: {s}");
    };
    n
}
```

### 8. Binding Patterns with `@`

```rust
match value {
    n @ 1..=9 => println!("single digit: {n}"),
    n @ 10..=99 => println!("two digits: {n}"),
    _ => println!("three or more"),
}
```

### 9. `impl Trait` in Return Position

```rust
// Returns something that implements Iterator without naming the concrete type
fn evens_up_to(n: u32) -> impl Iterator<Item = u32> {
    (0..=n).filter(|x| x % 2 == 0)
}
```

### 10. Trait Objects with `dyn`

```rust
// Heterogeneous collection of types that share a trait
fn notify(senders: &[Box<dyn Notifier>]) {
    for s in senders {
        s.send("hello");
    }
}
```

### 11. Interior Mutability with `RefCell`

```rust
use std::cell::RefCell;

let data = RefCell::new(vec![1, 2, 3]);
data.borrow_mut().push(4); // runtime borrow check — panics if already borrowed
println!("{:?}", data.borrow());
```

### 12. Shared Ownership with `Arc<Mutex<T>>` (Thread-Safe)

```rust
use std::sync::{Arc, Mutex};
use std::thread;

let counter = Arc::new(Mutex::new(0));
let handles: Vec<_> = (0..4).map(|_| {
    let c = Arc::clone(&counter);
    thread::spawn(move || { *c.lock().unwrap() += 1; })
}).collect();
for h in handles { h.join().unwrap(); }
println!("Result: {}", *counter.lock().unwrap()); // 4
```

### 13. Custom Iterator

```rust
struct Counter { count: u32, max: u32 }

impl Iterator for Counter {
    type Item = u32;
    fn next(&mut self) -> Option<Self::Item> {
        if self.count < self.max {
            self.count += 1;
            Some(self.count)
        } else {
            None
        }
    }
}
```

### 14. Async / Await

```rust
use tokio; // most common async runtime

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let body = reqwest::get("https://example.com")
        .await?
        .text()
        .await?;
    println!("{body}");
    Ok(())
}
```

### 15. Deriving Common Traits

```rust
#[derive(Debug, Clone, PartialEq, Eq, Hash, Default)]
struct Point {
    x: i32,
    y: i32,
}

// Now usable in HashSets, printable with {:?}, cloneable, comparable, etc.
```

### 16. Destructuring Structs, Tuples, and Enums

```rust
struct Point { x: f64, y: f64 }
let p = Point { x: 1.0, y: 2.0 };
let Point { x, y } = p;       // struct destructure

let (a, b, ..) = (1, 2, 3, 4); // tuple destructure, rest ignored

enum Shape { Circle(f64), Rect(f64, f64) }
match shape {
    Shape::Circle(r) => println!("area: {}", std::f64::consts::PI * r * r),
    Shape::Rect(w, h) => println!("area: {}", w * h),
}
```

## Part 6: Ownership & Borrowing Quick Rules

| Rule | Explanation |
| :--- | :--- |
| Each value has one owner. | When the owner goes out of scope, the value is dropped. |
| You can have many `&T` references at once. | But only if there are no active `&mut T` references. |
| You can have exactly one `&mut T` reference. | And no active `&T` references at the same time. |
| References must not outlive the value they point to. | Enforced by the borrow checker using lifetimes. |
| Moving transfers ownership. | The original binding is invalidated unless the type is `Copy`. |
| `.clone()` makes a deep copy. | Allows reuse after a move, at a performance cost. |

## Part 7: Commonly Used Attributes

| Attribute | Effect |
| :--- | :--- |
| `#[derive(...)]` | Auto-generates trait implementations (`Debug`, `Clone`, `PartialEq`, etc.). |
| `#[allow(dead_code)]` | Suppresses the unused-code warning for the annotated item. |
| `#[cfg(test)]` | Compiles the block only during `cargo test`. |
| `#[test]` | Marks a function as a unit test. |
| `#[inline]` | Hints to the compiler to inline the function at call sites. |
| `#[must_use]` | Generates a warning if the return value is discarded. |
| `#[deprecated]` | Marks an item as deprecated with an optional message. |
| `#[repr(C)]` | Lays out a struct with C-compatible field ordering and alignment. |
| `#[non_exhaustive]` | Prevents external crates from exhaustively matching the enum/struct. |
| `#![no_std]` | Opts the crate out of the standard library (embedded / OS development). |
