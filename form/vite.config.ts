import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'

export default defineConfig({
  server: {
    allowedHosts: ['6e8c-2409-8961-1f4a-bdc6-bcb8-5f45-fc3b-2cb.ngrok-free.app'],
  },
  // ...
  plugins: [
    // ...
    [vue()],
    AutoImport({
      resolvers: [ElementPlusResolver()],
    }),
    Components({
      resolvers: [ElementPlusResolver()],
    }),
  ],
})