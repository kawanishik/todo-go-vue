import { ref, computed } from 'vue';
import { defineStore } from 'pinia';
import type { Todo } from '../types/todo';

export const useTodoStore = defineStore('todo', () => {
  const todoList = ref<Todo[]>([]);
  const completedTodoList = computed(() => todoList.value.filter(todo => todo.Completed));
  const uncompletedTodoList = computed(() => todoList.value.filter(todo => !todo.Completed));

  const setTodoList = ((list: Todo[] | null)  => {
    if (list === null) return;
    todoList.value = list;
  });

  return { todoList, completedTodoList, uncompletedTodoList, setTodoList };
})
