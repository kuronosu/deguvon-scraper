package converter

import (
	"fmt"
	"strconv"
	"strings"
)

func floatOrInt(raw string) interface{} {
	if v, err := strconv.Atoi(raw); err == nil {
		return v
	}
	if v, err := strconv.ParseFloat(raw, 64); err == nil {
		return v
	}
	return nil
}

func formatEpisodes(raw, url string) []interface{} {
	if raw == "" {
		return nil
	}
	episodes := make([]interface{}, 0)
	for _, episode := range strings.Split(raw, "],") {
		tmp := strings.Split(strings.ReplaceAll(strings.ReplaceAll(episode, "[", ""), "]", ""), ",")
		flvid, err := strconv.Atoi(tmp[1])
		number := floatOrInt(tmp[0])
		if err != nil || number == nil {
			fmt.Println(url, " - ", err)
		} else {
			episodes = append(episodes, map[string]interface{}{
				"flvid":  flvid,
				"number": number,
			})
		}
	}
	return episodes
}

func delCharAt(sample string, index int) string {
	s := []rune(sample)
	return string(append(s[0:index], s[index+1:]...))
}

func getExtraInfo(raw string) []string {
	extraInfo := delCharAt(delCharAt(raw, len(raw)-1), 0)
	return strings.Split(fmt.Sprintf("%v", extraInfo), "\",\"")
}

func parseRelated(anime map[string]interface{}) []interface{} {
	newRelateds := make([]interface{}, 0)
	if anime["related"] != nil {
		for _, related := range anime["related"].([]interface{}) {
			if fmt.Sprintf("%v", related.(map[string]interface{})["name"]) != "" {
				newRelateds = append(newRelateds, related)
			}
		}
	}
	return newRelateds
}

func parseAnime(anime, genres map[string]interface{}, url string) map[string]interface{} {
	extraInfo := getExtraInfo(fmt.Sprintf("%v", anime["raw_data"]))

	flvID, err := strconv.Atoi(extraInfo[0])
	if err != nil {
		fmt.Println("Invalid flvid", url, extraInfo[0])
	}
	anime["flvid"] = flvID
	anime["slug"] = extraInfo[2]
	if len(extraInfo) > 3 {
		anime["next_episode"] = extraInfo[3]
	}
	anime["url"] = url

	anime["episodes"] = formatEpisodes(fmt.Sprintf("%v", anime["raw_episodes"]), url)

	delete(anime, "raw_episodes")
	delete(anime, "raw_data")

	animeGenres := make([]string, 0)
	for _, genre := range anime["genres"].([]interface{}) {
		genre := genre.(map[string]interface{})
		name := fmt.Sprintf("%v", genre["name"])
		url := fmt.Sprintf("%v", genre["url"])
		if _, ok := (genres)[name]; !ok {
			genres[name] = map[string]interface{}{
				"key": name,
				"url": url,
			}
		}
		animeGenres = append(animeGenres, name)
	}
	anime["genres"] = animeGenres
	anime["key"] = anime["slug"]
	anime["related"] = parseRelated(anime)
	return anime
}
