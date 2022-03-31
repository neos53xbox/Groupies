package API

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type API struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConsertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type LocationsAll struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type LocationsArtist struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type DatesAll struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	}
}

type DatesArtist struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type RelationsAll struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type RelationsArtist struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func DataAPI() API {
	DataAPI := API{}
	APIJson, _ := ioutil.ReadFile("./api/api.json")
	json.Unmarshal([]byte(APIJson), &DataAPI)
	return DataAPI
}

func ReadURL(URL string) []byte {
	resp, _ := http.Get(URL)
	jsonData, _ := ioutil.ReadAll(resp.Body)
	return jsonData
}

func ArtistsList() []Artist {
	ArtistsList := []Artist{}
	json.Unmarshal(ReadURL(DataAPI().Artists), &ArtistsList)
	return ArtistsList
}

func RelationsArtistList(ID int) map[string][]string {
	RelationsArtistList := RelationsArtist{}
	json.Unmarshal(ReadURL(ArtistsList()[ID].Relations), &RelationsArtistList)
	return RelationsArtistList.DatesLocations
}

func LocationsArtistList(ID int) []string {
	LocationsArtistList := LocationsArtist{}
	json.Unmarshal(ReadURL(ArtistsList()[ID].Locations), &LocationsArtistList)
	return LocationsArtistList.Locations
}
