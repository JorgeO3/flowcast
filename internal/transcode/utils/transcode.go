// Package utils provides utility functions for the transcoding service.
package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

// Chunk represents a chunk of a song.
type Chunk struct {
	Number   int
	Start    float64
	Duration float64
}

// GetSongName returns the name of a song file without the extension.
func GetSongName(filePath string) string {
	return strings.TrimSuffix(filepath.Base(filePath), filepath.Ext(filePath))
}

// CmdExec Execute a command
func CmdExec(args ...string) (string, error) {
	baseCmd := args[0]
	cmdArgs := args[1:]

	cmd := exec.Command(baseCmd, cmdArgs...)
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(out), nil
}

// GetDuration returns the duration of a media file.
func GetDuration(filePath string) (float64, error) {
	cmd, err := CmdExec("ffprobe",
		"-i", filePath,
		"-show_entries", "format=duration",
		"-v", "quiet",
		"-of", "csv=p=0",
	)
	if err != nil {
		return 0, err
	}
	durationStr := strings.TrimSpace(cmd)
	return strconv.ParseFloat(durationStr, 64)
}

// SplitFile splits a media file into chunks of a given size.
func SplitFile(filePath string, chunkSize int) ([]Chunk, error) {
	duration, err := GetDuration(filePath)
	if err != nil {
		return nil, err
	}

	numChunks := int(duration) / chunkSize
	if int(duration)%chunkSize != 0 {
		numChunks++
	}

	chunks := make([]Chunk, numChunks)
	for i := 0; i < numChunks; i++ {
		start := float64(i * chunkSize)
		dur := float64(chunkSize)
		if start+dur > duration {
			dur = duration - start
		}
		chunks[i] = Chunk{
			Number:   i + 1,
			Start:    start,
			Duration: dur,
		}
	}
	return chunks, nil
}

// GenerateChunk generates a chunk of a media file.
func GenerateChunk(filePath string, chunk Chunk, bitrate, ffmpegBin, outputDir string) error {
	bitrateFolder := filepath.Join(outputDir, bitrate)
	segmentFilename := "chunk_%03d.m4s"
	playlistFilename := fmt.Sprintf("playlist_%03d.m3u8", chunk.Number)
	playlistPath := filepath.Join(bitrateFolder, playlistFilename)

	if err := os.MkdirAll(bitrateFolder, 0755); err != nil {
		return err
	}

	_, err := CmdExec(ffmpegBin,
		"-hide_banner",
		"-loglevel", "error",
		"-vn",
		"-i", filePath,
		"-c:a", "aac",
		"-b:a", bitrate,
		"-ac", "2",
		"-ar", "48000",
		"-profile:a", "aac_low",
		"-ss", fmt.Sprintf("%.0f", chunk.Start),
		"-t", fmt.Sprintf("%.0f", chunk.Duration),
		"-f", "hls",
		"-hls_time", "10",
		"-hls_playlist_type", "vod",
		"-hls_segment_type", "fmp4",
		"-hls_segment_filename", filepath.Join(bitrateFolder, segmentFilename),
		playlistPath,
	)

	return err
}

// CreateMasterManifest creates a master manifest file for a set of media chunks.s
func CreateMasterManifest(bitrates []string, numChunks int, outputDir string) error {
	masterManifest := filepath.Join(outputDir, "master.m3u8")
	file, err := os.Create(masterManifest)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString("#EXTM3U\n"); err != nil {
		return err
	}

	for _, bitrate := range bitrates {
		for i := 1; i <= numChunks; i++ {
			playlistPath := fmt.Sprintf("%s/playlist_%03d.m3u8", bitrate, i)
			if _, err := fmt.Fprintf(file, "#EXT-X-STREAM-INF:BANDWIDTH=%s\n%s\n", bitrate, playlistPath); err != nil {
				return err
			}
		}
	}

	return nil
}
