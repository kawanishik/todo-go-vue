import { ref, computed } from 'vue';
import { defineStore } from 'pinia';
import type { Todo } from '../types/todo';

export const useTodoStore = defineStore('todo', () => {
  const todoList = ref<Todo[]>([]);
  const completedTodoList = computed(() => todoList.value.filter(todo => todo.Completed));
  const uncompletedTodoList = computed(() => todoList.value.filter(todo => !todo.Completed));
  const isAddTodoDialog = ref(false);

  const setTodoList = ((list: Todo[] | null)  => {
    if (list === null) return;
    todoList.value = list;
  });

  const openAddTodoDialog = () => {
    isAddTodoDialog.value = true;
  };

  const closeAddTodoDialog = () => {
    isAddTodoDialog.value = false;
  };

  return { todoList, completedTodoList, uncompletedTodoList, setTodoList, isAddTodoDialog, openAddTodoDialog, closeAddTodoDialog};
})
