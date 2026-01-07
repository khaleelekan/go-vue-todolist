<template>
  <li class="flex items-center gap-3 p-4 bg-gray-50 rounded-lg hover:bg-gray-100 transition transform hover:translate-x-1">
    <input 
      type="checkbox" 
      :checked="todo.completed"
      @change="$emit('toggle', todo)"
      class="w-5 h-5 text-purple-600 border-gray-300 rounded focus:ring-purple-500 cursor-pointer"
    >
    <span 
      :class="[
        'flex-1 text-lg transition',
        todo.completed 
          ? 'line-through text-gray-400' 
          : 'text-gray-800'
      ]"
    >
      {{ todo.title }}
    </span>
    <button 
      @click="$emit('delete', todo.id)"
      class="px-4 py-2 bg-red-500 text-white text-sm font-medium rounded-lg hover:bg-red-600 transition transform hover:scale-105 active:scale-95"
    >
      Delete
    </button>
  </li>
</template>

<script setup>
defineProps({
  todo: {
    type: Object,
    required: true
  }
});

defineEmits(['toggle', 'delete']);
</script>