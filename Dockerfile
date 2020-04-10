FROM golang:1.13.9-alpine3.11 AS builder

RUN apk update && apk add \
  make \
  git

WORKDIR /go/src/github.com/yanyi/generate-multifields
ENV GO111MODULE=auto

COPY . .
RUN go get -v ./...

RUN GOOS=linux GOARCH=386 go build -ldflags="-w -s" -o /go/bin/generate-multifields

FROM scratch

COPY --from=builder /go/bin/generate-multifields /bin/generate-multifields

ENTRYPOINT [ "/bin/generate-multifields" ]
