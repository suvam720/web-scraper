FROM golang:1.18.1-alpine3.15
WORKDIR /app
COPY . .
RUN go mod download
RUN apk add git
RUN go build -o /my-app
CMD ["/my-app"]