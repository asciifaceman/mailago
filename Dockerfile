FROM golang:1.8 as builder
WORKDIR /go/src/github.com/asciifaceman/mailago
RUN mkdir /dist
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /dist/mailago-linux-x86 .

FROM scratch
COPY --from=builder /dist/mailago-linux-x86 /bin/mailago
COPY /static /static
ENV PATH=/bin

ENTRYPOINT ["/bin/mailago", "run", "-s", "/static"]