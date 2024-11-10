package utils

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/JorgeO3/flowcast/internal/catalog/entity"
	"github.com/JorgeO3/flowcast/internal/catalog/repository"
)

// MockRepository es un mock del repositorio que implementa el método GeneratePresignedURL
type MockRepository struct {
	GeneratePresignedURLFunc func(ctx context.Context, fileName string, expiration time.Duration) (string, error)
}

func (m *MockRepository) GeneratePresignedURL(ctx context.Context, fileName string, expiration time.Duration) (string, error) {
	if m.GeneratePresignedURLFunc != nil {
		return m.GeneratePresignedURLFunc(ctx, fileName, expiration)
	}
	return "", errors.New("not implemented")
}

// TestAssetsProcessor_AddAlbum prueba la adición de un nuevo álbum a un Act existente
func TestAssetsProcessor_AddAlbum(t *testing.T) {
	ctx := context.Background()

	// Mock de los repositorios
	rawAudioRepo := &MockRepository{
		GeneratePresignedURLFunc: func(_ context.Context, fileName string, _ time.Duration) (string, error) {
			return "https://presigned-url.com" + fileName, nil
		},
	}
	assetsRepo := &MockRepository{
		GeneratePresignedURLFunc: func(_ context.Context, fileName string, _ time.Duration) (string, error) {
			return "https://presigned-url.com" + fileName, nil
		},
	}

	repos := &repository.Repositories{
		RawAudio: rawAudioRepo,
		Assets:   assetsRepo,
	}

	// Act antiguo
	oldAct := &entity.Act{
		ID:   "act1",
		Name: "Test Act",
		ProfilePicture: entity.Asset{
			ID:   "asset1",
			Type: entity.ImageAct,
			Name: "profile_pic_old",
			Ext:  "jpg",
			Size: 1024,
			URL:  "/assets/act1.jpg",
		},
		Albums: []entity.Album{
			{
				ID:    "album1",
				Title: "Existing Album",
				CoverArt: entity.Asset{
					ID:   "asset2",
					Type: entity.ImageAlbum,
					Name: "cover_art_old",
					Ext:  "jpg",
					Size: 2048,
					URL:  "/assets/act1/album1.jpg",
				},
				Songs: []entity.Song{
					{
						ID:    "song1",
						Title: "Existing Song",
						File: entity.Asset{
							ID:   "asset3",
							Type: entity.Audio,
							Name: "song_file_old",
							Ext:  "wav",
							Size: 4096,
							URL:  "/raw-audio/act1/album1/song1.wav",
						},
						CoverArt: entity.Asset{
							ID:   "asset4",
							Type: entity.ImageAlbum,
							Name: "song_cover_old",
							Ext:  "jpg",
							Size: 1024,
							URL:  "/assets/act1/album1/song1.jpg",
						},
					},
				},
			},
		},
	}

	// Act nuevo con un álbum adicional
	newAct := &entity.Act{
		ID:   "act1",
		Name: "Test Act",
		ProfilePicture: entity.Asset{
			ID:   "asset1",
			Type: entity.ImageAct,
			Name: "profile_pic_old",
			Ext:  "jpg",
			Size: 1024,
			URL:  "", // Se generará
		},
		Albums: []entity.Album{
			oldAct.Albums[0], // Álbum existente
			{
				ID:    "album2",
				Title: "New Album",
				CoverArt: entity.Asset{
					ID:   "asset5",
					Type: entity.ImageAlbum,
					Name: "cover_art_new_album",
					Ext:  "jpg",
					Size: 2048,
					URL:  "", // Se generará
				},
				Songs: []entity.Song{
					{
						ID:    "song2",
						Title: "New Song",
						File: entity.Asset{
							ID:   "asset6",
							Type: entity.Audio,
							Name: "song_file_new",
							Ext:  "wav",
							Size: 4096,
							URL:  "", // Se generará
						},
						CoverArt: entity.Asset{
							ID:   "asset7",
							Type: entity.ImageAlbum,
							Name: "song_cover_new",
							Ext:  "jpg",
							Size: 1024,
							URL:  "", // Se generará
						},
					},
				},
			},
		},
	}

	// Generamos las URLs para el nuevo act
	GenerateURLs(newAct)

	// Instancia del AssetsProcessor
	processor := NewAssetsProcessor(ctx, repos)

	// Llamada al método Update
	output, err := processor.Update(oldAct, newAct)
	if err != nil {
		t.Fatalf("Error en Update: %v", err)
	}

	// Validaciones
	if len(output.AssetsURLs) != 3 {
		t.Errorf("Se esperaban 3 URLs generadas, se obtuvieron %d", len(output.AssetsURLs))
	}

	if len(output.AddedAssets) != 3 {
		t.Errorf("Se esperaban 3 assets agregados, se obtuvieron %d", len(output.AddedAssets))
	}

	if len(output.DeletedAssets) != 0 {
		t.Errorf("No se esperaban assets eliminados, se obtuvieron %d", len(output.DeletedAssets))
	}
}

// TestAssetsProcessor_DeleteSong prueba la eliminación de una canción de un álbum
func TestAssetsProcessor_DeleteSong(t *testing.T) {
	ctx := context.Background()

	// Mock de los repositorios (no se usarán en Delete)
	rawAudioRepo := &MockRepository{}
	assetsRepo := &MockRepository{}

	repos := &repository.Repositories{
		RawAudio: rawAudioRepo,
		Assets:   assetsRepo,
	}

	// Act antiguo con dos canciones
	oldAct := &entity.Act{
		ID:   "act1",
		Name: "Test Act",
		ProfilePicture: entity.Asset{
			ID:   "asset1",
			Type: entity.ImageAct,
			Name: "profile_pic",
			Ext:  "jpg",
			Size: 1024,
			URL:  generateActPictureURL("act1"),
		},
		Albums: []entity.Album{
			{
				ID:    "album1",
				Title: "Test Album",
				CoverArt: entity.Asset{
					ID:   "asset2",
					Type: entity.ImageAlbum,
					Name: "cover_art",
					Ext:  "jpg",
					Size: 2048,
					URL:  generateAlbumCoverArtURL("act1", "album1"),
				},
				Songs: []entity.Song{
					{
						ID:    "song1",
						Title: "Song to Keep",
						File: entity.Asset{
							ID:   "asset3",
							Type: entity.Audio,
							Name: "song_file_keep",
							Ext:  "wav",
							Size: 4096,
							URL:  generateSongFileURL("act1", "album1", "song1"),
						},
					},
					{
						ID:    "song2",
						Title: "Song to Delete",
						File: entity.Asset{
							ID:   "asset4",
							Type: entity.Audio,
							Name: "song_file_delete",
							Ext:  "wav",
							Size: 4096,
							URL:  generateSongFileURL("act1", "album1", "song2"),
						},
					},
				},
			},
		},
	}

	// Act nuevo sin la segunda canción
	newAct := &entity.Act{
		ID:   "act1",
		Name: "Test Act",
		ProfilePicture: entity.Asset{
			ID:   "asset1",
			Type: entity.ImageAct,
			Name: "profile_pic",
			Ext:  "jpg",
			Size: 1024,
			URL:  "", // Se generará
		},
		Albums: []entity.Album{
			{
				ID:    "album1",
				Title: "Test Album",
				CoverArt: entity.Asset{
					ID:   "asset2",
					Type: entity.ImageAlbum,
					Name: "cover_art",
					Ext:  "jpg",
					Size: 2048,
					URL:  "", // Se generará
				},
				Songs: []entity.Song{
					oldAct.Albums[0].Songs[0], // Solo mantenemos la primera canción
				},
			},
		},
	}

	// Generamos las URLs para el nuevo act
	GenerateURLs(newAct)

	// Instancia del AssetsProcessor
	processor := NewAssetsProcessor(ctx, repos)
	output, err := processor.Update(oldAct, newAct)
	if err != nil {
		t.Fatalf("Error en Update: %v", err)
	}

	// Validaciones
	if len(output.AssetsURLs) != 0 {
		t.Errorf("No se esperaban URLs generadas, se obtuvieron %d", len(output.AssetsURLs))
	}

	if len(output.AddedAssets) != 0 {
		t.Errorf("No se esperaban assets agregados, se obtuvieron %d", len(output.AddedAssets))
	}

	if len(output.DeletedAssets) != 1 {
		t.Errorf("Se esperaba 1 asset eliminado, se obtuvieron %d", len(output.DeletedAssets))
	}
}

// TestAssetsProcessor_UpdateAlbumCover prueba la actualización del CoverArt de un álbum
func TestAssetsProcessor_UpdateAlbumCover(t *testing.T) {
	ctx := context.Background()

	// Mock de los repositorios
	rawAudioRepo := &MockRepository{
		GeneratePresignedURLFunc: func(_ context.Context, fileName string, _ time.Duration) (string, error) {
			return "https://presigned-url.com" + fileName, nil
		},
	}
	assetsRepo := &MockRepository{
		GeneratePresignedURLFunc: func(_ context.Context, fileName string, _ time.Duration) (string, error) {
			return "https://presigned-url.com" + fileName, nil
		},
	}

	repos := &repository.Repositories{
		RawAudio: rawAudioRepo,
		Assets:   assetsRepo,
	}

	// Act antiguo
	oldAct := &entity.Act{
		ID:   "act1",
		Name: "Test Act",
		ProfilePicture: entity.Asset{
			ID:   "asset1",
			Type: entity.ImageAct,
			Name: "profile_pic",
			Ext:  "jpg",
			Size: 1024,
			URL:  generateActPictureURL("act1"),
		},
		Albums: []entity.Album{
			{
				ID:    "album1",
				Title: "Test Album",
				CoverArt: entity.Asset{
					ID:   "asset2",
					Type: entity.ImageAlbum,
					Name: "cover_art_old",
					Ext:  "jpg",
					Size: 2048,
					URL:  generateAlbumCoverArtURL("act1", "album1"),
				},
				Songs: []entity.Song{},
			},
		},
	}

	// Act nuevo con CoverArt actualizado
	newAct := &entity.Act{
		ID:   "act1",
		Name: "Test Act",
		ProfilePicture: entity.Asset{
			ID:   "asset1",
			Type: entity.ImageAct,
			Name: "profile_pic",
			Ext:  "jpg",
			Size: 1024,
			URL:  "", // Se generará
		},
		Albums: []entity.Album{
			{
				ID:    "album1",
				Title: "Test Album",
				CoverArt: entity.Asset{
					ID:   "asset2",
					Type: entity.ImageAlbum,
					Name: "cover_art_new",
					Ext:  "jpg",
					Size: 2048,
					URL:  "", // Se generará
				},
				Songs: []entity.Song{},
			},
		},
	}

	// Generamos las URLs para el nuevo act
	GenerateURLs(newAct)

	// Llamada al método Update
	processor := NewAssetsProcessor(ctx, repos)
	output, err := processor.Update(oldAct, newAct)
	if err != nil {
		t.Fatalf("Error en Update: %v", err)
	}

	// Validaciones
	if len(output.AssetsURLs) != 1 {
		t.Errorf("Se esperaba 1 URL generada, se obtuvieron %d", len(output.AssetsURLs))
	}

	if len(output.AddedAssets) != 1 {
		t.Errorf("Se esperaba 1 asset agregado, se obtuvieron %d", len(output.AddedAssets))
	}

	if len(output.DeletedAssets) != 0 {
		t.Errorf("No se esperaban assets eliminados, se obtuvieron %d", len(output.DeletedAssets))
	}
}

// TestAssetsProcessor_EmptyAssets prueba el procesamiento cuando los assets están vacíos
func TestAssetsProcessor_EmptyAssets(t *testing.T) {
	ctx := context.Background()

	// Mock de los repositorios
	rawAudioRepo := &MockRepository{}
	assetsRepo := &MockRepository{}

	repos := &repository.Repositories{
		RawAudio: rawAudioRepo,
		Assets:   assetsRepo,
	}

	// Act sin assets
	act := &entity.Act{
		ID:   "act1",
		Name: "Test Act",
		// ProfilePicture está vacío
		Albums: []entity.Album{
			{
				ID:    "album1",
				Title: "Test Album",
				// CoverArt está vacío
				Songs: []entity.Song{
					{
						ID:    "song1",
						Title: "Test Song",
						// File y CoverArt están vacíos
					},
				},
			},
		},
	}

	// Generamos las URLs para el act
	GenerateURLs(act)

	// Llamada al método Create
	processor := NewAssetsProcessor(ctx, repos)
	output, err := processor.Create(act)
	if err != nil {
		t.Fatalf("Error en Create: %v", err)
	}

	// Validaciones
	if len(output.AssetsURLs) != 0 {
		t.Errorf("No se esperaban URLs generadas, se obtuvieron %d", len(output.AssetsURLs))
	}

	if len(output.AddedAssets) != 0 {
		t.Errorf("No se esperaban assets agregados, se obtuvieron %d", len(output.AddedAssets))
	}

	if len(output.DeletedAssets) != 0 {
		t.Errorf("No se esperaban assets eliminados, se obtuvieron %d", len(output.DeletedAssets))
	}
}

// TestAssetsProcessor_MultipleAlbumsAndSongs prueba un Act con múltiples álbumes y canciones
func TestAssetsProcessor_MultipleAlbumsAndSongs(t *testing.T) {
	ctx := context.Background()

	// Mock de los repositorios
	rawAudioRepo := &MockRepository{
		GeneratePresignedURLFunc: func(_ context.Context, fileName string, _ time.Duration) (string, error) {
			return "https://presigned-url.com" + fileName, nil
		},
	}
	assetsRepo := &MockRepository{
		GeneratePresignedURLFunc: func(_ context.Context, fileName string, _ time.Duration) (string, error) {
			return "https://presigned-url.com" + fileName, nil
		},
	}

	repos := &repository.Repositories{
		RawAudio: rawAudioRepo,
		Assets:   assetsRepo,
	}

	// Act de prueba con múltiples álbumes y canciones
	act := &entity.Act{
		ID:   "act1",
		Name: "Test Act",
		ProfilePicture: entity.Asset{
			ID:   "asset1",
			Type: entity.ImageAct,
			Name: "profile_pic",
			Ext:  "jpg",
			Size: 1024,
			URL:  "", // Se generará
		},
		Albums: []entity.Album{
			{
				ID:    "album1",
				Title: "Album One",
				CoverArt: entity.Asset{
					ID:   "asset2",
					Type: entity.ImageAlbum,
					Name: "cover_art_album1",
					Ext:  "jpg",
					Size: 2048,
					URL:  "", // Se generará
				},
				Songs: []entity.Song{
					{
						ID:    "song1",
						Title: "Song One",
						File: entity.Asset{
							ID:   "asset3",
							Type: entity.Audio,
							Name: "song_file1",
							Ext:  "wav",
							Size: 4096,
							URL:  "", // Se generará
						},
					},
					{
						ID:    "song2",
						Title: "Song Two",
						File: entity.Asset{
							ID:   "asset4",
							Type: entity.Audio,
							Name: "song_file2",
							Ext:  "wav",
							Size: 4096,
							URL:  "", // Se generará
						},
					},
				},
			},
			{
				ID:    "album2",
				Title: "Album Two",
				CoverArt: entity.Asset{
					ID:   "asset5",
					Type: entity.ImageAlbum,
					Name: "cover_art_album2",
					Ext:  "jpg",
					Size: 2048,
					URL:  "", // Se generará
				},
				Songs: []entity.Song{
					{
						ID:    "song3",
						Title: "Song Three",
						File: entity.Asset{
							ID:   "asset6",
							Type: entity.Audio,
							Name: "song_file3",
							Ext:  "wav",
							Size: 4096,
							URL:  "", // Se generará
						},
					},
				},
			},
		},
	}

	// Generamos las URLs para el act
	GenerateURLs(act)

	// Llamada al método Create
	processor := NewAssetsProcessor(ctx, repos)
	output, err := processor.Create(act)
	if err != nil {
		t.Fatalf("Error en Create: %v", err)
	}

	// Calculamos el número esperado de URLs
	expectedURLsCount := 1 // ProfilePicture
	expectedURLsCount += 2 // CoverArt de álbumes
	expectedURLsCount += 3 // Archivos de canciones

	if len(output.AssetsURLs) != expectedURLsCount {
		t.Errorf("Se esperaban %d URLs generadas, se obtuvieron %d", expectedURLsCount, len(output.AssetsURLs))
	}

	if len(output.AddedAssets) != expectedURLsCount {
		t.Errorf("Se esperaban %d assets agregados, se obtuvieron %d", expectedURLsCount, len(output.AddedAssets))
	}

	if len(output.DeletedAssets) != 0 {
		t.Errorf("No se esperaban assets eliminados, se obtuvieron %d", len(output.DeletedAssets))
	}
}
