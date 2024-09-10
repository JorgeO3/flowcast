// NOTE: This file contains the template for creating a new command.
import { Args } from "jsr:@std/cli/parse-args";
import { Action, Command, CommandFlags, Flag } from "../utils/types.ts";

const REQUIRED_FLAGS: string[] = [];

const ALL_FLAGS: Flag[] = [];

const OPTIONS: CommandFlags = {
	all: ALL_FLAGS,
	required: ALL_FLAGS.filter(({ flag }) => REQUIRED_FLAGS.includes(flag)),
};

const ACTIONS: Action[] = [];

const EXAMPLES: string[] = [
	"deno run -A requests/main.ts help",
	"deno run -A requests/main.ts request --help",
	"deno run -A requests/main.ts request --url http://localhost:3000/acts --type createAct --path ./data/acts.json  --index 1",
];

function handler(_args: Args) {}

export const GENERATE_COMMAND: Command = {
	name: "",
	options: OPTIONS,
	actions: ACTIONS,
	examples: EXAMPLES,
	run: handler,
};
