package events

import "encoding/json"

type EventType int

const (
	// CreateAudioProcessings is the event that is triggered when a new audio processing is created.
	CreateAudioProcessings EventType = iota
	// DeleteAudioProcessings is the event that is triggered when an audio processing is deleted.
	DeleteAudioProcessings
	// UpdateAudioProcessings is the event that is triggered when an audio processing is updated.
	UpdateAudioProcessings
)

type BaseAudioEvent struct {
	Type    EventType       `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

func DeserializePayload[T any](event BaseAudioEvent) (T, error) {
	var payload T
	err := json.Unmarshal(event.Payload, &payload)
	return payload, err
}
