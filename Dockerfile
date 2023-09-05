FROM golang:1.21-alpine

WORKDIR /app

COPY . .
RUN go mod download

RUN go build -o todo-app main.go

EXPOSE 3000

CMD [ "./todo-app" ]
