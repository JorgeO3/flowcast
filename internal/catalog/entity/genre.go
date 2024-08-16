package entity

// Genre represent a value object
type Genre struct {
	Name        string `json:"name,omitempty" bson:"name,omitempty" valid:"required,alpha"`
	Description string `json:"description,omitempty" bson:"description,omitempty" valid:"required,alpha"`
}

// GenreOption represent the functional options for the genre entity
type GenreOption func(*Genre)

// WithGenreName set the name of the genre
func WithGenreName(name string) GenreOption {
	return func(g *Genre) {
		g.Name = name
	}
}

// WithGenreDescription set the description of the genre
func WithGenreDescription(description string) GenreOption {
	return func(g *Genre) {
		g.Description = description
	}
}

// NewGenre create a new genre entity
func NewGenre(opts ...GenreOption) *Genre {
	genre := &Genre{}
	for _, opt := range opts {
		opt(genre)
	}
	return genre
}
