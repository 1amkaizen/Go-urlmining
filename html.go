// html.go
package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func getHTMLBody(url string, client *http.Client) string {
	resp, err := client.Get(url)
	if err != nil {
		log.Printf("Failed to retrieve URL %s: %v", url, err)
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to retrieve URL %s. Status code: %v", url, resp.StatusCode)
		return ""
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("Failed to parse HTML for URL %s: %v", url, err)
		return ""
	}

	// Menghapus elemen-elemen yang tidak diinginkan (script, style, iframe, dll.)
	doc.Find("script, style, iframe, link").Each(func(index int, item *goquery.Selection) {
		item.Remove()
	})

	// Mendapatkan teks dari elemen body
	bodyText := doc.Find("body").Text()
	return bodyText
}
