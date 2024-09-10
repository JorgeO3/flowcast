import { Args } from "jsr:@std/cli/parse-args";
import {
    Action,
    type Command,
    type CommandFlags,
    type Flag,
} from "../utils/index.ts";

const REQUIRED_FLAGS: string[] = ["-u", "-a"];

const ALL_FLAGS: Flag[] = [
    {
        flag: "-u, --url <url:string>",
        description: "Specify the URL for the request (required)",
        flagType: "string",
    },
    {
        flag: "-p, --path <path:string>",
        description:
            "Specify the path to read generated data (default: ./data/acts.json)",
        flagType: "string",
    },
    {
        flag: "-a, --action <action:string>",
        description: "Specify the action for request command (required)",
        flagType: "string",
    },
];

const OPTIONS: CommandFlags = {
    all: ALL_FLAGS,
    required: ALL_FLAGS.filter(({ flag }) => REQUIRED_FLAGS.includes(flag)),
};

const ACTIONS: Action[] = [
    { name: "getActs", description: "Get all acts" },
    { name: "createAct", description: "Create a new act" },
    { name: "getActById", description: "Get an act by ID" },
    { name: "updateAct", description: "Update an act by ID" },
    { name: "deleteAct", description: "Delete an act by ID" },
];

const EXAMPLES: string[] = [
    "deno run -A requests/main.ts generate --type act --count 10 --path ./data/acts.json",
];

function handler(_args: Args) {}

export const GENERATE_COMMAND: Command = {
    name: "request, r <action:string>",
    options: OPTIONS,
    actions: ACTIONS,
    examples: EXAMPLES,
    run: handler,
};
