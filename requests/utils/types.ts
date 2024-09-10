import { Args } from "jsr:@std/cli/parse-args";

type FlagType = "string" | "number" | "boolean";

export type Flag = {
    flag: string;
    description: string;
    flagType: FlagType;
};

export type CommandFlags = {
    required: Flag[];
    all: Flag[];
};

export type Action = {
    name: string;
    description: string;
};

type Handler = (args: Args) => void;

export type Command = {
    name: string;
    options: CommandFlags;
    actions?: Action[];
    examples?: string[];
    run: Handler;
};
