<template>
    <el-menu
      :default-active="activeIndex"
      class="el-menu-demo"
      mode="horizontal"
      :ellipsis="false"
      @select="handleSelect"
      style="display: flex; align-items: center; height: 70px;"
    >
      <el-menu-item index="0" style="width: 100px; height: 50px;  margin-left: 300px;" @click="gotoHome">
        <img
          style="width: 100px;;"
          src="../static/image/529bdbf8f6b8d51fb3aaa89d56970ba4.png"
          alt="Element logo"
        />
      </el-menu-item>
      <div style="flex: 1; display: flex;">
        <el-menu-item index="1">关于</el-menu-item>
        <el-sub-menu index="2">
          <template #title>分类</template>
          <el-menu-item index="2-1">item one</el-menu-item>
          <el-menu-item index="2-2">item two</el-menu-item>
          <el-menu-item index="2-3">item three</el-menu-item>
          <el-sub-menu index="2-4">
            <template #title>item four</template>
            <el-menu-item index="2-4-1">item one</el-menu-item>
            <el-menu-item index="2-4-2">item two</el-menu-item>
            <el-menu-item index="2-4-3">item three</el-menu-item>
          </el-sub-menu>
        </el-sub-menu>
      </div>
      <div style="margin-right: 0px">
        <el-input
          v-model="input2"
          style="width: 240px; height: 35px; margin-right: 20px;"
          placeholder="搜索商品"
          :prefix-icon="Search"
        />
      </div>
      <el-icon :size="25" style="margin-right: 20px"><ShoppingCart /></el-icon>
      <div v-if="hasToken">
        <el-dropdown trigger="click" placement="bottom-end" style="margin-right: 300px" @command="handleCommand">
          <el-avatar style="margin-right: 100px"
            src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
          />
          <template #dropdown>
            <el-dropdown-menu>
              <el-dropdown-item command="orders">订单中心</el-dropdown-item>
              <el-dropdown-item command="logout" divided>退出登录</el-dropdown-item>
            </el-dropdown-menu>
          </template>
        </el-dropdown>
      </div>
      <div v-else>
        <el-button type="primary" style="margin-right: 300px" @click="gotoLogin">登录</el-button>
      </div>
    </el-menu>
  </template>
  
  <script lang="ts" setup>
  import { ref, onMounted} from 'vue'
  import Cookies from 'js-cookie'
  import { useRouter } from 'vue-router'
  
  const activeIndex = ref('1')
  const handleSelect = (key: string, keyPath: string[]) => {
    console.log(key, keyPath)
  }
  
  import {Search} from '@element-plus/icons-vue'
  const input2 = ref('')
  
  const hasToken = ref(!!localStorage.getItem('token'))
  
  const router = useRouter() // 使用 useRouter 获取路由实例
  
  const gotoLogin = () => {
    // 获取当前路由的完整路径
    const currentPath = router.currentRoute.value.fullPath;
    console.log(currentPath);
    // 使用 router.push 跳转到登录页面，并将当前路径作为 redirect 参数
    router.push({ path: '/login', query: { redirect: currentPath } });
}

  const gotoHome = () => {
    router.push('/home') // 使用 router.push 跳转到登录页面
  }
  
  onMounted(() => {
          hasToken.value = !!localStorage.getItem('token')
  });

  const handleCommand = (command: string | number | object) => {
  switch (command) {
    case 'orders':
      router.push('/user/orders')
      break
    case 'logout':
      handleLogout()
      break
  }
}

// 新增退出登录处理
const handleLogout = () => {
    localStorage.removeItem('token')
    router.push('/') // 使用已定义的路由实例
    hasToken.value = false
}
  </script>
  
  <!-- <style scoped>
  .el-menu--horizontal > .el-menu-item:nth-child(1) {
    margin-right: auto;
  }
  </style> -->
  