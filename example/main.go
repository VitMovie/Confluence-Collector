package main

import (
	"fmt"
	confCollector "github.com/VitMovie/Confluence-Collector"
)

const (
	username = "username"
	token  = "token"
	domain = "linkerbot"
)

func main() {
	// Collect articles with confluence id and title
	articles := confCollector.CollectArticles(username, token, domain)
	fmt.Println(articles)
}
