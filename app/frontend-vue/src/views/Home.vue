<template>
  <Header></Header>
  <div class="content-container">
    <el-row :gutter="24" class="card-row">
      <el-col
        v-for="(item, index) in tableData"
        :key="index"
        :span="6"
        :xs="24"
        :sm="12"
        :md="6"
        :lg="6"
        :xl="6"
      >
        <el-card :body-style="{ padding: '0px' }" class="card-item">
          <el-image
            :src=getFullUrl(item.picture)
            class="image"
            fit="cover"
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
          <div style="padding: 14px">
            <span class="card-title">{{ item.name }}</span>
            <div class="product-info">
              <span class="price">￥{{ formatPrice(item.price) }}</span>
              <span class="stock">库存 {{ item.stock }} 件</span>
            </div>
            <div class="bottom">
              <el-button 
                type="primary" 
                text 
                class="button"
                @click="$router.push(`/product-detail/${item.id}`)">
                查看详情
              </el-button>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 分页组件 -->
    <div class="pagination-container">
      <div class="demo-pagination-block">
        <el-pagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :small="small"
          :disabled="disabled"
          :background="background"
          layout="prev, pager, next, jumper"
          :total="Number(total)"
          @size-change="handleSizeChange"
          @current-change="handlePageChange"
        />
      </div>
    </div>
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
  data() {
    return {
      currentType: 'normal',
      tableData: [],
      currentPage: 1,
      pageSize: 12,
      total: 0,
      baseUrl: import.meta.env.VITE_IMAGE_BASE_URL 
    };
  },
  watch: {
    // 监听路由参数变化
    '$route.query.type': {
      immediate: true,
      handler(newType) {
        if (newType === 'discount' || newType === 'normal') {
          this.currentType = newType;
          this.currentPage = 1; // 切换类型时重置页码
          this.fetchData();
        }
      }
    }
  },
  mounted() {
    this.fetchData();
  },
  methods: {

    getFullUrl(relativePath) {
      if (!relativePath) return '';
      
      // 使用URL对象处理路径
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
      return Number(price).toFixed(2)
    },
    // 获取数据
    async fetchData() {
      try {
        const categoryType = this.$route.query.type || 'normal'; // 获取类型参数
        const response = await this.$axios.get('/', {
          params: {
            page: this.currentPage,
            size: this.pageSize,
            type: categoryType // 添加类型参数到请求
          }
        });
        
        this.tableData = response.items;
        this.total = response.total_page;
      } catch (error) {
        console.error('数据获取失败:', error);
        this.$message.error('数据加载失败');
      }
    },

    // 分页变化处理
    handlePageChange(newPage) {
      this.currentPage = newPage;
      this.fetchData();
    },

    // 刷新数据
    handleRefresh() {
      this.currentPage = 1;
      this.fetchData();
    },
  }
}
</script>

<style scoped>
.active-category {
  color: var(--el-color-primary);
  font-weight: bold;
}

.product-info {
  margin: 12px 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.price {
  color: #e4393c;
  font-size: 18px;
  font-weight: 700;
}

.stock {
  color: #999;
  font-size: 12px;
}

.image-error {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 200px;
  background: var(--el-fill-color-light);
  border-radius: 8px 8px 0 0;
}

.error-text {
  margin-top: 10px;
  color: var(--el-text-color-secondary);
  font-size: 14px;
}

/* 修改原image样式 */
.image {
  width: 100%;
  height: 200px;
  border-radius: 8px 8px 0 0;
}

.content-container {
  padding: 20px;
  max-width: 1400px;
  margin: 0 auto;
  min-height: calc(100vh - 160px);
}

.card-row {
  margin-bottom: 20px;
}

.card-item {
  margin-bottom: 20px;
  transition: transform 0.3s;
  margin: 12px;
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
}

.card-item:hover {
  transform: translateY(-5px);
}

.image {
  width: 100%;
  height: 200px;
  object-fit: cover;
  border-radius: 8px 8px 0 0;
}

.card-title {
  font-size: 16px;
  font-weight: bold;
  margin-bottom: 8px;
  display: block;
}

.pagination-container {
  display: flex;
  justify-content: center;
  padding: 20px 0;
}

/* 响应式调整 */
@media (max-width: 768px) {
  .content-container {
    padding: 15px;
  }
  
  .card-item {
    margin: 8px;
  }
}
</style>