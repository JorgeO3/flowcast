type ContentType =
  | "album"
  | "act"
  | "playlist"
  | "podcast"
  | "radio"
  | "mixed"
  | "song";

export interface Cover {
  src: string;
}

interface BaseMediaContent {
  cover: Cover;
  title: string;
  text: string;
  id: string;
  link: string;
  contentType: ContentType;
}

// Album content
export interface AlbumContent extends BaseMediaContent {
  contentType: "album";
}

// Act content (e.g. artist)
export interface ActContent extends BaseMediaContent {
  contentType: "act";
}

// Playlist content
export interface PlaylistContent extends BaseMediaContent {
  contentType: "playlist";
}

// Podcast content
export interface PodcastContent extends BaseMediaContent {
  contentType: "podcast";
}

// Radio content
export interface RadioContent extends BaseMediaContent {
  contentType: "radio";
}

// Radio song
export interface SongContent extends BaseMediaContent {
  contentType: "song";
}

export type MediaContent =
  | ActContent
  | SongContent
  | AlbumContent
  | RadioContent
  | PodcastContent
  | PlaylistContent;

export interface Section {
  id: string;
  text: string;
  link: string;
  title: string;
  content: MediaContent[];
}
