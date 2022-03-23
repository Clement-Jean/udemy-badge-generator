FROM golang:1.17.8-alpine3.15 as build-env

WORKDIR /go/src/app
COPY *.go ./
COPY assets assets/

RUN go mod init
RUN go get -d -v ./...

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static

COPY --from=build-env /go/bin/app /
ENTRYPOINT ["/app"]