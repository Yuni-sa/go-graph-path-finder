FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go build main.go dijkstra.go graph.go priorityqueue.go handler.go

FROM alpine

WORKDIR /app

ENV PORT="8080"

COPY --from=builder /build/main /app/main
COPY --from=builder /build/templates/ /app/templates/
COPY --from=builder /build/static/ /app/static/

CMD ["./main"]