FROM ubuntu:trusty

RUN apt-get update && apt-get install -y ca-certificates

add mailago mailago

ENTRYPOINT ["./mailago"]