package output

import (
	"movie/entities"
	"time"
)

type Movie struct {
	ID          string    `json:"_id"`
	Title       string    `json:"title"`
	Image       string    `json:"image"`
	Language    string    `json:"language"`
	Genre       string    `json:"genre"`
	Director    string    `json:"director"`
	Cast        string    `json:"cast"`
	Description string    `json:"description"`
	Duration    string    `json:"duration"`
	ReleaseDate time.Time `json:"releaseDate"`
	EndDate     time.Time `json:"endDate"`
}

func ParseToOutput(ent *entities.Movie) (out *Movie) {
	return &Movie{
		ID:          ent.ID,
		Title:       ent.Title,
		Image:       ent.Image,
		Language:    ent.Language,
		Genre:       ent.Genre,
		Director:    ent.Director,
		Cast:        ent.Cast,
		Description: ent.Description,
		Duration:    ent.Duration,
		ReleaseDate: ent.ReleaseDate,
		EndDate:     ent.EndDate,
	}
}
