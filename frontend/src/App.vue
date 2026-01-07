<script setup>
import { ref, computed, onMounted } from 'vue';
import { useTodos } from './composables/useTodos';
import TodoInput from './components/TodoInput.vue';
import TodoFilter from './components/TodoFilter.vue';
import TodoList from './components/TodoList.vue';

const currentFilter = ref('All');

const {
  todos,
  loading,
  error,
  activeTodos,
  completedTodos,
  fetchTodos,
  addTodo,
  updateTodo,
  deleteTodo
} = useTodos();

const filteredTodos = computed(() => {
  if (currentFilter.value === 'Active') {
    return todos.value.filter(t => !t.completed);
  } else if (currentFilter.value === 'Completed') {
    return todos.value.filter(t => t.completed);
  }
  return todos.value;
});

onMounted(() => {
  fetchTodos();
});
</script>

<template>
  <div class="bg-gradient-to-br from-purple-600 via-purple-700 to-indigo-800 min-h-screen p-6">
    <div class="max-w-2xl mx-auto">
      <div class="bg-white rounded-2xl shadow-2xl p-8">
        <h1 class="text-4xl font-bold text-gray-800 mb-8 text-center">
          📝 My Todos
        </h1>

        <div 
          v-if="error" 
          class="bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded-lg mb-6"
        >
          {{ error }}
        </div>

        <TodoInput 
          :disabled="loading"
          @add="addTodo"
        />

        <TodoFilter v-model="currentFilter" />

        <TodoList 
          :todos="filteredTodos"
          :loading="loading"
          @toggle="updateTodo"
          @delete="deleteTodo"
        />

        <div 
          v-if="todos.length > 0" 
          class="mt-6 pt-6 border-t-2 border-gray-100 text-center text-sm text-gray-600"
        >
          <span class="font-semibold text-purple-600">{{ activeTodos }}</span> active • 
          <span class="font-semibold text-green-600">{{ completedTodos }}</span> completed • 
          <span class="font-semibold text-gray-700">{{ todos.length }}</span> total
        </div>
      </div>
    </div>
  </div>
</template>