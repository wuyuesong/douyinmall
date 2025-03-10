<template>
    
  </template>
  
  <script lang="ts" setup>
  import { ref, computed } from 'vue'
  
  // 模拟数据生成（实际项目应从接口获取）
  const mockData = () => {
    const data = []
    for(let i = 1; i <= 100; i++) {
      data.push({
        title: `Item ${i}`,
        date: new Date().toLocaleDateString(),
        id: i
      })
    }
    return data
  }
  
  // 分页参数
  const currentPage = ref(1)
  const pageSize = 20
  const total = ref(mockData().length)
  
  // 分页数据计算
  const paginatedData = computed(() => {
    const start = (currentPage.value - 1) * pageSize
    const end = start + pageSize
    return mockData().slice(start, end)
  })
  </script>
  
  <style scoped>
  .pagination-container {
    padding: 20px;
  }
  
  .card-item {
    margin-bottom: 20px;
  }
  
  .image {
    width: 100%;
    height: 200px;
    object-fit: cover;
  }
  
  .pagination-wrapper {
    margin-top: 30px;
    display: flex;
    justify-content: center;
  }
  
  /* 响应式调整 */
  @media (max-width: 768px) {
    .el-col {
      margin-bottom: 15px;
    }
    
    .image {
      height: 150px;
    }
  }
  </style>