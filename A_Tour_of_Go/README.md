# A Tour of Go — code exercises

Hands-on Go snippets written while working through the official [A Tour of Go](https://go.dev/tour/) (`golang.org/x/tour`). Each numbered folder groups the files by the section of the tour it corresponds to; files inside are numbered in the order they were written.

| Folder | Topic | Files |
|---|---|---|
| [1](1) | Getting started — Hello, World and the Go Playground/sandbox | `001`–`002` |
| [2](2) | Packages, imports, functions, variables, basic types, constants | `003`–`017` |
| [3](3) | Flow control — `for`, `if`, `switch`, `defer` | `018`–`030` |
| [4](4) | More types — pointers, structs, arrays, slices, maps, function closures | `031`–`056` |
| [5](5) | Methods and interfaces — pointer vs. value receivers, `Stringer`, errors, `io.Reader`, images | `057`–`076` |
| [6](6) | Generics — type parameters for functions and types | `077`–`078` |
| [7](7) | Concurrency — goroutines, channels, `select`, `sync.Mutex` | `079`–`087` |

## Running an example

Each file declares `package atourofgo`, so pick a file, and run/build it with the module root's `go.mod`:

```bash
cd A_Tour_of_Go
go build ./...
```

Snippets are kept as written while going through the tour — logic is untouched; only `gofmt` formatting is enforced.
