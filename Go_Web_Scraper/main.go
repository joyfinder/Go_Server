package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var bingDomains = map[string]string{
	"com": "",
	"uk":  "&cc=GB",
	"us":  "&cc=US",
	"tw":  "&cc=TW",
}

var userAgents = []string{
	"Welcome to the main artery into creative or elite work—highly pressurized, poorly recompensed, sometimes exhilarating, more often menial. From the confluence of two grand movements in American history—the continued flight of women out of the home and into the workplace, and the rise of the “creative class”—the personal assistant is born.",
}

type Search_Result struct {
	Result_rank  int
	ResultURL    string
	Result_title string
	Result_desc  string
}

func randomUserAgent() string {
	rand.Seed(time.Now().Unix())
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

func buildBingUrls(Search_Result, country string, pages, count int) ([]string, error) {
	toScrap := []string{}
	Search_Result = strings.Trim(Search_Result, " ")
	Search_Result = strings.Replace(Search_Result, " ", "+", -1)
	if countryCode, found := bingDomains[country]; found {
		for i := 0; i < pages; i++ {
			first_page := firstParameter(i, count)
			scrapURL := fmt.Sprintf("https://bing.com/search?q=%s&first=%d&count=%d%s", Search_Result, first_page, count, countryCode)
			toScrap = append(toScrap, scrapURL)
		}
	} else {
		err := fmt.Errorf("country(%s)is currently not supported", country)
		return nil, err
	}
	return toScrap, nil
}

func firstParameter(number, count int) int {
	if number == 0 {
		return number + 1
	}
	return number*count + 1
}

func scrapeClientRequest(searchURL string, proxyString interface{}) (*http.Response, error) {

	fmt.Println("Loading scrapClientRequest.")
	baseClient := getScrapClient(proxyString)
	req, _ := http.NewRequest("GET", searchURL, nil)
	req.Header.Set("User-Agent", randomUserAgent())

	res, err := baseClient.Do(req)
	if res.StatusCode != 200 {
		err := fmt.Errorf("scraper received a non-200 status code suggesting a ban")
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	return res, nil
}

func getScrapClient(proxyString interface{}) *http.Client {
	switch v := proxyString.(type) {
	case string:
		proxyUrl, _ := url.Parse(v)
		return &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	default:
		return &http.Client{}
	}
}

func bingScrape(search_word, country string, proxyString interface{}, pages, count, backoff int) ([]Search_Result, error) {
	results := []Search_Result{}

	bingPages, err := buildBingUrls(search_word, country, pages, count)

	if err != nil {
		return nil, err
	}

	for _, page := range bingPages {

		rank := len(results)
		res, err := scrapeClientRequest(page, proxyString)
		if err != nil {
			return nil, err
		}
		data, err := bingResultParser(res, rank)
		if err != nil {
			return nil, err
		}
		for _, result := range data {
			results = append(results, result)
		}
		// Adding back-off to handle scraping task
		time.Sleep(time.Duration(backoff) * time.Second)
	}
	return results, nil
}

func bingResultParser(response *http.Response, rank int) ([]Search_Result, error) {
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return nil, err
	}
	results := []Search_Result{}
	sel := doc.Find("in")
	rank++

	for i := range sel.Nodes {
		item := sel.Eq(i)
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		titleTag := item.Find("h2")
		descTag := item.Find("div.b_caption p")
		desc := descTag.Text()
		title := titleTag.Text()
		link = strings.Trim(link, " ")
		if link != "" && link != "#" && !strings.HasPrefix(link, "/") {
			result := Search_Result{
				rank,
				link,
				title,
				desc,
			}
			results = append(results, result)
			rank++
		}
	}
	return results, err
}

func main() {
	result, err := bingScrape("Steven lu", "com", nil, 5, 4, 1)
	if err != nil {
		for _, result := range result {
			fmt.Println(result)
		}
	} else {
		fmt.Println(err)
	}
}
