package animeflv

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/kuronosu/schema_scraper/pkg/scrape"
	"github.com/kuronosu/schema_scraper/pkg/utils"
)

const BASE_ANIME_FLV = "https://www3.animeflv.net"

func GetLastPageNumber() (int, error) {
	lastPageUrlResults := scrape.ScrapeListFlat(AnimeFLVBrowseSchema, BASE_ANIME_FLV+"/browse")
	if len(lastPageUrlResults) != 1 {
		return 0, fmt.Errorf("more than one last page url found")
	}
	lastPageUrl := lastPageUrlResults[0]
	lastPageNumber, err := strconv.Atoi(strings.ReplaceAll(lastPageUrl, "/browse?page=", ""))
	if err != nil {
		return 0, err
	}
	return lastPageNumber, nil
}

func GetAnimesUrlsAsync() ([]string, error) {
	lastPageNumber, err := GetLastPageNumber()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	urls := make([]string, 0)
	var wg sync.WaitGroup
	for i := 1; i <= lastPageNumber; i++ {
		wg.Add(1)
		go func(pageNumber int) {
			defer wg.Done()
			urls = append(urls, scrape.ScrapeListFlat(AnimeFLVSchema, BASE_ANIME_FLV+"/browse?page="+strconv.Itoa(pageNumber))...)
		}(i)
	}
	wg.Wait()
	return urls, nil
}

func ScrapeDirectory() (*scrape.MemoryDetailsCollector, map[string]string) {
	details := scrape.NewMemoryDetailsCollector()
	urls, err := GetAnimesUrlsAsync()
	if err != nil {
		panic(err)
	}
	urls = utils.RemoveDuplicatesUrls(urls)
	options := scrape.ScrapeDetailsOptions{
		Async:            true,
		Schema:           AnimeFLVSchema,
		URLs:             urls,
		DetailsCollector: details,
		BatchSize:        500,
	}
	errors := scrape.ScrapeDetails(options)
	okCount := len(details.Items)
	errCount := len(errors)
	trueErrors := utils.GetErrorUrlsWithoutNotFound(errors)
	fmt.Println(
		"[OK]", okCount,
		"\n[Errors]", errCount,
		"\n[Total]", okCount+errCount,
		"\n[Errors without 404]", len(trueErrors),
	)
	return details, trueErrors
}
