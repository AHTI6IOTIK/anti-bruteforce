FROM golang:latest AS builder
ENV NAME "anti-bruteforce"

WORKDIR /app
ENV CGO_ENABLED=0
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o bin/${NAME} ./cmd/${NAME}
RUN go build -o bin/abf ./cmd/cli


FROM alpine
ENV NAME "anti-bruteforce"

COPY --from=builder /app/config.yml /app/config.yml
COPY --from=builder /app/bin/${NAME} /app/
COPY --from=builder /app/bin/abf /app/

WORKDIR /app

EXPOSE 50051
CMD ./${NAME}
