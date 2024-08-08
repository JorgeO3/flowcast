// Package app provides the entry point to the catalog service.
package app

import (
	"github.com/asaskevich/govalidator"
	"gitlab.com/JorgeO3/flowcast/configs"
	"gitlab.com/JorgeO3/flowcast/pkg/logger"
)

// Run starts the catalog service.
func Run(cfg *configs.CatalogConfig, logg logger.Interface) {
	govalidator.SetFieldsRequiredByDefault(true)
}
