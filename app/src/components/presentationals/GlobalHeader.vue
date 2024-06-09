<template>
  <div class="header-content">
    <div class="title">
      Todo List
    </div>
    <button-small @on-click="addTodo">追加</button-small>
    <dialog-add-todo :is-visible="isVisibleAddTodoDialog" />
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue";

import ButtonSmall from '@/components/presentationals/ButtonSmall.vue';
import DialogAddTodo from "@/components/presentationals/dialogs/DialogAddTodo.vue";

import { useTodoStore } from "@/stores/todo";
const todoStore = useTodoStore();

const isVisibleAddTodoDialog = computed(() => todoStore.isAddTodoDialog);

const addTodo = () => {
  todoStore.openAddTodoDialog();
};
</script>

<style lang="scss" scoped>
@import '@/assets/scss/style.scss';

.header-content {
  display: flex;
  justify-content: center;
  align-items: center;
  height: $header-height;
}

.title {
  display: flex;
  justify-content: center;
  font-family: 'Caveat', cursive;
  font-size: $font-xlarge;
  width: 90%;
}
</style>