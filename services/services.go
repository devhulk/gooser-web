package services

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/fatih/color"
)

type SiteCheck struct {
	Name   string
	URL    string
	Exists bool
}

const file_uri = "https://raw.githubusercontent.com/WebBreacher/WhatsMyName/main/wmn-data.json"

func GetSiteMap(username string) ([]SiteCheck, error) {

	resp, err := http.Get(file_uri)
	if err != nil {
		return []SiteCheck{}, err
	}

	defer resp.Body.Close()

	byteValue, _ := io.ReadAll(resp.Body)

	var result WhatsMyName

	err3 := json.Unmarshal(byteValue, &result)
	if err3 != nil {
		log.Fatalln(err3)
		return []SiteCheck{}, err
	}

	hits, err := CheckSites(result, username)
	if err != nil {
		log.Fatalln(err)
	}

	return hits, nil

}

func CheckSites(w WhatsMyName, u string) ([]SiteCheck, error) {

	var s []SiteCheck

	for _, v := range w.Sites {

		topten := []string{"TikTok", "Twitter", "Facebook", "Instagram", "YouTube User2", "Telegram", "Pinterest", "Snapchat", "Reddit", "Twitch"}

		if slices.Contains(topten, v.Name) {

			uri := strings.Replace(v.URICheck, "{account}", u, 1)
			resp, err := http.Get(uri)
			if err != nil {
				log.Println(err)
				continue
			}

			if resp.StatusCode == v.ECode {
				color.Green("Hit: %v\n", uri)
				check := SiteCheck{
					Name:   v.Name,
					URL:    uri,
					Exists: true,
				}
				s = append(s, check)

			} else if resp.StatusCode == v.MCode {
				color.Red("Miss: %v\n", uri)
				check := SiteCheck{
					Name:   v.Name,
					URL:    uri,
					Exists: false,
				}
				s = append(s, check)
				continue
			}

		}
	}

	return s, nil

}
