export interface Genre {
    name: string;
    description: string;
}

export interface AudioFeatures {
    tempo: number;
    audioKey: string;
    mode: number;
    loudness: number;
    energy: number;
    danceability: number;
    speechiness: number;
    acousticness: number;
    instrumentalness: number;
    liveness: number;
    valance: number;
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
    releaseDate: Date;
    duration: number;
    lyrics: string;
    explicit: boolean;
    bitrates: AudioBitrate[];
}

export interface Album {
    id: string;
    title: string;
    releaseDate: Date;
    genre: Genre;
    coverArtUrl: string;
    totalTracks: number;
    songs: Song[];
}

export interface Member {
    name: string;
    biography: string;
    birthDate: Date;
    profilePictureUrl: string;
    startDate: Date;
    endDate: Date | null;
}

export interface Act {
    id: string;
    name: string;
    type: string;
    biography: string;
    formationDate: Date;
    disbandDate: Date | null;
    profilePictureUrl: string;
    genres: Genre[];
    albums: Album[];
    members: Member[];
}
