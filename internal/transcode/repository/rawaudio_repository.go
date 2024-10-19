package repository

import (
	"context"
)

// RawAudioRepository is a repository for the raw audio files in the transcoding service.
type RawAudioRepository interface {
	DownloadSong(ctx context.Context, songName string) (string, error)
}
