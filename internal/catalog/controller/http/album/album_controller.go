package http

import (
	"github.com/JorgeO3/flowcast/configs"
	"github.com/JorgeO3/flowcast/pkg/logger"
)

// AlbumController handles HTTP requests for the catalog service and delegates operations to use cases.
type AlbumController struct {
	Logger logger.Interface
	Cfg    *configs.CatalogConfig
}
