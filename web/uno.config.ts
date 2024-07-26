import { defineConfig } from 'unocss'

export default defineConfig({
    theme: {
        colors: {
            border: "var(--border)",
            input: "var(--input)",
            ring: "var(--ring)",
            background: "var(--background)",
            foreground: "var(--foreground)",
            primary: {
                DEFAULT: "var(--primary)",
                foreground: "var(--primary-foreground)",
            },
            secondary: {
                DEFAULT: "var(--secondary)",
                foreground: "var(--secondary-foreground)",
            },
            destructive: {
                DEFAULT: "var(--destructive)",
                foreground: "var(--destructive-foreground)",
            },
            muted: {
                DEFAULT: "var(--muted)",
                foreground: "var(--muted-foreground)",
            },
            accent: {
                DEFAULT: "var(--accent)",
                foreground: "var(--accent-foreground)",
            },
            popover: {
                DEFAULT: "var(--popover)",
                foreground: "var(--popover-foreground)",
            },
            card: {
                DEFAULT: "var(--card)",
                foreground: "var(--card-foreground)",
            },
        },
        borderRadius: {
            xl: "calc(var(--radius) + 4px)",
            lg: "var(--radius)",
            md: "calc(var(--radius) - 2px)",
            sm: "calc(var(--radius) - 4px)",
        },
    }
})