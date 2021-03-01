package models

type Series struct {
	Model
	Name          string `json:"name"`
	SeriesNumber  int    `json:"int"`
	EpisodeNumber int    `json:"int"`
	AssetPath     string `json:"asset_path"`
}

/*

// Movie
{
	"name": "Titanic",
	"description": "Movie where video syncs",
}

// Series
{
	"id": "grand_tour",
	"seasons": [
		1: {
			episdoes: {
				1: {
					name: "Going to middle east",
					assets: "test/this/out"
				},
				2: {
					name: "Going to africa",
					assets: "test/this/out"
				}
			}
		}
	]
}

// Mini-Series
{
	"id": "flight_atendant",
	"episodes": {
		1: {
			name: "blash",
			assets: "test/this/out",
		}
	}
}

// Content
{
	"name": "Titanic | Flight Atendant | Stranger Things",
	"descriptions": "Some description",
	"assets": "test/this/out",
	"type": "Movie Collection | Movie | Series | Mini Series"
}

*/
