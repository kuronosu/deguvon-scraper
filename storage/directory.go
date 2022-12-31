package storage

import (
	"encoding/json"
	"fmt"

	"github.com/kuronosu/deguvon-scraper/converter"
	"github.com/kuronosu/schema_scraper/pkg/utils"
)

type RawDirectory map[string]map[string]interface{}

type RD = RawDirectory

func (d RawDirectory) DetaStore() []ErrorsMap {
	d, err := d.Clone()
	if err != nil {
		return []ErrorsMap{{err, nil}}
	}
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

func (d RawDirectory) SaveToFiles() error {
	d, err := d.Clone()
	if err != nil {
		return err
	}
	animes, genres := converter.FromMap(d)
	if err := utils.WriteJson(genres, "genres.json", false); err != nil {
		return err
	}
	return utils.WriteJson(animes, "animes.json", false)
}

func (d RawDirectory) Clone() (RawDirectory, error) {
	data, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	var clone RawDirectory
	err = json.Unmarshal(data, &clone)
	if err != nil {
		return nil, err
	}
	return clone, nil
}
