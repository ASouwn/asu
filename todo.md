# ToDo

## 🎯 框架目标再确认

你要实现的框架具备以下能力：

| 能力                 | 说明                              |
| ------------------ | ------------------------------- |
| ⛓️ 文件系统路由          | 根据 `app/` 目录结构自动注册路由（如 Next.js） |
| ⚛️ 支持 React        | 允许开发者写 `.tsx/.jsx` 页面，自动编译运行    |
| 🎨 支持 Tailwind CSS | 可以写类名如 `text-red-500`，自动处理样式生成  |
| 🔧 用 Go 实现         | 构建系统、开发服务器、路由注册等由 Go 编写         |
| 💡 可扩展性强           | 用来学习前端核心逻辑，验证自定义实现想法            |

---

## 📦 你要实现的核心模块（分工明确）

我们可以把整个系统拆成几个子系统：

### ✅ 1. **目录结构 + 路由系统（文件路由）**

解析 `app/` 目录结构，自动注册路由。

| 示例文件路径              | 对应路由        |
| ------------------- | ----------- |
| `app/index.tsx`     | `/`         |
| `app/about.tsx`     | `/about`    |
| `app/blog/[id].tsx` | `/blog/:id` |

你需要实现：

* Go 递归扫描 `app/` 文件夹
* 提取路径 + 动态参数（如 `[id].tsx`）
* 构建路由映射表 `{ path: "/blog/:id", file: "app/blog/[id].tsx" }`

---

### ✅ 2. **JSX/TSX 构建系统（esbuild 集成）**

把用户写的 React 页面编译成浏览器能执行的 JS。

你需要：

* 用 Go 调用 `esbuild`（通过 Go API 或子进程）
* 支持编译 `.tsx`, `.jsx`, `.ts`, `.js`
* 输出到 `.output/` 文件夹
* 支持开发模式自动重建

参考用法（调用 esbuild）：

```go
esbuild.Build(esbuild.BuildOptions{
  EntryPoints: []string{"app/index.tsx"},
  Bundle: true,
  Outfile: "dist/main.js",
  ...
})
```

---

### ✅ 3. **HTML 模板系统**

提供一个标准 HTML Shell，把 React bundle 插进去：

```html
<!-- template.html -->
<html>
  <head>
    <title>{{.Title}}</title>
    <script defer src="/static/{{.Script}}"></script>
  </head>
  <body>
    <div id="root"></div>
  </body>
</html>
```

你需要：

* 在每个路由请求时渲染这个模板
* 填充动态字段（如页面标题、JS 入口）
* 支持传递初始数据（可用于 SSR 扩展）

---

### ✅ 4. **Tailwind CSS 支持**

让用户可以使用 Tailwind 类名（如 `bg-blue-500`）并生成真实样式。

你有两个实现路径：

#### 方案 A（推荐）

* 用户写自己的 `tailwind.config.js`
* 你在打包阶段调用 `tailwindcss -i ./input.css -o ./output.css`
* 然后 `<link rel="stylesheet" href="/static/output.css">` 插入 HTML

#### 方案 B（进阶）

* 自己实现 Tailwind 类扫描（复杂）
* 生成只包含用到类名的 CSS（需解析 `.tsx` 文件）

---

### ✅ 5. **开发服务器（Go 实现）**

启动一个本地服务：

* 监听路由（如 `/about`）
* 返回对应编译后的 HTML + JS
* 处理静态资源（如 CSS、图片、字体）
* 实现自动重载（可选）

可以使用 `net/http`：

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
  // 查找路由 -> 返回页面
})
```

---

### ✅ 6. **热重载 / 自动构建（可选但提升体验）**

监听 `app/`、`public/`、`tailwind.config.js` 变化，重新构建。

用 Go 实现文件监听：

```go
fsnotify.Watch("app/")
```

配合浏览器端 JS 实现热重载（或最简单页面刷新）。

---

## 🧩 最小项目结构（建议）

```
your-framework/
├── app/                    # 所有页面
│   ├── index.tsx
│   ├── about.tsx
│   └── blog/[id].tsx
├── public/                 # 静态资源
│   └── favicon.png
├── template.html           # HTML shell 模板
├── tailwind.config.js
├── input.css               # Tailwind entry
├── go.mod
├── main.go                 # 服务主程序
└── build/                  # 编译结果
```

---

## 💡 可选扩展功能（根据你的学习需求）

| 功能            | 学习价值                              |
| ------------- | --------------------------------- |
| SSR 支持        | 学习 Next.js 如何用 Node.js 做服务端渲染     |
| layout.tsx 支持 | 学习布局嵌套、共享 UI 模型                   |
| 动态 props 传参   | 实现 Next.js 的 `getServerSideProps` |
| 图片资源优化        | 模拟 `next/image` 懒加载、自动压缩          |
| 插件机制          | 支持用户扩展编译器行为、注册生命周期钩子              |

---

## 🧠 易用性建议

1. **开箱即用：**

   * 提供默认模板项目
   * `go run main.go` 就能起服务
2. **清晰报错：**

   * 提示 TS 编译失败、页面路径冲突等
3. **开发文档：**

   * 清晰说明如何添加页面、写样式、调试
4. **调试友好：**

   * 输出页面路径表（类似 Vite 的启动日志）

---

## ✅ 最后总结

你要做的是一个简化版的「Next.js-like 框架」，用 Go 替代 Node.js 来承担**开发服务器、编译管理、静态服务**，这是非常适合系统性学习 Web 框架的实践。

---

如果你希望我：

* 写一个最小可运行的原型项目
* 给出具体模块代码模板（如 `router.go`, `builder.go` 等）
* 用 esbuild 编译 React 的 Go 调用样例
  只需要告诉我，我可以分模块帮你构建！

是否想从 “最小可运行框架” 开始写一版？我可以帮你逐步搭建。
