# MailaGo [![Build Status](https://travis-ci.org/asciifaceman/mailago.svg?branch=master)](https://travis-ci.org/asciifaceman/mailago)

##### Setup
- `go get -u github.com/kardianos/govendor` if necessary
- Designed around `Docker 17.12` & `docker-compose 1.18.0`

##### Development
- `make run`
- http://localost:3030

##### Deploy
- Linux
    - `make build`
- MacOS
    - `make buildosx`
- `make deploy`
- http://localhost:3030
##### API (todo)
- POST
    - `/send` 
- GET
    - `/health`
