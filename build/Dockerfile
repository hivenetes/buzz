FROM golang:1.18.0 as builder
# # Define build env
ENV GOOS linux
ENV CGO_ENABLED 0
# Add a work directory
WORKDIR /buzz
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod tidy
# Copy app files
COPY . .
# Build app
RUN go build -o buzz

FROM alpine:3.16 as production
# Add certificates
RUN apk add --no-cache ca-certificates
# Copy built binary from builder
COPY --from=builder buzz .
# Expose port
EXPOSE 3001
# Exec built binary
CMD ["./buzz"]