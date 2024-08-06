import librosa
import argparse

from mutagen.oggopus import OggOpus

def extract_metadata(file_path):
    audio = OggOpus(file_path)
    metadata = {k: v for k, v in audio.items()}
    return metadata

def edit_metadata(file_path, metadata):
    audio = OggOpus(file_path)
    
    if "title" in metadata:
        audio["TITLE"] = metadata["title"]
    if "artist" in metadata:
        audio["ARTIST"] = metadata["artist"]
    if "album" in metadata:
        audio["ALBUM"] = metadata["album"]
    if "genre" in metadata:
        audio["GENRE"] = metadata["genre"]
    if "date" in metadata:
        audio["DATE"] = metadata["date"]
    if "tracknumber" in metadata:
        audio["TRACKNUMBER"] = metadata["tracknumber"]
    
    audio.save()

def extract_audio_features(file_path):
    try:
        y, sr = librosa.load(file_path)
        
        tempo, _ = librosa.beat.beat_track(y=y, sr=sr)
        key = librosa.feature.chroma_stft(y=y, sr=sr).argmax()
        mode = librosa.feature.chroma_cqt(y=y, sr=sr).argmax()
        danceability = librosa.feature.rms(y=y).mean()
        energy = librosa.feature.mfcc(y=y, sr=sr).mean()
        speechiness = librosa.feature.spectral_contrast(y=y, sr=sr).mean()
        acousticness = librosa.feature.zero_crossing_rate(y=y).mean()
        instrumentalness = librosa.feature.spectral_bandwidth(y=y, sr=sr).mean()
        liveness = librosa.feature.spectral_flatness(y=y).mean()
        valence = librosa.feature.tonnetz(y=y, sr=sr).mean()

        return {
            'tempo': tempo,
            'key': key,
            'mode': mode,
            'danceability': danceability,
            'energy': energy,
            'speechiness': speechiness,
            'acousticness': acousticness,
            'instrumentalness': instrumentalness,
            'liveness': liveness,
            'valence': valence
        }
    except Exception as e:
        print(f"Error loading audio file with librosa: {e}")
        return {}

def main():
    parser = argparse.ArgumentParser(description='Extract and edit metadata for Opus files and get audio features using librosa.')
    parser.add_argument('file', type=str, help='Path to the Opus file')
    parser.add_argument('--edit', action='store_true', help='Edit metadata')
    parser.add_argument('--title', type=str, help='Title of the song')
    parser.add_argument('--artist', type=str, help='Artist of the song')
    parser.add_argument('--album', type=str, help='Album of the song')
    parser.add_argument('--genre', type=str, help='Genre of the song')
    parser.add_argument('--date', type=str, help='Release date of the song')
    parser.add_argument('--tracknumber', type=str, help='Track number of the song')
    parser.add_argument('--all', action='store_true', help='Show all metadata')
    parser.add_argument('--features', action='store_true', help='Extract audio features using librosa')

    args = parser.parse_args()

    if args.edit:
        metadata_to_edit = {
            "title": args.title,
            "artist": args.artist,
            "album": args.album,
            "genre": args.genre,
            "date": args.date,
            "tracknumber": args.tracknumber
        }
        metadata_to_edit = {k: v for k, v in metadata_to_edit.items() if v is not None}
        edit_metadata(args.file, metadata_to_edit)
        print(f"Metadata updated for {args.file}")
    else:
        metadata = extract_metadata(args.file)
        if args.all:
            print(f"All metadata for {args.file}:")
            for key, value in metadata.items():
                print(f"{key}: {value}")
        else:
            selected_metadata = {
                "title": metadata.get("TITLE"),
                "artist": metadata.get("ARTIST"),
                "album": metadata.get("ALBUM"),
                "genre": metadata.get("GENRE"),
                "date": metadata.get("DATE"),
                "tracknumber": metadata.get("TRACKNUMBER"),
            }
            print(f"Selected metadata for {args.file}:")
            for key, value in selected_metadata.items():
                print(f"{key}: {value}")

    if args.features:
        audio_features = extract_audio_features(args.file)
        print(f"Audio features for {args.file}:")
        for key, value in audio_features.items():
            print(f"{key}: {value}")

if __name__ == "__main__":
    main()
