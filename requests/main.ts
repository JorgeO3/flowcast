import { type Args, parseArgs } from "jsr:@std/cli/parse-args";
import { GENERATE_COMMAND, REQUEST_COMMAND } from "./commands/index.ts";
import {
  type Command,
  type CommandFlags,
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

  for (const { flag, flagType } of all) {
    const value = args[flag];
    if (!value) showErrorAndExit("Missing required flag", flag);
    // deno-lint-ignore valid-typeof
    if (typeof value !== flagType) showErrorAndExit("Invalid type", flag);
  }

  required.forEach(({ flag }) => {
    if (!args[flag]) showErrorAndExit("Missing required flag", flag);
  });
}

function main(): void {
  const parsedArgs = parseArgs(Deno.args);
  const commandName = aliasToCommandName(parsedArgs._[0] as string);
  const command = COMMANDS.get(commandName);

  if (commandName === "help" || parsedArgs.help || parsedArgs.h) {
    return showGeneralHelp();
  }

  if (!command) {
    return showErrorAndExit("Command not found", commandName);
  }

  validateFlags(parsedArgs, command.options);
  command.run(parsedArgs);
}

if (import.meta.main) {
  main();
}

// Execute the following commands to test the code:
// deno run -A requests/main.ts help
// deno run -A requests/main.ts request --help
// deno run -A requests/main.ts request --url http://localhost:3000/acts --type createAct --path ./data/acts.json  --index 1
// deno run -A requests/main.ts generate --type act --count 10 --path ./data/acts.json
