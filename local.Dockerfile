# Start with the golang base image
FROM golang:1.22.4-alpine3.20 as base

#ENV GO111MODULE=on
ENV config=local

# Set the current working directory inside the container 
WORKDIR /build

# Copy go mod and sum files 
COPY go.mod go.sum ./

# Install ca-certificates for X509 validation
RUN apk update && apk upgrade && apk add --no-cache ca-certificates
RUN update-ca-certificates

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download && go mod verify

# # Copy the source from the current directory to the working Directory inside the container 
COPY . .

WORKDIR /build/cmd/api

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main

# Start a new stage from scratch
FROM scratch
# RUN apk --no-cache add ca-certificates

# Run the config as a different user than root for safety
WORKDIR /root/
ENV config=local

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=base /build/cmd/api/main .
COPY --from=base /build/configs/local.yaml .
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Expose port 5000 to the outside world
EXPOSE 8008
EXPOSE 5555
EXPOSE 7070

#Command to run the executable
CMD ["./main"]