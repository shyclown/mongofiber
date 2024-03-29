FROM golang:1.22

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

ENV HOST="0.0.0.0"
ENV PORT="8080"

COPY . .
RUN go build -v -o /usr/local/bin/ ./...

CMD ["mongofiber"]

EXPOSE 8080