package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"reflect"
	"runtime"

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

	log.Println("[connect] Connecting to Discord...")
	err = bot.Open()
	if err != nil {
		log.Printf("[connect] Failed to establish connection with Discord: %s\n", err)
		os.Exit(1)
	}
	log.Println("[connect] Connected to Discord successfully!")

	return bot
}


var registeredHandlers []interface{}

func init() {
	registeredHandlers = append(registeredHandlers, 
		Handler.ShortcutConverterHandler, 
		Handler.BasicHandler,
		Handler.GoogleSearchHandler, 
		Handler.YoutubeSearchHandler,
		Handler.UrbanDictionarySearchHandler,
		Handler.PingPongHandler,
		loggerHandler,
	)
}


func registerHandlers(bot *discordgo.Session) {
	bot.UpdateStatus(1, "Registering handlers")
	for _, handler := range registeredHandlers {
		log.Println("[registerHandlers] Registering", runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name())
		bot.AddHandler(handler)
	}
	bot.UpdateStatus(0, "")
}


func loggerHandler(bot *discordgo.Session, message *discordgo.MessageCreate) {
	log.Println(message.Author.Username, "-", message.ContentWithMentionsReplaced())
}