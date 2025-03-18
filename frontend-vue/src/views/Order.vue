<template>
  <Header></Header>
  <div class="content-container">
    <el-row v-if="orders.length === 0" class="empty-state">
      <el-empty description="暂无订单" />
    </el-row>
    <el-row v-else :gutter="24" class="card-row">
      <el-col
        v-for="(order, index) in orders"
        :key="index"
        :span="24"
      >
        <el-card class="order-card" shadow="hover">
          <div class="order-header">
            <div class="order-info">
              <span class="order-id">订单号：{{ order.OrderId }}</span>
              <span class="order-time">下单时间：{{ formatTime(order.CreatedDate) }}</span>
            </div>
            <div class="order-status">
              <span :class="statusClass(order.Status)">
                {{ getStatusText(order.Status) }}
              </span>
              <el-button 
                v-if="order.Status === 0" 
                type="primary" 
                size="small"
                @click="goToPayment(order.OrderId)"
              >
                立即支付
              </el-button>
            </div>
          </div>
          
          <el-row :gutter="20" class="order-content">
            <el-col :span="24">
              <div class="items-section">
                <h4>商品清单</h4>
                <div 
                  v-for="(item, idx) in order.Items"
                  :key="idx"
                  class="order-item"
                >
                  <span class="product-name">商品名称：{{ item.ProductName }}</span>
                  <span class="quantity">数量：{{ item.Qty }}</span>
                  <span class="cost">金额：¥{{ item.Cost.toFixed(2) }}</span>
                </div>
                <div class="order-total">
                  订单总额：¥{{ order.Cost.toFixed(2) }}
                </div>
              </div>
            </el-col>
          </el-row>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import Header from '../components/Header.vue';
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { ElMessage, ElEmpty } from 'element-plus';
import $axios from '@/utils/axios'

const orders = ref([]);
const router = useRouter();

// 状态文本映射
const statusTextMap = {
  0: { text: '待支付', class: 'status-pending' },
  1: { text: '支付成功', class: 'status-success' },
  2: { text: '已取消', class: 'status-cancelled' }
}

// 获取状态文本
const getStatusText = (status) => {
  return statusTextMap[status]?.text || '未知状态'
}

// 获取状态样式类
const statusClass = (status) => {
  return statusTextMap[status]?.class || ''
}

// 跳转到支付页面
const goToPayment = (orderId) => {
  router.push(`/payment/${orderId}`)
}
const formatTime = (dateString) => {
  const date = new Date(dateString);
  return date.toLocaleString();
};

const fetchOrders = async () => {
  try {
    const response = await $axios.get('/user/order');
    orders.value = response.data.orders;
  } catch (error) {
    console.error('获取订单失败', error);
    ElMessage.error('获取订单失败');
  }
};

onMounted(() => {
  if (!localStorage.getItem('token')) {
    router.push('/login');
  }
  fetchOrders();
});
</script>

<style scoped>
.order-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-bottom: 15px;
  margin-bottom: 15px;
  border-bottom: 1px solid #eee;
}

.order-status {
  display: flex;
  align-items: center;
  gap: 15px;
}

.status-pending {
  color: #e6a23c;
  font-weight: 500;
}

.status-success {
  color: #67c23a;
  font-weight: 500;
}

.status-cancelled {
  color: #909399;
  font-weight: 500;
}

.content-container {
  padding: 24px;
  max-width: 1200px;
  margin: 20px auto 0;
}

.empty-state {
  margin-top: 100px;
}

.order-card {
  margin-bottom: 20px;
  border-radius: 8px;
  transition: transform 0.3s;
}

.order-card:hover {
  transform: translateY(-3px);
}

.order-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-bottom: 15px;
  margin-bottom: 15px;
  border-bottom: 1px solid #eee;
}

.order-id {
  font-size: 16px;
  font-weight: bold;
  color: var(--el-color-primary);
}

.order-time {
  color: #666;
  font-size: 14px;
}

.order-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 0;
  border-bottom: 1px solid #f5f5f5;
}

.quantity {
  width: 30%;
  text-align: center;
}

.cost {
  width: 30%;
  text-align: right;
  color: #e4393c;
}

.order-total {
  text-align: right;
  padding: 15px 0;
  font-size: 16px;
  font-weight: bold;
  color: #e4393c;
  border-top: 1px solid #eee;
  margin-top: 15px;
}

/* 响应式调整 */
@media (max-width: 768px) {
  .content-container {
    padding: 16px;
  }
  
  .order-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
  }
}
</style>