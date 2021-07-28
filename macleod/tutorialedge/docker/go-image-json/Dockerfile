FROM golang:latest

LABEL maintainer ="Onkar json practice" 

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build  -o main .

EXPOSE 3333

CMD ["./main"]


