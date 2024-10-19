// Package repository provides the implementation of the catalog repository interface
package repository

import (
	"github.com/JorgeO3/flowcast/internal/catalog/repository/act"
	"github.com/JorgeO3/flowcast/internal/catalog/repository/assets"
	"github.com/JorgeO3/flowcast/internal/catalog/repository/rawaudio"
)

// Repositories represents the repositories used by the catalog.
type Repositories struct {
	Act      act.Repository
	RawAudio rawaudio.Repository
	Assets   assets.Repository
}

// Builder helps to build the repositories used by the catalog.
type Builder struct {
	repos Repositories
}

// NewRepositoryBuilder creates a new repository builder.
func NewRepositoryBuilder() *Builder {
	return &Builder{}
}

// WithActRepository sets the ActRepository in the builder.
func (rb *Builder) WithActRepository(repo act.Repository) *Builder {
	rb.repos.Act = repo
	return rb
}

// WithRawAudioRepository sets the RawAudioRepository in the builder.
func (rb *Builder) WithRawAudioRepository(repo rawaudio.Repository) *Builder {
	rb.repos.RawAudio = repo
	return rb
}

// WithAssetsRepository sets the AssetsRepository in the builder.
func (rb *Builder) WithAssetsRepository(repo assets.Repository) *Builder {
	rb.repos.Assets = repo
	return rb
}

// Build creates the repositories.
func (rb *Builder) Build() Repositories {
	return rb.repos
}
