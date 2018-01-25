FROM ubuntu:trusty

#RUN apt-get update && apt-get install -y ca-certificates

add target/mailago mailago

ENTRYPOINT ["./mailago"]