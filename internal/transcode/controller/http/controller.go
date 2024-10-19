// Package http provides the HTTP Controllers for the transcoding service.
package http

import (
	"net/http"

	"github.com/JorgeO3/flowcast/configs"
	"github.com/JorgeO3/flowcast/internal/transcode/usecase"
	"github.com/JorgeO3/flowcast/pkg/logger"
	"github.com/JorgeO3/flowcast/pkg/minio"
)

// Controller handles HTTP requests for the transcoding service and delegates operations to use cases.
type Controller struct {
	TranscodeUC *usecase.TranscodeSongUC
	Logger      logger.Interface
	Cfg         *configs.TranscodeConfig
}

// TranscodeSong handles transcoding a song.
func (c *Controller) TranscodeSong(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := c.withTimeout(r)
	defer cancel()

	var event minio.Event
	if !c.decodeJSON(w, r, &event) {
		return
	}

	input := usecase.TranscodeSongInput{Filename: event.Records[0].S3.Object.Key}
	output, err := c.TranscodeUC.Execute(ctx, input)
	if err != nil {
		c.Logger.Error("Error executing TranscodeSong use case - err %v", err)
		c.handleError(w, err)
		return
	}

	c.respondJSON(w, output)
}
