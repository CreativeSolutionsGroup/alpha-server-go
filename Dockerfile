FROM golang:alpine as builder

WORKDIR /app
COPY . .
# download git from alpine
RUN apk add --no-cache git
RUN go mod download
RUN go build -o /web

EXPOSE 8080

FROM alpine:latest 
WORKDIR /app
COPY --from=builder /web .

CMD [ "/app/web" ]