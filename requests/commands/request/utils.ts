import { Act } from "/commands/generate/generate-acts-types.ts";

// Asynchronous function to read data from a JSON file
export async function readData(path: string): Promise<Act[]> {
    try {
        const jsonData = await Deno.readTextFile(path);
        return JSON.parse(jsonData);
    } catch (error) {
        console.error("Error reading data from file:", error);
        Deno.exit(1);
    }
}

// Function to perform HTTP requests
export async function request(url: string, method: string, body?: string): Promise<any> {
    const headers = new Headers({
        "Content-Type": "application/json",
    });

    const options: RequestInit = {
        method,
        headers,
    };

    if (body) {
        options.body = body;
    }

    try {
        const response = await fetch(url, options);
        if (!response.ok) {
            const errorText = await response.text();
            throw new Error(`HTTP Error ${response.status}: ${errorText}`);
        }
        return await response.json();
    } catch (error) {
        console.error("Request error:", error);
        throw error;
    }
}

export function parseRange(range: string, items: number): [number, number] {
    if (range === "all") return [0, items - 1];
    const [start, end] = range.split("-").map(Number);

    if (isNaN(start) || isNaN(end)) {
        throw new Error("Invalid range format. Please use the format 'start-end'");
    }

    if (start < 0 || end < 0) {
        throw new Error("Invalid range values. Please use valid indices");
    }

    if (start >= items || end >= items) {
        throw new Error("Range values exceed the number of items: " + items);
    }

    return [start, end];
}
