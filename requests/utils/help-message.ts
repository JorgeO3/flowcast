import { bold, rgb8, underline } from "@std/fmt/colors";
import { GENERATE_COMMAND, REQUEST_COMMAND } from "../commands/index.ts";
import { Command } from "./types.ts";

// Configuración de colores y estilos
type ColorConfig = Record<string, number>;

const DEFAULT_COLORS: ColorConfig = {
  "grey": 15, // #FFFFFF
  "white": 231, // #E5E5E5
  "blue": 45, // #008DF8
  "green": 82, // #8CE10B
  "yellow": 214, // #FFB900
  "red": 160, // #FF000F
  "darkGray": 8, // #232323 (Referencial, no se aplica directamente)
};

const createColorizer = (colors: ColorConfig) => ({
  colorize: (text: string, colorCode: keyof ColorConfig) => rgb8(text, colors[colorCode]),
  colorBold: (text: string, colorCode: keyof ColorConfig) => bold(rgb8(text, colors[colorCode])),
  colorUnderline: (text: string, colorCode: keyof ColorConfig) =>
    underline(rgb8(text, colors[colorCode])),
});

const colorizer = createColorizer(DEFAULT_COLORS);

const COMMANDS: Command[] = [REQUEST_COMMAND, GENERATE_COMMAND];
const EXAMPLES: string[] = [
  "deno run -A requests/main.ts help",
  ...COMMANDS.flatMap(({ examples = [] }) => examples),
];

// Funciones de formato personalizables
const color = {
  usage: () => colorizer.colorBold("Usage:", "blue"),
  commands: () => colorizer.colorBold("COMMANDS", "blue"),
  flags: () => colorizer.colorBold("FLAGS", "blue"),
  description: (desc: string) => colorizer.colorize(desc, "grey"),
  optionFlag: (flag: string) => colorizer.colorize(flag, "yellow"),
  optionTitle: (opt: string) => colorizer.colorBold(opt, "blue"),
  example: () => colorizer.colorBold("EXAMPLES", "blue"),
  commandName: (name: string) => colorizer.colorBold(colorizer.colorize(name, "white"), "white"),
  actionDescription: (desc: string) => colorizer.colorize(desc, "grey"),
  exampleCommand: (cmd: string) => colorizer.colorize(cmd, "white"),
  required: () => colorizer.colorBold("required", "white"),
  error: () => colorizer.colorBold("Error:", "red"),
  errorCommand: (cmd: string) => colorizer.colorBold(cmd, "red"),
};

// Función para resaltar "required" en descripciones
const highlightRequired = (description: string) =>
  description.replace("required", color.required());

// Función para generar la ayuda de un comando
const generateCommandHelp = ({ name, flags, options, description }: Command) => {
  let helpMessage = `${color.commandName(name)}\n  ${
    color.description(description)
  }\n\n  ${color.flags()}\n`;

  flags.all.forEach(({ flag, description }) => {
    helpMessage += `    ${color.optionFlag(flag)}  ${
      description.includes("required") ? highlightRequired(description) : description
    }\n`;
  });

  if (options) {
    helpMessage += `\n  ${color.optionTitle(options.name)}\n`;
    options.values.forEach(({ name, description }) => {
      helpMessage += `    ${color.commandName(name + ":")} ${
        color.actionDescription(description)
      }\n`;
    });
  }

  helpMessage += "\n";
  return helpMessage;
};

// Función para generar el mensaje de ayuda general
const generateHelpMessage = (commands: Command[], examples: string[]) => {
  let helpMessage =
    `${color.usage()} deno run main.ts <command> [options]\n\n${color.commands()}\n`;

  commands.forEach((command) => {
    helpMessage += generateCommandHelp(command);
  });

  helpMessage += `${color.example()}\n`;
  examples.forEach((example) => {
    helpMessage += `  ${color.exampleCommand(example)}\n`;
  });

  return helpMessage;
};

// Función para mostrar la ayuda general
export const showGeneralHelp = () => console.log(generateHelpMessage(COMMANDS, EXAMPLES));

// Función para mostrar errores crudos
export const showRawError = (msg: string, command: string) =>
  console.error(`\n ${color.error()} ${msg} ${color.errorCommand(command)}`);

// Permite agregar más comandos dinámicamente o cambiar el comportamiento sin modificar el código principal
export const addCommand = (command: Command) => {
  COMMANDS.push(command);
};

// Función para establecer colores personalizados
export const setCustomColors = (colors: Partial<ColorConfig>) => {
  Object.assign(DEFAULT_COLORS, colors);
};
