<template>
  <div class="mb-8">
    <!-- Filter buttons -->
    <div class="flex gap-2 mb-4">
      <button 
        v-for="filter in filters" 
        :key="filter"
        @click="$emit('update:modelValue', filter)"
        :class="[
          'group relative px-5 py-3 rounded-xl font-medium transition-all duration-200 active:scale-95 flex items-center justify-center gap-2 overflow-hidden',
          modelValue === filter 
            ? 'bg-gradient-to-r from-purple-500 to-indigo-500 text-white shadow-lg shadow-purple-200' 
            : 'bg-gray-100 text-gray-600 hover:bg-gray-200 hover:shadow-md'
        ]"
      >
        <!-- Active indicator (for active filter) -->
        <div 
          v-if="modelValue === filter"
          class="absolute top-0 left-0 w-full h-1 bg-gradient-to-r from-purple-400 to-indigo-400"
        ></div>
        
        <!-- Filter icon -->
        <svg 
          :class="[
            'w-5 h-5 transition-transform duration-200',
            modelValue === filter ? 'transform group-hover:scale-110' : 'group-hover:scale-110'
          ]"
          fill="none" 
          stroke="currentColor" 
          viewBox="0 0 24 24"
        >
          <path v-if="filter === 'All'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2V6zM14 6a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2V6zM4 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2H6a2 2 0 01-2-2v-2zM14 16a2 2 0 012-2h2a2 2 0 012 2v2a2 2 0 01-2 2h-2a2 2 0 01-2-2v-2z" />
          <path v-if="filter === 'Active'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
          <path v-if="filter === 'Completed'" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
        </svg>
        
        <!-- Filter text -->
        <span class="font-semibold">{{ filter }}</span>
        
        <!-- Count badge -->
        <span 
          v-if="getCount(filter) > 0"
          :class="[
            'px-2 py-0.5 text-xs rounded-full font-bold transition-all duration-200',
            modelValue === filter 
              ? 'bg-white/20 text-white' 
              : 'bg-gray-200 text-gray-700'
          ]"
        >
          {{ getCount(filter) }}
        </span>
      </button>
    </div>
    
    <!-- Active filter indicator -->
    <div class="flex items-center justify-between px-1 text-sm text-gray-500">
      <div class="flex items-center gap-2">
        <div class="flex items-center gap-1">
          <div class="w-2 h-2 bg-purple-500 rounded-full"></div>
          <span>Currently viewing: </span>
        </div>
        <span class="font-semibold text-purple-600">
          {{ modelValue }} ({{ getCount(modelValue) }})
        </span>
      </div>
      
      <!-- Stats summary -->
      <div class="flex items-center gap-4">
        <div class="flex items-center gap-1">
          <svg class="w-4 h-4 text-green-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7" />
          </svg>
          <span>{{ completedCount }} completed</span>
        </div>
        <div class="flex items-center gap-1">
          <svg class="w-4 h-4 text-blue-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          <span>{{ activeCount }} active</span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed, defineProps, defineEmits } from 'vue';

const props = defineProps({
  modelValue: {
    type: String,
    default: 'All'
  },
  todos: {
    type: Array,
    default: () => []
  }
});

const emit = defineEmits(['update:modelValue']);

const filters = ['All', 'Active', 'Completed'];

// Compute counts for each filter
const getCount = (filter) => {
  if (!props.todos || props.todos.length === 0) return 0;
  
  switch(filter) {
    case 'All':
      return props.todos.length;
    case 'Active':
      return props.todos.filter(t => !t.completed).length;
    case 'Completed':
      return props.todos.filter(t => t.completed).length;
    default:
      return 0;
  }
};

// Compute total counts for stats
const completedCount = computed(() => {
  return props.todos.filter(t => t.completed).length;
});

const activeCount = computed(() => {
  return props.todos.filter(t => !t.completed).length;
});
</script>