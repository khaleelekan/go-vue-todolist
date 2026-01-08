<template>
  <div>
    <!-- Loading state -->
    <div v-if="loading" class="py-6 space-y-4">
      <!-- Skeleton loader -->
      <div 
        v-for="i in 3" 
        :key="i" 
        class="p-4 border-2 border-gray-100 animate-pulse rounded-xl"
      >
        <div class="flex items-center justify-between">
          <div class="flex items-center flex-1 gap-4">
            <div class="w-6 h-6 bg-gray-200 rounded-lg"></div>
            <div class="flex-1 space-y-2">
              <div class="w-3/4 h-4 bg-gray-200 rounded"></div>
              <div class="w-1/4 h-3 bg-gray-100 rounded"></div>
            </div>
          </div>
          <div class="w-20 h-8 bg-gray-200 rounded-lg"></div>
        </div>
      </div>
      
      <!-- Loading text -->
      <div class="pt-4 text-center">
        <div class="inline-flex items-center justify-center p-2 mb-3 rounded-full bg-gradient-to-r from-purple-100 to-indigo-100">
          <svg class="w-8 h-8 text-purple-600 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
          </svg>
        </div>
        <p class="font-medium text-gray-600">Loading your tasks...</p>
        <p class="mt-1 text-sm text-gray-500">Almost there!</p>
      </div>
    </div>

    <!-- Todo list -->
    <ul v-else-if="todos.length > 0" class="space-y-3">
      <TodoItem
        v-for="todo in todos"
        :key="todo.id"
        :todo="todo"
        @toggle="$emit('toggle', $event)"
        @delete="$emit('delete', $event)"
        @update="$emit('update', $event)"
      />
      
      <!-- List summary -->
      <div class="pt-4 mt-4 border-t border-gray-100">
        <div class="flex items-center justify-between text-sm">
          <div class="text-gray-600">
            Showing <span class="font-semibold text-purple-600">{{ todos.length }}</span> 
            {{ todos.length === 1 ? 'task' : 'tasks' }}
          </div>
          <div class="flex items-center gap-2">
            <div class="flex items-center gap-1">
              <div class="w-3 h-3 bg-green-500 rounded-full"></div>
              <span class="text-gray-600">
                {{ completedCount }} done
              </span>
            </div>
            <div class="w-px h-4 bg-gray-300"></div>
            <div class="flex items-center gap-1">
              <div class="w-3 h-3 bg-purple-500 rounded-full"></div>
              <span class="text-gray-600">
                {{ activeCount }} active
              </span>
            </div>
          </div>
        </div>
      </div>
    </ul>

    <!-- Empty state -->
    <div v-else class="px-4 py-16 text-center">
      <div class="max-w-md mx-auto">
        <div class="inline-flex items-center justify-center w-24 h-24 mb-6 rounded-full bg-gradient-to-r from-purple-100 to-indigo-100">
          <svg class="w-12 h-12 text-purple-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2" />
          </svg>
        </div>
        
        <h3 class="mb-3 text-2xl font-bold text-gray-800">
          {{ getEmptyStateTitle() }}
        </h3>
        
        <p class="mb-8 text-gray-600">
          {{ getEmptyStateMessage() }}
        </p>
        
        <!-- Quick action buttons -->
        <div class="flex flex-col justify-center gap-3 sm:flex-row">
          <button
            @click="$emit('add-example')"
            class="px-6 py-3 font-medium text-white transition-all duration-200 bg-gradient-to-r from-purple-500 to-indigo-500 rounded-xl hover:shadow-lg active:scale-95"
          >
            <svg class="inline-block w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v6m0 0v6m0-6h6m-6 0H6" />
            </svg>
            Add Example Tasks
          </button>
          
          <button
            @click="$emit('clear-all')"
            class="px-6 py-3 font-medium text-gray-700 transition-colors duration-200 border-2 border-gray-300 rounded-xl hover:bg-gray-50"
          >
            Clear All
          </button>
        </div>
        
        <!-- Tips -->
        <div class="pt-6 mt-8 border-t border-gray-100">
          <p class="mb-3 text-sm text-gray-500">Quick tips:</p>
          <div class="grid grid-cols-1 gap-3 text-sm sm:grid-cols-2">
            <div class="flex items-center gap-2 text-gray-600">
              <svg class="w-4 h-4 text-green-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
              </svg>
              <span>Click to mark complete</span>
            </div>
            <div class="flex items-center gap-2 text-gray-600">
              <svg class="w-4 h-4 text-purple-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
              </svg>
              <span>Double-click to edit</span>
            </div>
            <div class="flex items-center gap-2 text-gray-600">
              <svg class="w-4 h-4 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
              </svg>
              <span>Hover to delete</span>
            </div>
            <div class="flex items-center gap-2 text-gray-600">
              <svg class="w-4 h-4 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7h12m0 0l-4-4m4 4l-4 4m0 6H4m0 0l4 4m-4-4l4-4" />
              </svg>
              <span>Drag to reorder</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue';
import TodoItem from './TodoItem.vue';

const props = defineProps({
  todos: {
    type: Array,
    default: () => []
  },
  loading: {
    type: Boolean,
    default: false
  },
  filter: {
    type: String,
    default: 'All'
  }
});

const emit = defineEmits(['toggle', 'delete', 'update', 'add-example', 'clear-all']);

// Compute counts for summary
const completedCount = computed(() => {
  return props.todos.filter(todo => todo.completed).length;
});

const activeCount = computed(() => {
  return props.todos.filter(todo => !todo.completed).length;
});

// Dynamic empty state messages
const getEmptyStateTitle = () => {
  switch (props.filter) {
    case 'Active':
      return 'No active tasks';
    case 'Completed':
      return 'No completed tasks';
    default:
      return 'No tasks yet';
  }
};

const getEmptyStateMessage = () => {
  switch (props.filter) {
    case 'Active':
      return 'All tasks are completed! Great job! 🎉';
    case 'Completed':
      return 'Complete some tasks to see them here';
    default:
      return 'Start by adding your first task above!';
  }
};
</script>