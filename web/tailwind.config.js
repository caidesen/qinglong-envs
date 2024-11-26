import { nextui } from "@nextui-org/theme"

/** @type {import('tailwindcss').Config} */
export default {
  darkMode: ["class"],
  content: [
    "./src/pages/**/*.{ts,tsx}",
    "./src/components/**/*.{ts,tsx}",
    "./node_modules/@nextui-org/theme/dist/components/(button|tabs|ripple|spinner).js",
  ],
  theme: {
    extend: {},
  },
  plugins: [nextui()],
}
