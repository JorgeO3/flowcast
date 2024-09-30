type FlagType = "string" | "number" | "boolean";

export type Flag = {
    flag: string;
    description: string;
    default: string | number | boolean;
    short: string;
    long: string;
    flagType: FlagType;
};

export type CommandFlags = {
    required: Flag[];
    all: Flag[];
};

export type Option = {
    name: string;
    description: string;
};

export type Options = {
    name: string;
    values: Option[];
};

export type Args = Record<string, string | number | boolean>;

type Handler = (args: Args) => void;

export type Command = {
    name: string;
    description: string;
    flags: CommandFlags;
    options?: Options;
    examples?: string[];
    run: Handler;
};
