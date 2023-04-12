FROM golang:1.20.3
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN go install github.com/swaggo/swag/cmd/swag@latest
COPY . ./
RUN go build -v -o ./app

EXPOSE 8787

# CMD sleep infinity
CMD swag init && ./app
