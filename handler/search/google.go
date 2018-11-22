package search

import (
	"strings"
	"github.com/bwmarrin/discordgo"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	Constants "../../global"
	"../../cache"
)


func GoogleSearch(bot *discordgo.Session, message *discordgo.MessageCreate, query string) {
	const COMMAND = Constants.COMMAND_PREFIX + "google"

	if len(query) == 0 {
		bot.ChannelMessageSend(message.ChannelID, "**USAGE:** `"+COMMAND+" <search terms>`")
	} else {
		bot.UpdateStatus(1, "| :mag_right: '"+query+"' on Google")
		var results = googleSearchScraper(query)
		cache.Put("google", query, results)
		for _, url := range results {
			bot.ChannelMessageSend(message.ChannelID, url)
		}
		bot.UpdateStatus(0, "")
	}
}


func googleSearchScraper(searchTerm string) []string {
	res, err := fetchGoogleSearchPage(buildGoogleSearchUrl(searchTerm))
	if err != nil {
		return nil
	}
	return parseGoogleSearchResult(res)
}


func buildGoogleSearchUrl(searchTerm string) string {
	return fmt.Sprintf("https://www.google.com/search?q=%s&num=5&hl=en&safe=active", strings.Replace(strings.Trim(searchTerm, " "), " ", "+", -1))
}


func fetchGoogleSearchPage(url string) (*http.Response, error) {
	baseClient := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	res, err := baseClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}


func parseGoogleSearchResult(response *http.Response) []string {
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
		if len(results) >= 2 {
			break
		}
	}
	return results
}
