FROM golang:alpine as app-builder
WORKDIR /go/src
COPY . .
# Static build required so that we can safely copy the binary over.
# `-tags timetzdata` embeds zone info from the "time/tzdata" package.
RUN CGO_ENABLED=0 go build -o app main.go

FROM ubuntu:21.04

COPY --from=app-builder /go/src/app /app
EXPOSE 5050
ENTRYPOINT ["/app"]