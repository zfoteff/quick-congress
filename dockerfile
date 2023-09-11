FROM golang:1.20

# COPY destination
WORKDIR /quick-congress

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o container-quick-congress-application

EXPOSE 8080

# Run application in web-app mode
CMD ["container-quick-congress-application"]
