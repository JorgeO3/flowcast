package grpc

import (
	"context"
	"errors"
	"time"

	pb "github.com/JorgeO3/flowcast/gen/catalog"
	"github.com/JorgeO3/flowcast/internal/catalog/entity"
	e "github.com/JorgeO3/flowcast/internal/catalog/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Principal catalog errors
// Validation
// NotFound
// Conflict
// Internal
// Unauthorized
// Forbidden
// BadRequest
func mapCatalogErrorToGRPCError(err e.CatalogError) error {
	switch err.(type) {
	case e.Validation:
		return status.Error(codes.InvalidArgument, err.Msg())
	case e.NotFound:
		return status.Error(codes.NotFound, err.Msg())
	case e.Conflict:
		return status.Error(codes.AlreadyExists, err.Msg())
	case e.Internal:
		return status.Error(codes.Internal, err.Msg())
	case e.Unauthorized:
		return status.Error(codes.Unauthenticated, err.Msg())
	case e.Forbidden:
		return status.Error(codes.PermissionDenied, err.Msg())
	case e.BadRequest:
		return status.Error(codes.InvalidArgument, err.Msg())
	default:
		return status.Error(codes.Internal, "Internal Server Error")
	}
}

func (c *Controller) handleError(err error) error {
	var catalogErr e.CatalogError
	if errors.As(err, &catalogErr) {
		c.Logger.Error("Error executing CreateAct use case - err", err)
		return mapCatalogErrorToGRPCError(catalogErr)
	}

	c.Logger.Error("Unexpected error - err", err)
	return status.Error(codes.Internal, "Internal Server Error")
}

// withTimeout crea un contexto con un timeout de 5 segundos.
func withTimeout(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, 5*time.Second)
}

func convertPbActToEntity(pbAct *pb.Act) entity.Act {
	return entity.Act{
		ID:                entity.GetObjectID(pbAct.Id),
		Name:              pbAct.Name,
		Type:              pbAct.Type,
		Biography:         pbAct.Biography,
		FormationDate:     pbAct.FormationDate,
		DisbandDate:       pbAct.DisbandDate,
		ProfilePictureURL: pbAct.ProfilePictureUrl,
		Genres:            convertPbGenresToEntity(pbAct.Genres),
		Albums:            convertPbAlbumsToEntity(pbAct.Albums),
		Members:           convertPbMembersToEntity(pbAct.Members),
	}
}

func convertPbActsToEntity(pbActs []*pb.Act) []*entity.Act {
	var acts []*entity.Act
	for _, pbAct := range pbActs {
		act := convertPbActToEntity(pbAct)
		acts = append(acts, &act)
	}
	return acts
}

func convertPbGenreToEntity(pbGenre *pb.Genre) entity.Genre {
	return entity.Genre{
		Name:        pbGenre.Name,
		Description: pbGenre.Description,
	}
}

func convertPbGenresToEntity(pbGenres []*pb.Genre) []entity.Genre {
	var genres []entity.Genre
	for _, pbGenre := range pbGenres {
		genres = append(genres, convertPbGenreToEntity(pbGenre))
	}
	return genres
}

func convertPbAlbumToEntity(pbAlbum *pb.Album) entity.Album {
	return entity.Album{
		ID:          entity.GetObjectID(pbAlbum.Id),
		Title:       pbAlbum.Title,
		ReleaseDate: pbAlbum.ReleaseDate,
		Genre:       convertPbGenreToEntity(pbAlbum.Genre),
		CoverArtURL: pbAlbum.CoverArtUrl,
		TotalTracks: int(pbAlbum.TotalTracks),
		Songs:       convertPbSongsToEntity(pbAlbum.Songs),
	}
}

func convertPbAlbumsToEntity(pbAlbums []*pb.Album) []entity.Album {
	var albums []entity.Album
	for _, pbAlbum := range pbAlbums {
		albums = append(albums, convertPbAlbumToEntity(pbAlbum))
	}
	return albums
}

func convertPbSongToEntity(pbSong *pb.Song) entity.Song {
	return entity.Song{
		ID:            entity.GetObjectID(pbSong.Id),
		Title:         pbSong.Title,
		AudioFeatures: convertPbAudioFeaturesToEntity(pbSong.AudioFeatures),
		Genre:         convertPbGenreToEntity(pbSong.Genre),
		ReleaseDate:   pbSong.ReleaseDate,
		Duration:      int(pbSong.Duration),
		Lyrics:        pbSong.Lyrics,
		Explicit:      pbSong.Explicit,
		Bitrates:      convertPbAudioBitratesToEntity(pbSong.Bitrates),
	}
}

func convertPbSongsToEntity(pbAlbums []*pb.Song) []entity.Song {
	var songs []entity.Song
	for _, pbSong := range pbAlbums {
		songs = append(songs, convertPbSongToEntity(pbSong))
	}
	return songs
}

func convertPbAudioBitrateToEntity(pbAudioBitrate *pb.AudioBitrate) entity.AudioBitrate {
	return entity.AudioBitrate{
		Bitrate:  int(pbAudioBitrate.Bitrate),
		AudioURL: pbAudioBitrate.AudioUrl,
	}
}

func convertPbAudioBitratesToEntity(pbAudioBitrates []*pb.AudioBitrate) []entity.AudioBitrate {
	var bitrates []entity.AudioBitrate
	for _, pbAudioBitrate := range pbAudioBitrates {
		bitrates = append(bitrates, convertPbAudioBitrateToEntity(pbAudioBitrate))
	}
	return bitrates
}

func convertPbAudioFeaturesToEntity(pbAudioFeatures *pb.AudioFeatures) entity.AudioFeatures {
	return entity.AudioFeatures{
		Tempo:            int(pbAudioFeatures.Tempo),
		AudioKey:         pbAudioFeatures.AudioKey,
		Mode:             pbAudioFeatures.Mode,
		Loudness:         pbAudioFeatures.Loudness,
		Energy:           pbAudioFeatures.Energy,
		Danceability:     pbAudioFeatures.Danceability,
		Speechiness:      pbAudioFeatures.Speechiness,
		Acousticness:     pbAudioFeatures.Acousticness,
		Instrumentalness: pbAudioFeatures.Instrumentalness,
		Liveness:         pbAudioFeatures.Liveness,
		Velence:          pbAudioFeatures.Valence,
	}
}

func convertPbMemberToEntity(pbMember *pb.Member) entity.Member {
	return entity.Member{
		Name:              pbMember.Name,
		Biography:         pbMember.Biography,
		BirthDate:         pbMember.BirthDate,
		StartDate:         pbMember.StartDate,
		EndDate:           pbMember.EndDate,
		ProfilePictureURL: pbMember.ProfilePictureUrl,
	}
}

func convertPbMembersToEntity(pbMembers []*pb.Member) []entity.Member {
	var members []entity.Member
	for _, pbMember := range pbMembers {
		members = append(members, convertPbMemberToEntity(pbMember))
	}
	return nil
}

func convertEntityActToPb(act *entity.Act) *pb.Act {
	return &pb.Act{
		Id:                act.ID.String(),
		Name:              act.Name,
		Type:              act.Type,
		Biography:         act.Biography,
		FormationDate:     act.FormationDate,
		DisbandDate:       act.DisbandDate,
		ProfilePictureUrl: act.ProfilePictureURL,
		Genres:            convertEntityGenresToPb(act.Genres),
		Albums:            convertEntityAlbumsToPb(act.Albums),
		Members:           convertEntityMembersToPb(act.Members),
	}
}

func convertEntityActsToPb(acts []*entity.Act) []*pb.Act {
	var pbActs []*pb.Act
	for _, act := range acts {
		pbActs = append(pbActs, convertEntityActToPb(act))
	}
	return pbActs
}

func convertEntityGenreToPb(genre entity.Genre) *pb.Genre {
	return &pb.Genre{
		Name:        genre.Name,
		Description: genre.Description,
	}
}

func convertEntityGenresToPb(genres []entity.Genre) []*pb.Genre {
	var pbGenres []*pb.Genre
	for _, genre := range genres {
		pbGenres = append(pbGenres, convertEntityGenreToPb(genre))
	}
	return nil
}

func convertEntityAudioFeaturesToPb(audioFeatures entity.AudioFeatures) *pb.AudioFeatures {
	return &pb.AudioFeatures{
		Tempo:            int32(audioFeatures.Tempo),
		AudioKey:         audioFeatures.AudioKey,
		Mode:             audioFeatures.Mode,
		Loudness:         audioFeatures.Loudness,
		Energy:           audioFeatures.Energy,
		Danceability:     audioFeatures.Danceability,
		Speechiness:      audioFeatures.Speechiness,
		Acousticness:     audioFeatures.Acousticness,
		Instrumentalness: audioFeatures.Instrumentalness,
		Liveness:         audioFeatures.Liveness,
		Valence:          audioFeatures.Velence,
	}
}

func convertEntityAudioBitrateToPb(audioBitrate entity.AudioBitrate) *pb.AudioBitrate {
	return &pb.AudioBitrate{
		Bitrate:  int32(audioBitrate.Bitrate),
		AudioUrl: audioBitrate.AudioURL,
	}
}

func convertEntityAudioBitratesToPb(audioBitrates []entity.AudioBitrate) []*pb.AudioBitrate {
	var pbAudioBitrates []*pb.AudioBitrate
	for _, audioBitrate := range audioBitrates {
		pbAudioBitrates = append(pbAudioBitrates, convertEntityAudioBitrateToPb(audioBitrate))
	}
	return pbAudioBitrates
}

func convertEntitySongToPb(song entity.Song) *pb.Song {
	return &pb.Song{
		Id:            song.ID.Hex(),
		Title:         song.Title,
		AudioFeatures: convertEntityAudioFeaturesToPb(song.AudioFeatures),
		Genre:         convertEntityGenreToPb(song.Genre),
		ReleaseDate:   song.ReleaseDate,
		Duration:      int32(song.Duration),
		Lyrics:        song.Lyrics,
		Explicit:      song.Explicit,
		Bitrates:      convertEntityAudioBitratesToPb(song.Bitrates),
	}
}

func convertEntitySongsToPb(songs []entity.Song) []*pb.Song {
	var pbSongs []*pb.Song
	for _, song := range songs {
		pbSongs = append(pbSongs, convertEntitySongToPb(song))
	}
	return pbSongs
}

func convertEntityAlbumToPb(album entity.Album) *pb.Album {
	return &pb.Album{
		Id:          album.ID.Hex(),
		Title:       album.Title,
		ReleaseDate: album.ReleaseDate,
		Genre:       convertEntityGenreToPb(album.Genre),
		CoverArtUrl: album.CoverArtURL,
		TotalTracks: int32(album.TotalTracks),
		Songs:       convertEntitySongsToPb(album.Songs),
	}
}

func convertEntityAlbumsToPb(albums []entity.Album) []*pb.Album {
	var pbAlbums []*pb.Album
	for _, album := range albums {
		pbAlbums = append(pbAlbums, convertEntityAlbumToPb(album))
	}
	return pbAlbums
}

func convertEntityMemberToPb(member entity.Member) *pb.Member {
	return &pb.Member{
		Name:              member.Name,
		Biography:         member.Biography,
		BirthDate:         member.BirthDate,
		StartDate:         member.StartDate,
		EndDate:           member.EndDate,
		ProfilePictureUrl: member.ProfilePictureURL,
	}
}

func convertEntityMembersToPb(members []entity.Member) []*pb.Member {
	var pbMembers []*pb.Member
	for _, member := range members {
		pbMembers = append(pbMembers, convertEntityMemberToPb(member))
	}
	return pbMembers
}
