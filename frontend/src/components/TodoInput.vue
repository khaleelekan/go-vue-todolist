<template>
  <div class="relative mb-6">
    <div class="relative flex items-center">
      <!-- Input field -->
      <input 
        v-model="inputValue" 
        @keyup.enter="handleAdd"
        type="text" 
        placeholder="What needs to be done?"
        :disabled="disabled"
        class="flex-1 py-4 pr-12 text-lg text-gray-800 placeholder-gray-400 transition-all duration-200 bg-white border-2 border-gray-200 shadow-sm pl-14 rounded-2xl focus:outline-none focus:border-purple-500 focus:ring-2 focus:ring-purple-200 disabled:bg-gray-100 disabled:cursor-not-allowed"
      >
      
      <!-- Plus icon on left -->
      <div class="absolute left-4">
        <svg class="w-6 h-6 text-purple-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
        </svg>
      </div>
      
      <!-- Add button on right -->
      <button 
        @click="handleAdd" 
        :disabled="!inputValue.trim() || disabled"
        class="absolute flex items-center px-4 py-2 space-x-2 font-semibold text-white transition-all duration-200 transform right-2 bg-gradient-to-r from-purple-500 to-indigo-500 rounded-xl hover:shadow-lg focus:outline-none focus:ring-2 focus:ring-purple-500 focus:ring-offset-2 disabled:bg-gray-300 disabled:cursor-not-allowed hover:scale-105 active:scale-95 disabled:hover:scale-100"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
        </svg>
        <span class="hidden sm:inline">Add</span>
      </button>
    </div>
    
    <!-- Helper text and character counter -->
    <div class="flex items-center justify-between mt-3 text-sm">
      <div class="flex items-center space-x-4">
        <span class="flex items-center text-gray-500">
          <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
          </svg>
          Press Enter to add
        </span>
        <span class="flex items-center text-gray-500">
          <svg class="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4" />
          </svg>
          Esc to cancel
        </span>
      </div>
      
      <span 
        class="px-2 py-1 font-medium rounded-full"
        :class="[
          inputValue.length > 100 
            ? 'bg-red-100 text-red-600'
            : inputValue.length > 50
              ? 'bg-yellow-100 text-yellow-600'
              : 'bg-gray-100 text-gray-600'
        ]"
      >
        {{ inputValue.length }}/100
      </span>
    </div>
  </div>
</template>

<script setup>
import { ref, watch } from 'vue';

const props = defineProps({
  disabled: {
    type: Boolean,
    default: false
  }
});

const emit = defineEmits(['add']);

const inputValue = ref('');

// Optional: Limit input length
watch(inputValue, (newValue) => {
  if (newValue.length > 100) {
    inputValue.value = newValue.substring(0, 100);
  }
});

function handleAdd() {
  if (inputValue.value.trim() && !props.disabled) {
    emit('add', inputValue.value.trim());
    inputValue.value = '';
  }
}
</script>