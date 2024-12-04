import { defineConfig } from "vite"
import react from "@vitejs/plugin-react"
import { TanStackRouterVite } from "@tanstack/router-plugin/vite"

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    react(),
    TanStackRouterVite({
      routesDirectory: "src/pages",
      routeFileIgnorePattern: "components",
    }),
  ],
  server: {
    proxy: {
      "/api": "http://localhost:3000",
    },
  },
  resolve: {
    alias: {
      "@": "/src",
    },
  },
  build: {
    outDir: "dist",
  },
})
