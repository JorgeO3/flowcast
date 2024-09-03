import { parseArgs } from "jsr:@std/cli/parse-args";

const cmds = [];

function parse(args: string[]) {
  parseArgs(args, {});
}

// deno run -A requests/main.ts <cmd> <args> <flags>
// deno run -A requests/main.ts generate acts
// deno run -A requests/main.ts generate acts --count 10
// deno run -A requests/main.ts request --variant create_act
// deno run -A requests/main.ts request --varaint get_act_by_id
// deno run -A requests/main.ts request --variant update_act
// deno run -A requests/main.ts request --variant delete_act

if (import.meta.main) {
  console.log(Deno.args);
}
