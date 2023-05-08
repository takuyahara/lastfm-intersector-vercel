import fs from "fs";
import { defineConfig } from "vite";
import solidPlugin from "vite-plugin-solid";
import liveReload from "vite-plugin-live-reload";

export default defineConfig({
  plugins: [solidPlugin(), liveReload(["**/*"])],
  server: {
    port: 3000,
    hmr: false,
  },
  build: {
    target: "esnext",
  },
});
