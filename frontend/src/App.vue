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
  deleteTodo,
  clearCompleted,
  toggleAll
} = useTodos();

const filteredTodos = computed(() => {
  if (currentFilter.value === 'Active') {
    return todos.value.filter(t => !t.completed);
  } else if (currentFilter.value === 'Completed') {
    return todos.value.filter(t => t.completed);
  }
  return todos.value;
});

const allCompleted = computed(() => {
  return todos.value.length > 0 && todos.value.every(t => t.completed);
});

onMounted(() => {
  fetchTodos();
});
</script>

<template>
  <div class="min-h-screen p-4 bg-gradient-to-br from-indigo-50 via-white to-purple-50 md:p-8">
    <!-- Animated background elements -->
    <div class="fixed inset-0 overflow-hidden pointer-events-none">
      <div class="absolute bg-purple-200 rounded-full -top-40 -right-40 w-80 h-80 mix-blend-multiply filter blur-xl opacity-70 animate-blob"></div>
      <div class="absolute bg-yellow-200 rounded-full -bottom-40 -left-40 w-80 h-80 mix-blend-multiply filter blur-xl opacity-70 animate-blob animation-delay-2000"></div>
      <div class="absolute bg-pink-200 rounded-full top-1/2 left-1/3 w-80 h-80 mix-blend-multiply filter blur-xl opacity-70 animate-blob animation-delay-4000"></div>
    </div>
    <div class="relative max-w-2xl mx-auto">
      <!-- Header -->
      <header class="mb-8 text-center">
        <div class="flex items-center justify-between mb-6">
          <!-- Left side placeholder for symmetry -->
          <div class="w-20"></div>

          <!-- Logo/Icon -->
          <div class="inline-flex items-center justify-center p-2 shadow-lg bg-gradient-to-r from-purple-600 to-indigo-600 rounded-2xl">
            <div class="p-2 bg-white rounded-xl">
              <svg class="w-8 h-8 text-purple-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
              </svg>
            </div>
          </div>

          <!-- Right side placeholder for symmetry -->
          <div class="w-20"></div>
        </div>

        <h1 class="mb-3 text-5xl font-bold text-transparent bg-gradient-to-r from-purple-600 to-indigo-600 bg-clip-text">
          TaskMaster Pro
        </h1>
        <p class="text-lg text-gray-600">Stay organized and productive</p>
      </header>

      <!-- Main card -->
      <div class="overflow-hidden border shadow-2xl bg-white/90 backdrop-blur-xl rounded-3xl border-white/50">
        <!-- Stats bar -->
        <div class="px-6 py-4 border-b border-gray-100 bg-gradient-to-r from-purple-50 to-indigo-50">
          <div class="flex flex-wrap items-center justify-between gap-4">
            <div class="flex items-center space-x-6">
              <div class="text-center">
                <div class="text-2xl font-bold text-purple-700">{{ activeTodos }}</div>
                <div class="text-xs tracking-wider text-gray-600 uppercase">Active</div>
              </div>
              <div class="w-px h-8 bg-gray-300"></div>
              <div class="text-center">
                <div class="text-2xl font-bold text-green-600">{{ completedTodos }}</div>
                <div class="text-xs tracking-wider text-gray-600 uppercase">Done</div>
              </div>
              <div class="w-px h-8 bg-gray-300"></div>
              <div class="text-center">
                <div class="text-2xl font-bold text-gray-700">{{ todos.length }}</div>
                <div class="text-xs tracking-wider text-gray-600 uppercase">Total</div>
              </div>
            </div>
            
            <div class="flex items-center space-x-2">
              <button
                v-if="completedTodos > 0"
                @click="clearCompleted"
                class="flex items-center px-4 py-2 space-x-2 text-sm font-semibold text-white transition-all duration-200 bg-gradient-to-r from-red-500 to-pink-500 rounded-xl hover:shadow-lg active:scale-95"
              >
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                </svg>
                <span>Clear Completed</span>
              </button>
              
              <button
                v-if="todos.length > 0"
                @click="toggleAll"
                class="flex items-center px-4 py-2 space-x-2 text-sm font-semibold text-white transition-all duration-200 bg-gradient-to-r from-indigo-500 to-purple-500 rounded-xl hover:shadow-lg active:scale-95"
              >
                <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
                </svg>
                <span>{{ allCompleted ? 'Uncheck All' : 'Check All' }}</span>
              </button>
            </div>
          </div>
        </div>

        <!-- Error alert -->
        <div 
          v-if="error" 
          class="p-4 mx-6 mt-6 border border-red-200 shadow-sm bg-gradient-to-r from-red-50 to-pink-50 rounded-xl"
        >
          <div class="flex items-center">
            <div class="flex-shrink-0">
              <svg class="w-5 h-5 text-red-400" fill="currentColor" viewBox="0 0 20 20">
                <path fill-rule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clip-rule="evenodd" />
              </svg>
            </div>
            <div class="ml-3">
              <p class="text-sm text-red-700">{{ error }}</p>
            </div>
          </div>
        </div>

        <!-- Main content -->
        <div class="p-6 md:p-8">
          <TodoInput 
            :disabled="loading"
            @add="addTodo"
          />

          <div class="mt-8">
            <TodoFilter v-model="currentFilter" />
            
            <!-- Loading state -->
            <div v-if="loading" class="mt-6 space-y-4">
              <div v-for="i in 3" :key="i" class="animate-pulse">
                <div class="h-16 bg-gray-100 rounded-xl"></div>
              </div>
            </div>

            <TodoList 
              v-else
              :todos="filteredTodos"
              :loading="loading"
              @toggle="updateTodo"
              @delete="deleteTodo"
            />

            <!-- Empty state -->
            <div 
              v-if="!loading && filteredTodos.length === 0"
              class="py-12 mt-8 text-center"
            >
              <div class="flex items-center justify-center w-24 h-24 mx-auto mb-4 rounded-full bg-gradient-to-r from-purple-100 to-indigo-100">
                <svg class="w-12 h-12 text-purple-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                </svg>
              </div>
              <h3 class="mb-2 text-xl font-semibold text-gray-700">
                {{
                  currentFilter === 'Active' ? 'No active tasks' :
                  currentFilter === 'Completed' ? 'No completed tasks' :
                  'No tasks yet'
                }}
              </h3>
              <p class="max-w-sm mx-auto text-gray-500">
                {{
                  currentFilter === 'All' ? 'Start by adding your first task above!' :
                  currentFilter === 'Active' ? 'All tasks are completed! Great job! 🎉' :
                  'Complete some tasks to see them here!'
                }}
              </p>
            </div>
          </div>

          <!-- Tips section -->
          <div v-if="todos.length > 0" class="pt-6 mt-8 border-t border-gray-100">
            <div class="flex flex-col items-center justify-between gap-4 text-sm text-gray-600 md:flex-row">
              <div class="flex flex-wrap items-center gap-4">
                <div class="flex items-center space-x-1">
                  <svg class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                  <span>Click task to toggle</span>
                </div>
                <div class="flex items-center space-x-1">
                  <svg class="w-4 h-4 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                  </svg>
                  <span>Hover to delete</span>
                </div>
              </div>
              <div class="text-right">
                <div class="font-medium text-gray-700">{{ todos.length }} tasks</div>
                <div class="text-xs text-gray-500">Double-click to edit</div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Footer -->
      <footer class="mt-8 text-sm text-center text-gray-600">
        <p class="mt-1 text-xs text-gray-500">Press Enter to add, Esc to cancel, Delete to remove</p>
      </footer>
    </div>
  </div>
</template>

<style>
@keyframes blob {
  0% {
    transform: translate(0px, 0px) scale(1);
  }
  33% {
    transform: translate(30px, -50px) scale(1.1);
  }
  66% {
    transform: translate(-20px, 20px) scale(0.9);
  }
  100% {
    transform: translate(0px, 0px) scale(1);
  }
}

.animate-blob {
  animation: blob 7s infinite;
}

.animation-delay-2000 {
  animation-delay: 2s;
}

.animation-delay-4000 {
  animation-delay: 4s;
}
</style>