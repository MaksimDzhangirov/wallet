FROM golang:1.16-alpine as builder
RUN apk add --no-cache git
RUN go get github.com/golang/mock/mockgen

FROM golang:1.16-alpine
COPY . /app
WORKDIR /app
COPY --from=builder /go/bin/mockgen /bin/mockgen
CMD tail -f /dev/null