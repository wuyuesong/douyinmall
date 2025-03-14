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
            
            <div class="actions">
              <el-button type="primary" size="large" class="buy-btn">
                立即购买
              </el-button>
              <el-button type="success" size="large" class="cart-btn">
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
          <!-- <el-tab-pane label="规格参数" name="spec">
            <el-table :data="product.specifications" border>
              <el-table-column prop="name" label="参数名称" width="180" />
              <el-table-column prop="value" label="参数值" />
            </el-table>
          </el-tab-pane> -->
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
        baseUrl: import.meta.env.VITE_IMAGE_BASE_URL
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
          this.product = response.item;
        } catch (error) {
          console.error('获取商品详情失败:', error);
          this.$message.error('商品信息加载失败');
          this.$router.replace('/');
        }
      }
    },
    mounted() {
      this.fetchProductDetail();
    }
  }
  </script>
  
  <style scoped>
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
    .main-image {
      height: 300px;
    }
  
    .product-info {
      padding: 20px;
    }
  
    .product-title {
      font-size: 22px;
    }
  
    .price {
      font-size: 24px;
    }
  
    .buy-btn,
    .cart-btn {
      width: 100%;
      margin-left: 0;
      margin-top: 10px;
    }
  }
  </style>