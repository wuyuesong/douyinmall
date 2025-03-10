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
            :src="item.image"
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
            <div class="bottom">
              <el-button type="primary" text class="button">查看详情</el-button>
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
      tableData: [],
      currentPage: 1,
      pageSize: 12,
      total: 0
    };
  },
  mounted() {
    this.fetchData();
  },
  methods: {
    // 获取数据
    async fetchData() {
      try {
        const response = await this.$axios.get('/admin/home', {
          params: {
            page: this.currentPage,
            size: this.pageSize
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