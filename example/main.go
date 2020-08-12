package main

import (
	"fmt"
	confCollector "github.com/VitMovie/Confluence-Collector"
)

const (
	username = "username"
	token  = "token"
	url = "https://linkerbot.atlassian.net/wiki/rest/api/content?spaceKey=TS&type=page&start=0&end=99999"
)

func main() {
	// Collect articles with confluence id and title
	art := confCollector.CollectArticles(username, token, url)
	fmt.Println(art)
}
