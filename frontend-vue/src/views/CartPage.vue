<template>
    <Header />
    <div class="cart-container">
      <el-card class="cart-card">
        <template #header>
          <div class="cart-header">
            <h2>我的购物车（{{ cartItems.length }}件）</h2>
          </div>
        </template>
  
        <div v-if="cartItems.length === 0" class="empty-cart">
          <el-empty description="购物车空空如也~">
            <el-button type="primary" @click="$router.push('/home')">去逛逛</el-button>
          </el-empty>
        </div>
  
        <div v-else>
          <div v-for="item in cartItems" :key="item.id" class="cart-item">
            <el-image
              :src="getFullUrl(item.product.picture)"
              class="item-image"
              fit="contain"
            ></el-image>
            
            <div class="item-info">
              <h3 class="item-title">{{ item.product.name }}</h3>
              <p class="item-price">单价：￥{{ formatPrice(item.product.price) }}</p>
            </div>
  
            <div class="item-actions">
              <el-input-number
                v-model="item.quantity"
                :min="1"
                :max="item.product.stock"
                @change="(val) => updateQuantity(item.id, val)"
                size="small"
              />
              <el-button
                type="danger"
                icon="Delete"
                circle
                @click="removeItem(item.id)"
              />
            </div>
  
            <div class="item-total">
              ￥{{ formatPrice(item.product.price * item.quantity) }}
            </div>
          </div>
  
          <div class="cart-summary">
            <div class="total-price">
              总计：￥{{ formatPrice(totalPrice) }}
            </div>
            <el-button 
              type="primary" 
              size="large"
              @click="checkout"
            >
              去结算（{{ cartItems.length }}件）
            </el-button>
          </div>
        </div>
      </el-card>
    </div>
  </template>
  
  <script setup>
  import { ref, onMounted, computed } from 'vue'
    import { useRouter } from 'vue-router'
    import { ElMessage, ElMessageBox } from 'element-plus'
    import Header from '@/components/Header.vue'
    import $axios from '@/utils/axios'
    import { useStore } from 'vuex' // 导入store

    const store = useStore() // 获取store实例
    const cartItems = ref([])
    const router = useRouter()

    const getFullUrl = (relativePath) => {
    // 根据实际图片路径处理，示例：
    return `${import.meta.env.VITE_IMAGE_BASE_URL}${relativePath}`
    }

    const formatPrice = (price) => {
    return Number(price).toFixed(2)
    }

    const totalPrice = computed(() => {
    return cartItems.value.reduce((sum, item) => {
        return sum + (item.product.price * item.quantity)
    }, 0)
    })

    const fetchCartItems = async () => {
    try {
        const response = await $axios.get('/cart')
        cartItems.value = response.data.items.map(item => ({
        id: item.id || Date.now(), // 使用唯一ID（示例，应根据实际数据调整）
        product: {
            name: item.Name,
            picture: item.Picture,
            price: parseFloat(item.Price),
            stock: item.Stock || 99 // 默认库存
        },
        quantity: parseInt(item.Qty)
        }))
    } catch (error) {
        ElMessage.error('获取购物车数据失败')
    }
    }

  // 更新商品数量
  const updateQuantity = async (itemId, newQuantity) => {
    try {
      await $axios.put(`/cart/${itemId}`, {
        quantity: newQuantity
      })
      ElMessage.success('数量更新成功')
    } catch (error) {
      ElMessage.error('更新数量失败')
    }
  }
  
  // 删除商品
  const removeItem = async (itemId) => {
    try {
      await ElMessageBox.confirm('确定要删除该商品吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
      
      await $axios.delete(`/cart/${itemId}`)
      cartItems.value = cartItems.value.filter(item => item.id !== itemId)
      ElMessage.success('删除成功')
      
      // 更新Header中的购物车数量
      store.commit('SET_CART_NUM', cartItems.value.length)
    } catch (error) {
      if (error !== 'cancel') {
        ElMessage.error('删除失败')
      }
    }
  }
  
  // 结算
  const checkout = () => {
    router.push('/checkout')
  }
  
  onMounted(() => {
    fetchCartItems()
  })
  </script>
  
  <style scoped>
  .cart-container {
    max-width: 1200px;
    margin: 20px auto;
    padding: 0 20px;
  }
  
  .cart-card {
    min-height: 500px;
  }
  
  .cart-item {
    display: flex;
    align-items: center;
    padding: 20px;
    border-bottom: 1px solid #eee;
  }
  
  .item-image {
    width: 100px;
    height: 100px;
    margin-right: 20px;
  }
  
  .item-info {
    flex: 1;
    margin-right: 20px;
  }
  
  .item-title {
    margin: 0;
    font-size: 16px;
  }
  
  .item-price {
    color: #666;
    margin: 5px 0;
  }
  
  .item-actions {
    display: flex;
    align-items: center;
    gap: 10px;
  }
  
  .item-total {
    width: 120px;
    text-align: right;
    font-weight: bold;
    color: #e4393c;
  }
  
  .cart-summary {
    margin-top: 20px;
    padding: 20px;
    background: #f8f8f8;
    border-radius: 8px;
    display: flex;
    justify-content: flex-end;
    align-items: center;
    gap: 30px;
  }
  
  .total-price {
    font-size: 18px;
    color: #e4393c;
    font-weight: bold;
  }
  
  .empty-cart {
    padding: 50px 0;
    text-align: center;
  }
  
  @media (max-width: 768px) {
    .cart-item {
      flex-wrap: wrap;
      gap: 15px;
    }
    
    .item-info {
      order: 1;
      flex: 0 0 100%;
    }
  }
  </style>