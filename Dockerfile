FROM golang:1.14-alpine3.11
WORKDIR /go/src/github.com/orisano/dignore
COPY *.go .
RUN go build

FROM alpine:3.11
COPY --from=0 /go/src/github.com/orisano/dignore/dignore /bin/dignore
ENTRYPOINT ["/bin/dignore"]
