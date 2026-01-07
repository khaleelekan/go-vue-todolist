import { ref, computed } from 'vue';
import { todoAPI } from '../services/api';

export function useTodos() {
  const todos = ref([]);
  const loading = ref(false);
  const error = ref(null);

  const activeTodos = computed(() => 
    todos.value.filter(t => !t.completed).length
  );

  const completedTodos = computed(() => 
    todos.value.filter(t => t.completed).length
  );

  async function fetchTodos() {
    loading.value = true;
    error.value = null;
    try {
      const data = await todoAPI.getAll();
      todos.value = data || [];
    } catch (err) {
      error.value = 'Failed to load todos. Make sure the Go server is running.';
      console.error(err);
    } finally {
      loading.value = false;
    }
  }

  async function addTodo(title) {
    if (!title.trim()) return;
    
    error.value = null;
    try {
      const newTodo = await todoAPI.create(title.trim());
      todos.value.push(newTodo);
    } catch (err) {
      error.value = 'Failed to add todo';
      console.error(err);
    }
  }

  async function updateTodo(todo) {
    error.value = null;
    try {
      await todoAPI.update(todo.id, todo);
    } catch (err) {
      error.value = 'Failed to update todo';
      console.error(err);
      // Revert the change
      todo.completed = !todo.completed;
    }
  }

  async function deleteTodo(id) {
    error.value = null;
    try {
      await todoAPI.delete(id);
      todos.value = todos.value.filter(t => t.id !== id);
    } catch (err) {
      error.value = 'Failed to delete todo';
      console.error(err);
    }
  }

  return {
    todos,
    loading,
    error,
    activeTodos,
    completedTodos,
    fetchTodos,
    addTodo,
    updateTodo,
    deleteTodo
  };
}