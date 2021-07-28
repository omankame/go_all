FROM golang:latest

LABEL maintainer ="Onkar Practice"

WORKDIR /build

COPY . .

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .

EXPOSE 8080


CMD ["/dist/main"]


