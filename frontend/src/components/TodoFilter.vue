<template>
  <div class="flex gap-2 mb-6 pb-4 border-b-2 border-gray-100">
    <button 
      v-for="filter in filters" 
      :key="filter"
      @click="$emit('update:modelValue', filter)"
      :class="[
        'px-4 py-2 rounded-lg font-medium transition',
        modelValue === filter 
          ? 'bg-purple-600 text-white' 
          : 'bg-transparent text-gray-600 hover:bg-gray-100'
      ]"
    >
      {{ filter }}
    </button>
  </div>
</template>

<script setup>
defineProps({
  modelValue: {
    type: String,
    default: 'All'
  }
});

defineEmits(['update:modelValue']);

const filters = ['All', 'Active', 'Completed'];
</script>