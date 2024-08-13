package repository

import "gitlab.com/JorgeO3/flowcast/internal/catalog/entity"

// AlbumRepository provides an interface to interact with the album repository.
type AlbumRepository interface {
	GetByID(actID, albumID string) (*entity.Album, error)
	UpdateAlbum(actID string, album entity.Album) error
	AddSong(actID, albumID string, song entity.Song) error
	RemoveSong(actID, albumID, songID string) error
}
