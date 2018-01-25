# MailaGo [![Build Status](https://travis-ci.org/asciifaceman/mailago.svg?branch=master)](https://travis-ci.org/asciifaceman/mailago)

##### Setup
- `go get -u github.com/kardianos/govendor` if necessary
- Designed around `Docker 17.12` & `docker-compose 1.18.0`
- Written for `go1.8.3`

##### Development
- `make run`
- http://localhost:3031

##### Front-end Development
- `./static`

```
yarn start
  Starts the development server.

yarn build
  Bundles the app into static files for production.

yarn test
  Starts the test runner.

yarn eject
  Removes this tool and copies build dependencies, configuration files
  and scripts into the app directory. If you do this, you can’t go back!
```

##### Deploy
- Linux
    - `make build`
- MacOS
    - `make buildosx`
- `make deploy`
- http://localhost:3031
##### API (todo)
- POST
    - `/send` 
- GET
    - `/health`
