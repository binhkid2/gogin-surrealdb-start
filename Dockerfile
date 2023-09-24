# Build stage
FROM docker.io/library/golang:1.21.1 as builder

WORKDIR /app
COPY . .

# Check if go.mod exists, if not, then initialize the module
RUN test -f go.mod || go mod init github.com/binhkid2/gogin-surrealdb-start

RUN CGO_ENABLED=0 GOOS=linux go build -o gogin-surrealdb-start

# Final stage
FROM docker.io/library/alpine:3.18.3

WORKDIR /app
COPY --from=builder /app/gogin-surrealdb-start /app/
# Copy the config.ini file into the /app directory in the container
COPY config.ini /app/
EXPOSE 5678
CMD ["/app/gogin-surrealdb-start"]
