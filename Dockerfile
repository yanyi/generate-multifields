FROM golang:1.13.9-alpine3.11 AS builder

RUN apk update && apk add \
  make \
  git

WORKDIR /go/src/github.com/yanyi/generate-multifields
ENV GO111MODULE=auto

COPY . .
RUN go get -v ./...

ARG CLI_VERSION
ENV CLI_VERSION=${CLI_VERSION:-nil}
ARG BUILD_SHA
ENV BUILD_SHA=${BUILD_SHA:-nil}

RUN GOOS=linux GOARCH=386 go build \
  -ldflags="-w -s \
  -X 'github.com/yanyi/generate-multifields/cmd.Version=$CLI_VERSION' \
  -X 'github.com/yanyi/generate-multifields/cmd.Build=$BUILD_SHA'" \
  -o /go/bin/generate-multifields

##########

FROM scratch

COPY --from=builder /go/bin/generate-multifields /bin/generate-multifields

ENTRYPOINT [ "/bin/generate-multifields" ]
