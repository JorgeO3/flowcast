import { faker } from "npm:@faker-js/faker";

type MediaContentTypes = "Album" | "Act" | "Playlist" | "Podcast" | "Radio";

interface Cover {
    src: string;
    type: "Circle" | "Square";
}

interface Playlist {
    name: string;
    id: string;
}

interface MediaContent {
    type: MediaContentTypes;
    cover: Cover;
    playlists: Playlist[];
    title: string;
    text: string;
    id: string;
    link: string;
}

interface Section {
    title: string;
    text: string;
    link: string;
    type: "Album" | "Act" | "Playlist" | "Podcast" | "Radio" | "Mixed";
    id: string;
    boxes: MediaContent[];
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

function generateCover(): Cover {
    return {
        src: `https://picsum.photos/id/${faker.number.int(1000)}/200`,
        type: faker.helpers.arrayElement(["Circle", "Square"]),
    };
}

function generatePlaylist(): Playlist {
    return {
        name: faker.music.genre(),
        id: faker.string.uuid(),
    };
}

function generateMediaBox(): MediaContent {
    const type = faker.helpers.arrayElement([
        "Album",
        "Act",
        "Playlist",
        "Podcast",
        "Radio",
    ]);
    const id = faker.string.uuid();
    return {
        id,
        type: type as MediaContentTypes,
        cover: generateCover(),
        playlists: Array.from(
            { length: faker.number.int({ min: 1, max: 5 }) },
            generatePlaylist,
        ),
        title: faker.music.songName(),
        text: faker.helpers.arrayElement([
            "Feel the rhythm of timeless classics.",
            "An eclectic mix of melodies for every mood.",
            "Experience the latest hits that are shaping the world of music.",
            "Immerse yourself in the sounds that defined an era.",
            "Discover the artists pushing the boundaries of music.",
            "An emotional journey through heartfelt lyrics and powerful beats.",
        ]),
        link: generateLink(type as MediaContentTypes, id),
    };
}

function getDynamicRoute(type: MediaContentTypes): string {
    switch (type) {
        case "Radio":
        case "Playlist":
            return "Playlist";
        case "Podcast":
            return "show";
        case "Album":
            return "Album";
        case "Act":
            return "Act";
        default:
            return "play";
    }
}

function generateLink(type: MediaContentTypes | "sec", id: string): string {
    if (type === "sec") {
        return `/play/section/${id}`;
    }
    const routeType = getDynamicRoute(type);
    return `/play/${routeType}/${id}`;
}

function generateSection(): Section {
    const type = faker.helpers.arrayElement([
        "Album",
        "Act",
        "Playlist",
        "Podcast",
        "Radio",
        "Mixed",
    ]);
    const boxes = Array.from({ length: 10 }, generateMediaBox);
    const id = faker.string.uuid();
    return {
        id,
        title: faker.helpers.arrayElement([
            "The Ultimate Collection",
            "Timeless Sounds",
            "Journey Through Beats",
            "Hits of the Moment",
            "Classics Reimagined",
            "Sounds of the Future",
        ]),
        text: faker.helpers.arrayElement([
            "Explore this carefully curated collection of musical gems.",
            "A selection of the best tunes for every occasion.",
            "Dive into a world of melodies and rhythms.",
            "A must-listen for music enthusiasts of all kinds.",
            "Feel the power of music through every beat.",
            "Rediscover your favorite songs and artists in a new way.",
        ]),
        link: generateLink("sec", id), // Generar link basado en el primer MediaBox
        type: type as MediaContentTypes,
        boxes,
    };
}

function generateSections(count: number): Section[] {
    const sections: Section[] = [];

    for (let i = 0; i < count; i++) {
        sections.push(generateSection());
    }

    const rootDir = Deno.env.get("PWD");
    const sectionsFile = `${rootDir}/utils/tabAllSections.ts`;

    // and also write the data types to the file
    const data = `${TYPES}export const sections: Section[] = ${
        JSON.stringify(sections, null, 2)
    };\n`;

    Deno.writeTextFileSync(sectionsFile, data);
    // deno-lint-ignore no-deprecated-deno-api
    Deno.run({ cmd: ["deno", "fmt", sectionsFile] });
    return sections;
}

// Cambia el número de secciones generadas aquí
generateSections(5);
