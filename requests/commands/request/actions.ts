import { parseRange, readData, request } from "./utils.ts";

export enum Action {
    GetActs = "getActs",
    CreateAct = "createAct",
    GetActById = "getActById",
    UpdateAct = "updateAct",
    DeleteAct = "deleteAct",
    CreateManyActs = "createManyActs",
}

// Definition of the arguments required for action handlers
export type HandlerArgs = {
    url: string;
    path: string;
    index: number;
    action: Action;
    range: string;
};

// Type for action handlers
export type ActionHandler = (args: HandlerArgs) => Promise<void>;

// Handler for the "getActs" action
export async function getActs(args: HandlerArgs): Promise<void> {
    const { url } = args;
    try {
        const acts = await request(url, "GET");
        console.log("Retrieved acts:", acts);
    } catch {
        Deno.exit(1);
    }
}

// Handler for the "createAct" action
export async function createAct(args: HandlerArgs): Promise<void> {
    const { url, path, index } = args;
    try {
        const acts = await readData(path);
        const act = acts[index];
        const body = JSON.stringify(act);
        const result = await request(url, "POST", body);
        console.log("Created act:", result);
    } catch {
        Deno.exit(1);
    }
}

// Handler for the "getActById" action
export async function getActById(args: HandlerArgs): Promise<void> {
    const { url, index, path } = args;
    try {
        const act = await readData(path);
        const actId = act[index].id;
        const actResult = await request(`${url}/${actId}`, "GET");
        console.log("Retrieved act by ID:", actResult);
    } catch {
        Deno.exit(1);
    }
}

// Handler for the "updateAct" action
export async function updateAct(args: HandlerArgs): Promise<void> {
    const { url, path, index } = args;
    try {
        const acts = await readData(path);
        const act = acts[index];
        const actId = act.id;
        const body = JSON.stringify(act);
        const result = await request(`${url}/${actId}`, "PUT", body);
        console.log("Updated act:", result);
    } catch {
        Deno.exit(1);
    }
}

// Handler for the "deleteAct" action
export async function deleteAct(args: HandlerArgs): Promise<void> {
    const { url, index, path } = args;
    try {
        const act = await readData(path);
        const actId = act[index].id;
        const result = await request(`${url}/${actId}`, "DELETE");
        console.log("Deleted act:", result);
    } catch {
        Deno.exit(1);
    }
}

// Handler for the "createManyActs" action
export async function createManyActs(args: HandlerArgs): Promise<void> {
    const { url, path, range } = args;
    try {
        const allActs = await readData(path);
        console.log({ url, path, range });

        const [start, end] = parseRange(range, allActs.length);
        const acts = allActs.slice(start, end);
        const body = JSON.stringify({ acts });
        const result = await request(url, "POST", body);
        console.log("Created acts:", result);
    } catch (error) {
        console.error("Error creating acts:", error);
        Deno.exit(1);
    }
}
