import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import path from 'path';

// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
      '@assets': path.resolve(__dirname, './src/assets'),
      '@models': path.resolve(__dirname, './src/models'),
      '@services': path.resolve(__dirname, './src/services'),
      '@stores': path.resolve(__dirname, './src/stores'),
      '@icons': path.resolve(__dirname, './src/shared/icons'),
      '@ui-kit': path.resolve(__dirname, './src/shared/ui-kit'),
    },
  },
  plugins: [react()],
  server: {
    proxy: {
      '/api': {
        target: 'http://57b41e19c067.vps.myjino.ru:49472',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ''),
      },
      '/ws': {
        target: 'ws://57b41e19c067.vps.myjino.ru:49472',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/ws/, ''),
      },
    },
  },
})
