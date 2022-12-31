package storage

import (
	"fmt"

	"github.com/kuronosu/deguvon-scraper/converter"
)

type Directory map[string]map[string]interface{}

type D = Directory

func (d Directory) DetaStore() []ErrorsMap {
	parsedAnimes, parsedGenres := converter.FromMap(d)
	sm, err := NewDetaStorageManager(true)
	if err != nil {
		return []ErrorsMap{{err, nil}}
	}
	if errs := sm.Animes().UploadManyAsync(parsedAnimes); len(errs) > 0 {
		fmt.Println(errs)
	}
	if errs := sm.Genres().UploadManyAsync(parsedGenres); len(errs) > 0 {
		fmt.Println(errs)
	}
	return nil
}
