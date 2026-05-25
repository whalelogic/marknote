# Go `fmt` Reference Table

A quick reference for the `fmt` package in Go.

## 1. The Functions (Print vs. Sprint vs. Fprint)

The `fmt` package has three main families of functions based on **where** the output goes.

### The "S" Family (Returns a String)
Use these when you want to assign the result to a variable string.

| Function | Description | Example Code | Result |
| :--- | :--- | :--- | :--- |
| **`fmt.Sprint()`** | Concatenates arguments. No spaces added between non-strings. | `s := fmt.Sprint("Age:", 42)` | `"Age:42"` |
| **`fmt.Sprintf()`** | Formats according to a specifier (like C's `printf`). **Most common.** | `s := fmt.Sprintf("Age: %d", 42)` | `"Age: 42"` |
| **`fmt.Sprintln()`** | Adds spaces between args and appends a newline. | `s := fmt.Sprintln("Age:", 42)` | `"Age: 42\n"` |

### The "Print" Family (Writes to Stdout)
Use these for logging to the console.

| Function | Description | Example Code | Output |
| :--- | :--- | :--- | :--- |
| **`fmt.Print()`** | Prints to Stdout. No spaces, no newline. | `fmt.Print("Age:", 42)` | `Age:42` |
| **`fmt.Printf()`** | Prints to Stdout with format specifiers. | `fmt.Printf("Age: %d", 42)` | `Age: 42` |
| **`fmt.Println()`** | Prints to Stdout. Adds spaces between args and a newline. | `fmt.Println("Age:", 42)` | `Age: 42` |

### The "F" Family (Writes to an `io.Writer`)
Use these to write directly to files, web responses (`http.ResponseWriter`), or buffers (`bytes.Buffer`).

| Function | Description | Example Code |
| :--- | :--- | :--- |
| **`fmt.Fprint(w, ...)`** | Writes to `w`. No spaces, no newline. | `fmt.Fprint(w, "Hello")` |
| **`fmt.Fprintf(w, ...)`** | Writes to `w` with format specifiers. | `fmt.Fprintf(w, "Status: %d", 200)` |
| **`fmt.Fprintln(w, ...)`** | Writes to `w` with spaces and newline. | `fmt.Fprintln(w, "Hello World")` |

### The "Error" Family
| Function | Description | Example |
| :--- | :--- | :--- |
| **`fmt.Errorf()`** | Formats a string and returns it as an `error` type. Wraps errors with `%w`. | `return fmt.Errorf("db fail: %w", err)` |

---

## 2. Formatting Verbs (`%`)

Used primarily with `Sprintf`, `Printf`, and `Fprintf`.

### General & Debugging
| Verb | Use Case | Example Input | Output |
| :--- | :--- | :--- | :--- |
| **`%v`** | **Default format** (works for almost anything). | `User{Name: "Alice"}` | `{Alice}` |
| **`%+v`** | Prints struct fields with **field names**. | `User{Name: "Alice"}` | `{Name:Alice}` |
| **`%#v`** | **Go-syntax representation** (type + value). | `User{Name: "Alice"}` | `main.User{Name:"Alice"}` |
| **`%T`** | Prints the **Type** of the value. | `User{}` | `main.User` |
| **`%%`** | Prints a literal percent sign. | | `%` |

### Integers & Numbers
| Verb | Description | Example Input | Output |
| :--- | :--- | :--- | :--- |
| **`%d`** | Base 10 Integer. | `42` | `42` |
| **`%b`** | Binary (base 2). | `42` | `101010` |
| **`%x`** | Hexadecimal (base 16, lowercase). | `42` | `2a` |
| **`%X`** | Hexadecimal (base 16, uppercase). | `42` | `2A` |
| **`%c`** | The character represented by the integer. | `65` | `A` |
| **`%p`** | Pointer address (in hex). | `&x` | `0xc0000...` |

### Strings & Bytes
| Verb | Description | Example Input | Output |
| :--- | :--- | :--- | :--- |
| **`%s`** | Basic string or byte slice. | `"Hello"` | `Hello` |
| **`%q`** | Double-quoted string (safely escapes newlines/tabs). | `"Hello\nWorld"` | `"Hello\nWorld"` |
| **`%x`** | Base 16 (hex) dump of byte slice/string. | `"ABC"` | `414243` |

### Booleans
| Verb | Description | Output |
| :--- | :--- | :--- |
| **`%t`** | The word true or false. | `true` |

### Floats
| Verb | Description | Example Input | Output |
| :--- | :--- | :--- | :--- |
| **`%f`** | Decimal point, no exponent. | `123.456` | `123.456000` |
| **`%.2f`** | **Precision**: Limit to 2 decimal places. | `123.456` | `123.46` |
| **`%e`** | Scientific notation. | `123.456` | `1.234560e+02` |

---

## 3. Padding and Width
You can control the width of the output for nice table alignment.

| Code | Effect | Result |
| :--- | :--- | :--- |
| **`%5d`** | Pad with spaces on left (width 5). | `|   42|` |
| **`%05d`** | Pad with **zeros** on left (width 5). | `|00042|` |
| **`%-5d`** | Pad with spaces on **right** (left-align). | `|42   |` |
| **`%10s`** | String padding (left). | `|     hello|` |
| **`%-10s`** | String padding (right). | `|hello     |` |

## 4. Input Scanning (Reading)
The `fmt` package also handles reading input from Stdin.

| Function | Description |
| :--- | :--- |
| **`fmt.Scan(&x)`** | Reads space-separated values into variables. Newlines are treated as spaces. |
| **`fmt.Scanln(&x)`** | Reads until a newline is encountered. |
| **`fmt.Scanf("%d:%d", &h, &m)`** | Reads input matching a specific format string. |
