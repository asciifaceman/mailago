# MailaGo [![Build Status](https://travis-ci.org/asciifaceman/mailago.svg?branch=master)](https://travis-ci.org/asciifaceman/mailago)

A very basic insecure mailer API.

##### Setup
- `go get -u github.com/kardianos/govendor` if necessary
- Designed around `Docker 17.12` & `docker-compose 1.18.0`
- Written for `go1.8.3`
- Frontend development requires `npm` for dependencies and scripts. Not required to use Mailago (just convenient)

##### Mailgun
```
export MAILGUN_DOMAIN=yourdomain
export MAILGUN_API_KEY=yourkey
export MAILGUN_PUB_KEY=yourkey
```
or add to docker-compose.yaml for `make deploy`

##### Development & Quick Start
- `make run`
- http://localhost:3031
    - for frontend see below, otherwise you can `curl` the API

##### Front-end Development

- Development
    - `cd frontend`
    - `npm start`
    - Develop some stuff
- Deploy / Build
    - `make frontend`
        - Builds react project and places static content in `./static`
        - This will work with `go run main.go run` but is not required for Mailago
    - `make deploy`
        - Adds `./static` to `/static` in Docker image
        - Also deploys the whole application to docker using docker-compose

##### Build
- Linux
    - `make build` or `make all`
- MacOS
    - `make buildosx`

##### Deploy
- Requires docker-compose & docker
- `make deploy`
- http://localhost:8080 or see API below for `curl`

##### API (todo)
You can utilize `/etc/hosts` to fake an outage:
- Mailgun
    - `127.0.0.1    api.mailgun.net`
- POST
    - `/send`
        - Sends an email using Mailgun, however if an error is met, will attempt Sendgrid.
        - REQUIRED: `From, To, Subject, Body`
        - EX: `curl -d '{"From": "tester@mailago.io", "To": "tester@gmail.com"a test email."}' -H "Content-Type: application/json"  -X POST localhost:3031/send`
    - `/send/mailgun` Deprecated
        - Utilizes the mailgun API to send an email. Will not retry.
        - REQUIRED: `From, To, Subject, Body`
        - EX: `curl -d '{"From": "tester@mailago.io", "To": "tester@gmail.com"a test email."}' -H "Content-Type: application/json"  -X POST localhost:3031/send/mailgun`
    - `/send/sendgrid` Deprecated
        - Utilizes the sendgrid API to send an email. Will not retry.
        - REQUIRED: `From, To, Subject, Body`
        - EX: `curl -d '{"From": "tester@mailago.io", "To": "tester@gmail.com"a test email."}' -H "Content-Type: application/json"  -X POST localhost:3031/send/sendgrid`
- GET
    - `/` & `/#/dashboard`
        - Frontend UI (if applicable)
    - `/health`
        - Returns the status of mailer APIs used by Mailago
