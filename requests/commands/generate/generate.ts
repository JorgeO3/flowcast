import type { Args, Command, CommandFlags, Flag } from "/utils/index.ts";
import { generateActs, saveActsToFile } from "./generate-acts.ts";

// Define available data types for generation
enum DataType {
    Act = "act",
    // Future types can be added here
}

// Required flags by their long names
const REQUIRED_FLAGS: string[] = ["type"];

// Definition of all flags
const ALL_FLAGS: Flag[] = [
    {
        flag: "-c, --count <count:number>",
        default: 10,
        short: "c",
        long: "count",
        flagType: "number",
        description: "Specify the number of items to generate (default: 10)",
    },
    {
        flag: "-p, --path <path:string>",
        description: "Specify the path to write generated data (default: ./data/acts.json)",
        short: "p",
        long: "path",
        default: "./data/acts.json",
        flagType: "string",
    },
    {
        flag: "-t, --type <type:string>",
        description: "Specify the type of data to generate (required)",
        short: "t",
        long: "type",
        default: "",
        flagType: "string",
    },
];

// Configuration of command flags
const FLAGS: CommandFlags = {
    all: ALL_FLAGS,
    required: ALL_FLAGS.filter(({ long }) => REQUIRED_FLAGS.includes(long)),
};

// Examples of how to use the command
const EXAMPLES: string[] = [
    "deno run -A requests/main.ts generate --type act --count 10 --path ./data/acts.json",
];

// Definition of the arguments required for action handlers
type HandlerArgs = {
    count: number;
    path: string;
    type: DataType;
};

// Handler function for the generate command
async function handler(args: Args): Promise<void> {
    const count = args.count as number;
    const path = args.path as string;
    const type = args.type as string;

    // Validate the data type
    if (!Object.values(DataType).includes(type as DataType)) {
        console.error(
            `Invalid type: "${type}". Supported types are: ${Object.values(DataType).join(", ")}`,
        );
        Deno.exit(1);
    }

    const dataType = type as DataType;

    try {
        let data: unknown[];

        switch (dataType) {
            case DataType.Act:
                data = await generateActs(count);
                break;
            // Future cases for other data types can be added here
            default:
                console.error(`Unsupported data type: "${dataType}"`);
                Deno.exit(1);
        }

        await saveActsToFile(path, data);
    } catch (error) {
        console.error("Error during data generation:", error);
        Deno.exit(1);
    }
}

// Definition of the GENERATE_COMMAND
export const GENERATE_COMMAND: Command = {
    name: "generate, g <type:string>",
    description: "Generate data for testing",
    examples: EXAMPLES,
    flags: FLAGS,
    run: handler,
};
