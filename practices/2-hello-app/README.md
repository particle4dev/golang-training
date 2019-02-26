---
layout: post
title: Hello App
tags: [go, hello, app]
---

### [Hello app](https://gobyexample.com/hello-world)

Our first program will print the classic “hello world” message. Here’s the full source code.

```go
package main
import "fmt"
func main() {
    fmt.Println("hello world")
}
```

To run the program, put the code in hello-world.go and use go run.

```
$ go run hello-world.go
hello world
```

Sometimes we’ll want to build our programs into binaries. We can do this using go build.

```
$ go build hello-world.go
$ ls
hello-world	hello-world.go
```

We can then execute the built binary directly.

``
$ ./hello-world
hello world
```

### Links

- https://gobyexample.com/hello-world
