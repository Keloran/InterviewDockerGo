package main

import (
  "github.com/keloran/go-healthcheck"
  "github.com/go-chi/chi"

  "fmt"
  "net/http"
  "os"
)

func _main(args []string) int {
  router := chi.NewRouter()

  router.Get("/healthcheck", healthcheck.HTTP)

  router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
    _, err := w.Write([]byte("Hello"))
    if err != nil {
      fmt.Printf("write err: %+v\n", err)
      return
    }
  })

  port := "80"
  if len(os.Getenv("PORT")) > 2 {
    port = os.Getenv("PORT")
  }
  fmt.Printf("Listening: %s", port)
  if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router); err != nil {
    fmt.Printf("Server err: %+v\n", err)
    return 1
  }

  return 0
}

func main() {
  os.Exit(_main(os.Args[1:]))
}

