FROM golang:1.21.4-alpine
WORKDIR /app
COPY  . .
RUN go mod tidy
RUN go build -o svr cmd/httpserver/main.go
EXPOSE 8080
CMD [ "./svr" ]