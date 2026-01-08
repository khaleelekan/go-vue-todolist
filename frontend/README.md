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

    Backend Files Breakdown
/backend/cmd/server/main.go
Entry point - initializes server and routes
/backend/internal/models/todo.go
Todo struct and model definitions
/backend/internal/store/memory_store.go
In-memory data storage implementation
/backend/internal/handlers/todo_handler.go
HTTP handlers for CRUD operations
/backend/internal/middleware/cors.go
CORS middleware configuration
Frontend Files Breakdown
/frontend/src/components/

TodoList.vue - Renders list of todos
TodoItem.vue - Individual todo item component
TodoInput.vue - Input field for new todos
TodoFilter.vue - Filter buttons (All/Active/Completed)

/frontend/src/services/api.js
API service layer for backend communication
/frontend/src/composables/useTodos.js
Vue composition API for todo logic (reusable)
/frontend/src/App.vue
Main application component
/frontend/src/main.js
Vue app initialization