```go
package main

import (

    "fmt"
)

func main() {
    m := make(map[int]int)
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
        fmt.Println("Exiting gracefully...")
    }()

    m[1] = 1
    m[2] = 2
    val, ok := m[3] //Check for existence
    if ok {
        fmt.Println("Value at key 3:", val)
    } else {
        fmt.Println("Key 3 not found in the map")
    }
    m[3] = 3
}
```