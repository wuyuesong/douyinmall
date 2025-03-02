<template>
  <Header></Header>
  <div class="pagination-container">
    <div class="demo-pagination-block">
      <el-pagination
        v-model:current-page="currentPage3"
        v-model:page-size="pageSize3"
        :small="small"
        :disabled="disabled"
        :background="background"
        layout="prev, pager, next, jumper"
        :total="1000"
        @size-change="handleSizeChange"
        @current-change="handleCurrentChange"
      />
    </div>
  </div>
</template>

<script>
import Header from '../components/Header.vue';

export default {
  components: {
    Header
  },
  data() {
    return {
      currentPage3: 1,
      pageSize3: 10,
      small: false,
      disabled: false,
      background: false
    };
  },
  methods: {
    handleSizeChange(val) {
      console.log(`${val} items per page`);
    },
    handleCurrentChange(val) {
      console.log(`current page: ${val}`);
    }
  }
}
</script>

<style scoped>
.pagination-container {
  min-height: calc(100vh - 150px); /* 减去Header高度 */
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
  align-items: center;
  padding-bottom: 20px; /* 底部留白 */
}

.demo-pagination-block {
  margin-top: auto; /* 将分页器推到底部 */
}
</style>