package events

// EventType represents the type of event.
type EventType int

const (
	// CreateAudioProcessings is the event that is triggered when a new audio processing is created.
	CreateAudioProcessings EventType = iota
	// DeleteAudioProcessings is the event that is triggered when an audio processing is deleted.
	DeleteAudioProcessings
	// UpdateAudioProcessings is the event that is triggered when an audio processing is updated.
	UpdateAudioProcessings
)

// BaseAudioEvent represents the base event structure for audio events.
type BaseAudioEvent struct {
	Type    EventType   `json:"type"`
	Payload interface{} `json:"payload"`
}
