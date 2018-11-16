FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get github.com/PuerkitoBio/goquery github.com/bwmarrin/discordgo github.com/TwinProduction/go-away && \
	go build -o main .
CMD ["/app/main"]
