package roleplay

import (
	"math/rand"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var greetings = []string{
		"https://media1.tenor.com/images/90975d9b5cc7a3b48147514308fd1e17/tenor.gif?itemid=8390761",
		"https://image.myanimelist.net/ui/G-Sm6d0qIwQxUGHIp-m2WDtxdXe6XQHxhSWRtaSrIvuRbzghf_SAS_kNOM09Afi_0OJQSOa3KVo8nbVqnbRm4_2Et2wQeSeehopIp0q-FIXDy-HBdTjCN9qyvWRwf6vw",
	}


func Greet(b *discordgo.Session, m *discordgo.MessageCreate) {
	img := greetings[rand.Intn(len(greetings))]
	msg := &discordgo.MessageEmbed{}
	msg.Description = m.Author.Mention() + " has greeted **" + strings.Replace(m.Message.Content, "!greet ", "", -1) + "**"
	msg.Image = &discordgo.MessageEmbedImage{URL:img}
	b.ChannelMessageSendEmbed(m.ChannelID, msg)
}
