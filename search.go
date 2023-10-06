// search.go
package main

import (
	"fmt"
	"log"
	"strings"
	"net/http"
	"github.com/PuerkitoBio/goquery"
)

func processGoogleSearch(searchTerm string, pages int) {
	for page := 0; page < pages; page++ {
		startResults := page * 100
		url := fmt.Sprintf("https://www.google.com/search?q=%s&start=%d", strings.Replace(searchTerm, " ", "+", -1), startResults)

		// Membuat permintaan HTTP ke Google
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("HTTP request failed: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			log.Fatalf("HTTP request failed with status code: %v", resp.StatusCode)
		}

		// Menganalisis halaman HTML
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatalf("Failed to parse HTML: %v", err)
		}

		// Mendapatkan URL dari hasil pencarian
		doc.Find("a").Each(func(index int, item *goquery.Selection) {
			href, exists := item.Attr("href")
			if exists && strings.HasPrefix(href, "/url?q=") {
				url := strings.TrimPrefix(href, "/url?q=")
				url = strings.Split(url, "&")[0] // Hapus parameter tambahan setelah URL

				// Ambil body HTML dari URL
				bodyText := getHTMLBody(url, &http.Client{})
				fmt.Printf("URL: %s\n", url)
				fmt.Printf("Body Text:\n%s\n", bodyText)
				fmt.Println("------")
			}
		})
	}
}
