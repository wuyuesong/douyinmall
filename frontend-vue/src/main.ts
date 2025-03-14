// import './assets/main.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

import router from './router/index'

import { createApp } from 'vue'
import App from './App.vue'

import axiosInstance from './utils/axios'



const app = createApp(App)
app.config.globalProperties.$axios = axiosInstance 

// import { ElButton } from 'element-plus'
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
  }
app.use(ElementPlus).use(router).mount('#app')