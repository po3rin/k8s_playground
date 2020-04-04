FROM golang:1.13-alpine as builder

WORKDIR /src/
COPY main.go /src/
RUN DGO_ENABLED=0 go build -o /bin/demo

FROM alpine:latest
COPY --from=builder /bin/demo /bin/
ENTRYPOINT ["/bin/demo"]

