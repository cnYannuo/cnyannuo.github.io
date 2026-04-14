# 仓库指南

## 项目结构与模块组织
本仓库是一个基于 Vite 构建的 Vue 3 应用。
- `src/main.js`：应用启动与挂载入口。
- `src/App.vue`：当前顶层页面组件。
- `src/components/`：可复用 Vue 组件（文件名使用 PascalCase，例如 `HelloWorld.vue`）。
- `src/components/icons/`：仅用于展示图标的组件。
- `src/assets/`：全局 CSS 与在源码中引用的静态资源。
- `public/`：原样对外提供的静态文件（例如 `favicon.ico`）。
- `index.html`：Vite 入口 HTML。

从 `src` 导入模块时使用 `@` 别名（在 `vite.config.js` 中配置），例如：`import Foo from '@/components/Foo.vue'`。

## 构建、测试与开发命令
- `npm install`：安装依赖（Node 版本要求 `^20.19.0 || >=22.12.0`）。
- `npm run dev`：启动本地 Vite 开发服务器（支持热更新）。
- `npm run build`：生成生产构建，输出到 `dist/`。
- `npm run preview`：本地预览已构建产物，用于验收。

提交 PR 前，至少执行一次 `npm run build`，以尽早发现编译期错误。

## 编码风格与命名规范
- Vue、JavaScript、CSS 统一使用 2 空格缩进。
- 优先使用 Vue 单文件组件（`.vue`），保持组件职责单一、体量适中。
- 组件文件名使用 `PascalCase.vue`；变量/函数使用 `camelCase`；常量使用 `UPPER_SNAKE_CASE`。
- 样式尽量限定在组件内部；公共/基础样式放在 `src/assets/`。
- 当前未配置 lint/format 工具；请遵循现有代码风格并保持最小改动面。

## 测试指南
当前仓库尚未配置自动化测试框架。现阶段请：
- 通过 `npm run dev` 与 `npm run build` 验证修改。
- 若涉及 UI 变更，请在 PR 中附上手动验证说明（测试页面/路由、浏览器、预期行为）。

若后续新增测试，建议使用 Vitest + Vue Test Utils，测试文件放在 `src/__tests__/`，命名采用 `*.spec.js`。

## 提交与合并请求规范
近期提交信息风格不统一，后续请统一规范：
- 提交信息使用清晰的祈使句，建议采用 Conventional Commits（例如：`feat: add hero section`）。
- 每次提交只包含一个逻辑变更，避免混入无关修改。
- PR 应包含：变更摘要、关键文件说明、已执行的验证步骤；涉及视觉改动时附截图。
- 如有对应 issue 或任务，请在 PR 中关联。
