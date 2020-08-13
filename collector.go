package confluencego

import (
	"encoding/json"
	articlesgo "github.com/vitmovie/articlesgo"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const articlesPath = ".atlassian.net/wiki/rest/api/content?spaceKey=TS&type=page&start=0&end=99999"

var myClient = &http.Client{Timeout: 10 * time.Second}

func articlesUrl(domain string) string {
	return "https://" + domain + articlesPath
}

func getResponse(username, token, domain string) (interface{}, string) {
	client := &http.Client{}
	url := articlesUrl(domain)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(username, token)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	var jsonResp map[string]interface{}
	if err = json.Unmarshal(data, &jsonResp); err != nil {
		panic(err)
	}
	links := jsonResp["_links"].(map[string]interface{})
	return jsonResp["results"], links["base"].(string)
}

func getArticles(in interface{}, baseUrl string) []articlesgo.Article {
	var articles []articlesgo.Article
	switch results := in.(type) {
	case []interface {}:
		for _, result := range results {
			data := result.(map[string]interface{})
			a := articlesgo.Article{}
			confId, err := strconv.Atoi(data["id"].(string))
			if err != nil {
				panic(err)
			}
			a.ConfluenceID = confId
			a.Title = data["title"].(string)
			links := data["_links"].(map[string]interface{})
			a.Url = baseUrl + links["webui"].(string)
			articles = append(articles, a)
		}
	}
	return articles
}

func CollectArticles(username, token, domain string) []articlesgo.Article {
	jsonResp, baseUrl := getResponse(username, token, domain)
	articles := getArticles(jsonResp, baseUrl)
	return articles
}
