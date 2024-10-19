// Package usecase provides the use cases for the transcoding service.
package usecase

import (
	"context"
	"sync"

	"github.com/JorgeO3/flowcast/internal/transcode/errors"
	"github.com/JorgeO3/flowcast/internal/transcode/repository"
	"github.com/JorgeO3/flowcast/internal/transcode/utils"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/validator"
)

const (
	chunkSize  = 10
	outputDir  = "/tmp/hls_output"
	ffprobeBin = "ffprobe"
	ffmpegBin  = "ffmpeg"
)

var bitrates = [3]string{"64000", "128000", "192000"}

type (
	// TranscodeSongInput represents the input for the TranscodeSong use case.
	TranscodeSongInput struct {
		Filename string
	}

	// TranscodeSongOutput represents the output for the TranscodeSong use case.
	TranscodeSongOutput struct{}

	// TranscodeSongUC is the use case for transcoding a song.
	TranscodeSongUC struct {
		EoRepo    repository.EncodedOpusRepository
		RaRepo    repository.RawAudioRepository
		Logger    logger.Interface
		Validator validator.Interface
	}
)

// TranscodeSongUCOpts represents the functional options for the TranscodeSongUC.
type TranscodeSongUCOpts func(uc *TranscodeSongUC)

// WithTranscodeSongRepository sets the RawAudioRepository in the TranscodeSongUC.
func WithTranscodeSongRepository(repo repository.RawAudioRepository) TranscodeSongUCOpts {
	return func(uc *TranscodeSongUC) {
		uc.RaRepo = repo
	}
}

// WithTranscodeSongLogger sets the
func WithTranscodeSongLogger(logger logger.Interface) TranscodeSongUCOpts {
	return func(uc *TranscodeSongUC) {
		uc.Logger = logger
	}
}

// WithTranscodeSongValidator sets the validator in the TranscodeSongUC.
func WithTranscodeSongValidator(validator validator.Interface) TranscodeSongUCOpts {
	return func(uc *TranscodeSongUC) {
		uc.Validator = validator
	}
}

// WithTranscodeSongEncodedOpusRepository sets the EncodedOpus
func WithTranscodeSongEncodedOpusRepository(repo repository.EncodedOpusRepository) TranscodeSongUCOpts {
	return func(uc *TranscodeSongUC) {
		uc.EoRepo = repo
	}
}

// NewTranscodeSongUC creates a new TranscodeSongUC.
func NewTranscodeSongUC(opts ...TranscodeSongUCOpts) *TranscodeSongUC {
	uc := &TranscodeSongUC{}
	for _, opt := range opts {
		opt(uc)
	}
	return uc
}

func (uc *TranscodeSongUC) generateChunks(filePath string, chunks []utils.Chunk) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(bitrates)*len(chunks))

	for _, bitrate := range bitrates {
		for _, chunk := range chunks {
			wg.Add(1)
			go func(b string, c utils.Chunk) {
				defer wg.Done()
				if err := utils.GenerateChunk(filePath, c, b, ffmpegBin, outputDir); err != nil {
					errCh <- err
				}
			}(bitrate, chunk)
		}
	}

	wg.Wait()
	close(errCh)
	return <-errCh
}

// Execute performs the TranscodeSong use case.
func (uc *TranscodeSongUC) Execute(ctx context.Context, input TranscodeSongInput) (*TranscodeSongOutput, error) {
	uc.Logger.Info("Transcoding song %s", input.Filename)

	if err := uc.Validator.Validate(input); err != nil {
		uc.Logger.Error("invalid input - err %v", err)
		return nil, errors.NewValidation("invalid input", err)
	}

	filePath, err := uc.RaRepo.DownloadSong(ctx, input.Filename)
	if err != nil {
		uc.Logger.Error("failed to download song - err %v", err)
		return nil, errors.NewInternal("failed to download song", err)
	}

	chunks, err := utils.SplitFile(filePath, chunkSize)
	if err != nil {
		uc.Logger.Error("failed to split file - err %v", err)
		return nil, errors.NewInternal("failed to split file", err)
	}

	if err := uc.generateChunks(filePath, chunks); err != nil {
		uc.Logger.Error("failed to generate chunkss - err %v", err)
		return nil, errors.NewInternal("failed to generate chunks", err)
	}

	if err := utils.CreateMasterManifest(bitrates[:], len(chunks), outputDir); err != nil {
		uc.Logger.Error("failed to create master manifest - err %v", err)
		return nil, errors.NewInternal("failed to create master manifest", err)
	}

	songName := utils.GetSongName(filePath)
	if err := uc.EoRepo.PutSong(ctx, songName, outputDir); err != nil {
		uc.Logger.Error("failed to put song - err %v", err)
		return nil, errors.NewInternal("failed to put song", err)
	}

	uc.Logger.Info("Song transcoded and uploaded successfully")

	return &TranscodeSongOutput{}, nil
}
