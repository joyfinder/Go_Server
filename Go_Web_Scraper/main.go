package main

import (
	"fmt"
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

func buildBingUrls() {

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
		time.Sleep(time.Duration(backoff)*time.Second)
	}
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
