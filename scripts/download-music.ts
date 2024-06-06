// @ts-nocheck
import { join } from "https://deno.land/std@0.224.0/path/mod.ts";
import { Text } from "npm:domhandler";
import * as cheerio from "npm:cheerio@latest";

/**
 * Pauses execution for a specified number of seconds.
 * @param seconds The number of seconds to pause.
 * @returns A promise that resolves after the specified time.
 */
function sleep(seconds: number): Promise<void> {
    return new Promise((resolve) => setTimeout(resolve, seconds * 1000));
}

/**
 * Retrieves an environment variable.
 * @param key The key of the environment variable.
 * @returns The value of the environment variable.
 * @throws An error if the environment variable is not found.
 */
function getEnvVariable(key: string): string {
    const value = Deno.env.get(key);
    if (!value) throw new Error(`Environment variable ${key} not found`);
    return value;
}

/**
 * Cleans the artist name to create a valid URL segment.
 * @param name The name to clean.
 * @returns The cleaned name.
 */
function cleanArtistNames(name: string): string {
    return name
        .replaceAll(" ", "-")
        .replaceAll("&", "")
        .replaceAll("--", "-")
        .replaceAll(".", "-")
        .replaceAll(":", "")
        .toLowerCase();
}

/**
 * Cleans the song name to create a valid URL segment.
 * @param name The name to clean.
 * @returns The cleaned name.
 */
function cleanSongName(name: string) {
    return name.replaceAll("–", "-");
}

/**
 * Extracts artist names from the given URL.
 * @param url The URL of the artists page.
 * @returns A promise that resolves to an array of artist names.
 */
async function extractArtists(url: string): Promise<string[]> {
    const artistsEndPoint = `${url}/artistas`;
    const response = await fetch(artistsEndPoint);
    const rawHtml = await response.text();

    const $ = cheerio.load(rawHtml);
    const $titles = $('span[id="titulo"]');

    const artists: string[] = [];
    $titles.each((_, element) => {
        const artist = $(element).text();
        artists.push(artist);
    });
    return artists;
}

/**
 * Retrieves the HTML of the artist's songs page.
 * @param url The base URL.
 * @param path The path to the artist's songs page.
 * @returns A promise that resolves to the HTML of the artist's songs page.
 */
async function fetchHtml(url: string, path: string): Promise<string> {
    try {
        const response = await fetch(`${url}/${path}/`);
        if (!response.ok) {
            throw new Error(`Failed to fetch HTML: ${response.statusText}`);
        }
        return await response.text();
    } catch (error) {
        console.error(`Error fetching HTML for ${path}:`, error);
        throw error; // Re-throw the error to be handled by the caller
    }
}

/**
 * Extracts song names from the given HTML.
 * @param html The HTML of the songs page.
 * @returns An array of song names.
 */
function extractSongNames(html: string): string[] {
    const $ = cheerio.load(html);
    const $links = $("ul li a");

    const songNames: string[] = [];
    $links.each((_, element) => {
        if (!element.attribs.href.includes("https://mp3teca.co/mp3/")) return;
        const songName = (element.children[0] as Text).data;
        if (songName) songNames.push(songName);
    });

    return songNames;
}

/**
 * Downloads a song from the given URL and saves it as an MP3 file.
 * @param songUrl The URL of the song.
 * @param fileName The name to save the file as.
 * @returns A promise that resolves when the song has been downloaded and saved.
 */
async function downloadSong(
    songUrl: string,
    targetPath: string,
): Promise<void> {
    const response = await fetch(songUrl);
    if (!response.ok) {
        throw new Error(`Failed to download ${songUrl}: ${response.statusText}`);
    }

    const contentType = response.headers.get("Content-Type");
    if (contentType !== "application/octet-stream") {
        throw new Error(`Unexpected content type ${contentType} for ${songUrl}`);
    }

    const destFile = await Deno.open(targetPath, {
        create: true,
        write: true,
        truncate: true,
    });

    await response.body?.pipeTo(destFile.writable);
    console.log(`Downloaded and saved ${targetPath}`);
}

/**
 * Downloads songs for the given artists.
 * @param webUrl The base web URL.
 * @param serverUrl The server URL.
 * @param targetDir The directory to save the downloaded songs.
 * @param artists The array of artist names.
 * @returns A promise that resolves when all songs have been downloaded.
 */
async function downloadSongs(
    webUrl: string,
    serverUrl: string,
    targetDir: string,
    artists: string[],
): Promise<void> {
    for (const artist of artists) {
        const cleanedArtist = cleanArtistNames(artist);

        let html: string;
        try {
            html = await fetchHtml(webUrl, cleanedArtist);
        } catch (error) {
            console.error(`Skipping artist ${cleanedArtist} due to fetch error.`);
            continue; // Skip to the next artist
        }

        const songNames = extractSongNames(html);

        for (const songName of songNames) {
            const cleanedSongName = cleanSongName(songName);
            const songUrl = `${serverUrl}/${cleanedSongName}.mp3`;
            const targetSongPath = join(targetDir, `${cleanedSongName}.mp3`);
            console.log(songUrl);

            try {
                await downloadSong(songUrl, targetSongPath);
            } catch (error) {
                console.error(`Error downloading ${songUrl}:`, error);
            }

            await sleep(2);
        }
    }
}

if (import.meta.main) {
    try {
        const songsDir = getEnvVariable("SONGS_DIR");
        const webUrl = getEnvVariable("WEB_ENDPOINT");
        const serverUrl = getEnvVariable("SERVER_ENDPOINT");

        const artists = await extractArtists(webUrl);
        await downloadSongs(webUrl, serverUrl, songsDir, artists);
    } catch (error) {
        console.error("Error:", error);
    }
}
