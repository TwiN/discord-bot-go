FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -mod vendor .
CMD ["/app/discord-bot-go"]
