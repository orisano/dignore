FROM golang:1.12-alpine3.9
WORKDIR /go/src/github.com/orisano/dignore
COPY *.go .
RUN go build

FROM alpine:3.9
COPY --from=0 /go/src/github.com/orisano/dignore/dignore /bin/dignore
ENTRYPOINT ["/bin/dignore"]
