FROM golang:latest

WORKDIR /app
COPY . .
RUN go build -o api-service .

ENV JWT_SECRET="hs256-devkey-GuOd"
ENV API_KEY="7d572541135f588432728f7fa12b3b320abde07bab2dbe84"
EXPOSE 8082
CMD ["./api-service"]
