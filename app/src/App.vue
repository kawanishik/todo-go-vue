<template>
  <div class="todo-app">
    <global-header />
    <div class="todo-list">
      <todo-card-list  />
    </div>
  </div>
</template>

<script setup lang="ts">
import GlobalHeader from '@/components/presentationals/GlobalHeader.vue';
import TodoCardList from '@/components/presentationals/TodoCardList.vue';
</script>

<style lang="scss" scoped>
@import '@/assets/scss/color.scss';
@import '@/assets/scss/style.scss';

.todo-app {
  height: 100vh;
  width: 100vw;
  background-color: $color-bg;

  .todo-list {
    height: calc(100% - $header-height);
    overflow-y: scroll;
  }
}
</style>
