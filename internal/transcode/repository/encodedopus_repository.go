// Package repository provides the different repositories for the transcoding service
package repository

import (
	"context"
)

// EncodedOpusRepository is a repository for the encoded Opus files in the transcoding service.
type EncodedOpusRepository interface {
	PutSong(ctx context.Context, songName, outputDir string) error
}
