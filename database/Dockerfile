FROM golang:latest as build-env

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...
RUN go build -o /go/bin/app

FROM gcr.io/distroless/base
USER nonroot
COPY --chown=nonroot:nonroot --from=build-env /go/bin/app /
CMD ["/app"]
