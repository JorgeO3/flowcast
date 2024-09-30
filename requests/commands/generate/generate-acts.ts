import { faker } from "faker";

// Interfaces are assumed to be imported or defined elsewhere
import type {
    Act,
    Album,
    AudioBitrate,
    AudioFeatures,
    Genre,
    Member,
    Song,
} from "./generate-acts-types.ts";

// List of possible musical keys and modes
const KEYS = ["C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"] as const;
const MODES = ["major", "minor"] as const;

// Generate a musical key
function generateKey(): string {
    return `${faker.helpers.arrayElement(KEYS)} ${faker.helpers.arrayElement(MODES)}`;
}

// Generate a music genre
function generateGenre(): Genre {
    return {
        name: faker.music.genre(),
        description: faker.lorem.sentence(),
    };
}

// Generate audio features for a song
function generateAudioFeatures(): AudioFeatures {
    return {
        tempo: faker.number.int({ min: 60, max: 200 }),
        audioKey: generateKey(),
        mode: faker.number.int({ min: 0, max: 1 }).toString(),
        loudness: faker.number.float({ min: -60, max: 0 }),
        energy: faker.number.float({ min: 0, max: 1 }),
        danceability: faker.number.float({ min: 0, max: 1 }),
        speechiness: faker.number.float({ min: 0, max: 1 }),
        acousticness: faker.number.float({ min: 0, max: 1 }),
        instrumentalness: faker.number.float({ min: 0, max: 1 }),
        liveness: faker.number.float({ min: 0, max: 1 }),
        valence: faker.number.float({ min: 0, max: 1 }),
    };
}

// Generate audio bitrate information
function generateAudioBitrate(): AudioBitrate {
    return {
        bitrate: faker.number.int({ min: 64, max: 320 }),
        audioUrl: faker.internet.url(),
    };
}

function formatDateString(date: Date): string {
    return date.toISOString().split("T")[0];
}

function generatePastDate(yearsAgo: number): string {
    const date = faker.date.past({ years: yearsAgo });
    return formatDateString(date);
}

function generateFutureDate(yearsAhead: number): string {
    const date = faker.date.future({ years: yearsAhead });
    return formatDateString(date);
}

function generateRecentDate(): string {
    const date = faker.date.recent();
    return formatDateString(date);
}

// Generate a song
function generateSong(): Song {
    return {
        id: faker.database.mongodbObjectId(),
        title: faker.lorem.words(),
        audioFeatures: generateAudioFeatures(),
        genre: generateGenre(),
        releaseDate: generatePastDate(10),
        duration: faker.number.int({ min: 120, max: 600 }),
        lyrics: faker.lorem.paragraph(),
        explicit: faker.datatype.boolean(),
        bitrates: Array.from({ length: 3 }, generateAudioBitrate),
    };
}

// Generate an album
function generateAlbum(): Album {
    return {
        id: faker.database.mongodbObjectId(),
        title: faker.lorem.words(),
        releaseDate: generatePastDate(20),
        genre: generateGenre(),
        coverArtUrl: faker.internet.url(),
        totalTracks: faker.number.int({ min: 1, max: 20 }),
        songs: Array.from(
            { length: faker.number.int({ min: 5, max: 15 }) },
            generateSong,
        ),
    };
}

// Generate a band member
function generateMember(): Member {
    return {
        name: faker.person.fullName(),
        biography: faker.lorem.paragraph(),
        birthDate: generatePastDate(50),
        profilePictureUrl: faker.image.avatar(),
        startDate: generatePastDate(20),
        endDate: faker.datatype.boolean() ? generateRecentDate() : null,
    };
}

// Generate an act (band or artist)
function generateAct(): Act {
    return {
        id: faker.database.mongodbObjectId(),
        name: faker.company.name(),
        type: faker.helpers.arrayElement(["Band", "Solo Artist", "Duo"]),
        biography: faker.lorem.paragraph(),
        formationDate: generatePastDate(30),
        disbandDate: faker.datatype.boolean() ? generateRecentDate() : null,
        profilePictureUrl: faker.image.avatar(),
        genres: Array.from(
            { length: faker.number.int({ min: 1, max: 3 }) },
            generateGenre,
        ),
        albums: Array.from(
            { length: faker.number.int({ min: 1, max: 10 }) },
            generateAlbum,
        ),
        members: Array.from(
            { length: faker.number.int({ min: 1, max: 5 }) },
            generateMember,
        ),
    };
}

// Generate a specified number of acts
export function generateActs(numberOfActs: number = 1000) {
    return Array.from({ length: numberOfActs }, generateAct);
}

// Save generated data to a file asynchronously
export async function saveActsToFile(filePath: string, data: unknown[]): Promise<void> {
    try {
        const jsonData = JSON.stringify(data, null, 2);
        const dir = filePath.split("/").slice(0, -1).join("/");

        await Deno.mkdir(dir, { recursive: true });
        await Deno.writeTextFile(filePath, jsonData);

        console.log(`Data successfully saved to ${filePath}`);
    } catch (error) {
        console.error("Error saving data to file:", error);
        throw error;
    }
}
