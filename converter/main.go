package converter

import (
	"encoding/json"
	"os"
)

func FromMap(inMap map[string]map[string]interface{}) (map[string]interface{}, map[string]interface{}) {
	parsedGenres := make(map[string]interface{})
	parsedAnimes := make(map[string]interface{})

	for url, anime := range inMap {
		parsedAnime := parseAnime(anime, parsedGenres, url)
		parsedAnimes[parsedAnime["slug"].(string)] = parsedAnime
	}

	return parsedAnimes, parsedGenres
}

func FromFile(inFile string) (map[string]interface{}, map[string]interface{}, error) {
	dat, err := os.ReadFile("/tmp/dat")
	if err != nil {
		return nil, nil, err
	}
	var result map[string]map[string]interface{}
	err = json.Unmarshal(dat, &result)
	if err != nil {
		return nil, nil, err
	}
	animes, genres := FromMap(result)
	return animes, genres, nil
}

func FromString(inString string) (map[string]interface{}, map[string]interface{}) {
	var result map[string]map[string]interface{}
	json.Unmarshal([]byte(inString), &result)

	return FromMap(result)
}

func FromBytes(inBytes []byte) (map[string]interface{}, map[string]interface{}) {
	var result map[string]map[string]interface{}
	json.Unmarshal(inBytes, &result)

	return FromMap(result)
}
