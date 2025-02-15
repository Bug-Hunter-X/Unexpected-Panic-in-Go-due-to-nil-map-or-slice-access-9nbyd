```go
func main() {
    m := make(map[int]int)
    defer func() {
        fmt.Println("Exiting...")
        recover()
    }()

    m[1] = 1
    m[2] = 2
    fmt.Println(m[3]) // Accessing non-existent key. This will panic
    m[3] = 3
}
```