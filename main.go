package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: pingmate <url>")
        return
    }

    url := os.Args[1]
    resp, err := http.Get(url)
    if err != nil {
        fmt.Printf("❌ %s is down: %s\n", url, err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode == 200 {
        fmt.Printf("✅ %s is UP!\n", url)
    } else {
        fmt.Printf("⚠️ %s returned status: %d\n", url, resp.StatusCode)
    }
}

