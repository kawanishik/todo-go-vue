<template>
  <teleport to="#modal-distination">
    <div v-if="isVisible" class="dialog-add-todo">
      <div class="dialog-contents">
        <div class="header-title">Add Todo</div>
        <div class="form-title">
          <div class="title">タイトルを入力してください</div>
          <input type="text" placeholder="Title" v-model="formData.title" class="input-title" />
          <div v-if="v$.title.$invalid && formData.title.length" class="error">入力文字数を超えています</div>
        </div>
        <div class="form-description">
          <div class="description" @click="changeMarkdown">詳細を入力してください(MarkDown形式)</div>
          <textarea v-if="!isMarkdown" placeholder="Description" v-model="formData.description" class="input-description" />
          <div v-if="isMarkdown" v-html="formData.description" class="markdown"></div>
        </div>
      </div>
      <div class="btns">
        <button-medium @on-click="onClose" class="close-btn">Close</button-medium>
        <button-medium @on-click="onAddClick">追加</button-medium>
      </div>
    </div>
  </teleport>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useVuelidate } from '@vuelidate/core';
import { required, maxLength } from '@vuelidate/validators';
import TodoRepository from "@/repositories/todo";
import type { Todo } from "@/types/todo";

import ButtonMedium from '@/components/presentationals/ButtonMedium.vue';

import { useTodoStore } from "@/stores/todo";
const todoStore = useTodoStore();

withDefaults(defineProps<{
  isVisible: boolean;
}>(), {
  isVisible: false,
});

const formData = ref({
  title: '',
  description: '',
});
const isMarkdown = ref(false);

const rules = {
  title: {
    required,
    maxLength: maxLength(255),
  },
};
const v$ = useVuelidate(rules, formData);

const changeMarkdown = () => {
  isMarkdown.value = !isMarkdown.value;
};
const onClose = () => {
  todoStore.closeAddTodoDialog();
};
const onAddClick = async () => {
  TodoRepository.addTodo({
    Title: formData.value.title,
    Description: formData.value.description,
  }).then(() => {
    TodoRepository.getTodoList().then((res: Todo[] | null) => {
      todoStore.setTodoList(res);
      todoStore.closeAddTodoDialog();
    });
  });
};
</script>

<style lang="scss" scoped>
@import '@/assets/scss/style.scss';

$padding-up-down: 80px;
$padding-left-right: 120px;

.dialog-add-todo {
  position: fixed;
  top: 0;
  left: 0;
  height: calc(100% - $padding-up-down * 2);
  width: calc(100% - $padding-left-right * 2);
  background-color: rgba(59, 58, 58, 0.5);
  padding: $padding-up-down $padding-left-right;

  .dialog-contents {
    background-color: #fff;
    border-radius: 10px;
    padding: 20px;

    .header-title {
      font-size: $font-large;
      font-weight: bold;
      margin-bottom: 20px;
      text-align: center;
    }

    .form-title {
      margin: 24px;

      .title {
        font-size: $font-large;
        font-weight: bold;
        margin-bottom: 8px;
      }

      .input-title {
        width: 100%;
        height: 40px;
        font-size: $font-medium;
      }

      .error {
        color: red;
        font-size: $font-small;
      }
    }

    .form-description {
      margin: 24px;

      .description {
        font-size: $font-large;
        font-weight: bold;
        margin-bottom: 8px;
      }

      .input-description {
        resize: none;
        width: 100%;
        height: 240px;
      }

      .markdown {
        height: 240px;
        overflow-y: scroll;
      }
    }
  }

  .btns {
    display: flex;
    justify-content: center;
    margin-top: 20px;

    .close-btn {
      margin-right: 20px;
    }
  }
}
</style>
