package repository

import "gitlab.com/JorgeO3/flowcast/internal/catalog/entity"

type SongRepository interface {
	GetByID(actID, albumID, songID string) (*entity.Song, error)
	UpdateSong(actID, albumID, songID string, song entity.Song) error
}
