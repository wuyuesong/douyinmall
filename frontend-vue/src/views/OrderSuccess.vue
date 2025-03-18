<template>
    <Header />
    <div class="success-container">
      <el-card class="success-card">
        <div class="success-icon">
          <el-icon color="#67C23A" :size="80">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1024 1024">
              <path fill="currentColor" d="M512 64a448 448 0 1 1 0 896A448 448 0 0 1 512 64m-55.808 536.384-99.52-99.584a38.4 38.4 0 0 0-54.336 54.336l126.72 126.72a38.272 38.272 0 0 0 54.336 0l275.84-275.84a38.4 38.4 0 0 0-54.272-54.336L456.192 600.384z"/>
            </svg>
          </el-icon>
        </div>
        <h1 class="success-title">订单 #{{ orderId }} 支付成功！</h1>
        <div class="success-actions">
          <el-button 
            type="primary" 
            size="large" 
            @click="goToHome"
          >
            返回首页
          </el-button>
          <el-button 
            size="large" 
            @click="goToOrders"
          >
            查看订单
          </el-button>
        </div>
      </el-card>
    </div>
  </template>
  
  <script setup>
  import { useRoute, useRouter } from 'vue-router'
  import Header from '@/components/Header.vue'
  
  const route = useRoute()
  const router = useRouter()
  const orderId = route.query.orderId
  
  const goToHome = () => {
    router.push('/')
  }
  
  const goToOrders = () => {
    router.push('/orders')
  }
  </script>
  
  <style scoped>
  .success-container {
    max-width: 800px;
    margin: 50px auto;
    padding: 0 20px;
  }
  
  .success-card {
    padding: 40px 30px;
    text-align: center;
  }
  
  .success-icon {
    margin-bottom: 30px;
  }
  
  .success-title {
    color: #333;
    margin-bottom: 40px;
    font-size: 24px;
  }
  
  .success-actions {
    margin-top: 30px;
    display: flex;
    justify-content: center;
    gap: 20px;
  }
  
  .el-button {
    width: 150px;
  }
  </style>