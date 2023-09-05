FROM golang:1.21-alpine

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o todo-app main.go

CMD [ "./todo-app" ]
