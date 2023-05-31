package main

import (
  "fmt"
  "log"

  "github.com/FarStep131/go-simple-server/server"
)

const (
  host = "0.0.0.0"
  port = "8080"
)

func main() {
  s := server.New()
  err := s.Start(fmt.Sprintf("%s:%s", host, port))
  if err != nil {
    s.Stop()
    log.Fatalf("server stopped with error: %s", err)
  }
}
