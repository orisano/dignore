FROM golang:1.12-alpine3.10
WORKDIR /go/src/github.com/orisano/dignore
COPY *.go .
RUN go build

FROM alpine:3.10
COPY --from=0 /go/src/github.com/orisano/dignore/dignore /bin/dignore
ENTRYPOINT ["/bin/dignore"]
