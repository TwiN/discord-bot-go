package handler

import (
	"strings"
	"github.com/bwmarrin/discordgo"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func buildSearchUrl(searchTerm string) string {
	return fmt.Sprintf("https://www.google.com/search?q=%s&num=10&hl=en", strings.Replace(strings.Trim(searchTerm, " "), " ", "+", -1))
}

func googleRequest(searchURL string) (*http.Response, error) {
	baseClient := &http.Client{}
	req, _ := http.NewRequest("GET", searchURL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	res, err := baseClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func googleResultParser(response *http.Response) []string {
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	var results []string
	sel := doc.Find("div.g")
	for i := range sel.Nodes {
		item := sel.Eq(i)
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		link = strings.Trim(link, " ")
		if link != "" && link != "#" && strings.HasPrefix(link, "http") {
			result := link
			results = append(results, result)
		}
		if len(results) >= 3 {
			break
		}
	}
	return results
}

func GoogleScraper(searchTerm string) []string {
	res, err := googleRequest(buildSearchUrl(searchTerm))
	if err != nil {
		return nil
	}
	return googleResultParser(res)
}

func GoogleHandler(bot *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == bot.State.User.ID {
		return
	}
	if strings.HasPrefix(message.Content, "!google ") {
		var results = GoogleScraper(strings.Replace(message.Content, "!google ", "", 1))
		for _, url := range results {
			bot.ChannelMessageSend(message.ChannelID, url)
		}
	}
}

