<template>
    <div class="d-flex justify-content-center align-items-center">
      <!-- 卡片列表 -->
      <el-row :gutter="20">
        <el-col
          v-for="(item, index) in paginatedData"
          :key="index"
          :span="6"
          :xs="24"
          :sm="12"
          :md="8"
          :lg="6"
          :xl="4"
        >
          <el-card :body-style="{ padding: '0px' }" class="card-item">
            <img
              src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png"
              class="image"
            />
            <div style="padding: 14px">
              <span>{{ item.title }}</span>
              <div class="bottom">
                <time class="time">{{ item.date }}</time>
                <el-button type="primary" text class="button">操作</el-button>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>
  
      <!-- 分页组件 -->
      <div class="pagination-wrapper">
        <el-pagination
          background
          layout="prev, pager, next"
          :total="total"
          :page-size="pageSize"
          v-model:current-page="currentPage"
        />
      </div>
    </div>
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