<template>
  <div class="todo-card">
    <div class="header">
      <div class="title">{{ props.todo.Title }}</div>
      <button-small @on-click="editTodo" class="edit-btn">編集</button-small>
    </div>
    <div v-html="props.todo.Description" class="description"></div>
    <dialog-edit-todo :is-visible="isEditDialog" :todo="props.todo" @close="closeEditDialog" />
  </div>
</template>

<script setup lang="ts">
import { ref }from 'vue';
import type { Todo } from "@/types/todo";

import ButtonSmall from '@/components/presentationals/ButtonSmall.vue';
import DialogEditTodo from "@/components/presentationals/dialogs/DialogEditTodo.vue";

const props = defineProps<{
  todo: Todo;
}>();

const isEditDialog = ref(false);

const editTodo = () => {
  isEditDialog.value = true;
};

const closeEditDialog = () => {
  isEditDialog.value = false;
};
</script>

<style lang="scss" scoped>
@import '@/assets/scss/color.scss';
@import '@/assets/scss/style.scss';

.todo-card {
  width: 300px;
  height: 300px;
  text-align: center;
  border: solid black 1px;

  .header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    font-size: $font-large;
    height: $font-large + 8px;
    background-color: $color-list-title;

    .title {
      flex: 1;
      text-align: center;
      font-family: 'Caveat', cursive;
    }

    .edit-btn {
      width: 48px;
    }
  }

  .description {
    margin-top: 8px;
    overflow-y: scroll;
  }
}
</style>
