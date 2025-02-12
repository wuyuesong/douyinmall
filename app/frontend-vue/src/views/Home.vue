<template>
  <el-menu
    :default-active="activeIndex"
    class="el-menu-demo"
    mode="horizontal"
    :ellipsis="false"
    @select="handleSelect"
    style="display: flex; align-items: center;"
  >
    <el-menu-item index="0" style="width: 100px; height: 50px;  margin-left: 100px;">
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
    <div style="margin-right: 10px;">
      <el-input
        v-model="input2"
        style="width: 240px; height: 35px; "
        placeholder="搜索商品"
        :prefix-icon="Search"
      />
    </div>
  </el-menu>
</template>

<script lang="ts" setup>
import { ref } from 'vue'

const activeIndex = ref('1')
const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
}

import {Search} from '@element-plus/icons-vue'
const input2 = ref('')
</script>

<!-- <style scoped>
.el-menu--horizontal > .el-menu-item:nth-child(1) {
  margin-right: auto;
}
</style> -->
