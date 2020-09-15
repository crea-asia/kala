# Change to golang:alpine so the binary can run on alpine seemlessly
FROM golang

WORKDIR /go/src/github.com/crea-asia/kala
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o kala

FROM alpine:latest
WORKDIR /app
# Copy the binary file from the first image
COPY ./webui /app/webui
COPY --from=0 /go/src/github.com/crea-asia/kala/kala /app/kala

CMD ["/app/kala", "serve"]
EXPOSE 8000
