import { faker } from "npm:@faker-js/faker";

// type MediaContentTypes = "Album" | "Act" | "Playlist" | "Podcast" | "Radio";

// interface Cover {
//     src: string;
//     type: "Circle" | "Square";
// }

// interface Playlist {
//     name: string;
//     id: string;
// }

// interface MediaContent {
//     type: MediaContentTypes;
//     cover: Cover;
//     playlists: Playlist[];
//     title: string;
//     text: string;
//     id: string;
//     link: string;
// }

// interface Section {
//     title: string;
//     text: string;
//     link: string;
//     type: "Album" | "Act" | "Playlist" | "Podcast" | "Radio" | "Mixed";
//     id: string;
//     boxes: MediaContent[];
// }

interface Section {
    title: string;
    text: string;
    link: string;
}

const TYPES = `\
export interface Cover {
  src: string;
  type: "Circle" | "Square";
}

export interface Playlist {
  name: string;
  id: string;
}

export interface MediaContent {
  type: 'Album' | 'Act' | 'Playlist' | 'Podcast' | 'Radio';
  cover: Cover;
  playlists: Playlist[];
  title: string;
  text: string;
  id: string;
  link: string;
}

export interface Section {
  title: string;
  text: string;
  link: string;
  type: 'Album' | 'Act' | 'Playlist' | 'Podcast' | 'Radio' | 'Mixed';
  id: string;
  boxes: MediaContent[];
}\n\n`;
