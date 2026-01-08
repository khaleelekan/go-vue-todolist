<template>
  <li 
    @click="$emit('toggle', todo.id)"
    @dblclick="startEditing"
    @mouseenter="isHovered = true"
    @mouseleave="isHovered = false"
    class="group p-4 rounded-xl border-2 transition-all duration-200 cursor-pointer transform hover:-translate-y-0.5 active:scale-[0.98] flex items-center justify-between"
    :class="[
      todo.completed
        ? 'border-green-200 bg-green-50'
        : 'border-gray-200 bg-white hover:border-purple-300',
      isEditing ? 'ring-2 ring-purple-500 ring-offset-1' : ''
    ]"
  >
    <!-- Left side: Checkbox and content -->
    <div class="flex items-center flex-1 min-w-0 gap-4">
      <!-- Custom checkbox -->
      <div class="relative">
        <div 
          class="flex items-center justify-center w-6 h-6 transition-all duration-200 border-2 rounded-lg"
          :class="[
            todo.completed
              ? 'bg-gradient-to-r from-green-500 to-emerald-500 border-green-500'
              : 'border-gray-300 group-hover:border-purple-500'
          ]"
        >
          <svg 
            v-if="todo.completed"
            class="w-4 h-4 text-white"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
          >
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7" />
          </svg>
        </div>
      </div>
      
      <!-- Content area -->
      <div v-if="!isEditing" class="flex-1 min-w-0">
        <!-- Todo text -->
        <span 
          class="text-lg break-words"
          :class="[
            todo.completed
              ? 'line-through text-gray-500'
              : 'text-gray-800'
          ]"
        >
          {{ todo.title }}
        </span>
        
        <!-- Optional metadata -->
        <div v-if="todo.createdAt || todo.category" class="flex items-center gap-3 mt-2">
          <span 
            v-if="todo.createdAt"
            class="px-2 py-1 text-xs rounded-full"
            :class="todo.completed 
              ? 'bg-gray-200 text-gray-500' 
              : 'bg-purple-100 text-purple-600'"
          >
            {{ formatDate(todo.createdAt) }}
          </span>
          
          <span 
            v-if="todo.category"
            class="px-2 py-1 text-xs text-blue-600 bg-blue-100 rounded-full"
          >
            {{ todo.category }}
          </span>
        </div>
      </div>
      
      <!-- Edit input -->
      <input
        v-else
        ref="editInput"
        v-model="editText"
        @keyup.enter="saveEdit"
        @keyup.esc="cancelEdit"
        @blur="saveEdit"
        class="flex-1 text-lg text-gray-800 bg-transparent border-none outline-none"
        type="text"
        :placeholder="todo.title"
      />
    </div>
    
    <!-- Right side: Actions -->
    <div class="flex items-center gap-2">
      <!-- Edit button (shown on hover) -->
      <button
        v-if="isHovered && !isEditing"
        @click.stop="startEditing"
        class="p-2 text-gray-500 transition-colors duration-200 rounded-lg hover:text-purple-600 hover:bg-gray-100"
        title="Edit"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z" />
        </svg>
      </button>
      
      <!-- Delete button (always visible when editing, otherwise on hover) -->
      <button
        v-if="isHovered || isEditing"
        @click.stop="$emit('delete', todo.id)"
        class="p-2 text-red-500 transition-colors duration-200 rounded-lg hover:bg-red-50"
        :class="isEditing ? 'bg-red-50' : ''"
        title="Delete"
      >
        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
        </svg>
      </button>
      
      <!-- Save button (only when editing) -->
      <button
        v-if="isEditing"
        @click.stop="saveEdit"
        class="px-3 py-2 text-sm font-medium text-white transition-all duration-200 rounded-lg bg-gradient-to-r from-green-500 to-emerald-500 hover:shadow-lg active:scale-95"
      >
        Save
      </button>
      
      <!-- Cancel button (only when editing) -->
      <button
        v-if="isEditing"
        @click.stop="cancelEdit"
        class="px-3 py-2 text-sm font-medium text-gray-700 transition-colors duration-200 bg-gray-200 rounded-lg hover:bg-gray-300"
      >
        Cancel
      </button>
    </div>
  </li>
</template>

<script setup>
import { ref, nextTick } from 'vue';

const props = defineProps({
  todo: {
    type: Object,
    required: true,
    default: () => ({
      id: '',
      title: '',
      completed: false,
      createdAt: null,
      category: null
    })
  }
});

const emit = defineEmits(['toggle', 'delete', 'update']);

const isHovered = ref(false);
const isEditing = ref(false);
const editText = ref('');
const editInput = ref(null);

const startEditing = () => {
  isEditing.value = true;
  editText.value = props.todo.title;
  nextTick(() => {
    if (editInput.value) {
      editInput.value.focus();
      editInput.value.select();
    }
  });
};

const saveEdit = () => {
  if (editText.value.trim() && editText.value !== props.todo.title) {
    emit('update', { id: props.todo.id, title: editText.value.trim() });
  }
  isEditing.value = false;
};

const cancelEdit = () => {
  isEditing.value = false;
  editText.value = '';
};

const formatDate = (dateString) => {
  if (!dateString) return '';
  try {
    const date = new Date(dateString);
    const now = new Date();
    const diffMs = now - date;
    const diffDays = Math.floor(diffMs / (1000 * 60 * 60 * 24));
    
    if (diffDays === 0) {
      return 'Today';
    } else if (diffDays === 1) {
      return 'Yesterday';
    } else if (diffDays < 7) {
      return `${diffDays} days ago`;
    } else {
      return date.toLocaleDateString('en-US', { 
        month: 'short', 
        day: 'numeric'
      });
    }
  } catch {
    return '';
  }
};
</script>