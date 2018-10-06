FROM golang:1.11

WORKDIR $GOPATH/src/github.com/nigeltiany/micro_istio

COPY . .

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -mod=vendor -o app

FROM scratch

COPY --from=0 /go/src/github.com/nigeltiany/micro_istio/app /usr/bin/app

ENTRYPOINT ["app"]

