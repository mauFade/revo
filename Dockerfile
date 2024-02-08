FROM golang:1.20-alpine

RUN mkdir -p /usr/src/app

WORKDIR /usr/src/app

COPY go.mod .
COPY go.sum .
COPY . .

RUN go mod download

EXPOSE 8081

RUN chmod +x entrypoint.sh

CMD [ "./entrypoint.sh" ]
