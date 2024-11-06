// Package repository provides all the repositories for the audsync service.
package repository

import (
	"github.com/JorgeO3/flowcast/internal/audsync/repository/assets"
	"github.com/JorgeO3/flowcast/internal/audsync/repository/audprocess"
	"github.com/JorgeO3/flowcast/internal/audsync/repository/transcodaudio"
)

// Repositories holds all the repositories.
type Repositories struct {
	Process       audprocess.Repository
	Assets        assets.Repository
	Transcodaudio transcodaudio.Repository
}

// Opts is a type for setting the repositories.
type Opts func(r *Repositories)

// WithProcess sets the process repository.
func WithProcess(r audprocess.Repository) Opts {
	return func(repos *Repositories) {
		repos.Process = r
	}
}

// WithAssets sets the assets repository.
func WithAssets(r assets.Repository) Opts {
	return func(repos *Repositories) {
		repos.Assets = r
	}
}

// WithTranscodaudio sets the transcodaudio repository.
func WithTranscodaudio(r transcodaudio.Repository) Opts {
	return func(repos *Repositories) {
		repos.Transcodaudio = r
	}
}

// NewRepositories creates a new instance of Repositories.
func NewRepositories(opts ...Opts) *Repositories {
	r := &Repositories{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}
