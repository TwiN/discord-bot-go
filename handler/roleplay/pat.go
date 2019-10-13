package roleplay

import (
	"github.com/bwmarrin/discordgo"
	"math/rand"
	"strings"
)

var pats = []string{
	"https://media.tenor.com/images/0d9d44e6a9577eb28c47b22f5acd7d69/tenor.gif",
	"https://thumbs.gfycat.com/AgileHeavyGecko-max-1mb.gif",
	"https://i.imgur.com/4ssddEQ.gif",
	"https://media1.tenor.com/images/1e92c03121c0bd6688d17eef8d275ea7/tenor.gif?itemid=9920853",
	"https://data.whicdn.com/images/297125626/original.gif",
	"https://78.media.tumblr.com/755e34a2e227e153ba64d4a4f27848fe/tumblr_nahgozuIwM1rbnx7io1_500.gif",
	"https://media1.tenor.com/images/c0bcaeaa785a6bdf1fae82ecac65d0cc/tenor.gif?itemid=7453915",
	"http://i.imgur.com/laEy6LU.gif",
	"https://archive-media-0.nyafuu.org/c/image/1483/55/1483553008493.gif",
}

func Pat(bot *discordgo.Session, message *discordgo.MessageCreate) {
	img := pats[rand.Intn(len(pats))]
	msg := &discordgo.MessageEmbed{}
	msg.Description = message.Author.Mention() + " has patted **" + strings.Replace(message.Message.Content, "!pat ", "", -1) + "**"
	msg.Image = &discordgo.MessageEmbedImage{URL: img}
	bot.ChannelMessageSendEmbed(message.ChannelID, msg)
}
