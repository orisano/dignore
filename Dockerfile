FROM golang:1.11-alpine3.8
WORKDIR /go/src/github.com/orisano/dignore
COPY *.go .
RUN go build

FROM scratch
COPY --from=0 /go/src/github.com/orisano/dignore/dignore /bin/dignore
ENTRYPOINT ["/bin/dignore"]
