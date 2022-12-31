package main

import (
	"github.com/kuronosu/deguvon-scraper/animeflv"
	"github.com/kuronosu/deguvon-scraper/storage"
	"github.com/kuronosu/deguvon-scraper/utils"
	"github.com/kuronosu/schema_scraper/pkg/scrape"
	// "github.com/kuronosu/deguvon-scraper/converter"
	// schemaUtils "github.com/kuronosu/schema_scraper/pkg/utils"
)

func main() {
	utils.MeasureTime("Get directory and store in Deta", func() {
		scrape.SetVerbose(true)
		directory := animeflv.ScrapeDirectory()
		storage.D(directory.Items).DetaStore()
		// animes, genres := converter.FromMap(directory.Items)
		// schemaUtils.WriteJson(animes, "animes.json", false)
		// schemaUtils.WriteJson(genres, "genres.json", false)
	})
}
