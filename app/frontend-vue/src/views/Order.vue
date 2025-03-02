<template>
    <Header></Header>
  </template>

<script>
import Header from '../components/Header.vue';
export default {
  // 在 created 生命周期中调用（推荐用于数据初始化）
  
  components: {
            Header
        },
  created() {
    this.handleOrder() // 页面加载时自动执行
  },
  methods: {
    async handleOrder() {
      try {
        const response = await this.$axios.get('/user/order', {

        })

      } catch (error) {
        console.error('登录失败', error)
      }
    }
  }
}
</script>