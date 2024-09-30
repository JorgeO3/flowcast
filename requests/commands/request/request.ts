import type { Args, Command, CommandFlags, Flag, Options } from "/utils/index.ts";
import { createAct, createManyActs, deleteAct, getActById, getActs, updateAct } from "./actions.ts";
import { Action, ActionHandler } from "./actions.ts";

// Definition of available actions using an enum for type safety

// Required flags
const REQUIRED_FLAGS: string[] = ["url", "action"];

// Definition of all flags
const ALL_FLAGS: Flag[] = [
    {
        flag: "-u, --url <url:string>",
        description: "Specify the URL for the request (required)",
        short: "u",
        long: "url",
        default: "",
        flagType: "string",
    },
    {
        flag: "-p, --path <path:string>",
        description: "Specify the path to read generated data (default: ./data/acts.json)",
        short: "p",
        long: "path",
        default: "./data/acts.json",
        flagType: "string",
    },
    {
        flag: "-a, --action <action:string>",
        description: "Specify the action for request command (required)",
        short: "a",
        long: "action",
        default: "",
        flagType: "string",
    },
    {
        flag: "-i, --index <index:number>",
        description: "Specify the index of the data for request command (default: 0)",
        short: "i",
        long: "index",
        default: 0,
        flagType: "number",
    },
    {
        flag: "r, --range <range:string>",
        description: "Specify the range of data to read (default: all)",
        short: "r",
        long: "range",
        default: "all",
        flagType: "string",
    },
];

// Configuration of command flags
const FLAGS: CommandFlags = {
    all: ALL_FLAGS,
    required: ALL_FLAGS.filter(({ long }) => REQUIRED_FLAGS.includes(long)),
};

// Available options for the command
const OPTIONS: Options = {
    name: "ACTIONS",
    values: [
        { name: Action.GetActs, description: "Get all acts" },
        { name: Action.CreateAct, description: "Create a new act" },
        { name: Action.GetActById, description: "Get an act by ID" },
        { name: Action.UpdateAct, description: "Update an act by ID" },
        { name: Action.DeleteAct, description: "Delete an act by ID" },
        { name: Action.CreateManyActs, description: "Create multiple acts" },
    ],
};

// Examples of how to use the command
const EXAMPLES: string[] = [
    "deno run -A requests/main.ts request --url http://localhost:3000/acts --action createAct --path ./data/acts.json --index 1",
    "deno run -A requests/main.ts request -u http://localhost:3000/acts -a createManyActs -r 1-5",
];

// Map of actions to their respective handlers
const actionsMap: Record<Action, ActionHandler> = {
    [Action.GetActs]: getActs,
    [Action.CreateAct]: createAct,
    [Action.GetActById]: getActById,
    [Action.UpdateAct]: updateAct,
    [Action.DeleteAct]: deleteAct,
    [Action.CreateManyActs]: createManyActs,
};

// Main handler function that manages the command logic
async function handler(args: Args): Promise<void> {
    const url = args.url as string;
    const path = args.path as string;
    const range = args.range as string || "all";
    const index = args.index as number;
    const actionName = args.action as string;

    // Validate that the action is valid
    if (!Object.values(Action).includes(actionName as Action)) {
        console.error("Invalid action:", actionName);
        Deno.exit(1);
    }

    const action = actionName as Action;
    const actionHandler = actionsMap[action];

    try {
        await actionHandler({ url, path, index, action, range });
    } catch (error) {
        console.error("Error executing action:", error);
        Deno.exit(1);
    }
}

// Definition of the REQUEST_COMMAND
export const REQUEST_COMMAND: Command = {
    name: "request, r <action:string>",
    description: "Perform API requests",
    examples: EXAMPLES,
    options: OPTIONS,
    flags: FLAGS,
    run: handler,
};
