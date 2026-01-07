<template>
  <div class="flex gap-3 mb-6">
    <input 
      v-model="inputValue" 
      @keyup.enter="handleAdd"
      type="text" 
      placeholder="What needs to be done?"
      :disabled="disabled"
      class="flex-1 px-4 py-3 border-2 border-gray-200 rounded-lg focus:outline-none focus:border-purple-500 transition disabled:bg-gray-100 disabled:cursor-not-allowed"
    >
    <button 
      @click="handleAdd" 
      :disabled="!inputValue.trim() || disabled"
      class="px-6 py-3 bg-purple-600 text-white font-semibold rounded-lg hover:bg-purple-700 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:ring-offset-2 disabled:bg-gray-300 disabled:cursor-not-allowed transition transform hover:scale-105 active:scale-95"
    >
      Add
    </button>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const props = defineProps({
  disabled: {
    type: Boolean,
    default: false
  }
});

const emit = defineEmits(['add']);

const inputValue = ref('');

function handleAdd() {
  if (inputValue.value.trim()) {
    emit('add', inputValue.value);
    inputValue.value = '';
  }
}
</script>