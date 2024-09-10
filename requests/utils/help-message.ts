import { bold, rgb8 } from "@std/fmt/colors";
import { GENERATE_COMMAND, REQUEST_COMMAND } from "../commands/index.ts";
import { Command } from "./types.ts";

// Configuración de colores y estilos
type ColorConfig = Record<string, number>;

const DEFAULT_COLORS: ColorConfig = {
  green: 118,
  orange: 214,
  yellow: 220,
  blue: 39,
  cyan: 45,
  white: 252,
  red: 196,
  darkOrange: 208,
  lightGreen: 83,
};

// deno-fmt-ignore
const createColorizer = (colors: ColorConfig) => ({
  colorize: (text: string, colorCode: keyof ColorConfig) => rgb8(text, colors[colorCode]),
  colorBold: (text: string, colorCode: keyof ColorConfig) => bold(rgb8(text, colors[colorCode])),
});

const colorizer = createColorizer(DEFAULT_COLORS);

const COMMANDS: Command[] = [REQUEST_COMMAND, GENERATE_COMMAND];
const EXAMPLES: string[] = [
  "deno run -A requests/main.ts help",
  ...COMMANDS.flatMap(({ examples = [] }) => examples),
];

// Funciones de formato personalizables
const color = {
  usage: () => colorizer.colorBold("Usage:", "green"),
  commands: () => colorizer.colorBold("COMMANDS", "orange"),
  description: () => colorizer.colorize("Perform API requests", "white"),
  optionFlag: (flag: string) => colorizer.colorize(flag, "yellow"),
  actionsTitle: () => colorizer.colorBold("ACTIONS", "orange"),
  example: () => colorizer.colorBold("EXAMPLES", "blue"),
  commandName: (name: string) => colorizer.colorize(name, "cyan"),
  actionDescription: (desc: string) => colorizer.colorize(desc, "white"),
  exampleCommand: (cmd: string) => colorizer.colorize(cmd, "lightGreen"),
  required: () => colorizer.colorBold("required", "red"),
  error: () => colorizer.colorBold("Error:", "red"),
  errorCommand: (cmd: string) => colorizer.colorBold(cmd, "darkOrange"),
};

// Función para resaltar "required" en descripciones
const highlightRequired = (description: string) =>
  description.replace("required", color.required());

// Función para generar la ayuda de un comando
const generateCommandHelp = ({ name, options, actions }: Command) => {
  let helpMessage = `${
    color.commandName(name)
  }\n ${color.description()}\n ${color.commands()}\n`;

  options.all.forEach(({ flag, description }) => {
    helpMessage += ` ${color.optionFlag(flag)} ${
      description.includes("required")
        ? highlightRequired(description)
        : description
    }\n`;
  });

  if (actions) {
    helpMessage += ` ${color.actionsTitle()}\n`;
    actions.forEach(({ name, description }) => {
      helpMessage += ` ${color.commandName(name + ":")} ${
        color.actionDescription(description)
      }\n`;
    });
  }

  return helpMessage + "\n";
};

// Función para generar el mensaje de ayuda general
const generateHelpMessage = (commands: Command[], examples: string[]) => {
  let helpMessage =
    `${color.usage()} deno run main.ts <command> [options]\n\n${color.commands()}\n\n`;

  commands.forEach((command) => {
    helpMessage += generateCommandHelp(command);
  });

  helpMessage += `${color.example()}\n`;
  examples.forEach((example) => {
    helpMessage += ` ${color.exampleCommand(example)}\n`;
  });

  return helpMessage;
};

// Función para mostrar la ayuda general
export const showGeneralHelp = () =>
  console.log(generateHelpMessage(COMMANDS, EXAMPLES));

// Función para mostrar errores crudos
export const showRawError = (msg: string, command: string) =>
  console.error(`\n ${color.error()} ${msg} ${color.errorCommand(command)}`);

// Permite agregar más comandos dinámicamente o cambiar el comportamiento sin modificar el código principal
export const addCommand = (command: Command) => {
  COMMANDS.push(command);
};

export const setCustomColors = (colors: Partial<ColorConfig>) => {
  Object.assign(DEFAULT_COLORS, colors);
};
