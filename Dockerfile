FROM golang:1.27
RUN go install github.com/air-verse/air@latest
WORKDIR /app
CMD ["air", "-c", ".air.toml"]
