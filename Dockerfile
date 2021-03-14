FROM golang:1.16-alpine3.13
WORKDIR /go/src/github.com/orisano/dignore
COPY * ./
RUN go build

FROM alpine:3.13
COPY --from=0 /go/src/github.com/orisano/dignore/dignore /bin/dignore
ENTRYPOINT ["/bin/dignore"]
