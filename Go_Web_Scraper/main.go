package main

import "fmt"


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

func bingScrape(){

}

func bingResultParser(search_word, country string)([]Search_Result){

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