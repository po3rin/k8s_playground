FROM golang:alpine AS builder

WORKDIR /work
COPY main.go .
RUN go build -o server .

FROM alpine

WORKDIR /exec
COPY --from=builder /work/server .
CMD ["./server"]