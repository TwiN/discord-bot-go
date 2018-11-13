package roleplay

import (
	"math/rand"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var pats = []string{
	"https://media.tenor.com/images/0d9d44e6a9577eb28c47b22f5acd7d69/tenor.gif",
	"https://thumbs.gfycat.com/AgileHeavyGecko-max-1mb.gif",
	"https://i.imgur.com/4ssddEQ.gif",
	"https://media1.tenor.com/images/1e92c03121c0bd6688d17eef8d275ea7/tenor.gif?itemid=9920853",
	"https://data.whicdn.com/images/297125626/original.gif",
	"https://78.media.tumblr.com/755e34a2e227e153ba64d4a4f27848fe/tumblr_nahgozuIwM1rbnx7io1_500.gif",
	}


func Pat(b *discordgo.Session, m *discordgo.MessageCreate) {
	img := pats[rand.Intn(len(pats))]
	msg := &discordgo.MessageEmbed{}
	msg.Description = m.Author.Mention() + " has patted **" + strings.Replace(m.Message.Content, "!pat ", "", -1) + "**"
	msg.Image = &discordgo.MessageEmbedImage{URL:img}
	b.ChannelMessageSendEmbed(m.ChannelID, msg)
}
