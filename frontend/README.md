# Vue 3 + Vite

This template should help get you started developing with Vue 3 in Vite. The template uses Vue 3 `<script setup>` SFCs, check out the [script setup docs](https://v3.vuejs.org/api/sfc-script-setup.html#sfc-script-setup) to learn more.

Learn more about IDE Support for Vue in the [Vue Docs Scaling up Guide](https://vuejs.org/guide/scaling-up/tooling.html#ide-support).

todo-app/
├── backend/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   ├── internal/
│   │   ├── handlers/
│   │   │   └── todo_handler.go
│   │   ├── models/
│   │   │   └── todo.go
│   │   ├── store/
│   │   │   └── memory_store.go
│   │   └── middleware/
│   │       └── cors.go
│   ├── go.mod
│   ├── go.sum
│   └── README.md
│
└── frontend/
    ├── public/
    │   └── index.html
    ├── src/
    │   ├── components/
    │   │   ├── TodoList.vue
    │   │   ├── TodoItem.vue
    │   │   ├── TodoInput.vue
    │   │   └── TodoFilter.vue
    │   ├── services/
    │   │   └── api.js
    │   ├── composables/
    │   │   └── useTodos.js
    │   ├── App.vue
    │   └── main.js
    ├── package.json
    ├── vite.config.js
    ├── tailwind.config.js
    ├── postcss.config.js
    └── README.md