package main

import (
	"fmt"
	"strings"
	"time"
)


bingDomains = map[string]string{
	"com":""
}

var userAgents = []string{

}
type Search_Result struct{
	Result_rank int
	ResultURL string
	Result_title string
	Result_desc string
}

func randomUserAgent() string{

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

}

func scrapeClientRequest() {

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
