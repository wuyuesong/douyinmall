<template>
  <Header></Header>
  <div class="action-buttons">
    <!-- 非删除模式显示 -->
    <template v-if="!isDeleting">
      <el-button type="danger" @click="startDelete">删除商品</el-button>
      <el-button type="primary" @click="openAddDialog">添加商品</el-button>
    </template>
    
    <!-- 删除模式显示 -->
    <div v-else class="confirm-actions">
      <el-button type="danger" @click="confirmDelete">确认删除</el-button>
      <el-button @click="cancelDelete">取消删除</el-button>
    </div>
  </div>
  
  <div class="content-container">
    <el-table 
      :data="tableData" 
      style="width: 100%" 
      class="table-wrapper"
      @selection-change="handleSelectionChange">
      <!-- 动态显示选择列 -->
      <el-table-column 
        type="selection" 
        width="55" 
        v-if="showSelection">
      </el-table-column>
      
      <!-- 原有表格列 -->
      <el-table-column prop="id" label="商品id" width="200" />
      <el-table-column prop="name" label="名称" width="200" />
      <el-table-column prop="description" label="描述" width="800" />
      <el-table-column prop="price" label="价格" width="200" />
      <el-table-column prop="stock" label="库存" width="200" />
      <el-table-column fixed="right" label="编辑" min-width="200">
        <template #default="scope">
          <el-button link type="primary" size="small">Edit</el-button>
        </template>
      </el-table-column>
    </el-table>

  </div>

  <div class="pagination-container">
    <div class="demo-pagination-block">
      <el-pagination
        v-model:current-page="currentPage3"
        v-model:page-size="pageSize3"
        :small="small"
        :disabled="disabled"
        :background="background"
        layout="prev, pager, next, jumper"
        :total="tatalPage3"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>

<script>
import Header from '../../components/AdminHeader.vue';

export default {
  components: {
    Header
  },
  data() {
    return {
      // 新增状态管理
      isDeleting: false,
      showSelection: false,
      selectedIds: [],
      // 原有数据
      currentPage3: 1,
      pageSize3: 16,
      tatalPage3: 1,
      tableData: []
    };
  },
  mounted() {
    this.fetchTableData(); // 3. 组件挂载时自动获取数据
  },
  methods: {
    // 新增删除相关方法
    startDelete() {
      this.isDeleting = true;
      this.showSelection = true;
    },
    
    async confirmDelete() {
      if (this.selectedIds.length === 0) {
        this.$message.warning('请选择要删除的商品');
        return;
      }
      
      try {
        await this.$axios.post('/admin/delete', { 
          ids: this.selectedIds 
        });
        this.$message.success('删除成功');
        this.fetchTableData(); // 刷新数据
      } catch (error) {
        console.error('删除失败:', error);
      } finally {
        this.cancelDelete();
      }
    },
    
    cancelDelete() {
      this.isDeleting = false;
      this.showSelection = false;
      this.selectedIds = [];
    },
    
    // 处理选择变化
    handleSelectionChange(selection) {
      this.selectedIds = selection.map(item => item.id);
    },
    
    // 添加商品方法（示例）
    openAddDialog() {
      // 这里实现添加商品逻辑
      this.$router.push('/admin/addProduct')
      console.log('打开添加商品对话框');
    },
    
    // 原有方法保持不变
    async fetchTableData() {
      try {
        const response = await this.$axios.get('/admin/home', {
          params: { // 添加参数对象
            page: this.currentPage3, // 假设页码变量是 currentPage
            size: this.pageSize3     // 假设每页条数变量是 pageSize
          }
        });
        this.tableData = response.data.items;
        this.totalPage3 = response.data.total_page;
      } catch (error) {
        console.error('数据获取失败:', error);
      }
    }
  }
}
</script>

<style scoped>
/* 按钮容器样式 */
.action-buttons {
  padding: 0 300px;
  margin: 20px 0;
  display: flex;
  justify-content: flex-end;
  gap: 10px;
}

/* 确认操作容器样式 */
.confirm-actions {
  display: flex;
  gap: 10px;
}

/* 分页居中 */
.pagination-container {
  display: flex;
  justify-content: center;
  margin-top: 20px;
}

/* 原有表格容器样式保持不变 */
.content-container {
  padding-left: 300px;
  padding-right: 300px;
  display: flex;
  flex-direction: column;
  min-height: calc(100vh - 300px);
}

.table-wrapper {
  margin: 0 0;
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}
</style>