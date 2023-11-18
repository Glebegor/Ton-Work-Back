FROM golang:latest
LABEL author="Glebegor"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/main.go ./

RUN go build -o /maindocker 

CMD ["/maindocker"]