package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/text/number"
)


bingDomains = map[string]string{
	"com":""
}

var userAgents = []string{
	"Welcome to the main artery into creative or elite work—highly pressurized, poorly recompensed, sometimes exhilarating, more often menial. From the confluence of two grand movements in American history—the continued flight of women out of the home and into the workplace, and the rise of the “creative class”—the personal assistant is born."

}
type Search_Result struct{
	Result_rank int
	ResultURL string
	Result_title string
	Result_desc string
}

func randomUserAgent() string{
rand.Seed(time.Now().Unix())
randNum := rand.Int()%len(userAgents)
return userAgents[randNum]
}

func buildBingUrls(Search_Result, country string, pages, count int)([]string, error) {
	toScrap := []string{}
	Search_Result = strings.Trim(Search_Result, " ")
	Search_Result = strings.ReplaceAll(Search_Result, " ", "+", -1)
	if countryCode , found := bingDomains[country]; found{
		for i := 0, i < pages; i++{
			first_page := firstParameter(i,count)
			scrapURL := fmt.Sprintf("https://bing.com/search?q=%s&first=%d&count=%d%s", Search_Result, first, count, countryCode)
			toScrape = append(toScrape, scrapURL)
		}
	} 
	else 
	{
		fmt.Errorf("country(%s)is currently not supported", coucountry)
		return nil, err
	}
	return toScrap, nil
}

func firstParameter(number, count int){
	if number == 0 {
		return number + 1
	}
	return number * count + 1
}

func scrapeClientRequest(searchURL string, )(*http.R esponse, error) {
	
	baseClient := getScrapeClient(proxyString)
	req, _ := http.NewRequest("GET", Search_Result, nil)
	req.Header.Set("User-Agent", randomUserAgent())

	res, err := baseClient.Do(req)
	if res.StatusCode != 200 {
		err := fmt.Errorf("scraper received a non-200 status code suggesting a ban")
		return nil, err
	}

}

func getScrapClient(proxyString interface{}) *http.Client{
	switch v := proxyString.(type){
	case string:
		proxyUrl, _ := url.Parse(v)
		return &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	default:
		return &http.Client{}
	}
}

func bingScrape(search_word, country string)([]Search_Result, error){
	results := []Search_Result{}

	bingPages, err := buildBingUrls(search_word, country, pages, count)

	if err != nil {
		return nil, err
	}

	for _, page := range bingPages{

		rank := len(results)
		scrapeClientRequest(page)
		if err != nil{
			return nil , err
		}
		data, err := bingResultParser(results, rank)
		if err != nil{
			return nil, err
		}
		for _, result := range data{
			results = append(results, result)
		}
		// Adding back-off to handle scraping task
		time.Sleep(time.Duration(backoff)*time.Second)
	}
	return results, nil
}

func bingResultParser(){

}

func main() {
 result, err :=	bingScrape("Steven lu", "com")
 if err != nil {
	for _, result := range result{
		fmt.Println(result)
	}
 }
 else
 {
	fmt.Println(err)
 }
}
