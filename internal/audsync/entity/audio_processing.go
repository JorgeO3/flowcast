// Package entity provides the entities for the audsync service.
package entity

// AudioProcessing is a entity that represents the processing of an asset.
type AudioProcessing struct {
	EventID             string `json:"eventId"`
	AudioID             string `json:"audioId"`
	ActID               string `json:"actId"`
	AlbumID             string `json:"albumId"`
	SongID              string `json:"songId"`
	FilePath            string `json:"filePath"`
	Name                string `json:"name"`
	ActName             string `json:"actName"`
	AlbumName           string `json:"albumName"`
	CoverArtURL         string `json:"coverArtUrl"`
	Status              string `json:"status"`
	ProcessingStartTime string `json:"processingStartTime"`
	ProcessingEndTime   string `json:"processingEndTime"`
	ErrorMessage        string `json:"errorMessage"`
	CreatedAt           string `json:"createdAt"`
}
