<template>
  <Header />
  <div class="checkout-container">
    <el-card class="checkout-card">
      <!-- 地址填写模块 -->
      <div class="address-section">
        <h3>收货地址</h3>
        <el-form 
          ref="formRef" 
          :model="formData" 
          :rules="rules" 
          label-width="100px"
          label-position="left"
        >
          <el-form-item label="街道地址" prop="street_address">
            <el-input
              v-model="formData.street_address"
              type="textarea"
              :rows="2"
              placeholder="例如：滨海街道新湖小区17幢"
            />
          </el-form-item>

          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="城市" prop="city">
                <el-input v-model="formData.city" placeholder="例如：杭州" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="州/省" prop="state">
                <el-input v-model="formData.state" placeholder="例如：浙江" />
              </el-form-item>
            </el-col>
          </el-row>

          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="国家" prop="country">
                <el-input v-model="formData.country" placeholder="例如：中国" />
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="邮编" prop="zip_code">
                <el-input 
                  v-model="formData.zip_code" 
                  placeholder="例如：10001"
                  maxlength="10"
                />
              </el-form-item>
            </el-col>
          </el-row>
        </el-form>
      </div>

      <!-- 商品信息模块 -->
      <div class="goods-section">
        <h3>订单商品</h3>
        <div v-for="item in cartItems" :key="item.id" class="goods-item">
          <el-image 
            :src="getFullUrl(item.product.picture)" 
            class="goods-image"
            fit="contain"
          />
          <div class="goods-info">
            <p class="title">{{ item.product.name }}</p>
            <p class="specs">数量：{{ item.quantity }}</p>
            <p class="price">单价：￥{{ formatPrice(item.product.price) }}</p>
          </div>
          <div class="goods-total">
            ￥{{ formatPrice(item.product.price * item.quantity) }}
          </div>
        </div>
      </div>

      <!-- 订单确认栏 -->
      <div class="order-summary">
        <div class="total">
          实付款：<span class="price">￥{{ formatPrice(totalPrice) }}</span>
        </div>
        <el-button 
          type="primary" 
          size="large" 
          @click="submitOrder"
          :loading="submitting"
        >
          提交订单
        </el-button>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ElMessage } from 'element-plus'
import { ref, computed, onMounted } from 'vue'  
import { useRouter } from 'vue-router'
import Header from '@/components/Header.vue'
import $axios from '@/utils/axios'
import { useStore } from 'vuex';
const store = useStore();

const router = useRouter()
const formRef = ref()
const submitting = ref(false)
const paymentMethod = ref('alipay')
const cartItems = ref([])

const formData = ref({
  street_address: '',
  city: '',
  state: '',
  country: '',
  zip_code: ''
})

const rules = {
  street_address: [
    { required: true, message: '请输入街道地址', trigger: 'blur' },
    { min: 5, message: '地址至少5个字符', trigger: 'blur' }
  ],
  city: [
    { required: true, message: '请输入城市名称', trigger: 'blur' }
  ],
  state: [
    { required: true, message: '请输入州/省名称', trigger: 'blur' }
  ],
  country: [
    { required: true, message: '请输入国家名称', trigger: 'blur' }
  ],
  zip_code: [
    { required: true, message: '请输入邮政编码', trigger: 'blur' },
    { pattern: /^\d{3,10}$/, message: '请输入3-10位数字邮编', trigger: 'blur' }
  ]
}



const totalPrice = computed(() => {
  return cartItems.value.reduce((sum, item) => {
    return sum + (item.product.price * item.quantity)
  }, 0)
})

const getFullUrl = (relativePath) => {
  return `${import.meta.env.VITE_IMAGE_BASE_URL}${relativePath}`
}

const formatPrice = (price) => {
  return Number(price).toFixed(2)
}

const fetchCartItems = async () => {
  try {
    const response = await $axios.get('/cart')
    cartItems.value = response.data.items.map(item => ({
      id: item.id,
      product: {
        name: item.Name,
        picture: item.Picture,
        price: parseFloat(item.Price),
        stock: item.Stock
      },
      quantity: parseInt(item.Qty)
    }))
  } catch (error) {
    ElMessage.error('获取购物车数据失败')
  }
}

const submitOrder = async () => {
  try {
    await formRef.value.validate()
    
    submitting.value = true
    // 构造符合后端要求的请求体
    const requestData = {
      street: formData.value.street_address,   // 前端字段 -> street_address
      city: formData.value.city,                // 字段名称一致
      province: formData.value.state,           // 前端字段 -> state
      country: formData.value.country,          // 注意拼写一致性（country）
      zipcode: formData.value.zip_code          // 前端字段 -> zip_code
    }
    
    // 发送到正确的端点
    const response = await $axios.post('/checkout', requestData)

    if (response.code == 200) {
      ElMessage.success('订单创建成功')
      router.push({
        path: `/payment/${response.data.orderId}`,
        query: { total: totalPrice.value } // 添加金额查询参数
      })
      store.commit('SET_CART_NUM', 0)
      localStorage.setItem('cartNum', 0)
    } else {
      ElMessage.error('订单创建失败: ' + response.message)
    }


  } catch (error) {
    if (!error.response) {
      console.log(error)
      ElMessage.warning('请正确填写地址信息')
    } else {
      ElMessage.error('订单创建失败: ' + error.response.data.message)
    }
  } finally {
    submitting.value = false
  }
}

onMounted(async () => {
  await fetchCartItems()
  if (cartItems.value.length === 0) {
    ElMessage.warning('购物车为空，请先添加商品')
    router.push('/cart')
  }
})
</script>

<style scoped>
.checkout-container {
  max-width: 1200px;
  margin: 20px auto;
  padding: 0 20px;
}

.address-section {
  margin-bottom: 30px;
  padding: 20px;
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0,0,0,0.1);
}

.goods-item {
  display: flex;
  align-items: center;
  padding: 20px;
  border-bottom: 1px solid #eee;
}

.goods-image {
  width: 80px;
  height: 80px;
  margin-right: 20px;
}

.goods-info {
  flex: 1;
  margin-right: 20px;
}

.goods-info .title {
  font-size: 16px;
  margin: 0 0 8px 0;
}

.goods-info .specs {
  color: #666;
  margin: 4px 0;
}

.goods-info .price {
  color: #e4393c;
  font-weight: bold;
}

.goods-total {
  width: 120px;
  text-align: right;
  font-weight: bold;
  color: #e4393c;
}

.order-summary {
  margin-top: 30px;
  padding-top: 20px;
  border-top: 1px solid #eee;
  text-align: right;
}

.el-input, .el-textarea {
  width: 100%;
}
</style>