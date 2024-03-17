import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import { fileURLToPath, URL } from 'url'



// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: [
      { find: '@', replacement: fileURLToPath(new URL('./src', import.meta.url)) },
      { find: '@sass', replacement: fileURLToPath(new URL('./src/assets/sass', import.meta.url)) },
      { find: '@components', replacement: fileURLToPath(new URL('./src/components', import.meta.url)) },
      { find: '@assets', replacement: fileURLToPath(new URL('./src/assets', import.meta.url)) },
      { find: '@images', replacement: fileURLToPath(new URL('./src/assets/images', import.meta.url)) },
      { find: '@icons', replacement: fileURLToPath(new URL('./src/assets/icons', import.meta.url)) },
    ],
  },
  plugins: [react()],
  server: {
    host: '0.0.0.0',
    port: 3000
  }
})