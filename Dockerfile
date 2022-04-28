FROM golang:alpine

WORKDIR /app
COPY . .
# download git from alpine
RUN apk add --no-cache git
RUN go mod download
RUN go build -o /web

EXPOSE 8080

CMD [ "/web" ]