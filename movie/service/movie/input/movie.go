package input

import (
	"movie/entities"
	"time"
)

type MovieInput struct {
	ID          string    `json:"-"`
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

func (input *MovieInput) ParseToEntities() (ent *entities.Movie) {
	return &entities.Movie{
		ID:          input.ID,
		Title:       input.Title,
		Image:       input.Image,
		Language:    input.Language,
		Genre:       input.Genre,
		Director:    input.Director,
		Cast:        input.Cast,
		Description: input.Description,
		Duration:    input.Duration,
		ReleaseDate: input.ReleaseDate,
		EndDate:     input.EndDate,
	}
}
