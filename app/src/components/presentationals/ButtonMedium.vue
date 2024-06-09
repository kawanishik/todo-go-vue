<template>
    <div class="btn" @click="onClick">
      <slot />
    </div>
  </template>
  
  <script setup lang="ts">
  interface Emits {
    (e: 'on-click'): void;
  }
  
  const emits = defineEmits<Emits>();
  
  const onClick = () => {
    emits('on-click');
  };
  </script>
  
  <style lang="scss" scoped>
  .btn {
    display: block;
    color: #fff;
    background-color: #f7dd37;
    border-radius: 30px;
    padding: 4px 8px;
    width: 140px;
    text-align: center;
  }
  
  .btn:hover {
    color: #fff;
    background: #f3d620;
  }
  </style>
  