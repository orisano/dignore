FROM golang:1.17-alpine3.14
WORKDIR /go/src/github.com/orisano/dignore
COPY * ./
RUN go build

FROM alpine:3.14
COPY --from=0 /go/src/github.com/orisano/dignore/dignore /bin/dignore
ENTRYPOINT ["/bin/dignore"]
