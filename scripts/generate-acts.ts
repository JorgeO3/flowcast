// @ts-nocheck
import { faker } from "npm:@faker-js/faker";

// Lista de posibles claves musicales
const keys = ["C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"];
const modes = ["major", "minor"];

function generateKey() {
    return `${faker.helpers.arrayElement(keys)} ${
        faker.helpers.arrayElement(modes)
    }`;
}

function generateGenre() {
    return {
        name: faker.music.genre(),
        description: faker.lorem.sentence(),
    };
}

function generateAudioFeatures() {
    return {
        tempo: faker.number.int({ min: 60, max: 200 }),
        audioKey: generateKey(),
        mode: faker.number.int({ min: 0, max: 1 }),
        loudness: faker.number.float({ min: -60, max: 0 }),
        energy: faker.number.float({ min: 0, max: 1 }),
        danceability: faker.number.float({ min: 0, max: 1 }),
        speechiness: faker.number.float({ min: 0, max: 1 }),
        acousticness: faker.number.float({ min: 0, max: 1 }),
        instrumentalness: faker.number.float({ min: 0, max: 1 }),
        liveness: faker.number.float({ min: 0, max: 1 }),
        valance: faker.number.float({ min: 0, max: 1 }),
    };
}

function generateAudioBitrate() {
    return {
        bitrate: faker.number.int({ min: 64, max: 320 }),
        audioUrl: faker.internet.url(),
    };
}

function generateSong() {
    return {
        id: crypto.randomUUID(),
        title: faker.lorem.words(),
        audioFeatures: generateAudioFeatures(),
        genre: generateGenre(),
        releaseDate: faker.date.past({ years: 10 }),
        duration: faker.number.int({ min: 120, max: 600 }),
        lyrics: faker.lorem.paragraph(),
        explicit: faker.datatype.boolean(),
        bitrates: Array.from({ length: 3 }, generateAudioBitrate),
    };
}

function generateAlbum() {
    return {
        id: crypto.randomUUID(),
        title: faker.lorem.words(),
        releaseDate: faker.date.past({ years: 20 }),
        genre: generateGenre(),
        coverArtUrl: faker.internet.url(),
        totalTracks: faker.number.int({ min: 1, max: 20 }),
        songs: Array.from(
            { length: faker.number.int({ min: 5, max: 15 }) },
            generateSong,
        ),
    };
}

function generateMember() {
    return {
        name: faker.person.fullName(),
        biography: faker.lorem.paragraph(),
        birthDate: faker.date.past({ years: 50 }),
        profilePictureUrl: faker.image.avatar(),
        startDate: faker.date.past({ years: 20 }),
        endDate: faker.datatype.boolean() ? faker.date.recent() : null,
    };
}

function generateAct() {
    return {
        id: crypto.randomUUID(),
        name: faker.company.name(),
        type: faker.helpers.arrayElement(["Band", "Solo Artist", "Duo"]),
        biography: faker.lorem.paragraph(),
        formationDate: faker.date.past({ years: 30 }),
        disbandDate: faker.datatype.boolean() ? faker.date.recent() : null,
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

function generateActs(numberOfActs: number = 1000) {
    return Array.from({ length: numberOfActs }, generateAct);
}

function saveActsToFile(filePath: string, data: unknown[]): void {
    const jsonData = JSON.stringify(data, null, 2);
    Deno.writeTextFileSync(filePath, jsonData);
    console.log(`Data successfully saved to ${filePath}`);
}

const musicalActs = generateActs();

// Define el path donde se guardará el archivo JSON
const filePath = "./musicalActs.json";

// Guarda los datos generados en un archivo JSON
saveActsToFile(filePath, musicalActs);
