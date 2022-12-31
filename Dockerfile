FROM golang:1.17.6-alpine
# Add a work directory
WORKDIR /ditalrepublic
# Cache and install dependencies
COPY go.mod go.sum ./
RUN go mod download
# Copy app files
COPY . .
# Build app
RUN go build cmd
# Expose port
EXPOSE 8080
# Start app
CMD ./paintcalculator