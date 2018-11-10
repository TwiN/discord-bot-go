package handler

import (
	"strings"
	"github.com/bwmarrin/discordgo"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	Constants "../global"
	Cache "../cache"
)


func YoutubeSearchHandler(bot *discordgo.Session, message *discordgo.MessageCreate) {
	const COMMAND = Constants.COMMAND_PREFIX + "youtube"
	if message.Author.ID == bot.State.User.ID {
		return
	}
	if strings.HasPrefix(message.Content, COMMAND) {
		var query = strings.Trim(strings.Replace(message.Content, COMMAND, "", 1), " ")
		if Cache.Has("youtube", query) {
			for _, url := range Cache.Get("youtube", query) {
				bot.ChannelMessageSend(message.ChannelID, "[cached] " + url)
			}
			return
		}
		if len(query) == 0 {
			bot.ChannelMessageSend(message.ChannelID, "**USAGE:** `" + COMMAND + " <search terms>`")
		} else {
			bot.UpdateStatus(1, "| :mag_right: '" + query + "' on Youtube")
			var results= YoutubeSearchScraper(query)
			Cache.Put("youtube", query, results)
			for _, url := range results {
				bot.ChannelMessageSend(message.ChannelID, url)
			}
			bot.UpdateStatus(0, "")
		}
	}
}


func YoutubeSearchScraper(searchTerm string) []string {
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
		if len(results) >= 3 {
			break
		}
	}
	return results
}
