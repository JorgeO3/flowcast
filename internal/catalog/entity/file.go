package entity

// AudioFile represents the value object for an audio file.
type AudioFile struct {
	Name string `json:"name"`
	Ext  string `json:"ext"`
	Size int    `json:"size"`
}
