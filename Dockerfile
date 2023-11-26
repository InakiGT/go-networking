FROM golang:1.21

WORKDIR /usr/src

COPY ["go.mod", "go.sum", "/usr/src/"]

RUN go mod download && go mod verify

COPY [".", "/usr/src/"]
RUN go build -v -o /usr/local/bin/app ./...

EXPOSE 5050

CMD [ "app" ]