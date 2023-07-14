package main


import (
  "fmt"
  "net/http"
)

func main() {
 for i := 0; i < 400; i++ {
    resp, err := http.Get("https://example.com")
    fmt.Printf("\rSend: %d%%", i)
    fmt.Println(resp, err)
  }
}
