FROM golang

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN chmod -R 777 "$GOPATH"

WORKDIR /go/src/golang-todo-api

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o ./todo-api

RUN chmod +x todo-api

EXPOSE $APP_PORT

CMD ["./todo-api"]
