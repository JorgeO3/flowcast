package proxy

import (
	"gitlab.com/JorgeO3/flowcast/configs"
	"gitlab.com/JorgeO3/flowcast/pkg/logger"
)

func Run(cfg *configs.Config) {
	l := logger.New(cfg.LogLevel)
}
