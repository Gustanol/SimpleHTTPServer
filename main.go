package main
import (
  "fmt"
  "net/http"
)

func main() {
  http.HandleFunc("/", handler) // associate the `handler` function to the root route
  fmt.Println("Server started in http://localhost:8080")
  http.ListenAndServe(":8080", nil) // defines the port "8080"
}

func handler(writer http.ResponseWriter, request *http.Request) {
  // writes in the browser
  fmt.Fprintf(writer, "Hi! You are current in: %s\n", request.URL.Path)
}