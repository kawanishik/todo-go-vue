<template>
  <div class="todo-app">
    <global-header />
  </div>
</template>

<script setup lang="ts">
import GlobalHeader from '@/components/presentationals/GlobalHeader.vue';

import type { Todo } from "@/types/todo";

import TodoRepository from "@/repositories/todo";
import { useTodoStore } from "@/stores/todo";

const todoStore = useTodoStore();
TodoRepository.getTodoList().then((res: Todo[] | null) => todoStore.setTodoList(res));
</script>

<style lang="scss" scoped>
@import '@/assets/scss/color.scss';

.todo-app {
  height: 100vh;
  width: 100vw;
  background-color: $color-bg;
}
</style>
