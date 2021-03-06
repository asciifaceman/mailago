package mailago

import (
  "encoding/json"
  "errors"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "strings"
  "time"

  "github.com/gorilla/handlers"
  "github.com/gorilla/mux"
  "github.com/goware/emailx"
  sendgrid "github.com/sendgrid/sendgrid-go"
  "github.com/sendgrid/sendgrid-go/helpers/mail"
  mailgun "gopkg.in/mailgun/mailgun-go.v1"
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

// EmailPayload represents an email structure
type EmailPayload struct {
  From    string
  Subject string
  Body    string
  To      string
}

func formatHostPort(host string, port int) string {
  return fmt.Sprintf("%s:%d", host, port)
}

// Validate that all the fields are present
// We want all of them even tho Subject and Body can technically be empty
func validateEmailInput(email *EmailPayload) (err error) {
  if email.From == "" {
    err = fmt.Errorf("Missing [From] field in payload: %v", email)
    return
  }
  err = emailx.Validate(email.From)
  if err != nil {
    err = fmt.Errorf("[FROM] Email is not a valid email: %v", email.From)
    return
  }

  if email.To == "" {
    err = fmt.Errorf("Missing [To] field in payload: %v", email)
    return
  }
  err = emailx.Validate(email.To)
  if err != nil {
    err = fmt.Errorf("[TO] Email is not a valid email: %v", email.To)
    return
  }

  if email.Subject == "" {
    err = fmt.Errorf("Missing [Subject] field in payload: %v", email)
    return
  }
  if email.Body == "" {
    err = fmt.Errorf("Missing [Body] field in payload: %v", email)
    return
  }
  return nil
}

// New returns an instance of Mailago
func New(host string, port int, staticdir string) *Mailago {
  r := mux.NewRouter()
  r.HandleFunc("/health", healthHandler).Methods("GET")
  r.HandleFunc("/send", sendHandler).Methods("POST")
  r.PathPrefix("/").Handler(http.FileServer(http.Dir(staticdir)))

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

// Create new mailgun object
func newMailgun() (mailgun.Mailgun, error) {
  domain := os.Getenv("MAILGUN_DOMAIN")
  if domain == "" {
    return nil, errors.New("required env var MAILGUN_DOMAIN missing")
  }
  apiKey := os.Getenv("MAILGUN_API_KEY")
  if apiKey == "" {
    return nil, errors.New("required env var MAILGUN_API_KEY missing")
  }
  publicAPIKey := os.Getenv("MAILGUN_PUB_KEY")
  if publicAPIKey == "" {
    return nil, errors.New("required env var MAILGUN_PUB_KEY missing")
  }

  mg := mailgun.NewMailgun(
    domain,
    apiKey,
    publicAPIKey,
  )
  return mg, nil
}

// Send email via mailgun
func sendMailgun(payload *EmailPayload) error {
  // Nasty dirty, should keep one mailgun but my Mailago instantiation won't support it
  mg, err := newMailgun()
  if err != nil {
    msg := fmt.Sprintf("There was an error in setting up the Mailgun connection: %v", err.Error())
    log.Print(msg)
    return fmt.Errorf("Mailgun connection failed: %v", msg)
  }

  msg, id, err := mg.Send(mg.NewMessage(payload.From, payload.Subject, payload.Body, payload.To))
  if err != nil {
    msg := fmt.Sprintf("Could not send message: %v, ID %v, %+v", err, id, msg)
    log.Print(msg)
    return fmt.Errorf("Mailgun send failed: %v", msg)
  }

  return nil
}

// Create a new sendgrid object
func newSendgrid() (*sendgrid.Client, error) {
  key := os.Getenv("SENDGRID_KEY")
  if key == "" {
    return nil, errors.New("required env var SENDGRID_KEY missing")
  }
  sg := sendgrid.NewSendClient(key)
  return sg, nil
}

func sendSG(payload *EmailPayload) error {
  // Nasty dirty, should keep one send grid but my Mailago instantiation won't support it
  sg, err := newSendgrid()
  if err != nil {
    msg := fmt.Errorf("There was an error in setting up the Sendgrid connection: %v", err.Error())
    log.Print(msg)
    return msg
  }
  from := mail.NewEmail(strings.Split(payload.From, "@")[0], payload.From)
  to := mail.NewEmail(strings.Split(payload.To, "@")[0], payload.To)
  message := mail.NewSingleEmail(from, payload.Subject, to, payload.Body, payload.Body)

  response, err := sg.Send(message)
  if err != nil {
    msg := fmt.Errorf("Could not send message: %v", response)
    log.Print(msg)
    return msg
  }

  return nil
}

func extractAndValidate(r *http.Request) (*EmailPayload, error) {
  defer r.Body.Close()
  body, err := ioutil.ReadAll(r.Body)
  if err != nil {
    return nil, err
  }
  payload := &EmailPayload{}
  err = json.Unmarshal(body, payload)
  if err != nil {
    return nil, err
  }
  err = validateEmailInput(payload)
  if err != nil {
    return nil, err
  }

  return payload, nil
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
  // Read the body and set up our payload, validating the data
  payload, err := extractAndValidate(r)
  if err != nil {
    log.Print(err.Error())
    respondError(w, 400, err)
    return
  }
  // Attempt MailGun send
  err = sendMailgun(payload)
  if err != nil {
    log.Print(fmt.Errorf("Could not send via Mailgun: [%v]. Attempting SendGrid", err.Error()))
  } else {
    rem := ResponseMessage{Status: "Ok", Body: "Message has been successfully sent."}
    log.Print(rem)
    respondJSON(w, 200, rem)
    return
  }

  err = sendSG(payload)
  if err != nil {
    log.Print(fmt.Errorf("could not send via Sendgrid: [%v]. complete failure", err.Error()))
    respondError(w, 500, err)
    return
  }
  log.Print("Message sent!")
  rem := ResponseMessage{Status: "Ok", Body: "Message has been successfully sent."}
  respondJSON(w, 200, rem)
  return

}

func healthHandler(w http.ResponseWriter, r *http.Request) {
  log.Print("Health hit. Checking env vars!")
  _, err := newMailgun()
  if err != nil {
    respondError(w, 500, err)
    return
  }
  _, err = newSendgrid()
  if err != nil {
    respondError(w, 500, err)
    return
  }

  rem := ResponseMessage{Status: "Ok", Body: "Server Living. All required ENV Vars are set."}
  respondJSON(w, 200, rem)

}

func respondError(w http.ResponseWriter, code int, message error) {
  respondJSON(w, code, map[string]string{"error": message.Error()})
}

func respondJSON(w http.ResponseWriter, code int, payload interface{}) {
  response, _ := json.Marshal(payload)
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(code)
  w.Write(response)
}
