FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/github.com/boringmary/gomem

COPY . .

RUN go mod download
RUN go build -o /go/bin/gomem

FROM alpine

COPY --from=builder /go/bin/gomem /go/bin/gomem
RUN chmod +x /go/bin/gomem
EXPOSE 5000
ENTRYPOINT ["/go/bin/gomem"]