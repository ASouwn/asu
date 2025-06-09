import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig({
  plugins: [react()],
  build: {
    lib: {
      entry: 'src/RemoteComponent.jsx', // 远程组件入口
      name: 'RemoteComponent',          // 导出名称（可省略）
      fileName: () => 'RemoteComponent.js', // 构建输出的文件名
      formats: ['es'],                  // 构建为 ESM 格式
    },
    rollupOptions: {
      // 不打包 React 和 ReactDOM，让主站提供运行时环境
      external: ['react', 'react-dom'],
      output: {
        globals: {
          react: 'React',
          'react-dom': 'ReactDOM',
        },
      },
    },
    emptyOutDir: true,  // 每次清理 dist
    outDir: 'dist',     // 输出目录
  },
});
