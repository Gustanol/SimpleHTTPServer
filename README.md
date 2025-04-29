# This is a simple HTTP server builded in Go

### How does it work? 🤔

We use the `net/http` default package to build the server.

The `HandleFunction` associate another function to the root route, which means it will be called in every route.
```go
http.HandleFunc("/", handler) // calls `handler` function
```

The defined listening port is `8080`

---

Execute the server using:
```go
go run main.go
```

Then, visit [http://localhost:8080](http://localhost:8080) in your browser