<template>
  <div>
    <div v-if="loading" class="text-center py-12">
      <div class="inline-block animate-spin rounded-full h-12 w-12 border-4 border-purple-600 border-t-transparent"></div>
      <p class="mt-4 text-gray-600">Loading todos...</p>
    </div>

    <ul v-else-if="todos.length > 0" class="space-y-3">
      <TodoItem
        v-for="todo in todos"
        :key="todo.id"
        :todo="todo"
        @toggle="$emit('toggle', $event)"
        @delete="$emit('delete', $event)"
      />
    </ul>

    <div v-else class="text-center py-16">
      <svg class="w-16 h-16 mx-auto mb-4 text-gray-300" fill="currentColor" viewBox="0 0 20 20">
        <path d="M9 2a1 1 0 000 2h2a1 1 0 100-2H9z"/>
        <path fill-rule="evenodd" d="M4 5a2 2 0 012-2 3 3 0 003 3h2a3 3 0 003-3 2 2 0 012 2v11a2 2 0 01-2 2H6a2 2 0 01-2-2V5zm3 4a1 1 0 000 2h.01a1 1 0 100-2H7zm3 0a1 1 0 000 2h3a1 1 0 100-2h-3zm-3 4a1 1 0 100 2h.01a1 1 0 100-2H7zm3 0a1 1 0 100 2h3a1 1 0 100-2h-3z" clip-rule="evenodd"/>
      </svg>
      <p class="text-gray-400 text-lg">No todos yet. Add one above!</p>
    </div>
  </div>
</template>

<script setup>
import TodoItem from './TodoItem.vue';

defineProps({
  todos: {
    type: Array,
    default: () => []
  },
  loading: {
    type: Boolean,
    default: false
  }
});

defineEmits(['toggle', 'delete']);
</script>
```

---

Make sure your folder structure looks like this:
```
frontend/src/
├── components/
│   ├── TodoInput.vue
│   ├── TodoFilter.vue
│   ├── TodoItem.vue
│   └── TodoList.vue
├── composables/
│   └── useTodos.js
├── services/
│   └── api.js
├── App.vue
├── main.js
└── style.css