<template>
    <Header />
    <div class="payment-container">
      <el-card class="payment-card">
        <h2 class="payment-title">支付订单 #{{ orderId }}</h2>
        <div class="total-amount">
          支付金额：<span class="price">￥{{ formatPrice(total) }}</span>
        </div>
  
        <el-form 
          ref="paymentFormRef" 
          :model="paymentForm" 
          :rules="paymentRules"
          label-width="120px"
        >
          <!-- 信用卡信息 -->
          <el-form-item label="信用卡号" prop="cardNumber">
            <el-input 
              v-model="paymentForm.cardNumber"
              placeholder="1234 5678 9012 3456"
              maxlength="19"
              @input="formatCardNumber"
            />
          </el-form-item>
  
          <el-row :gutter="20">
            <el-col :span="8">
              <el-form-item label="有效期月份" prop="expMonth">
                <el-select 
                  v-model="paymentForm.expMonth" 
                  placeholder="MM"
                  class="w-full"
                >
                  <el-option 
                    v-for="month in 12" 
                    :key="month"
                    :label="String(month).padStart(2, '0')"
                    :value="month"
                  />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="有效期年份" prop="expYear">
                <el-select 
                  v-model="paymentForm.expYear" 
                  placeholder="YYYY"
                  class="w-full"
                >
                  <el-option 
                    v-for="year in yearOptions" 
                    :key="year"
                    :label="year"
                    :value="year"
                  />
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="8">
              <el-form-item label="CVV" prop="cvv">
                <el-input 
                  v-model="paymentForm.cvv" 
                  placeholder="123" 
                  maxlength="4"
                  type="password"
                />
              </el-form-item>
            </el-col>
          </el-row>
  
          <div class="payment-actions">
            <el-button 
              type="primary" 
              size="large" 
              @click="handlePayment"
              :loading="processing"
            >
              立即支付
            </el-button>
            <el-button 
              size="large" 
              @click="$router.push('/')"
            >
              返回首页
            </el-button>
          </div>
        </el-form>
      </el-card>
    </div>
  </template>
  
  <script setup>
  import { ref, computed, onMounted } from 'vue'
  import { ElMessage } from 'element-plus'
  import { useRoute, useRouter } from 'vue-router'
  import Header from '@/components/Header.vue'
  import $axios from '@/utils/axios'

  const route = useRoute()
  const router = useRouter()
  const processing = ref(false)
  const orderId = route.params.orderId
  
  const total = ref(route.query.total ? parseFloat(route.query.total) : 0)
  // 支付表单数据
  const paymentForm = ref({
    cardNumber: '',
    expMonth: '',
    expYear: '',
    cvv: ''
  })

  
  // 生成年份选项（当前年份往后10年）
  const currentYear = new Date().getFullYear()
  const yearOptions = Array.from({ length: 10 }, (_, i) => currentYear + i)
  
  // 表单验证规则
  const paymentRules = {
    cardNumber: [
      { required: true, message: '请输入信用卡号', trigger: 'blur' },
      { 
        validator: (_, value, callback) => {
          const cleanValue = value.replace(/\s/g, '')
          if (!/^\d{13,19}$/.test(cleanValue)) {
            callback(new Error('请输入有效的信用卡号'))
          } else {
            callback()
          }
        },
        trigger: 'blur'
      }
    ],
    expMonth: [
      { required: true, message: '请选择有效期月份', trigger: 'change' }
    ],
    expYear: [
      { required: true, message: '请选择有效期年份', trigger: 'change' }
    ],
    cvv: [
      { required: true, message: '请输入CVV码', trigger: 'blur' },
      { 
        validator: (_, value, callback) => {
          if (!/^\d{3,4}$/.test(value)) {
            callback(new Error('CVV应为3或4位数字'))
          } else {
            callback()
          }
        },
        trigger: 'blur'
      }
    ]
  }
  
  // 格式化卡号显示
  const formatCardNumber = () => {
    paymentForm.value.cardNumber = paymentForm.value.cardNumber
      .replace(/\D/g, '')
      .replace(/(\d{4})/g, '$1 ')
      .trim()
  }
  
  // 金额格式化
  const formatPrice = (price) => {
    return Number(price).toFixed(2)
  }
  
  // 处理支付
  const handlePayment = async () => {
    try {
      processing.value = true
      
      const payload = {
        credit_card: {
          credit_card_number: paymentForm.value.cardNumber.replace(/\s/g, ''),
          credit_card_cvv: parseInt(paymentForm.value.cvv),
          credit_card_expiration_month: parseInt(paymentForm.value.expMonth),
          credit_card_expiration_year: parseInt(paymentForm.value.expYear)
        }
      }
  
      const response = await $axios.post(`/payment/${orderId}`, payload)
      
      if (response.code === 200) {
        ElMessage.success('支付成功')
        router.push({
          path: '/order-success',
          query: { orderId }
        })
      } else {
        ElMessage.error(`支付失败: ${response.message}`)
      }
    } catch (error) {
      ElMessage.error(`支付失败: ${error.response?.data?.message || error.message}`)
    } finally {
      processing.value = false
    }
  }
  
  // 获取订单金额（根据实际情况实现）
  onMounted(async () => {
  })

  </script>
  
  <style scoped>
  .payment-container {
    max-width: 800px;
    margin: 20px auto;
    padding: 0 20px;
  }
  
  .payment-card {
    padding: 30px;
  }
  
  .payment-title {
    text-align: center;
    margin-bottom: 20px;
  }
  
  .total-amount {
    text-align: right;
    font-size: 18px;
    margin-bottom: 30px;
  }
  
  .price {
    color: #e4393c;
    font-weight: bold;
    font-size: 24px;
  }
  
  .payment-actions {
    margin-top: 30px;
    text-align: center;
  }
  
  .el-select {
    width: 100%;
  }
  </style>