<template>
    <Header></Header>
    <div class="detail-container" v-if="product">
      <el-row :gutter="40" class="detail-row">
        <!-- 商品图片 -->
        <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
          <div class="image-container">
            <el-image 
              :src="getFullUrl(product.picture)"
              fit="contain"
              class="main-image"
              :preview-src-list="[getFullUrl(product.picture)]"
            >
              <template #error>
                <div class="image-error">
                  <el-icon :size="40" color="var(--el-text-color-secondary)">
                    <Picture />
                  </el-icon>
                  <span class="error-text">图片加载失败</span>
                </div>
              </template>
            </el-image>
          </div>
        </el-col>
  
        <!-- 商品信息 -->
        <el-col :xs="24" :sm="12" :md="12" :lg="12" :xl="12">
          <div class="product-info">
            <h1 class="product-title">{{ product.name }}</h1>
            <div class="price-section">
              <span class="price">￥{{ formatPrice(product.price) }}</span>
            </div>
            <div class="meta-info">
              <span class="stock">库存：{{ product.stock }}件</span>
            </div>
            
            <!-- 数量选择 -->
            <div class="quantity-selector">
              <span class="quantity-label">数量：</span>
              <el-input-number 
                v-model="quantity"
                :min="1" 
                :max="product.stock"
                :step="1"
                size="large"
                controls-position="right"
              />
            </div>
            
            <div class="actions">
              <el-button 
                type="primary" 
                size="large" 
                class="buy-btn"
                @click="handleBuy"
              >
                立即购买
              </el-button>
              <el-button 
                type="success" 
                size="large" 
                class="cart-btn"
                @click="handleAddToCart"
              >
                加入购物车
              </el-button>
            </div>
          </div>
        </el-col>
      </el-row>
  
      <!-- 商品详情 -->
      <div class="product-detail">
        <el-tabs v-model="activeTab" class="detail-tabs">
          <el-tab-pane label="商品详情" name="detail">
            <div class="detail-content" v-html="product.description"></div>
          </el-tab-pane>
        </el-tabs>
      </div>
    </div>
  
    <div v-else class="loading-container">
      <el-skeleton :rows="6" animated />
    </div>
  </template>
  
  <script>
  import Header from '../components/Header.vue';
  import { Picture } from '@element-plus/icons-vue';
  import { ElMessageBox } from 'element-plus';
  
  export default {
    components: {
      Header,
      Picture
    },
    props: {
      id: {
        type: String,
        required: true
      }
    },
    data() {
      return {
        product: null,
        activeTab: 'detail',
        baseUrl: import.meta.env.VITE_IMAGE_BASE_URL,
        quantity: 1  // 新增数量状态
      };
    },
    methods: {
      getFullUrl(relativePath) {
        // 复用列表页的图片处理方法
        if (!relativePath) return '';
        try {
          const base = this.baseUrl.endsWith('/') 
            ? this.baseUrl.slice(0, -1)
            : this.baseUrl;
          const path = relativePath.startsWith('/')
            ? relativePath.slice(1)
            : relativePath;
          return `${base}/${encodeURIComponent(path)}`;
        } catch (e) {
          console.error('URL构造错误:', e);
          return '';
        }
      },
      formatPrice(price) {
        return Number(price).toFixed(2);
      },
      async fetchProductDetail() {
        try {
          const response = await this.$axios.get(`/product/${this.id}`);
          this.product = response.data.item;
        } catch (error) {
          console.error('获取商品详情失败:', error);
          this.$message.error('商品信息加载失败');
          this.$router.replace('/');
        }
      },
      handleBuy() {
        if (this.quantity > this.product.stock) {
          this.$message.error('购买数量不能超过库存');
          return;
        }
        // 处理立即购买逻辑
        this.$message.success(`立即购买 ${this.quantity} 件商品`);
        console.log('购买数量:', this.quantity);
      },

      async handleAddToCart() {
            try {
                // 添加确认对话框
                await ElMessageBox.confirm('确定要将此商品加入购物车吗？', '操作确认', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
                });

                if (this.quantity > this.product.stock) {
                this.$message.error('添加数量不能超过库存');
                return;
                }

                console.log('添加数量:', this.quantity);
                const response = await this.$axios.post('/cart', {
                    product_id: Number(this.id.trim()),
                    Product_num: this.quantity
                });

                if (response.code === 200) {
                    this.$message.success('已成功加入购物车');
                    this.$store.commit('SET_CART_NUM', response.data.cart_num.toString())
                    localStorage.setItem('cartNum', response.data.cart_num.toString())
                } else {
                    this.$message.error(response.message || '操作失败，请重试');
                }
            } catch (error) {
                // 用户点击取消时忽略错误
                if (error !== 'cancel') {
                console.error('加入购物车失败:', error);
                const msg = error.response?.data?.message || '网络错误，请稍后重试';
                
                if (error.response?.status === 401) {
                    this.$message.error('登录已过期，请重新登录');
                    return this.$router.push('/login');
                }
                
                this.$message.error(msg);
                }
            }
        }
    },
    mounted() {
      this.fetchProductDetail();
    }
  }
  </script>
  
  <style scoped>
  /* 新增数量选择样式 */
  .quantity-selector {
    margin: 25px 0;
    display: flex;
    align-items: center;
  }

  .quantity-label {
    font-size: 16px;
    color: #666;
    margin-right: 15px;
  }

  /* 调整Element输入数字组件样式 */
  :deep(.el-input-number--large) {
    width: 150px;
  }

  :deep(.el-input-number__decrease),
  :deep(.el-input-number__increase) {
    width: 32px;
    background-color: #f5f7fa;
  }

  .detail-container {
    max-width: 1400px;
    margin: 0 auto;
    padding: 30px 20px;
  }
  
  .detail-row {
    margin-bottom: 40px;
  }
  
  .image-container {
    border: 1px solid #eee;
    border-radius: 8px;
    padding: 20px;
    background: #fff;
  }
  
  .main-image {
    width: 100%;
    height: 500px;
  }
  
  .product-info {
    padding: 20px 40px;
  }
  
  .product-title {
    font-size: 28px;
    margin-bottom: 20px;
    color: #333;
  }
  
  .price-section {
    margin: 25px 0;
    padding: 20px;
    background: #f8f8f8;
    border-radius: 8px;
  }
  
  .price {
    color: #e4393c;
    font-size: 32px;
    font-weight: 700;
    margin-right: 20px;
  }
  
  .original-price {
    color: #999;
    text-decoration: line-through;
    font-size: 16px;
  }
  
  .meta-info {
    margin: 25px 0;
    font-size: 14px;
    color: #666;
  }
  
  .meta-info span {
    margin-right: 30px;
  }
  
  .actions {
    margin-top: 40px;
  }
  
  .buy-btn {
    width: 180px;
    height: 50px;
    font-size: 18px;
  }
  
  .cart-btn {
    width: 180px;
    height: 50px;
    font-size: 18px;
    margin-left: 20px;
  }
  
  .product-detail {
    margin-top: 40px;
    background: #fff;
    padding: 20px;
    border-radius: 8px;
  }
  
  .detail-tabs {
    margin-top: 20px;
  }
  
  .detail-content {
    padding: 20px;
    line-height: 1.8;
    font-size: 16px;
    color: #666;
  }
  
  .loading-container {
    max-width: 1400px;
    margin: 0 auto;
    padding: 50px 20px;
  }
  
  @media (max-width: 768px) {
    /* 移动端样式调整 */
    .quantity-selector {
      flex-direction: column;
      align-items: flex-start;
    }

    .quantity-label {
      margin-bottom: 10px;
    }

    :deep(.el-input-number--large) {
      width: 100%;
    }

    /* 移动端按钮调整 */
    .buy-btn,
    .cart-btn {
      width: 100%;
      margin-left: 0;
      margin-top: 10px;
    }
  }
  </style>