package mailago

import (
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "time"

  "github.com/gorilla/handlers"
  "github.com/gorilla/mux"
)

// Mailago represents the mailer instance
type Mailago struct {
  router *mux.Router
  Srv    *http.Server
}

// ResponseMessage represents a json body response
type ResponseMessage struct {
  Status string
  Body   string
}

func formatHostPort(host string, port int) string {
  return fmt.Sprintf("%s:%d", host, port)
}

// New returns an instance of Mailago
func New(host string, port int) *Mailago {
  r := mux.NewRouter()
  r.HandleFunc("/health", healthHandler)

  s := &http.Server{
    Handler:      handlers.RecoveryHandler()(r),
    Addr:         formatHostPort(host, port),
    WriteTimeout: 15 * time.Second,
    ReadTimeout:  15 * time.Second,
  }

  return &Mailago{
    router: r,
    Srv:    s,
  }
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
  log.Print("Health hit")
  rem := ResponseMessage{Status: "Ok", Body: "Server Living."}
  mess, err := json.Marshal(rem)
  if err != nil {
    panic(err.Error())
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(mess)

}
