export interface Genre {
    name: string;
    description: string;
}

export interface AudioFeatures {
    tempo: number;
    audioKey: string;
    mode: string;
    loudness: number;
    energy: number;
    danceability: number;
    speechiness: number;
    acousticness: number;
    instrumentalness: number;
    liveness: number;
    valence: number;
}

export interface AudioBitrate {
    bitrate: number;
    audioUrl: string;
}

export interface Song {
    id: string;
    title: string;
    audioFeatures: AudioFeatures;
    genre: Genre;
    releaseDate: string;
    duration: number;
    lyrics: string;
    explicit: boolean;
    bitrates: AudioBitrate[];
}

export interface Album {
    id: string;
    title: string;
    releaseDate: string;
    genre: Genre;
    coverArtUrl: string;
    totalTracks: number;
    songs: Song[];
}

export interface Member {
    name: string;
    biography: string;
    birthDate: string;
    profilePictureUrl: string;
    startDate: string;
    endDate: string | null;
}

export interface Act {
    id: string;
    name: string;
    type: string;
    biography: string;
    formationDate: string;
    disbandDate: string | null;
    profilePictureUrl: string;
    genres: Genre[];
    albums: Album[];
    members: Member[];
}
