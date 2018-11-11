package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"reflect"
	"runtime"
	"time"

	"github.com/bwmarrin/discordgo"

	Handler "./handler"
	"./config"
)

var TOKEN = os.Getenv("SECRETS_DISCORD_BOT_TOKEN")
var registeredHandlers []interface{}


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
	log.Println("[main][connect] Connecting to Discord...")
	err = bot.Open()
	attempts := 3
	for err != nil && attempts > 0 {
		log.Printf("[main][connect] Failed to establish connection with Discord: %s\n", err)
		log.Println("[main][connect] Retrying in 3 second")
		time.Sleep(3 * time.Second)
		err = bot.Open()
		attempts--
	}
	if attempts == 0 {
		log.Fatalln("[main][connect] Unable to establish connection")
		os.Exit(1)
	}
	log.Println("[main][connect] Connected to Discord successfully!")
	return bot
}

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
	config.Load()
}


func registerHandlers(bot *discordgo.Session) {
	bot.UpdateStatus(1, "Registering handlers")
	for _, handler := range registeredHandlers {
		log.Println("[main][registerHandlers] Registering", runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name())
		bot.AddHandler(handler)
	}
	bot.UpdateStatus(0, "")
}


func loggerHandler(bot *discordgo.Session, message *discordgo.MessageCreate) {
	log.Println(message.Author.Username, "-", message.ContentWithMentionsReplaced())
}