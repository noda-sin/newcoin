package main

import (
	"fmt"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type coin struct {
	Name   string
	Symbol string
	Price  float64
	URL    string
}

func getNewCoins() (newCoins []coin, err error) {
	newCoins = make([]coin, 0)
	doc, err := goquery.NewDocument("https://coinmarketcap.com/new/")
	if err != nil {
		fmt.Print("url scarapping failed")
		return
	}
	doc.Find("table#trending-recently-added > tbody > tr").Each(func(_ int, row *goquery.Selection) {
		url, _ := row.Find("td.currency-name > a").First().Attr("href")
		name := row.Find("td.currency-name > a").First().Text()
		symbol := row.Find("td.text-left").First().Text()
		priceStr, _ := row.Find("td > a.price").Attr("data-usd")
		price, err := strconv.ParseFloat(priceStr, 32)
		if err != nil {
			return
		}
		coin := coin{Name: name, Symbol: symbol, Price: price, URL: url}
		newCoins = append(newCoins, coin)
	})
	return
}
