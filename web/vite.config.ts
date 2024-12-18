import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import path from 'path';

// https://vitejs.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      '@': path.resolve(__dirname, './src'),
      '@api': path.resolve(__dirname, './src/api'),
      '@assets': path.resolve(__dirname, './src/assets'),
      '@models': path.resolve(__dirname, './src/models'),
      '@services': path.resolve(__dirname, './src/services'),
      '@stores': path.resolve(__dirname, './src/stores'),
      '@icons': path.resolve(__dirname, './src/shared/icons'),
      '@ui-kit': path.resolve(__dirname, './src/shared/ui-kit'),
      '@components': path.resolve(__dirname, './src/components'),
    },
  },
  plugins: [react()],
})
