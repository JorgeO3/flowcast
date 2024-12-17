package entity

// AssetType represents the type of an asset
type AssetType string

const (
	// Audio represents an audio asset
	Audio AssetType = "audio"
	// ImageAct represents an image asset
	ImageAct AssetType = "imageAct"
	// ImageSong represents an image asset
	ImageSong AssetType = "imageSong"
	// ImageAlbum represents an image asset
	ImageAlbum AssetType = "imageAlbum"
)

// AssetExt represents the extension of an asset
type AssetExt string

const (
	// PNG represents a PNG image
	PNG AssetExt = "png"
	// JPEG represents a JPEG image
	JPEG AssetExt = "jpeg"
	// MP3 represents an MP3 audio
	MP3 AssetExt = "mp3"
	// WAV represents a WAV audio
	WAV AssetExt = "wav"
)

// Asset represents any file (image or audio) associated with an entity
type Asset struct {
	ID   string    `json:"id,omitempty" validate:"required" bson:"_id,omitempty"`
	Type AssetType `json:"type,omitempty" validate:"required,onof=image audio" bson:"type,omitempty"`
	Name string    `json:"name,omitempty" validate:"required,alphanum,min=1,max=255" bson:"name,omitempty"`
	Ext  AssetExt  `json:"ext,omitempty" validate:"required,oneof=png jpeg mp3 wav" bson:"ext,omitempty"`
	Size int       `json:"size,omitempty" validate:"required,assetsize" bson:"size,omitempty"`
	URL  string    `json:"url,omitempty" validate:"required,url" bson:"url,omitempty"`
}

// IsEmpty checks if the asset is empty
func (a Asset) IsEmpty() bool {
	return a == Asset{}
}
