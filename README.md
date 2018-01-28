# MailaGo [![Build Status](https://travis-ci.org/asciifaceman/mailago.svg?branch=master)](https://travis-ci.org/asciifaceman/mailago)

A very basic, naive, insecure mailer API with automatic failover.

##### Setup
- `go get -u github.com/kardianos/govendor` if necessary
- Designed around `Docker 17.12` & `docker-compose 1.18.0`
- Written for `go1.8.3`
- Frontend development requires `npm` for dependencies and scripts. Not required to use Mailago (just convenient)

##### Mailgun & Sendgrid
for `make run`:

```
export MAILGUN_DOMAIN=yourdomain
export MAILGUN_API_KEY=yourkey
export MAILGUN_PUB_KEY=yourkey
export SENDGRID_KEY=sendgridapikey
```
or update docker-compose.yaml for `make deploy`.

```
    environment:
      - MAILGUN_API_KEY=key-
      - MAILGUN_PUB_KEY=pubkey-
      - MAILGUN_DOMAIN=sandboxBLAH.mailgun.org
      - SENDGRID_KEY=sendgridapikey
```

##### Development & Quick Start
- Set env vars
- `make run`
- http://localhost:3031

##### Build
- Linux
    - `make build` or `make all`
- MacOS
    - `make buildosx`

##### Deploy
- Requires docker-compose & docker
- Make sure to add your mailgun & sendgrid information to docker-compose.yaml
- `make deploy`
- `curl http://localhost:8080/health` to check if it is running properly and to check env vars

MANUALLY:
- Requires docker-compose & docker
- Make sure to add your mailgun & sendgrid information to docker-compose.yaml
- `docker-compose -f docker-compose.yaml up --build`
- `curl http://localhost:8080/health` to check if it is running properly and to check env vars

##### API

- You can utilize `/etc/hosts` to fake an outage:
    - `127.0.0.1    api.mailgun.net`
    - `127.0.0.1    api.sendgrid.com`
- POST
    - `/send`
        - Sends an email using Mailgun, however if an error is met, will attempt Sendgrid.
        - REQUIRED: `From, To, Subject, Body`
        - Ex: 
        ```
        curl -d '{"From": "tester@mailago.io", "To": "you@gmail.com", "Subject": "Test email", "Body": "This is just a test email."}' -H "Content-Type: application/json"  -X POST localhost:8080/send
        ```
- GET
    - `/health`
        - Returns status of Mailago and checks if ENV VARs are set.