FROM golang:alpine3.19

WORKDIR /code

RUN apk update 

COPY . /code/

RUN go mod tidy

RUN mkdir -p /app

RUN go build -tags netgo -ldflags="-s -w" -o /app/mvc -v .

RUN rm -rf /code

CMD ["/app/mvc"]