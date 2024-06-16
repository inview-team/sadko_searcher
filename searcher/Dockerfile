FROM golang:1.21 AS builder
ENV PROJECT_PATH=/app/searcher
ENV CGO_ENABLED=0
ENV GOOS=linux
COPY . ${PROJECT_PATH}
WORKDIR ${PROJECT_PATH}
RUN go build cmd/searcher/main.go

FROM golang:alpine
WORKDIR /app/cmd/searcher
COPY --from=builder /app/searcher/main .
EXPOSE 30002
CMD ["./main"]