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


func UrbanDictionarySearchHandler(bot *discordgo.Session, message *discordgo.MessageCreate) {
	const COMMAND = Constants.COMMAND_PREFIX + "urban"
	if message.Author.ID == bot.State.User.ID {
		return
	}
	if strings.HasPrefix(message.Content, COMMAND) {
		var query = strings.Trim(strings.Replace(message.Content, COMMAND, "", 1), " ")
		if Cache.Has("urban", query) {
			for _, url := range Cache.Get("urban", query) {
				bot.ChannelMessageSend(message.ChannelID, "[cached] " + url)
			}
			return
		}
		if len(query) == 0 {
			bot.ChannelMessageSend(message.ChannelID, "**USAGE:** `" + COMMAND + " <search terms>`")
		} else {
			bot.UpdateStatus(1, "| :mag_right: '" + query + "' on UrbanDictionary")
			result := "**Urban Dictionary search result for `" + query + "`:**" + UrbanDictionarySearchScraper(query)
			Cache.Put("urban", query, []string{result})
			bot.ChannelMessageSend(message.ChannelID, result)
			bot.UpdateStatus(0, "")
		}
	}
}


func UrbanDictionarySearchScraper(searchTerm string) string {
	res, _ := fetchUrbanDictionarySearchPage(buildUrbanDictionarySearchUrl(searchTerm))
	return parseUrbanDictionarySearchResult(res)
}


func buildUrbanDictionarySearchUrl(searchTerm string) string {
	return fmt.Sprintf("https://www.urbandictionary.com/define.php?term=%s", strings.Replace(strings.Trim(searchTerm, " "), " ", "+", -1))
}


func fetchUrbanDictionarySearchPage(url string) (*http.Response, error) {
	baseClient := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	res, err := baseClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}


func parseUrbanDictionarySearchResult(response *http.Response) string {
	doc, _ := goquery.NewDocumentFromReader(response.Body)
	sel := doc.Find("div.meaning")
	for i := range sel.Nodes {
		item := sel.Eq(i)
		definition := item.Text()
		definition = strings.Trim(definition, " ")
		if definition != "" {
			return definition
		}
	}
	return "¯\\_(ツ)_/¯"
}
