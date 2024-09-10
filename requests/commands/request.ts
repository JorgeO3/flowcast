import { Args } from "jsr:@std/cli/parse-args";
import { Command, CommandFlags, Flag } from "../utils/types.ts";

const REQUIRED_FLAGS: string[] = [];

const ALL_FLAGS: Flag[] = [];

const OPTIONS: CommandFlags = {
    all: ALL_FLAGS,
    required: ALL_FLAGS.filter(({ flag }) => REQUIRED_FLAGS.includes(flag)),
};

const EXAMPLES: string[] = [
    "deno run -A requests/main.ts request --url http://localhost:3000/acts --type createAct --path ./data/acts.json  --index 1",
];

function handler(_args: Args) {
}

export const REQUEST_COMMAND: Command = {
    name: "request, r <action:string>",
    options: OPTIONS,
    examples: EXAMPLES,
    run: handler,
};
