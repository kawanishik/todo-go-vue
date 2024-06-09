<template>
  <div class="todo-card-list">
    <todo-card v-for="todo in todoStore.todoList" :key="todo.id" :todo="todo" class="unit" />
  </div>
</template>

<script setup lang="ts">
import TodoCard from '@/components/presentationals/TodoCard.vue';

import type { Todo } from "@/types/todo";

import TodoRepository from "@/repositories/todo";
import { useTodoStore } from "@/stores/todo";

const todoStore = useTodoStore();
TodoRepository.getTodoList().then((res: Todo[] | null) => todoStore.setTodoList(res));
</script>

<style lang="scss" scoped>
.todo-card-list {
  display: flex;
  flex-wrap: wrap;

  .unit {
    margin: 16px;
  }
}
</style>
