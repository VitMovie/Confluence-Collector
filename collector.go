package confluencego

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type article map[string]string

var myClient = &http.Client{Timeout: 10 * time.Second}

func getResponse(username, token, url string) interface{} {
	client := &http.Client{}
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
	return jsonResp["results"]
}

func getArticles(in interface{}) []article {
	var articles []article
	switch results := in.(type) {
	case []interface {}:
		for _, result := range results {
			data := result.(map[string]interface{})
			a := make(article)
			a["id"] = data["id"].(string)
			a["title"] = data["title"].(string)
			articles = append(articles, a)
		}
	}
	return articles
}

func CollectArticles(username, token, url string) []article {
	jsonResp := getResponse(username, token, url)
	articles := getArticles(jsonResp)
	return articles
}
