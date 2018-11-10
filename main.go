package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	Handler "./handler"
)

var TOKEN = os.Getenv("SECRETS_DISCORD_BOT_TOKEN")


func main() {
	var bot = connect()
	registerHandlers(bot)

	// Wait for a CTRL-C
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Close()
}


func connect() *discordgo.Session {
	var bot, err = discordgo.New("Bot " + TOKEN)

	fmt.Println("[connect] Connecting to Discord...")
	err = bot.Open()
	if err != nil {
		log.Printf("[connect] Failed to establish connection with Discord: %s\n", err)
		os.Exit(1)
	}
	fmt.Println("[connect] Connected to Discord successfully!")

	return bot
}


func registerHandlers(bot *discordgo.Session) {
	bot.AddHandler(Handler.ShortcutConverterHandler)
	bot.AddHandler(Handler.SayHandler)
	bot.AddHandler(Handler.GoogleSearchHandler)
	bot.AddHandler(Handler.YoutubeSearchHandler)
	bot.AddHandler(Handler.PingPongHandler)
	bot.AddHandler(loggerHandler)
}


func loggerHandler(bot *discordgo.Session, message *discordgo.MessageCreate) {
	fmt.Println("[" + message.Author.Username + "] " + message.ContentWithMentionsReplaced())
}