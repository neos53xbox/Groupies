package API

import (
	"encoding/json"
)

type GeocodeGoogle struct { // Struct GeocodeAPI from Google
	Results []struct {
		Formatted_Address string `json:"formatted_address"`
		Geometry          struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
}

type GeocodeGoogle_Location struct { // Struct GeocodeAPI for location with
	Lat float64 `json:"lat"` // with coordinates on the map
	Lng float64 `json:"lng"`
}

type Return_GeocodeGoogleAPI struct {
	GeocodeGoogle_Location []GeocodeGoogle_Location
	GeocodeGoogle_Relation map[string][]string
}

func GeocodeGoogleAPI(relation map[string][]string) Return_GeocodeGoogleAPI { // The function processes the Geocode API and displays
	URL := "https://maps.googleapis.com/maps/api/geocode/json" // the structure of cities with dates and geolocation
	Address := "?address="
	APIKey := "&key=AIzaSyA4ZaCtIOqBH4OwpiPw3zMwuWFUvwE8GC4"
	Language := "&language=en"

	GeocodeGoogleAPI := GeocodeGoogle{}
	GeocodeGoogleAPI_RelationList := make(map[string][]string)
	GeocodeGoogleAPI_LocationList := []GeocodeGoogle_Location{}

	for city, dates := range relation {
		FullURL := input(URL + Address + city + Language + APIKey)
		bytes, _ := FullURL.ReadURL()
		json.Unmarshal(bytes, &GeocodeGoogleAPI)
		GeocodeGoogleAPI_RelationList[GeocodeGoogleAPI.Results[0].Formatted_Address] = dates
		GeocodeGoogleAPI_LocationList = append(GeocodeGoogleAPI_LocationList, GeocodeGoogleAPI.Results[0].Geometry.Location)
	}

	return Return_GeocodeGoogleAPI{
		GeocodeGoogle_Location: GeocodeGoogleAPI_LocationList,
		GeocodeGoogle_Relation: GeocodeGoogleAPI_RelationList,
	}
}
