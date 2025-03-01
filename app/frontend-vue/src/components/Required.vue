<!-- components/Required.vue -->
<template>
    <span v-if="show" class="text-danger">*</span>
  </template>
  
  <script>
  export default {
    props: {
      show: {
        type: Boolean,
        default: true  // 默认显示星号
      }
    }
  }
  </script>