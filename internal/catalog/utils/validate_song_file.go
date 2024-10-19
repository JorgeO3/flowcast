package utils

import (
	"github.com/JorgeO3/flowcast/internal/catalog/entity"
	"github.com/JorgeO3/flowcast/internal/catalog/errors"
)

// ValidateSongFile validates the song file.
func ValidateSongFile(file *entity.AudioFile) error {
	if file.Ext != "wav" {
		return errors.NewValidation("invalid file extension", nil)
	}

	maxFileSize := 1024 * 1024 * 50 // 10MB
	if file.Size > maxFileSize {
		return errors.NewValidation("invalid file size", nil)
	}
	return nil
}
