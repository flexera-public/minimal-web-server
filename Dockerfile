FROM golang:1.20-alpine3.17

LABEL label-source=container-label resource-type=container app-version=1.0

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN go build -o main .

CMD ["/app/main"]
