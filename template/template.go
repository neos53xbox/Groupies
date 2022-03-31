package template

import (
	"html/template"
	"net/http"
	"strconv"

	API "groupetracker/api"
)

type DataMainPage struct {
	FirstLineArtist  []API.GroupeTracker_Artist
	SecondLineArtist []API.GroupeTracker_Artist
	ThirdLineArtist  []API.GroupeTracker_Artist
	FourthLineArtist []API.GroupeTracker_Artist
	FifthLineArtist  []API.GroupeTracker_Artist
	SixthLineArtist  []API.GroupeTracker_Artist
	SearchArtist     []API.GroupeTracker_Artist
	NextButtonJS     template.JS
	BackButtonJS     template.JS
}

type DataArtistPage struct {
	DeezerApi        API.DeezerIndex
	Artist           []API.GroupeTracker_Artist
	SearchArtist     []API.GroupeTracker_Artist
	GeocodeGoogleAPI API.Return_GeocodeGoogleAPI
}

func HandleTemplateMainPage(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("./index.html"))
	var data = DataMainPage{
		FirstLineArtist:  API.ArtistsLists()[0:5],
		SecondLineArtist: API.ArtistsLists()[5:9],
		ThirdLineArtist:  API.ArtistsLists()[9:14],
		FourthLineArtist: API.ArtistsLists()[14:18],
		FifthLineArtist:  API.ArtistsLists()[18:23],
		SixthLineArtist:  API.ArtistsLists()[23:27],
		SearchArtist:     API.ArtistsLists(),
		NextButtonJS:     template.JS(PageButtonJS(API.ArtistsLists()[27:])),
		BackButtonJS:     template.JS(PageButtonJS(API.ArtistsLists()[0:27]))}
	tmpl.Execute(w, data)
}

func HandleTemplateArtistPage(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("./artist.html"))

	for id := range API.ArtistsList() {
		if r.URL.Path == "/artist/"+strconv.Itoa(id+1) {
			var data = DataArtistPage{
				DeezerApi:        API.DeezerAPI()[id : id+1],
				Artist:           API.ArtistsLists()[id : id+1],
				GeocodeGoogleAPI: API.GeocodeGoogleAPI(API.RelationsArtistList(id)),
				SearchArtist:     API.ArtistsLists(),
			}
			tmpl.Execute(w, data)
			break
		}
	}
}

func PageButtonJS(input []API.GroupeTracker_Artist) template.JS {
	output := ``
	for id, artist := range input {
		output += `
		var A` + strconv.Itoa(id+1) + `= [document.getElementById("I` + strconv.Itoa(id+1) + `"), document.getElementById("A` + strconv.Itoa(id+1) + `"), document.getElementById("AName` + strconv.Itoa(id+1) + `")];
		A` + strconv.Itoa(id+1) + `[0].style.backgroundImage = "url(` + artist.Image + `)";
		A` + strconv.Itoa(id+1) + `[1].href = "/artist/` + strconv.Itoa(artist.ID) + `";
		A` + strconv.Itoa(id+1) + `[2].innerHTML = "` + artist.Name + `";
		`
	}
	return template.JS(output)
}
