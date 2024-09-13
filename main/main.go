package main

import (
	// importing Colly
	"fmt"

	"github.com/gocolly/colly"
)

type Product struct {
	Url, Image, Name, Price string
}

var products []Product

func main() {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36"),
		colly.AllowedDomains("www.alaskaair.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.Visit("https://www.alaskaair.com")
}
