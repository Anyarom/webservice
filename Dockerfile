FROM golang:1.16.3-alpine as builder
RUN apk add --no-cache gcc musl-dev
WORKDIR /app
COPY . .

# соберем бинарь
RUN go build -o webservice .

# соберем контейнер и запустим бинарь
FROM alpine
COPY --from=builder /app/webservice /app/
RUN chmod +x /app/webservice
EXPOSE 8080
CMD ["/app/webservice"]