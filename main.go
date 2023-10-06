// main.go
package main

import (
	"flag"
	"fmt"
)

func main() {
	// Definisikan flag
	searchTerm := flag.String("q", "", "Search Query")
	pages := flag.Int("p", 1, "Number of pages")
	help := flag.Bool("h", false, "Show help")

	// Parse flag
	flag.Parse()

	// Tampilkan bantuan jika -h digunakan
	if *help {
		flag.Usage()
		return
	}

	// Pastikan search term telah diisi
	if *searchTerm == "" {
		fmt.Println("Search term is required. Use -h for help.")
		return
	}

	processGoogleSearch(*searchTerm, *pages)
}

