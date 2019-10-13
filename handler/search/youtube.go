package search

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/TwinProduction/discord-bot-go/cache"
	"github.com/TwinProduction/discord-bot-go/global"
	"github.com/bwmarrin/discordgo"
	"net/http"
	"strings"
)

func YoutubeSearch(bot *discordgo.Session, message *discordgo.MessageCreate, query string) {
	const Command = global.CommandPrefix + "youtube"
	if len(query) == 0 {
		bot.ChannelMessageSend(message.ChannelID, "**USAGE:** `"+Command+" <search terms>`")
	} else {
		bot.UpdateStatus(1, "| :mag_right: '"+query+"' on Youtube")
		var results = youtubeSearchScraper(query)
		cache.Youtube.Put(query, results)
		for _, url := range results {
			bot.ChannelMessageSend(message.ChannelID, url)
		}
		bot.UpdateStatus(0, "")
	}
}

func youtubeSearchScraper(searchTerm string) []string {
	res, err := fetchYoutubeSearchPage(buildYoutubeSearchUrl(searchTerm))
	if err != nil {
		return nil
	}
	return parseYoutubeSearchResult(res)
}

func buildYoutubeSearchUrl(searchTerm string) string {
	return fmt.Sprintf("https://www.youtube.com/results?search_query=%s", strings.Replace(strings.Trim(searchTerm, " "), " ", "+", -1))
}

func fetchYoutubeSearchPage(url string) (*http.Response, error) {
	baseClient := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	res, err := baseClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func parseYoutubeSearchResult(response *http.Response) []string {
	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil
	}
	var results []string
	sel := doc.Find("div#img-preload img")
	for i := range sel.Nodes {
		item := sel.Eq(i)
		thumbnailUrl, _ := item.Attr("src")
		parts := strings.Split(thumbnailUrl, "/")
		videoId := parts[4]
		if len(videoId) > 20 {
			continue
		}
		link := "https://www.youtube.com/watch?v=" + videoId
		if link != "" && link != "#" && strings.HasPrefix(link, "http") {
			result := link
			results = append(results, result)
		}
		if len(results) >= 2 {
			break
		}
	}
	return results
}
