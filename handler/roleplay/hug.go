package roleplay

import (
	"math/rand"
	"github.com/bwmarrin/discordgo"
	"strings"
)

var hugs = []string{
		"https://media1.tenor.com/images/49a21e182fcdfb3e96cc9d9421f8ee3f/tenor.gif?itemid=3532079",
		"https://vignette.wikia.nocookie.net/yandere-simulator/images/d/d6/Anime-hug-gif-16.gif/revision/latest?cb=20180223151556",
		"https://media.giphy.com/media/l2QDM9Jnim1YVILXa/giphy.gif",
		"https://media.giphy.com/media/xJlOdEYy0r7ZS/giphy.gif",
		"http://s.orzzzz.com/news/a3/85//586604e56348c.gif",
	}


func Hug(b *discordgo.Session, m *discordgo.MessageCreate) {
	img := hugs[rand.Intn(len(hugs))]
	msg := &discordgo.MessageEmbed{}
	msg.Description = m.Author.Mention() + " has hugged **" + strings.Replace(m.Message.Content, "!hug ", "", -1) + "**"
	msg.Image = &discordgo.MessageEmbedImage{URL:img}
	b.ChannelMessageSendEmbed(m.ChannelID, msg)
}
