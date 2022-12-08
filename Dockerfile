FROM golang:1.18.1-alpine
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go mod tidy
# COPY *.go .
COPY . .
EXPOSE 8080
ENTRYPOINT [ "go", "run", "main.go" ]