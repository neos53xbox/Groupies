package API

import (
	"encoding/json"
	"io/ioutil"
)

type Deezer struct {
	Index []struct {
		ID      int    `json:"id"`
		Image   string `json:"image"`
		Extrait string `json:"extrait"`
	}
}

type DeezerIndex []struct {
	ID      int    `json:"id"`
	Image   string `json:"image"`
	Extrait string `json:"extrait"`
}

func DeezerAPI() DeezerIndex {
	DeezerAPI := Deezer{}
	APIJson, _ := ioutil.ReadFile("./api/deezer-api.json")
	json.Unmarshal([]byte(APIJson), &DeezerAPI)
	return DeezerAPI.Index
}
