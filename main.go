package main

import (
	"fmt"

	"github.com/kuronosu/deguvon-scraper/animeflv"
	"github.com/kuronosu/deguvon-scraper/storage"
	"github.com/kuronosu/deguvon-scraper/utils"
	"github.com/kuronosu/schema_scraper/pkg/scrape"
	scrapeUtils "github.com/kuronosu/schema_scraper/pkg/utils"
)

func confirmUpload() bool {
	var input string
	fmt.Print("Upload to Deta? (y/n): ")
	fmt.Scanln(&input)
	return input == "y" || input == "Y"
}

func main() {
	utils.MeasureTime("Get directory and store in Deta", func() {
		scrape.SetVerbose(true)
		_directory, errors := animeflv.ScrapeDirectory()
		directory := storage.RD(_directory.Items)
		directory.SaveToFiles()
		scrapeUtils.WriteJson(errors, "true_anime_errors.json", true)
		fmt.Println("Directory saved to files")
		if confirmUpload() {
			directory.DetaStore()
		}
	})
}
