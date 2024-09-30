import { type Args, parseArgs } from "@std/cli/parse-args";
import { GENERATE_COMMAND, REQUEST_COMMAND } from "./commands/index.ts";
import {
  type Command,
  type CommandFlags,
  Flag,
  showGeneralHelp,
  showRawError,
} from "./utils/index.ts";

const COMMANDS: Map<string, Command> = new Map([
  [REQUEST_COMMAND.name, REQUEST_COMMAND],
  [GENERATE_COMMAND.name, GENERATE_COMMAND],
]);

function aliasToCommandName(alias: string): string {
  for (const [name, cmd] of COMMANDS) {
    if (cmd.name.includes(alias)) return name;
  }
  return alias;
}

function showErrorAndExit(message: string, details: string): never {
  showGeneralHelp();
  showRawError(message, details);
  Deno.exit(1);
}

function validateFlags(args: Args, flags: CommandFlags): void {
  const { required, all } = flags;

  for (const { long, short, flagType, flag } of all) {
    const value = args[short] || args[long];
    if (!value) continue;
    // deno-lint-ignore valid-typeof
    if (typeof value !== flagType) showErrorAndExit("Invalid type", flag);
  }

  // validate required fields
  for (const { long, short, flag } of required) {
    if (!args[short] && !args[long]) {
      showErrorAndExit("Missing required flag", flag);
    }
  }
}

function getLongFlag(rawFlag: string, allFlags: Flag[]): string {
  const formattedFlag = allFlags.find(({ short, long }) => short === rawFlag || long === rawFlag);
  if (!formattedFlag) {
    throw new Error(`Flag ${rawFlag} not found`);
  }
  return formattedFlag.long;
}

function parseFlags(args: Args, { flags }: Command): Record<string, string | number | boolean> {
  return Object.entries(args)
    .filter(([key]) => key !== "_")
    .reduce((acc, [key, value]) => {
      acc[getLongFlag(key, flags.all)] = value;
      return acc;
    }, {} as Record<string, string | number | boolean>);
}

function main(): void {
  const parsedArgs = parseArgs(Deno.args);
  const commandName = aliasToCommandName(parsedArgs._[0] as string);
  const command = COMMANDS.get(commandName);

  console.log({ parsedArgs });

  if (commandName === "help" || parsedArgs.help || parsedArgs.h || !command) {
    return showGeneralHelp();
  }

  if (!command) {
    return showErrorAndExit("Command not found", commandName);
  }

  validateFlags(parsedArgs, command.flags);
  const args = parseFlags(parsedArgs, command);
  console.log({ args });
  command.run(args);
}

if (import.meta.main) {
  main();
}

// Execute the following commands to test the code:
// deno run -A requests/main.ts help
// deno run -A requests/main.ts request --help
// deno run -A requests/main.ts request --url http://localhost:3000/acts --type createAct --path ./data/acts.json  --index 1
// deno run -A requests/main.ts generate --type act --count 10 --path ./data/acts.json
