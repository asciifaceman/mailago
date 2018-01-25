package mailago

import (
  "net/http"

  "github.com/gorilla/mux"
)

// Mailago represents the mailer router
type Mailago struct {
  router *mux.Router
}

func main() {
  r := mux.NewRouter()
  http.Handle("/", r)
  http.ListenAndServe(":8080", nil)
}
