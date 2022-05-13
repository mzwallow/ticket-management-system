FROM golang:alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build ./cmd/ticket/

EXPOSE 5000

CMD [ "./ticket" ]