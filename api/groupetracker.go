package API

import (
	"encoding/json"
	"io/ioutil"
)

type GroupeTracker struct {
	Artists   string `json:"artists"`
	Locations string `json:"locations"`
	Dates     string `json:"dates"`
	Relation  string `json:"relation"`
}

type GroupeTracker_Artist struct {
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

type GroupeTracker_LocationsAll struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

type GroupeTracker_LocationsArtist struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type GroupeTracker_DatesAll struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	}
}

type GroupeTracker_DatesArtist struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type GroupeTracker_RelationsAll struct {
	Index []struct {
		ID             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

type GroupeTracker_RelationsArtist struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func GroupeTrackerAPI() GroupeTracker {
	GroupeTrackerAPI := GroupeTracker{}
	APIJson, _ := ioutil.ReadFile("./api/groupetracker-api.json")
	json.Unmarshal([]byte(APIJson), &GroupeTrackerAPI)
	return GroupeTrackerAPI
}

func ArtistsLists() []GroupeTracker_Artist {
	ArtistsList := []GroupeTracker_Artist{}
	URL := input(GroupeTrackerAPI().Artists)
	bytes, _ := URL.ReadURL()

	json.Unmarshal(bytes, &ArtistsList)
	return ArtistsList
}

func RelationsArtistLists(ID int) map[string][]string {
	RelationsArtistList := GroupeTracker_RelationsArtist{}
	URL := input(ArtistsList()[ID].Relations)
	bytes, _ := URL.ReadURL()

	json.Unmarshal(bytes, &RelationsArtistList)
	return RelationsArtistList.DatesLocations
}

func LocationsArtistLists(ID int) []string {
	LocationsArtistList := GroupeTracker_LocationsArtist{}
	URL := input(ArtistsList()[ID].Locations)
	bytes, _ := URL.ReadURL()

	json.Unmarshal(bytes, &LocationsArtistList)
	return LocationsArtistList.Locations
}
