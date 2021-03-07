package models

type MiniSeries struct {
	Model
	MediaMetadata
	EpisodeNumber int  `json:"episode_number"`
	ContentID     uint `json:"content_id"`
}

func (MiniSeries) TableName() string {
	return "mini_series"
}

type Content struct {
	Model
	MediaMetadata
	MiniSeries []MiniSeries `json:"mini_series"`
}

func (Content) TableName() string {
	return "content"
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
