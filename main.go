package main

import (
	"github.com/kuronosu/deguvon-scraper/animeflv"
	"github.com/kuronosu/deguvon-scraper/storage"
	"github.com/kuronosu/deguvon-scraper/utils"
	"github.com/kuronosu/schema_scraper/pkg/scrape"
)

func main() {
	utils.MeasureTime("Get directory and store in Deta", func() {
		scrape.SetVerbose(true)
		directory := storage.RD(animeflv.ScrapeDirectory().Items)
		directory.SaveToFiles()
		directory.DetaStore()
	})
}
