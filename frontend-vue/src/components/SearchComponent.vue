<template>
    <div>
      <!-- 搜索框 -->
      <input
        type="text"
        v-model="searchQuery"
        placeholder="输入搜索内容..."
        @keyup.enter="handleSearch"
      />
      <!-- 搜索按钮 -->
      <button @click="handleSearch">搜索</button>
  
      <!-- 搜索结果 -->
      <div v-if="searchResults.length > 0">
        <ul>
          <li v-for="(result, index) in searchResults" :key="index">
            {{ result }}
          </li>
        </ul>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        searchQuery: '', // 绑定到输入框的搜索词
        searchResults: [], // 存储搜索结果
        // 假设这是我们要搜索的数据列表
        dataList: ['Apple', 'Banana', 'Cherry', 'Date', 'Fig', 'Grape']
      };
    },
    methods: {
      handleSearch() {
        // 这里实现搜索逻辑，比如过滤数据列表
        if (this.searchQuery.trim() === '') {
          this.searchResults = [];
        } else {
          this.searchResults = this.dataList.filter(item =>
            item.toLowerCase().includes(this.searchQuery.toLowerCase())
          );
        }
      }
    }
  };
  </script>
  
  <style>
  /* 这里可以添加一些样式 */
  input[type="text"] {
    margin-right: 10px;
  }
  </style>
  